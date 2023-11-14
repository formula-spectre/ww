package main

import (
	"encoding/csv"
	"fmt"
	"os"
	//	"time"
)

func main() {
	// Get user inputs for date and number
	dateInput := getUserInput("Enter date (YYYY-MM-DD): ")
	numberInput := getUserInputInt("Enter a number: ")

	// Open or create the CSV file for appending
	file, err := os.OpenFile("transactions.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the user inputs to the CSV file
	if err := writer.Write([]string{dateInput, fmt.Sprintf("%d", numberInput)}); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Transaction successfully stored in transactions.csv")
}

// getUserInput prompts the user for input and returns the entered value as a string
func getUserInput(prompt string) string {
	var userInput string
	fmt.Print(prompt)
	fmt.Scanln(&userInput)
	return userInput
}

// getUserInputInt prompts the user for input and returns the entered value as an integer
func getUserInputInt(prompt string) int {
	var userInput int
	fmt.Print(prompt)
	fmt.Scanln(&userInput)
	return userInput
}
