package router

import (
	"apiserver/handler/ovn"
	"net/http"

	_ "apiserver/docs" // docs is generated by Swag CLI, you have to import it.
	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// The user handlers, requiring authentication
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	// API For OVN Logical switch
	ovnLs := g.Group("/api/v1/esix/ovn/LS")
	ovnLs.Use(middleware.AuthMiddleware())
	{
		ovnLs.PUT("/:name", ovn.LSAdd)
		ovnLs.DELETE("/:name", ovn.LSDel)
		ovnLs.GET("/:name", ovn.LSGet)
		ovnLs.GET("", ovn.LSList)
	}


	return g
}
