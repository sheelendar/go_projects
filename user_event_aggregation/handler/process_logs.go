package main

import (
	"bufio"
	"fmt"
	"os"
)

// processLiveLogs is used to process the file only onece read input file data and write aggregation into output file.
func processLogs(inputFileName, outputFileName string) {

	fmt.Println("Process data from log file")
	readFile, err := os.Open("../" + inputFileName)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fmt.Println("starting reading")

	var input *Input
	// we can also read whole file once and ummarshal into Array[]Input
	// For simplicity i am reading lime by line from file and parsing it into Input
	// Might be case where json is written into miltiple line then we start reading s string from "{"
	// will end until not received "}" ans consider one jons from { to }.
	for fileScanner.Scan() {
		input = validateInputParseInputData(fileScanner.Text())
		if input != nil {
			updateOutputStream(input)
		}
	}
	writeDataInOutputFile(outputFileName)
	fmt.Println("Output file generated successfully")
}
