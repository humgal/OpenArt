package bid

import (
	"github.com/shopspring/decimal"

	_ "github.com/go-sql-driver/mysql"
)

type Bid struct {
	Id            int             `json:"id"`
	ItemId        int             `json:"itemId"`
	ServiceFee    decimal.Decimal `json:"serviceFee"`
	Total         decimal.Decimal `json:"total"`
	UserID        string          `json:"userId"`
	Username      string          `json:"username"`
	BidDate       string          `json:"bidDate"`
	ItemPriceType int             `json:"itemPriceType"`
	OnsaleType    int             `json:"onsaleType"`
	BidEndDate    *string         `json:"bidEndDate"`
}
