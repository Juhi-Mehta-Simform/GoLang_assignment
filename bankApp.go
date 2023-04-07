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

// var user1 Users
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

	fmt.Println("Welcome to Bank. Please enter correct details")
	// usr := userDetails(&users)
	userDetails(&users)
	// if usr.PinNo == 0 {
	// 	os.Exit(0)
	// }
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
		panic("Please enter the valid choice. The choice must be number & between 1 to 4")
	}
	switch choice {
	case 1:
		withdrawMoney(&users)
		goto Services
	case 2:
		// depositAmount(&usr)
		goto Services
	case 3:
		goto Services
	case 4:
		fmt.Println("Thank you...Visit Again")
		break
	}
}

// var user Users

func userDetails(u *[]Users) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			userDetails(u)
		}
	}()
	for _, value := range *u {
	CardNo:
		fmt.Print("\nPlease enter your Card Number: ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		cardNo, err := strconv.Atoi(strings.TrimSpace(input))
		x := len(input) - 1

		if err != nil {
			panic("Please enter valid input")
		}
		fmt.Println(cardNo)
		fmt.Println(value.CardNo)
		if cardNo == value.CardNo {
			attemps := 3
			for _, value := range *u {
			Pin:
				fmt.Print("\nEnter your Pin Number: ")
				input, _ = reader.ReadString('\n')
				pinNo, err := strconv.Atoi(strings.TrimSpace(input))
				x = len(input) - 1
				defer func() {
					if r := recover(); r != nil {
						fmt.Println(r)
					}
				}()
				if err != nil {
					panic("Please enter valid pin")
				}
				if pinNo != value.PinNo {
					attemps--
					if attemps == 0 {
						fmt.Println("You have exceeded the maximum number of attempts. Please try again after sometimes")

					} else if x != 4 {
						fmt.Println("Please enter the valid 4 digit Pin Number")
						fmt.Printf("%d attempts left", attemps)
						goto Pin
					} else {
						fmt.Println("Please enter the correct Pin Number")
						fmt.Printf("%d attempts left", attemps)
						goto Pin
					}
					break
				} else {
					fmt.Println("You are successfully logged in.")
					break
				}
				break
			}
			break
		} else if x != 8 {
			fmt.Println("Please enter the valid 8 digit Card Number")
			goto CardNo
		} else {
			fmt.Println("user not found")
			os.Exit(0)
		}
	}
}

func withdrawMoney(u *[]Users) bool {
Withdraw:
	fmt.Print("Please enter the amount you want to withdraw: ")
	input, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			withdrawMoney(u)
		}
	}()
	if err != nil {
		panic("Please enter the valid amount.")
	}
	if float64(amount) > u.AccountBalance {
		fmt.Println("Your account does not have enough balance.")
		fmt.Println("Do you want to continue Y/N?")
		input, _ = reader.ReadString('\n')
		ans := strings.TrimSpace(input)
		if ans == "y" || ans == "Y" {
			goto Withdraw
		} else {
			return false
		}
	} else if amount > 0 {
		amt := math.Mod(amount, 100)
		if amt == 0 {
			u.AccountBalance -= amount
			fmt.Println("Current Balance: ", u.AccountBalance)
			return true
		} else {
			fmt.Println("Withdrawal amount should be in multiplication of 100")
			return false
		}
	} else {
		return false
	}
}

func depositAmount(u *Users) {
Deposit:
	fmt.Print("Please enter the amount you want to deposit: ")
	input, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			depositAmount(u)
		}
	}()
	if err != nil {
		panic("Please enter the valid amount.")
	}
	if amount < 0 {
		fmt.Println("Invalid amount")
	} else if amount > 0 && amount < 100 {
		fmt.Println("The minimum amount should be 100")
		fmt.Println("Do you want to continue Y/N?")
		input, _ = reader.ReadString('\n')
		ans := strings.TrimSpace(input)
		if ans == "y" || ans == "Y" {
			goto Deposit
		}
	} else if amount > 50000 {
		fmt.Println("You can't deposit more than 50000.")
		fmt.Println("Do you want to continue Y/N?")
		input, _ = reader.ReadString('\n')
		ans := strings.TrimSpace(input)
		if ans == "y" || ans == "Y" {
			goto Deposit
		}
	} else {
		fmt.Println("Your current balance is: ", u.AccountBalance)
		u.AccountBalance += amount
		fmt.Println("Yo have successfully deposited amount: ", amount)
		fmt.Println("Your new balance is: ", u.AccountBalance)
	}
}
