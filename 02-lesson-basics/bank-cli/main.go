package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/appliedgocourses/bank"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	var err error = nil

	// Restore the bank data.
	// Call bank.Load() here and handle any error.
	if err = bank.Load(); err != nil {
		log.Fatal("Could not load bank data.", err)
	}

	// Save the bank data.
	// Use a deferred function for calling bank.Save().
	// (See the lecture on function behavior.)
	// If Save() returns an error, print it.

	defer func() {
		if err = bank.Save(); err != nil {
			log.Fatal("Could not save bank data.", err)
		}
	}()

	// Perform the action.
	// os.Args[0] is the path to the executable.
	// os.Args[1] is the first parameter - the action we want to perform:
	// create, list, update, transfer, or history.

	switch os.Args[1] {

	case "create":
		// bank create <name>                     Create an account.
		name := os.Args[2]
		if name == "" {
			log.Fatal("Must provide a name.")
		}
		account := create(name)
		fmt.Println("Account created.", account)
	case "list":
		fmt.Println(list())
	case "update":
		name := os.Args[2]
		amount, err := strconv.Atoi(os.Args[3])
		if name == "" || err != nil {
			log.Fatal("Must provide a name and amount.")
		}
		balance, err := update(name, amount)
		if err != nil {
			log.Fatal("Could not update account", err)
		} else {
			fmt.Printf("Account updated. Balance: %d\n", balance)
		}
	case "transfer":
		// bank transfer <name> <name> <amount>   Transfer money between two accounts.
		name, name2 := os.Args[2], os.Args[3]
		amount, err := strconv.Atoi(os.Args[4])
		if name == "" || name2 == "" || err != nil {
			log.Fatal("Must provide two names and amount.")
		}
		if balanceFrom, balanceTo, err := transfer(name, name2, amount); err != nil {
			log.Fatal("Could not transfer funds.", err)
		} else {
			fmt.Printf("Trnasfer complete. balanceFrom: %d, balanceTo: %d\n", balanceFrom, balanceTo)
		}
	case "history":
		name := os.Args[2]
		transactions, err := history(name)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(transactions)
		}
	default:
		log.Fatal("Action is not valid.")
	}
}

func usage() {
	fmt.Println(`Usage:

bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
}

func create(name string) *bank.Account {
	return bank.NewAccount(name)
}

func list() string {
	return bank.ListAccounts()
}

// update takes a name and an amount, deposits the amount if it
// is greater than zero, or withdraws it if it is less than zero,
// and returns the new balance and any error that occurred.
func update(name string, amount int) (int, error) {
	account, err := bank.GetAccount(name)
	if err != nil {
		return -1, err
	}

	switch {
	case amount > 0:
		return bank.Deposit(account, amount)
	case amount < 0:
		return bank.Withdraw(account, -amount)
	default:
		return bank.Balance(account), nil
	}
}

// transfer takes two names and an amount, transfers the amount from
// the account belonging to name #1 to the account belonging to name #2,
// and returns the new balances of both accounts and any error that occurred.
func transfer(name, name2 string, amount int) (int, int, error) {
	accountFrom, err := bank.GetAccount(name)
	if err != nil {
		return -1, -1, err
	}
	accountTo, err := bank.GetAccount(name2)
	if err != nil {
		return -1, -1, err
	}
	return bank.Transfer(accountFrom, accountTo, amount)
}

// history takes an account name, retrieves the account, and calls bank.History()
// to get the history closure function. Then it calls the closure in a loop,
// formatting the return values and appending the result to the output string, until the boolean return parameter of the closure is `false`.
func history(name string) (string, error) {
	account, err := bank.GetAccount(name)
	if err != nil {
		return "", err
	}
	var transactions string
	iterHistory := bank.History(account)
	for {
		amt, bal, more := iterHistory()
		transactions += fmt.Sprintf("amt: %d, bal: %d\n", amt, bal)
		if more == false {
			break
		}
	}
	return transactions, nil
}
