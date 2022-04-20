package gamesocketrouter

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/controllers/gamesocketcontroller"
	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	router := engine.Group("/ws/game")
	router.GET("/", gamesocketcontroller.Controller)
}
