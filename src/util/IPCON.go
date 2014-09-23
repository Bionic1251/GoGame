package util

import (
	"fmt"
	"log"
	"net"
)

/*
* Add gate
* IP1 needs address from IP2
 */
var (
	Gate string
)

type IPConn struct {
	Ip      string
	netmask int
	network string
	Port    string
	Gate    string // ip + port #2
}

type IPConns struct {
	Ips []IPConn
}

func (hosts *IPConns) AddIp(ip IPConn) []IPConn {

	addr := net.ParseIP(ip.Ip)
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, _ := mask.Size()

	ip.network = network.String()
	ip.netmask = ones

	Gate = ip.Ip + ":" + ip.Port

	hosts.Ips = append(hosts.Ips, ip)
	return hosts.Ips
}

func (ip IPConn) String() string {
	return fmt.Sprintf("ip: %s \nNetwork: %s \nPort: %s \nGate %s ", ip.Ip, ip.network, ip.Port, ip.Gate)
}

func IterateList(ips []IPConn) int {
	count := 0
	for i, _ := range ips {
		count = i
	}
	return count
}

func PrintIpDetails(ip string) {
	//uses string array loop
	addrs, err := net.LookupHost(ip)

	addr := net.ParseIP(ip)
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size() //

	log.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())

	//ip lookup details
	log.Printf("%v, %s", addrs, err)
}
