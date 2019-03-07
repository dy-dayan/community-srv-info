package dal

import "github.com/dy-dayan/community-srv-info/global"

type Community struct {
	ID             string    `bson:"_id"`
	Name           string    `bson:"name"`
	SerialNumber   string    `bson:"serial_number"`
	Provinces      string    `bson:"provinces"`
	City           string    `bson:"city"`
	Region         string    `bson:"region"`
	Street         string    `bson:"street"`
	CouncilID      string    `bson:"council_id"` //居委会ID
	organizationID string    `bson:"organization_id"`
	house_count    int32     `bson:"house_count"`
	CheckInCount   int32     `bson:"check_in_count"`
	BuildingArea   int32     `bson:"building_area"`
	GreeningArea   int32     `bson:"greening_area"`
	SealedState    int32     `bson:"sealed_state"`
	Loc            []float32 `bson:"loc"`
	State          int32     `bson:"state"`
	Del            int32     `bson:"del"`
	CreatedAt      int       `bson:"created_at"`
	UpdatedAt      int       `bson:"updated_at"`
	OperatorID     int64     `bson:"operator_id"`
}

func InsertOne(cInfo *Community) error {
	ses := global.Mgo().Copy()
	defer ses.Close()
	return nil
}
