package main

import (
	"fmt"
	"time"
	"net"
	"encoding/gob"
)

type Proceso struct {
	Id int64
	count int64
}


var enviar = false
var p Proceso

func inicioCliente() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(enviar)
	if err != nil {
		fmt.Println(err)
	} 

	if !enviar {
		err = gob.NewDecoder(c).Decode(&p)
		if err != nil {
			fmt.Println(err)
			return
		}
		enviar = true
	}

	c.Close()
	
}

func finCliente()  {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(enviar)
	if err != nil {
		fmt.Println(err)
	}

	err = gob.NewEncoder(c).Encode(p)
	if err != nil {
		fmt.Println(err)
	}
}

func correrProceso()  {
	for {
		fmt.Println("::::::::::::::::::")
		fmt.Println(p.Id, " - ", p.count)
		p.count = p.count + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go inicioCliente()
	go correrProceso()

	var input string
	fmt.Scanln(&input)

	finCliente()

}