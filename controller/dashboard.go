package controller

import (
	framework "github.com/aripstheswike/swikefw"
	"github.com/gin-gonic/gin"

	"backend_skripsi/lib"
	"backend_skripsi/model"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	nohp := c.PostForm("nohp")
	kota := c.PostForm("kota")
	kecamatan := c.PostForm("kecamatan")
	kelurahan := c.PostForm("kelurahan")
	password := c.PostForm("password")

	dataToInsert := map[string]interface{}{
		"username":  username,
		"email":     email,
		"nohp":      nohp,
		"kota":      kota,
		"kecamatan": kecamatan,
		"kelurahan": kelurahan,
		"password":  framework.Password(password),
	}

	err := model.InsertDataToTable("user", dataToInsert)
	if err != nil {
		lib.ReturnError(c, err)
		return
	}

	lib.Json(c, 200, "success", gin.H{})
}
