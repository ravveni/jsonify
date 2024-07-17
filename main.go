package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run jsonify input_file.csv")
		return
	}

	inputFile := os.Args[1]

	// Copying input filename as output filename
	outputFilename := strings.Split(inputFile, ".")[0]

	// Open the CSV file for reading
	csvFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Parse the CSV file using the standard library's csv package
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Use the first row as keys for JSON objects
	keys := records[0]

	// Convert each row (CSV record) to a map and create an array of maps
	var jsonArray []map[string]any
	for _, record := range records[1:] { // Skip header row
		item := make(map[string]any)
		for j, value := range record {
			intValue, err := strconv.ParseInt(value, 0, 64) // Check if value is an int
			if err == nil {
				item[keys[j]] = intValue
				continue
			}

			floatValue, err := strconv.ParseFloat(value, 64) // Check if value is a float
			if err == nil {
				item[keys[j]] = floatValue
				continue
			}

			boolValue, err := strconv.ParseBool(value) // Check if value is a bool
			if err == nil {
				item[keys[j]] = boolValue
				continue
			}

			item[keys[j]] = value // Value is a string
		}

		jsonArray = append(jsonArray, item)
	}

	// Convert the array of maps to JSON format and write it to standard output
	outputJson, err := json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create the output file with .json extension
	outputFilename = outputFilename + ".json"
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write JSON to the output file
	_, err = io.WriteString(outputFile, string(outputJson))
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("JSON data saved to:", outputFilename)
}
