package main

func main() {
}

//export Fibonacci
func Fibonacci(in uint32) uint32 {
	if in <= 1 {
		return in
	}
	return Fibonacci(in-1) + Fibonacci(in-2)
}
