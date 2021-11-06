package main

// go mod edit -replace aud/aud=../aud
// go mod tidy

import (
	"aud/aud"
)

func main() {
	aud.Connect()
	aud.Process("../media/mixkit-tech-house-vibes-130.mp3")
	aud.Disconnect()
}
