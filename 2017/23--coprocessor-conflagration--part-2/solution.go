// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/23.
package solution

func Part2() int {
	b := 84
	c := b

	b *= 100
	b += 100000
	c = b
	c += 17000

	h := 0

	for {
		f := 1
		g := 0

		// BEGIN Instructions in original input, removed and replaced with isPrime:
		// d := 2
		// for {
		// 	e := 2
		// 	for {
		// 		g = d*e - b
		// 		if g == 0 {
		// 			f = 0
		// 		}
		// 		e++
		// 		g = e - b
		// 		if g == 0 {
		// 			break
		// 		}
		// 	}
		// 	d++
		// 	g = d - b
		// 	if g == 0 {
		// 		break
		// 	}
		// }
		// END

		// BEGIN Optimized instructions replacing code above:
		if !isPrime(b) {
			f = 0
		}
		// END

		if f == 0 {
			h++
		}

		g = b - c

		if g == 0 {
			return h
		}
		b += 17
	}
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
