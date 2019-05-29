package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	s := &Config{
		BroadcastPort: 21000,
		//  Simple Service Discovery Protocol address
		MulticastAddress: "239.255.255.250",
		Delay:            1 * time.Second,
	}

	listAddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		fmt.Println("Unable to get listen address")
		return
	}

	address := net.JoinHostPort(s.MulticastAddress, "21000")
	mcAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Cannot resolve address", err)
		return
	}

	// get network interfaces
	/*ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Failed", err)
		return
	}*/

	conn, err := net.ListenMulticastUDP("udp", nil, mcAddr)
	if err != nil {
		fmt.Println("Cannot listen MC")
		return
	}
	defer conn.Close()

	lconn, err := net.ListenUDP("udp", listAddr)
	if err != nil {
		fmt.Println("Cannot listen")
		return
	}
	defer lconn.Close()

	go listen(conn)

	time.Sleep(10 * time.Second)

	for {
		b := []byte("Hello")
		_, err = lconn.WriteToUDP(b, mcAddr)

		time.Sleep(10 * time.Second)
	}
}

func listen(conn *net.UDPConn) {
	for {
		b := make([]byte, 256)
		_, _, err := conn.ReadFromUDP(b)
		if err != nil {
			panic(err)
		}
		fmt.Println("read", string(b))
	}
}
