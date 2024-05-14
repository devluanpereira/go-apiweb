package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Path string
	DB* gorm.DB
}

var RootDatabase Database

func InitDatabase(path string) {
	var err error
	RootDatabase = Database{}
	RootDatabase.Path = path

	RootDatabase.DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		errMessage := fmt.Sprintf("Erro aoo abrir Banco De Dados => %s\n", err.Error())
		panic(errMessage)
	}

	fmt.Println("Banco de Dados iniciado...", path)

}
