/*
	Lookup port for given service
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type(udp/tcp etc) service\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	// func LookupPort(network, service string) (port int, err os.Error)
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	fmt.Println("Service port: ", port)
	os.Exit(0)
}
