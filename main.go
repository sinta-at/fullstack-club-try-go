package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting HTTP server...")

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        textResponse := "Hello, world!"
        w.Write([]byte(textResponse))
    })

    http.ListenAndServe(":80", nil)
}