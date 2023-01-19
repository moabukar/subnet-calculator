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

	// if ipnet.Mask.Size < 32 { // check if the mask size is less than 32
	// 	fmt.Println("Invalid mask, /32, /31, and /30 masks are not supported.")
	// 	os.Exit(1)
	// }

	fmt.Print("Select the Cloud provider - AWS,GCP or Azure): ")
	var cloudProvider string
	fmt.Scanln(&cloudProvider)

	if cloudProvider == "AWS" || cloudProvider == "GCP" {
		vpcReservation = 5
	} else if cloudProvider == "Azure" {
		vpcReservation = 4
	}

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		usableIPs = append(usableIPs, ip.String())
	}
	if subnetAddress != nil {
		fmt.Println("Subnet address: ", subnetAddress)
		fmt.Println("Usable IPs: ", usableIPs[vpcReservation:len(usableIPs)])
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
