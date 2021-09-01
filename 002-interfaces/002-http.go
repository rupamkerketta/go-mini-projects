package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
    res, err := http.Get("http://google.com")
    if err != nil {
        log.Fatal("Error:", err)
        os.Exit(1)
    }

    io.Copy(os.Stdout, res.Body)
    fmt.Println("HTTP request")
}
