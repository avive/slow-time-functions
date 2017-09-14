package main

import (
    "github.com/minio/sha256-simd" // simd optimized sha256 computation
	//"crypto/sha256" // use the go crypto lib for comparison
    "fmt"
	"bytes"
	"math"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	t1 := time.Now().Unix()

	buff:= bytes.Buffer{}
	buff.Write([]byte("Seed data goes here"))
	out := [32]byte{}
	n := uint64(math.Pow(10, 8))

	fmt.Printf("Computing %d seriel sha-256s...\n", n)

	for i := uint64(0); i < n; i++ {
		out = sha256.Sum256(buff.Bytes())
		buff.Reset()
		buff.Write(out[:])
	}

	// my 2016 mbp w 2.5 GHz Intel Core i7 does ~33m hashes per second
	fmt.Printf("Final hash: %x took:%d seconds\n", buff.Bytes(), time.Now().Unix() - t1)
}


