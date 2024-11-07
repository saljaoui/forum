package main 

import (
    "fmt"
    "net/http"
    "forum-project/backend/internal/handlers"
)

func main() {


    mux := http.NewServeMux()
    mux.HandleFunc("/", handlers.TestHandlers)

    err := http.ListenAndServe(":3333", mux)
    fmt.Println(err)
}

