package controllers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lwinmgmg/chat/controllers/v1"
	"github.com/lwinmgmg/chat/middlewares"
)

func DefineRoutes(app *gin.Engine) {

	v1Router := app.Group("/api/v1", middlewares.JwtAuthMiddleware("Bearer"))
	v1Controller := v1.ControllerV1{
		Router: v1Router,
	}
	v1Controller.Serve()
}
