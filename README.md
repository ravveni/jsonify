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

## pre-requirements
- [golang 1.22.5+](https://go.dev/doc/install)

## csv preparation
- first row **must** contain only string headers/keys
- no extra attention needed for most basic data types (float, int, string)
- bool values must be either `TRUE` or `FALSE` (0 and 1 are used in int conversion)
- array items in csv cells **must** be separated with `|` (ex. `"item_one|item_two"`), as it prevents array assignment with comma-containing strings.

## usage
### basic
```
$ git clone https://github.com/ravveni/jsonify
$ cd jsonify
$ go run jsonify input_file.csv
```
`input_file.json` (output) is saved to same directory as `input_file.csv`

### compiled
```
$ git clone https://github.com/ravveni/jsonify
$ cd jsonify
$ go build
```
create a symlink from the executable to your PATH for use system-wide
```
$ sudo ln -s absolute/path/to/jsonify/jsonify /usr/local/bin/jsonify
```
