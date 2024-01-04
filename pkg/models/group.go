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
	Name      string    `json:"name"`
	Usernames Usernames `json:"usernames" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []*User   `gorm:"many2many:user_groups;" json:"users"`
	Messages  []Message `json:"messages"` // Every group can have a lot of messages
	// Need to create a new type Admins because GORM doesn't support []string
	Admins Admins `gorm:"type:text" json:"admins"`
	IsDm   bool   `json:"is_dm" gorm:"not null"`
}

type GroupRequest struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"user_ids"`
	Admins  Admins   `json:"admins"`
	IsDm    bool     `json:"is_dm"`
}

type GetExistingGroupsByUsersAndAdminsRequest struct {
	UserIds []string `json:"user_ids"`
	Admins  Admins   `json:"admins"`
}

type Admins []string
type Usernames map[string]string

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

func (m *Usernames) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		if !strings.HasPrefix(v, "{") {
			newString := fmt.Sprintf("{\"%s\"}", v)
			return json.Unmarshal([]byte(newString), m)
		}
		return json.Unmarshal([]byte(v), m)
	case []byte:
		return json.Unmarshal(v, m)
	default:
		return errors.New("unsupported Scan, storing driver.Value type " + string(value.([]byte)) + " into type *Usernames")
	}
}

func (m Usernames) Value() (driver.Value, error) {
	// Convert the map to a string value
	return json.Marshal(m)
}
