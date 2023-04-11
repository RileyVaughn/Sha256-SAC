package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type FunctiionName int

const (
	XOR FunctiionName = iota
	Kfunc
	CHOOSE
	MAJOR
	SIGMA_0
	SIGMA_1
)

func main() {
	// Test()
	rand.Seed(time.Now().UnixNano())

	means := MeasureMean(10000, "H", make([]FunctiionName, 0))
	fmt.Println(means)
	Write("H_normal", means)
	means = MeasureMean(10000, "ZERO", make([]FunctiionName, 0))
	fmt.Println(means)
	Write("ZERO_normal", means)
	means = MeasureMean(10000, "Random", make([]FunctiionName, 0))
	fmt.Println(means)
	Write("Random_normal", means)
	fmt.Println()

	means = MeasureMean(10000, "H", []FunctiionName{XOR})
	fmt.Println(means)
	Write("H_XOR", means)
	fmt.Println()
	// means = MeasureMean(2000, "ZERO", []FunctiionName{XOR})
	// fmt.Println(means)
	// means = MeasureMean(2000, "Random", []FunctiionName{XOR})
	// fmt.Println(means)
	// fmt.Println()

	means = MeasureMean(10000, "ZERO", []FunctiionName{XOR, MAJOR, CHOOSE})
	fmt.Println(means)
	Write("ZERO_XOR_MAJ_CH", means)

}

func Sha256Verbose(msg string) (string, [][64][8]uint32) {

	msgBSlice := preprocess(msg)
	hash := H
	var roundsList [][64][8]uint32
	for _, chunk := range msgBSlice {
		var rounds [64][8]uint32
		hash, rounds = Sha256_compress_verbose(chunk, hash, make([]FunctiionName, 0))
		roundsList = append(roundsList, rounds)
	}
	hashString := fmt.Sprintf("%08x%08x%08x%08x%08x%08x%08x%08x", hash[0], hash[1], hash[2], hash[3], hash[4], hash[5], hash[6], hash[7])
	return hashString, roundsList

}

func Sha256_compress_verbose(chunk [16]uint32, iv [8]uint32, remove []FunctiionName) ([8]uint32, [64][8]uint32) {
	msgSchedule := createMessageSchedule(chunk)
	useFunc := FNStoBS(remove)

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

		if useFunc[XOR] {
			a, b, c, d, e, f, g, h = Sha256XOR_compress_round(a, b, c, d, e, f, g, h, K[t], msgSchedule[t], useFunc)
		} else {
			a, b, c, d, e, f, g, h = Sha256_compress_round(a, b, c, d, e, f, g, h, K[t], msgSchedule[t], useFunc)
		}

		rounds[t][0] = a
		rounds[t][1] = b
		rounds[t][2] = c
		rounds[t][3] = d
		rounds[t][4] = e
		rounds[t][5] = f
		rounds[t][6] = g
		rounds[t][7] = h

	}

	if useFunc[XOR] {
		iv[0] = (iv[0] ^ a)
		iv[1] = (iv[1] ^ b)
		iv[2] = (iv[2] ^ c)
		iv[3] = (iv[3] ^ d)
		iv[4] = (iv[4] ^ e)
		iv[5] = (iv[5] ^ f)
		iv[6] = (iv[6] ^ g)
		iv[7] = (iv[7] ^ h)
	} else {
		iv[0] = (iv[0] + a)
		iv[1] = (iv[1] + b)
		iv[2] = (iv[2] + c)
		iv[3] = (iv[3] + d)
		iv[4] = (iv[4] + e)
		iv[5] = (iv[5] + f)
		iv[6] = (iv[6] + g)
		iv[7] = (iv[7] + h)
	}

	return iv, rounds

}

