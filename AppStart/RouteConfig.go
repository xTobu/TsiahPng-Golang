package appstart

import (
	controllers "TsiahPng-Golang/Controllers"
	middlewares "TsiahPng-Golang/Middlewares"

	"github.com/gin-gonic/gin"

	//mysql
	_ "github.com/go-sql-driver/mysql"
)

// ========== server

//Config struct
type Config struct {
	Port         string
	StaticFolder string
	IndexFile    string
}

//SetDefault Sever data
func (config *Config) SetDefault() {
	config.Port = ":8080"
	config.StaticFolder = "../dist"
	config.IndexFile = "../index.html"
}

////////////////////

// RouteConfig blablaba
func RouteConfig() {
	// set config
	config := Config{}
	config.SetDefault()

	// Creates a default gin router
	router := gin.Default() // Grouping routes
	router.Use(middlewares.CORSMiddleware())

	//group： api
	api := router.Group("/api")
	// api.Use(middlewares.CORSMiddleware())
	// api.Use(middlewares.ValidateTokenMiddleware())
	{
		api.GET("/get", controllers.APIStudent)
		//api.OPTIONS("/get", controllers.APIStudent)

		api.POST("/post", controllers.APIInsert)

		api.GET("/checktoken", controllers.AuthCheckToken)
		//api.OPTIONS("/checktoken", controllers.AuthCheckToken)
		//group： api
		tsiahpng := api.Group("/tsiahpng")
		{
			tsiahpng.GET("/RestaurantsList", controllers.TsiahPngGetList)

		}
	}

	//group： auth
	auth := router.Group("/auth")
	// auth.Use(middlewares.CORSMiddleware())
	{
		auth.POST("/login", controllers.AuthLogin)
		//auth.OPTIONS("/login", controllers.AuthLogin)

	}

	router.Run(config.Port) // listen and serve on 0.0.0.0:8000

}
