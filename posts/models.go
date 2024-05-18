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

	UserID int `json:"user"`
}

func MakePostModel() PostModel {
	return PostModel{}
}

func (p *PostModel) Migrate() {
	database.RootDatabase.DB.AutoMigrate(&PostModel{})
	fmt.Println("Migrando Posts...")
}
