package entuser

import (
	"github.com/atpuxiner/gin-layout/app/datatype/model"
	"time"
)

type GetResp struct {
	model.BaseUser
	ID        uint      `json:"id"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Gender    int       `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetListResp struct {
	model.BaseUser
	ID        uint      `json:"id"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Gender    int       `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
