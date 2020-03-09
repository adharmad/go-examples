package main

// Shows dynamic routing using gorilla/mux package
// http://localhost:8888/pages/1
// http://localhost:8888/pages/2

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

const (
    PORT = ":8888"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Println("Vars = ", vars)
    
    pageID := vars["id"]
    fmt.Println("Page ID = " + pageID)

    fileName := "pages/" + pageID + ".html"
    fmt.Println("File name = " + fileName)
   
    http.ServeFile(w, r, fileName)

}


func main() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
    http.Handle("/", rtr)
    http.ListenAndServe(PORT, nil)

}
