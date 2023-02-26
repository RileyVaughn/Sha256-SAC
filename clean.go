package main

import (
	io "io/ioutil"
	"log"
	"strings"
)

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
