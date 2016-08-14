package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type InfluxJson struct {
	Results []struct {
		Series []struct {
			Columns []string        `json:"columns"`
			Name    string          `json:"name"`
			Values  [][]interface{} `json:"values"`
		} `json:"series"`
	} `json:"results"`
}

func GetTemperature() string {
	request, err := http.NewRequest("GET", "http://192.168.1.252:8086/query?pretty", nil)
	q := request.URL.Query()
	q.Add("db", "weather")
	q.Add("q", "SELECT value FROM temperature WHERE time > now() - 1m  LIMIT 1")
	request.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	var infl InfluxJson
	err = json.Unmarshal(body, &infl)
	fmt.Printf(string(body))

	if err != nil {
		fmt.Printf(err.Error())
	}
	var toRet float64
	inte := infl.Results[0].Series[0].Values[0]
	toRet = inte[1].(float64)
	tStr := strconv.FormatFloat(toRet, 'f', 1, 64)
	return tStr
}
