package server

import (
	"testing"
)

func TestGetTemperature(t *testing.T) {

	s := GetTemperature()
	t.Fatal("Fatal: %s", s)
}
