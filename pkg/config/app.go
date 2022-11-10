package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// this method connect to the mysql db and
// assign created db instance to the db object
func Connect() {
	d, err := gorm.Open("mysql", "root:mysql@/bookstore_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// to get db object
func GetDB() *gorm.DB {
	return db
}

// ALTER USER 'root'@'localhost' IDENTIFIED BY 'mysql';
