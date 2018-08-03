package main

import (
	"fmt"
	"github.com/vasiliy249/copeer"
	"net"
)

func localAddresses() net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print("error during resolve self ip")
		return nil
	}
	for i, _ := range ifaces {
		fmt.Println(ifaces[i].Name)
	}
	return nil
}

func main() {
	fmt.Println("Enter network interface name (in double-quoted form):")
	fmt.Scanf("%q", &copeer.NetIfaceName)

	fmt.Println("Starting on network interface :", copeer.NetIfaceName)

	cfg := copeer.NewConfig()
	cop := copeer.NewCopeer(cfg)

	for {
		fmt.Println("Enter command (exit/master/bs/publish/start/ping/show):")
		var cmd string
		fmt.Scanln(&cmd)

		if cmd == "exit" {
			cop.Stop()
			return
		} else if cmd == "master" {
			cop.Config.SetMaster()
			continue
		} else if cmd == "bs" {
			cop.Config.SetBs()
			continue
		} else if cmd == "publish" {
			cop.Publish()
			continue
		} else if cmd == "start" {
			go cop.Start()
			continue
		} else if cmd == "ping" {
			fmt.Println("Enter ip address: ")
			var strIP string
			fmt.Scanln(&strIP)
			if check := net.ParseIP(strIP); check == nil {
				fmt.Println("Wrong IP address")
				continue
			}
			fmt.Println("Enter UDP port: ")
			var strPort string
			fmt.Scanln(&strPort)
			cop.PingNode(strIP + ":" + strPort)
		} else if cmd == "show" {
			cop.ShowRouteState()
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
