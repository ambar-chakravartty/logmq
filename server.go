package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)


func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := reader.ReadString('\n')
	check(err)

	// tokenize request
	tokens := strings.Split(req," ")

	var resp string
	if tokens[0] == "PRODUCE" {
		// We assume everything goes well if we return.
		off := Produce(tokens[1], tokens[2])
		resp = fmt.Sprintf("OK %d\n",off)
	}

	if tokens[0] == "CONSUME" {
		offset,e := strconv.Atoi(strings.TrimSuffix(tokens[2], "\n"))
		check(e)
		l,msg := Consume(tokens[1],uint(offset))
		resp = fmt.Sprintf("MESSAGE %d %d %s\n",offset,l,msg)

	}

	_,err = conn.Write([]byte(resp))
	check(err)
}
