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

	ip, ipnet, _ := net.ParseCIDR(networkAddress + "/" + mask)
	subnetAddress = ipnet

	fmt.Print("Select the Cloud provider - AWS,GCP or Azure): ")
	var cloudProvider string
	fmt.Scanln(&cloudProvider)

	if cloudProvider == "AWS" || cloudProvider == "GCP" {
		vpcReservation = 5
	} else if cloudProvider == "Azure" {
		vpcReservation = 4
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
