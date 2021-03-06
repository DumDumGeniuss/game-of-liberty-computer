package router

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/router/gamesocketrouter"
	"github.com/gin-gonic/gin"
)

func SetRouters() {
	router := gin.Default()

	gamesocketrouter.SetRouter(router)
	router.Run()
}
