package posts

import "github.com/gin-gonic/gin"

type PostHandler struct {
	model PostModel
}

func MakePostHandler() PostHandler {
	return PostHandler{}
}

func (p *PostHandler) CreatePostHandler(c *gin.Context) {

}
