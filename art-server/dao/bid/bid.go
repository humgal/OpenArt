package bid

type Bid struct {
	ID            int     `json:"id"`
	ItemID        int     `json:"itemId"`
	ServiceFee    float64 `json:"serviceFee"`
	Total         float64 `json:"total"`
	UserID        string  `json:"userId"`
	Username      string  `json:"username"`
	BidDate       string  `json:"bidDate"`
	ItemPriceType int     `json:"itemPriceType"`
	OnsaleType    int     `json:"onsaleType"`
	BidEndDate    *string `json:"bidEndDate"`
}
