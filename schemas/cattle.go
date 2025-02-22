package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Cattle struct {
	gorm.Model
	Name    string
	Earring string
	Gender  string
}

type CattleResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Name      string    `json:"name"`
	Earring   string    `json:"earring"`
	Gender    string    `json:"gender"`
}
