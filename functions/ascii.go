package asciiwebkood

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AsciiProgram(input string, style string) string {

	if input == "" {
		return " "
	} else if input == "\n" {
		fmt.Println() // Print a newline if the argument is "\n"
		return " "
	}

	var asciiData *os.File
	var err error

	switch style {
	case "shadow":
		asciiData, err = os.Open("shadow.txt")
	case "standard":
		asciiData, err = os.Open("standard.txt")
	case "thinkertoy":
		asciiData, err = os.Open("thinkertoy.txt")
	default:
		fmt.Printf("Unvalid banner argument\n")
		return " "
	}

	if err != nil { // Error check in case standard.txt is not available
		fmt.Println("Error accessing the file")
		panic(err)
	}
	lineArr := make(map[int][]string)      // Creation of map which will have int as keys and a default data of an array of strings
	scanner := bufio.NewScanner(asciiData) // Reading of every newline
	lineNum := 0                           // Used for the next for loop

	for scanner.Scan() { // For every iteration of the scan...
		line := scanner.Text()           // Gathering of line for splitting
		parts := strings.Split(line, "") // Splitting into characters inside the array
		lineNum++                        // Going down a line

		lineArr[lineNum] = parts // Addition of keys and values into lineMap(array) map
	}

	// * Conversion

	split := strings.Split(input, "\\n") // Split

	var result strings.Builder // This will provide a template of ASCII values to reflect on the index

	for i, word := range split {
		if i > 0 {
			result.WriteString("\n") // Add a newline before each non-empty word (after the first word)
		}
		if word != "" { // Check if the value is empty (space)
			for i := 0; i < 8; i++ { // Setting the limit on 8, while that's the height of the ASCII characters
				for _, lint := range word {
					lines, exists := lineArr[int((lint)-' ')*9+2+i]
					/* 32 is the value of space in the ASCII table, which is the one we start with in the text files, we multiply for 9 (amount of lines per character + space line which will add up for every iteration), exists will return a boolean value if the value exists or not, if it does, it will print the line
					 */
					if exists {
						for _, line := range lines {
							result.WriteString(line)
						}
					}
				}
				result.WriteString("\n") // Newline for each iteration for structuring the word, otherwise, it would be a straight line
			}
		} else {
			result.WriteString("\n") // Newline
		}
	}
	return result.String()
}
