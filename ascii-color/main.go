package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"ascii-art/ascii"
)

func main() {
	lenArgs := len(os.Args)
	fileName := "standard.txt"
	var str string
	var toBeColored string

	var color colorFlagValue
	flag.Var(&color, "color", "Set the color")
	flag.Parse()
	c, _ := color.Color.(string)
	rgb, _ := color.Color.([]string)
	if strings.HasPrefix(c, "#") {
		rgb = convertHex2RGB(c[1:])
		c = ""
	}

	usage := `
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --color=<color> <letters to be colored> "something" <banner>
`
	if lenArgs < 2 || lenArgs > 5 {
		log.Fatal(usage)
	}

	if strings.HasPrefix(os.Args[1], "--color=") {
		switch lenArgs {
		case 2:
			log.Fatal(usage)
		case 3:
			str = os.Args[2]
		case 4:
			if strings.Contains("standard thinkertoy shadow", os.Args[3]) {
				fileName = os.Args[3] + ".txt"
				str = os.Args[2]
			} else {
				toBeColored = os.Args[2]
				str = os.Args[3]
			}
		case 5:
			toBeColored = os.Args[2]
			str = os.Args[3]
			fileName = os.Args[4] + ".txt"
		}
	} else {
		switch lenArgs {
		case 2:
			str = os.Args[1]
		case 3:
			str = os.Args[1]
			fileName = os.Args[2] + ".txt"
		default:
			log.Fatalf(usage)
		}
	}

	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	err := ascii.IsPrintableAscii(str)
	if err != nil {
		log.Fatal(err)
	}

	filePath := os.DirFS("./banner")
	contentByte, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		log.Fatal(err)
	}
	if len(contentByte) == 0 {
		log.Fatal("Banner file is empty")
	}
	er := ascii.CheckFileTamper(fileName, contentByte)
	if er != nil {
		log.Fatal(er)
	}

	contentString := string(contentByte[1:])
	if fileName == "thinkertoy.txt" {
		contentString = strings.ReplaceAll(string(contentByte[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")
	words := strings.Split(str, "\\n")
	count := 0
	for _, str := range words {
		tbcInStr := false
		if strings.Index(str, toBeColored) != -1 {
			tbcInStr = true
		}
		if str == "" {
			count++
			if count < len(words) {
				fmt.Println()
			}
		} else {
			ascii.PrintAscii(str, contentSlice, 0, c, toBeColored, rgb, tbcInStr)
		}
	}
}
