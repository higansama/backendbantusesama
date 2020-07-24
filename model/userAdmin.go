package model

import (
	"errors"

	framework "github.com/aripstheswike/swikefw"
)

func Login(email string, password string) (map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()
	db.Select("*")
	db.From("user")
	db.Where("email", email)
	user, err := db.Row()
	if err != nil {
		return nil, err
	}
	if framework.ValidPassword(password, user["password"].(string)) {
		return user, nil
	}
	return nil, errors.New("Invalid password")
}
