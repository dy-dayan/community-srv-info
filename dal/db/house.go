package db

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CHouse = "house"
)

type House struct {
	ID         int64   `bson:"_id"`
	BuildingID int64   `bson:"building_id"`
	Unit       string  `bson:"unit"`
	Acreage    float32 `bson:"acreage"`
	State      int32   `bson:"state"`
	Rental     int32   `bson:"rental"`
	CreatedAt  int64   `bson:"created_at"`
	UpdatedAt  int64   `bson:"updated_at"`
	OperatorID int64   `bson:"operator_id"`
}

//UpsertHouse 插入房屋信息
func UpsertHouse(house *House) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		return errors.New("mgo session is nil")
	}
	defer ses.Close()

	now := time.Now().Unix()
	query := bson.M{
		"_id": house.ID,
	}

	change := mgo.Change{
		Update: bson.M{
			"$set": house,
			"$setOnInsert": bson.M{
				"created_at": now,
			},
		},
		Upsert:    true,
		Remove:    false,
		ReturnNew: false,
	}

	_, err := ses.DB(DCommunity).C(CHouse).Find(query).Apply(change, nil)
	return err
}

//DelHouseByID 删除房屋信息
func DelHouseByID(id int64) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		return errors.New("mgo session is nil")
	}
	defer ses.Close()

	now := time.Now().Unix()
	query := bson.M{
		"_id": id,
	}

	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"updated_at": now,
				"del":        1,
			},
		},
		Upsert:    false,
		Remove:    false,
		ReturnNew: false,
	}

	_, err := ses.DB(DCommunity).C(CHouse).Find(query).Apply(change, nil)
	return err
}

//GetHouseByID 获得具体社区信息
func GetHouseByID(id int64) (*House, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()
	query := bson.M{
		"_id": id,
	}
	ret := &House{}
	err := ses.DB(DCommunity).C(CHouse).Find(query).One(ret)
	return ret, err
}

//GetHouse 获得社区房屋信息
func GetHouse(limit, offset int, communityID int64) (*[]House, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"community_id": communityID,
	}

	ret := &[]House{}
	err := ses.DB(DCommunity).C(CHouse).Find(query).Skip(offset).Limit(limit).All(ret)
	return ret, err
}
