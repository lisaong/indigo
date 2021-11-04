package main

// go mod edit -replace aud/aud=../aud
// go mod tidy

import (
	"aud/aud"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("aud:")
	fmt.Println(("Aud!"))
	aud.Load("path")
}
