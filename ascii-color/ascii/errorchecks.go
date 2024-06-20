package ascii

import (
	"fmt"
	"strings"
)

func IsPrintableAscii(str string) error {
	var nonPrintables string
	var foundEscapes string
	errMessage := ": Not within the printable ascii range"
	for index, char := range str {

		escapes := "avrfb"
		var next byte
		if index < len(str)-1 {
			next = str[index+1]
		}

		NextIsAnEscapeLetter := strings.ContainsAny(string(next), escapes)
		isAnEscape := (char == '\\' && NextIsAnEscapeLetter)
		isNonPrintable := ((char < ' ' || char > '~') && char != '\n')

		if isAnEscape {
			foundEscapes += "\\" + string(next)
		}
		if isNonPrintable {
			nonPrintables += string(char)
		}
	}

	if foundEscapes != "" {
		return fmt.Errorf("%s%s", foundEscapes, errMessage)
	} else if nonPrintables != "" {
		return fmt.Errorf("%s%s", nonPrintables, errMessage)
	}
	return nil
}

func CheckFileTamper(fileName string, content []byte) error {
	errMessage := " is tampered"
	lengthContent := len(content)

	if fileName == "standard.txt" && lengthContent != 6623 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	} else if fileName == "thinkertoy.txt" && lengthContent != 5558 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	} else if fileName == "shadow.txt" && lengthContent != 7465 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	}

	return nil
}
