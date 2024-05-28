// this package basically has functions which handles errors at different stages
package ascii

import (
	"fmt"
	"os"
	"strings"
)

// the function takes string and returns an error
// the function check if the string passed is printable or has any of the escapeds values,
// and if so the nonasciivalues  are stored in the varriable nonprintable and  the escapes value in founEscapes respectively
// if both founescapes and nonprintable varriable is not empty ,the error message is printed  along with the found values
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
			foundEscapes += string(next)
		}
		if isNonPrintable {
			nonPrintables += string(char)
		}
	}

	if foundEscapes != "" {
		escSlash := ""
		for _, es := range foundEscapes {
			escSlash += "\\" + string(es)
		}
		return fmt.Errorf("%s%s", escSlash, errMessage)
	} else if nonPrintables != "" {
		return fmt.Errorf("%s%s", nonPrintables, errMessage)
	}

	return nil
}

// function takes string and returns an error
// function opens the path containing files/ folder containing file
// then reads names of the files in the directory or path stored and then store them as slice of string
// join slicenames to form one block string
// checks if the filename is among the names on the folder banner and if not error message is printed "not a valid banner file name"
func CheckFileValidity(dirPath, fileName string) error {
	openPath, err := os.Open(dirPath)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer openPath.Close()

	filenames, err := openPath.Readdirnames(0)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	fileNameString := strings.Join(filenames, " ")
	if !strings.Contains(fileNameString, fileName) {
		return fmt.Errorf("%s is not a valid banner style\n"+
			"Try \"standard\", \"shadow\", or \"thinkertoy\"",
			fileName[:len(fileName)-4])
	}
	return nil
}

// function takes string  which is filename , content of filename respectively and then returns an error message
// if the name obtained does not match the expected length respectively, then error message is displayed
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
