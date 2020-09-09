package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var sourceUrl string = "mongodb://root:jBK8HL7OGVllZk96@39.107.244.43:8007/db_ipc?authSource=admin"
var targetUrl string = "mongodb://root:jBK8HL7OGVllZk96@39.107.244.43:8007/db_dmp?authSource=admin"
/**

 */
/**
源对象结构体
*/
type Source struct {
	Id bson.ObjectId `bson:"_id"`
	CameraSn string `bson:"cameraSn"`
	Name string `bson:"name"`
	IpCameraType string `bson:"ipCameraType"`
	DeviceStatus string `bson:"deviceStatus"`
	ProjectId bson.ObjectId `bson:"projectId"`
	ServerIp string `bson:"serverIp"`
	Iccid string `bson:"iccis"`
	GmtCreated mgo.Database `bson:"gmtCreated"`
	GmtUpdated mgo.Database `bson:"gmtUpdated"`
}

func (source Source) toString() {
	fmt.Printf("名称 : %s , 摄像机类型 : %s \n",source.Name, source.IpCameraType)
}

/**
目标对象结构体
*/
type Target struct {
}

type test1 struct {
	Id bson.ObjectId `bson:"_id"`
}
/**
测试对象结构体
 */
type Test struct {
	test1
	X int `bson:"x"`
	Y int `bson:"y"`

}
func (test Test) toString() {
	fmt.Printf("id:%s,x:%d,y:%d",test.Id.String(), test.X, test.Y)
}

func startSource(url string, dataBaseName string, tableName string) []Source{
	session, err := mgo.Dial(url)
	if err != nil {
		log.Println(err)
	}
	//延迟函数
	defer session.Close()
	//设置模式
	session.SetMode(mgo.Monotonic, true)
	//选择访问数据库,表
	collection := session.DB(dataBaseName).C(tableName)
	var sourceList []Source
	collection.Find(nil).All(&sourceList)
	return sourceList
}

func startTarget(url string, dataBaseName string, tableName string) []Target{
	session, err := mgo.Dial(url)
	if err != nil {
		log.Println(err)
	}
	//延迟函数
	defer session.Close()
	//设置模式
	session.SetMode(mgo.Monotonic, true)
	//选择访问数据库,表
	collection := session.DB(dataBaseName).C(tableName)
	var targetList []Target
	collection.Find(nil).All(&targetList)
	return targetList
}

func main() {
	sourceDataBaseName := "db_ipc"
	sourceTableName := "t_ip_camera"
	sourceList := startSource(sourceUrl, sourceDataBaseName, sourceTableName)
	for i := range sourceList {
		sourceList[i].toString()
	}
}
