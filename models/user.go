package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"gorm.io/gorm"
)

type User struct {
	PrimaryKey
	Name     string  `json:"name"`
	Username string  `json:"username" gorm:"not null"`
	Email    string  `json:"email" gorm:"not null" binding:"required"`
	Password string  `json:"password,omitempty" gorm:"not null" binding:"required"`
	Assets   []Asset `json:"assets"`
	UpdatedCreated
	DeletedAt
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var count = new(int64)
	db.DB.Model(&User{}).Where("email = ?", u.Email).Where("deleted_at is null").Count(count)

	fmt.Println(*count)

	if *count > 0 {
		return errors.New("duplicate email")
	}

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}
