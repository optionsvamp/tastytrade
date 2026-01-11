package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/optionsvamp/tastytrade"
)

func main() {
	// Parse command-line flags
	symbol := flag.String("symbol", "SPX", "Underlying symbol to fetch options for (e.g., SPX, AAPL, TSLA)")
	batchSize := flag.Int("batch-size", 100, "Number of symbols to request per API call (default: 100, max: 100)")
	delayMs := flag.Int("delay-ms", 500, "Delay in milliseconds between batch requests to avoid rate limiting (default: 500ms)")
	startDate := flag.String("start-date", "", "Filter options by expiration date (start) in YYYY-MM-DD format (inclusive). If not set, no start filter is applied.")
	endDate := flag.String("end-date", "", "Filter options by expiration date (end) in YYYY-MM-DD format (inclusive). If not set, no end filter is applied.")
	perDay := flag.Bool("per-day", false, "Create separate CSV file for each expiration date (default: false, creates single combined file)")
	flag.Parse()

	// Cap batch size at 100 (server limit)
	if *batchSize > 100 {
		log.Printf("Warning: Batch size capped at 100 (server limit). Requested: %d\n", *batchSize)
		*batchSize = 100
	}
	if *batchSize < 1 {
		log.Fatal("Error: Batch size must be at least 1")
	}

	// Parse and validate date filters
	var startDateParsed, endDateParsed *time.Time
	if *startDate != "" {
		parsed, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			log.Fatalf("Error: Invalid start-date format. Use YYYY-MM-DD: %v", err)
		}
		startDateParsed = &parsed
	}
	if *endDate != "" {
		parsed, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			log.Fatalf("Error: Invalid end-date format. Use YYYY-MM-DD: %v", err)
		}
		endDateParsed = &parsed
	}
	if startDateParsed != nil && endDateParsed != nil && startDateParsed.After(*endDateParsed) {
		log.Fatal("Error: start-date must be before or equal to end-date")
	}

	// Get credentials from environment
	username := os.Getenv("TT_USER")
	password := os.Getenv("TT_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("Error: TT_USER and TT_PASSWORD environment variables must be set")
	}

	// Normalize symbol (uppercase, no spaces)
	normalizedSymbol := strings.ToUpper(strings.TrimSpace(*symbol))
	if normalizedSymbol == "" {
		log.Fatal("Error: Symbol cannot be empty")
	}

	fmt.Printf("%s Options Chain Market Data Dump\n", normalizedSymbol)
	fmt.Println(strings.Repeat("=", len(normalizedSymbol)+32))
	fmt.Println()

	// Initialize API client
	api := tastytrade.NewTastytradeAPI()
	api.SetAPIVersion("20250715")

	// Authenticate
	fmt.Println("Authenticating...")
	err := api.Authenticate(username, password)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}
	fmt.Println("✓ Authenticated")
	fmt.Println()

	// Get option chain for the specified symbol
	fmt.Printf("Fetching %s option chain...\n", normalizedSymbol)
	if startDateParsed != nil || endDateParsed != nil {
		dateRange := "all dates"
		if startDateParsed != nil && endDateParsed != nil {
			dateRange = fmt.Sprintf("%s to %s", startDateParsed.Format("2006-01-02"), endDateParsed.Format("2006-01-02"))
		} else if startDateParsed != nil {
			dateRange = fmt.Sprintf("from %s", startDateParsed.Format("2006-01-02"))
		} else if endDateParsed != nil {
			dateRange = fmt.Sprintf("until %s", endDateParsed.Format("2006-01-02"))
		}
		fmt.Printf("  Date filter: %s\n", dateRange)
	}
	optionChain, err := api.ListOptionsChainsDetailed(normalizedSymbol)
	if err != nil {
		log.Fatalf("Failed to fetch option chain: %v", err)
	}
	fmt.Printf("✓ Found %d option contracts", len(optionChain.Data.Items))
	if startDateParsed != nil || endDateParsed != nil {
		fmt.Printf(" (will filter by date range)")
	}
	fmt.Println()

	// Filter options by date range BEFORE fetching quotes
	filteredOptions := make([]tastytrade.OptionDataDetailed, 0)
	for _, option := range optionChain.Data.Items {
		// Apply date range filter if specified
		if startDateParsed != nil || endDateParsed != nil {
			expDate, err := time.Parse("2006-01-02", option.ExpirationDate)
			if err != nil {
				// Skip if we can't parse the expiration date
				continue
			}

			// Check start date filter
			if startDateParsed != nil && expDate.Before(*startDateParsed) {
				continue
			}

			// Check end date filter
			if endDateParsed != nil && expDate.After(*endDateParsed) {
				continue
			}
		}
		filteredOptions = append(filteredOptions, option)
	}

	fmt.Printf("✓ Filtered to %d options within date range\n", len(filteredOptions))
	fmt.Println()

	// Collect option symbols for quote fetching (only for filtered options)
	optionSymbols := make([]string, 0, len(filteredOptions))
	optionMap := make(map[string]tastytrade.OptionDataDetailed) // Map symbol to option data

	for _, option := range filteredOptions {
		optionSymbols = append(optionSymbols, option.Symbol)
		optionMap[option.Symbol] = option
	}

	// Fetch quotes for all options (in batches to maximize data per request and minimize API calls)
	fmt.Printf("Fetching option quotes (batching %d symbols per request, %dms delay between requests)...\n", *batchSize, *delayMs)
	totalSymbols := len(optionSymbols)
	expectedRequests := (totalSymbols + *batchSize - 1) / *batchSize
	fmt.Printf("  Will make approximately %d API requests for %d total symbols\n", expectedRequests, totalSymbols)

	quoteMap := make(map[string]tastytrade.QuoteData)
	delayDuration := time.Duration(*delayMs) * time.Millisecond

	for i := 0; i < len(optionSymbols); i += *batchSize {
		end := i + *batchSize
		if end > len(optionSymbols) {
			end = len(optionSymbols)
		}
		batch := optionSymbols[i:end]

		quotes, err := api.GetQuotesByType(&tastytrade.QuoteQueryParams{
			EquityOption: batch,
		})
		if err != nil {
			log.Printf("Warning: Failed to fetch quotes for batch %d-%d: %v\n", i, end, err)
			// Still add delay even on error to respect rate limits
			if i+*batchSize < len(optionSymbols) {
				time.Sleep(delayDuration)
			}
			continue
		}

		for _, quote := range quotes.Data.Items {
			quoteMap[quote.Symbol] = quote
		}

		requestsCompleted := (i + *batchSize) / *batchSize
		fmt.Printf("  Fetched quotes for %d/%d options (%d requests completed)...\r", end, totalSymbols, requestsCompleted)

		// Add delay between requests to avoid rate limiting (except after the last batch)
		if i+*batchSize < len(optionSymbols) {
			time.Sleep(delayDuration)
		}
	}
	fmt.Printf("\n✓ Fetched quotes for %d options in %d API requests\n\n", len(quoteMap), expectedRequests)

	// Group options by expiration - include ALL available data
	type OptionRow struct {
		// Option chain data
		Symbol                         string
		StreamerSymbol                 string
		ExpirationDate                 string
		ExpiresAt                      string
		StrikePrice                    float64
		OptionType                     string
		RootSymbol                     string
		UnderlyingSymbol               string
		DaysToExpiration               int
		Active                         bool
		IsClosingOnly                  bool
		ExpirationType                 string
		SettlementType                 string
		ExerciseStyle                  string
		SharesPerContract              int
		OptionChainType                string
		ListedMarket                   string
		HaltedAt                       string
		StopsTradingAt                 string
		OldSecurityNumber              string
		MarketTimeInstrumentCollection string

		// Quote data - prices
		Bid               float64
		BidSize           float64
		Ask               float64
		AskSize           float64
		Mid               float64
		Mark              float64
		Last              float64
		LastMkt           float64
		Open              float64
		Close             float64
		PrevClose         float64
		DayHighPrice      float64
		DayLowPrice       float64
		YearHighPrice     float64
		YearLowPrice      float64
		Beta              float64
		DividendAmount    float64
		DividendFrequency float64
		LowLimitPrice     float64
		HighLimitPrice    float64

		// Quote data - volume and interest
		OpenInterest float64
		Volume       float64

		// Quote data - other
		UpdatedAt       string
		SummaryDate     string
		PrevCloseDate   string
		IsTradingHalted bool
		HaltStartTime   int64
		HaltEndTime     int64

		// Greeks (if available in quote response)
		Delta     float64
		Gamma     float64
		Theta     float64
		Vega      float64
		Rho       float64
		IV        float64 // Implied Volatility (from "volatility" field)
		TheoPrice float64 // Theoretical price
		DxMark    float64 // DX mark price
		TickSize  float64 // Tick size
	}

	expirationMap := make(map[string][]OptionRow)

	// Process filtered options (date filtering already done above)
	for _, option := range filteredOptions {
		// Parse strike price
		strike, err := strconv.ParseFloat(option.StrikePrice, 64)
		if err != nil {
			continue
		}

		expiration := option.ExpirationDate

		// Get quote data
		quote, hasQuote := quoteMap[option.Symbol]

		// Parse Open Interest
		var oi float64
		if hasQuote && len(quote.OpenInterest) > 0 {
			var oiValue interface{}
			if err := json.Unmarshal(quote.OpenInterest, &oiValue); err == nil {
				switch v := oiValue.(type) {
				case float64:
					oi = v
				case int64:
					oi = float64(v)
				case string:
					oi, _ = strconv.ParseFloat(v, 64)
				}
			}
		}

		// Helper function to parse string to float64
		parseFloat := func(s string) float64 {
			if s == "" {
				return 0
			}
			val, _ := strconv.ParseFloat(s, 64)
			return val
		}

		// Parse all quote price fields
		bid := parseFloat(quote.Bid)
		bidSize := parseFloat(quote.BidSize)
		ask := parseFloat(quote.Ask)
		askSize := parseFloat(quote.AskSize)
		mid := parseFloat(quote.Mid)
		mark := parseFloat(quote.Mark)
		last := parseFloat(quote.Last)
		lastMkt := parseFloat(quote.LastMkt)
		open := parseFloat(quote.Open)
		close := parseFloat(quote.Close)
		prevClose := parseFloat(quote.PrevClose)
		dayHigh := parseFloat(quote.DayHighPrice)
		dayLow := parseFloat(quote.DayLowPrice)
		yearHigh := parseFloat(quote.YearHighPrice)
		yearLow := parseFloat(quote.YearLowPrice)
		volume := parseFloat(quote.Volume)

		// Extract Greeks from quote data (they're now in the struct)
		var delta, gamma, theta, vega, rho, iv, theoPrice, dxMark, tickSize float64
		if hasQuote {
			delta = parseFloat(quote.Delta)
			gamma = parseFloat(quote.Gamma)
			theta = parseFloat(quote.Theta)
			vega = parseFloat(quote.Vega)
			rho = parseFloat(quote.Rho)
			iv = parseFloat(quote.Volatility) // Implied volatility is in "volatility" field
			theoPrice = parseFloat(quote.TheoPrice)
			dxMark = parseFloat(quote.DxMark)
			tickSize = parseFloat(quote.TickSize)
		}

		row := OptionRow{
			// Option chain data
			Symbol:                         option.Symbol,
			StreamerSymbol:                 option.StreamerSymbol,
			ExpirationDate:                 expiration,
			ExpiresAt:                      option.ExpiresAt,
			StrikePrice:                    strike,
			OptionType:                     option.OptionType,
			RootSymbol:                     option.RootSymbol,
			UnderlyingSymbol:               option.UnderlyingSymbol,
			DaysToExpiration:               option.DaysToExpiration,
			Active:                         option.Active,
			IsClosingOnly:                  option.IsClosingOnly,
			ExpirationType:                 option.ExpirationType,
			SettlementType:                 option.SettlementType,
			ExerciseStyle:                  option.ExerciseStyle,
			SharesPerContract:              option.SharesPerContract,
			OptionChainType:                option.OptionChainType,
			ListedMarket:                   option.ListedMarket,
			HaltedAt:                       option.HaltedAt,
			StopsTradingAt:                 option.StopsTradingAt,
			OldSecurityNumber:              option.OldSecurityNumber,
			MarketTimeInstrumentCollection: option.MarketTimeInstrumentCollection,

			// Quote prices
			Bid:           bid,
			BidSize:       bidSize,
			Ask:           ask,
			AskSize:       askSize,
			Mid:           mid,
			Mark:          mark,
			Last:          last,
			LastMkt:       lastMkt,
			Open:          open,
			Close:         close,
			PrevClose:     prevClose,
			DayHighPrice:  dayHigh,
			DayLowPrice:   dayLow,
			YearHighPrice: yearHigh,
			YearLowPrice:  yearLow,

			// Volume and interest
			OpenInterest: oi,
			Volume:       volume,

			// Other quote data
			UpdatedAt:       quote.UpdatedAt,
			SummaryDate:     quote.SummaryDate,
			PrevCloseDate:   quote.PrevCloseDate,
			IsTradingHalted: quote.IsTradingHalted,
			HaltStartTime:   quote.HaltStartTime,
			HaltEndTime:     quote.HaltEndTime,

			// Greeks
			Delta:     delta,
			Gamma:     gamma,
			Theta:     theta,
			Vega:      vega,
			Rho:       rho,
			IV:        iv,
			TheoPrice: theoPrice,
			DxMark:    dxMark,
			TickSize:  tickSize,
		}

		expirationMap[expiration] = append(expirationMap[expiration], row)
	}

	// Sort expirations
	expirations := make([]string, 0, len(expirationMap))
	for exp := range expirationMap {
		expirations = append(expirations, exp)
	}
	sort.Strings(expirations)

	fmt.Println()

	// Write CSV files
	fmt.Println("Writing CSV files...")
	csvFilesCreated := 0

	// Helper function to format float64 for CSV
	formatFloat := func(f float64, decimals int) string {
		if f == 0 {
			return ""
		}
		return strconv.FormatFloat(f, 'f', decimals, 64)
	}

	// Helper function to format bool as 1/0 for CSV (saves space)
	formatBool := func(b bool) string {
		if b {
			return "1"
		}
		return "0"
	}

	// Write comprehensive header
	header := []string{
		// Option chain identifiers
		"Symbol", "StreamerSymbol", "ExpirationDate", "ExpiresAt", "StrikePrice", "OptionType",
		"RootSymbol", "UnderlyingSymbol", "DaysToExpiration",
		// Option chain metadata
		"Active", "IsClosingOnly", "StopsTradingAt",
		// Quote prices
		"Bid", "BidSize", "Ask", "AskSize", "Mid", "Mark", "Last", "LastMkt",
		"Open", "Close", "PrevClose",
		"DayHighPrice", "DayLowPrice", "YearHighPrice", "YearLowPrice",
		// Volume and interest
		"OpenInterest", "Volume",
		// Quote metadata
		"UpdatedAt", "SummaryDate", "PrevCloseDate", "IsTradingHalted",
		"HaltStartTime", "HaltEndTime",
		// Greeks
		"Delta", "Gamma", "Theta", "Vega", "Rho", "ImpliedVolatility", "TheoPrice", "DxMark", "TickSize",
	}

	// Helper function to write a row to CSV
	writeRow := func(writer *csv.Writer, row OptionRow) error {
		record := []string{
			// Option chain identifiers
			row.Symbol,
			row.StreamerSymbol,
			row.ExpirationDate,
			row.ExpiresAt,
			formatFloat(row.StrikePrice, 2),
			row.OptionType,
			row.RootSymbol,
			row.UnderlyingSymbol,
			strconv.Itoa(row.DaysToExpiration),
			// Option chain metadata
			formatBool(row.Active),
			formatBool(row.IsClosingOnly),
			row.StopsTradingAt,
			// Quote prices
			formatFloat(row.Bid, 2),
			formatFloat(row.BidSize, 0),
			formatFloat(row.Ask, 2),
			formatFloat(row.AskSize, 0),
			formatFloat(row.Mid, 2),
			formatFloat(row.Mark, 2),
			formatFloat(row.Last, 2),
			formatFloat(row.LastMkt, 2),
			formatFloat(row.Open, 2),
			formatFloat(row.Close, 2),
			formatFloat(row.PrevClose, 2),
			formatFloat(row.DayHighPrice, 2),
			formatFloat(row.DayLowPrice, 2),
			formatFloat(row.YearHighPrice, 2),
			formatFloat(row.YearLowPrice, 2),
			// Volume and interest
			formatFloat(row.OpenInterest, 0),
			formatFloat(row.Volume, 0),
			// Quote metadata
			row.UpdatedAt,
			row.SummaryDate,
			row.PrevCloseDate,
			formatBool(row.IsTradingHalted),
			strconv.FormatInt(row.HaltStartTime, 10),
			strconv.FormatInt(row.HaltEndTime, 10),
			// Greeks
			formatFloat(row.Delta, 6),
			formatFloat(row.Gamma, 6),
			formatFloat(row.Theta, 6),
			formatFloat(row.Vega, 6),
			formatFloat(row.Rho, 6),
			formatFloat(row.IV, 6),
			formatFloat(row.TheoPrice, 2),
			formatFloat(row.DxMark, 2),
			formatFloat(row.TickSize, 2),
		}
		return writer.Write(record)
	}

	if *perDay {
		// Write individual CSV files per expiration
		for _, expiration := range expirations {
			rows := expirationMap[expiration]

			// Sort by strike price, then by option type
			sort.Slice(rows, func(i, j int) bool {
				if rows[i].StrikePrice != rows[j].StrikePrice {
					return rows[i].StrikePrice < rows[j].StrikePrice
				}
				return rows[i].OptionType < rows[j].OptionType
			})

			// Create individual CSV file for this expiration
			filename := fmt.Sprintf("%s_options_%s.csv", strings.ToLower(normalizedSymbol), expiration)
			file, err := os.Create(filename)
			if err != nil {
				log.Printf("Failed to create CSV file %s: %v", filename, err)
				continue
			}

			writer := csv.NewWriter(file)

			// Write header
			if err := writer.Write(header); err != nil {
				log.Printf("Failed to write header to %s: %v", filename, err)
				file.Close()
				continue
			}

			// Write rows
			for _, row := range rows {
				if err := writeRow(writer, row); err != nil {
					log.Printf("Failed to write record to %s: %v", filename, err)
					continue
				}
			}

			writer.Flush()
			file.Close()

			fmt.Printf("✓ Created CSV: %s (%d rows)\n", filename, len(rows))
			csvFilesCreated++
		}

		fmt.Printf("\n✓ Market data dump complete! Created %d individual files.\n", csvFilesCreated)
	} else {
		// Write single combined CSV file
		// Get today's date for the combined filename
		today := time.Now().Format("2006-01-02")

		combinedFile, err := os.Create(fmt.Sprintf("%s_options_all_expirations_%s.csv", strings.ToLower(normalizedSymbol), today))
		if err != nil {
			log.Fatalf("Failed to create combined CSV file: %v", err)
		}
		defer combinedFile.Close()

		combinedWriter := csv.NewWriter(combinedFile)
		defer combinedWriter.Flush()

		// Write header
		if err := combinedWriter.Write(header); err != nil {
			log.Fatalf("Failed to write header: %v", err)
		}

		// Write all rows from all expirations
		totalRows := 0
		for _, expiration := range expirations {
			rows := expirationMap[expiration]

			// Sort by strike price, then by option type
			sort.Slice(rows, func(i, j int) bool {
				if rows[i].StrikePrice != rows[j].StrikePrice {
					return rows[i].StrikePrice < rows[j].StrikePrice
				}
				return rows[i].OptionType < rows[j].OptionType
			})

			// Write rows to combined file
			for _, row := range rows {
				if err := writeRow(combinedWriter, row); err != nil {
					log.Printf("Failed to write record to combined file: %v", err)
					continue
				}
				totalRows++
			}
		}

		combinedWriter.Flush()
		fmt.Printf("✓ Created combined CSV: %s_options_all_expirations_%s.csv (%d rows)\n", strings.ToLower(normalizedSymbol), today, totalRows)
		fmt.Printf("\n✓ Market data dump complete! Created 1 combined file.\n")
	}
}
