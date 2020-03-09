package main

import (
    "net/http"
)

const (
    PORT = ":8888"
)

func main() {
    http.ListenAndServe(PORT, http.FileServer(http.Dir("/Users/adharmad")))
}
