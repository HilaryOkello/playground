package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type colorFlagValue struct {
	Color interface{}
}

func (c *colorFlagValue) String() string {
	switch v := c.Color.(type) {
	case string:
		return v
	case []string:
		return fmt.Sprintf("%v", v)
	default:
		return fmt.Sprintf("Unsupported color type: %T", v)
	}
}

func (c *colorFlagValue) Set(value string) error {
	rgbRegex := regexp.MustCompile(`^rgb\((\d+)\s+(\d+)\s+(\d+)\)$`)
	hexRegex := regexp.MustCompile(`^#[0-9a-eA-F]{6}|[0-9a-fA-F]{3}$`)

	if strings.HasPrefix(value, "rgb") {
		matches := rgbRegex.FindStringSubmatch(value)
		if len(matches) == 4 {
			c.Color = []string{matches[1], matches[2], matches[3]}
		}
	} else if hexRegex.MatchString(value) {
		c.Color = (value)
	} else if strings.HasPrefix(value, "#") {
		log.Fatalf("Invalid hex: %s", value)
	} else {
		c.Color = value
	}
	return nil
}
