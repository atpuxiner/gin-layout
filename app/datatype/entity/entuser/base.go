package entuser

import (
	"github.com/atpuxiner/gin-layout/app/datatype/model"
	"time"
)

type User struct {
	model.BaseUser
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Age       uint      `json:"age"`
	Gender    uint      `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
