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
	SerialNumber string    `bson:"serial_number"`
	Province     string    `bson:"province"`
	City         string    `bson:"city"`
	Region       string    `bson:"region"`
	Street       string    `bson:"street"`
	SealedState  int32     `bson:"sealed_state"`
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

//UpsertCommnunityInfo 插入一个社区信息
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

//DelCommunityInfo 删除一个资产信息
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

//GetCommunityInfoByID 获得具体社区信息
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

//GetCommnunityInfo 获得所有社区信息
func GetCommunityInfo(limit, offset int) (*[]CommunityInfo, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	defer ses.Close()

	ret := &[]CommunityInfo{}
	err := ses.DB(DCommunity).C(CCommunityInfo).Find("{}").Skip(offset).Limit(limit).All(ret)
	return ret, err
}

//GetCommunityInfoByLoc 获得具体经纬度，范围内的社区信息
func GetCommunityInfoByLoc(limit, offset int, loc []float32, distance float32) (*[]CommunityInfo, error) {
	ses := defaultMgo.Copy()
	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}

	defer ses.Close()
	query := bson.M{
		"loc": bson.M{
			"$near":        loc,
			"$maxDistance": distance,
		},
	}
	ret := &[]CommunityInfo{}
	err := ses.DB(DCommunity).C(CCommunityInfo).Find(query).Skip(offset).Limit(limit).All(ret)
	return ret, err
}
