package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	fileName := "SRV_LINEAS.DAT"

	fmt.Println(time.Now())
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	originalLines := make([]string, len(lines))
	copy(originalLines, lines[:])
	fmt.Println(len(originalLines))
	fmt.Println(cap(originalLines))
	fmt.Println(originalLines[1100:1110])
	fmt.Println(time.Now())
	//for index := 0; index < len(lines); index++ {
	//fmt.Println(lines[index])
	//}

	/*
		f, _ := os.Open(fileName)
		// Create new Scanner.
		scanner := bufio.NewScanner(f)
		result := []string{}
		// Use Scan.
		for scanner.Scan() {
			line := scanner.Text()
			// Append line to result.
			result = append(result, line)
		}
		fmt.Println(result)
	*/
	/*
		fptr := flag.String("fpath", fileName, "file path to read from")
		flag.Parse()

		f, err := os.Open(*fptr)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = f.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		r := bufio.NewReader(f)
		b := make([]byte, 12)
		for {
			_, err := r.Read(b)
			if err != nil {
				fmt.Println("Error reading file:", err)
				break
			}
			fmt.Println(string(b))
		}
	*/
}
