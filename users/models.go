package users

import (
	"apigo/database"
	"fmt"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model

	Nome  string `json:"nome"`
	Senha string `json:"senha"`

	Seguidores int `json:"seguidores"`
}

func Migrate() {
	database.RootDatabase.DB.AutoMigrate(&UserModel{})
	fmt.Println("Migrando Usuarios...")
}

func (u *UserModel) Save() error {
	result := database.RootDatabase.DB.Create(u)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
