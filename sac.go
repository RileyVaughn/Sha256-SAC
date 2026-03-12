package main

var INIT_MSGS [][]uint32


func main() {


	INIT_MSGS = CSVtoUint32(ReadCSV("./init_vals/init_vals_512"))

	// R1()
	// R2()
	// R3()
	// R4()
	// R5()
	// R6()


	

}


func R1(){

	RoundsSac("removed_1/C", []FunctionName{CHOOSE})
	RoundsSac("removed_1/M", []FunctionName{MAJOR})
	RoundsSac("removed_1/K", []FunctionName{Kfunc})
	RoundsSac("removed_1/X", []FunctionName{XOR})
	RoundsSac("removed_1/R", []FunctionName{SCHEDULE})
	RoundsSac("removed_1/S0", []FunctionName{SIGMA_0})
	RoundsSac("removed_1/S1", []FunctionName{SIGMA_1})

}


func R2(){

	RoundsSac("removed_2/CM", []FunctionName{CHOOSE,MAJOR})
	RoundsSac("removed_2/CK", []FunctionName{CHOOSE,Kfunc})
	RoundsSac("removed_2/CX", []FunctionName{CHOOSE,XOR})
	RoundsSac("removed_2/CR", []FunctionName{CHOOSE,SCHEDULE})
	RoundsSac("removed_2/CS0", []FunctionName{CHOOSE,SIGMA_0})
	RoundsSac("removed_2/CS1", []FunctionName{CHOOSE,SIGMA_1})

	RoundsSac("removed_2/MK", []FunctionName{MAJOR,Kfunc})
	RoundsSac("removed_2/MX", []FunctionName{MAJOR,XOR})
	RoundsSac("removed_2/MR", []FunctionName{MAJOR,SCHEDULE})
	RoundsSac("removed_2/MS0", []FunctionName{MAJOR,SIGMA_0})
	RoundsSac("removed_2/MS1", []FunctionName{MAJOR,SIGMA_1})

	RoundsSac("removed_2/KX", []FunctionName{Kfunc,XOR})
	RoundsSac("removed_2/KR", []FunctionName{Kfunc,SCHEDULE})
	RoundsSac("removed_2/KS0", []FunctionName{Kfunc,SIGMA_0})
	RoundsSac("removed_2/KS1", []FunctionName{Kfunc,SIGMA_1})

	RoundsSac("removed_2/XR", []FunctionName{XOR,SCHEDULE})
	RoundsSac("removed_2/XS0", []FunctionName{XOR,SIGMA_0})
	RoundsSac("removed_2/XS1", []FunctionName{XOR,SIGMA_1})

	RoundsSac("removed_2/RS0", []FunctionName{SCHEDULE,SIGMA_0})
	RoundsSac("removed_2/RS1", []FunctionName{SCHEDULE,SIGMA_1})

	RoundsSac("removed_2/S0S1", []FunctionName{SIGMA_0,SIGMA_1})

}


