package main

import (
	"fmt"
)

type supplier func() int

type IntGenerator struct {
	Supplier chan int   // Supplier<T>
	Next     []chan int // bifunc<T,T>
	Consumer chan int   // Conumser<T>
	Size     int
}

func NewIntGenerator(size int) IntGenerator {
	return IntGenerator{
		Supplier: make(chan int, size),
		Next:     []chan int{},
		Supplier: make(chan int, size),
		Size:     size,
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
	//res := limit(10, mapi(square, filter(isEven, numbers)))
	gen := NewIntGenerator(len(numbers))
	go ranger(numbers, &gen)
	go lazyFilter(isEven, &gen)
	go lazyMap(square, &gen)
	s := sum(&gen)
	fmt.Printf("DID WE GET THE RES?? %v\n", s)
}

// terminator func?
func sum(ig *IntGenerator) int {
	myChan := ig.Next[len(ig.Next)-1]
	sum := 0
	for n := range myChan {
		sum += n
	}
	return sum
}

func ranger(input []int, ig *IntGenerator) {
	myChan := ig.Next[0]
	for _, i := range input {
		fmt.Println("pushed data")
		myChan <- i
	}
	close(myChan)
}

func lazyFilter(f filterfunc, ig *IntGenerator) {
	myChanIndex := len(ig.Next)
	if myChanIndex == 0 {
		panic("We need a source for input")
	}
	myChan := ig.Next[myChanIndex-1]
	nextChan := make(chan int, cap(ig.Next)+10)
	ig.Next = append(ig.Next, nextChan)
	fmt.Println("listening")
	for n := range myChan {
		fmt.Println("lazily filtering")
		if f(n) {
			nextChan <- n
		}
	}
	fmt.Println("done filtering")
	ig.Result <- 1336
	close(nextChan)
}

func lazyMap(f func(int) int, ig *IntGenerator) {
	myChanIndex := len(ig.Next)
	myChan := ig.Next[myChanIndex-1]
	nextChan := make(chan int)
	ig.Next = append(ig.Next, nextChan)
	for n := range myChan {
		fmt.Println("lazily mapping")
		n = f(n)
	}
	fmt.Println("done mapping")
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
