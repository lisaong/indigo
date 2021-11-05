// Applies audio effects using Audacity
// Requires Audacity installed with mod-script-pipe enabled
// See docker/README.md for build instructions
package aud

import (
	"fmt"
	"os"
)

// Loads a file into Audacity
func Load(path string) {
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}
	fmt.Println(fileinfo)
}