func R3() {

	RoundsSac("removed_3/CMK", []FunctionName{CHOOSE,MAJOR, Kfunc})
	RoundsSac("removed_3/CMX", []FunctionName{CHOOSE,MAJOR, XOR})
	RoundsSac("removed_3/CMR", []FunctionName{CHOOSE,MAJOR, SCHEDULE})
	RoundsSac("removed_3/CMS0", []FunctionName{CHOOSE,MAJOR, SIGMA_0})
	RoundsSac("removed_3/CMS1", []FunctionName{CHOOSE,MAJOR, SIGMA_1})

	RoundsSac("removed_3/CKX", []FunctionName{CHOOSE,Kfunc,XOR})
	RoundsSac("removed_3/CKR", []FunctionName{CHOOSE,Kfunc,SCHEDULE})
	RoundsSac("removed_3/CKS0", []FunctionName{CHOOSE,Kfunc,SIGMA_0})
	RoundsSac("removed_3/CKS1", []FunctionName{CHOOSE,Kfunc,SIGMA_1})

	RoundsSac("removed_3/CXR", []FunctionName{CHOOSE,XOR,SCHEDULE})
	RoundsSac("removed_3/CXS0", []FunctionName{CHOOSE,XOR,SIGMA_0})
	RoundsSac("removed_3/CXS1", []FunctionName{CHOOSE,XOR,SIGMA_1})

	RoundsSac("removed_3/CRS0", []FunctionName{CHOOSE,SCHEDULE,SIGMA_0})
	RoundsSac("removed_3/CRS1", []FunctionName{CHOOSE,SCHEDULE,SIGMA_1})

	RoundsSac("removed_3/CS0S1", []FunctionName{CHOOSE,SIGMA_0,SIGMA_1})

	///////////////

	RoundsSac("removed_3/MKX", []FunctionName{MAJOR, Kfunc, XOR})
	RoundsSac("removed_3/MKR", []FunctionName{MAJOR, Kfunc, SCHEDULE})
	RoundsSac("removed_3/MKS0", []FunctionName{MAJOR, Kfunc, SIGMA_0})
	RoundsSac("removed_3/MKS1", []FunctionName{MAJOR, Kfunc, SIGMA_1})
	
	RoundsSac("removed_3/MXR", []FunctionName{MAJOR, XOR, SCHEDULE})
	RoundsSac("removed_3/MXS0", []FunctionName{MAJOR, XOR, SIGMA_0})
	RoundsSac("removed_3/MXS1", []FunctionName{MAJOR, XOR, SIGMA_1})

	RoundsSac("removed_3/MRS0", []FunctionName{MAJOR, SCHEDULE, SIGMA_0})	
	RoundsSac("removed_3/MRS1", []FunctionName{MAJOR, SCHEDULE, SIGMA_1})

	RoundsSac("removed_3/MS0S1", []FunctionName{MAJOR, SIGMA_0, SIGMA_1})

	///////////////

	RoundsSac("removed_3/KXR", []FunctionName{Kfunc, XOR, SCHEDULE})
	RoundsSac("removed_3/KXS0", []FunctionName{Kfunc, XOR, SIGMA_0})
	RoundsSac("removed_3/KXS1", []FunctionName{Kfunc, XOR, SIGMA_1})

	RoundsSac("removed_3/KRS0", []FunctionName{Kfunc, SCHEDULE, SIGMA_0})
	RoundsSac("removed_3/KRS1", []FunctionName{Kfunc, SCHEDULE, SIGMA_1})

	RoundsSac("removed_3/KS0S1", []FunctionName{Kfunc, SIGMA_0, SIGMA_1})

	///////////////

	RoundsSac("removed_3/XRS0", []FunctionName{XOR, SCHEDULE, SIGMA_0})
	RoundsSac("removed_3/XRS1", []FunctionName{XOR, SCHEDULE, SIGMA_1})

	RoundsSac("removed_3/XS0S1", []FunctionName{XOR, SIGMA_0, SIGMA_1})

	///////////////

	RoundsSac("removed_3/RS0S1", []FunctionName{SCHEDULE, SIGMA_0, SIGMA_1})

}


