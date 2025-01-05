package main

import (
	"fmt"
	"github.com/phentrox/db-pg-trunc/internal/help"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			help.PrintHelp()
		case "-v", "--version":
			println("0.1")
		case "--init":
			runInit()
		case "--trunc":
			runTruncate()
		default:
			fmt.Println("unknown arg: " + arg)
		}
	}
}
