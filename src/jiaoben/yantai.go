package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {
	recentConnection()
	//dayConnection()
}

//数据库地址
var dataUrl string = "mongodb://skywalker:eKYzUvcKT28Py@192.168.1.154:27017/db_yantai?authSource=admin"
//var dataUrl string = "mongodb://root:jBK8HL7OGVllZk96@39.107.244.43:8007/db_yantai?authSource=admin"

//对象结构体
type DevStat struct {
	Id                            bson.ObjectId `bson:"_id"`
	AccumulatedQCooling           float64       `bson:"accumulatedQCooling"`
	AccumulatedQCoolingSource     float64       `bson:"accumulatedQCoolingSource"`
	AccumulatedQHeat              float64       `bson:"accumulatedQHeat"`
	AccumulatedQHeatSource        float64       `bson:"accumulatedQHeatSource"`
	AccumulatedWindCoolingMeter       float64       `bson:"accumulatedWindCoolingMeter"`
	AccumulatedWindCoolingMeterSource float64       `bson:"accumulatedWindCoolingMeterSource"`
	AccumulatedWindHeatMeter         float64       `bson:"accumulatedWindHeatMeter"`
	AccumulatedWindHeatMeterSource    float64       `bson:"accumulatedWindHeatMeterSource"`
	DeviceSn                      string        `bson:"deviceSn"`
	HeatPumpInstantaneousHeat     float64       `bson:"heatPumpInstantaneousHeat"`
	InstantaneousHeat             float64       `bson:"instantaneousHeat"`
	OutdoorTemperature            float64       `bson:"outdoorTemperature"`
	StartTime                     time.Time     `bson:"startTime"`
	GmtCreated                    time.Time     `bson:"gmtCreated"`
	GmtUpdated                    time.Time     `bson:"gmtUpdated"`
	AcState                       int64          `bson:"acState"`
}

type EleStat struct {
	Id                        bson.ObjectId `bson:"_id"`
	DeviceSn                  string        `bson:"deviceSn"`
	StartTime                 time.Time     `bson:"startTime"`
	GmtCreated                time.Time     `bson:"gmtCreated"`
	GmtUpdated                time.Time     `bson:"gmtUpdated"`
	BeforeOptimizationConsume float64       `bson:"beforeOptimizationConsume"`
	BeforeOptimizationMoney   float64       `bson:"beforeOptimizationMoney"`
	Consume                   float64       `bson:"consume"`
	ForecastSavingsConsume    float64       `bson:"forecastSavingsConsume"`
	Load                      float64       `bson:"load"`
	Money                     float64       `bson:"money"`
	TotalEnergy               float64       `bson:"totalEnergy"`
}

func (this EleStat) setValue(eleMeterEnergyStatDO EleStat) EleStat {
	var result EleStat
	result.TotalEnergy = eleMeterEnergyStatDO.TotalEnergy
	result.Consume = this.Consume + eleMeterEnergyStatDO.Consume
	result.BeforeOptimizationConsume = this.BeforeOptimizationConsume + eleMeterEnergyStatDO.BeforeOptimizationConsume
	result.BeforeOptimizationMoney = this.BeforeOptimizationMoney + eleMeterEnergyStatDO.BeforeOptimizationMoney
	result.ForecastSavingsConsume = this.ForecastSavingsConsume + eleMeterEnergyStatDO.ForecastSavingsConsume
	result.Money = this.Money + eleMeterEnergyStatDO.Money
	result.Load = (this.Load + eleMeterEnergyStatDO.Load) / 2
	result.StartTime = eleMeterEnergyStatDO.StartTime
	result.GmtCreated = this.GmtCreated
	result.GmtUpdated = this.GmtUpdated
	result.DeviceSn = this.DeviceSn
	result.Id = this.Id
	return result
}

func dayConnection() {
	collection := getConnection("db_yantai", "t_ele_meter_energy_stat")
	var eleDeviceSn  = []string{"000201004", "000201005", "000201002", "000201007", "000201006", "000201001", "000202012", "000203001", "000203075", "000203010", "000201002", "000203084"}
	for x := 0; x < len(eleDeviceSn); x++ {
		funcName(collection, eleDeviceSn[x])
	}
}

func funcName(collection *mgo.Collection, deviceSn string) {
	var result []EleStat
	collection.Find(bson.M{"deviceSn":deviceSn}).Sort("startTime").All(&result)
	fmt.Println(len(result))
	nowData := make([]EleStat, len(result)/10)
	//当天数据累加为一条数据
	for i := 0; i < len(result); i++ {
		eleStat := result[i]
		if i == 0 {
			nowData[0] = eleStat
			continue
		}

		lastEle, index := getLastData(nowData, eleStat.DeviceSn)
		if index == -1 {
			fmt.Println("error ......")
			break
		}
		nowDay := eleStat.StartTime.Day()
		oldDay := lastEle.StartTime.Day()
		if nowDay == oldDay {
			value := lastEle.setValue(eleStat)
			nowData[index] = value
		} else {
			nowData[index+1] = eleStat
		}
		fmt.Println((i + 1) / (len(result) - 1))
	}

	dayCollection := getConnection("db_yantai", "t_day_ele_meter_energy_stat")
	dayCollection.RemoveAll(bson.M{"deviceSn":deviceSn})
	for j := 0; j < len(nowData); j++ {

		err := dayCollection.Insert(nowData[j])
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	}
}

func getLastData(eleStatList []EleStat, deviceSn string) (EleStat ,int){
	for i := len(eleStatList)-1;i>=0 ;i--  {
		if eleStatList[i].DeviceSn != "" {
			return eleStatList[i], i
		}
	}
	return eleStatList[0],-1
}

func recentConnection() {

	var devStat []DevStat
	//var eleStat []EleStat
	nowData := time.Now()
	nowData.Month().String()
	year, month, day := nowData.Date()
	month = time.Month(month - 1)
	date := time.Date(year, month, day, 0, 0, 0, 0, nowData.Location())

	//collection1 := getConnection("db_yantai", "t_ele_meter_energy_stat")
	////查询一个月内的数据
	//collection1.Find(bson.M{"startTime": bson.M{"$gt": date}}).All(&eleStat)
	//fmt.Println(len(eleStat))
	//
	//recent1Collection := getConnection("db_yantai", "t_recent_ele_meter_energy_stat")
	//
	//for _, v := range eleStat {
	//	err := recent1Collection.Insert(v)
	//	if err != nil {
	//		fmt.Printf(err.Error())
	//	}
	//}

	collection := getConnection("db_yantai", "t_dev_meter_energy_stat")

	//查询一个月内的数据
	collection.Find(bson.M{"startTime": bson.M{"$gt": date}}).All(&devStat)
	fmt.Println(len(devStat))

	recentCollection := getConnection("db_yantai", "t_recent_dev_meter_energy_stat")

	recentCollection.RemoveAll(nil)
	for _, v := range devStat {
		err := recentCollection.Insert(v)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

}

func getConnection(db_name string, collection_name string) *mgo.Collection {
	session, err := mgo.Dial(dataUrl)
	if err != nil {
		println(err)
	}
	collection := session.DB(db_name).C(collection_name)
	return collection
}
