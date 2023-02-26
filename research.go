package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Test()
	rand.Seed(time.Now().UnixNano())

	ZERO_IV := [8]uint32{0, 0, 0, 0, 0, 0, 0, 0}

	means := MeasureMean(1000, H)
	fmt.Println(means)
	means = MeasureMean(1000, ZERO_IV)
	fmt.Println(means)
	//This iv is constant throught all test, we would like ot rye with random evert time
	means = MeasureMean(1000, generateIV())
	fmt.Println(means)

}

func Sha256_compress_verbose(chunk [16]uint32, iv [8]uint32) ([8]uint32, [64][8]uint32) {
	msgSchedule := createMessageSchedule(chunk)

	a := iv[0]
	b := iv[1]
	c := iv[2]
	d := iv[3]
	e := iv[4]
	f := iv[5]
	g := iv[6]
	h := iv[7]

	var rounds [64][8]uint32

	for t := 0; t < 64; t++ {

		a, b, c, d, e, f, g, h = Sha256_compress_round(a, b, c, d, e, f, g, h, K[t], msgSchedule[t])
		rounds[t][0] = a
		rounds[t][1] = b
		rounds[t][2] = c
		rounds[t][3] = d
		rounds[t][4] = e
		rounds[t][5] = f
		rounds[t][6] = g
		rounds[t][7] = h

	}

	iv[0] = (iv[0] + a)
	iv[1] = (iv[1] + b)
	iv[2] = (iv[2] + c)
	iv[3] = (iv[3] + d)
	iv[4] = (iv[4] + e)
	iv[5] = (iv[5] + f)
	iv[6] = (iv[6] + g)
	iv[7] = (iv[7] + h)

	return iv, rounds

}

func MeasureMean(count int, iv [8]uint32) [64]int {
	var means [64]int
	for i := 0; i < count; i++ {
		roundCounts := MeasurePseudo(generateMsg(), iv)
		for j := 0; j < 64; j++ {
			means[j] += roundCounts[j]
		}

	}
	for i := 0; i < 64; i++ {
		means[i] = means[i] / count
	}
	return means
}

func MeasurePseudo(msg [16]uint32, iv [8]uint32) [64]int {

	_, rounds := Sha256_compress_verbose(msg, iv)
	_, roundsFlip := Sha256_compress_verbose(FlipRandBit(msg), iv)

	var roundCount [64]int
	for i := 0; i < 64; i++ {
		roundCount[i] = countOnes(xorHash(rounds[i], roundsFlip[i]))
	}

	return roundCount

}

func FlipRandBit(msg [16]uint32) [16]uint32 {
	byteChoice := rand.Intn(16)
	bitChoice := rand.Intn(32)
	flip := uint32(1) << uint32(bitChoice)
	msg[byteChoice] = msg[byteChoice] ^ flip
	return msg
}

func TestFlip() {
	testMsg := [16]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	flipMsg := FlipRandBit(testMsg)
	for i := range testMsg {
		if testMsg[i] != flipMsg[i] {
			fmt.Printf("%032b\n%032b\n", testMsg[i], flipMsg[i])
		}

	}
}

func generateMsg() [16]uint32 {
	var msg [16]uint32
	for i := 0; i < 16; i++ {
		msg[i] = rand.Uint32()
	}
	return msg
}

func generateIV() [8]uint32 {
	var msg [8]uint32
	for i := 0; i < 8; i++ {
		msg[i] = rand.Uint32()
	}
	return msg
}

func xorHash(hash1 [8]uint32, hash2 [8]uint32) [8]uint32 {

	var rv [8]uint32
	for i := 0; i < 8; i++ {
		rv[i] = hash1[i] ^ hash2[i]
	}
	return rv
}

func countOnes(xorMsgs [8]uint32) int {
	var count int
	for i := 0; i < 8; i++ {
		sXor := fmt.Sprintf("%032b", xorMsgs[i])

		for _, num := range sXor {
			if num == '1' {
				count = count + 1
			}
		}
	}

	return count
}
