package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func Produce(topic string, msg string) {		
	filename :=  topic+".log"
	file, err := os.OpenFile(filename,os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	check(err)
	defer file.Close()
	
	err = binary.Write(file, binary.LittleEndian, uint32(len(msg)))
	check(err)

	_,err = file.Write([]byte(msg))
	check(err)
}


func Consume(topic string, offset uint) {
	filename := topic+".log"
	file,err := os.Open(filename)
	check(err)
	defer file.Close()

	file.Seek(int64(offset),0)
	var msgLen uint32
	err = binary.Read(file,binary.LittleEndian,&msgLen)
	check(err)

	var buf []byte = make([]byte,msgLen)
	_,err = file.Read(buf)
	fmt.Println(string(buf))
}
