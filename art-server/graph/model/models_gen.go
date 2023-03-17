// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Bid struct {
	ItemID     int      `json:"itemId"`
	Balance    float64  `json:"balance"`
	ServiceFee *float64 `json:"serviceFee"`
	Total      *float64 `json:"total"`
}

type Collection struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Items      []*Item `json:"items"`
	CreateDate *string `json:"createDate"`
}

type Item struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Tag         *string    `json:"tag"`
	Description *string    `json:"description"`
	UploadURL   string     `json:"uploadUrl"`
	SaleStatus  int        `json:"saleStatus"`
	Price       *ItemPrice `json:"price"`
	CreatorID   int        `json:"creatorId"`
	Creator     string     `json:"creator"`
	CreateDate  *string    `json:"createDate"`
}

type ItemPrice struct {
	Type           int        `json:"type"`
	Onsale         OnsaleCoin `json:"onsale"`
	InitPrice      float64    `json:"initPrice"`
	StartDate      *string    `json:"startDate"`
	ExpirationDate *string    `json:"expirationDate"`
}

type Link struct {
	Type LinkType `json:"type"`
	URL  string   `json:"url"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type PriceRange struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type SearchParm struct {
	Param   string      `json:"param"`
	Type    *SearchType `json:"type"`
	Price   *PriceRange `json:"price"`
	Chain   *Blockchain `json:"chain"`
	Onsale  *OnsaleCoin `json:"onsale"`
	Creator *string     `json:"creator"`
}

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	UserName   string      `json:"userName"`
	Photo      *string     `json:"photo"`
	Phone      *string     `json:"phone"`
	Company    *string     `json:"company"`
	Email      *string     `json:"email"`
	Bio        *string     `json:"bio"`
	Img        *string     `json:"img"`
	JoinDate   *string     `json:"joinDate"`
	VerifyType *VerifyType `json:"verifyType"`
	VerifyName *string     `json:"verifyName"`
	Links      []*Link     `json:"links"`
}

type Payment struct {
	ItemID  int     `json:"itemId"`
	PayType int     `json:"payType"`
	Price   float64 `json:"price"`
	PayDate *string `json:"payDate"`
}

type Wallet struct {
	Type     WalletType `json:"type"`
	PubToken string     `json:"pubToken"`
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
		return fmt.Errorf("%s is not a valid blockchain", str)
	}
	return nil
}

func (e Blockchain) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OnsaleCoin string

const (
	OnsaleCoinEth   OnsaleCoin = "ETH"
	OnsaleCoinWeth  OnsaleCoin = "WETH"
	OnsaleCoinOxBtc OnsaleCoin = "oxBTC"
)

var AllOnsaleCoin = []OnsaleCoin{
	OnsaleCoinEth,
	OnsaleCoinWeth,
	OnsaleCoinOxBtc,
}

func (e OnsaleCoin) IsValid() bool {
	switch e {
	case OnsaleCoinEth, OnsaleCoinWeth, OnsaleCoinOxBtc:
		return true
	}
	return false
}

func (e OnsaleCoin) String() string {
	return string(e)
}

func (e *OnsaleCoin) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OnsaleCoin(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid onsaleCoin", str)
	}
	return nil
}

func (e OnsaleCoin) MarshalGQL(w io.Writer) {
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
