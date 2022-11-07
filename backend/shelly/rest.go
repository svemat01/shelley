package shelly

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/svemat01/shelley/pkg"
	"io"
	"net/http"
	"time"
)

type RESTGetRequestArgs struct {
	Ip   string
	Path string
	Args string
}

func restGet(data RESTGetRequestArgs, response *map[string]interface{}) error {
	url := fmt.Sprintf("http://%s/%s%s", data.Ip, data.Path, data.Args)

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

func GetDeviceStateREST(baseUrl string, spec pkg.DeviceSpec) (pkg.DeviceState, error) {
	// Fetch Switches
	switches := make([]pkg.SwitchState, spec.SwitchCount)
	for i := 0; i < spec.SwitchCount; i++ {
		var response map[string]interface{}
		err := restGet(RESTGetRequestArgs{
			Ip:   baseUrl,
			Path: "relay",
			Args: fmt.Sprintf("/%d", i),
		}, &response)
		if err != nil {
			return pkg.DeviceState{}, err
		}

		switches[i] = pkg.SwitchState{
			IsOn: response["ison"].(bool),
		}
	}

	// Fetch Lights
	lights := make([]pkg.LightState, spec.LightCount)
	for i := 0; i < spec.LightCount; i++ {
		var response map[string]interface{}
		err := restGet(RESTGetRequestArgs{
			Ip:   baseUrl,
			Path: "light",
			Args: fmt.Sprintf("/%d", i),
		}, &response)
		if err != nil {
			return pkg.DeviceState{}, err
		}

		lights[i] = pkg.LightState{
			IsOn:       response["ison"].(bool),
			Brightness: int(response["brightness"].(float64)),
		}
	}

	return pkg.DeviceState{
		Online:   true,
		Switches: switches,
		Lights:   lights,
	}, nil
}

func SetSwitchStateREST(baseUrl string, switchIndex string, state bool) error {
	var response map[string]interface{}

	stateString := "off"
	if state {
		stateString = "on"
	}
	err := restGet(RESTGetRequestArgs{
		Ip:   baseUrl,
		Path: "relay",
		Args: fmt.Sprintf("/%s?turn=%s", switchIndex, stateString),
	}, &response)

	return err
}

func SetLightStateREST(baseUrl string, lightIndex string, state bool, brightness int) error {
	var response map[string]interface{}

	stateString := "off"
	if state {
		stateString = "on"
	}

	args := RESTGetRequestArgs{
		Ip:   baseUrl,
		Path: "light",
		Args: fmt.Sprintf("/%s?turn=%s", lightIndex, stateString),
	}

	if brightness > 0 {
		args.Args += fmt.Sprintf("&brightness=%d", brightness)
	}

	err := restGet(args, &response)

	return err
}
