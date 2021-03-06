package main

import (
	"fmt"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/apis"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/config"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/middleware"
	brotli "github.com/anargu/gin-brotli"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())
	r.Use(brotli.Brotli(brotli.DefaultCompression))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/posts/:id", apis.GetPost)
		v1.GET("/posts/", apis.GetPosts)
		v1.GET("/tags/:post_id", apis.GetTags)
		v1.GET("/sections/:post_id", apis.GetSections)
		v1.GET("/books/", apis.GetBooks)
		v1.GET("/projects/", apis.GetProjects)
		v1.POST("/newsletter/subscribe/", apis.AddSubscriber)
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	// config.Config.DB.AutoMigrate(&models.Post{}, &models.Project{}, &models.Section{}, &models.Tag{}, &models.Book{}) // This is needed for generation of schema for postgres image.

	defer config.Config.DB.Close()

	fmt.Println(fmt.Sprintf("Successfully connected to :%v", config.Config.DSN))

	r.RunTLS(fmt.Sprintf(":%v", config.Config.ServerPort), config.Config.CertFile, config.Config.KeyFile)
}
