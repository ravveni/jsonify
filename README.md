```
   _                 _  __       
  (_)               (_)/ _|      
   _ ___  ___  _ __  _| |_ _   _ 
  | / __|/ _ \| '_ \| |  _| | | |
  | \__ \ (_) | | | | | | | |_| |
  | |___/\___/|_| |_|_|_|  \__, |
 _/ |                       __/ |
|__/                       |___/ 

a simple csv to json converter
```

## Pre-Requirements
- [Go 1.22.5+](https://go.dev/doc/install)

## CSV Preparation
- First row **must** contain only string headers/keys
- No extra attention needed for most basic data types (float, int, string)
- Bool values must be either `TRUE` or `FALSE` (0 and 1 are used in int conversion)
- Array items in csv cells **must** be separated with `|` (ex. `"item_one|item_two"`), as it prevents array assignment with comma-containing strings.

## Usage
### Basic
```
$ go run jsonify input_file.csv
```
`input_file.json` (output) is saved to same directory as `input_file.csv`

### Compiled
```
$ go build
```
Create a symlink from the executable to your PATH for use system-wide
```
$ sudo ln -s absolute/path/to/jsonify /usr/local/bin/jsonify
```

---

Example csv by [datablist/sample-csv-files](https://github.com/datablist/sample-csv-files)
