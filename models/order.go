package models

import (
	"errors"
	"time"

	"github.com/konrad-marzec/fiber-api/database"
)

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer uint    `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    uint    `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}

func (u *Order) FindById(id int) error {
	database.Database.Db.Joins("User").Joins("Product").Find(u, "orders.id = ?", id)

	if u.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}
