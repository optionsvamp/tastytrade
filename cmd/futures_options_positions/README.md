# Futures Options Positions Viewer

This application logs into Tastytrade, fetches a specific account from your account list, and displays all current futures options positions for that account.

## Features

- Authenticates with Tastytrade API using environment variables
- Lists all your accounts and finds the specified account by account number
- Fetches all positions for the specified account
- Filters and displays only futures options positions
- Prints positions in a readable table format to the console

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
   go build -o futures_options_positions ./cmd/futures_options_positions
   ```

3. Run the application with an account number:
   ```bash
   ./futures_options_positions -account 3AA14199
   ```

   Or run directly:
   ```bash
   go run ./cmd/futures_options_positions/main.go -account 3AA14199
   ```

## Command-Line Arguments

- `-account` (required): The account number to fetch futures options positions for (e.g., `3AA14199`)

## Output

The application displays a table with the following columns for each futures options position:

- **Symbol**: The futures option symbol
- **Underlying**: The underlying futures symbol
- **Quantity**: Position quantity
- **Direction**: "Long" or "Short"
- **Avg Open**: Average opening price
- **Close Price**: Current close price
- **Realized Today**: Realized P&L for today

## Example Output

```
Futures Options Positions Viewer
==================================================

Authenticating...
✓ Authenticated

Fetching account list...
✓ Found account: 3AA14199 (My Trading Account)

Fetching positions for account 3AA14199...
✓ Found 15 total positions

Futures Options Positions (3):
----------------------------------------------------------------------------------------------------
Symbol               Underlying      Quantity         Direction    Avg Open     Close Price     Realized Today
----------------------------------------------------------------------------------------------------
/ESM6 EW2G6 260213C  /ESM6           2                Long         125.50       130.25          +950.00
/NQM6 NQ2G6 260213P  /NQM6           1                Short        85.00        80.50           +450.00
/CLM6 CL2G6 260213C  /CLM6           5                Long         2.50         3.00            +250.00
----------------------------------------------------------------------------------------------------

✓ Displayed 3 futures options positions
```

## Notes

- The account number is case-insensitive (automatically converted to uppercase)
- If the account number is not found in your account list, the application will exit with an error
- If there are no futures options positions, the application will display a message and exit
- Only positions with `InstrumentType == "Future Option"` are displayed
