package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Group struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []*User   `gorm:"many2many:user_groups;" json:"users"`
	Messages  []Message `json:"messages"` // Every group can have a lot of messages
	Admins    Admins    `gorm:"type:text" json:"admins"`
	IsDm      bool      `json:"isDm" gorm:"not null"`
}

type GroupRequest struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"user_ids"`
	Admins  Admins   `json:"admins"`
	IsDm    bool     `json:"isDm"`
}

type Admins []string

func (a *Admins) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		if !strings.HasPrefix(v, "[") {
			newString := fmt.Sprintf("[\"%s\"]", v)
			return json.Unmarshal([]byte(newString), a)
		}
		return json.Unmarshal([]byte(v), a)
	case []byte:
		return json.Unmarshal(v, a)
	default:
		return errors.New("unsupported Scan, storing driver.Value type " + string(value.([]byte)) + " into type *Admins")
	}
}

func (a Admins) Value() (driver.Value, error) {
	return json.Marshal(a)
}
