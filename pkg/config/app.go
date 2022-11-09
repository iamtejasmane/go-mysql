package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// this method connect to the mysql db and
// assign created db instance to the db object
func Connect() {
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// to get db object
func GetDB() *gorm.DB {
	return db
}
