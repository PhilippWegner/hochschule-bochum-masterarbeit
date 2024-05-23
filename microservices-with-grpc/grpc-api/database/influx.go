package database

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/grpc-api/data"
	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/grpc-api/model"
	"github.com/influxdata/influxdb/client/v2"
)

var (
	dbHost = "192.168.0.247"
	dbPort = "8086"
	dbUser = "root"
	dbPass = "root"
	dbName = "reich"
)

type Influxdb struct {
	client client.HTTPClient
}

func ConnectInfluxdb() *Influxdb {
	client, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://" + dbHost + ":" + dbPort,
		Username: dbUser,
		Password: dbPass,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	return &Influxdb{client: client}
}

func (db *Influxdb) CreateStates(states []*model.State) error {
	log.Println("CreateStates")
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbName,
		Precision: "ns",
	})
	if err != nil {
		return err
	}
	for _, state := range states {
		tags := map[string]string{
			"machine": state.Machine,
			"state":   state.State,
			"color":   state.Color,
		}
		fields := map[string]interface{}{
			"value": state.Value,
		}
		time_int64, _ := strconv.ParseInt(state.Time, 10, 64)
		pt, err := client.NewPoint("statemachine", tags, fields, time.Unix(0, time_int64))
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}
	err = db.client.Write(bp)
	if err != nil {
		return err
	}
	return nil
}

func (db *Influxdb) GetStates(machine string, limit int) ([]*model.State, error) {
	log.Println("GetStates")
	stmt := fmt.Sprintf("SELECT time, machine, state, color, value from statemachine WHERE machine = '%s' ORDER BY time DESC LIMIT %d", machine, limit)
	// log.Println(stmt)
	response, err := db.executeQuery(stmt)
	if err != nil {
		return nil, err
	}
	if len(response.Results[0].Series) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	values := response.Results[0].Series[0].Values
	var states []*model.State
	for _, value := range values {
		time := value[0].(json.Number).String()
		machine, _ := value[1].(string)
		state_name, _ := value[2].(string)
		color, _ := value[3].(string)
		value_int64, _ := value[4].(json.Number).Int64()
		value := int(value_int64)
		state := model.State{
			Time:    time,
			Machine: machine,
			State:   state_name,
			Color:   color,
			Value:   int32(value),
		}
		states = append(states, &state)
	}
	return states, nil
}

func (db *Influxdb) GetPlcs(machine string, time string, limit int) ([]*model.Plc, error) {
	log.Println("GetPlcs")
	stmt := ""
	stmt = fmt.Sprintf("SELECT time, maschine, bezeichner, value from data WHERE maschine = '%s' AND time >= %s ORDER BY time ASC LIMIT %d", machine, time, limit)
	// execute query
	response, err := db.executeQuery(stmt)
	if err != nil {
		return nil, err
	}
	// get values
	values := response.Results[0].Series[0].Values
	// pivotize
	plcs := data.Pivotize(values)
	return plcs, nil
}

func (db *Influxdb) executeQuery(stmt string) (*client.Response, error) {
	query := client.NewQuery(stmt, dbName, "ns")
	response, err := db.client.Query(query)
	if err != nil {
		return nil, err
	}
	if response.Error() != nil {
		return nil, response.Error()
	}
	return response, nil
}
