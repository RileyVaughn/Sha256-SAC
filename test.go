package main

import (
	"fmt"
	io "io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Test() {

	if TestSha256("SHA256ShortMsg.rsp") {
		fmt.Println("Sha256 passed all short test vectors")
	} else {
		fmt.Println("Sha256 failed a short test vectors")
	}

	if TestSha256("SHA256LongMsg.rsp") {
		fmt.Println("Sha256 passed all Long test vectors")
	} else {
		fmt.Println("Sha256 failed a Long test vectors")
	}

}

func TestSha256(filename string) bool {

	pass := true

	Clean(filename)

	lengths, msgs, hashs := ReadClean(filename)

	// The test input is represented in hex where evry 2 items make a byte. This converts their style to my style.
	var repMsgs []string
	for _, line := range msgs {
		var temp []byte
		for i, _ := range line {
			if i%2 == 0 {
				temp = append(temp, (asciiToNum(line[i])<<4)+asciiToNum(line[i+1]))

			}

		}
		// Special case where input is empty
		if len(temp) == 1 && temp[0] == 0 {
			repMsgs = append(repMsgs, "")
		} else {
			repMsgs = append(repMsgs, string(temp))
		}
	}

	//Finally tests
	for i := range lengths {
		if hashs[i] != Sha256(repMsgs[i]) {
			fmt.Printf("%v   %v \n", hashs[i], Sha256(repMsgs[i]))
			pass = false
		}
	}

	return pass
}

func ReadClean(filename string) ([]int64, []string, []string) {
	// Read file
	file, err := io.ReadFile("./cleanTV/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	// Read data as string and remove top comments
	data := string(file)
	//Split data by line
	dataSlice := strings.Split(data, "\n")

	var lengths []int64
	var msgs []string
	var hashs []string

	for i, line := range dataSlice {
		if i%3 == 0 {
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			lengths = append(lengths, num)
		} else if i%3 == 1 {
			msgs = append(msgs, line)
		} else {
			hashs = append(hashs, line)
		}

	}

	return lengths, msgs, hashs
}

func asciiToNum(a byte) byte {
	if a < 58 {
		return byte(a - '0')
	} else {
		return byte(a - 87)
	}
}

func Clean(filename string) {

	// Read file
	file, err := io.ReadFile("./testvectors/" + filename)

	if err != nil {
		log.Fatal(err)
	}
	// Read data as string and remove top comments
	data := string(file)
	dataSlice := strings.Split(data, "\n")
	dataSlice = dataSlice[7:]

	//Remove empty strings
	var temp []string
	for _, line := range dataSlice {
		if len(line) > 1 {
			temp = append(temp, line)
		}
	}
	dataSlice = temp

	//Remove prefix and leave only the values
	for i, line := range dataSlice {
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "\r", "")
		dataSlice[i] = strings.Split(line, "=")[1]
	}

	data = strings.Join(dataSlice, "\n")

	// Write new file
	err = io.WriteFile("./cleanTV/"+filename, []byte(data), 0777)
	if err != nil {
		log.Fatal(err)
	}

}
