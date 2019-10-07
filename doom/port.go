package doom

import (
	"fmt"
	"net"
)

func PortIsOpen(port int) bool {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}

	defer server.Close()

	return true
}
