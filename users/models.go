package users

import (
	"fmt"
	"gorm.io/gorm"
	"apigo/database"
)

type UserModel struct {
	gorm.Model

	Nome string
	Senha string

	Seguidores int
}

func Migrate() {
	database.RootDatabase.DB.AutoMigrate(&UserModel{})
	fmt.Println("Migrando Usuarios...")
}