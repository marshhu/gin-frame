package utils

import "testing"

func TestGetIPAddress(t *testing.T) {
	ip := "121.33.145.21"
	address, err := GetIPAddress(ip)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(address)
}
