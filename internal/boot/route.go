package boot

import (
	"github.com/gin-gonic/gin"
	"go-egg/internal/server"
	"go-egg/internal/server/routes"
)

func SetupRoute(router *gin.Engine) {
	server.RegisterGlobalMiddleWare(router)

	routes.RegisterAPIRoutes(router)

	server.Setup404Handler(router)
}
