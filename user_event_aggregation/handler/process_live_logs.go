package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

// processLiveLogs is used to process the live upcoming logs input file and update aggregation into output file every time.
func processLiveLogs(inputFileName, outputFileName string) {
	fmt.Println("Process data from live upcoming log file")

	// For simplicity i am reading lime by line from file and parsing it into Input
	// Might be case where json is written into miltiple line then we start reading s string from "{"
	// will end until not received "}" ans consider one jons from { to }.
	t, err := tail.TailFile("../"+inputFileName, tail.Config{Follow: true})
	fmt.Println("starting reading")
	if err != nil {
		fmt.Println("error while reading data from file err : ", err)
	}
	var input *Input
	for line := range t.Lines {
		input = validateInputParseInputData(line.Text)
		if input != nil {
			updateOutputStream(input)
			writeDataInOutputFile(outputFileName)
		}
	}
	fmt.Println("Output file generated successfully")
}