func R4() {

	RoundsSac("removed_4/CMKX", []FunctionName{CHOOSE,MAJOR, Kfunc, XOR})
	RoundsSac("removed_4/CMKR", []FunctionName{CHOOSE,MAJOR, Kfunc, SCHEDULE})
	RoundsSac("removed_4/CMKS0", []FunctionName{CHOOSE,MAJOR, Kfunc, SIGMA_0})
	RoundsSac("removed_4/CMKS1", []FunctionName{CHOOSE,MAJOR, Kfunc, SIGMA_1})

	RoundsSac("removed_4/CMXR", []FunctionName{CHOOSE,MAJOR, XOR, SCHEDULE})
	RoundsSac("removed_4/CMXS0", []FunctionName{CHOOSE,MAJOR, XOR, SIGMA_0})
	RoundsSac("removed_4/CMXS1", []FunctionName{CHOOSE,MAJOR, XOR, SIGMA_1})

	RoundsSac("removed_4/CMRS0", []FunctionName{CHOOSE,MAJOR, SCHEDULE, SIGMA_0})
	RoundsSac("removed_4/CMRS1", []FunctionName{CHOOSE,MAJOR, SCHEDULE, SIGMA_1})


	RoundsSac("removed_4/CMS0S1", []FunctionName{CHOOSE,MAJOR, SIGMA_0, SIGMA_1})

	/////

	RoundsSac("removed_4/CKXR", []FunctionName{CHOOSE,Kfunc,XOR, SCHEDULE})
	RoundsSac("removed_4/CKXS0", []FunctionName{CHOOSE,Kfunc,XOR, SIGMA_0})
	RoundsSac("removed_4/CKXS1", []FunctionName{CHOOSE,Kfunc,XOR, SIGMA_1})

	RoundsSac("removed_4/CKRS0", []FunctionName{CHOOSE,Kfunc,SCHEDULE, SIGMA_0})
	RoundsSac("removed_4/CKRS1", []FunctionName{CHOOSE,Kfunc,SCHEDULE, SIGMA_1})

	RoundsSac("removed_4/CKS0S1", []FunctionName{CHOOSE,Kfunc,SIGMA_0, SIGMA_1})

	/////

	RoundsSac("removed_4/CXRS0", []FunctionName{CHOOSE,XOR,SCHEDULE, SIGMA_0})
	RoundsSac("removed_4/CXRS1", []FunctionName{CHOOSE,XOR,SCHEDULE, SIGMA_1})

	RoundsSac("removed_4/CXS0S1", []FunctionName{CHOOSE,XOR,SIGMA_0, SIGMA_1})

	/////

	RoundsSac("removed_4/CRS0S1", []FunctionName{CHOOSE,SCHEDULE,SIGMA_0, SIGMA_1})

	///////////////

	RoundsSac("removed_4/MKXR", []FunctionName{MAJOR, Kfunc, XOR,SCHEDULE})
	RoundsSac("removed_4/MKXS0", []FunctionName{MAJOR, Kfunc, XOR,SIGMA_0})
	RoundsSac("removed_4/MKXS1", []FunctionName{MAJOR, Kfunc, XOR,SIGMA_1})

	RoundsSac("removed_4/MKRS0", []FunctionName{MAJOR, Kfunc, SCHEDULE, SIGMA_0})
	RoundsSac("removed_4/MKRS1", []FunctionName{MAJOR, Kfunc, SCHEDULE, SIGMA_1})

	RoundsSac("removed_4/MKS0S1", []FunctionName{MAJOR, Kfunc, SIGMA_0, SIGMA_1})

	/////

	RoundsSac("removed_4/MXRS0", []FunctionName{MAJOR, XOR, SCHEDULE, SIGMA_0})
	RoundsSac("removed_4/MXRS1", []FunctionName{MAJOR, XOR, SCHEDULE, SIGMA_1})

	RoundsSac("removed_4/MXS0S1", []FunctionName{MAJOR, XOR, SIGMA_0, SIGMA_1})

	/////

	RoundsSac("removed_4/MRS0S1", []FunctionName{MAJOR, SCHEDULE, SIGMA_0, SIGMA_1})

	///////////////

	RoundsSac("removed_4/KXRS0", []FunctionName{Kfunc, XOR, SCHEDULE, SIGMA_0})
	RoundsSac("removed_4/KXRS1", []FunctionName{Kfunc, XOR, SCHEDULE, SIGMA_1})

	RoundsSac("removed_4/KXS0S1", []FunctionName{Kfunc, XOR, SIGMA_0, SIGMA_1})

	/////

	RoundsSac("removed_4/KRS0S1", []FunctionName{Kfunc, SCHEDULE, SIGMA_0, SIGMA_1})

	///////////////

	RoundsSac("removed_4/XRS0S1", []FunctionName{XOR, SCHEDULE, SIGMA_0, SIGMA_1})


}

