package controller

import (
	"backend_skripsi/lib"
	"backend_skripsi/model"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func ShowProvinsi(c *gin.Context) {
	data, err := model.GetAllProvinsi()
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}
func ShowKota(c *gin.Context) {
	id_Kota := c.Query("kota")
	data, err := model.GetKota(id_Kota)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}
func ShowKecamatan(c *gin.Context) {
	id_Kecamatan := c.Query("kecamatan")
	data, err := model.GetKecamatan(id_Kecamatan)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}
func ShowKelurahan(c *gin.Context) {
	id_Kelurahan := c.Query("kelurahan")
	data, err := model.GetKelurahan(id_Kelurahan)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}
	lib.Json(c, 200, "success", data)
}

func Percobaan(c *gin.Context) {
	// data := map[string]interface{}{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// 	"d": 4,
	// }
	// err := model.InsertDataToTable("test", data)
	// if err != nil {
	// 	lib.ReturnError(c, err)
	// 	return
	// }
	port := os.Getenv("portHost")
	fmt.Println(port)
	lib.Json(c, 200, "success", port)
}
