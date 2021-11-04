// Applies audio effects using Audacity
// Requires Audacity installed with mod-script-pipe enabled
// cf. https://github.com/audacity/audacity/blob/master/scripts/piped-work/pipe_test.py
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
