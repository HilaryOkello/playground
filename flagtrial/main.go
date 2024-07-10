package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type customFlagValue string

func (c *customFlagValue) String() string {
	return string(*c)
}

func (c *customFlagValue) Set(value string) error {
	*c = customFlagValue(value)
	return nil
}

func main() {
	var color string

	flag.StringVar(&color, "color", "black", "Set the color")

	flag.Parse()

	color =parseColor(color)

	fmt.Println("Color:", color)
}

// Predefined ANSI color codes
var ansiCodes = map[string]string{
	"black":          "30",
	"red":            "31",
	"green":          "32",
	"yellow":         "33",
	"blue":           "34",
	"magenta":        "35",
	"cyan":           "36",
	"white":          "37",
	"gray":           "90",
	"bright red":     "91",
	"bright green":   "92",
	"bright yellow":  "93",
	"bright blue":    "94",
	"bright magenta": "95",
	"bright cyan":    "96",
	"bright white":   "97",
	"purple":         "38;5;128", // 256-color mode
	"orange":         "38;5;208", // 256-color mode
	"brown":          "38;5;94",  // 256-color mode
	"pink":           "38;5;218", // 256-color mode
	"light gray":     "38;5;250", // 256-color mode
	"dark gray":      "38;5;238", // 256-color mode
	"light red":      "38;5;203", // 256-color mode
	"light green":    "38;5;120", // 256-color mode
	"light yellow":   "38;5;229", // 256-color mode
	"light blue":     "38;5;153", // 256-color mode
	"light magenta":  "38;5;207", // 256-color mode
	"light cyan":     "38;5;159", // 256-color mode
	"light white":    "38;5;231", // 256-color mode
}

// Converts a hex color to its ANSI escape code in 24-bit format
func hexToANSI(hex string) (string, error) {
	// Remove the '#' character
	hex = strings.TrimPrefix(hex, "#")
	// Parse the hex values
	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return "", err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return "", err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("38;2;%d;%d;%d", r, g, b), nil
}

// Converts RGB values to the ANSI escape code in 24-bit format
func rgbToANSI(r, g, b int) string {
	return fmt.Sprintf("38;2;%d;%d;%d", r, g, b)
}

func parseColor(color string) string {
	if strings.HasPrefix(color, "#") {
		code, err := hexToANSI(color)
		if err == nil {
			return code
		}
	} else if strings.HasPrefix(color, "rgb") {
		parts := strings.Split(strings.TrimPrefix(color, "rgb("), ",")
		if len(parts) == 3 {
			r, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			g, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			b, err3 := strconv.Atoi(strings.TrimSuffix(strings.TrimSpace(parts[2]), ")"))
			if err1 == nil && err2 == nil && err3 == nil {
				return rgbToANSI(r, g, b)
			}
		}
	} else {
		if code, exists := ansiCodes[color]; exists {
			return code
		}
	}
	return ""
}