func R5() {

	RoundsSac("removed_5/CMKXR", []FunctionName{CHOOSE,MAJOR, Kfunc, XOR, SCHEDULE})
	RoundsSac("removed_5/CMKXS0", []FunctionName{CHOOSE,MAJOR, Kfunc, XOR, SIGMA_0})
	RoundsSac("removed_5/CMKXS1", []FunctionName{CHOOSE,MAJOR, Kfunc, XOR, SIGMA_1})
	
	RoundsSac("removed_5/CMKRS0", []FunctionName{CHOOSE,MAJOR, Kfunc, SCHEDULE, SIGMA_0})
	RoundsSac("removed_5/CMKRS1", []FunctionName{CHOOSE,MAJOR, Kfunc, SCHEDULE, SIGMA_1})

	RoundsSac("removed_5/CMKS0S1", []FunctionName{CHOOSE,MAJOR, Kfunc, SIGMA_0, SIGMA_1})

	////

	RoundsSac("removed_5/CMXRS0", []FunctionName{CHOOSE,MAJOR, XOR, SCHEDULE, SIGMA_0})
	RoundsSac("removed_5/CMXRS1", []FunctionName{CHOOSE,MAJOR, XOR, SCHEDULE, SIGMA_1})

	RoundsSac("removed_5/CMXS0S1", []FunctionName{CHOOSE,MAJOR, XOR, SIGMA_0, SIGMA_1})

	////

	RoundsSac("removed_5/CMRS0S1", []FunctionName{CHOOSE,MAJOR, SCHEDULE, SIGMA_0,SIGMA_1})

	////////

	RoundsSac("removed_5/CKXRS0", []FunctionName{CHOOSE,Kfunc,XOR, SCHEDULE, SIGMA_0})
	RoundsSac("removed_5/CKXRS1", []FunctionName{CHOOSE,Kfunc,XOR, SCHEDULE, SIGMA_1})

	RoundsSac("removed_5/CKXS0S1", []FunctionName{CHOOSE,Kfunc,XOR, SIGMA_0, SIGMA_1})

	////
	
	RoundsSac("removed_5/CKRS0S1", []FunctionName{CHOOSE,Kfunc,SCHEDULE, SIGMA_0, SIGMA_1})

	////////

	RoundsSac("removed_5/CXRS0S1", []FunctionName{CHOOSE,XOR,SCHEDULE, SIGMA_0, SIGMA_1})
	
	///////////////

	RoundsSac("removed_5/MKXRS0", []FunctionName{MAJOR, Kfunc, XOR,SCHEDULE, SIGMA_0})
	RoundsSac("removed_5/MKXRS1", []FunctionName{MAJOR, Kfunc, XOR,SCHEDULE, SIGMA_1})

	RoundsSac("removed_5/MKXS0S1", []FunctionName{MAJOR, Kfunc, XOR,SIGMA_0, SIGMA_1})

	////

	RoundsSac("removed_5/MKRS0S1", []FunctionName{MAJOR, Kfunc, SCHEDULE, SIGMA_0, SIGMA_1})

	////////

	RoundsSac("removed_5/MXRS0S1", []FunctionName{MAJOR, XOR, SCHEDULE, SIGMA_0, SIGMA_1})
	
	///////////////

	RoundsSac("removed_5/KXRS0S1", []FunctionName{Kfunc, XOR, SCHEDULE, SIGMA_0, SIGMA_1})




}


func R6() {

	//Sigma1
	RoundsSac("removed_6/CMKXRS0", []FunctionName{CHOOSE,MAJOR, Kfunc, XOR, SCHEDULE, SIGMA_0})

	//Sigm0
	RoundsSac("removed_6/CMKXRS1", []FunctionName{CHOOSE,MAJOR, Kfunc, XOR,SCHEDULE, SIGMA_1})

	//Schedule
	RoundsSac("removed_6/CMKXS0S1", []FunctionName{CHOOSE,MAJOR, Kfunc, XOR, SIGMA_0, SIGMA_1})

	//Xor
	RoundsSac("removed_6/CMKRS0S1", []FunctionName{CHOOSE,MAJOR, Kfunc, SCHEDULE, SIGMA_0, SIGMA_1})

	//kfunc
	RoundsSac("removed_6/CMXRS0S1", []FunctionName{CHOOSE,MAJOR, XOR, SCHEDULE, SIGMA_0, SIGMA_1})

	//Major
	RoundsSac("removed_6/CKXRS0S1", []FunctionName{CHOOSE,Kfunc,XOR, SCHEDULE, SIGMA_0, SIGMA_1})

	//Chooose
	RoundsSac("removed_6/MKXRS0S1", []FunctionName{MAJOR, Kfunc, XOR,SCHEDULE, SIGMA_0, SIGMA_1})





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

	msgs := INIT_MSGS
	// msgs := CSVtoUint32(ReadCSV("./init_vals/init_vals_512"))
	//msgs = msgs[0:10]
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
