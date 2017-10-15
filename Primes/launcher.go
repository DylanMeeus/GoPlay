package main


import (
    "fmt"
    "./Sieves"
)

func main(){
    fmt.Println("Generating primes!")
    primes := Primes.Sieve(50000000)
    fmt.Println(primes[len(primes)-1])
}
