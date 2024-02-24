package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var result []recipe

func main() {
	var userInput string
	fmt.Println("Entre an Ingredient")
	fmt.Scan(&userInput)
	fmt.Println("You entered:", userInput)

	// Open the CSV file
	file, err := os.Open("Recipe_dataset.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	outp := fixUserInput(userInput)
	fmt.Print(outp)
	fmt.Print("length: ")
	fmt.Print(len(outp))
	fmt.Print("\n")
	fmt.Print("\n")
	var recipes []recipe = fillRecipe(records)
	result = getRecipe(&recipes, outp)

	/*for _, r := range recipes {
		fmt.Printf("Title: %s\nIngredients: %s\nCleaned Ingredients: %s\n\n",
			r.Title, r.Ingredients, r.CleanedIngredients)
	}*/

	for _, r := range result {
		fmt.Printf("Title: %s\nIngredients: %s\nCleaned Ingredients: %s\n\n",
			r.Title, r.Ingredients, r.CleanedIngredients)
	}
}

func fixUserInput(str1 string) string {
	var output string
	for _, char := range str1 {
		if char != ' ' {
			output += string(char)
		}
	}
	return output
}
