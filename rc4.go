package main

import (
	"log"
)

type RC4 []byte

func initRC4(sec uint64) RC4 {
	key := ToSliceUint64(sec)
	b := make([]byte, 256)
	for i := 0; i < 255; i++ {
		b[i] = byte(i)
	}
	j := uint8(0)
	for i := 0; i < 255; i++ {
		j = uint8(int((j + b[i] + key[i%8])) % 256)
		//swap is just so easy
		b[i], b[j] = b[j], b[i]
	}
	return b
}

var (
	i = uint64(0)
	j = uint64(0)
)

//Mutate byte is for mutating bytes against the rc4 stream
func (rc4 RC4) MutateByte(b byte) byte {
	i = (i + 1) % 256
	j = (j + uint64(rc4[i])) % 256
	rc4[i], rc4[j] = rc4[j], rc4[i]
	idx := int(rc4[i]+rc4[j]) % 256
	log.Println(idx)
	val := int(rc4[idx]) % 256

	key := uint8(val)
	return b ^ key
}
