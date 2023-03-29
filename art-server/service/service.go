package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/humgal/art-server/dao/collection"
	"github.com/humgal/art-server/dao/item"
	"github.com/humgal/art-server/dao/users"
	"github.com/humgal/art-server/db"
	"github.com/humgal/art-server/graph/model"
	"github.com/humgal/art-server/util"
	"github.com/humgal/art-server/util/redis"
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
		err := rows.Scan(&res_item.ID, &res_item.Name, &res_item.Tag, &res_item.Description, &res_item.SaleStatus, &res_item.CreatorID, &res_item.Creator, &res_item.CreateDate, &res_item.CollectionID)
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

func PlaceBid(bid *model.BidParm) (*model.Bid, error) {

	return &model.Bid{}, nil
}

// CreateCollection is the resolver for the createCollection field.
func CreateCollection(param model.CollectionParm) (bool, error) {
	var coll collection.Collection
	coll.Name = *param.Name
	coll.Creator = &param.Creator
	coll.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	res, err := db.DB.Exec(db.GenInsertSql(coll, "collection"))
	if err != nil {
		util.Logger.Println(err)
		return false, err
	}
	util.Logger.Println(res.LastInsertId())
	return true, err
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

	var user model.User
	res, err := redis.Rdb.HGet(redis.Rdb.Context(), "openart:user", id).Result()
	if err != nil {
		util.Logger.Println(err)
	}
	if res != "" {
		json.Unmarshal([]byte(res), &user)
		return &user, err
	}
	sql := "select id,realname,username,avatar,phone,company,email,bio,img,verify_type,verify_name,is_creator,follower_num,following_num,links from user where id=?"

	row := db.DB.QueryRow(sql, id)

	var linkstr *string
	var linksjson []*model.Link
	err = row.Scan(&user.ID, &user.Realname, &user.Username, &user.Avatar, &user.Phone, &user.Company, &user.Email, &user.Bio, &user.Img, &user.VerifyType, &user.VerifyName, &user.IsCreator, &user.FollowerNum, &user.FollowingNum, &linkstr)
	if err != nil {
		util.Logger.Println(err)
		return &user, err
	}
	if linkstr != nil {
		json.Unmarshal([]byte(*linkstr), &linksjson)
		user.Links = linksjson
	}

	userbyte, err := json.Marshal(user)
	if err != nil {
		util.Logger.Println(err)
	}
	redis.Rdb.HSet(redis.Rdb.Context(), "openart:user", id, userbyte)
	return &user, err
}

// Item is the resolver for the item field.
func Item(id string) (*model.Item, error) {
	var item model.Item
	res, err := redis.Rdb.HGet(redis.Rdb.Context(), "openart:item", id).Result()
	if err != nil {
		util.Logger.Println(err)
	}
	if res != "" {
		json.Unmarshal([]byte(res), &item)
		return &item, err
	}

	sql := "select id,name,tag,description,upload_url,sale_status,creator_id,creator,create_date,collection_id from item where id=?"
	row := db.DB.QueryRow(sql, id)
	err = row.Scan(&item.ID, &item.Name, &item.Tag, &item.Description, &item.UploadURL, &item.SaleStatus, &item.CollectionID, &item.Creator, &item.CreateDate, &item.CollectionID)
	if err != nil {
		util.Logger.Println(err)
		return &item, err
	}
	itembyte, err := json.Marshal(item)
	redis.Rdb.HSet(redis.Rdb.Context(), "openart:item", id, itembyte)

	return &item, err

}

// Collection is the resolver for the collection field.
func Collection(createor string) (*model.Collection, error) {
	var coll model.Collection

	res, err := redis.Rdb.HGet(redis.Rdb.Context(), "openart:collection", createor).Result()
	if err != nil {
		util.Logger.Println(err)
	}
	if res != "" {
		json.Unmarshal([]byte(res), &coll)
		return &coll, err
	}

	sql := "select  id,name,creator,create_date from collection where creator=?"
	row := db.DB.QueryRow(sql, createor)
	err = row.Scan(&coll.ID, &coll.Name, &coll.Createor, &coll.CreateDate)
	if err != nil {
		util.Logger.Println(err)
		return &coll, err
	}
	collbyte, err := json.Marshal(coll)
	redis.Rdb.HSet(redis.Rdb.Context(), "openart:collection", createor, collbyte)

	return &coll, err
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
