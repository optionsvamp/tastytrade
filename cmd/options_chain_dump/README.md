# Options Chain Market Data Dump

This application fetches option chain data for any underlying symbol from the Tastytrade API and exports it to CSV files for later analysis. By default, it fetches SPX (S&P 500 Index) options, but you can specify any symbol.

## Features

- Fetches option chain data for any underlying symbol from Tastytrade API (default: SPX)
- Retrieves **ALL** available market data from quotes including:
  - All price data (Bid, Ask, Mid, Mark, Last, Open, Close, etc.)
  - Volume and Open Interest
  - Greeks (Delta, Gamma, Theta, Vega, Rho, Implied Volatility) - if available
  - All option chain metadata
- Exports data to CSV format (configurable):
  - Default: Single combined CSV file with all expirations: `{SYMBOL}_options_all_expirations_YYYY-MM-DD.csv`
  - Optional: One CSV file per expiration date: `{SYMBOL}_options_YYYY-MM-DD.csv` (use `-per-day` flag)
- Includes **comprehensive data fields** (50+ columns):
  - Option identifiers and metadata
  - All quote prices (Bid, Ask, Mid, Mark, Last, Open, Close, High, Low, etc.)
  - Volume and Open Interest
  - Greeks (Delta, Gamma, Theta, Vega, Rho, IV)
  - Trading status and timestamps

## Requirements

- Go 1.22 or later
- Tastytrade API credentials (set as environment variables)

## Usage

1. Set your Tastytrade API credentials:
   ```bash
   export TT_USER=your_username
   export TT_PASSWORD=your_password
   ```

2. Build the application:
   ```bash
   go build -o options_chain_dump ./cmd/options_chain_dump
   ```

3. Run the application:
   ```bash
   ./options_chain_dump
   ```

   Or run directly:
   ```bash
   go run ./cmd/options_chain_dump/main.go
   ```

   To specify a different underlying symbol:
   ```bash
   go run ./cmd/options_chain_dump/main.go -symbol AAPL
   go run ./cmd/options_chain_dump/main.go -symbol TSLA
   go run ./cmd/options_chain_dump/main.go -symbol SPX
   ```
   
   To filter by expiration date range:
   ```bash
   # Options expiring from a specific date onwards
   go run ./cmd/options_chain_dump/main.go -symbol SPX -start-date 2026-01-15
   
   # Options expiring up to a specific date
   go run ./cmd/options_chain_dump/main.go -symbol SPX -end-date 2026-02-15
   
   # Options expiring within a date range (inclusive)
   go run ./cmd/options_chain_dump/main.go -symbol SPX -start-date 2026-01-15 -end-date 2026-02-15
   
   # Combine with other options
   go run ./cmd/options_chain_dump/main.go -symbol AAPL -start-date 2026-01-01 -end-date 2026-01-31 -batch-size 50
   ```

   To control rate limiting and batch size:
   ```bash
   # Reduce batch size for more conservative approach (default is 100, max is 100)
   go run ./cmd/options_chain_dump/main.go -symbol SPX -batch-size 50
   
   # Add more delay between requests (1000ms instead of default 500ms)
   go run ./cmd/options_chain_dump/main.go -symbol SPX -delay-ms 1000
   
   # More conservative settings for large option chains
   go run ./cmd/options_chain_dump/main.go -symbol SPX -batch-size 50 -delay-ms 1000
   ```
   
   To control output format:
   ```bash
   # Default: Single combined CSV file (all expirations in one file)
   go run ./cmd/options_chain_dump/main.go -symbol SPX
   
   # Per-day: Separate CSV file for each expiration date
   go run ./cmd/options_chain_dump/main.go -symbol SPX -per-day
   ```
   
   **Note**: Batch size is capped at 100 symbols per request (server limit). Values above 100 will be automatically reduced.

   Or with the built binary:
   ```bash
   ./options_chain_dump -symbol AAPL
   ./options_chain_dump -symbol SPX -batch-size 500 -delay-ms 100
   ```

## Output Files

