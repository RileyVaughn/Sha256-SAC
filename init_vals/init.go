package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	N := 1000000

	var msgs [][]string
	for i := 0; i < N; i++ {
		msgs = append(msgs, generateMsg())
	}

	Write("init_vals_512", msgs)
}

func Write(filename string, data [][]string) {

	f, err := os.Create(filename + ".csv")
	if err != nil {
		log.Fatalln(err)
	}
	writer := csv.NewWriter(f)

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatalln(err)
	}

}

// 512-bit message
func generateMsg() []string {
	var msg [16]uint32
	for i := 0; i < 16; i++ {
		msg[i] = rand.Uint32()
	}

	msg_str := strings.Split(strings.Trim(fmt.Sprint(msg), "[]"), " ")

	return msg_str
}
