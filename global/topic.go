package global

const (
	SubRawDataTopic      string = "$ROOTEDGE/datasource/rawdata/+"
	SubCreateDeviceTopic string = "$ROOTEDGE/thing/model/+"
	SubRealtimeDataTopic string = "$ROOTEDGE/thing/realtimedata/+"

	PubRealtimeDataTopic string = "$ROOTEDGE/thing/realtimedata/%s"
)
