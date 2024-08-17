package model

/*****************************************************************/

type OrderBook struct {
	ID       uint  `gorm:"primaryKey" json:"orderbook_id"`
	OrderID  uint  `gorm:"primaryKey" json:"order_id"`
	BookID   uint  `gorm:"primaryKey" json:"book_id"`
	Quantity int   `gorm:"not null;default:1" json:"quantity"`
	Order    Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Book     Book  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

/*****************************************************************/
