package main

import "math/rand"

const (
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bits    = 6
	mask    = 1<<bits - 1
	max     = 63 / bits
)

type randIndexSource struct {
	src    int64
	remain int
}

func (r *randIndexSource) cutDown() {
	r.src >>= bits
	r.remain--
}

func (r *randIndexSource) resetIfNeeded() {
	if r.remain == 0 {
		r.src, r.remain = rand.Int63(), max
	}
}

func (r *randIndexSource) generateIndex() int {
	return int(r.src & mask)
}

func generateRandString(length int) string {
	bytes := make([]byte, length)
	randIndexSource := randIndexSource{src: rand.Int63(), remain: max}
	for i := 0; i < length; {
		randIndexSource.resetIfNeeded()
		randIndex := randIndexSource.generateIndex()
		if randIndex < len(letters) {
			bytes[i] = letters[randIndex]
			i++
		}
		randIndexSource.cutDown()
	}
	return string(bytes)
}
