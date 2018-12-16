package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	sum := 0

	for {
		lineBytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		line := string(lineBytes)

		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += value
	}

	fmt.Println("Solution:", sum)
}
