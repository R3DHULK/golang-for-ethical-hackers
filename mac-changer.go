package main

//usage: go run mac-changer.go -i <interface> -m <new-mac>
//example: go run mac-changer.go -i eth0 -m 00:11:22:33:44:55:66

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func executeCommand(command string, argsArray []string) (err error) {
	args := argsArray
	// create object
	cmdObj := exec.Command(command, args...)
	// stdout to display the output on the screen
	cmdObj.Stdout = os.Stdout
	// process errors
	cmdObj.Stderr = os.Stderr
	// stdin to add input commands
	cmdObj.Stdin = os.Stdin

	// run the command
	err = cmdObj.Run()

	if err != nil {
		log.Fatal(err)
		return
	}
	return nil
}

func main() {

	iface := flag.String("i", "eth0", "Interface to change MAC address")
	newmac := flag.String("m", "", "Type new MAC address")
	flag.Parse()

	command := "sudo"
	executeCommand(command, []string{"ifconfig", *iface, "down"})
	executeCommand(command, []string{"ifconfig", *iface, "hw", "ether", *newmac})
	executeCommand(command, []string{"ifconfig", *iface, "up"})
}
