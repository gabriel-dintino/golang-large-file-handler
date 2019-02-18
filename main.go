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

	horaInicio := time.Now()
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	originalLines := make([]string, len(lines))
	copy(originalLines, lines[:])
	fmt.Println(len(originalLines))
	fmt.Println(cap(originalLines))
	// tamaño de chunk
	chunk := 10000
	inicio := 0
	fin := 10000
	cantidadChunk := len(originalLines) / chunk
	var wg sync.WaitGroup
	wg.Add(cantidadChunk)
	for i := 1; i <= cantidadChunk; i++ {
		//fmt.Println(originalLines[inicio:fin])
		go sendDataToSDL(originalLines, inicio, fin, &wg)
		inicio = fin
		fin = chunk * i
	}
	fmt.Println(originalLines[fin : len(originalLines)-1])
	wg.Wait()
	horaFin := time.Now()
	fmt.Println(horaInicio)
	fmt.Println(horaFin)
	fmt.Println(cantidadChunk)
}

func sendDataToSDL(originalLines []string, inicio int, fin int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(originalLines[inicio:fin])
}
