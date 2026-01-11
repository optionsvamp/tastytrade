package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/optionsvamp/tastytrade"
)

func main() {
	// Parse command-line flags
	accountNumber := flag.String("account", "", "Account number to fetch futures options positions for (e.g., 3AA14199)")
	flag.Parse()

	if *accountNumber == "" {
		log.Fatal("Error: -account flag is required. Specify the account number (e.g., -account 3AA14199)")
	}

	// Normalize account number (uppercase, no spaces)
	normalizedAccount := strings.ToUpper(strings.TrimSpace(*accountNumber))
	if normalizedAccount == "" {
		log.Fatal("Error: Account number cannot be empty")
	}

	// Get credentials from environment
	username := os.Getenv("TT_USER")
	password := os.Getenv("TT_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("Error: TT_USER and TT_PASSWORD environment variables must be set")
	}

	fmt.Printf("Futures Options Positions Viewer\n")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()

	// Initialize API client
	api := tastytrade.NewTastytradeAPI()

	// Authenticate
	fmt.Println("Authenticating...")
	err := api.Authenticate(username, password)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}
	fmt.Println("✓ Authenticated")
	fmt.Println()

	// List customer accounts to find the matching account
	fmt.Printf("Fetching account list...\n")
	accountsResponse, err := api.ListCustomerAccounts()
	if err != nil {
		log.Fatalf("Failed to fetch accounts: %v", err)
	}

	// Find the account matching the provided account number
	var targetAccount *tastytrade.AccountContainer
	for i := range accountsResponse.Data.Items {
		if accountsResponse.Data.Items[i].Account.AccountNumber == normalizedAccount {
			targetAccount = &accountsResponse.Data.Items[i]
			break
		}
	}

	if targetAccount == nil {
		log.Fatalf("Error: Account '%s' not found in your account list", normalizedAccount)
	}

	fmt.Printf("✓ Found account: %s (%s)\n", targetAccount.Account.AccountNumber, targetAccount.Account.Nickname)
	fmt.Println()

	// Get positions for the account
	fmt.Printf("Fetching positions for account %s...\n", normalizedAccount)
	positionsResponse, err := api.GetPositions(normalizedAccount)
	if err != nil {
		log.Fatalf("Failed to fetch positions: %v", err)
	}

	fmt.Printf("✓ Found %d total positions\n", len(positionsResponse.Data.Items))
	fmt.Println()

	// Filter for futures options positions
	futuresOptionsPositions := make([]tastytrade.Position, 0)
	for _, position := range positionsResponse.Data.Items {
		if position.InstrumentType == "Future Option" {
			futuresOptionsPositions = append(futuresOptionsPositions, position)
		}
	}

	if len(futuresOptionsPositions) == 0 {
		fmt.Println("No futures options positions found.")
		return
	}

	// Collect symbols for quote fetching
	futuresOptionSymbols := make([]string, 0, len(futuresOptionsPositions))
	for _, pos := range futuresOptionsPositions {
		futuresOptionSymbols = append(futuresOptionSymbols, pos.Symbol)
	}

	// Fetch quotes for all futures options positions to get Greeks
	fmt.Printf("Fetching quotes for %d futures options positions...\n", len(futuresOptionSymbols))
	quotesResponse, err := api.GetQuotesByType(&tastytrade.QuoteQueryParams{
		FutureOption: futuresOptionSymbols,
	})
	if err != nil {
		log.Printf("Warning: Failed to fetch quotes: %v\n", err)
		log.Println("Continuing without Greeks data...")
	}

	// Create a map of symbol to quote for quick lookup
	quoteMap := make(map[string]tastytrade.QuoteData)
	for _, quote := range quotesResponse.Data.Items {
		quoteMap[quote.Symbol] = quote
	}

	fmt.Printf("✓ Fetched quotes for %d positions\n", len(quoteMap))
	fmt.Println()

	// Helper function to parse string to float64
	parseFloat := func(s string) float64 {
		if s == "" {
			return 0
		}
		val, _ := strconv.ParseFloat(s, 64)
		return val
	}

	// Helper function to parse quantity (handles both string and number formats)
	parseQuantity := func(s string) float64 {
		if s == "" {
			return 0
		}
		// Remove any commas
		s = strings.ReplaceAll(s, ",", "")
		val, _ := strconv.ParseFloat(s, 64)
		return val
	}

	// Calculate net delta and net theta
	var netDelta, netTheta float64
	type positionWithGreeks struct {
		Position      tastytrade.Position
		Delta         float64
		Theta         float64
		Quantity      float64
	}

	positionsWithGreeks := make([]positionWithGreeks, 0, len(futuresOptionsPositions))

	for _, pos := range futuresOptionsPositions {
		quote, hasQuote := quoteMap[pos.Symbol]
		
		var delta, theta float64
		if hasQuote {
			delta = parseFloat(quote.Delta)
			theta = parseFloat(quote.Theta)
		}

		quantity := parseQuantity(pos.Quantity)
		multiplier := float64(pos.Multiplier)
		
		// Determine direction multiplier: +1 for Long, -1 for Short
		directionMultiplier := 1.0
		if strings.ToUpper(pos.QuantityDirection) == "SHORT" {
			directionMultiplier = -1.0
		}

		// Calculate position delta and theta
		// Greeks are per contract, so we multiply by:
		// - quantity (number of contracts)
		// - multiplier (contract multiplier, e.g., 50 for ES, 20 for NQ)
		// - direction (positive for long, negative for short)
		positionDelta := delta * quantity * multiplier * directionMultiplier
		positionTheta := theta * quantity * multiplier * directionMultiplier

		netDelta += positionDelta
		netTheta += positionTheta

		positionsWithGreeks = append(positionsWithGreeks, positionWithGreeks{
			Position:      pos,
			Delta:         delta,
			Theta:         theta,
			Quantity:      quantity,
		})
	}

	// Print futures options positions with Greeks
	fmt.Printf("Futures Options Positions (%d):\n", len(futuresOptionsPositions))
	fmt.Println(strings.Repeat("-", 130))

	// Print header
	fmt.Printf("%-20s %-15s %-10s %-10s %-8s %-12s %-12s %-12s %-12s %-15s\n",
		"Symbol",
		"Underlying",
		"Quantity",
		"Direction",
		"Mult",
		"Pos Delta",
		"Pos Theta",
		"Avg Open",
		"Close Price",
		"Realized Today",
	)
	fmt.Println(strings.Repeat("-", 140))

	// Print each position
	for _, pg := range positionsWithGreeks {
		// Calculate position Greeks for display
		directionMultiplier := 1.0
		if strings.ToUpper(pg.Position.QuantityDirection) == "SHORT" {
			directionMultiplier = -1.0
		}
		positionDelta := pg.Delta * pg.Quantity * float64(pg.Position.Multiplier) * directionMultiplier
		positionTheta := pg.Theta * pg.Quantity * float64(pg.Position.Multiplier) * directionMultiplier

		positionDeltaStr := fmt.Sprintf("%.2f", positionDelta)
		positionThetaStr := fmt.Sprintf("%.2f", positionTheta)
		if pg.Delta == 0 && pg.Theta == 0 {
			positionDeltaStr = "N/A"
			positionThetaStr = "N/A"
		}

		fmt.Printf("%-20s %-15s %-10.0f %-10s %-8d %-12s %-12s %-12s %-12s %-15s\n",
			pg.Position.Symbol,
			pg.Position.UnderlyingSymbol,
			pg.Quantity,
			pg.Position.QuantityDirection,
			pg.Position.Multiplier,
			positionDeltaStr,
			positionThetaStr,
			pg.Position.AverageOpenPrice,
			pg.Position.ClosePrice,
			pg.Position.RealizedToday,
		)
	}

	fmt.Println(strings.Repeat("-", 140))

	// Print net Greeks summary
	fmt.Println()
	fmt.Println("Net Greeks Summary:")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Net Delta: %12.4f\n", netDelta)
	fmt.Printf("Net Theta: %12.4f\n", netTheta)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("\n✓ Displayed %d futures options positions\n", len(futuresOptionsPositions))
}
