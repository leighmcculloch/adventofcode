// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/15.
package solution

func Part1(startingA, startingB int) int {
	prevA := startingA
	prevB := startingB
	count := 0
	for i := 0; i < 40000000; i++ {
		prevA = prevA * 16807 % 2147483647
		prevB = prevB * 48271 % 2147483647
		if uint16(prevA) == uint16(prevB) {
			count++
		}
	}
	return count
}

func Part2(startingA, startingB int) int {
	stop := make(chan struct{})
	defer close(stop)

	streamA := stream(startingA, 16807, 4, stop)
	streamB := stream(startingB, 48271, 8, stop)

	matched := 0
	for i := 0; i < 5000000; i++ {
		valueA := <-streamA
		valueB := <-streamB
		if uint16(valueA) == uint16(valueB) {
			matched++
		}
	}

	return matched
}

func stream(start, factor, divisor int, stop <-chan struct{}) <-chan int {
	stream := make(chan int, 1000)
	go func() {
		defer close(stream)
		prev := start
		for {
			prev = prev * factor % 2147483647
			if prev%divisor == 0 {
				stream <- prev
			}
			select {
			case <-stop:
				break
			default:
			}
		}
	}()
	return stream
}
