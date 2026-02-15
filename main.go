package main

import (
	"fmt"
	"net"
)


func main(){
	listener, err := net.Listen("tcp", ":6666")
	fmt.Println("Listening on port 6666.....")
	check(err)

	defer listener.Close()

	for {
		conn,err := listener.Accept()
		check(err)
		go handleConnection(conn)

	}

}
