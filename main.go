package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

func main() {
	fileName := "./files/SRV_LINEAS.DAT"

	processStartTime := time.Now()
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lineContent := strings.Split(string(fileContent), "\n")
	lines := make([]string, len(lineContent))
	copy(lines, lineContent[:])
	fmt.Println(len(lines))
	fmt.Println(cap(lines))
	chunkLen := 10000
	chunkStart := 0
	chunkEnd := chunkLen
	quantityChunk := len(lines) / chunkLen
	var wg sync.WaitGroup
	wg.Add(quantityChunk)
	for i := 1; i <= quantityChunk; i++ {
		go sendDataToSDL(lines, chunkStart, chunkEnd, &wg)
		chunkStart = chunkEnd
		chunkEnd = chunkLen * i
	}
	fmt.Println(lines[chunkEnd : len(lines)-1])
	wg.Wait()
	processEndTime := time.Now()
	fmt.Println(processStartTime)
	fmt.Println(processEndTime)
	fmt.Println(quantityChunk)
}

func sendDataToSDL(lines []string, chunkStart int, chunkEnd int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(lines[chunkStart:chunkEnd])
}
