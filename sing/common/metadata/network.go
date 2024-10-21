package metadata

import (
	"net/netip"
)

// Package netip defines an IP address type that's a small value type. Building on that Addr type, the package also defines AddrPort (an IP address and a port) and Prefix (an IP address and a bit length prefix).

func NetWorkFromNetAddr(network string, addr netip.Addr) string {
	if addr == netip.IPv4Unspecified() {
		return network + "4"
	}
	return network
}
