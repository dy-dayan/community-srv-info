package main

import (
	"github.com/dy-dayan/community-srv-info/dal/db"
	"github.com/dy-dayan/community-srv-info/idl/dayan/community/srv-info"
	"github.com/dy-gopkg/kit"
	"github.com/dy-dayan/community-srv-info/handler"
	"github.com/dy-dayan/community-srv-info/util/config"
	"github.com/sirupsen/logrus"
)

func main(){
	kit.Init()

	// 初始化配置
	uconfig.Init()

	// 初始化数据库
	db.Init()

	//TODO 初始化缓存
	//cache.CacheInit()

	err := dayan_community_srv_info.RegisterCommunityInfoHandler(kit.Server(), &handler.Handle{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}

	kit.Run()
}
