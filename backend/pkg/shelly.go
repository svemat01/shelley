package pkg

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
