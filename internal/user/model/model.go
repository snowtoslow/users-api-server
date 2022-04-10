package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string   `json:"id" gorm:"primary_key"`
	FirstName string   `json:"firstname" gorm:" column:firstname"`
	LastName  string   `json:"lastname" gorm:" column:lastname"`
	Age       uint8    `json:"age"`
	Position  Position `json:"position"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	if err = tx.Preload("Position").Error; err != nil {
		return
	}

	return
}

func UserFromReq(req UserReq) User {
	return User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Position:  MapToPosition(req.LatLong),
	}
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Age:         u.Age,
		LatLongResp: u.Position.ToLatLongResponse(),
	}
}

type Position struct {
	ID     string  `json:"id" gorm:"primaryKey"`
	Lat    float64 `json:"lat"`
	Long   float64 `json:"long"`
	UserID string  `json:"user_id"`
}

func (p *Position) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

func (p Position) ToLatLongResponse() LatLong {
	return LatLong{
		Latitude:  p.Lat,
		Longitude: p.Long,
	}
}

func MapToPosition(latLongMap LatLong) Position {
	return Position{
		Lat:  latLongMap.Latitude,
		Long: latLongMap.Longitude,
	}
}
