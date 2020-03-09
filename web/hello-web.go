package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    PORT = ":8888"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static.html")
}

func serveDynamic(w http.ResponseWriter, r *http.Request) {
    response := "The time is now " + time.Now().String()
    fmt.Fprintln(w, response)
}

func main() {
    http.HandleFunc("/static", serveStatic)
    http.HandleFunc("/dynamic", serveDynamic)
    http.ListenAndServe(PORT, nil)
}
