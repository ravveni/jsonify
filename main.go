package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run csv_to_json.go input_file.csv")
		return
	}

	inputFile := os.Args[1]

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
	var jsonArray []map[string]interface{}
	for _, record := range records[1:] { // Skip header row
		item := make(map[string]interface{})
		for j, value := range record {
			item[keys[j]] = value
		}

		jsonArray = append(jsonArray, item)
	}

	// Convert the array of maps to JSON format and write it to standard output
	outputJson, err := json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	io.WriteString(os.Stdout, string(outputJson)) // Write the JSON output to stdout
}
