package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(bytes.NewReader(b))

	sum := 0
	sumsSeen := map[int]bool{0: true}

	for {
		lineBytes, _, err := r.ReadLine()
		if err == io.EOF {
			r = bufio.NewReader(bytes.NewReader(b))
			continue
		} else if err != nil {
			panic(err)
		}

		line := string(lineBytes)

		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += value

		hasSeen := sumsSeen[sum]
		if hasSeen {
			fmt.Println("Solution:", sum)
			return
		}

		sumsSeen[sum] = true
	}

	fmt.Println("Finished parsing file without finding repeating frequency.")
}
