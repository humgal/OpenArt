package payment

type Payment struct {
	ID         int     `json:"id"`
	ItemID     int     `json:"itemId"`
	OnsaleType int     `json:"onsaleType"`
	Price      float64 `json:"price"`
	ServiceFee float64 `json:"serviceFee"`
	Createor   string  `json:"createor"`
	CreateDate *string `json:"createDate"`
	PayStatus  int     `json:"payStatus"`
	PayUser    *string `json:"payUser"`
	PayDate    *string `json:"payDate"`
}
