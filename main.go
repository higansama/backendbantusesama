package main

import (
	"backend_skripsi/router"
	"net/http"
	"os"

	"github.com/bandros/framework"
)

func main() {
	framework := framework.Init{}
	framework.Get()

	r := framework.Begin
	router.Router(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
	// framework.Run()
}
