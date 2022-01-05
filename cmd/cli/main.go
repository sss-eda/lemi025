package main

import (
	"flag"
	"fmt"
)

var subcommand *string

func init() {
	subcommand = flag.String("sub", "", "Sub-Command for LEMI-025.")

	flag.Parse()
}

func main() {
	fmt.Println(*subcommand)
}
