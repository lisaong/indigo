package aud

import "testing"

func TestSendCommand(t *testing.T) {
	connection := Connect()
	SendCommand(connection, "Help: Command=\"GetInfo\"")
	Disconnect(connection)
}
