package models

import (
	"errors"

	"github.com/Kamva/mgm/v2"
)

// Player struct
type Player struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
}

// CreatePlayer is a wrapper that creates a new player entry
func CreatePlayer(name, email string) *Player {
	return &Player{
		Name:  name,
		Email: email,
	}
}

// GetPlayerByID gets it
func GetPlayerByID(id string) (*Player, error) {
	player := &Player{}
	collection := mgm.Coll(player)

	err := collection.FindByID(id, player)
	if err != nil {
		return player, errors.New("Player not found")
	}

	return player, nil
}

var (
	players []*Player
)

// GetPlayers here
func GetPlayers() []*Player {
	return players
}
