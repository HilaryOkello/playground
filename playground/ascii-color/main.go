package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"

	"ascii-art/ascii"
)

// The main function checks if the length of arguments passed on Terminal
// is between two and three. If so , then value of os.args[1] is stored in str.
// It then checks if all the characters of the string at argument 1 are printable
// ascii characters. If not, the error is displayed and program terminates.
// It again checks if the second argument passed,which represents the banner filename,
// is among  the files in the banner banner directory and throws an error if it's not.
// If fileName is banner, it again checks if the file passed is either empty or it's length tampered.
// if the filename is "thinkertoy", "\r\n "are replaced with "\n" and content is split as from index 2:
// Other file is from 1: because of the newline at the index 0
// We do some replacement to deal with special characters
// We split str splitted using "\\n" to get words
// While looping through words, count value is initialized to 0 to track empty strings, and call
// PrintAscii()for every str in words
func main() {
	lenArgs := len(os.Args)
	if lenArgs < 2 || lenArgs > 7 {
		fmt.Printf("Incorrect no. of arguments.\n" +
			"Expects: \"go run . <string> | cat -e\"\n" +
			"or\n" +
			"\"go run . <string> <banner name> | cat -e\"\n")
		return
	}
	str, color, toBeColored, r, g, b, fileName := processArgs(os.Args)

	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	err := ascii.IsPrintableAscii(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	errFile := ascii.CheckFileValidity("./banner", fileName)
	if errFile != nil {
		fmt.Println(errFile)
		return
	}

	filePath := os.DirFS("./banner")
	contentByte, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(contentByte) == 0 {
		fmt.Println("Banner file is empty")
		return
	}
	er := ascii.CheckFileTamper(fileName, contentByte)
	if er != nil {
		fmt.Println(er)
		return
	}

	contentString := string(contentByte[1:])
	if fileName == "thinkertoy.txt" {
		contentString = strings.ReplaceAll(string(contentByte[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")
	words := strings.Split(str, "\\n")
	count := 0
	w := false
	for _, str := range words {
		if strings.Index(str, toBeColored) != -1 {
			w = true
		} else{
			w = false
		}
		if str == "" {
			count++
			if count < len(words) {
				fmt.Println()
			}
		} else {
			ascii.PrintAscii(str, contentSlice, 0, color, toBeColored, r, g, b, w)
		}
	}
}

func processArgs(args []string) (str, color, toBeColored, r, g, b, fileName string) {
	lenArgs := len(args)
	fileName = "standard.txt"
	// go run . "Hello There"
	if lenArgs == 2 {
		str = os.Args[1]
		if strings.HasPrefix(str, "--color=") {
			log.Fatal("missing string")
		}
	}
	if lenArgs > 2 && strings.HasPrefix(os.Args[1], "--color=") {
		switch lenArgs {
		// go run . --color=#666AB1 "Hello There"
		case 3:
			str = os.Args[2]
			color = os.Args[1]
		case 4:
			// go run . --color=#666AB1 There "Hello There"
			if strings.Contains(os.Args[3], os.Args[2]) {
				color = os.Args[1]
				toBeColored = os.Args[2]
				str = os.Args[3]
			} else { // go run . --color=#666AB1 "Hello There" thinkertoy
				color = os.Args[1]
				str = os.Args[2]
				fileName = os.Args[3] + ".txt"
			}
		case 5:
			// go run . --color=rgb\(255 255 0\) "Hello There thinkertoy"
			if strings.HasPrefix(os.Args[1], "--color=rgb") {
				r, g, b = os.Args[1][12:], os.Args[2], os.Args[3][:len(os.Args[3])-1]
				str = os.Args[4]
			} else { // go run . --color=#666AB1 "Hello There" thinkertoy
				color = os.Args[1]
				toBeColored = os.Args[2]
				str = os.Args[3]
				fileName = os.Args[4] + ".txt"
			}
		case 6:
			// go run . --color=rgb\(255 255 0\) There "Hello There"
			if strings.Contains(os.Args[5], os.Args[4]) {
				r, g, b = os.Args[1][12:], os.Args[2], os.Args[3][:len(os.Args[3])-1]
				toBeColored = os.Args[4]
				str = os.Args[5]

			} else { // go run . --color=rgb\(255 255 0\) "Hello There" thinkertoy
				r, g, b = os.Args[1][12:], os.Args[2], os.Args[3][:len(os.Args[3])-1]
				str = os.Args[4]
				fileName = os.Args[5] + ".txt"
			}
		case 7: // go run . --color=rgb\(255 255 0\) There "Hello There" thinkertoy
			r, g, b = os.Args[1][12:], os.Args[2], os.Args[3][:len(os.Args[3])-1]
			toBeColored = os.Args[4]
			str = os.Args[5]
			fileName = os.Args[6] + ".txt"

		}
	} else if lenArgs == 3 { // go run . "Hello There" thinkertoy
		str = os.Args[1]
		fileName = os.Args[2] + ".txt"
	} else {
		str = os.Args[1]
	}
	var hex string

	// Extract hex from --color=#666AB1"
	if strings.HasPrefix(color, "--color=#") {
		hex = color[9:]
		rgbSlice, err := convertHex2RGB(hex)
		if err != nil {
			fmt.Println(err)
			return
		}
		r, g, b = rgbSlice[0], rgbSlice[1], rgbSlice[2]
		color = ""
	} else if strings.HasPrefix(color, "--color=") {
		color = color[8:]
	}

	return
}

func convertHex2RGB(h string) ([]string, error) {
	if len(h) != 6 && len(h) != 6 {
		return nil, fmt.Errorf("color hex can only be made up of 3 0r 6 hex characters")
	}
	step := 1
	if len(h) == 6 {
		step = 2
	}
	var result []string
	for i := 0; i < len(h); i = i + step {
		subString := h[i:min(i+step, len(h))]
		dec, err := strconv.ParseInt(subString, 16, 32)
		if err != nil {
			return nil, err
		}
		result = append(result, fmt.Sprint(dec))

	}
	return result, nil
}
