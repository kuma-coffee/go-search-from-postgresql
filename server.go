package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kuma-coffee/go-search-from-postgresql/controller"
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	router "github.com/kuma-coffee/go-search-from-postgresql/http"
	"github.com/kuma-coffee/go-search-from-postgresql/repository"
	"github.com/kuma-coffee/go-search-from-postgresql/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {

	return gin
}

func main() {
	// connect to postgresql
	const (
		host     = "localhost"
		username = "postgres"
		password = "postgres"
		dbName   = "test"
		port     = 5432
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbName)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Post{})

	httpRouter := router.NewChiRouter()
	postRepository := repository.NewPostgresRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)

	// httpRouter.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Hello World"})
	// })

	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(":8080")
}
