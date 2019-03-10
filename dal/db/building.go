package db

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CBuilding = "building"
)

//Building 楼宇
type Building struct {
	ID          int64     `bson:"_id"`
	Name        string    `bson:"name"`
	Loc         []float32 `bson:"loc"`
	ElevatorIDs []int64   `bson:"elevator_ids"`
	CommunityID int64     `bson:"community_id"`
	Period      int32     `bson:"period"`
	CreatedAt   int64     `bson:"created_at"`
	UpdatedAt   int64     `bson:"updated_at"`
	OperatorID  int64     `bson:"operator_id"`
}

//UpSertBuilding 插入一个建筑物信息
func UpsertBuilding(building *Building) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		return errors.New("mgo session is nil")
	}
	defer ses.Close()
	query := bson.M{
		"_id": building.ID,
	}
	now := time.Now().Unix()
	change := mgo.Change{
		Update: bson.M{
			"$set": building,
			"SsetOnInsert": bson.M{
				"created_at": now,
			},
		},
		Upsert:    true,
		Remove:    false,
		ReturnNew: false,
	}

	_, err := ses.DB(DCommunity).C(CBuilding).Find(query).Apply(change, nil)
	return err
}

//DelBuildingByID 删除建筑物
func DelBuildingByID(id int64) error {
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
				"del":        1,
				"updated_at": now,
			},
		},
		Upsert:    false,
		Remove:    false,
		ReturnNew: false,
	}
	_, err := ses.DB(DCommunity).C(CBuilding).Find(query).Apply(change, nil)
	return err
}

//GetBuildingByID 获得具体的建筑物信息
func GetBuildingByID(id int64) (*Building, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": id,
	}
	ret := &Building{}
	err := ses.DB(DCommunity).C(CBuilding).Find(query).One(ret)
	return ret, err
}

//GetBuilding  获得的社区的建筑物
func GetBuilding(limit, offset int, communityID int64) (*[]Building, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()
	query := bson.M{
		"community_id": communityID,
	}

	ret := &[]Building{}
	err := ses.DB(DCommunity).C(CBuilding).Find(query).Skip(offset).Limit(limit).All(ret)
	return ret, err
}
