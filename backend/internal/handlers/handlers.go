package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func TestHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
