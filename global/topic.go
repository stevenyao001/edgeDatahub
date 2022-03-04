package global

const (
	SubRawDataTopic      string = "$ROOTEDGE/datasource/rawdata/+"
	SubCreateDeviceTopic string = "$ROOTEDGE/datasource/cmd/createdevice/+"
	SubRealtimeDataTopic string = "$ROOTEDGE/thing/realtimedata/+"

	PubRealtimeDataTopic string = "$ROOTEDGE/thing/realtimedata/%s"
)
