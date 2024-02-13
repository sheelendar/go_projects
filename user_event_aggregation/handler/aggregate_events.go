package main

import (
	"fmt"
	"os"
	"strings"
)

// streamOutput to track record of each user on daily basis
var streamOutput = make(map[string]Output)

// maintain user order as their logs comes into system
var userKeyOrder []string

func main() {

	args := os.Args
	size := len(args)
	if size < 5 {
		fmt.Println("Please provide command line agrs input file and output file names")
		return
	}
	// take all args from command line and put these for further use.
	inputFileName := args[2]
	outputFileName := args[4]
	if size == 6 && strings.ContainsAny(args[size-1], "update") {
		processLiveLogs(inputFileName, outputFileName)
	} else {
		processLogs(inputFileName, outputFileName)
	}
}
