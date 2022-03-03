package controller

import (
	"edgeDatahub/mqtt/logic"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stevenyao001/edgeCommon/logger"
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
)

type DataSourceC struct {
}

func (d *DataSourceC) RawDataReport(_ mqtt.Client, msg mqtt.Message) {
	msgEntity := mqtt2.MsgPool.Get().(mqtt2.Msg)
	defer mqtt2.MsgPool.Put(msgEntity)

	err := json.Unmarshal(msg.Payload(), &msgEntity)
	if err != nil {
		logger.ErrorLog("DataSourceC-RawDataReport", "消息解析失败", "", err)
		return
	}
	if msgEntity.DeviceId == "" {
		logger.ErrorLog("DataSourceC-RawDataReport", "设备id不能为空", "", err)
		return
	}

	go logic.CollectorInsM.Get(msgEntity.DeviceId).MsgPutQueue(msgEntity)
}
