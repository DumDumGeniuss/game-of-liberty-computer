package gamesocketcontroller

import (
	"encoding/json"
	"math/rand"
)

func getActionTypeFromMessage(msg []byte) (*actionType, error) {
	var newAction action
	err := json.Unmarshal(msg, &newAction)
	if err != nil {
		return nil, err
	}

	return &newAction.Type, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomHash(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
