package global

import "gopkg.in/mgo.v2"

var defaultMgo *mgo.Session

//Mgo 返回MongoDB数据库链接
func Mgo()*mgo.Session{
	return defaultMgo
}

//逻辑初始化
func Init(){

}
