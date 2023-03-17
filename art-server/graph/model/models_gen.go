// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Bid struct {
	TemID         int           `json:"temId"`
	ServiceFee    float64       `json:"serviceFee"`
	Total         float64       `json:"total"`
	User          string        `json:"user"`
	Username      string        `json:"username"`
	BidDate       string        `json:"bidDate"`
	ItemPriceType ItemPriceType `json:"itemPriceType"`
	BidEndDate    *string       `json:"bidEndDate"`
}

type BidParm struct {
	ItemID int     `json:"itemId"`
	Total  float64 `json:"total"`
}

type Collection struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Items      []*Item `json:"items"`
	CreateDate *string `json:"createDate"`
	Createor   *string `json:"createor"`
}

type CollectionParm struct {
	Name    *string `json:"name"`
	Creator string  `json:"creator"`
}

type Creator struct {
	UserID      string  `json:"userId"`
	Username    string  `json:"username"`
	Description *string `json:"description"`
	Followers   []*int  `json:"followers"`
	FollowerNum *int    `json:"followerNum"`
}

type Item struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Tag         *string    `json:"tag"`
	Description *string    `json:"description"`
	UploadURL   string     `json:"uploadUrl"`
	SaleStatus  int        `json:"saleStatus"`
	Price       *ItemPrice `json:"price"`
	Creator     string     `json:"creator"`
	CreateDate  *string    `json:"createDate"`
}

type ItemPrice struct {
	ItemID         string        `json:"itemId"`
	ItemPriceType  ItemPriceType `json:"itemPriceType"`
	Onsale         OnsaleType    `json:"onsale"`
	InitPrice      float64       `json:"initPrice"`
	ServiceFee     float64       `json:"serviceFee"`
	StartDate      *string       `json:"startDate"`
	ExpirationDate *string       `json:"expirationDate"`
}

