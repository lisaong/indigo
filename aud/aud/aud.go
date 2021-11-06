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

func GetResponse(rd *bufio.Reader) (response string, err error) {
	const END_SEQUENCE = "BatchCommand finished: OK\n"
	var line string

	for err == nil && line != END_SEQUENCE {
		line, err = rd.ReadString('\n')
		if err == nil {
			response += line
		}
	}

	// one last newline after the end sequence
	if line == END_SEQUENCE {
		line, err = rd.ReadString('\n')
		if err == nil {
			response += line
		}
	}
	return
}

// Sends a single command
func SendCommand(conn Connection, command string) {
	fmt.Println("Send: >>> \n" + command)
	conn.toPipe.WriteString(command + "\n")
	conn.toPipe.Sync()

	// Note: default buffer size is 4K
	rd := bufio.NewReader(conn.fromPipe)
	response, err := GetResponse(rd)
	if err == nil {
		fmt.Printf("Rcvd: <<< \n" + response)
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
