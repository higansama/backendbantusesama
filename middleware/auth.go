package middleware

import (
	"fmt"

	framework "github.com/aripstheswike/swikefw"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthUser(c *gin.Context) {
	var tokenString string
	session := sessions.Default(c)
	if v := session.Get(framework.Config("jwt")); v == nil {
		notfound(session, c)
		c.Abort()
		return
	} else {
		tokenString = v.(string)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(framework.Config("jwtKey")), nil
	})
	if token == nil {
		notfound(session, c)
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
		data_jwt := map[string]interface{}{}
		data_jwt["id"] = claims["id"]
		data_jwt["username"] = claims["username"]
		data_jwt["nohp"] = claims["nohp"]
		data_jwt["kota"] = claims["kota"]
		data_jwt["kecamatan"] = claims["kecamatan"]
		data_jwt["kelurahan"] = claims["kelurahan"]
		c.Set("jwt", data_jwt)
	} else {
		notfound(session, c)
		c.Abort()
		return
	}
}

func notfound(session sessions.Session, c *gin.Context) {
	// session.Delete(framework.Config("jwtName"))
	// session.Save()
	// c.Redirect(http.StatusFound, "/")
}
