package ascii

import (
	"fmt"
	"strings"
)

func PrintAscii(str string, contentSlice []string, index int, color string, toBeColored string, rgb []string, tbcInStr bool) {
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
		var r, g, b string
		if len(rgb) == 3 {
			r, g, b = string(rgb[0]), string(rgb[1]), string(rgb[2])
		}

		if tbcInStr {
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
			} else if len(rgb) == 3 && toBeColored == "" {
				fmt.Printf("\033[38;2;%s;%s;%sm%s\033[0m", r, g, b, lines[index])
			} else {
				fmt.Print(lines[index])
			}
		}
	}

	fmt.Println()
	PrintAscii(str, contentSlice, index+1, color, toBeColored, rgb, tbcInStr)
}
