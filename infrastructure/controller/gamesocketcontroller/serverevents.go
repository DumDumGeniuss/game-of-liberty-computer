package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/areadto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/unitdto"
)

type eventType string

const (
	errorHappenedEventType      eventType = "ERROR"
	informationUpdatedEventType eventType = "INFORMATION_UPDATED"
	areaUpdatedEventType        eventType = "AREA_UPDATED"
	unitsUpdatedEventType       eventType = "UNITS_UPDATED"
)

type errorHappenedEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type errorHappenedEvent struct {
	Type    eventType                 `json:"type"`
	Payload errorHappenedEventPayload `json:"payload"`
}

func constructErrorHappenedEvent(clientMessage string) *errorHappenedEvent {
	return &errorHappenedEvent{
		Type: errorHappenedEventType,
		Payload: errorHappenedEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

type informationUpdatedEventPayload struct {
	MapSize      mapsizedto.MapSizeDTO `json:"mapSize"`
	PlayersCount int                   `json:"playersCount"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

func constructInformationUpdatedEvent(mapSizeDTO mapsizedto.MapSizeDTO, playersCount int) *informationUpdatedEvent {
	return &informationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: informationUpdatedEventPayload{
			MapSize:      mapSizeDTO,
			PlayersCount: playersCount,
		},
	}
}

type unitsUpdatedEventPayloadItem struct {
	Coordinate coordinatedto.CoordinateDTO `json:"coordinate"`
	Unit       unitdto.UnitDTO             `json:"unit"`
}

type unitsUpdatedEventPayload struct {
	Items []unitsUpdatedEventPayloadItem `json:"items"`
}
type unitsUpdatedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsUpdatedEventPayload `json:"payload"`
}

func constructUnitsUpdatedEvent(coordinateDTOs []coordinatedto.CoordinateDTO, unitDTOs []unitdto.UnitDTO) *unitsUpdatedEvent {
	unitsUpdatedEventPayloadItems := []unitsUpdatedEventPayloadItem{}
	for i := 0; i < len(unitDTOs); i += 1 {
		unitsUpdatedEventPayloadItems = append(
			unitsUpdatedEventPayloadItems,
			unitsUpdatedEventPayloadItem{
				Coordinate: coordinateDTOs[i],
				Unit:       unitDTOs[i],
			},
		)
	}
	return &unitsUpdatedEvent{
		Type: unitsUpdatedEventType,
		Payload: unitsUpdatedEventPayload{
			Items: unitsUpdatedEventPayloadItems,
		},
	}
}

type areaUpdatedEventPayload struct {
	Area  areadto.AreaDTO     `json:"area"`
	Units [][]unitdto.UnitDTO `json:"units"`
}
type areaUpdatedEvent struct {
	Type    eventType               `json:"type"`
	Payload areaUpdatedEventPayload `json:"payload"`
}

func constructAreaUpdatedEvent(gameAreaDTO areadto.AreaDTO, units [][]unitdto.UnitDTO) *areaUpdatedEvent {
	return &areaUpdatedEvent{
		Type: areaUpdatedEventType,
		Payload: areaUpdatedEventPayload{
			Area:  gameAreaDTO,
			Units: units,
		},
	}
}
