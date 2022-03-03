package redis

import (
	"github.com/stevenyao001/edgeCommon/logger"
	"github.com/stevenyao001/edgeCommon/redis"
)

type DataRds struct {
}

func (d *DataRds) ins() string {
	return "rootcloud_edge"
}

func (d *DataRds) CollectDataPushQueue(deviceId, msg string) error {
	rds, err := redis.GetRds(d.ins())
	if err != nil {
		logger.ErrorLog("DataRds-CollectDataPushQueue", "获取连接报错", "", err)
		return err
	}

	_, err = rds.LPush("new_collect_data:"+deviceId, msg).Result()
	if err != nil {
		logger.ErrorLog("DataRds-CollectDataPushQueue", "设置报错", "", err)
		return err
	}

	return nil
}

func (d *DataRds) CollectDataOutQueue(deviceId string) ([]string, error) {
	rds, err := redis.GetRds(d.ins())
	if err != nil {
		logger.ErrorLog("DataRds-CollectDataOutQueue", "获取连接报错", "", err)
		return nil, err
	}

	data, err := rds.BRPop(0, "new_collect_data:"+deviceId).Result()
	if err != nil {
		logger.ErrorLog("DataRds-CollectDataOutQueue", "获取数据报错", "", err)
		return nil, err
	}

	return data, nil
}
