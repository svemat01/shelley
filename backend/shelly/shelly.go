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
	case "REST":
		return GetDeviceStateREST(device.Ip, deviceSpec)
	default:
		return pkg.DeviceState{}, InvalidShellyType(device.Type)
	}
}

func SetSwitchState(deviceId string, switchIndex string, state bool) error {
	// TODO check if device type has switch
	device, err := redisDB.GetDevice(deviceId)

	if err != nil {
		return err
	}

	deviceSpec := pkg.DeviceTypes[device.Type]

	switch deviceSpec.Protocol {
	case "RPC":
		return SetSwitchStateRPC(device.Ip, switchIndex, state)
	case "REST":
		return SetSwitchStateREST(device.Ip, switchIndex, state)
	default:
		return InvalidShellyType(device.Type)
	}
}

func SetLightState(deviceId string, lightIndex string, state bool, brightness int) error {
	// TODO check if device type has light
	device, err := redisDB.GetDevice(deviceId)

	if err != nil {
		return err
	}

	deviceSpec := pkg.DeviceTypes[device.Type]

	switch deviceSpec.Protocol {
	case "RPC":
		return SetLightStateRPC(device.Ip, lightIndex, state, brightness)
	case "REST":
		return SetLightStateREST(device.Ip, lightIndex, state, brightness)
	default:
		return InvalidShellyType(device.Type)
	}
}
