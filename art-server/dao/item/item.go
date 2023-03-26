package item

type Item struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Tag          *string    `json:"tag"`
	Description  *string    `json:"description"`
	UploadURL    string     `json:"uploadUrl"`
	SaleStatus   int        `json:"saleStatus"`
	Price        *ItemPrice `json:"price"`
	CreateorID   string     `json:"createorId"`
	Creator      string     `json:"creator"`
	CreateDate   *string    `json:"createDate"`
	CollectionID *int       `json:"collectionId"`
	OwnerID      *int       `json:"ownerID"`
	Owner        *string    `json:"owner"`
	Status       *int       `json:"status"`
}

type ItemPrice struct {
	ID             int     `json:"id"`
	ItemID         int     `json:"itemId"`
	ItemPriceType  int     `json:"itemPriceType"`
	OnsaleType     int     `json:"onsaleType"`
	InitPrice      float64 `json:"initPrice"`
	ServiceFee     float64 `json:"serviceFee"`
	StartDate      *string `json:"startDate"`
	ExpirationDate *string `json:"expirationDate"`
}
