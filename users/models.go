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

func (u UserModel) GetAll() []UserModel {
	userModels := []UserModel{}
	database.RootDatabase.DB.Find(&userModels)

	return userModels
}

func (u UserModel) NameExists() bool {
	var tempUser UserModel
	result := database.RootDatabase.DB.Where("nome=?", u.Nome).First(&tempUser)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}

		// @TODO:
		panic(result.Error.Error())
	}

	return true
}
