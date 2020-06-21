package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Account struct {
	income, percentageUsed int
	expenses, splits map[string]int
}

func scanLine() string {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func (account *Account) initInitialSplit() {
	if len(account.expenses) == 0 {
		return
	}

	var expensesSum int

	for _, value := range account.expenses {
		expensesSum += value
	}


	expensesPercentageUsed := int(math.Ceil(float64(expensesSum) / float64(account.income) * 100))
	account.percentageUsed = expensesPercentageUsed
	account.splits["Expenses"] = expensesPercentageUsed
}

func (account *Account) scanExpenses() bool {
	var expense string

	fmt.Println("Please enter expenses! In this format \"type:amount -- Example: Food:300\"! You can also finish with expenses by typing finish!")
	expense = scanLine()

	if strings.Contains(expense, "finish") {
		account.initInitialSplit()
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
		fmt.Println("Error")
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
	fmt.Printf("Percentage used: %v --- %v left \n", account.percentageUsed, 100 - account.percentageUsed)
	fmt.Println("Please enter splits! In this format \"type:amount -- Example: House:20\"! You can also finish with splits by typing finish!")

	split := scanLine()

	if strings.ContainsRune(split, ':') == false {
		fmt.Println("Format for split is not compatible!")
		return false
	}

	splittedSplit := strings.Split(split, ":")
	name := splittedSplit[0]
	amount, err := strconv.Atoi(splittedSplit[1])

	if err != nil {
		return false
	}

	_, ok := account.splits[name]

	// if we do not have the name in the expense map
	// just set it
	if ok == true {
		fmt.Printf("%v is already used", name)
		return false
	}

	// if user has enter amount less than 0 and percent greater than 100 with the sum of other percentage
	// we will say the they exceed the percentage limit
	if amount < 0 || amount + account.percentageUsed > 100 {
		fmt.Println("You are exceeding the percentage limit!")
		return false
	}

	account.percentageUsed += amount
	account.splits[name] = amount

	if account.percentageUsed >= 100 {
		return true
	}

	return false
}

func (account *Account) showPlanning() {
	fmt.Println("Here is your planning")

	for name, amount := range account.splits {
		fmt.Printf("%v - %v = %vlv\n", name, amount, ((account.income / 100) * amount))
	}
}

type Stage string

const (
	income Stage = "Income"
	expense = "Expense"
	percentageSplit = "Percentage Split"
)

func main() {
	account := Account{
		expenses:       make(map[string]int),
		splits:         make(map[string]int),
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

				sum := 0
				for name, amount := range account.expenses {
					fmt.Printf("%v - %vlv\n", name, amount)
					sum += amount
				}

				fmt.Printf("Total Expenses: %vlv\n", sum)
				fmt.Println("Press any key to continue!")
				scanLine()

				continue
			}
		} else if stage == percentageSplit {
			if account.scanPercentages() {
				break
			}
		}
	}

	account.showPlanning()
}
