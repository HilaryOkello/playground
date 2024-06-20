package main

import (
	"fmt"
	"log"
	"strconv"
)

func convertHex2RGB(h string) []string {
	step := 1
	if len(h) == 6 {
		step = 2
	}
	var result []string
	for i := 0; i < len(h); i = i + step {
		subString := h[i:min(i+step, len(h))]
		if len(subString) == 1 {
			subString += subString // Double the character
		}
		dec, err := strconv.ParseInt(subString, 16, 32)
		if err != nil {
			log.Fatalf("Error converting %s to rgb", h)
		}
		result = append(result, fmt.Sprint(dec))

	}
	return result
}
