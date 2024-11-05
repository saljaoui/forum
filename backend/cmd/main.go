package main 

import (
    "fmt"
    "net/http"
    "io"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", getRoot)
    
    err := http.ListenAndServe(":3333", mux)
    fmt.Println(err)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request\n")
    io.WriteString(w, "This is my website!\n")
}