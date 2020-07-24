package main

import (
	"backend_skripsi/router"
	"fmt"
	"net/http"

	"github.com/bandros/framework"
)

func main() {
	framework := framework.Init{}
	framework.Get()

	r := framework.Begin
	router.Router(r)
	port := process.env.PORT
	if port == "" {
		port = "8080"
	}
	fmt.Println("Ini port yang dipakai => ", port)
	http.ListenAndServe(":"+port, nil)
	// framework.Run()
}
