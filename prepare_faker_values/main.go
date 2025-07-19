package main

import (
	"flag"
	"fmt"
	"os"
)

var count int
var valueLength int

func main() {
	flag.IntVar(&count, "count", 100_000, "count")
	flag.IntVar(&valueLength, "value-length", 100, "value length")
	flag.Parse()
	fmt.Printf("start to generate %d faker values, value length: %d\n", count, valueLength)

	userHomeDir := os.Getenv("HOME")
	filePath := userHomeDir + "/faker-value-as-json.txt"
	filePath2 := userHomeDir + "/faker-value-as-text.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	file2, err := os.OpenFile(filePath2, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	mocker := NewMocker()

	lenJsonSum := 0
	lenTextSum := 0
	for i := 0; i < count; i++ {
		json := mocker.FakerValueAsJson(valueLength)
		_, err := file.Write([]byte(json + "\n"))
		if err != nil {
			panic(err)
		}
		lenJsonSum += len(json)

		text := mocker.FakerValue(valueLength)
		_, err = file2.Write([]byte(text + "\n"))
		if err != nil {
			panic(err)
		}
		lenTextSum += len(text)
	}
	fmt.Printf("Avg json length: %d\n", lenJsonSum/count)
	fmt.Printf("Avg text length: %d\n", lenTextSum/count)

	_ = file.Close()
	_ = file2.Close()
}
