package models

import (
	"github.com/Kamva/mgm/v2"
)

// Request struct
type Request struct {
	mgm.DefaultModel `bson:",inline"`
	Interest         string `json:"interest" bson:"interest"`
	PlayerID         string `json:"playerID" bson:"playerID"`

	GeoJSONType string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

// CreateRequest is a wrapper that creates a new request entry
func CreateRequest(playerID string, interest string, lat, lon float64) *Request {
	return &Request{
		Interest:    interest,
		PlayerID:    playerID,
		GeoJSONType: "Point",
		Coordinates: []float64{lat, lon},
	}
}
