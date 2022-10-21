package shelly

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/svemat01/shelley/pkg"
	"io"
	"log"
	"net/http"
	"time"
)

type RpcGetRequestArgs struct {
	Ip     string
	Method string
	Args   string
}

func rpcGet(data RpcGetRequestArgs, response *map[string]interface{}) error {
	url := fmt.Sprintf("http://%s/rpc/%s%s", data.Ip, data.Method, data.Args)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return err
	}

	//fmt.Printf("client: got response!\n")
	//fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		//os.Exit(1)
		return err
	}

	return json.Unmarshal(resBody, response)
}

func GetDeviceStateRPC(baseUrl string, spec pkg.DeviceSpec) (pkg.DeviceState, error) {
	// Fetch Switches
	switches := make([]pkg.SwitchState, spec.SwitchCount)
	for i := 0; i < spec.SwitchCount; i++ {
		var response map[string]interface{}
		err := rpcGet(RpcGetRequestArgs{
			Ip:     baseUrl,
			Method: "Switch.GetStatus",
			Args:   fmt.Sprintf("?id=%d", i),
		}, &response)
		if err != nil {
			return pkg.DeviceState{
				Online: false,
			}, nil
		}

		switches[i] = pkg.SwitchState{
			IsOn: response["output"].(bool),
		}
	}

	// Fetch Lights
	lights := make([]pkg.LightState, 0)
	for i := 0; i < spec.LightCount; i++ {
		log.Printf("light %d", i)
		var response map[string]interface{}
		err := rpcGet(RpcGetRequestArgs{
			Ip:     baseUrl,
			Method: "Light.GetStatus",
			Args:   fmt.Sprintf("?id=%d", i),
		}, &response)
		if err != nil {
			return pkg.DeviceState{
				Online: false,
			}, nil
		}

		lights = append(lights, pkg.LightState{
			IsOn:       response["ison"].(bool),
			Brightness: int(response["brightness"].(float64)),
		})
	}

	return pkg.DeviceState{
		Online:   true,
		Switches: switches,
		Lights:   lights,
	}, nil
}

func SetSwitchStateRPC(baseUrl string, switchIndex string, state bool) error {
	var response map[string]interface{}
	err := rpcGet(RpcGetRequestArgs{
		Ip:     baseUrl,
		Method: "Switch.Set",
		Args:   fmt.Sprintf("?id=%s&on=%t", switchIndex, state),
	}, &response)
	if err != nil {
		return err
	}

	//fmt.Printf("client: response: %v\n", response)

	return nil
}

func SetLightStateRPC(baseUrl string, lightIndex string, state bool, brightness int) error {
	var response map[string]interface{}
	args := RpcGetRequestArgs{
		Ip:     baseUrl,
		Method: "Light.Set",
		Args:   fmt.Sprintf("?id=%s&on=%t", lightIndex, state),
	}

	if brightness != 0 {
		args.Args += fmt.Sprintf("&brightness=%d", brightness)
	}

	err := rpcGet(args, &response)

	return err
}
