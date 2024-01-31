package main

import (
	"fmt"
)

type FunctionName int

const (
	XOR      FunctionName = iota // Defualts "+",  xor when false
	Kfunc                        //Defualts true, use when true
	CHOOSE                       //Defualts true, use when true
	MAJOR                        //Defualts true, use when true
	SIGMA_0                      //Defualts true, use when true
	SIGMA_1                      //Defualts true, use when true
	SCHEDULE                     //Defualts true, use when true
)

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
		msgSchedule = plainMessageSchedule(chunk)
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

func plainMessageSchedule(chunk [16]uint32) [64]uint32 {

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
// Hash Auxilary
/////////////////////////////////////////////////////////////////////////////////

// XOR is opposite the rest, as the rest are default
func FNStoBS(names []FunctionName) [7]bool {

	var out [7]bool = [7]bool{false, true, true, true, true, true, true}

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
