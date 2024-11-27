package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	port := 1234
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: port})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Listening on port %d\n", port)

	br := bufio.NewReader(conn)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		log.Printf("Received: %s", line)
	}
}
