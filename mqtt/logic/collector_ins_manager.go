package logic

import (
	mqtt2 "github.com/stevenyao001/edgeCommon/mqtt"
	"sync"
)

var CollectorInsM *collectorInsManager

type collectorInsManager struct {
	mutex sync.RWMutex
	ins   map[string]*collectorIns
}

func init() {
	CollectorInsM = &collectorInsManager{
		mutex: sync.RWMutex{},
		ins:   make(map[string]*collectorIns),
	}
}

//
func (manager *collectorInsManager) New(deviceId string) *collectorIns {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	if ins, exists := manager.ins[deviceId]; exists {
		return ins
	}

	ins := &collectorIns{
		deviceId: deviceId,
		msgQueue: make(chan mqtt2.Msg, 1000),
		close:    make(chan struct{}, 1),
	}

	go ins.msgOutQueue()

	manager.ins[deviceId] = ins

	return ins
}

//
func (manager *collectorInsManager) Del(deviceId string) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	ins, exists := manager.ins[deviceId]

	if !exists {
		return
	}

	ins.close <- struct{}{}

	delete(manager.ins, deviceId)
}

func (manager *collectorInsManager) Get(deviceId string) *collectorIns {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	return manager.ins[deviceId]
}