### Default Mode (Single Combined File)
- Format: `{SYMBOL}_options_all_expirations_YYYY-MM-DD.csv` (e.g., `spx_options_all_expirations_2026-01-10.csv`)
- Contains all option contracts across all expiration dates in a single file
- Date in filename is the date when the dump was created (today's date)
- Sorted by expiration date, then by strike price, then by option type

### Per-Day Mode (`-per-day` flag)
- Format: `{SYMBOL}_options_YYYY-MM-DD.csv` (e.g., `spx_options_2026-01-12.csv` or `aapl_options_2026-01-12.csv`)
- One file per expiration date
- Contains all option contracts for that specific expiration date
- Sorted by strike price, then by option type (Calls before Puts)

## CSV Format

Each CSV file contains comprehensive data with the following columns:

### Option Chain Identifiers
| Column | Description |
|--------|-------------|
| Symbol | Option symbol in OCC format |
| StreamerSymbol | Streamer symbol format |
| ExpirationDate | Expiration date (YYYY-MM-DD) |
| ExpiresAt | Expiration timestamp |
| StrikePrice | Strike price of the option |
| OptionType | "C" for Call, "P" for Put |
| RootSymbol | Root symbol for the option |
| UnderlyingSymbol | Underlying equity symbol (SPX) |
| DaysToExpiration | Days until expiration |

### Option Chain Metadata
| Column | Description |
|--------|-------------|
| Active | Whether the option is currently active |
| IsClosingOnly | Whether only closing positions are allowed |
| InstrumentType | Type of instrument (e.g., "Equity Option") |
| ExpirationType | Expiration type (e.g., "Standard", "Weekly") |
| SettlementType | Settlement type (e.g., "Physical", "Cash") |
| ExerciseStyle | Exercise style: "American" or "European" |
| SharesPerContract | Number of shares per contract (typically 100) |
| OptionChainType | Type of option chain |
| ListedMarket | Market where the option is listed |
| HaltedAt | Time when trading was halted |
| StopsTradingAt | Time when trading stops |
| OldSecurityNumber | Old security number |
| MarketTimeInstrumentCollection | Market time instrument collection identifier |

### Quote Prices
| Column | Description |
|--------|-------------|
| Bid | Bid price |
| BidSize | Bid size |
| Ask | Ask price |
| AskSize | Ask size |
| Mid | Mid price (average of bid and ask) |
| Mark | Mark price |
| Last | Last trade price |
| LastMkt | Last market price |
| Open | Opening price |
| Close | Closing price |
| ClosePriceType | Close price type (e.g., "Final", "Regular") |
| PrevClose | Previous close price |
| PrevClosePriceType | Previous close price type |
| DayHighPrice | Day high price |
| DayLowPrice | Day low price |
| YearHighPrice | Year high price |
| YearLowPrice | Year low price |

### Volume and Interest
| Column | Description |
|--------|-------------|
| OpenInterest | Open interest (number of contracts) |
| Volume | Trading volume |

### Quote Metadata
| Column | Description |
|--------|-------------|
| UpdatedAt | Timestamp of last update |
| SummaryDate | Summary date |
| PrevCloseDate | Previous close date |
| IsTradingHalted | Whether trading is halted |
| HaltStartTime | Halt start time (-1 if not halted) |
| HaltEndTime | Halt end time (-1 if not halted) |

### Greeks (if available)
| Column | Description |
|--------|-------------|
| Delta | Delta (price sensitivity to underlying) |
| Gamma | Gamma (rate of change of delta) |
| Theta | Theta (time decay) |
| Vega | Vega (volatility sensitivity) |
| Rho | Rho (interest rate sensitivity) |
| ImpliedVolatility | Implied volatility |

## Rate Limiting & API Efficiency

The application is designed to minimize API requests while respecting rate limits:

- **Batching**: Requests multiple symbols (default: 100, max: 100) in a single API call to maximize data per request
- **Rate Limiting**: Adds configurable delays (default: 500ms) between batch requests
- **Efficiency**: For SPX with ~3000 options, this results in ~30 API requests instead of 3000 individual requests
- **Server Limit**: The API server enforces a maximum of 100 symbols per request. Values above 100 are automatically capped.

You can adjust these settings:
- `-batch-size`: Reduce below 100 for more conservative approach (e.g., 50)
- `-delay-ms`: Increase if you encounter rate limiting (e.g., 1000ms for more conservative approach)

## Notes

- The application fetches quotes in batches to maximize data per request and minimize API calls
- If Open Interest or Greeks are not available in the quote data, they will be empty/0
- All numeric fields are formatted appropriately:
  - Prices: 2 decimal places
  - Volume/OI: integers (0 decimals)
  - Greeks: 6 decimal places for precision
- Empty/zero values are left blank in CSV for cleaner data
- Files are created in the current working directory
- Greeks are extracted from the raw JSON response if available (may vary by API response format)