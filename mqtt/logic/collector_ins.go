package logic

import (
	"edgeDatahub/rule"
	"encoding/json"
	"fmt"
	"github.com/stevenyao001/edgeCommon/logger"
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
)

type collectorIns struct {
	deviceId string
	msgQueue chan mqtt2.Msg
	close    chan struct{}
}

//消息入队
func (ins *collectorIns) MsgPutQueue(msg mqtt2.Msg) {
	if msg.DeviceId == "" {
		return
	}

	ins.msgQueue <- msg
}

//消息出队
func (ins *collectorIns) msgOutQueue() {
	//退出后删除程序标识
	defer func() {
		if r := recover(); r != nil {
			logger.ErrorLog("CollectorIns-msgOutQueue", "异常退出", ins.deviceId, r)
		}
		go ins.msgOutQueue()
	}()

	for {
		select {
		case msg, ok := <-ins.msgQueue:
			if !ok {
				return
			}

			//input := &rule.Input{
			//	Ts:         time.Now().UnixNano(),
			//	Properties: &rule.Properties{
			//		AddressNames1: 1,
			//		AddressNames2: 2,
			//		AddressNames3: 3,
			//	},
			//}

			buf, _ := json.Marshal(msg.Content)

			data, _ := rule.Computing(buf,"./rule/rule11.txt")
			fmt.Println("data------", string(data))
			_ = json.Unmarshal(data, &msg.Content)

			_, _ = mqtt2.GetClient("rootcloud").Publish("datasource/computingdata/"+msg.DeviceId, msg, 0, false)
		case <-ins.close:
			return
		}
	}
}
