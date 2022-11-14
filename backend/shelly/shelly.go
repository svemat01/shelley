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

func UpdateDeviceState(deviceId string) error {
	device, err := redisDB.GetDevice(deviceId)

	if err != nil {
		return err
	}

	deviceState, err := GetDeviceState(device)

	if err != nil {
		return err
	}

	err = redisDB.SetDeviceState(deviceId, deviceState)

	if err != nil {
		return err
	}

	return nil
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
		err = SetSwitchStateRPC(device.Ip, switchIndex, state)
		if err != nil {
			return err
		}

		err = UpdateDeviceState(deviceId)
		if err != nil {
			return err
		}
	case "REST":
		err = SetSwitchStateREST(device.Ip, switchIndex, state)
		if err != nil {
			return err
		}

		err = UpdateDeviceState(deviceId)
		if err != nil {
			return err
		}
	default:
		return InvalidShellyType(device.Type)
	}

	return nil
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
		err = SetLightStateRPC(device.Ip, lightIndex, state, brightness)
		if err != nil {
			return err
		}

		err = UpdateDeviceState(deviceId)
		if err != nil {
			return err
		}
	case "REST":
		err = SetLightStateREST(device.Ip, lightIndex, state, brightness)
		if err != nil {
			return err
		}

		err = UpdateDeviceState(deviceId)
		if err != nil {
			return err
		}
	default:
		return InvalidShellyType(device.Type)
	}

	return nil
}
