package main

import (
	"log"
	"net"
)

func readyListener(port string) net.Listener {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot start listener %s\n", err)
	}
	return l
}

func acceptCon(c net.Conn) {
	basekey := exchangeKeys(BOBROOT, c)
	keystream := initRC4(basekey)
	defer c.Close()
	for {
		//this is a bit of a sticky wicket, because we don't
		//preface a message with it's length, so we just gotta
		//keep reading until a demutated byte matches '\n'
		var inbuf []byte
		for {
			b := make([]byte, 1)
			_, err := c.Read(b)
			if err != nil {
				log.Printf("Error reading in byte: %s\n", err.Error())
				return
			}
			b[0] = keystream.MutateByte(b[0])
			inbuf = append(inbuf, b[0])
			if b[0] == '\n' {
				break
			}

		}
		for i := 0; i < len(inbuf); i++ {
			inbuf[i] = keystream.MutateByte(inbuf[i])
		}
		_, err := c.Write(inbuf)
		if err != nil {
			log.Printf("error in writing to conn %s\n", err.Error())
			return
		}
	}
}

func serverMain() {
	l := readyListener(SERV_PORT)
	for {
		con, err := l.Accept()
		if err != nil {
			log.Fatalf("Error in accepting con: %s\n", err.Error())
		}
		go acceptCon(con)
	}
}
