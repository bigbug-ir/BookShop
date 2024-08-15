package model

/*****************************************************************/

type OrderBook struct {
	OrderID  uint  `gorm:"primaryKey"`
	BookID   uint  `gorm:"primaryKey"`
	Quantity int   `gorm:"not null"`
	Order    Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Book     Book  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

/*****************************************************************/
