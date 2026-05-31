package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"unit-converter/converter"
)

func main() {
	// Reader to capture user input from the terminal
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop so the program keeps running until user exits
	for {
		// Display menu options
		fmt.Println("\n=== UNIT CONVERTER ===")
		fmt.Println("1. Kg → Lbs")
		fmt.Println("2. Lbs → Kg")
		fmt.Println("3. Km → Miles")
		fmt.Println("4. Miles → Km")
		fmt.Println("5. Liters → Gallons")
		fmt.Println("6. Gallons → Liters")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		// Read menu choice from user
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		// Convert input string to integer
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 0-6.")
			continue
		}

		// Exit condition
		if choice == 0 {
			fmt.Println("Goodbye 👋")
			break
		}

		// Ask user for the value to convert
		fmt.Print("Enter value: ")
		valueStr, _ := reader.ReadString('\n')
		valueStr = strings.TrimSpace(valueStr)

		// Convert input string to float
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println("Invalid number input. Try again.")
			continue
		}

		// Perform conversion based on selected option
		switch choice {
		case 1:
			fmt.Println("Result:", converter.KgToLbs(value))
		case 2:
			fmt.Println("Result:", converter.LbsToKg(value))
		case 3:
			fmt.Println("Result:", converter.KmToMiles(value))
		case 4:
			fmt.Println("Result:", converter.MilesToKm(value))
		case 5:
			fmt.Println("Result:", converter.LitersToGallons(value))
		case 6:
			fmt.Println("Result:", converter.GallonsToLiters(value))
		default:
			fmt.Println("Invalid option selected.")
		}
	}
}