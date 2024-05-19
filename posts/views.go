package posts

import (
	"apigo/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	model PostModel
}

func MakePostHandler() PostHandler {
	return PostHandler{}
}

func (p *PostHandler) CreatePostHandler(c *gin.Context) {
	err := c.BindJSON(&p.model)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "Failed",
				"message": err.Error(),
			},
		)

		return
	}

	userId := p.model.UserId
	var userModel users.UserModel

	_, found := userModel.FindByID(userId)

	if found == !found {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "Failed",
				"message": "ID do usuario nao foi encontrado",
			},
		)

		return
	}

	err = p.model.Save()
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "Failed",
				"message": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated, gin.H{
			"status": "Sucesso",
		},
	)

}
