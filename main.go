package main

import (
	"crawler/colours"
	"crawler/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		colours.PrintRed(err)
	}
}
