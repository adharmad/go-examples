package main

import (
    "fmt"
    "strings"
)

func main() {
    // string.Contains(str, substr string) bool
    fmt.Println(strings.Contains("Hello, World!", "Hell"))

    // count number of occurances of substrings
    fmt.Println("Number of times space occurs in the string is: ", strings.Count("This is a test", " "))

    // check if a string begins with another string
    fmt.Println(strings.HasPrefix("Pax Vobiscum", "Pa"))

    // check if a string ends with another string
    fmt.Println(strings.HasSuffix("Pax Vobiscum", "scum"))

    // index of
    fmt.Println("Index of substring is: ", strings.Index("Play Framework", "work"))

    // Replace a substring of a string
    //strings.Replace(str, old, new, num)
    fmt.Println(strings.Replace("abracadabra", "ab", "xx", 2))

    // concat multiple strings to form a new one interspersed with separator
    s := strings.Join([]string{"Hello","World", "Join", "These"}, "-")
    fmt.Println(s)

    // split string by separator
    fmt.Printf("%q\n", strings.Split("a,b,c", ","))

}
