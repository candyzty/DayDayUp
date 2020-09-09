package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DevStat1 struct {
	Id                            bson.ObjectId `bson:"_id"`
	AccumulatedQCooling           float64       `bson:"accumulatedQCooling"`
	AccumulatedQCoolingSource     float64       `bson:"accumulatedQCoolingSource"`
	AccumulatedQHeat              float64       `bson:"accumulatedQHeat"`
	AccumulatedQHeatSource        float64       `bson:"accumulatedQHeatSource"`
	AccumulatedQCoolingMete       float64       `bson:"accumulatedQCoolingMete"`
	AccumulatedQCoolingMeteSource float64       `bson:"accumulatedQCoolingMeteSource"`
	AccumulatedQHeatMete          float64       `bson:"accumulatedQHeatMete"`
	AccumulatedQHeatMeteSource    float64       `bson:"accumulatedQHeatMeteSource"`
	DeviceSn                      string        `bson:"deviceSn"`
	HeatPumpInstantaneousHeat     float64       `bson:"heatPumpInstantaneousHeat"`
	InstantaneousHeat             float64       `bson:"instantaneousHeat"`
	StartTime                     time.Time     `bson:"startTime"`
	GmtCreated                    time.Time     `bson:"gmtCreated"`
	GmtUpdated                    time.Time     `bson:"gmtUpdated"`
}

func main() {

	var devList = make([]DevStat1, 10)

	fmt.Println(devList)

	//nowData := time.Now()
	//nowData.Month().String()
	//year, month, day := nowData.Date()
	//month =time.Month(month-1)
	//date := time.Date(year, month, day, 0, 0, 0, 0, nowData.Location())
	//
	//format := date.Format("2006-01-02 15:04:05")
	//
	//fmt.Print(format)
}
