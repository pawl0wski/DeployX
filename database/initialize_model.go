package database

import (
	"fmt"
	"reflect"
)

func InitializeModel(model interface{}) {
	err := DBConn.AutoMigrate(model)
	if err != nil {
		panic(fmt.Sprintf("Can't migrate %s", reflect.TypeOf(model)))
	}
}
