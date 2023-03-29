package service

import (
	"fmt"
	"strconv"

	"github.com/humgal/art-server/dao/item"
	"github.com/humgal/art-server/dao/users"
	"github.com/humgal/art-server/db"
	"github.com/humgal/art-server/graph/model"
	"github.com/humgal/art-server/util"
	"github.com/shopspring/decimal"
)

func UpdateUser(user *model.UpdateUser, username string) (string, error) {
	var dao users.User
	dao.Username = username

	dao.Realname = user.Realname
	dao.Avatar = user.Avatar
	dao.Phone = user.Phone
	dao.Company = user.Company
	dao.Email = user.Email
	dao.Bio = user.Bio
	dao.Img = user.Img
	dao.VerifyName = user.VerifyName
	dao.Links = user.Links

	err := dao.UpdateUser()
	if err != nil {
		return "update user fail", err
	}
	return "update user success", err
}

// UploadArt is the resolver for the uploadArt field.
func UploadArt(items []*model.UploadItem) ([]*model.Item, error) {
	sql := ""
	for i := 0; i < len(items); i++ {

		var item item.Item
		item.Name = items[i].Name
		item.Tag = items[i].Tag
		item.Description = items[i].Description
		item.UploadUrl = items[i].UploadURL
		item.CreatorId = items[i].CreateID
		item.Creator = items[i].Creator
		item.SaleStatus = items[i].SaleStatus

		sql += db.GenInsertSql(item, "item") + ";"
	}
	res, err := db.DB.Exec(sql)
	if err != nil {
		return []*model.Item{}, err
	}
	insid, _ := res.LastInsertId()
	rows, err := db.DB.Query("select id,name,tag,description,sale_status,creator_id,creator,create_date,collection_id from item where id=?", insid)
	var res_items []*model.Item
	for rows.Next() {
		var res_item model.Item
		err := rows.Scan(&res_item.ID, &res_item.Name, &res_item.Tag, &res_item.Description, &res_item.SaleStatus, &res_item.CreateorID, &res_item.Creator, &res_item.CreateDate, &res_item.CollectionID)
		if err != nil {
			util.Logger.Fatal(err)
		}
		res_items = append(res_items, &res_item)
	}

	return res_items, err
}

// SetPrice is the resolver for the setPrice field.
func SetPrice(param *model.PriceParam) (bool, error) {
	var price item.ItemPrice
	price.InitPrice = decimal.NewFromFloat(param.InitPrice)
	id, err := strconv.ParseInt(param.ItemID, 0, 32)
	if err != nil {
		util.Logger.Println(err)
	}
	price.ItemId = int(id)
	if param.ItemPriceType == "FIXED" {
		price.ItemPriceType = 1
	} else {
		price.ItemPriceType = 2
	}
	price.OnsaleType = param.OnsaleType.String()
	price.StartDate = param.StartDate
	res, err := db.DB.Exec(db.GenInsertSql(price, "item_price"))
	if err != nil {
		util.Logger.Println(err)
	}
	insertid, _ := res.LastInsertId()
	db.DB.Exec("update item set price=" + strconv.Itoa(int(insertid)))

	return true, err
}

// MintArt is the resolver for the mintArt field.
func MintArt(items []string) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: MintArt - mintArt"))
}

func PlaceBid(bid *model.BidParm) (*model.Bid, error) {

	return &model.Bid{}, nil
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
