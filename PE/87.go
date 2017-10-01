package main


import (
    "fmt"
    "../Primes/Sieves"
)


func main(){
    fmt.Println("Solving power triples")

    // 7071 is the square root of 50.000.000 which is the max of the sum
    limit := 50000000
    primes := Primes.Sieve(7072)

    results := make([]bool,limit)
    for double := 0; double < len(primes); double++{
        for triple := 0; triple < len(primes); triple++{
            for quad := 0; quad < len(primes); quad++{
                a := primes[double] * primes[double]
                b := primes[triple] * primes[triple] * primes[triple]
                c := primes[quad] * primes[quad] * primes[quad] * primes[quad]
                sum := a + b + c
                if sum < limit {
                    results[sum] = true
                }
            }
        }
    }

    sum := 0
    for i := 0; i < len(results); i++{
        if results[i]{
            sum++
        }
    }
    fmt.Println(sum)
}

func contains(arr *[]int, num int) bool{
    for i := 0; i < len(*arr); i++{
        if (*arr)[i] == num{
            return true
        }
    }
    return false
}