package config

import (
	"github.com/jinzhu/gorm"
	// "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConnection(url string)  {
	db, err := gorm.Open("mysql",url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
