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
	ID          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Loc         []float32 `bson:"loc"`
	ElevatorIDs []int64   `bson:"elevator_ids"`
	CommunityID int64     `bson:"community_id"`
	Period      int32     `bson:"period"`
	CreatedAt   int64     `bson:"created_at"`
	UpdatedAt   int64     `bson:"updated_at"`
	OperatorID  int64     `bson:"operator_id"`
}

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

func GetBUildingByID(id int64) (*Building, error) {
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
