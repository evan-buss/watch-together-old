package transcode

import (
	"fmt"
	"io"
	"log"
	"net"
)

// Progress implements a basic TCP server
// that the ffmpeg transcoder sends progress updates to
func main() {

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			var buf [512]byte
			for {
				n, err := conn.Read(buf[:])
				if err != nil {
					if err == io.EOF {
						log.Println("EOF")
						conn.Close()
						return
					}
					log.Println(err)
				}
				fmt.Println(string(buf[:n]))
			}
		}(conn)
	}
}
