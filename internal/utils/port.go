package utils

import (
	"fmt"
	"net"
)

//CheckPort takes a port number and checks if its available.
//If available, will return the port as it is. If not, error will be thrown.
func CheckPort(port int) int {

	address := fmt.Sprintf(":%d", port)
	connection, err := net.Listen("tcp", address)
	if err != nil {
		TerminateWithCustomError(
			fmt.Sprintf("Port %d is already in use. Please free it first!", port))
	} else {
		_ = connection.Close()
	}
	return port
}