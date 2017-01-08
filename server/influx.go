package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

type Temperature uint8
type Measurement struct {
	temp float64
	time string
}

type InfluxServ struct {
	Server string
}

const (
	MIN = iota + 1
	MAX
	NOW
)

func buildQuery(t Temperature, d time.Time) string {
	var q string
	switch t {
	case MIN:
		next := d.AddDate(0, 0, 1)
		q = "SELECT bottom(value,1) FROM temperature WHERE time < '" + next.Format("2006-01-02") + "' and time > '" + d.Format("2006-01-02") + "'"
	case MAX:
		next := d.AddDate(0, 0, 1)
		q = "SELECT top(value,1) FROM temperature WHERE time < '" + next.Format("2006-01-02") + "' and time > '" + d.Format("2006-01-02") + "'"
	case NOW:
		q = "SELECT value FROM temperature WHERE time > now() - 1m LIMIT 1"
	}
	return q

}

func (inf InfluxServ) GetTemperature(query string) (Measurement, error) {
	var meas Measurement
	fmt.Printf("Server: %s", inf.Server)
	request, err := http.NewRequest("GET", inf.Server+"/query", nil)
	q := request.URL.Query()
	q.Add("db", "weather")
	q.Add("q", query)
	request.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return meas, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	var infl InfluxJson
	err = json.Unmarshal(body, &infl)
	fmt.Printf(string(body))

	if err != nil {
		fmt.Printf(err.Error())
	}
	var toRet Measurement
	inte := infl.Results[0].Series[0].Values[0]
	toRet.temp = inte[1].(float64)
	toRet.time = timeToHours(inte[0].(string))
	return toRet, nil
}

func timeToHours(s string) string {
	layout := time.RFC3339Nano
	t, err := time.Parse(layout, s)
	if err != nil {
		return ""
	}
	dur := time.Hour * 3
	t = t.Add(dur)
	return t.Format("15:04")
}
