package users

import (
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
			http.StatusForbidden, gin.H{
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
		http.StatusOK, gin.H{
			"status":  "Sucesso",
			"message": "Usuarios criado...",
		},
	)
}

func (u *UserHandler) GetUsersHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"users": u.model.GetAll(),
		},
	)

}

func (u *UserHandler) LoginHandler(c *gin.Context) {
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

	if u.model.Exists() {
		c.JSON(
			http.StatusOK, gin.H{
				"status": "Sucesso",
			},
		)

		return
	}

	c.JSON(
		http.StatusBadRequest, gin.H{
			"status": "Faied",
		},
	)

}
