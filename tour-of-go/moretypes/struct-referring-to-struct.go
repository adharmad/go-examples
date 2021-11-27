package main

import "fmt"

type Name struct {
    first string
    middle string
    last string
}

type Email struct {
    email string
    emailType string
    isPrimary bool
    isVerified bool
}

type User struct {
    name *Name
    email *Email
    age int
    ssn string
}

func main() {
    user := make_user("Richard", "P", "Feynman", "rpf@caltech.edu", 137, "999-999-9999")
    fmt.Println(user)
    fmt.Println(user.name)
    fmt.Println(user.email)
}

func make_user(first1 string, middle1 string, last1 string, email1 string, age int, ssn string) *User {
    name := Name{first: first1, last: last1, middle: middle1}
    email := Email{email: email1, emailType: "work", isPrimary: true, isVerified: false}

    user := new(User)
    user.name = &name
    user.email = &email
    user.age = age
    user.ssn = ssn

    return user
}
