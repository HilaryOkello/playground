package main

import (
	"flag"
	"fmt"
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
	
	fmt.Println("Color:", color)
}
