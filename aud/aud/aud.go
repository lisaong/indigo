// Applies audio effects using Audacity
// Requires Audacity running with mod-script-pipe enabled
// See docker/README.md for build instructions
//
// Reference: http://www.albertoleal.me/posts/golang-pipes.html

package aud

import (
	"fmt"
	"os"
)

// Connects to Audacity
func Connect() {

}

// Disconnects from Audacity
func Disconnect() {

}

// Processes a file
func Process(path string) {
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}
	fmt.Println(fileinfo)

	// TODO
}
