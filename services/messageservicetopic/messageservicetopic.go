package messageservicetopic

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto"
)

const GameWorkerTickedMessageTopic = "GAME_WORKER_TICKED"

const GameUnitsUpdatedMessageTopic = "GAME_UNITS_UPDATED"

type GameUnitsUpdatedMessageTopicPayloadUnit struct {
	Coordinate dto.CoordinateDTO
	Unit       dto.GameUnitDTO
}
type GameUnitsUpdatedMessageTopicPayload []GameUnitsUpdatedMessageTopicPayloadUnit

const GamePlayerJoinedMessageTopic = "GAME_PLAYER_JOINED"

const GamePlayerLeftMessageTopic = "GAME_PLAYER_LEFT"
