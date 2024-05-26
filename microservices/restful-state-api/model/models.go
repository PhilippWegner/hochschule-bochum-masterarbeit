package model

type Plc struct {
	Time       string        `json:"time"`
	Machine    string        `json:"machine"`
	Identifier []*Identifier `json:"identifier"`
}

type Identifier struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type State struct {
	Time    string `json:"time"`
	Machine string `json:"machine"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Value   int    `json:"value"`
}
