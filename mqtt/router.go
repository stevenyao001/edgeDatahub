package mqtt

import (
	"edgeDatahub/global"
	"edgeDatahub/mqtt/controller"
	"github.com/stevenyao001/edgeCommon/mqtt"
)

var Subscribes = map[string][]mqtt.SubscribeOpts{
	"rootcloud": {
		//原始数据采集并且计算
		{
			Topic:    global.SubRawDataTopic,
			Qos:      0,
			Callback: new(controller.DataSourceC).RawDataReport,
		},

		//创建设备同时创建实例
		{
			Topic:    global.SubCreateDeviceTopic,
			Qos:      0,
			Callback: new(controller.DeviceManagerC).CreateDevice,
		},
	},
}
