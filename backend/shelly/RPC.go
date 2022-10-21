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
	switches := make([]pkg.SwitchState, spec.SwitchCount)
	for i := 0; i < spec.SwitchCount; i++ {
		var response map[string]interface{}
		err := rpcGet(RpcGetRequestArgs{
			Ip:     baseUrl,
			Method: "Switch.GetStatus",
			Args:   fmt.Sprintf("?id=%d", i),
		}, &response)
		if err != nil {
			switches[i] = pkg.SwitchState{
				Online: false,
			}
			continue
		}

		//fmt.Printf("client: response: %v\n", response)

		switches[i] = pkg.SwitchState{
			Online: true,
			IsOn:   response["output"].(bool),
		}
	}

	return pkg.DeviceState{
		Switches: switches,
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
