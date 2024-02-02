package main

func main() {

	FullSAC("no_sched_fullCF", []FunctionName{SCHEDULE})

	//RoundsSac("no_sched", []FunctionName{SCHEDULE})
}

func FullSAC(fileName string, rmvFuncs []FunctionName) {

	msgs := CSVtoUint32(ReadCSV("./init_vals/init_vals_512"))

	var depMatrix [512][256]float32
	for _, msg := range msgs {
		AddToDepMat(&depMatrix, MeasureSAC([16]uint32(msg), rmvFuncs))
	}
	DepMatDiv(&depMatrix, float32(len(msgs)))
	WriteCSV(fileName, DepMatToCSV((&depMatrix)))

}

func RoundsSac(dirName string, rmvFuncs []FunctionName) {

	msgs := CSVtoUint32(ReadCSV("./init_vals/init_vals_512"))
	msgs = msgs[0:1000]
	var depMatrices [64][512][256]float32
	for _, msg := range msgs {
		AddToDepMat64(&depMatrices, MeasureSac64([16]uint32(msg), rmvFuncs))
	}
	DepMatDiv64(&depMatrices, float32(len(msgs)))
	WriteCSV64(dirName, DepMatToCSV64(&depMatrices))

}

func MeasureSAC(msg [16]uint32, rmvFuncs []FunctionName) *[512][256]uint8 {

	hash, _ := Sha256_compress_verbose(msg, H, rmvFuncs)

	var depMatrix [512][256]uint8

	for i := uint32(0); i < 512; i++ {
		msg2 := FlipBit(msg, i)
		hash2, _ := Sha256_compress_verbose(msg2, H, rmvFuncs)
		depMatrix[i] = *Uint32x8ToUint8x256(XorHash(&hash, &hash2))
	}

	return &depMatrix
}

func MeasureSac64(msg [16]uint32, rmvFuncs []FunctionName) *[64][512][256]uint8 {

	_, hash := Sha256_compress_verbose(msg, H, rmvFuncs)

	var depMatrices [64][512][256]uint8

	for i := uint32(0); i < 512; i++ {
		msg2 := FlipBit(msg, i)
		_, hash2 := Sha256_compress_verbose(msg2, H, rmvFuncs)
		for j := 0; j < 64; j++ {
			depMatrices[j][i] = *Uint32x8ToUint8x256(XorHash(&hash[j], &hash2[j]))
		}

	}

	return &depMatrices

}

func XorHash(hash1 *[8]uint32, hash2 *[8]uint32) *[8]uint32 {

	var rv [8]uint32
	for i := 0; i < 8; i++ {
		rv[i] = hash1[i] ^ hash2[i]
	}
	return &rv
}

func FlipBit(msg [16]uint32, index uint32) [16]uint32 {

	byteChoice := index / 32
	bitChoice := index % 32
	flip := uint32(1) << bitChoice
	msg[byteChoice] = msg[byteChoice] ^ flip
	return msg
}
