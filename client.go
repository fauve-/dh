package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func dialOut(port string) net.Conn {
	c, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatalf("Can't connect to server %s\n", err.Error())
	}
	return c
}

func clientMain() {
	c := dialOut(SERV_PORT)
	basekey := exchangeKeys(ALICEROOT, c)
	keystream := initRC4(basekey)
	bio := bufio.NewReader(os.Stdin)
	//we keep these guys in lockstep so that the traversals of
	//they keystream don't drop out of sync
	for {
		line, err := bio.ReadBytes('\n')
		if err != nil {
			log.Fatal("error reading line: %s\n", err.Error())
		}
		for i := 0; i < len(line); i++ {
			line[i] = keystream.MutateByte(line[i])
		}
		_, err = c.Write(line)
		if err != nil {
			log.Fatalf("error in writing to socket %s\n", err.Error())
		}
		//now we read back the echo
		in := make([]byte, len(line))
		_, err = c.Read(in)
		if err != nil {
			log.Fatalf("error in reading from server %s\n", err.Error())
		}
		for i := 0; i < len(in); i++ {
			in[i] = keystream.MutateByte(in[i])
		}
		fmt.Printf("the server says > %s\n", string(in))

	}
}
