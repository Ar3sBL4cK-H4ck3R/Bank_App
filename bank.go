package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const balanceFile = "Go_Balance.txt"

var choice int

func main() {
	var currentBalance, err = fileops.ReadFloatToFile(balanceFile)

	fmt.Println("\n$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("~~~~~~~~~~~~~~~^ Welcome To GO Bank! ^~~~~~~~~~~~~~~~")
	fmt.Print("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$\n")

	for {
		Welcome()
		switch choice {
		case 1:
			fmt.Println("Your Current Balance:", currentBalance)
			if err != nil {
				fmt.Println("< ERROR >")
				fmt.Println(err)
				fmt.Println("-----------------")
			}
			continue
		case 2:
			depositeAmount(currentBalance)
			continue
		case 3:
			withdrawalAmount(currentBalance)
			continue
		case 4:
			fmt.Println("Good Bye!")
			fmt.Println("Thanks For Chosing Our Bank!")
			return
		default:
			fmt.Println("Please Choose The Right Option!")
			continue
		}
	}

}

func depositeAmount(currentBalance float64) {
	var deposite float64
	fmt.Print("Enter Deposite Amount:")
	fmt.Scanln(&deposite)
	if deposite <= 0 {
		fmt.Println("Please Enter The Valid Amount!")
		return
	} else {
		deposite += currentBalance
		updatedbalance := fmt.Sprint(deposite)
		fileops.WirteFloatTofile(balanceFile, updatedbalance)
		fmt.Printf("Thank You For Depositing $%.2f", deposite)
	}

}

func withdrawalAmount(currentBalance float64) {
	var withdrawal float64
	fmt.Print("Enter Withdrawal Amount:")
	fmt.Scanln(&withdrawal)
	if withdrawal <= 0 {
		fmt.Println("Please Enter The Valid Amount!")
		return
	} else if withdrawal > currentBalance {
		fmt.Println("Your Account Balance is not Enough!")
		return
	} else {
		updatedAmount := currentBalance - withdrawal
		updatedWithdrawal := fmt.Sprint(updatedAmount)
		fileops.WirteFloatTofile(balanceFile, updatedWithdrawal)
		fmt.Printf("Thanks For Your Withdrawal $%.2f", updatedAmount)
	}
}
