package main

import (
	"encoding/binary"
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
