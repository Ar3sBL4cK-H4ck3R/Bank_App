package main

import "fmt"

func Welcome() {
	fmt.Println("\nWhat Do You Want to Do?")
	fmt.Println("1. Check Your Balance")
	fmt.Println("2. Deposite Amount")
	fmt.Println("3. Withdrawal Amount")
	fmt.Println("4. Exit")
	fmt.Print("Enter Your Choice:")
	fmt.Scanln(&choice)
}
