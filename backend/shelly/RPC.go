package shelly

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Args struct {
	A, B int
}

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

	res, err := http.DefaultClient.Do(req)
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

func GetDeviceStateRPC(baseUrl string, spec DeviceSpec) (DeviceState, error) {
	switches := make([]SwitchState, spec.SwitchCount)
	for i := 0; i < spec.SwitchCount; i++ {
		var response map[string]interface{}
		err := rpcGet(RpcGetRequestArgs{
			Ip:     baseUrl,
			Method: "Switch.GetStatus",
			Args:   fmt.Sprintf("?id=%d", i),
		}, &response)
		if err != nil {
			switches[i] = SwitchState{
				Online: false,
			}
			continue
		}

		//fmt.Printf("client: response: %v\n", response)

		switches[i] = SwitchState{
			Online: true,
			IsOn:   response["output"].(bool),
		}
	}

	return DeviceState{
		Switches: switches,
	}, nil
}
