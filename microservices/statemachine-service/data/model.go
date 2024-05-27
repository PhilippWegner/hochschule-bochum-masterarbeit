package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Plc struct {
	Time       string `json:"time"`
	Machine    string `json:"machine"`
	Identifier []*Identifier
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

type ApiRepository struct {
	restful_api string
}

func NewApiRepository(restful_api string) *ApiRepository {
	return &ApiRepository{restful_api: restful_api}
}

func (r *ApiRepository) GetPlcs(machine string, time string, limit int) ([]*Plc, error) {
	plcsApiURL := fmt.Sprintf("%v/plcs/%s/%s/%d", r.restful_api, machine, time, limit)
	request, err := http.NewRequest("GET", plcsApiURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var plcs []*Plc
	err = json.NewDecoder(response.Body).Decode(&plcs)
	if err != nil {
		return nil, err
	}
	return plcs, nil
}

func (r *ApiRepository) GetStates(machine string, limit int) ([]*State, error) {
	statesApiURL := fmt.Sprintf("%v/states/%s/%d", r.restful_api, machine, limit)
	request, err := http.NewRequest("GET", statesApiURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var states []*State
	err = json.NewDecoder(response.Body).Decode(&states)
	if err != nil {
		return nil, err
	}
	return states, nil
}

func (r *ApiRepository) CreateState(states []*State) error {
	statesApiURL := fmt.Sprintf("%v/states", r.restful_api)
	statesJson, err := json.Marshal(states)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", statesApiURL, bytes.NewBuffer(statesJson))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}
