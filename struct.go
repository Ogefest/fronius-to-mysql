package main

type Response struct {
	Body Body `json:"Body"`
	Head Head `json:"Head"`
}

type Body struct {
	Data Data `json:"Data"`
}

type Data struct {
	DAY_ENERGY   Energy       `json:"DAY_ENERGY"`
	DeviceStatus DeviceStatus `json:"DeviceStatus"`
	FAC          Frequency    `json:"FAC"`
	IAC          Current      `json:"IAC"`
	IDC          Current      `json:"IDC"`
	PAC          Power        `json:"PAC"`
	TOTAL_ENERGY Energy       `json:"TOTAL_ENERGY"`
	UAC          Voltage      `json:"UAC"`
	UDC          Voltage      `json:"UDC"`
	YEAR_ENERGY  Energy       `json:"YEAR_ENERGY"`
}

type Energy struct {
	Unit  string  `json:"Unit"`
	Value float64 `json:"Value"`
}

type DeviceStatus struct {
	ErrorCode              int  `json:"ErrorCode"`
	LEDColor               int  `json:"LEDColor"`
	LEDState               int  `json:"LEDState"`
	MgmtTimerRemainingTime int  `json:"MgmtTimerRemainingTime"`
	StateToReset           bool `json:"StateToReset"`
	StatusCode             int  `json:"StatusCode"`
}

type Frequency struct {
	Unit  string  `json:"Unit"`
	Value float64 `json:"Value"`
}

type Current struct {
	Unit  string  `json:"Unit"`
	Value float64 `json:"Value"`
}

type Power struct {
	Unit  string  `json:"Unit"`
	Value float64 `json:"Value"`
}

type Voltage struct {
	Unit  string  `json:"Unit"`
	Value float64 `json:"Value"`
}

type Head struct {
	RequestArguments RequestArguments `json:"RequestArguments"`
	Status           Status           `json:"Status"`
	Timestamp        string           `json:"Timestamp"`
}

type RequestArguments struct {
	DataCollection string `json:"DataCollection"`
	DeviceClass    string `json:"DeviceClass"`
	DeviceId       string `json:"DeviceId"`
	Scope          string `json:"Scope"`
}

type Status struct {
	Code        int    `json:"Code"`
	Reason      string `json:"Reason"`
	UserMessage string `json:"UserMessage"`
}
