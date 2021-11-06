// Applies audio effects using Audacity
// Requires Audacity running with mod-script-pipe enabled
// See docker/README.md for build instructions
//
// Reference: http://www.albertoleal.me/posts/golang-pipes.html

package aud

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	_ = iota
	READ
	READ_WRITE
)

func CreateNamedPipe(mode int) (file *os.File) {

	const TO_PIPE_PREFIX = "audacity_script_pipe.to."
	const FROM_PIPE_PREFIX = "audacity_script_pipe.from."

	var flag int
	var prefix string
	if mode == READ {
		flag = os.O_RDONLY
		prefix = FROM_PIPE_PREFIX
	} else {
		flag = os.O_RDWR
		prefix = TO_PIPE_PREFIX
	}

	path := filepath.Join(os.TempDir(), prefix+strconv.Itoa(os.Getuid()))
	file, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		fmt.Println(fmt.Errorf("Could not create file: %s %w. Make sure Audacity is running with mod-script-pipe enabled.", path, err))
	} else {
		fmt.Printf("Opened %s\n", file.Name())
	}
	return
}

// Connects to Audacity's scripting interface
func Connect() (toPipe *os.File, fromPipe *os.File) {
	toPipe = CreateNamedPipe(READ_WRITE)
	fromPipe = CreateNamedPipe(READ)
	return
}

// Disconnects from Audacity
func Disconnect(files ...*os.File) {
	for _, f := range files {
		if f != nil {
			fmt.Printf("Closed %s\n", f.Name())
			f.Close()
		}
	}
}

// Processes a file
func Process( /*file object, */ path string) {
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}
	fmt.Println(fileinfo)

	// TODO
}
