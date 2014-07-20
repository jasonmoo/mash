package mash

func Uint32(c, n uint32) uint32 {

	return n * (n * 7351)
	// return n ^ (n << 1)
	// return n ^ (c^n)<<1

	// return n * 0xA54FF53A * (n<<5 ^ n)
}

func Uint64(n uint64) uint64 {
	return n * 0x510e527fade682d1 * (n<<5 ^ n)
}

func BytesUint32(buf []byte) uint32 {
	r := uint32(buf[0])
	for _, b := range buf {
		r *= 0xA54FF53A * (uint32(b)<<5 ^ uint32(b))
	}
	return r
}

func BytesUint64(buf []byte) uint64 {
	r := uint64(buf[0])
	for _, b := range buf {
		r *= 0x510e527fade682d1 * (uint64(b)<<5 ^ uint64(b))
	}
	return r
}
