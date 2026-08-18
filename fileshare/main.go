package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

func main() {
	p := flag.String("path", "", "directory to share")
	flag.Parse()
	if *p == "" {
		fmt.Println("path not provided")
		return
	}

	dir := http.Dir(*p)
	handler := http.FileServer(dir)

	port := ":8080"
	fmt.Printf("url: %s%s, directory: %s\n", hostIPv4(), port, *p)
	fmt.Println(http.ListenAndServe(port, handler))
}

func hostIPv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if ip != nil {
			return ip.String()
		}
	}

	return ""
}
