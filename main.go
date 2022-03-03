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

	mqttConfs := make([]mqtt.Conf, 0)
	for k := range global.Conf.Mqtt {
		mqttConfs = append(mqttConfs, mqtt.Conf{
			InsName:  global.Conf.Mqtt[k].InsName,
			ClientId: global.Conf.Mqtt[k].ClientId,
			Username: global.Conf.Mqtt[k].Username,
			Password: global.Conf.Mqtt[k].Password,
			Addr:     global.Conf.Mqtt[k].Addr,
			Port:     global.Conf.Mqtt[k].Port,
		})
	}
	edge.RegisterMqtt(mqttConfs, mqtt2.Subscribes)

	tdConfs := make([]tdengine.Conf, 0)
	for k := range global.Conf.Tdengine {
		tdConfs = append(tdConfs, tdengine.Conf{
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
	edge.RegisterTdEngine(tdConfs)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	_ = <-quit

}
