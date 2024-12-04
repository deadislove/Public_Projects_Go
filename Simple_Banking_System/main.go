package main

import (
	"Simple_Banking_System/services"
	"fmt"
)

func main() {
	account := services.NewAccount("12345", "John Doe")

	// Menu options
	for {
		fmt.Println("\nSimple Banking System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)

		case 2:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println(err)
			}

		case 3:
			fmt.Printf("Balance: %.2f\n", account.GetBalance())

		case 4:
			fmt.Println("Exiting... Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
