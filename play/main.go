package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	str, color, toBeColored, r, g, b, fileName := processArgs(os.Args)
	colors := map[string]string{
		"reset":   "\033[0m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"gray":    "\033[37m",
		"white":   "\033[97m",
	}

	if toBeColored != "" {
		startIndex := strings.Index(str, toBeColored)
		lastIndex := startIndex + len(toBeColored)
		for index, char := range str {
			if index >= startIndex && index <= lastIndex {
				if color != "" {
					fmt.Printf("%s%c\033[0m", colors[color], char)
				} else {
					fmt.Printf("\033[38;2;%s;%s;%sm%c\033[0m", r, g, b, char)
				}
			} else {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()

	} else {
		if color != "" {
			fmt.Printf("%s%s\n\033[0m", colors[color], str)
		} else if r != "" {
			fmt.Printf("\033[38;2;%s;%s;%sm%s\n\033[0m", r, g, b, str)
		} else {
			fmt.Println(str)
		}
	}

	fmt.Println(fileName)
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
			fileName = os.Args[6]

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
