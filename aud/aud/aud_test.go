package aud

import "testing"

func TestConnectDisconnect(t *testing.T) {
	connection := Connect()
	if connection.toPipe == nil || connection.fromPipe == nil {
		t.Fatalf("Could not connect")
	}
	Disconnect(connection)
}
