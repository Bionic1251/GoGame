package util

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

var (
	IpConnections IPConns
)

func LoadFile(file string) (err error) {
	log.Println("Load file")
	in, err := os.Open(file)
	if err != nil {
		return errors.New("error, file could not load")
	}

	defer in.Close()
	lines, ok := Load(in)

	var ip IPConn
	var ips IPConns
	var gate string
	if ok == nil {
		for _, line := range lines {
			values := strings.Split(line, ",")
			log.Println(values)
			for j, value := range values {
				switch j {
				case 0:
					ip.Ip = value
				case 1:
					ip.Port = value
				default:
					log.Println("default")
				}
			}
			//outer for loop
			ips.AddIp(ip)
		}

		//swap, build gate
		for i, ip := range ips.Ips {
			if i == 0 {
				gate = ip.Ip + ":" + ip.Port
			} else if i == 1 {
				ip.Gate = gate
				gate = ip.Ip + ":" + ip.Port
				ips.Ips[i-1].Gate = gate
				IpConnections.AddIp(ips.Ips[i-1])
				IpConnections.AddIp(ip)
			}
		}
	}
	return nil
}

func Load(in io.Reader) (info []string, err error) {
	var lines []string
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
