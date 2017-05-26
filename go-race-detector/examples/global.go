package main

import (
	"fmt"
	"log"
	"net"
)

var (
	service map[string]net.Addr
)

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}

func main() {
	service = make(map[string]net.Addr)
	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for i, iface := range interfaces {
		fmt.Println(iface)
		go RegisterService(fmt.Sprintf("iface-%d", i), iface)
	}

}
