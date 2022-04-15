package gamecontroller

import (
	"strconv"
	"strings"

	"github.com/DumDumGeniuss/game-of-liberty-computer/entities/gameentity"
	"github.com/gin-gonic/gin"
)

func GetController(c *gin.Context) {
	fromXY := strings.Split(c.Query("from"), ",")
	toXY := strings.Split(c.Query("to"), ",")
	fromX, err := strconv.Atoi(fromXY[0])
	if err != nil {
		c.JSON(400, err)
		return
	}
	fromY, err := strconv.Atoi(fromXY[1])
	if err != nil {
		c.JSON(400, err)
		return
	}
	toX, err := strconv.Atoi(toXY[0])
	if err != nil {
		c.JSON(400, err)
		return
	}
	toY, err := strconv.Atoi(toXY[1])
	if err != nil {
		c.JSON(400, err)
		return
	}
	units, err := gameentity.GetUnits(fromX, fromY, toX, toY)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(
		200,
		units,
	)
}
