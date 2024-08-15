package model

import (
	"bmacharia/jwt-go-rbac/database"

	"gorm.io/gorm"
)

/*****************************************************************/

type Order struct {
	gorm.Model
	ID         uint        `gorm:"primaryKey"`
	UserID     uint        `gorm:"not null"`
	User       User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Item       []OrderBook `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalPrice float64     `gorm:"not null"`
	Status     string      `gorm:"size:100;not null"`
}

/*****************************************************************/

func (order *Order) Save() (*Order, error) {
	err := database.Database.DB.Create(&order).Error
	if err != nil {
		return &Order{}, err
	}
	return order, nil
}

/*****************************************************************/

func GetOrders(order *[]Order) (err error) {
	err = database.Database.DB.Find(order).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetOrder(order *Order, id int) (err error) {
	err = database.Database.DB.Where("id=?", id).First(&order).Error
	if err != nil {
		return err
	}
	return
}

/*****************************************************************/

func GetOrdersCustomer(order *[]Order, id int) (err error) {
	err = database.Database.DB.Where("user_id=?", id).Find(&order).Error
	if err != nil {
		return err
	}
	return
}

/*****************************************************************/

func GetOrderCustomer(order *Order, id int, userId int) (err error) {
	err = database.Database.DB.Where("id=? AND user_id=?", id, userId).First(&order).Error
	if err != nil {
		return err
	}
	return
}

/*****************************************************************/

func UpdateOrder(order *Order) (err error) {
	err = database.Database.DB.Updates(order).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func DeleteOrder(order *Order) (err error) {
	err = database.Database.DB.Delete(order).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
