package main

import (
	"backend_skripsi/router"
	"fmt"
	"net/http"
	"os"

	"github.com/bandros/framework"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Port => ", port)

	os.Setenv("PORT", port)

	framework := framework.Init{}
	framework.Get()

	r := framework.Begin
	router.Router(r)
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
	// framework.Run()
}
