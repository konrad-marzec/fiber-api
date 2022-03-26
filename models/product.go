package models

import (
	"errors"
	"time"

	"github.com/konrad-marzec/fiber-api/database"
)

type Product struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func (u *Product) FindById(id int) error {
	database.Database.Db.Find(u, "id = ?", id)

	if u.ID == 0 {
		return errors.New("product does not exist")
	}

	return nil
}
