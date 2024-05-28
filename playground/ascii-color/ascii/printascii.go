package ascii

import (
	"fmt"
	"strings"
)

// function takes str which is string passed at argument one ,contentslice which is filename that  is sliced,and index which is the lenght of character in that string
// if the index value(lenght of each character in a string) is not equal to 8  the printing of each line of character continues,else loop stop and program is terminated
// while looping over each character in string  we subtract 32 from  the character location in the contentslice which is the sliced value of content in filename used
// after obtaining the character,character is splitted using line inorder to print line by line then after printing each line ,we print a new line to separate each line
func PrintAscii(str string, contentSlice []string, index int, color string, toBeColored string, r string, g string, b string, w bool) {
	if index == 8 {
		return
	}

	for i, char := range str {
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
		character := contentSlice[int(char)-32]
		character = strings.ReplaceAll(character, "\r\n", "\n")
		lines := strings.Split(character, "\n")
		if toBeColored != "" && w {
			startIndex := strings.Index(str, toBeColored)
			lastIndex := startIndex + len(toBeColored) - 1
			if i >= startIndex && i <= lastIndex {
				if color != "" {
					fmt.Printf("%s%s\033[0m", colors[color], lines[index])
				} else {
					fmt.Printf("\033[38;2;%s;%s;%sm%s\033[0m", r, g, b, lines[index])
				}
			} else {
				fmt.Print(lines[index])
			}
		} else {
			if color != "" && toBeColored == "" {
				fmt.Printf("%s%s\033[0m", colors[color], lines[index])
			} else if r != "" {
				fmt.Printf("\033[38;2;%s;%s;%sm%s\033[0m", r, g, b, lines[index])
			} else {
				fmt.Print(lines[index])
			}
		}

	}
	fmt.Println()
	PrintAscii(str, contentSlice, index+1, color, toBeColored, r, g, b, w)
}
