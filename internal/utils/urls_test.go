package utils

import (
	"testing"
)

func TestHostAliveWithOutSchema(t *testing.T) {
	host := "google.com"

	if _, err := HostAlive(host); err != nil {
		t.Error(err)
	}
}

func TestHostAliveWithWrongSchema(t *testing.T) {
	host := "htp:/google.com"

	if ok, _ := HostAlive(host); ok {
		t.FailNow()
	}
}

func TestHostAliveWithRightSchema(t *testing.T) {
	host := "http://google.com"

	if ok, err := HostAlive(host); err != nil {
		t.Error(err)
	} else if !ok {
		t.FailNow()
	}
}
