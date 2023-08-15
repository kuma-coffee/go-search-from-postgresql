package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kuma-coffee/go-search-from-postgresql/controller"
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	"github.com/kuma-coffee/go-search-from-postgresql/repository"
	"github.com/kuma-coffee/go-search-from-postgresql/scraper"
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

	postRepository := repository.NewPostgresRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	webScraper := scraper.NewScraper().Scraper()

	server := gin.Default()

	for _, item := range webScraper {
		if item.Ndex == "#0000" {
			continue
		}
		err := postRepository.Store(&entity.Post{
			Ndex:       item.Ndex,
			Pokemon:    item.Pokemon,
			PokemonURL: item.PokemonURL,
			Type:       item.Type,
		})
		if err != nil {
			panic(err)
		}
	}

	server.POST("/posts", postController.AddPost)
	server.GET("/posts", postController.FindAll)

	server.Run(":8080")
}
