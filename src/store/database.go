package store

import (
	"app/src/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


var Db *gorm.DB

// OpenDb initializes database connection
func OpenDb(conString  string) {

	fmt.Println(conString)
	var err error
	Db, err = gorm.Open(utils.DatabaseName, conString)
	if err != nil {
		fmt.Print("failed to connect database")
		panic("failed to connect database")
	}
}

// CloseDb closes database connection
func CloseDb() {
	err := Db.Close()
	if err != nil {
		panic("failed to close database")
	}
}