package router

import (
	"backend_skripsi/controller"
	"net/http"

	framework "github.com/bandros/framework"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Static("/asset", "./asset")
	r.Static("/public", "./public")
	r.LoadHTMLGlob("./pages/**/*")

	r.NoRoute(error404)
	r.NoMethod(error404)
	r.POST("/login/proses", controller.LoginProses)
	r.POST("/user/register", controller.RegisterUser)
	r.GET("/list/campaign", controller.GetAllCampaign)

	r.GET("/get/provinsi", controller.ShowProvinsi)
	r.GET("/get/kota", controller.ShowKota)
	r.GET("/get/kecamatan", controller.ShowKecamatan)
	r.GET("/get/kelurahan", controller.ShowKelurahan)
	r.GET("/percobaan", controller.Percobaan)
	RouterHome(r)

}

func error404(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	login := false
	if v != nil {
		login = true
	}

	js := []string{
		"/asset/js/popper.min",
		"/asset/js/errorpagenotfound",
	}

	c.HTML(http.StatusNotFound, "error/404", gin.H{
		"title": "Error 404",
		"login": login,
		"js":    js,
	})
}
