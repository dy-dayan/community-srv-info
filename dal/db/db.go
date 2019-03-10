package db

import (
	"github.com/dy-dayan/community-srv-info/util"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	defaultMgo *mgo.Session
)

const (
	DCommunity = "dayan_community"
)

func Mgo() *mgo.Session {
	return defaultMgo
}

func Init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     util.DefaultMgoConf.Addr,
		Direct:    false,
		Timeout:   time.Second * 3,
		PoolLimit: util.DefaultMgoConf.PoolLimit,
		Username:  util.DefaultMgoConf.Username,
		Password:  util.DefaultMgoConf.Password,
	}

	ses, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Fatalf("dail mgo server error:%v", err)
	}

	ses.SetMode(mgo.Monotonic, true)
	defaultMgo = ses
}
