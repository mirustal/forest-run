package runs

import (
	"time"
)

type (
	PermissionsType int
)

const (
	FreePermissionsType = PermissionsType(0)
)

type (
	Id            int
	EssentialInfo struct {
		Id             Id              `json:"id,omitempty"`
		Route          Route           `json:"route"`
		StartTime      time.Time       `json:"start_time"`
		StartPlace     Place           `json:"start_place"`
		RunPermissions PermissionsType `json:"permissions,omitempty"`
	}
	Place struct {
		Address string `json:"address"`
		Point   Point  `json:"point"`
	}
	Route struct {
		Points []Point `json:"points,omitempty"`
	}
	Point struct {
		Latitude  float64 `json:"latitude" yaml:"latitude"`
		Longitude float64 `json:"longitude" yaml:"longitude"`
	}
)
