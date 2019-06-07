package main

import "fmt"

type IPAddr [4]byte

func (host IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", host[0], host[1], host[2], host[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for _, ip := range hosts {
		p := &ip
		fmt.Println(ip)
		// it seems that even if String() method has a value receiver
		// it still interprets this in the right way
		fmt.Println(p)
	}
}

