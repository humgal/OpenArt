package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/humgal/art-server/dao/bid"
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
	var id int
	row := db.DB.QueryRow("select id from user where username=?", username)

	row.Scan(&id)
	dao.ID = strconv.Itoa(id)
	redis.Rdb.HSet(context.Background(), "openart:user"+dao.ID, "id", dao.ID, "username", dao.Username, "realname", dao.Realname, "avatar", dao.Avatar, "phone", dao.Phone, "company", dao.Company, "email", dao.Email, "bio", dao.Bio, "img", dao.Img, "verifyType", dao.VerifyType, "verifyName", dao.VerifyName, "isCreator", dao.IsCreator, "followerNum", dao.FollowerNum, "followingNum", dao.FollowingNum, "links", dao.Links)

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
		var item model.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Tag, &item.Description, &item.SaleStatus, &item.CreatorID, &item.Creator, &item.CreateDate, &item.CollectionID)
		if err != nil {
			util.Logger.Fatal(err)
		}
		redis.Rdb.HSet(context.Background(), "openart:item:"+item.ID, "id", item.ID, "name", item.Name, "tag", item.Tag, "description", item.Description, "uploadUrl", item.UploadURL, "saleStatus", item.SaleStatus, "creatorId", item.CreatorID, "creator", item.Creator, "createDate", item.CreateDate, "collectionId", item.CollectionID)

		res_items = append(res_items, &item)
	}

	return res_items, err
}

// SetPrice is the resolver for the setPrice field.
func SetPrice(param *model.PriceParam) (bool, error) {
	var price item.ItemPrice
	price.InitPrice = decimal.NewFromFloat(param.InitPrice)

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
	db.DB.Exec("update item set price=" + strconv.Itoa(int(insertid)) + " where id=" + param.ItemID)

	return true, err
}

func PlaceBid(param *model.BidParm, username string) (bool, error) {
	var biddao bid.Bid
	biddao.ItemId = param.ItemID
	biddao.Total = decimal.NewFromFloat(param.Total)
	biddao.Username = username
	row := db.DB.QueryRow("select service_fee,item_price_type,onsale_type,start_date,expiration_date from item_price left join item on item_price.id=item.price where item.id=" + strconv.Itoa(param.ItemID))
	var startdatestr string
	var expdate int

	row.Scan(&biddao.ServiceFee, &biddao.ItemPriceType, &biddao.OnsaleType, &startdatestr, &expdate)
	t, _ := time.Parse("2006-01-02 15:04:05", startdatestr)
	t = t.Add(time.Duration(expdate))
	enddate := t.Format("2006-01-02 15:04:05")
	biddao.BidEndDate = &enddate
	biddao.BidDate = time.Now().Format("2006-01-02 15:04:05")
	_, err := db.DB.Exec(db.GenInsertSql(biddao, "bid"))
	if err != nil {
		util.Logger.Println(err)
		return false, err
	}
	return true, nil
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

	var dao model.User
	res, err := redis.Rdb.HGetAll(context.Background(), "openart:user"+id).Result()
	if err != nil {
		util.Logger.Println(err)
	}
	if res["id"] != "" {
		dao.ID = res["id"]
		dao.Username = res["username"]
		dao.Realname = res["realname"]
		*dao.Avatar = res["avatar"]
		*dao.Phone = res["phone"]
		*dao.Company = res["company"]
		*dao.Email = res["email"]
		*dao.Bio = res["bio"]
		*dao.Img = res["img"]
		// dao.VerifyType = res["verifyType"]
		*dao.VerifyName = res["verifyName"]
		// *dao.IsCreator = res["isCreator"]
		a, _ := strconv.Atoi(res["followerNum"])
		dao.FollowerNum = &a
		b, _ := strconv.Atoi(res["followingNum"])
		dao.FollowingNum = &b
		// dao.Links = res["links"]

		return &dao, err
	}
	sql := "select id,realname,username,avatar,phone,company,email,bio,img,verify_type,verify_name,is_creator,follower_num,following_num,links from user where id=?"

	row := db.DB.QueryRow(sql, id)

	var linkstr *string
	var linksjson []*model.Link
	err = row.Scan(&dao.ID, &dao.Realname, &dao.Username, &dao.Avatar, &dao.Phone, &dao.Company, &dao.Email, &dao.Bio, &dao.Img, &dao.VerifyType, &dao.VerifyName, &dao.IsCreator, &dao.FollowerNum, &dao.FollowingNum, &linkstr)
	if err != nil {
		util.Logger.Println(err)
		return &dao, err
	}
	if linkstr != nil {
		json.Unmarshal([]byte(*linkstr), &linksjson)
		dao.Links = linksjson
	}

	redis.Rdb.HSet(context.Background(), "openart:user"+dao.ID, "id", dao.ID, "username", dao.Username, "realname", dao.Realname, "avatar", dao.Avatar, "phone", dao.Phone, "company", dao.Company, "email", dao.Email, "bio", dao.Bio, "img", dao.Img, "verifyType", dao.VerifyType, "verifyName", dao.VerifyName, "isCreator", dao.IsCreator, "followerNum", dao.FollowerNum, "followingNum", dao.FollowingNum, "links", dao.Links)
	return &dao, err
}

