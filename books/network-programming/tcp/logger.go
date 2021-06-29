package tcp

import (
	"fmt"
)

func client_start_logs(c *Client) {
	fmt.Printf("Connecting client to port %s\n", c.Port)
}

func server_start_logs(s *Server) {
	fmt.Printf("Running server on port %s\n", s.Port)
}
