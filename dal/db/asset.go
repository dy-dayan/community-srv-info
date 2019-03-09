package db

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CAsset = "asset"
)

//Asset 资产表
type Asset struct {
	ID          int64  `bson:"_id"`
	Serial      string `bson:"serial"`
	Category    int32  `bson:"category"`
	Loc         string `bson:"loc"`
	State       int32  `bson:"state"`
	CommunityID int64  `bson:"community_id"`
	Brand       string `bson:"brand"`
	Desc        string `bson:"desc"`
	CreatedAt   int64  `bson:"created_at"`
	UpdatedAt   int64  `bson:"updated_at"`
	OperatorID  int64  `bson:"operator_id"`
}

func UpsertAsset(asset *Asset) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		return errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": asset.ID,
	}

	now := time.Now().Unix()
	asset.UpdatedAt = now
	change := mgo.Change{
		Update: bson.M{
			"$set": asset,
			"$setOnInsert": bson.M{
				"created_at": now,
			},
		},
		Upsert:    true,
		Remove:    false,
		ReturnNew: false,
	}
	_, err := ses.DB(DCommunity).C(CAsset).Find(query).Apply(change, nil)
	return err
}

func DelAssetByID(id int64) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		return errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": id,
	}

	now := time.Now().Unix()
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
	_, err := ses.DB(DCommunity).C(CAsset).Find(query).Apply(change, nil)
	return err
}

func GetAssetByID(id int64) (*Asset, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": id,
	}
	ret := &Asset{}
	err := ses.DB(DCommunity).C(CAsset).Find(query).One(ret)
	return ret, err
}
