package Primes


import (
    "math"
)

func Sieve(n int) []int{
    revisedLimit := n + 2
    a := make([]bool, revisedLimit)
    setTrue(&a)

    // the first 2 are false as 0, 1 != prime
    a[0] = false
    a[1] = false

    limit := math.Sqrt(float64(revisedLimit))
    for i := 2; float64(i) < limit; i++{
        if a[i] {
            k := 0
            for j := i*i;j < revisedLimit; j = (i*i) + (k * i){
                a[j] = false
                k++
            }
        }
    }

    output := []int{}
    for i := 0; i < len(a); i++{
        if a[i]{
            output = append(output,i)
        }
    }
    return output
}


func setTrue(arr *[]bool){
    for i:= 0; i < len(*arr); i++{
        (*arr)[i] = true;
    }
}
