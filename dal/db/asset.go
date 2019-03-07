package db

type Asset struct {
	ID          string `bson:"_id"`
	Serial      string `bson:"serial"`
	Category    int32  `bson:"category"`
	location    string `bson:"location"`
	state       int32  `bson:"state"`
	CommunityID string `bson:"community_id"`
	Brand       string `bson:"brand"`
	Desc        string `bson:"desc"`
}
