package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
    if len(os.Args) == 1 {
        fmt.Println("File name required as argument !!")
        os.Exit(1)
    }

    file, err := os.Open(string(os.Args[1]))
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    io.Copy(os.Stdout, file)
}
