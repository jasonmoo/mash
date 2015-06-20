package mash

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"os"
	"testing"
	// go test should output images or even ascii patterns
)

// func hashfunc(data []byte) uint64 {
// 	const full_width_prime uint64 = 0xFFFFFFFFFFFFFFCF
// 	var h uint64 = full_width_prime
// 	for i, b := range data {
// 		var ib = uint64(i) * (uint64(b) ^ full_width_prime)
// 		h ^= uint64(b)*
// 			(full_width_prime^(ib<<16))*(full_width_prime^(ib<<40)) ^
// 			(full_width_prime^(ib<<32))*(full_width_prime^(ib<<8)) ^
// 			(full_width_prime^(ib<<48))*(full_width_prime^(ib<<56)) ^
// 			(full_width_prime^(ib<<24))*ib
// 	}
// 	return h
// }

func hashfunc(data []byte) uint64 {
	const mask uint64 = 0xFFFFFFFFFFFFFFFF
	var h uint64 = mask
	for i, b := range data {
		h ^= uint64(b)*(mask^(h<<8))*(mask^(h<<40)) ^
			(mask^(h<<16))*(mask^(h<<8)) ^
			(mask^(h<<24))*(mask^(h<<56)) ^
			(mask^(h<<32))*(uint64(i)^mask^uint64(b))
	}
	return h
}

func TestNewHash(t *testing.T) {

	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer file.Close()

	hashes := map[uint64]string{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		h := hashfunc(sc.Bytes())
		if v, exists := hashes[h]; exists {
			fmt.Printf("Collision: %x, %q, %q\n", h, v, sc.Text())
		} else {
			hashes[h] = sc.Text()
			// fmt.Printf("%x\n", h)
		}
	}
	if sc.Err() != nil {
		t.Error(sc.Err())
	}

	inthashes := map[uint64]int{}

	for i := 0; i < 1<<20; i++ {
		var tmp [8]byte
		binary.BigEndian.PutUint64(tmp[:], uint64(i))
		h := hashfunc(tmp[:])
		if v, exists := inthashes[h]; exists {
			fmt.Printf("Collision: %x, %v, %v\n", h, v, i)
		} else {
			inthashes[h] = i
			// fmt.Printf("%x\n", h)
		}
	}

	randhashes := map[uint64][]byte{}

	for i := 0; i < 1<<20; i++ {
		var tmp [8]byte
		rand.Read(tmp[:])
		h := hashfunc(tmp[:])
		if v, exists := randhashes[h]; exists {
			fmt.Printf("Collision: %x, %v, %v\n", h, v, i)
		} else {
			randhashes[h] = tmp[:]
			// fmt.Printf("%x\n", h)
		}
	}

}

func TestNewHashDistribution(t *testing.T) {

	dist := map[byte]int{}

	for i := 0; i < 10<<20; i++ {
		sum := hashfunc([]byte(string(i)))
		var tmp [8]byte
		binary.BigEndian.PutUint64(tmp[:], sum)
		dist[tmp[0]]++
		dist[tmp[1]]++
		dist[tmp[2]]++
		dist[tmp[3]]++
		dist[tmp[4]]++
		dist[tmp[5]]++
		dist[tmp[6]]++
		dist[tmp[7]]++
	}

	for c, ct := range dist {
		fmt.Printf("%v: %d\n", c, ct)
	}

}

func BenchmarkNewHashFile(b *testing.B) {

	wordsdata, _ := ioutil.ReadFile("/usr/share/dict/words")
	words := bytes.Split(wordsdata, []byte{'\n'})

	b.ResetTimer()

	for _, word := range words {
		hashfunc(word)
	}

	b.N = len(words)

}

func BenchmarkNewHashIntegers(b *testing.B) {

	var (
		tmps [][]byte
	)

	for i := 0; i < 235887; i++ {
		var tmp [8]byte
		binary.BigEndian.PutUint64(tmp[:], uint64(i))
		tmps = append(tmps, tmp[:])
	}

	b.ResetTimer()

	for _, tmp := range tmps {
		hashfunc(tmp)
	}

	b.N = len(tmps)

}

func BenchmarkMd5File(b *testing.B) {

	wordsdata, _ := ioutil.ReadFile("/usr/share/dict/words")
	words := bytes.Split(wordsdata, []byte{'\n'})

	b.ResetTimer()

	for _, word := range words {
		md5.Sum(word)
	}

	b.N = len(words)

}

func BenchmarkFnv1aFile(b *testing.B) {

	wordsdata, _ := ioutil.ReadFile("/usr/share/dict/words")
	words := bytes.Split(wordsdata, []byte{'\n'})
	h := fnv.New64a()

	b.ResetTimer()

	for _, word := range words {
		h.Write(word)
		h.Sum64()
	}

	b.N = len(words)

}
