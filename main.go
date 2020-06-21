package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Account struct {
	income int
	expenses, splits map[string]int
}

func (account *Account) scanExpenses() bool {
	var expense string

	fmt.Println("Please enter expenses! In this format \"type:amount -- Example: Food:300\"! You can also finish with expenses by typing finish!")
	_, err := fmt.Scan(&expense)
	
	if err != nil {
		return false
	}

	if strings.Contains(expense, "finish") {
		return true
	}

	if strings.ContainsRune(expense, ':') == false {
		fmt.Println("Format for expense is not compatible!")
		return false
	}
	
	splittedExpense := strings.Split(expense, ":")
	name := splittedExpense[0]
	amount, err := strconv.Atoi(splittedExpense[1])

	if err != nil {
		return false
	}

	_, ok := account.expenses[name]

	// if we do not have the name in the expense map
	// just set it
	if ok == false {
		account.expenses[name] = amount
		return false
	}

	account.expenses[name] += amount

	return false
}

func (account *Account) scanPercentages() bool {
	return false
}

type Stage string

const (
	income Stage = "Income"
	expense = "Expense"
	percentageSplit = "Percentage Split"
)

func main() {
	account := Account{
		expenses: make(map[string]int),
		splits: make(map[string]int),
	}
	stage := income

	for account.income == 0 {
		fmt.Println("Please enter your income!")
		fmt.Scan(&account.income)
	}

	stage = expense

	for {
		if stage == expense {
			if account.scanExpenses() {
				stage = percentageSplit

				fmt.Println("These are the expenses you entered")

				for name, amount := range account.expenses {
					fmt.Printf("%v - %vlv\n", name, amount)
				}

				fmt.Println("Press any key to continue!")
				fmt.Scan()

				continue
			}
		} else if stage == percentageSplit {
			if account.scanPercentages() {
				break
			}
		}
	}
}
