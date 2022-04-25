package gameprovider

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

type GameProvider interface {
	GenerateNextUnits()
	GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error)
	GetGameSize() *GameSize
}

type gameProviderImplement struct {
	gameOfLiberty ggol.Game[GameUnit]
	gameDAO       gamedao.GameDAO
}

var gameStore GameProvider = nil

func CreateGameProvider(gameDAO gamedao.GameDAO) (GameProvider, error) {
	if gameStore != nil {
		return nil, &errGameProviderHasBeenCreated{}
	}
	initialUnit := GameUnit{
		Alive: false,
		Age:   0,
	}
	gameFieldSize, _ := gameDAO.GetGameSize()
	ggolSize := convertGameSizeToGgolSize(gameFieldSize)
	newGameOfLiberty, _ := ggol.NewGame(
		ggolSize,
		&initialUnit,
	)

	newGameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	gameField, _ := gameDAO.GetGameField()
	gameUnits := convertGameFieldToGameUnits(gameField)

	for x := 0; x < ggolSize.Width; x += 1 {
		for y := 0; y < ggolSize.Height; y += 1 {
			gameFieldUnit := &(*gameUnits)[x][y]
			coord := &ggol.Coordinate{X: x, Y: y}
			newGameOfLiberty.SetUnit(coord, gameFieldUnit)
		}
	}

	newGameProvider := &gameProviderImplement{
		gameOfLiberty: newGameOfLiberty,
		gameDAO:       gameDAO,
	}
	gameStore = newGameProvider
	return gameStore, nil
}

func (gsi *gameProviderImplement) GenerateNextUnits() {
	gsi.gameOfLiberty.GenerateNextUnits()
}

func (gsi *gameProviderImplement) GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error) {
	ggolArea := convertGameAreaToGgolArea(area)
	units, err := gsi.gameOfLiberty.GetUnitsInArea(ggolArea)
	if err != nil {
		return nil, err
	}
	return &units, nil
}

func (gsi *gameProviderImplement) GetGameSize() *GameSize {
	return convertGgolSizeToGameSize(gsi.gameOfLiberty.GetSize())
}
