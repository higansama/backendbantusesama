package controller

import (
	"backend_skripsi/model"
	"net/http"

	"github.com/bandros/framework"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {

	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	login := false
	if v != nil {
		login = true
	}

	c.HTML(http.StatusOK, "default/index", gin.H{
		"title":  "Beranda",
		"method": "ok",
		"login":  login,
	})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	login := false
	if v != nil {
		login = true
	}

	js := []string{
		"/asset/js/popper.min",
		"/asset/js/login",
	}

	css := []string{
		"/asset/css/login",
		"/asset/css/iofrm-style-login",
		"/asset/css/iofrm-theme1-login",
	}

	c.HTML(http.StatusOK, "default/login", gin.H{
		"title": "Login",
		"js":    js,
		"css":   css,
		"login": login,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(framework.Config("jwtName"))
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func LoginProses(c *gin.Context) {
	email := c.DefaultPostForm("email", "")
	password := c.DefaultPostForm("password", "")
	user, err := model.Login(email, password)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  err.Error(),
		})
		return
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["id"] = user["id"]
	claims["username"] = user["username"]
	claims["nohp"] = user["nohp"]
	claims["kota"] = user["kota"]
	claims["kecamatan"] = user["kecamatan"]
	claims["kelurahan"] = user["kelurahan"]
	token, err := sign.SignedString([]byte(framework.Config("jwtKey")))
	if err != nil {
		framework.ErrorJson("JWT "+err.Error(), c)
		return
	}
	sesion := sessions.Default(c)
	sesion.Set(framework.Config("jwtName"), token)
	sesion.Save()
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "correct login",
		"token": token,
	})

}
