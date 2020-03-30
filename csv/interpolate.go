package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
)

func main() {	
	// There must be InputCsvFile, FormatFile and OutputFile alongwith command itself.
	if (len(os.Args) != 4) {
		fmt.Printf("ERROR :: Too Few Arguments \n")
		os.Exit(1000);
	}

	// Now Check if input file is file and can open it
	inputFile, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("ERROR :: Cant Read Input File\n")
		os.Exit(1001);
	}

	// Now Check if input file is file and can open it
	formatFile, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Printf("ERROR :: Cant Read Format File\n")
		os.Exit(1002);
	}
	formatText := string(formatFile)


	// Now Check if input file is file and can open it
	outputFile, err := os.OpenFile(os.Args[3], os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("ERROR :: Cant Open Output File for Writing\n")
		inputFile.Close()
		os.Exit(1003);
	}
	createFile(inputFile,formatText,outputFile)
	

	

	inputFile.Close()
	outputFile.Close()
}

func createFile(csvFile *os.File, format string, resultFile *os.File ) {
	
	csvReader := csv.NewReader(csvFile)
	totalLines := 0;
	for {

		// Read A reccord of line from CSV File
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Read Error in CSV Inputfile!!!")
			os.Exit(1001);
		}

		// Set output text as initital format text. We will replace the variable in each pass.
		outputText := format;
		for i,value := range record {
			searchString := "{" + strconv.Itoa(i) + "}"
			outputText = strings.ReplaceAll(outputText,searchString,value)
		}
		outputText += "\n"
		
		io.WriteString(resultFile,outputText)

		log.Printf("Output for %v is %v \n",record,string(outputText))
		totalLines++;
	}
	log.Printf("%v Record Processed and put in %v.\n",totalLines,resultFile.Name())
}
