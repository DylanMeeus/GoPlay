package main

import (
	"fmt"
)

type IntGenerator struct {
	Next   chan int
	Result chan int
}

func NewIntGenerator(size int) IntGenerator {
	return IntGenerator{
		Next:   make(chan int, size),
		Result: make(chan int),
	}
}

type filterfunc func(i int) bool

func main() {
	/*
		i := limit(10, (map(mapfunc, filter(filterfunc, numbers))))

	*/
	numbers := []int{}
	for i := 0; i < 22; i++ {
		numbers = append(numbers, i)
	}
	res := limit(10, mapi(func(i int) int { return i * i }, filter(func(i int) bool { return i%2 == 0 }, numbers)))
	fmt.Printf("%v\n", res)
	gen := NewIntGenerator(len(numbers))
	go sum(gen)
	for _, n := range numbers {
		gen.Next <- n
	}
	close(gen.Next)
	i := <-gen.Result
	fmt.Printf("i %v\n", i)
}

func sum(g IntGenerator) {
	var s int
	for n := range g.Next {
		fmt.Println(n)
		s += n
	}
	g.Result <- s
}

func filter(f filterfunc, is []int) []int {
	out := []int{}
	for _, i := range is {
		fmt.Println("filtering")
		if f(i) {
			out = append(out, i)
		}
	}
	return out
}

func mapi(f func(int) int, is []int) []int {
	out := make([]int, len(is))
	for index, i := range is {
		fmt.Println("mapping")
		out[index] = f(i)
	}
	return out
}

func limit(max int, is []int) []int {
	return is[:max]
}
