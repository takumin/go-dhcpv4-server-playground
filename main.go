package main

import (
	"log"
	"net"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
)

func handler(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	log.Println("conn", conn.LocalAddr())
	log.Println("peer", peer)
	log.Println("dhcpv4", m.Summary())
}

func main() {
	addr, err := net.ResolveUDPAddr("udp4", ":67")
	if err != nil {
		log.Fatal(err)
	}

	srv, err := server4.NewServer("", addr, handler)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()

	log.Println("Starting DHCPv4 Server")
	srv.Serve()
}
