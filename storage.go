package main

import (
	"encoding/binary"
	"io"
	"os"
)

func check(err error){
	if err != nil {
		panic(err)
	}
}

func Produce(topic string, msg string) uint {		
	filename :=  topic+".log"
	file, err := os.OpenFile(filename,os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	check(err)
	defer file.Close()

	var offset int64 
	offset,err = file.Seek(0, io.SeekEnd)
	check(err)

	err = binary.Write(file, binary.LittleEndian, uint32(len(msg)))
	check(err)

	_,err = file.Write([]byte(msg))
	check(err)
	file.Sync()

	return uint(offset)
}


func Consume(topic string, offset uint) (uint,string) {
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

	return uint(msgLen),string(buf)
}
