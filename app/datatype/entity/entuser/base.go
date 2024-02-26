package entuser

import (
	"database/sql"
	"github.com/atpuxiner/gin-layout/app/datatype/model"
	"time"
)

type User struct {
	model.BaseUser
	ID        uint         `json:"id"`
	Phone     string       `json:"phone"`
	Password  string       `json:"password"`
	Name      string       `json:"name"`
	Age       int          `json:"age"`
	Gender    int          `json:"gender"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
