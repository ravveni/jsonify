// jsonify -- a simple csv to json converter
// Copyright (C) 2024 ravveni

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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

	// Copy input filename as output filename
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
			item[keys[j]] = getTypedValue(value)
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

func getTypedValue(value string) any {
	intValue, err := strconv.ParseInt(value, 0, 64) // Check if value is an int
	if err == nil {
		return intValue
	}

	floatValue, err := strconv.ParseFloat(value, 64) // Check if value is a float
	if err == nil {
		return floatValue
	}

	boolValue, err := strconv.ParseBool(value) // Check if value is a bool
	if err == nil {
		return boolValue
	}

	cellArray := strings.Split(value, "|")
	if len(cellArray) > 1 { // Check if value is an array
		cellArrayItems := []any{}
		for _, item := range cellArray {
			cellArrayItems = append(cellArrayItems, getTypedValue(item))
		}
		return cellArrayItems
	}

	return value // Value is a string
}
