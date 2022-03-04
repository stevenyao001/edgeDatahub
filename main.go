package main

import (
	"edgeDatahub/global"
	mqtt2 "edgeDatahub/mqtt"
	"github.com/stevenyao001/edgeCommon"
	"github.com/stevenyao001/edgeCommon/mqtt"
	"github.com/stevenyao001/edgeCommon/tdengine"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	edge := edgeCommon.New()

	filePath, _ := os.Getwd()
	edge.RegisterConfig(filePath+"/conf/local.yaml", &global.Conf)

	edge.RegisterLogger(global.Conf.Log.MainPath)

	initMqtt(edge, global.Conf.Mqtt)

	initTdEngine(edge, global.Conf.Tdengine)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	_ = <-quit

}

func initMqtt(edge edgeCommon.EdgeCommon, conf []global.MqttConf) {
	mqttConf := make([]mqtt.Conf, 0)
	for k := range conf {
		mqttConf = append(mqttConf, mqtt.Conf{
			InsName:  global.Conf.Mqtt[k].InsName,
			ClientId: global.Conf.Mqtt[k].ClientId,
			Username: global.Conf.Mqtt[k].Username,
			Password: global.Conf.Mqtt[k].Password,
			Addr:     global.Conf.Mqtt[k].Addr,
			Port:     global.Conf.Mqtt[k].Port,
		})
	}
	edge.RegisterMqtt(mqttConf, mqtt2.Subscribes)
}

func initTdEngine(edge edgeCommon.EdgeCommon, conf []global.TdengineConf) {
	tdConf := make([]tdengine.Conf, 0)
	for k := range conf {
		tdConf = append(tdConf, tdengine.Conf{
			InsName:      global.Conf.Tdengine[k].InsName,
			Driver:       global.Conf.Tdengine[k].Driver,
			Network:      global.Conf.Tdengine[k].Network,
			Addr:         global.Conf.Tdengine[k].Fqdn,
			Port:         global.Conf.Tdengine[k].Port,
			Username:     global.Conf.Tdengine[k].Username,
			Password:     global.Conf.Tdengine[k].Password,
			Db:           global.Conf.Tdengine[k].DbName,
			MaxIdleConns: global.Conf.Tdengine[k].MaxIdleConns,
			MaxIdleTime:  global.Conf.Tdengine[k].MaxIdleTime,
			MaxLifeTime:  global.Conf.Tdengine[k].MaxLifeTime,
			MaxOpenConns: global.Conf.Tdengine[k].MaxOpenConns,
		})
	}
	edge.RegisterTdEngine(tdConf)
}
