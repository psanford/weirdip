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

func weirdIPv4(ipIn net.IP) {
	ip := NewIPv4(ipIn)

	formats := []string{
		ip.String(),
		strconv.Itoa(int(ip.Uint32())),
		ip.Octal(),
		ip.Hex(),
		ip.ClassA(),
		ip.ClassB(),
		ip.ClassC(),
		ip.V6plusDoted(),
		ip.V6(),
	}

	for _, format := range formats {
		fmt.Println(format)
	}
}

func weirdIPv6(ip net.IP) {
	fmt.Printf("%s\n", ip)
}

type IPv4 net.IP

func NewIPv4(ip net.IP) IPv4 {
	return IPv4(ip.To4())
}

func (ip IPv4) String() string {
	return net.IP(ip).String()
}

func (ip IPv4) Octal() string {
	octals := make([]string, 4)
	for i := 0; i < 4; i++ {
		octals[i] = "0" + strconv.FormatUint(uint64(ip[i]), 8)
	}

	return strings.Join(octals, ".")
}

func (ip IPv4) Hex() string {
	hexes := make([]string, 4)
	for i := 0; i < 4; i++ {
		hexes[i] = "0x" + strconv.FormatUint(uint64(ip[i]), 16)
	}

	return strings.Join(hexes, ".")
}

func (ip IPv4) V6plusDoted() string {
	return fmt.Sprintf("::ffff:%s", ip)
}

func (ip IPv4) V6() string {
	return fmt.Sprintf("::ffff:%02x%02x:%02x%02x", ip[0], ip[1], ip[2], ip[3])
}

func (ip IPv4) Uint32() uint32 {
	return binary.BigEndian.Uint32(ip)
}

func (ip IPv4) ClassA() string {
	return ip.classN(1)
}

func (ip IPv4) ClassB() string {
	return ip.classN(2)
}

func (ip IPv4) ClassC() string {
	return ip.classN(3)
}

func (ip IPv4) classN(n int) string {
	parts := make([]string, 0, 4)
	uparts := make([]byte, 4)
	copy(uparts, ip)
	for i := 0; i < n; i++ {
		parts = append(parts, strconv.Itoa(int(ip[i])))
		uparts[i] = 0
	}

	parts = append(parts, strconv.FormatUint(uint64(binary.BigEndian.Uint32(uparts)), 10))
	return strings.Join(parts, ".")
}
