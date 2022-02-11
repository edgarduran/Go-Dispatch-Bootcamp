package main

import (
	"fmt"
	"net/http"

	"github.com/edgarduran/Go-Dispatch-Bootcamp/router"
)

func main() {
	fmt.Printf("Hello\n")
	router := router.Setup()

	fmt.Printf("Listening on port 3030\n")
	http.ListenAndServe(":3030", router)
}