func Sha256_compress_round(a uint32, b uint32, c uint32, d uint32, e uint32, f uint32, g uint32, h uint32, k uint32, msg uint32, useFunc [6]bool) (uint32, uint32, uint32, uint32, uint32, uint32, uint32, uint32) {

	T1 := h + msg
	if useFunc[Kfunc] {
		T1 += k
	}
	if useFunc[CHOOSE] {
		T1 += Ch(e, f, g)
	}
	if useFunc[SIGMA_1] {
		T1 += Σ1(e)
	}

	T2 := uint32(0)
	if useFunc[MAJOR] {
		T2 += Maj(a, b, c)
	}
	if useFunc[SIGMA_0] {
		T2 += Σ0(a)
	}

	h = g
	g = f
	f = e
	e = (d + T1)
	d = c
	c = b
	b = a
	a = (T1 + T2)

	return a, b, c, d, e, f, g, h
}

func Sha256XOR_compress_round(a uint32, b uint32, c uint32, d uint32, e uint32, f uint32, g uint32, h uint32, k uint32, msg uint32, useFunc [6]bool) (uint32, uint32, uint32, uint32, uint32, uint32, uint32, uint32) {

	T1 := h ^ msg
	if useFunc[Kfunc] {
		T1 ^= k
	}
	if useFunc[CHOOSE] {
		T1 ^= Ch(e, f, g)
	}
	if useFunc[SIGMA_1] {
		T1 ^= Σ1(e)
	}

	T2 := uint32(0)
	if useFunc[MAJOR] {
		T2 ^= Maj(a, b, c)
	}
	if useFunc[SIGMA_0] {
		T2 ^= Σ0(a)
	}

	h = g
	g = f
	f = e
	e = (d ^ T1)
	d = c
	c = b
	b = a
	a = (T1 ^ T2)

	return a, b, c, d, e, f, g, h
}

//XOR is opposite the rest, as the rest are default
func FNStoBS(names []FunctiionName) [6]bool {

	var out [6]bool = [6]bool{false, true, true, true, true, true}

	for i := range names {
		switch names[i] {
		case XOR:
			out[XOR] = true
		case Kfunc:
			out[Kfunc] = false
		case CHOOSE:
			out[CHOOSE] = false
		case MAJOR:
			out[MAJOR] = false
		case SIGMA_0:
			out[SIGMA_0] = false
		case SIGMA_1:
			out[SIGMA_1] = false

		}
	}

	return out
}

func MeasureMean(count int, ivType string, names []FunctiionName) [64]int {
	var means [64]int

	for i := 0; i < count; i++ {
		var roundCounts [64]int
		if ivType == "ZERO" {
			roundCounts = MeasurePseudo(generateMsg(), [8]uint32{0, 0, 0, 0, 0, 0, 0, 0}, names)
		} else if ivType == "H" {
			roundCounts = MeasurePseudo(generateMsg(), H, names)
		} else if ivType == "Random" {
			roundCounts = MeasurePseudo(generateMsg(), generateIV(), names)
		} else {
			log.Fatalln("Wrong IV input")
		}

		for j := 0; j < 64; j++ {
			means[j] += roundCounts[j]
		}

	}
	for i := 0; i < 64; i++ {
		means[i] = means[i] / count
	}
	return means
}

func MeasurePseudo(msg [16]uint32, iv [8]uint32, names []FunctiionName) [64]int {

	_, rounds := Sha256_compress_verbose(msg, iv, names)
	_, roundsFlip := Sha256_compress_verbose(FlipRandBit(msg), iv, names)

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

func Write(filename string, data [64]int) {

	f, err := os.Create("./data/" + filename + ".csv")
	if err != nil {
		log.Fatalln(err)
	}

	writer := csv.NewWriter(f)

	var dataString [][]string
	for i := range data {
		var line []string
		line = append(line, strconv.Itoa(i))
		line = append(line, strconv.Itoa(data[i]))
		dataString = append(dataString, line)
	}

	err = writer.WriteAll(dataString)

	if err != nil {
		log.Fatalln(err)
	}

}
