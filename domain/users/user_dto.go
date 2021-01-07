package users

import (
	"strings"
	"time"

	"github.com/ab3llo/bookstore_users-api/utils/errors"
	"gorm.io/gorm"
)

const (
	StatusActive = "active"
)

// User dto
type User struct {
	ID        int64          `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Status    string         `json:"status"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     string         `json:"email" sql:"not null;unique"`
	Password  string         `gorm:"type:varchar(255); not null" json:"password,omitempty"`
}

// Validate user struct
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
