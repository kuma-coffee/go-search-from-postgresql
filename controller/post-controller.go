package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	"github.com/kuma-coffee/go-search-from-postgresql/service"
)

type PostController interface {
	AddPost(c *gin.Context)
	FindAll(c *gin.Context)
	Search(c *gin.Context)
	Sort(c *gin.Context)
}

type controller struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &controller{postService}
}

func (cp *controller) AddPost(c *gin.Context) {
	var newPost entity.Post

	err := c.ShouldBindJSON(&newPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = cp.postService.Store(&newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Store data success!"})
}

func (cp *controller) FindAll(c *gin.Context) {
	var query entity.Query
	err := c.Bind(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if query.SearchQuery != "" {
		querySplit := strings.Fields(query.SearchQuery)

		itemList, err := cp.postService.Search(querySplit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, itemList)
		return
	}

	if query.SortQuery != "" {
		itemList, err := cp.postService.Sort(query.SortQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, itemList)
		return
	}

	itemList, err := cp.postService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemList)
}

func (cp *controller) Search(c *gin.Context) {
	var query entity.Query
	err := c.Bind(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	querySplit := strings.Fields(query.SearchQuery)

	itemList, err := cp.postService.Search(querySplit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemList)
}

func (cp *controller) Sort(c *gin.Context) {
	var query entity.Query
	err := c.Bind(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemList, err := cp.postService.Sort(query.SortQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemList)
}
