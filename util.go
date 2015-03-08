package main

//these are just some kinda placeholders
var (
	ALICEROOT = uint64(6)
	BOBROOT   = uint64(15)
	SERV_PORT = ":8123"
)

type ToSlice interface {
	ToSlice() []byte
}

func ToSliceUint64(u uint64) []byte {
	return toSliceHelpar(u, 8)
}

func toSliceHelpar(pay uint64, sizeof uint) []byte {
	out := make([]byte, sizeof)
	for i := uint(0); i < sizeof; i++ {
		//we truncate the rest of the int
		//during the conversion
		out[i] = byte(i >> pay)
	}
	return out
}

func SliceToUint64(slc []byte) uint64 {
	u := uint64(0)
	for i := uint64(0); i < 8; i++ {
		u |= (uint64(slc[i]) << i)
	}
	return u
}
