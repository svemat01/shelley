package shelly

import (
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
)

func GetDeviceState(device redisDB.Device) (pkg.DeviceState, error) {
	deviceSpec := pkg.DeviceTypes[device.Type]

	switch deviceSpec.Protocol {
	case "RPC":
		return GetDeviceStateRPC(device.Ip, deviceSpec)
	default:
		return pkg.DeviceState{}, InvalidShellyType(device.Type)
	}
}

func SetSwitchState(deviceId string, switchIndex string, state bool) error {
	device, err := redisDB.GetDevice(deviceId)

	if err != nil {
		return err
	}

	deviceSpec := pkg.DeviceTypes[device.Type]

	switch deviceSpec.Protocol {
	case "RPC":
		return SetSwitchStateRPC(device.Ip, switchIndex, state)
	default:
		return InvalidShellyType(device.Type)
	}
}
