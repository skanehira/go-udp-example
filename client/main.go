package main

import (
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: client <ip:port>")
		os.Exit(1)
	}

	laddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	raddr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	terminal := term.NewTerminal(os.Stdin, "> ")
	for {
		line, err := terminal.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			log.Printf("send error: %v", err)
		}
	}
}
