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
	// Reader for handling user input from terminal
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display menu options
		fmt.Println("\n===== UNIT CONVERTER =====")
		fmt.Println("1. Kilograms → Pounds")
		fmt.Println("2. Pounds → Kilograms")
		fmt.Println("3. Kilometers → Miles")
		fmt.Println("4. Miles → Kilometers")
		fmt.Println("5. Liters → Gallons")
		fmt.Println("6. Gallons → Liters")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		// Read user menu selection
		input := readInput(reader)

		// Convert menu input to integer
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 0–6.")
			continue
		}

		// Exit condition
		if choice == 0 {
			fmt.Println("Goodbye 👋")
			break
		}

		// Ask user for value to convert
		fmt.Print("Enter value: ")
		valueInput := readInput(reader)

		// Convert value input to float
		value, err := strconv.ParseFloat(valueInput, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		// Perform conversion based on selected option
		switch choice {

		case 1:
			fmt.Printf("Result: %.4f lbs\n", converter.KgToLbs(value))

		case 2:
			fmt.Printf("Result: %.4f kg\n", converter.LbsToKg(value))

		case 3:
			fmt.Printf("Result: %.4f miles\n", converter.KmToMiles(value))

		case 4:
			fmt.Printf("Result: %.4f km\n", converter.MilesToKm(value))

		case 5:
			fmt.Printf("Result: %.4f gallons\n", converter.LitersToGallons(value))

		case 6:
			fmt.Printf("Result: %.4f liters\n", converter.GallonsToLiters(value))

		// Handles invalid menu choices (e.g. 9, 99, -1)
		default:
			fmt.Println("Invalid option. Please select a number between 0–6.")
		}
	}
}

// readInput handles all user input and trims unnecessary spaces/newlines
func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}