type Link struct {
	Type LinkType `json:"type"`
	URL  string   `json:"url"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUser struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

type PayParam struct {
	ID         string     `json:"id"`
	OnsaleType OnsaleType `json:"onsaleType"`
	Balance    float64    `json:"balance"`
	ServiceFee float64    `json:"serviceFee"`
	PayAmount  float64    `json:"payAmount"`
}

type Payment struct {
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

type PriceParam struct {
	ItemID         string        `json:"itemId"`
	ItemPriceType  ItemPriceType `json:"itemPriceType"`
	Onsale         OnsaleType    `json:"onsale"`
	InitPrice      float64       `json:"initPrice"`
	StartDate      *string       `json:"startDate"`
	ExpirationDate *string       `json:"expirationDate"`
}

type PriceRange struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type SearchParm struct {
	Param   string      `json:"param"`
	Type    *SearchType `json:"type"`
	Price   *PriceRange `json:"price"`
	Chain   *Blockchain `json:"chain"`
	Onsale  *OnsaleType `json:"onsale"`
	Creator *string     `json:"creator"`
}

type SubscriptionEvent struct {
	Payments []*Payment `json:"payments"`
	Bids     []*Bid     `json:"bids"`
}

type User struct {
	ID         string      `json:"id"`
	Realname   string      `json:"realname"`
	Username   string      `json:"username"`
	Photo      *string     `json:"photo"`
	Phone      *string     `json:"phone"`
	Company    *string     `json:"company"`
	Email      *string     `json:"email"`
	Bio        *string     `json:"bio"`
	Img        *string     `json:"img"`
	JoinDate   *string     `json:"joinDate"`
	VerifyType *VerifyType `json:"verifyType"`
	VerifyName *string     `json:"verifyName"`
	IsCreator  *bool       `json:"isCreator"`
	Links      []*Link     `json:"links"`
}

type UploadItem struct {
	Name        string   `json:"name"`
	Tag         *string  `json:"tag"`
	Description *string  `json:"description"`
	UploadURL   []string `json:"uploadUrl"`
	Creator     string   `json:"creator"`
	SaleStatus  int      `json:"saleStatus"`
	Collection  *string  `json:"collection"`
}

type Wallet struct {
	Type     WalletType `json:"type"`
	PubToken string     `json:"pubToken"`
}

type Blockchain string

const (
	BlockchainEthereum Blockchain = "ETHEREUM"
	BlockchainMatic    Blockchain = "MATIC"
	BlockchainKlaytn   Blockchain = "KLAYTN"
	BlockchainSolana   Blockchain = "SOLANA"
	BlockchainBnb      Blockchain = "BNB"
)

var AllBlockchain = []Blockchain{
	BlockchainEthereum,
	BlockchainMatic,
	BlockchainKlaytn,
	BlockchainSolana,
	BlockchainBnb,
}

func (e Blockchain) IsValid() bool {
	switch e {
	case BlockchainEthereum, BlockchainMatic, BlockchainKlaytn, BlockchainSolana, BlockchainBnb:
		return true
	}
	return false
}

func (e Blockchain) String() string {
	return string(e)
}

func (e *Blockchain) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Blockchain(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Blockchain", str)
	}
	return nil
}

func (e Blockchain) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ItemPriceType string

const (
	ItemPriceTypeFixed   ItemPriceType = "FIXED"
	ItemPriceTypeAuction ItemPriceType = "AUCTION"
)

var AllItemPriceType = []ItemPriceType{
	ItemPriceTypeFixed,
	ItemPriceTypeAuction,
}

func (e ItemPriceType) IsValid() bool {
	switch e {
	case ItemPriceTypeFixed, ItemPriceTypeAuction:
		return true
	}
	return false
}

func (e ItemPriceType) String() string {
	return string(e)
}

func (e *ItemPriceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ItemPriceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ItemPriceType", str)
	}
	return nil
}

func (e ItemPriceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LinkType string

const (
	LinkTypeWebsite   LinkType = "WEBSITE"
	LinkTypeDiscord   LinkType = "DISCORD"
	LinkTypeInstagram LinkType = "INSTAGRAM"
	LinkTypeYoutube   LinkType = "YOUTUBE"
	LinkTypeFacebook  LinkType = "FACEBOOK"
	LinkTypeTiktok    LinkType = "TIKTOK"
)

var AllLinkType = []LinkType{
	LinkTypeWebsite,
	LinkTypeDiscord,
	LinkTypeInstagram,
	LinkTypeYoutube,
	LinkTypeFacebook,
	LinkTypeTiktok,
}

func (e LinkType) IsValid() bool {
	switch e {
	case LinkTypeWebsite, LinkTypeDiscord, LinkTypeInstagram, LinkTypeYoutube, LinkTypeFacebook, LinkTypeTiktok:
		return true
	}
	return false
}

func (e LinkType) String() string {
	return string(e)
}

func (e *LinkType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LinkType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LinkType", str)
	}
	return nil
}

func (e LinkType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OnsaleType string

const (
	OnsaleTypeEth   OnsaleType = "ETH"
	OnsaleTypeWeth  OnsaleType = "WETH"
	OnsaleTypeOxBtc OnsaleType = "oxBTC"
)

var AllOnsaleType = []OnsaleType{
	OnsaleTypeEth,
	OnsaleTypeWeth,
	OnsaleTypeOxBtc,
}

func (e OnsaleType) IsValid() bool {
	switch e {
	case OnsaleTypeEth, OnsaleTypeWeth, OnsaleTypeOxBtc:
		return true
	}
	return false
}

func (e OnsaleType) String() string {
	return string(e)
}

func (e *OnsaleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OnsaleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OnsaleType", str)
	}
	return nil
}

func (e OnsaleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SearchType string

const (
	SearchTypeGame       SearchType = "GAME"
	SearchTypeVideo      SearchType = "VIDEO"
	SearchTypeAnimation  SearchType = "ANIMATION"
	SearchTypePhotogrphy SearchType = "PHOTOGRPHY"
	SearchTypeAll        SearchType = "ALL"
)

var AllSearchType = []SearchType{
	SearchTypeGame,
	SearchTypeVideo,
	SearchTypeAnimation,
	SearchTypePhotogrphy,
	SearchTypeAll,
}

func (e SearchType) IsValid() bool {
	switch e {
	case SearchTypeGame, SearchTypeVideo, SearchTypeAnimation, SearchTypePhotogrphy, SearchTypeAll:
		return true
	}
	return false
}

func (e SearchType) String() string {
	return string(e)
}

func (e *SearchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SearchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SearchType", str)
	}
	return nil
}

func (e SearchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VerifyType string

const (
	VerifyTypeTwitter   VerifyType = "TWITTER"
	VerifyTypeInstagram VerifyType = "INSTAGRAM"
)

var AllVerifyType = []VerifyType{
	VerifyTypeTwitter,
	VerifyTypeInstagram,
}

func (e VerifyType) IsValid() bool {
	switch e {
	case VerifyTypeTwitter, VerifyTypeInstagram:
		return true
	}
	return false
}

func (e VerifyType) String() string {
	return string(e)
}

func (e *VerifyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VerifyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VerifyType", str)
	}
	return nil
}

func (e VerifyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WalletType string

const (
	WalletTypeBank   WalletType = "BANK"
	WalletTypeCoin   WalletType = "COIN"
	WalletTypeWallet WalletType = "WALLET"
)

var AllWalletType = []WalletType{
	WalletTypeBank,
	WalletTypeCoin,
	WalletTypeWallet,
}

func (e WalletType) IsValid() bool {
	switch e {
	case WalletTypeBank, WalletTypeCoin, WalletTypeWallet:
		return true
	}
	return false
}

func (e WalletType) String() string {
	return string(e)
}

func (e *WalletType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WalletType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid walletType", str)
	}
	return nil
}

func (e WalletType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
