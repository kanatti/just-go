package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	Port string
}

func (c *Client) Run() {
	client_start_logs(c)
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%s", c.Port))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	for {
		fmt.Print(">>")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))

		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		fmt.Printf("Recieved: %q\n", buf[:n])
	}
}

type Server struct {
	Port string
}

func (s *Server) Run() {
	server_start_logs(s)
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", s.Port))
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Printf("Bound to %q\n", listener.Addr())

	for {
		conn, err := listener.Accept() // Blocks till incoming conn and completes TCP handshake
		if err != nil {
			fmt.Println("Accept returned error")
			continue
		}

		fmt.Println("Got connection from " + conn.RemoteAddr().String())

		go func(c net.Conn) {
			defer c.Close() // Terminates by sending FIN packet

			buf := make([]byte, 1024)
			for {
				n, err := c.Read(buf)
				if err != nil {
					if err != io.EOF {
						panic(err)
					}
					break
				}

				c.Write([]byte("Ok!"))
				fmt.Printf("Recieved: %q\n", buf[:n])
			}
			fmt.Println("Closing connection")
		}(conn)
	}
}
