package main

import (
	"fmt"
)

type IntGenerator struct {
	Next   []chan int
	Result chan int
}

func NewIntGenerator(size int) IntGenerator {
	return IntGenerator{Next: []chan int{make(chan int, size)}}
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
	//res := limit(10, mapi(square, filter(isEven, numbers)))
	gen := NewIntGenerator(len(numbers))
	go ranger(numbers, &gen)
	lazyFilter(isEven, &gen)
	<-gen.Result

}

func sum(g IntGenerator) {
	var s int
	for n := range g.Next {
		fmt.Println(n)
		s += n
	}
	g.Result <- s
}

func ranger(input []int, ig *IntGenerator) {
	myChan := ig.Next[0]
	for _, i := range input {
		fmt.Println("pushed data")
		myChan <- i
	}
}

func lazyFilter(f filterfunc, ig *IntGenerator) {
	myChanIndex := len(ig.Next)
	if myChanIndex == 0 {
		panic("We need a source for input")
	}
	myChan := ig.Next[myChanIndex-1]
	nextChan := make(chan int)
	ig.Next = append(ig.Next, nextChan)
	fmt.Println("listening")
	for n := range myChan {
		fmt.Println("lazily filtering")
		if f(n) {
			//			nextChan <- n
			fmt.Println(n)
		}
	}
	fmt.Println("done filtering")
	close(nextChan)
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
