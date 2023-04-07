package main

import (
	"GoPractical/functions"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	users := []functions.Users{
		{Username: "Juhi Mehta", CardNo: 33802817, PinNo: 6402, AccountType: "Savings Account", AccountBalance: 43174.24},
		{Username: "Janvi Sureja", CardNo: 96452487, PinNo: 1308, AccountType: "Savings Account", AccountBalance: 36870.00},
		{Username: "Dhruvi Patel", CardNo: 20053476, PinNo: 1512, AccountType: "Salary Account", AccountBalance: 54166.00},
		{Username: "Meet Radadiya", CardNo: 45793425, PinNo: 1306, AccountType: "Recurring Deposit Account", AccountBalance: 14161.59},
		{Username: "Tirth Bhatt", CardNo: 76129863, PinNo: 3878, AccountType: "Business Account", AccountBalance: 78000.00},
		{Username: "Pakshal Ghiya", CardNo: 85148640, PinNo: 2803, AccountType: "Salary Account", AccountBalance: 40133.00},
		{Username: "Hinal Panchal", CardNo: 50709876, PinNo: 3402, AccountType: "Recurring Deposit Account", AccountBalance: 21786.34},
		{Username: "Khushali Shah", CardNo: 15863465, PinNo: 2382, AccountType: "Savings Account", AccountBalance: 35000.00},
	}
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("Welcome to Bank. Please enter correct details")
	functions.UserDetails(&users)

Services:
	fmt.Println("Please select the service you want.")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("1. Withdraw money from the ATM")
	fmt.Println("2. Deposit Money into the ATM")
	fmt.Println("3. Check Bank Balance")
	fmt.Println("4. Exit")
	fmt.Println(strings.Repeat("-", 60))
Choice:
	fmt.Print("Plese enter your choice: ")
	input, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(strings.Repeat("-", 60))
		fmt.Println("Please enter the valid choice. The choice must be a number & between 1 to 4")
		fmt.Println(strings.Repeat("-", 60))
		goto Choice
	}
	switch choice {
	case 1:
		functions.WithdrawMoney(&users)
		goto Services
	case 2:
		functions.DepositMoney(&users)
		goto Services
	case 3:
		functions.CheckBalance(&users)
		goto Services
	case 4:
		fmt.Println("Thank you...Visit Again")
		fmt.Println(strings.Repeat("-", 60))
		os.Exit(0)
	default:
		fmt.Println("Please enter the number between 1 to 4")
		fmt.Println(strings.Repeat("-", 60))
		goto Choice
	}

}
