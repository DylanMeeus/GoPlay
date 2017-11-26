package main


import (
    "fmt"
    "math"
)

func main(){
    fmt.Println("Hello World")
    squares := createSquares(int64(math.Pow(10, 8))) // (10^4)^2 == 10^8 :-)
    
}


// create all the squares below our limit
func createSquares(total int64) []int64{

    squares := make([]int64, 0)
    for i := int64(0); i*i < total; i++{
        square := i * i
        squares = append(squares,square)
    }

    return squares
}
