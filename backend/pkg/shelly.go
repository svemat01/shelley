package pkg

type DeviceSpec struct {
	Protocol    string `json:"protocol"`
	SwitchCount int    `json:"switch_count"`
	LightCount  int    `json:"light_count"`
}

type SwitchState struct {
	IsOn bool `json:"ison"`
}

type LightState struct {
	IsOn       bool `json:"ison"`
	Brightness int  `json:"brightness"`
}

type DeviceState struct {
	Online   bool          `json:"online"`
	Switches []SwitchState `json:"switches"`
	Lights   []LightState  `json:"lights"`
}

// map of device types to their respective structs
var DeviceTypes = map[string]DeviceSpec{
	"SHP1PM": DeviceSpec{
		Protocol:    "RPC",
		SwitchCount: 1,
	},
}
