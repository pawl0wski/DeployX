package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
)

func Initialize(databasePath string) {
	if databasePath == "" {
		databasePath = ":memory:"
	}
	var err error
	DBConn, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}

}

func InitializeModel(model interface{}) {
	err := DBConn.AutoMigrate(model)
	if err != nil {
		panic(fmt.Sprintf("Can't migrate %s", reflect.TypeOf(model)))
	}
}
