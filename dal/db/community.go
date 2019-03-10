package db

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (

	//CCommunityInfo  集合名称
	CCommunityInfo = "community_info"
)

//CommunityInfo 社区信息
type CommunityInfo struct {
	ID           int64     `bson:"_id"` // 小区编号
	Name         string    `bson:"name"`
	Province     string    `bson:"province"`
	City         string    `bson:"city"`
	Region       string    `bson:"region"`
	Street       string    `bson:"street"`
	OrgID        int64     `bson:"org_id"`
	HouseCount   int32     `bson:"house_count"`
	CheckInCount int32     `bson:"check_in_count"`
	BuildingArea float32   `bson:"building_area"`
	GreeningArea float32   `bson:"greening_area"`
	Loc          []float32 `bson:"loc"`
	State        int32     `bson:"state"`
	Del          int32     `bson:"del"`
	CreatedAt    int64     `bson:"created_at"`
	UpdatedAt    int64     `bson:"updated_at"`
	OperatorID   int64     `bson:"operator_id"`
}

func UpsertCommunityInfo(com *CommunityInfo) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		return errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": com.ID,
	}

	now := time.Now().Unix()
	com.UpdatedAt = now
	change := mgo.Change{
		Update: bson.M{
			"$set": com,
			"$setOnInsert": bson.M{
				"created_at": now,
			},
		},
		Upsert:    true,
		Remove:    false,
		ReturnNew: false,
	}

	_, err := ses.DB(DCommunity).C(CCommunityInfo).Find(query).Apply(change, nil)
	return err
}

func DelCommunityInfo(id int64) error {
	ses := defaultMgo.Copy()
	if ses == nil {
		return errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": id,
	}
	change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"del": 1,
			},
		},
		Upsert:    false,
		Remove:    false,
		ReturnNew: false,
	}
	_, err := ses.DB(DCommunity).C(CCommunityInfo).Find(query).Apply(change, nil)
	return err
}

func GetCommunityInfoByID(id int64) (*CommunityInfo, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()

	query := bson.M{
		"_id": id,
	}
	ret := &CommunityInfo{}
	err := ses.DB(DCommunity).C(CCommunityInfo).Find(query).One(ret)
	return ret, err
}

// todo: 添加根据地理位置获取小区列表
