package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	model UserModel
}

func MakeUserHandle() UserHandler {
	return UserHandler{model: UserModel{}}
}

func (u *UserHandler) SignUpHandler(c *gin.Context) {
	fmt.Print(u.model)
	err := c.BindJSON(&u.model)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "Failed",
				"message": err.Error(),
			},
		)

		return
	}

	if u.model.NameExists() {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":  "Failed",
				"message": "Erro, este usuario ja existe...",
			},
		)

		return
	}

	err = u.model.Save()
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
		http.StatusBadRequest, gin.H{
			"status":  "Sucesso",
			"message": "Usuarios criado...",
		},
	)

}

func (u *UserHandler) GetUsersHandler(c *gin.Context) {
	c.JSON(
		http.StatusBadRequest, gin.H{
			"users": u.model.GetAll(),
		},
	)

}
