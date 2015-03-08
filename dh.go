package main

import (
	"io"
	"log"
	"math"
)

//as per wiki
var (
	base  uint64 = 5
	prime uint64 = 23
)

func exchangeKeys(sec uint64, wtr io.ReadWriter) uint64 {
	tosend := uint64(math.Pow(float64(base), float64(sec))) % prime
	_, err := wtr.Write(ToSliceUint64(tosend))
	if err != nil {
		log.Fatalf("Unable to send data becase %s\n", err.Error())
	}
	//read the response from buffer, 8 bytes
	in := make([]byte, 8)
	_, err = wtr.Read(in)
	if err != nil {
		log.Fatalf("Error in reading in other's data %s \n", err.Error())
	}
	foreign := SliceToUint64(in)
	sec = uint64(math.Pow(float64(foreign), float64(base))) % prime
	return sec
}
