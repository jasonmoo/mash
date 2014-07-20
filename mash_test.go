package mash

import (
	// "math"
	"testing"
	// go test should output images or even ascii patterns
)

func TestUint32(t *testing.T) {

	size := uint32(5)

	hits := make([]uint32, size)

	for i := uint32(97); i < 127; i++ {
		hits[Fnv1a32(Uint32(0xa, i))%size]++
	}

	for i, ct := range hits {
		t.Logf("n: %d, ct: %d\n", i, ct)
	}

	// for c := uint32(0); c < 255; c++ {
	// 	for i := uint32(0); i < 255; i++ {
	// 		hits[Uint32(c, i)%size]++
	// 	}
	// 	var mean uint32
	// 	for _, v := range hits {
	// 		mean += v
	// 	}
	// 	mean /= uint32(len(hits))

	// 	var stddev uint32
	// 	for _, v := range hits {
	// 		stddev += uint32(math.Pow(float64(v-mean), 2))
	// 	}
	// 	stddev /= uint32(len(hits))
	// 	t.Logf("n: %d, mean: %d, stddev: %d\n", c, mean, stddev)
	// 	for i, ct := range hits {
	// 		t.Logf("%d %d\n", i, ct)
	// 	}
	// }

	// t.Log("dont")
}

func BenchmarkBytesUint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BytesUint64([]byte{byte(i)})
	}
}

func BenchmarkUint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uint64(uint64(i))
	}
}

func BenchmarkFnv1a64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv1a64(uint64(i))
	}

}

func BenchmarkFnv1aBytes64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv1aBytes64([]byte{byte(i)})
	}

}

func BenchmarkFnv1a32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv1a32(uint32(i))
	}

}

func BenchmarkFnv1aBytes32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv1aBytes32([]byte{byte(i)})
	}

}

func Fnv1a64(n uint64) uint64 {
	return (14695981039346656037 ^ n) * 1099511628211
}

func Fnv1aBytes64(data []byte) uint64 {
	h := uint64(14695981039346656037)
	for _, c := range data {
		h ^= uint64(c)
		h *= 14695981039346656037
	}
	return h
}

func Fnv1a32(n uint32) uint32 {
	return (2166136261 ^ n) * 16777619
}

func Fnv1aBytes32(data []byte) uint32 {
	h := uint32(2166136261)
	for _, c := range data {
		h ^= uint32(c)
		h *= 16777619
	}
	return h
}
