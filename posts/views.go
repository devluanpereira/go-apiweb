package posts

import (
	"apigo/users"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostHandler struct{}

func MakePostHandler() PostHandler {
	return PostHandler{}
}

// CreatePostHandler lida com a criação de um post
func (p *PostHandler) CreatePostHandler(c *gin.Context) {
	var postModel PostModel // Usar uma variável local para a vinculação

	// Vincular JSON a postModel
	err := c.BindJSON(&postModel)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":   "Falha",
				"mensagem": err.Error(),
			},
		)
		return
	}

	userId := postModel.UserId
	var userModel users.UserModel

	// Verificar se o usuário existe
	_, found := userModel.FindByID(userId)
	if !found {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":   "Falha",
				"mensagem": "ID do usuário não foi encontrado",
			},
		)
		return
	}

	// Salvar o post
	err = postModel.Save()
	if err != nil {
		// Lidar com erro de violação de restrição única
		if isUniqueConstraintError(err) {
			c.JSON(
				http.StatusConflict, gin.H{
					"status":   "Falha",
					"mensagem": "Post com este ID já existe",
				},
			)
		} else {
			c.JSON(
				http.StatusInternalServerError, gin.H{
					"status":   "Falha",
					"mensagem": err.Error(),
				},
			)
		}
		return
	}

	c.JSON(
		http.StatusCreated, gin.H{
			"status": "Sucesso",
		},
	)
}

// isUniqueConstraintError verifica se um erro é uma violação de restrição única
func isUniqueConstraintError(err error) bool {
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}
