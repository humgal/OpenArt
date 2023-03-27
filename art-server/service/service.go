package service

import (
	"fmt"

	"github.com/humgal/art-server/graph/model"
)

// add user in database and add userlist in redis
func NewUser(user *model.NewUser) (string, error) {
	panic("")
}

func CheckUser(username string) (int, error) {
	panic("")
}

func Login(input model.Login) (string, error) {
	panic("")
}
func PlaceBid(bid *model.BidParm) (*model.Bid, error) {
	return &model.Bid{}, nil
}
func UpdateUser(user *model.UpdateUser) (string, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// UploadArt is the resolver for the uploadArt field.
func UploadArt(items []*model.UploadItem) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: UploadArt - uploadArt"))
}

// SetPrice is the resolver for the setPrice field.
func SetPrice(param *model.PriceParam) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: SetPrice - setPrice"))
}

// MintArt is the resolver for the mintArt field.
func MintArt(items []string) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: MintArt - mintArt"))
}

// CreateCollection is the resolver for the createCollection field.
func CreateCollection(param model.CollectionParm) (*model.Collection, error) {
	panic(fmt.Errorf("not implemented: CreateCollection - createCollection"))
}

// Checkout is the resolver for the checkout field.
func Checkout(param *model.PayParam) (*string, error) {
	panic(fmt.Errorf("not implemented: Checkout - checkout"))
}

// ConnectWallet is the resolver for the connectWallet field.
func ConnectWallet(userID string, typeArg model.WalletType) (*model.Wallet, error) {
	panic(fmt.Errorf("not implemented: ConnectWallet - connectWallet"))
}

// Follow is the resolver for the follow field.
func Follow(param *model.FollowParam) (*string, error) {
	panic(fmt.Errorf("not implemented: Follow - follow"))
}

// SearchItems is the resolver for the searchItems field.
func SearchItems(param model.SearchParm) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: SearchItems - searchItems"))
}

// User is the resolver for the user field.
func User(id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Item is the resolver for the item field.
func Item(id string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: Item - item"))
}

// Collection is the resolver for the collection field.
func Collection(createor string) (*model.Collection, error) {
	panic(fmt.Errorf("not implemented: Collection - collection"))
}

// Items is the resolver for the items field.
func Items(createor *string, ids []string) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// SubscriptionPayment is the resolver for the subscriptionPayment field.
func SubscriptionPayment(itemid *string) (<-chan *model.SubscriptionEvent, error) {
	panic(fmt.Errorf("not implemented: SubscriptionPayment - subscriptionPayment"))
}

func SubscriptionBid(itemid *string) (<-chan []*model.Bid, error) {
	panic(fmt.Errorf("not implemented: SubscriptionBid - subscriptionBid"))
}
