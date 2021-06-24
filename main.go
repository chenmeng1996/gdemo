package main

func main() {
	for range N(4) {
	}
}

// N 该函数不引发堆内存分配，而是分配在栈上
func N(n int) []struct{} {
	return make([]struct{}, n)
}
