package main

import (
	"backend_skripsi/router"
	"os"

	"github.com/bandros/framework"
)

func main() {
	framework := framework.Init{}
	framework.Get()

	r := framework.Begin
	router.Router(r)
	// framework.Run()
	r.Run(":" + os.Getenv("portHost"))
}
