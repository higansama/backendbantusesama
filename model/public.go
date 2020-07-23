package model

import "github.com/bandros/framework"

func InsertDataToTable(table string, data map[string]interface{}) error {
	db := framework.Database{}
	defer db.Close()

	db.From(table)
	_,e := db.Insert(data)
	return e
}