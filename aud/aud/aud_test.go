package aud

import "testing"

func TestConnectDisconnect(t *testing.T) {
	toFile, fromFile := Connect()
	if toFile == nil || fromFile == nil {
		t.Fatalf("Could not connect")
	}
	Disconnect(toFile, fromFile)
}
