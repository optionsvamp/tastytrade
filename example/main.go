package main

import (
	"fmt"
	"github.com/optionsvamp/tastytrade"
	"os"
)

func main() {
	// Example usage:
	api := tastytrade.NewTastytradeAPI("")

	// Authenticate with Tastytrade API
	err := api.Authenticate(os.Getenv("USER"), os.Getenv("PWD"))
	if err != nil {
		fmt.Println("Authentication failed:", err)
		return
	}

	//customerInfo, err := api.GetCustomerInfo()
	//if err != nil {
	//	fmt.Println("Error fetching customer data:", err)
	//	return
	//}
	//fmt.Println("Customer data:", customerInfo)

	// Get a list of customer accounts
	customerAccounts, err := api.ListCustomerAccounts()
	if err != nil {
		fmt.Println("Error fetching customer accounts:", err)
		return
	}
	//fmt.Println("Customer accounts:", customerAccounts)

	// Get account trading status for the first account in the previous response
	status, err := api.GetAccountTradingStatus(customerAccounts.Data.Items[0].Account.AccountNumber)
	if err != nil {
		fmt.Println("Error fetching customer account status:", err)
		return
	}
	fmt.Println("Customer account status:", status)

	// Get account
	acct, err := api.GetAccount(customerAccounts.Data.Items[0].Account.AccountNumber)
	if err != nil {
		fmt.Println("Error fetching customer account:", err)
		return
	}
	fmt.Println("Customer account:", acct)

	// Get account balance
	bal, err := api.GetAccountBalances(customerAccounts.Data.Items[0].Account.AccountNumber)
	if err != nil {
		fmt.Println("Error fetching account balance:", err)
		return
	}
	fmt.Println("Account balance:", bal)

	snap, err := api.GetAccountBalanceSnapshots(customerAccounts.Data.Items[0].Account.AccountNumber, "2024-01-01", "BOD")
	if err != nil {
		fmt.Println("Error fetching account balance snapshot:", err)
		return
	}
	fmt.Println("Account balance snapshot:", snap)

	// Get symbol data
	symbolData, err := api.GetEquityData("AAPL")
	if err != nil {
		fmt.Println("Error fetching symbol data:", err)
		return
	}
	fmt.Println("Symbol data:", symbolData)

	transactions, err := api.GetTransactions(customerAccounts.Data.Items[2].Account.AccountNumber, &tastytrade.TransactionQueryParams{
		Symbol:         "/MCLN4",
		InstrumentType: "Future",
	})
	if err != nil {
		fmt.Println("Error fetching transactions data:", err)
		return
	}
	fmt.Println("Transactions data:", transactions)

	symbols, err := api.ListEquities(nil)
	if err != nil {
		fmt.Println("Error fetching symbol data:", err)
		return
	}
	_ = symbols

	// Get option chain
	optionChain, err := api.ListOptionsChainsDetailed("AAPL")
	if err != nil {
		fmt.Println("Error fetching option chain:", err)
		return
	}
	fmt.Println("Option chain:", optionChain)
}
