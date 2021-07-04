package router

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jtieri/Nostalgia/webserver/app"
	"github.com/jtieri/Nostalgia/webserver/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(app.WebApp.Config.WebServer.StaticFiles, true))) // Enable static client files
	api := router.Group("/api")
	v1 := api.Group("/v1")
	blogAPI := v1.Group("/blog")
	tvAPI := v1.Group("/tv")

	// Initialize routes for main site
	// nostalgiarewind.org/
	router.GET("/home", controllers.GetIndex)
	router.GET("/stream", controllers.GetStream)
	router.GET("/about", controllers.GetAbout)
	router.GET("/blog", controllers.GetBlog)

	// Initialize routes for blog API
	// nostalgiarewind.org/api/v1/blog/
	// nostalgiarewind.org/api/v1/blog/id/:id
	// nostalgiarewind.org/api/v1/blog/title/:title
	// nostalgiarewind.org/api/v1/blog/create/
	// nostalgiarewind.org/api/v1/blog/update/
	// nostalgiarewind.org/api/v1/blog/delete/
	blogAPI.GET("/", controllers.GetAllBlogPosts)
	blogAPI.GET("/id/:id", controllers.GetBlogPostById)
	blogAPI.GET("/title/:title", controllers.GetBlogPostByTitle)
	blogAPI.POST("/create/", controllers.CreateBlogPost)

	// Initialize routes for TV API
	// nostalgiarewind.org/api/v1/tv/
	// nostalgiarewind.org/api/v1/tv/network/:network
	// nostalgiarewind.org/api/v1/tv/year/:year
	// nostalgiarewind.org/api/v1/tv/date/:date
	// nostalgiarewind.org/api/v1/tv/show/:title
	tvAPI.GET("/", controllers.GetAllRecordings)
	tvAPI.GET("/year/:year", controllers.GetRecordingsByYear)
	tvAPI.POST("/create/", controllers.CreateRecording)
	tvAPI.GET("/network/:network", controllers.GetRecordingsByNetwork)

	return router
}
