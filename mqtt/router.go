package mqtt

import (
	"edgeDatahub/mqtt/controller"
	"github.com/stevenyao001/edgeCommon/mqtt"
)

var Subscribes = map[string][]mqtt.SubscribeOpts{
	"rootcloud": {
		//采集数据
		{
			Topic:    "datasource/rawdata/#",
			Qos:      0,
			Callback: new(controller.DataSourceC).RawDataReport,
		},

		//
		{
			Topic:    "deviceManager/createDevice",
			Qos:      0,
			Callback: new(controller.DeviceManagerC).CreateDevice,
		},
	},
}
