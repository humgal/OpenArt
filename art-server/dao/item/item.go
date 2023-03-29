package item

import "github.com/shopspring/decimal"

type Item struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Tag          *string    `json:"tag"`
	Description  *string    `json:"description"`
	UploadUrl    string     `json:"uploadUrl"`
	SalesStatus  int        `json:"saleStatus"`
	Price        *ItemPrice `json:"price"`
	CreatorId    string     `json:"createorId"`
	Creator      string     `json:"creator"`
	CreateDate   *string    `json:"createDate"`
	CollectionId *int       `json:"collectionId"`
	OwnerId      *int       `json:"ownerId"`
	Owner        *string    `json:"owner"`
	Status       *int       `json:"status"`
}

type ItemPrice struct {
	ID             int             `json:"id"`
	ItemID         int             `json:"itemId"`
	ItemPriceType  int             `json:"itemPriceType"`
	OnsaleType     int             `json:"onsaleType"`
	InitPrice      decimal.Decimal `json:"initPrice"`
	ServiceFee     decimal.Decimal `json:"serviceFee"`
	StartDate      *string         `json:"startDate"`
	ExpirationDate *string         `json:"expirationDate"`
}
