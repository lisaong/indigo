// Applies audio effects using Audacity
// Requires Audacity running with mod-script-pipe enabled
// See docker/README.md for build instructions
//
// Reference: http://www.albertoleal.me/posts/golang-pipes.html

package aud

import (
	"bufio"
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

type Connection struct {
	toPipe   *os.File
	fromPipe *os.File
}

func CreateNamedPipe(direction int) (file *os.File) {

	const TO_PIPE_PREFIX = "audacity_script_pipe.to."
	const FROM_PIPE_PREFIX = "audacity_script_pipe.from."

	var (
		flags  int
		prefix string
	)
	if direction == READ {
		flags = os.O_RDONLY
		prefix = FROM_PIPE_PREFIX
	} else {
		flags = os.O_RDWR
		prefix = TO_PIPE_PREFIX
	}

	path := filepath.Join(os.TempDir(), prefix+strconv.Itoa(os.Getuid()))
	file, err := os.OpenFile(path, flags, os.ModeNamedPipe)
	if err != nil {
		fmt.Println(fmt.Errorf("Could not create file: %s %w. Make sure Audacity is running with mod-script-pipe enabled.", path, err))
	} else {
		fmt.Printf("Opened %s\n", file.Name())
	}
	return
}

// Connects to Audacity's scripting interface
func Connect() (conn Connection) {
	conn.toPipe = CreateNamedPipe(READ_WRITE)
	conn.fromPipe = CreateNamedPipe(READ)
	return
}

// Disconnects from Audacity
func Disconnect(conn Connection) {
	close := func(f *os.File) {
		if f != nil {
			f.Close()
			fmt.Printf("Closed %s\n", f.Name())
		}
	}

	close(conn.toPipe)
	close(conn.fromPipe)
}

func ReadLine(rd *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = rd.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// Sends a single command
func SendCommand(conn Connection, command string) {
	fmt.Println("Send: >>> \n" + command)
	conn.toPipe.WriteString(command + "\n")
	conn.toPipe.Sync()

	// Note: default buffer size is 4K
	rd := bufio.NewReader(conn.fromPipe)
	line, e := ReadLine(rd)
	for e == nil {
		fmt.Printf("Rcvd: <<< \n" + line)
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
