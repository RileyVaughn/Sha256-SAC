package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadCSV(filename string) [][]string {

	f, err := os.Open(filename + ".csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(f)

	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	return data
}

func WriteCSV(filename string, data [][]string) {

	f, err := os.Create("./data/" + filename + ".csv")
	if err != nil {
		log.Fatalln(err)
	}
	writer := csv.NewWriter(f)

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatalln(err)
	}

}

func WriteCSV64(data [][][]string) {

	for i := 0; i < 64; i++ {
		WriteCSV(fmt.Sprintf("rounds/round_%v", i+1), data[i])
	}

}

func CSVtoUint32(data [][]string) [][]uint32 {

	var data32 [][]uint32
	for _, msg := range data {
		var msg32 []uint32
		for _, word := range msg {
			num, err := strconv.Atoi(word)
			if err != nil {
				log.Fatalln(err)
			}
			msg32 = append(msg32, uint32(num))
		}
		data32 = append(data32, msg32)
	}

	return data32
}

func Uint32x8ToUint8x256(msg *[8]uint32) *[256]uint8 {

	var extMsg [256]uint8

	for i := 0; i < 8; i++ {
		tempMsg := Uint32ToUint8x32(msg[i])
		for j := 0; j < 32; j++ {
			extMsg[i*32+j] = tempMsg[j]
		}

	}

	return &extMsg
}

func Uint32ToUint8x32(word uint32) *[32]uint8 {

	var wordBits [32]uint8
	for i := 0; i < 32; i++ {
		wordBits[i] = uint8(word % 2)
		word = word >> 1
	}
	return &wordBits
}

func AddToDepMat(depMat *[512][256]float32, newMat *[512][256]uint8) {

	for i := 0; i < 512; i++ {
		for j := 0; j < 256; j++ {
			depMat[i][j] = depMat[i][j] + float32(newMat[i][j])
		}
	}
}

func AddToDepMat64(depMat *[64][512][256]float32, newMat *[64][512][256]uint8) {
	for i := 0; i < 64; i++ {
		AddToDepMat(&depMat[i], &newMat[i])
	}
}

func DepMatDiv(depMat *[512][256]float32, size float32) {

	for i := 0; i < 512; i++ {
		for j := 0; j < 256; j++ {
			depMat[i][j] = depMat[i][j] / size
		}
	}

}

func DepMatDiv64(depMat *[64][512][256]float32, size float32) {
	for i := 0; i < 64; i++ {
		DepMatDiv(&depMat[i], size)
	}
}

func DepMatToCSV(depMat *[512][256]float32) [][]string {

	var csv [][]string
	for i := 0; i < 512; i++ {
		var line []string
		for j := 0; j < 256; j++ {
			line = append(line, fmt.Sprint(depMat[i][j]))
		}
		csv = append(csv, line)
	}
	return csv
}

func DepMatToCSV64(depMat *[64][512][256]float32) [][][]string {

	var csvs [][][]string

	for i := 0; i < 64; i++ {
		csvs = append(csvs, DepMatToCSV(&depMat[i]))
	}
	return csvs
}
