package payment

import "github.com/shopspring/decimal"

type Payment struct {
	ID         int             `json:"id"`
	ItemID     int             `json:"itemId"`
	OnsaleType int             `json:"onsaleType"`
	Price      decimal.Decimal `json:"price"`
	ServiceFee decimal.Decimal `json:"serviceFee"`
	Createor   string          `json:"createor"`
	CreateDate *string         `json:"createDate"`
	PayStatus  int             `json:"payStatus"`
	PayUser    *string         `json:"payUser"`
	PayDate    *string         `json:"payDate"`
}
