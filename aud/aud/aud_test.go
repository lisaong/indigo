package aud

import "testing"

func TestConnectDisconnect(t *testing.T) {
	connection := Connect()
	if connection.toPipe == nil || connection.fromPipe == nil {
		t.Fatalf("Could not connect")
	}
	Disconnect(connection)
}

func TestSendCommand(t *testing.T) {
	connection := Connect()
	SendCommand(connection, "Help: Command=\"GetInfo\"")

	Disconnect(connection)
}
