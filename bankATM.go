package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Users struct {
	Username       string
	CardNo         int
	PinNo          int
	AccountType    string
	AccountBalance float64
}

var usr Users

var reader = bufio.NewReader(os.Stdin)

func main() {
	users := []Users{
		{"Juhi Mehta", 33802817, 6402, "Savings Account", 43174.24},
		{"Janvi Sureja", 96452487, 1308, "Savings Account", 36870.00},
		{"Dhruvi Patel", 20053476, 1512, "Salary Account", 54166.00},
		{"Meet Radadiya", 45793425, 1306, "Recurring Deposit Account", 14161.59},
		{"Tirth Bhatt", 76129863, 3878, "Bussiness Account", 78000.00},
		{"Dev Kapadiya", 43560098, 2354, "Savings Account", 41567.63},
		{"Pakshal Ghiya", 85148640, 2803, "Salary Account", 40133.00},
		{"Hinal Panchal", 50709876, 3402, "Recurring Deposit Account", 21786.34},
		{"Khushali Shah", 15863465, 2382, "Savings Account", 35000.00},
	}
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("Welcome to Bank. Please enter correct details")
	fmt.Println(strings.Repeat("-", 60))
	userDetails(&users)

Services:
	fmt.Println("Please select the service you want.")
	fmt.Println("1. Withdraw money from the ATM")
	fmt.Println("2. Deposit Money into the ATM")
	fmt.Println("3. Check Bank Balance")
	fmt.Println("4. Exit")
	fmt.Print("Plese enter your choice: ")
	input, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic("Please enter the valid choice. The choice must be a number & between 1 to 4")
	}
	switch choice {
	case 1:
		withdrawMoney(&users)
		goto Services
	case 2:
		depositMoney(&users)
		goto Services
	case 3:
		checkBalance(&users)
		goto Services
	case 4:
		fmt.Println("Thank you...Visit Again")
		os.Exit(0)
	}

}
func userDetails(u *[]Users) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			userDetails(u)
		}
	}()
CardNo:
	fmt.Print("Please enter your Card Number: ")
	input, _ := reader.ReadString('\n')
	cardNo, err := strconv.Atoi(strings.TrimSpace(input))
	x := len(input) - 1
	if x != 8 {
		fmt.Println("Please enter 8 digit valid card number")
		goto CardNo
	}
	if err != nil {
		panic("Please enter valid input. Enter digits only")
	}
	for _, value := range *u {
		if cardNo == value.CardNo {
			usr = value
			// fmt.Println(value)
			break
		}
	}
	if (usr == Users{}) {
		fmt.Println("User not found")
		os.Exit(0)
	}
	attemps := 3
PinNo:
	fmt.Print("Enter your Pin Number: ")
	input, _ = reader.ReadString('\n')
	pinNo, err := strconv.Atoi(strings.TrimSpace(input))
	x = len(input) - 1
	if pinNo != usr.PinNo {
		attemps--
		if attemps == 0 {
			fmt.Println("You have exceeded the maximum number of attempts. Please try again after sometimes")
			os.Exit(0)
		} else if x != 4 {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Println("Please enter the valid 4 digit Pin Number")
			fmt.Printf("%d attempts left\n", attemps)
			fmt.Println(strings.Repeat("-", 60))
			goto PinNo
		} else {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Println("Please enter the correct Pin Number")
			fmt.Printf("%d attempts left\n", attemps)
			fmt.Println(strings.Repeat("-", 60))
			goto PinNo
		}
	} else {
		fmt.Println(strings.Repeat("-", 60))
		fmt.Println("You are successfully logged in.")
		fmt.Println(strings.Repeat("-", 60))
	}

}

func isContinue() bool {
	fmt.Printf("Do you want to continue Y/N? ")
	input, _ := reader.ReadString('\n')
	ans := strings.TrimSpace(input)
	if ans == "y" || ans == "Y" {
		return true
	} else {
		return false
	}
}

func withdrawMoney(u *[]Users) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			withdrawMoney(u)
		}
	}()
Withdraw:
	fmt.Println(strings.Repeat("-", 60))
	fmt.Print("Please enter the amount you want to withdraw: ")
	input, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	if err != nil {
		fmt.Println(strings.Repeat("-", 60))
		panic("Please enter the valid input. Input must be a Number.")
	}
	if float64(amount) > usr.AccountBalance {
		fmt.Println(strings.Repeat("-", 60))
		fmt.Println("Your account does not have enough balance.")
		fmt.Println(strings.Repeat("-", 60))
		if isContinue() {
			goto Withdraw
		} else {
			os.Exit(0)
		}
	} else if amount >= 0 && amount < 500 {
		fmt.Println("Withdrawal amount shoulde be greater than 500.")
		fmt.Println(strings.Repeat("-", 60))
		if isContinue() {
			goto Withdraw
		} else {
			os.Exit(0)
		}
	} else if amount > 0 {
		amt := math.Mod(amount, 500)
		if amt == 0 {
			usr.AccountBalance -= amount
			fmt.Println(strings.Repeat("-", 60))
			fmt.Println("Your Current Balance is: ", usr.AccountBalance)
			fmt.Println(strings.Repeat("-", 60))
			if isContinue() {
				return
			} else {
				os.Exit(0)
			}
		} else {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Println("Withdrawal amount should be in multiplication of 500")
			fmt.Println(strings.Repeat("-", 60))
			if isContinue() {
				goto Withdraw
			} else {
				os.Exit(0)
			}
		}
	} else {
		fmt.Println(strings.Repeat("-", 60))
		fmt.Println("Please enter the valid input. Input must be a Positive Number.")
		fmt.Println(strings.Repeat("-", 60))
		if isContinue() {
			goto Withdraw
		} else {
			os.Exit(0)
		}
	}
}

func depositMoney(u *[]Users) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			depositMoney(u)
		}
	}()
Deposit:
	fmt.Print("Please enter the amount you want to deposit: ")
	input, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	if err != nil {
		panic("Please enter the input amount. The input should be a number.")
	}
	if amount < 0 {
		fmt.Println("Please enter the valid input. Input must be a Positive Number.")
		if isContinue() {
			goto Deposit
		} else {
			os.Exit(0)
		}
	} else if amount >= 0 && amount < 500 {
		fmt.Println("Deposit amount shoulde be greater than 500.")
		if isContinue() {
			goto Deposit
		} else {
			os.Exit(0)
		}
	} else if amount > 50000 {
		fmt.Println("Deposit amount should be less than 50000.")
		if isContinue() {
			goto Deposit
		} else {
			os.Exit(0)
		}
	} else {
		amt := math.Mod(amount, 500)
		if amt == 0 {
			fmt.Println("Your current balance is: ", usr.AccountBalance)
			usr.AccountBalance += amount
			fmt.Println("Yo have successfully deposited amount: ", amount)
			fmt.Printf("Your new balance is: %.2f\n", usr.AccountBalance)
			if isContinue() {
				return
			} else {
				os.Exit(0)
			}
		} else {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Println("Deposit amount should be in multiplication of 500")
			fmt.Println(strings.Repeat("-", 60))
			if isContinue() {
				goto Deposit
			} else {
				os.Exit(0)
			}
		}

	}
}

func checkBalance(u *[]Users) {
	fmt.Println("Your Current Balance is: ", usr.AccountBalance)
	if isContinue() {
		return
	} else {
		os.Exit(0)
	}
}
