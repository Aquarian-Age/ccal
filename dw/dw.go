package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
)

// go run .\fileServer.go -p 9090 -d d:\  显示d:\盘目录下所有文件 含子目录
// go run .\fileServer.go -p 9090 只显示当前目录下的文件

func main() {
	flag.Usage = func() {
		fmt.Println("Usage go run FileServer.exe -p 9999 -d e:\\")
		flag.PrintDefaults()
	}
	p := flag.Int("p", 9999, "")
	d := flag.String("d", ".", "")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	//
	fmt.Printf("server: %s%s\n", ip.String(), port)
	http.ListenAndServe(port, http.FileServer(http.Dir(*d)))
}

func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
