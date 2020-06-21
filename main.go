package main

import (
	"fmt"
	"strings"
)

type Account struct {
	income int
	expenses map[string]int
	splits map[string]int
}

func (account *Account) scanExpenses() bool {
	var expense string

	fmt.Println("Please enter expenses! In this format \"type:amount -- Example: Food:300\"")
	fmt.Scan(&expense)

	if strings.Contains(expense, "exit") {
		return true
	}

	if strings.ContainsRune(expense, ':') == false {
		fmt.Println("Format for expense is not compatible!")
		return false
	}

	return false
}

type Stage string

const (
	income Stage = "Income"
	expense = "Expense"
	percentageSplit = "Percentage Split"
)

func main() {
	account := Account{}
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
				continue
			}
		} else if stage == percentageSplit {

		}
	}
}
