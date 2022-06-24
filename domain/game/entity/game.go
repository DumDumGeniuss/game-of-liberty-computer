package entity

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	Id         uuid.UUID
	UnitMatrix [][]valueobject.Unit
	MapSize    valueobject.MapSize
}

func NewGame() Game {
	id, _ := uuid.NewUUID()
	return Game{
		Id:         id,
		UnitMatrix: make([][]valueobject.Unit, 0),
		MapSize:    valueobject.NewMapSize(0, 0),
	}
}
