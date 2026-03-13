# ASCII Art 

## Description

This program prints text in **ASCII Art** using a banner file (`standard.txt`).
Each character in the input string is converted into a larger ASCII representation.

The program reads characters from the command line and renders them using the ASCII patterns stored in the banner file.

Supported characters are **ASCII 32 to 126**.

---

## Project Structure

```
ascii-art/
│
├── main.go        # Main program logic
├── standard.txt   # ASCII banner file
├── go.mod         # Go module file
└── README.md      # Project documentation
```

---

## How It Works

1. The program reads the input string from the command line.
2. It loads the ASCII banner file (`standard.txt`).
3. Each character corresponds to **8 lines of ASCII art** in the banner.
4. The program prints the characters **row by row** to construct the ASCII output.

---

## Usage

Run the program using:

```
go run . "your text"
```

Example:

```
go run . "Hello"
```

Output:

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
```

---

## Handling New Lines

The program supports newline characters using `\n`.

Example:

```
go run . "Hello\nThere"
```

Output:

```
(ASCII art for Hello)

(ASCII art for There)
```

Multiple new lines are also supported:

```
go run . "Hello\n\nThere"
```

---

## Edge Cases

The program correctly handles:

* Empty input
* Newline characters
* Spaces between words
* Numbers
* Special characters
* Multiple lines of text

Examples:

```
go run . ""
go run . "\n"
go run . "Hello There"
go run . "1Hello 2There"
go run . "{Hello There}"
```



