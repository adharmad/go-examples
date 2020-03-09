package main

import (
    "fmt"
    "crypto/sha1"
)

func main() {
    h := sha1.New()
    h.Write([]byte("the quick brown fox jumped over the lazy dog"))
    bs := h.Sum([]byte{})
    fmt.Printf("%x", bs)
}
