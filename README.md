# ASCII-ART

A Go program that converts text strings into ASCII art using predefined banner templates.

## Description

ASCII-ART receives a string as an argument and outputs a graphical representation of the text using ASCII characters. Each character is rendered as an 8-line tall ASCII art representation based on banner template files.

## Author

**Team Project**

## Project Structure
```
ASCII-ART/
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── MethodsAndTesting/
│   ├── file-handling.go
│   ├── printer.go
│   └── printer_test.go
├── main.go
├── go.mod
└── README.md
```

## Features

- Converts text to ASCII art using banner templates
- Supports all printable ASCII characters (32-126)
- Handles newline characters (`\n`)
- Multiple banner styles available (standard, shadow, thinkertoy)
- Comprehensive unit tests
- Clean modular code structure

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd ASCII-ART
```

2. Ensure you have Go installed (version 1.16 or higher):
```bash
go version
```

3. Initialize the module (if not already done):
```bash
go mod init ascii-art
```

## Usage

### Basic Usage
```bash
go run . "Your text here"
```

### Examples

**Simple text:**
```bash
go run . "hello"
```

**Output:**
```
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
```

**Uppercase text:**
```bash
go run . "HELLO"
```

**With newlines:**
```bash
go run . "Hello\nWorld"
```

**Output:**
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
__          __                 _       _ 
\ \        / /                | |     | |
 \ \  /\  / /    ___    _ __  | |   __| |
  \ \/  \/ /    / _ \  | '__| | |  / _` |
   \  /\  /    | (_) | | |    | | | (_| |
    \/  \/      \___/  |_|    |_|  \__,_|
```

**Special characters:**
```bash
go run . "[\]^_ 'a"
```

**Numbers and symbols:**
```bash
go run . "Hello123!"
```

## Supported Characters

The program supports all printable ASCII characters from space (32) to tilde (126):
- Lowercase letters: `a-z`
- Uppercase letters: `A-Z`
- Numbers: `0-9`
- Special characters: `!@#$%^&*()_+-=[]{}|;:'",.<>?/~` and space

## Banner Format

Each banner file contains ASCII representations of characters where:
- Each character is exactly **8 lines tall**
- Characters are separated by a **newline** (9 lines total per character)
- Characters are ordered by ASCII value (32-126)
- Banner files are stored in the `banners/` directory

## Code Structure

### `main.go`
Entry point of the application. Handles command-line arguments and calls the formatter.

### `MethodsAndTesting/file-handling.go`
Contains `FileHandler()` function that reads the banner template file.

**Key Function:**
```go
func FileHandler() ([]byte, bool)
```
- Reads `banners/standard.txt`
- Returns file contents and success status
- Handles file reading errors

### `MethodsAndTesting/printer.go`
Contains `FormatPrinter()` function that processes the input string and generates ASCII art.

**Key Function:**
```go
func FormatPrinter(input string) string
```
- Splits input by `\n` to handle multiple lines
- Maps each character to its ASCII art representation
- Builds output line by line (8 rows per text line)
- Returns formatted ASCII art string

**Algorithm:**
1. Load banner file using `FileHandler()`
2. Split banner content into lines
3. Split input string by `\n`
4. For each word/line:
   - If empty, add a newline
   - For each of 8 rows:
     - For each character, find its representation in banner
     - Calculate line index: `(ASCII_value - 32) * 9 + 1 + row`
     - Append character's row to output
5. Trim final newline

### `MethodsAndTesting/printer_test.go`
Unit tests for the `FormatPrinter()` function.

**Test Cases:**
- Simple text (`"hello"`)
- Special characters (`"[\]^_ 'a"`)
- Newline handling (`"hello\n"`)
- Multiple newlines (`"hello\n\n"`)

## Running Tests

Execute all tests:
```bash
go test ./MethodsAndTesting
```

Run tests with verbose output:
```bash
go test -v ./MethodsAndTesting
```

Run specific test:
```bash
go test -run TestFormatPrinter ./MethodsAndTesting
```

## Error Handling

The program handles the following errors:

1. **Incorrect number of arguments:**
```
   error: enter 2 arguments
```

2. **Banner file not found:**
```
   Error
```
   (Program terminates with fatal error)

3. **Empty input:**
   Program returns without output

4. **Invalid characters:**
   Characters outside ASCII 32-126 are silently skipped

## Implementation Details

### Character Mapping
Each ASCII character (32-126) occupies 9 lines in the banner file:
- Lines 1-8: ASCII art representation
- Line 9: Empty separator

**Formula:**
```
lineIndex = (charASCII - 32) * 9 + 1 + rowNumber
```

Where:
- `charASCII`: ASCII value of character
- `rowNumber`: Current row (0-7)
- `32`: ASCII value of space (first printable character)

### Newline Handling
Input string `"Hello\nWorld"` is split into:
```
["Hello", "World"]
```
Each segment is processed separately with 8-line output.

Empty segments (from `"\n\n"`) produce single newline in output.

## Technical Requirements

- **Language:** Go (Golang)
- **Go Version:** 1.16 or higher
- **Packages:** Standard library only
  - `os`: File operations and command-line arguments
  - `fmt`: Formatted I/O
  - `strings`: String manipulation
  - `log`: Error logging
  - `testing`: Unit tests

## Good Practices

✅ **Modular code** - Functions separated by responsibility  
✅ **Unit tests** - Comprehensive test coverage  
✅ **Error handling** - Proper error checking and reporting  
✅ **Clean code** - Readable and maintainable  
✅ **No external dependencies** - Uses only standard library  

## Learning Objectives

This project helps you learn about:
- Go file system (fs) API
- Data manipulation in Go
- String processing
- ASCII character encoding
- Command-line argument handling
- Unit testing in Go
- Code organization and modularity

## Future Enhancements

Potential improvements:
- [ ] Support for multiple banner styles (shadow, thinkertoy)
- [ ] Custom banner file selection via command-line flag
- [ ] Color output support
- [ ] Export to file option
- [ ] UTF-8 character support
- [ ] Interactive mode

## License

This project is part of an educational program.

## Troubleshooting

**Problem:** `Error` message when running
- **Solution:** Ensure `banners/standard.txt` exists in the correct location

**Problem:** Wrong output format
- **Solution:** Verify banner file format (8 lines per character + 1 separator)

**Problem:** Tests failing
- **Solution:** Run from project root directory where `banners/` folder is accessible

**Problem:** Characters not displaying
- **Solution:** Check if character is in printable ASCII range (32-126)

## Contact

For questions or issues, please contact the development team.

---

**Note:** This is an educational project demonstrating ASCII art generation and Go programming concepts.