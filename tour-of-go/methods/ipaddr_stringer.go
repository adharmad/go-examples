package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]int

// TODO: Add a "String() string" method to IPAddr.
func (ip *IPAddr) String() string {
	var s string

	for i:=0 ; i<4 ; i++ {
		s += strconv.Itoa(ip[0])
		if i < 3 {
			s += "."
		}
	}

	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.String())
	}

}
