package item

import "github.com/shopspring/decimal"

type Item struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Tag          *string    `json:"tag"`
	Description  *string    `json:"description"`
	UploadUrl    string     `json:"uploadUrl"`
	SaleStatus   int        `json:"saleStatus"`
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
	Id             int             `json:"id"`
	ItemPriceType  int             `json:"itemPriceType"`
	OnsaleType     string          `json:"onsaleType"`
	InitPrice      decimal.Decimal `json:"initPrice"`
	ServiceFee     decimal.Decimal `json:"serviceFee"`
	StartDate      *string         `json:"startDate"`
	ExpirationDate *string         `json:"expirationDate"`
}
