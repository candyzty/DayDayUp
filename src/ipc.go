package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"os/exec"
	"time"
	"gopkg.in/mgo.v2/bson"
)

//摄像机表结构
type Camera struct {
	Id           bson.ObjectId `bson:"_id"`
	CameraSn     string        `bson:"cameraSn"`
	Name         string        `bson:"name"`
	IpCameraType string        `bson:"ipCameraType"`
	IpCameraDesc string        `bson:"ipCameraDesc"`
	DeviceStatus string        `bson:"deviceStatus"`
	ProjectId    bson.ObjectId `bson:"projectId"`
	Iccid        string        `bson:"iccid"`
	GmtCreated   time.Time     `bson:"gmtCreated"`
	GmtUpdated   time.Time     `bson:"gmtUpdated"`
}

//设备表结构
type Device struct {
	Id           bson.ObjectId `bson:"_id"`
	ProjectId    bson.ObjectId `bson:"projectId"`
	ProductId    bson.ObjectId `bson:"productId"`
	Name         string        `bson:"name"`
	DeviceSn     string        `bson:"deviceSn"`
	Sensors      []string      `bson:"sensors"`
	Description  string        `bson:"description"`
	CityName     string        `bson:"cityName"`
	DeviceStatus string        `bson:"deviceStatus"`
	Iccid        string        `bson:"iccid"`
	GmtCreated   time.Time     `bson:"gmtCreated"`
	GmtUpdated   time.Time     `bson:"gmtUpdated"`
}

//摄像机资源表
type CameraResource struct {
	Id         bson.ObjectId `bson:"_id"`
	OssKey     string        `bson:"ossKey"`
	ResType    string        `bson:"resType"`
	ResClass   string        `bson:"resClass"`
	IpCameraSn string        `bson:"ipCameraSn"`
	GmtCreated time.Time     `bson:"gmtCreated"`
	GmtUpdated time.Time     `bson:"gmtUpdated"`
}

//ipc数据库连接
var c *mgo.Collection

//dmp数据库连接
var b *mgo.Collection
var d *mgo.Collection
var session *mgo.Session

func (camera Camera) ToString() string {
	return fmt.Sprintf("%#v", camera)
}

var opt string
var env string       //环境
var exportCmd string //环境
var importCmd string //环境
func init() {
	//添加一些动态修改的参数

	flag.StringVar(&env, "env", "test", "默认是test")
	flag.StringVar(&opt, "opt", "0", "默认是0")
	flag.Parse()
	fmt.Println("env=", env)

	switch env {
	case "test":
		//测试配置
		exportCmd = "mongoexport --uri=\"mongodb://root:jBK8HL7OGVllZk96@39.107.244.43:8007/db_ipc?authSource=admin&readPreference=primary\"  --collection=t_ip_camera_resource  --out=t_ip_camera_resource.dat"
		importCmd = "mongoimport --uri=\"mongodb://root:jBK8HL7OGVllZk96@39.107.244.43:8007/db_dmp?authSource=admin&readPreference=primary\"  --collection=t_ip_camera_resource  --file=t_ip_camera_resource.dat"

		session, _ := mgo.Dial("mongodb://root:jBK8HL7OGVllZk96@39.107.244.43:8007/db_ipc?authSource=admin")
		//切换数据库到IPC
		//切换到collection
		c = session.DB("db_ipc").C("t_ip_camera")
		//切换数据库到dmp
		dmp := session.DB("db_dmp")
		//切换到collection
		b = dmp.C("t_device")
		//切换到collection
		d = dmp.C("t_product")
	case "prod":
		exportCmd = "mongoexport --uri=\"mongodb://skywalker:eKYzUvcKT28Py@dds-2zec957064b792a41.mongodb.rds.aliyuncs.com:3717/db_ipc?replicaSet=mgset-12833163\"  --collection=t_ip_camera_resource  --out=t_ip_camera_resource.dat"
		importCmd = "mongoimport --uri=\"mongodb://skywalker:eKYzUvcKT28Py@dds-2zec957064b792a41.mongodb.rds.aliyuncs.com:3717/db_dmp?replicaSet=mgset-12833163\"  --collection=t_ip_camera_resource  --file=t_ip_camera_resource.dat"

		session, _ := mgo.Dial("mongodb://skywalker:eKYzUvcKT28Py@dds-2zec957064b792a41.mongodb.rds.aliyuncs.com:3717,dds-2zec957064b792a42.mongodb.rds.aliyuncs.com:3717/db_ipc?replicaSet=mgset-12833163")
		//切换数据库到IPC //切换到collection
		c = session.DB("db_ipc").C("t_ip_camera")
		//切换数据库到dmp
		dmp := session.DB("db_dmp")
		//切换到collection
		b = dmp.C("t_device")
		d = dmp.C("t_product")
	default:
		fmt.Println("opt  Command not support")
	}

}

//查询ipc摄像机数据   插入到dmp设备表里面
func find() {
	var cameraArr []Camera
	var product Device //用设备直接接收只需要可以封装ID即可
	//查出产品id是多少
	d.Find(bson.M{"deviceNodeType": "IPC_DEVICE"}).One(&product)
	fmt.Println("productId is =", product.Id.Hex())

	//查出所有的摄像机数据
	c.Find(nil).All(&cameraArr)
	for _, value := range cameraArr {
		fmt.Println(value.ToString())
		//组装设备然后进行插入
		var device = new(Device)
		device.Id = bson.NewObjectId()
		device.Name = value.Name
		device.DeviceSn = value.CameraSn
		device.ProjectId = value.ProjectId
		//产品ID
		device.ProductId = product.Id
		device.Iccid = value.Iccid
		device.DeviceStatus = value.DeviceStatus
		device.Description = value.IpCameraDesc
		device.GmtCreated = time.Now()
		device.GmtUpdated = time.Now()
		err := b.Insert(device)
		if err == nil {
			fmt.Println("摄像机插入成功", value.CameraSn)
		} else {
			fmt.Println("摄像机插入失败", value.CameraSn, err.Error())
			defer panic(err)
		}
	}

}

//查询所有的摄像机数据
func findAll() {
	var CameraArr []Camera
	//查出所有的摄像机数据
	c.Find(nil).All(&CameraArr)
	for _, value := range CameraArr {
		fmt.Println(value.ToString())
	}

}

//执行一个shell命令
func execCmd(exe string) {
	cmd := exec.Command("/bin/bash", "-c", exe)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute cmd failed:" + err.Error())
		return
	}
}

func main() {
	fmt.Println("opt handle is ", opt)
	// 0 查询数据库连接是否正常
	// 1 备份数据库表
	// 2 恢复到另一个数据库
	// 3 查询摄像机表并插入到设备表
	switch opt {
	case "0":
		//查出所有的摄像机测试连接是否异常
		findAll()
	case "1":
		//执行导出命令
		execCmd(exportCmd)
		fmt.Println("Execute exportCmd finished.")
	case "2":
		//执行导入命令
		execCmd(importCmd)
		fmt.Println("Execute importCmd finished.")
	case "3":
		//查找并插入
		find()
	default:
		fmt.Println("opt  Command not support")
	}

}
