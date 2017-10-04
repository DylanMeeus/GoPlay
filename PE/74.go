package main


import (
    "fmt"
    "math"
)


func main(){
    values := map[int] int{}
    s := 0
    for i := 0; i < 1000000; i++{
        chain := []int{}
        length := buildChain(i, &chain, &values)
        values[i] = length
        if length == 60{
            s++
        }
    }
    fmt.Println(s)
}

// returns the length of the build chain
func buildChain(n int, chain *[]int, values *map[int]int) int{
    digits := toDigits(n)
    sum := factSum(&digits)

    // first check the previous chains
    value, containsValue := (*values)[sum]
    if containsValue{
        // current chain + old chain
        return value + len(*chain)+1
    }

    if chainContainsNumber(sum,chain){
        return len(*chain)+1
    } else {
        *chain = append(*chain,sum)
        return buildChain(sum, chain,values)}

}

func chainContainsNumber(num int, arr *[]int) bool{
    for i := range *arr{
        if (*arr)[i] == num{
            return true
        }
    }
    return false
}

func factSum(digits *[]int) int{
    sum := 0
    for i := range *digits{
        factorial := fact((*digits)[i])
        sum += factorial
    }
    return sum
}

func toDigits(n int) []int{
    digits := []int{}
    for n != 0{
        lastDigit := math.Mod(float64(n),10)
        digits = append(digits,int(lastDigit))
        n = n / 10
    }
    return digits
}

func fact(n int) int{
    if n == 0{
        return 1
    }
    return n * fact(n -1)
}
