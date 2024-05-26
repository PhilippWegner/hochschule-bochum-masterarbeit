package model

type RequestPayload struct {
	Action    string          `json:"action"`
	Machine   string          `json:"machine,omitempty"`
	LastState StatePayload    `json:"lastState,omitempty"`
	Plcs      []*PlcPayload   `json:"plcs,omitempty"`
	States    []*StatePayload `json:"states,omitempty"`
	Log       LogPayload      `json:"log,omitempty"`
}

type StatePayload struct {
	Time    string `json:"time"`
	Machine string `json:"machine"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Value   int64  `json:"value"`
}

type PlcPayload struct {
	Time       string               `json:"time"`
	Machine    string               `json:"machine"`
	Identifier []*IdentifierPayload `json:"identifier"`
}

type IdentifierPayload struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
