package functions

import (
	"fmt"
	"os"
	"strings"
)

func CheckBalance(u *[]Users) {
	fmt.Printf("Your Current Balance is: %.2f\n", usr.AccountBalance)
	fmt.Println(strings.Repeat("-", 60))
	if IsContinue() {
		return
	} else {
		os.Exit(0)
	}
}
