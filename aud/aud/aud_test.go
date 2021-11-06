package aud

import "testing"

func TestSendCommand(t *testing.T) {
	connection := Connect()
	if connection.toPipe == nil || connection.fromPipe == nil {
		t.Fatalf("Could not connect")
	}
	SendCommand(connection, "Help: Command=\"GetInfo\"")
	Disconnect(connection)
}
