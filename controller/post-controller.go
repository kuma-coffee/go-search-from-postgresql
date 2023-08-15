package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	"github.com/kuma-coffee/go-search-from-postgresql/service"
)

type PostController interface {
	AddPost(c *gin.Context)
	FindAll(c *gin.Context)
}

type controller struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &controller{postService}
}

func (cp *controller) AddPost(c *gin.Context) {
	var newPost entity.Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := cp.postService.Store(&newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Store data success!"})
}

func (cp *controller) FindAll(c *gin.Context) {
	itemList, err := cp.postService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemList)
}
