// Convert go struct to typescript object type
// type Device struct {
	// Name string `json:"name" validate:"required"`
	// Ip   string `json:"ip" validate:"required,ip"`
	// Type string `json:"type" validate:"required,shellyType"`
// }
export type DeviceData = {
    name: string;
    ip: string;
    type: string;
}
// Convert go struct to typescript object type
// type SwitchState struct {
// 	IsOn bool `json:"ison"`
// }

// type LightState struct {
// 	IsOn       bool `json:"ison"`
// 	Brightness int  `json:"brightness"`
// }

// type DeviceState struct {
// 	Online   bool          `json:"online"`
// 	Switches []SwitchState `json:"switches"`
// 	Lights   []LightState  `json:"lights"`
// }

export type DeviceState = {
    online: boolean;
    switches: SwitchState[];
    lights: LightState[];
}

export type SwitchState = {
    ison: boolean;
}

export type LightState = {
    ison: boolean;
    brightness: number;
}