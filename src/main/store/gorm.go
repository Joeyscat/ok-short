package store

import (
	. "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type ORM struct {
}

var (
	MyDB = &DB{}
)

func init() {
	db, err := Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ok-short?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	log.Printf("MySQL Connected : %+v", db.DB().Stats())
	MyDB = db
}
