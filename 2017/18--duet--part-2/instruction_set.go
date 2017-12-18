package solution

var instructionSet = map[string]func(x int, y int, r <-chan int, s chan<- int) (newX int, jump int, skipped bool, yield bool){
	"set": set,
	"add": add,
	"mul": mul,
	"mod": mod,
	"snd": snd,
	"rcv": rcv,
	"jgz": jgz,
}

func set(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	return y, 1, false, false
}

func add(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	return x + y, 1, false, false
}

func mul(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	return x * y, 1, false, false
}

func mod(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	return x % y, 1, false, false
}

func snd(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	select {
	case s <- x:
		return x, 1, false, false
	default:
		return x, 0, false, true
	}
}

func rcv(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	select {
	case xr := <-r:
		return xr, 1, false, false
	default:
		return x, 0, false, true
	}
}

func jgz(x int, y int, r <-chan int, s chan<- int) (int, int, bool, bool) {
	if x > 0 {
		return x, y, false, false
	} else {
		return x, 1, true, false
	}
}
