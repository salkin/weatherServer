package server

import (
	"testing"
	"time"
)

func TestGetTemperature(t *testing.T) {
	inf := InfluxServ{}
	inf.Server = "http://127.0.0.1:8086"
}

func TestBuildQuery(t *testing.T) {

	tim := time.Date(2016, time.August, 9, 0, 0, 0, 0, time.UTC)

	q := "SELECT bottom(value,1) FROM temperature WHERE time < '2016-08-10' and time > '2016-08-09'"
	got := buildQuery(MIN, tim)
	if q != got {
		t.Fatalf("Expected: %s, got %s", q, got)
	}

}

func TestTimeParse(t *testing.T) {
	time := "2016-10-09T06:58:13.256652425Z"

	res := timeToHours(time)
	if res != "09:58" {
		t.Fatalf("Expected: %s, got %s", "09:58", res)
	}

}
