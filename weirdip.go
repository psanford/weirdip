package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// see: https://blog.dave.tf/post/ip-addr-parsing/
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <ip>", os.Args[0])
	}

	origIP := os.Args[1]

	ip := net.ParseIP(origIP)
	if ip == nil {
		log.Fatalf("invalid ip <%s>\n", ip)
	}

	maybeV4 := ip.To4()
	if maybeV4 != nil {
		weirdIPv4(maybeV4)
	} else {
		weirdIPv6(ip)
	}
}

func weirdIPv4(ip net.IP) {
	fmt.Printf("%s\n", ip)

	n := binary.BigEndian.Uint32(ip)
	fmt.Printf("%d\n", n)

	octals := make([]string, 4)
	hexes := make([]string, 4)
	for i := 0; i < 4; i++ {
		octals[i] = "0" + strconv.FormatUint(uint64(ip[i]), 8)
		hexes[i] = "0x" + strconv.FormatUint(uint64(ip[i]), 16)
	}

	fmt.Println(strings.Join(octals, "."))
	fmt.Println(strings.Join(hexes, "."))

	weirdIPv6(ip.To16())
}

func weirdIPv6(ip net.IP) {
	fmt.Printf("%s\n", ip)
}
