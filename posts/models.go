package posts

import (
	"apigo/database"
	"fmt"

	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model

	Texto   string `json:"texto"`
	Like    int    `json:"like"`
	Dislike string `json:"dislike"`

	UserId int `json:"user"`
}

func MakePostModel() PostModel {
	return PostModel{}
}

func (p *PostModel) Migrate() {
	database.RootDatabase.DB.AutoMigrate(&PostModel{})
	fmt.Println("Migrando Posts...")
}

func (p *PostModel) Save() error {
	result := database.RootDatabase.DB.Create(p)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PostModel) GetAll() []PostModel {
	posts := []PostModel{}
	database.RootDatabase.DB.Find(&posts)

	return posts
}
