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

type FunctionName int

const (
	XOR FunctionName = iota // Use "+" normal when false
	Kfunc //use when true
	CHOOSE //use when true
	MAJOR //use when true
	SIGMA_0 //use when true
	SIGMA_1 //use when true
	SCHEDULE //use when true
)

func main() {
	//Test()
	rand.Seed(time.Now().UnixNano())

	means := MeasureStrictMean(10000, "H", []FunctionName{})
	WriteFull("H_Normal",means)


	means = MeasureStrictMean(10000, "H", []FunctionName{XOR})
	WriteFull("H_XOR",means)

	means = MeasureStrictMean(10000, "Random", []FunctionName{XOR,Kfunc,CHOOSE,MAJOR,SIGMA_0,SIGMA_1})
	WriteFull("R_ALL-SCHEDULE",means)

	means = MeasureStrictMean(10000, "Random", []FunctionName{XOR,Kfunc,CHOOSE,MAJOR,SIGMA_0,SIGMA_1,SCHEDULE})
	WriteFull("R_ALL",means)

}

/////////////////////////////////////////////////////////////////////////////////
// Hash and compression functions
/////////////////////////////////////////////////////////////////////////////////

func Sha256Verbose(msg string) (string, [][64][8]uint32) {

	msgBSlice := preprocess(msg)
	hash := H
	var roundsList [][64][8]uint32
	for _, chunk := range msgBSlice {
		var rounds [64][8]uint32
		hash, rounds = Sha256_compress_verbose(chunk, hash, make([]FunctionName, 0))
		roundsList = append(roundsList, rounds)
	}
	hashString := fmt.Sprintf("%08x%08x%08x%08x%08x%08x%08x%08x", hash[0], hash[1], hash[2], hash[3], hash[4], hash[5], hash[6], hash[7])
	return hashString, roundsList

}

func Sha256_compress_verbose(chunk [16]uint32, iv [8]uint32, remove []FunctionName) ([8]uint32, [64][8]uint32) {
	
	useFunc := FNStoBS(remove)
	
	var msgSchedule [64]uint32

	if useFunc[SCHEDULE] {
		msgSchedule = createMessageSchedule(chunk)
	} else {
		msgSchedule = normalMessageSchedule(chunk)
	}
	

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

func Sha256_compress_round(a uint32, b uint32, c uint32, d uint32, e uint32, f uint32, g uint32, h uint32, k uint32, msg uint32, useFunc [7]bool) (uint32, uint32, uint32, uint32, uint32, uint32, uint32, uint32) {

	
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

func Sha256XOR_compress_round(a uint32, b uint32, c uint32, d uint32, e uint32, f uint32, g uint32, h uint32, k uint32, msg uint32, useFunc [7]bool) (uint32, uint32, uint32, uint32, uint32, uint32, uint32, uint32) {
	
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


func normalMessageSchedule(chunk [16]uint32) [64]uint32 {

	var msgSchedule [64]uint32
	for j := range chunk {
		msgSchedule[j] = chunk[j]
		msgSchedule[j+16] = chunk[j]
		msgSchedule[j+32] = chunk[j]
		msgSchedule[j+48] = chunk[j]
	}

	return msgSchedule
}



/////////////////////////////////////////////////////////////////////////////////
// Measurements
/////////////////////////////////////////////////////////////////////////////////

// This only measures avalanche, not complete-avalanche
func MeasureMean(count int, ivType string, names []FunctionName) [64]int {
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

//Heuristically measures complete-avalanche
func MeasureStrictMean(count int, ivType string, names []FunctionName) [64][256]float32 {
	var means [64][256]float32

	for i := 0; i < count; i++ {
		var roundCounts [64][256]bool
		if ivType == "ZERO" {
			roundCounts = MeasureStrict(generateMsg(), [8]uint32{0, 0, 0, 0, 0, 0, 0, 0}, names)
		} else if ivType == "H" {
			roundCounts = MeasureStrict(generateMsg(), H, names)
		} else if ivType == "Random" {
			roundCounts = MeasureStrict(generateMsg(), generateIV(), names)
		} else {
			log.Fatalln("Wrong IV input")
		}

		for j := 0; j < 64; j++ {
			for k := 0; k < 256; k++ {
				if roundCounts[j][k] {
					means[j][k] += 1
				}
			}
		}

	}
	for i := 0; i < 64; i++ {
		for j := 0; j < 256; j++ {
			means[i][j] = means[i][j] / float32(count)
		}
	}
	return means

}

// Returns amount of 1's in XOR of msg and msg with 1 random bit modified, for all 64 rounds
func MeasurePseudo(msg [16]uint32, iv [8]uint32, names []FunctionName) [64]int {

	_, rounds := Sha256_compress_verbose(msg, iv, names)
	_, roundsFlip := Sha256_compress_verbose(FlipRandBit(msg), iv, names)

	var roundCount [64]int
	for i := 0; i < 64; i++ {
		roundCount[i] = countOnes(xorHash(rounds[i], roundsFlip[i]))
	}

	return roundCount

}

//Returns XOR of msg and msg with 1 random bit modified, for all 64 rounds
func MeasureStrict(msg [16]uint32, iv [8]uint32, names []FunctionName) [64][256]bool {

	_, rounds := Sha256_compress_verbose(msg, iv, names)
	_, roundsFlip := Sha256_compress_verbose(FlipRandBit(msg), iv, names)

	var roundXOR [64][256]bool
	for i := 0; i < 64; i++ {
		roundXOR[i] = Uint32SToBoolS(xorHash(rounds[i], roundsFlip[i]))
	}

	return roundXOR

}

/////////////////////////////////////////////////////////////////////////////////
// Hash Auxilary
/////////////////////////////////////////////////////////////////////////////////

//XOR is opposite the rest, as the rest are default
func FNStoBS(names []FunctionName) [7]bool {

	var out [7]bool = [7]bool{false, true, true, true, true, true}

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
		case SCHEDULE:
			out[SCHEDULE] = false
		}
	}

	return out
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

func Uint32SToBoolS(uints [8]uint32) [256]bool {
	var rv [256]bool

	for i := uint32(0); i < 8; i++ {
		for j := uint32(0); j < 32; j++ {
			rv[i*32+j] = uints[i]&(uint32(1)<<j) != 0

		}
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

/////////////////////////////////////////////////////////////////////////////////
// File Utility
/////////////////////////////////////////////////////////////////////////////////


func WriteFull(filename string, data [64][256]float32) {

	f, err := os.Create("./data/" + filename + ".csv")
	if err != nil {
		log.Fatalln(err)
	}

	writer := csv.NewWriter(f)

	var dataString [][]string
	for _,round := range data {
		var line []string
		for _, val := range round {
			line = append(line, strconv.FormatFloat(float64(val),'f',-1,32))
		}
		dataString = append(dataString, line)
	}

	err = writer.WriteAll(dataString)

	if err != nil {
		log.Fatalln(err)
	}

}


/////////////////////////////////////////////////////////////////////////////////
// Math utility
/////////////////////////////////////////////////////////////////////////////////

func Min(array [256]float32) float32 {

	min := array[0]

	for _, val := range array {
		if val < min {
			min = val
		}
	}

	return min
}

func Max(array [256]float32) float32 {

	max := array[0]

	for _, val := range array {
		if val > max {
			max = val
		}
	}

	return max
}

func Mean(array [256]float32) float32 {

	mean := 0.0
	count := 0.0
	for i, val := range array {
		mean = mean + float64(val)
		count = float64(i)

	}

	return float32(mean/count)
}