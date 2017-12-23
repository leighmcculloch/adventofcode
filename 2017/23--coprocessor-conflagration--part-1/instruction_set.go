package solution

var instructionSet = map[string]func(x int, y int, r <-chan int, s chan<- int) (newX int, jump int, skipped bool, yield bool){
	"set": set,
	"sub": sub,
	"mul": mul,
	"jnz": jnz,
}

func set(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	return y, 1, false, false
}

func sub(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	return x - y, 1, false, false
}

func mul(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	return x * y, 1, false, false
}

func jnz(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	if x != 0 {
		return x, y, false, false
	} else {
		return x, 1, true, false
	}
}
