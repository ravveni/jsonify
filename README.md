```
   _                 _  __       
  (_)               (_)/ _|      
   _ ___  ___  _ __  _| |_ _   _ 
  | / __|/ _ \| '_ \| |  _| | | |
  | \__ \ (_) | | | | | | | |_| |
  | |___/\___/|_| |_|_|_|  \__, |
 _/ |                       __/ |
|__/                       |___/ 

simple csv to json converter
```

## csv requirements
no attention needed for basic data types (bool, float, int, string).

array items in the csv file **must** be separated with `|` (ex. `"item_one|item_two"`), as it prevents array assignment with comma-containing strings.

## usage
```
go run jsonify input_file.csv
```
`input_file.json` (output) is saved to same directory as `input_file.csv`
