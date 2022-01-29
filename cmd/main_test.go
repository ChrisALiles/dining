package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

var binName = "./dining"
var log = "./diningout"
var buffer []byte

func TestMain(m *testing.M) {
	// Build the program and execute it to produce the log file.
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintln(os.Stdin, "Cannot build executable", binName, err)
		os.Exit(1)
	}
	build = exec.Command(binName)
	if err := build.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot run program", binName, err)
		os.Exit(1)
	}
	var berr error
	buffer, berr = os.ReadFile(log)
	if berr != nil {
		fmt.Fprintln(os.Stderr, "Cannot read log file", log, berr)
	}
	result := m.Run()
	os.Remove(binName)
	os.Remove(log)
	os.Exit(result)
}
func TestDining(t *testing.T) {
	type test struct {
		name string
		want string
		msg  string
	}
	tests := []test{
		{name: "Phil 0 Eating", want: "Phil 0 has left the room", msg: "Phil 0 apparently did not eat."},
		{name: "Phil 1 Eating", want: "Phil 1 has left the room", msg: "Phil 1 apparently did not eat."},
		{name: "Phil 2 Eating", want: "Phil 2 has left the room", msg: "Phil 2 apparently did not eat."},
		{name: "Phil 3 Eating", want: "Phil 3 has left the room", msg: "Phil 3 apparently did not eat."},
		{name: "Phil 4 Eating", want: "Phil 4 has left the room", msg: "Phil 4 apparently did not eat."},
		{name: "Phil 0 Stopping", want: "Phil 0 is exiting", msg: "Phil 0 did not stop."},
		{name: "Phil 1 Stopping", want: "Phil 1 is exiting", msg: "Phil 1 did not stop."},
		{name: "Phil 2 Stopping", want: "Phil 2 is exiting", msg: "Phil 2 did not stop."},
		{name: "Phil 3 Stopping", want: "Phil 3 is exiting", msg: "Phil 3 did not stop."},
		{name: "Phil 4 Stopping", want: "Phil 4 is exiting", msg: "Phil 4 did not stop."},
		{name: "Room Stopping", want: "Room exiting", msg: "The room did not stop."},
		{name: "Fork 0 Usage", want: "Fork 0 picked up", msg: "Fork 0 was not used."},
		{name: "Fork 1 Usage", want: "Fork 1 picked up", msg: "Fork 1 was not used."},
		{name: "Fork 2 Usage", want: "Fork 2 picked up", msg: "Fork 2 was not used."},
		{name: "Fork 3 Usage", want: "Fork 3 picked up", msg: "Fork 3 was not used."},
		{name: "Fork 4 Usage", want: "Fork 4 picked up", msg: "Fork 4 was not used."},
		{name: "Fork 0 Stopping", want: "Fork 0 exiting", msg: "Fork 0 did not stop."},
		{name: "Fork 1 Stopping", want: "Fork 1 exiting", msg: "Fork 1 did not stop."},
		{name: "Fork 2 Stopping", want: "Fork 2 exiting", msg: "Fork 2 did not stop."},
		{name: "Fork 3 Stopping", want: "Fork 3 exiting", msg: "Fork 3 did not stop."},
		{name: "Fork 4 Stopping", want: "Fork 4 exiting", msg: "Fork 4 did not stop."},
	}
	for _, te := range tests {
		t.Run(te.name, func(t *testing.T) {
			match, err := regexp.Match(te.want, buffer)
			if err != nil {
				fmt.Fprintln(os.Stderr, "regexp error", err)
				os.Exit(1)
			}
			if !match {
				t.Errorf(te.msg)
			}
		})
	}
}
