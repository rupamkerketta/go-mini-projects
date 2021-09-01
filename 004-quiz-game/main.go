package main

import (
	"flag"
	"os"
    "fmt"
)

func main() {
    csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
    flag.Parse()

    file, err := os.Open(*csvFilename)
    if err != nil {
        fmt.Printf("Failed to open  the CSV file : %v", *csvFilename)
        os.Exit(1)
    }
    _ = file
}
