package models

import (
	"errors"
	"time"

	"github.com/konrad-marzec/fiber-api/database"
)

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Orders    []Order `gorm:"foreignKey:UserRefer"`
}

func (u *User) FindById(id int) error {
	database.Database.Db.Find(u, "id = ?", id)

	if u.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func (u *User) FindOrders(o *[]Order) {
	database.Database.Db.Find(o, "user_id = ?", u.ID)
}
