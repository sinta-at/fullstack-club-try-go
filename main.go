package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting HTTP server...")

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        textResponse := "Hello, world!"
        w.Write([]byte(textResponse))
    })

    http.HandleFunc("/foobar", func(w http.ResponseWriter, r *http.Request) {
        mapResponse := map[string]string{
            "foo": "bar",
        }
        responseBody, err := json.Marshal(mapResponse)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        w.Header().Add("Content-Type", "application/json")
        w.Write(responseBody)
    })

    http.ListenAndServe(":80", nil)
}