// Item is the resolver for the item field.
func Item(id string) (*model.Item, error) {
	var item model.Item
	res, err := redis.Rdb.HGetAll(context.Background(), "openart:item"+id).Result()
	if err != nil {
		util.Logger.Println(err)
	}
	if res["ID"] != "" {
		item.ID = res["id"]
		item.Name = res["name"]
		*item.Tag = res["tag"]
		*item.Description = res["description"]
		item.UploadURL = res["uploadUrl"]
		a, _ := strconv.Atoi(res["saleStatus"])
		item.SaleStatus = a
		item.CreatorID = res["creatorId"]
		item.Creator = res["creator"]
		b, err := strconv.Atoi(res["collectionId"])
		item.CollectionID = &b
		return &item, err
	}

	sql := "select id,name,tag,description,upload_url,sale_status,creator_id,creator,create_date,collection_id from item where id=?"
	row := db.DB.QueryRow(sql, id)
	err = row.Scan(&item.ID, &item.Name, &item.Tag, &item.Description, &item.UploadURL, &item.SaleStatus, &item.CollectionID, &item.Creator, &item.CreateDate, &item.CollectionID)
	if err != nil {
		util.Logger.Println(err)
		return &item, err
	}
	redis.Rdb.HSet(context.Background(), "openart:item:"+item.ID, "id", item.ID, "name", item.Name, "tag", item.Tag, "description", item.Description, "uploadUrl", item.UploadURL, "saleStatus", item.SaleStatus, "creatorId", item.CreatorID, "creator", item.Creator, "createDate", item.CreateDate, "collectionId", item.CollectionID)
	return &item, err

}

// Collection is the resolver for the collection field.
func Collection(creator string) ([]*model.Collection, error) {
	var colls []*model.Collection

	res, err := redis.Rdb.Get(context.Background(), "openart:collection:"+creator).Result()
	if err != nil {
		util.Logger.Println(err)
	}
	if res != "" {
		json.Unmarshal([]byte(res), &colls)
		return colls, err
	}

	sql := "select  id,name,creator,create_date from collection where creator=?"
	row, err := db.DB.Query(sql, creator)
	if err != nil {
		util.Logger.Println(err)
	}
	for row.Next() {
		var coll model.Collection
		err = row.Scan(&coll.ID, &coll.Name, &coll.Createor, &coll.CreateDate)
		colls = append(colls, &coll)
		if err != nil {
			util.Logger.Println(err)
			return colls, err
		}
	}
	if len(colls) > 0 {
		collbyte, _ := json.Marshal(colls)
		redis.Rdb.Set(context.Background(), "openart:collection:"+creator, collbyte, time.Hour*24*30*3)
	}

	return colls, err
}

// Items is the resolver for the items field.
func Items(creator string) ([]*model.Item, error) {

	var items []*model.Item

	res, err := redis.Rdb.Get(context.Background(), "openart:items:"+creator).Result()
	if err != nil {
		util.Logger.Println(err)
	}
	if res != "" {
		json.Unmarshal([]byte(res), &items)
		return items, err
	}

	sql := "select  id,name,tag,description, upload_url,sale_status,creator_id,creator,create_date,collection_id from item where creator=?"
	row, err := db.DB.Query(sql, creator)
	if err != nil {
		util.Logger.Println(err)
	}
	for row.Next() {
		var item model.Item
		err = row.Scan(&item.ID, &item.Name, &item.Tag, &item.Description, &item.UploadURL, &item.SaleStatus, &item.CreatorID, &item.Creator, &item.CreateDate, &item.CollectionID)
		items = append(items, &item)
		if err != nil {
			util.Logger.Println(err)
			return items, err
		}
	}
	if len(items) > 0 {
		collbyte, _ := json.Marshal(items)
		redis.Rdb.Set(context.Background(), "openart:items:"+creator, collbyte, time.Hour*24*30)
	}
	return items, err
}

// SubscriptionPayment is the resolver for the subscriptionPayment field.
func SubscriptionPayment(itemid *string) (<-chan *model.SubscriptionEvent, error) {
	panic(fmt.Errorf("not implemented: SubscriptionPayment - subscriptionPayment"))
}

func SubscriptionBid(itemid *string) (<-chan []*model.Bid, error) {
	panic(fmt.Errorf("not implemented: SubscriptionBid - subscriptionBid"))
}
