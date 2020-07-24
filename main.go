package main

import (
	"backend_skripsi/router"
	"fmt"
	"os"

	"github.com/bandros/framework"
)

func main() {
	framework := framework.Init{}
	framework.Get()

	r := framework.Begin
	router.Router(r)
	fmt.Println(os.Getenv("portHost"))

	// framework.Run()
	fmt.Println(os.Getenv("PORT"))
	r.Run(":" + os.Getenv("portHost"))
}
