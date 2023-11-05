package main

import (
	"os"
	"fmt"
	"github.com/harkaitz/go-parse-monetary"
	"github.com/pborman/getopt/v2"
)

const help string =
`Usage: parse-monetary [-a] -- <money>

Parse a monetary value from the command line arguments. The value
is expected to be a string in the spanish format, with a comma or
a dot as the decimal or thousands separator.

The return value is always in cents, if "-a" is set it prints
the absolute value.

Copyright (c) 2023 - Harkaitz Agirre - MIT License`

func main() {
	
	hFlag := getopt.BoolLong("help", 'h', "Show help")
	aFlag := getopt.BoolLong("absolute", 'a', "Print the absolute value")
	getopt.SetUsage(func () { fmt.Println(help) })
	getopt.Parse()
	if *hFlag || len(getopt.Args()) != 1 { getopt.Usage(); return }
	
	value, negative, err := money.Parse(getopt.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	
	if !negative || *aFlag {
		fmt.Printf("%d\n", value)
	} else {
		fmt.Printf("-%d\n", value)
	}
}
