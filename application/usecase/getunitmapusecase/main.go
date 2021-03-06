package getunitmapusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/areadto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/unitdto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/google/uuid"
)

type useCase struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
}

func New(gameRoomRepository gameroomrepository.GameRoomRepository) *useCase {
	return &useCase{
		gameRoomRepository: gameRoomRepository,
	}
}

func (uc *useCase) Execute(gameId uuid.UUID, areaDTO areadto.AreaDTO) ([][]unitdto.UnitDTO, error) {
	gameRoom, err := uc.gameRoomRepository.Get(gameId)
	if err != nil {
		return make([][]unitdto.UnitDTO, 0), err
	}

	area := areadto.FromDTO(areaDTO)
	unitMap, err := gameRoom.GetUnitMapWithArea(area)
	if err != nil {
		return make([][]unitdto.UnitDTO, 0), err
	}

	return unitdto.ToDTOMap(unitMap), nil
}
