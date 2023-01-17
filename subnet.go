package main

import (
	"fmt"
	"net"
)

func main() {
	var networkAddress string
	var mask string
	var vpcReservation int
	var subnetAddress *net.IPNet
	var usableIPs []string

	fmt.Print("Enter Network Address: ")
	fmt.Scanln(&networkAddress)

	fmt.Print("Enter Mask: ")
	fmt.Scanln(&mask)
}
