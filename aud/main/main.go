package main

// go mod edit -replace aud/aud=../aud
// go mod tidy

import (
	"aud/aud"
)

func main() {
	conn := aud.Connect()
	aud.Process(conn, "../media/mixkit-tech-house-vibes-130.mp3")
	aud.Disconnect(conn)
}
