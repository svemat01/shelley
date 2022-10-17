package shelly

import "github.com/svemat01/shelley/redisDB"

type DeviceSpec struct {
	Protocol    string `json:"protocol"`
	SwitchCount int    `json:"switch_count"`
}

type SwitchState struct {
	Online bool `json:"online"`
	IsOn   bool `json:"ison"`
}

type DeviceState struct {
	Switches []SwitchState `json:"switches"`
}

// map of device types to their respective structs
var DeviceTypes = map[string]DeviceSpec{
	"SHP1PM": DeviceSpec{
		Protocol:    "RPC",
		SwitchCount: 1,
	},
}

func GetDeviceState(deviceId string, device redisDB.Device) (DeviceState, error) {
	deviceSpec := DeviceTypes[device.Type]

	switch deviceSpec.Protocol {
	case "RPC":
		return GetDeviceStateRPC(deviceId, deviceSpec)
	default:
		return DeviceState{}, InvalidShellyType(device.Type)
	}
}
