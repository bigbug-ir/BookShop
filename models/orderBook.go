package model

import "gorm.io/gorm"

/*****************************************************************/

type OrderBook struct {
	gorm.Model
	OrderID  uint `gorm:"primaryKey"`
	BookID   uint `gorm:"primaryKey"`
	Quantity int  `gorm:"not null"`
	Book     Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

/*****************************************************************/
