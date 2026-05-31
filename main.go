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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== UNIT CONVERTER =====")
		fmt.Println("1. Kilograms → Pounds")
		fmt.Println("2. Pounds → Kilograms")
		fmt.Println("3. Kilometers → Miles")
		fmt.Println("4. Miles → Kilometers")
		fmt.Println("5. Liters → Gallons")
		fmt.Println("6. Gallons → Liters")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if choice == 0 {
			fmt.Println("Goodbye 👋")
			break
		}

		fmt.Print("Enter value: ")
		valueInput, _ := reader.ReadString('\n')
		valueInput = strings.TrimSpace(valueInput)

		value, err := strconv.ParseFloat(valueInput, 64)
		if err != nil {
			fmt.Println("Invalid number input.")
			continue
		}

		switch choice {

		case 1:
			fmt.Println("Result:", converter.KgToLbs(value), "lbs")

		case 2:
			fmt.Println("Result:", converter.LbsToKg(value), "kg")

		case 3:
			fmt.Println("Result:", converter.KmToMiles(value), "miles")

		case 4:
			fmt.Println("Result:", converter.MilesToKm(value), "km")

		case 5:
			fmt.Println("Result:", converter.LitersToGallons(value), "gallons")

		case 6:
			fmt.Println("Result:", converter.GallonsToLiters(value), "liters")

		default:
			fmt.Println("Invalid option. Please choose between 0–6.")
		}
	}
}