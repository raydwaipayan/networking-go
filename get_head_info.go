/*
	Get the head information after sending
	HEAD / HTTP/1.0\r\n\r\n to the server

	Connection methods:
	func (c *TCPConn) Write(b []byte) (n int, err os.Error)
	func (c *TCPConn) Read(b []byte) (n int, err os.Error)
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	// func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	handleError(err)

	// establish a tcp connection
	// func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	handleError(err)

	// send the tcp packcets
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	handleError(err)

	// read the recieved response
	result, err := ioutil.ReadAll(conn)
	handleError(err)

	fmt.Println("Result: ", string(result[:]))

}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
