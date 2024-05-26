package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices/graphql-plc-api/data"
	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices/graphql-plc-api/graph/model"
	"github.com/influxdata/influxdb/client/v2"
)

var (
	dbHost = "localhost"
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

func (db *Influxdb) GetPlcs(machine string, time string, limit int, filter *model.IdentifierFilterInput) ([]*model.Plc, error) {
	log.Println("GetPlcs")
	stmt := ""
	if filter == nil || filter.Identifier == nil {
		// log.Println("filter is nil")
		stmt = fmt.Sprintf("SELECT time, maschine, bezeichner, value from data WHERE maschine = '%s' AND time >= %s ORDER BY time ASC LIMIT %d", machine, time, limit)
	} else if len(filter.Identifier.In) > 0 {
		// log.Println("filter.Identifier.In is filled")
		// create empty list
		var identifiers_in []string
		// for all elements in the filter.Identifier.In array create a query
		for _, id := range filter.Identifier.In {
			identifier_in := fmt.Sprintf("bezeichner = '%s'", *id)
			identifiers_in = append(identifiers_in, identifier_in)
		}
		identifers_in_join := strings.Join(identifiers_in, " OR ")
		stmt = fmt.Sprintf("SELECT time, maschine, bezeichner, value from data WHERE maschine = '%s' AND time >= %s AND (%s) ORDER BY time ASC LIMIT %d", machine, time, identifers_in_join, limit)
	}
	// log.Println(stmt)
	// execute query
	response, err := db.executeQuery(stmt)
	if err != nil {
		return nil, err
	}
	// get values
	values := response.Results[0].Series[0].Values
	// log.Println(values)
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
