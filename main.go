package main

import (
	"bufio"
	"errors"
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

type Tuple struct {
	val1 string
	val2 int
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
	fmt.Println("Please enter expenses! In this format \"type:amount -- Example: Food:300\"! You can also finish with expenses by typing finish!")
	tuple, err := getNameValueFormatFromUser()

	if err != nil {
		if tuple.val1 == "finish" {
			account.initInitialSplit()
			return true
		}

		fmt.Println(err.Error())
		return false
	}

	_, ok := account.expenses[tuple.val1]

	// if we do not have the name in the expense map
	// just set it
	if ok == false {
		account.expenses[tuple.val1] = tuple.val2
		return false
	}

	account.expenses[tuple.val1] += tuple.val2

	return false
}

func (account *Account) scanPercentages() bool {
	fmt.Printf("Percentage used: %v --- %v left \n", account.percentageUsed, 100 - account.percentageUsed)
	fmt.Println("Please enter splits! In this format \"type:amount -- Example: House:20\"! You can also finish with splits by typing finish!")

	tuple, err := getNameValueFormatFromUser()

	if err != nil {
		if tuple.val1 == "finish" {
			return true
		}

		fmt.Println(err.Error())
		return false
	}

	_, ok := account.splits[tuple.val1]

	// if we do not have the name in the expense map
	// just set it
	if ok == true {
		fmt.Printf("%v is already used", tuple.val1)
		return false
	}

	// if user has enter amount less than 0 and percent greater than 100 with the sum of other percentage
	// we will say the they exceed the percentage limit
	if tuple.val2 < 0 || tuple.val2 + account.percentageUsed > 100 {
		fmt.Println("You are exceeding the percentage limit!")
		return false
	}

	account.percentageUsed += tuple.val2
	account.splits[tuple.val1] = tuple.val2

	if account.percentageUsed >= 100 {
		return true
	}

	return false
}

func getNameValueFormatFromUser() (tuple Tuple, error error) {
	line := scanLine()

	if strings.ContainsRune(line, ':') == false {
		return Tuple{line, 0}, errors.New("format for split is not compatible")
	}

	splittedSplit := strings.Split(line, ":")
	name := splittedSplit[0]
	amount, err := strconv.Atoi(splittedSplit[1])

	if err != nil {
		return Tuple{line, 0}, errors.New("couldn't convert number")
	}

	return Tuple{name, amount}, nil
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
