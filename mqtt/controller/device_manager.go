package controller

import (
	"edgeDatahub/mqtt/action"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stevenyao001/edgeCommon/logger"
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
)

type DeviceManagerC struct {
}

func (d *DeviceManagerC) CreateDevice(_ mqtt.Client, msg mqtt.Message) {
	msgEntity := mqtt2.MsgPool.Get().(mqtt2.Msg)
	defer mqtt2.MsgPool.Put(msgEntity)

	err := json.Unmarshal(msg.Payload(), &msgEntity)
	if err != nil {
		logger.ErrorLog("DeviceManage-CreateDevice", "消息解析失败", "", err)
		return
	}
	if msgEntity.DeviceId == "" {
		logger.ErrorLog("DeviceManage-CreateDevice", "设备id不能为空", "", err)
		return
	}

	new(action.DeviceManageA).CreateDevice(msgEntity)
}
