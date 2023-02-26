package main

func main() {
	Test()
}

func Sha256_compress_verbose(chunk []uint32, hash [8]uint32) [64][8]uint32 {
	msgSchedule := createMessageSchedule(chunk)

	a := hash[0]
	b := hash[1]
	c := hash[2]
	d := hash[3]
	e := hash[4]
	f := hash[5]
	g := hash[6]
	h := hash[7]

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

	hash[0] = (hash[0] + a)
	hash[1] = (hash[1] + b)
	hash[2] = (hash[2] + c)
	hash[3] = (hash[3] + d)
	hash[4] = (hash[4] + e)
	hash[5] = (hash[5] + f)
	hash[6] = (hash[6] + g)
	hash[7] = (hash[7] + h)

	return rounds

}
