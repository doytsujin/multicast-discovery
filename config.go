package main

import (
	"time"
)

type Config struct {
	BroadcastPort int
	// MulticastAddress specifies the multicast address.
	// You should be able to use any of 224.0.0.0/4 or ff00::/8.
	// By default it uses the Simple Service Discovery Protocol
	// address (239.255.255.250 for IPv4 or ff02::c for IPv6).
	MulticastAddress string
	// Dely is the time between broadcasts
	Delay time.Duration
}
