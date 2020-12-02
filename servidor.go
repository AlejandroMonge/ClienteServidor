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


var procesos = []Proceso {
	Proceso{Id: int64(0), count: int64(0)},
	Proceso{Id: int64(1), count: int64(0)}, 
	Proceso{Id: int64(2), count: int64(0)},
	Proceso{Id: int64(3), count: int64(0)},
	Proceso{Id: int64(4), count: int64(0)},
}
var recibir bool = false


func correrProcesos() {
	for {
		fmt.Println("::::::::::::::::::")
		for i, p := range procesos {
			fmt.Println(p.Id, " - ", p.count)
			procesos[i].count = procesos[i].count + 1
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	} 

	go correrProcesos()

	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleCliente(c)
	}
}

func handleCliente(c net.Conn) {

	if !recibir {
		err := gob.NewDecoder(c).Decode(&recibir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if recibir {

		var p Proceso
		err := gob.NewDecoder(c).Decode(&p)
		if err != nil {
			fmt.Println(err)
		}
		procesos = append(procesos, p)
		recibir = false

	} else {
		p := procesos[0]
		err := gob.NewEncoder(c).Encode(p)
		if err != nil {
			fmt.Println(err)
		} else {
			if len(procesos) > 0 {
				procesos = procesos[1:len(procesos)]
			}
		}
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}