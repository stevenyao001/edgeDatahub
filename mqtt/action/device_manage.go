package action

import (
	"edgeDatahub/mqtt/logic"
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
)

type DeviceManageA struct {
}

func (d *DeviceManageA) CreateDevice(msg mqtt2.Msg) {

	if msg.Cmd == mqtt2.CollectDeviceRegister {
		logic.CollectorInsM.New(msg.DeviceId)
	}

	//todo do some other things
}
