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

	reply, err := dhcpv4.NewReplyFromRequest(m)
	if err != nil {
		log.Printf("NewReplyFromRequest failed: %v", err)
		return
	}

	switch mt := m.MessageType(); mt {
	case dhcpv4.MessageTypeDiscover:
		reply.UpdateOption(dhcpv4.OptMessageType(dhcpv4.MessageTypeOffer))
	case dhcpv4.MessageTypeRequest:
		reply.UpdateOption(dhcpv4.OptMessageType(dhcpv4.MessageTypeAck))
	default:
		log.Printf("Unhandled message type: %v", mt)
		return
	}

	if _, err := conn.WriteTo(reply.ToBytes(), peer); err != nil {
		log.Printf("Cannot reply to client: %v", err)
	}
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
