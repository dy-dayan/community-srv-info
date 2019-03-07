package db

import (
	"github.com/dy-dayan/community-srv-info/util/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	defaultMgo *mgo.Session
)

const (
	DCommunity = "dayan_community"
	CCommunityInfo = "community_info"
)

func Mgo() *mgo.Session {
	return defaultMgo
}

func Init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     uconfig.DefaultMgoConf.Addr,
		Direct:    false,
		Timeout:   time.Second * 3,
		PoolLimit: uconfig.DefaultMgoConf.PoolLimit,
		Username:  uconfig.DefaultMgoConf.Username,
		Password:  uconfig.DefaultMgoConf.Password,
	}

	ses, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Fatalf("dail mgo server error:%v", err)
	}

	ses.SetMode(mgo.Monotonic, true)
	defaultMgo = ses
}
