package main


import(
    "fmt"
    "math"
)

type Fib struct{
    sequence []int
}

func main(){

    f := Fib{}
    f.sequence = []int{1,1,1}

    odds := 0
    for i := 3; odds < 2; i+=2{
        // generate more numbers!
        for f.last() < i{
            f.next()
        }
        if !hasDivisor(i,f.sequence[3:]){
            fmt.Println(i)
            odds += 1
        }
    }

    fmt.Println(f.sequence)
}

func hasDivisor(n int, arr []int) bool{
    for i := 0; i < len(arr); i++{
        mod := math.Mod(float64(n),float64(arr[i]))
        if mod == 0{
            return true
        }
    }
    return false
}

func (fib *Fib) next(){
    length := len(fib.sequence)-1
    new := fib.sequence[length] + fib.sequence[length-1] + fib.sequence[length-2]
    fib.sequence = append(fib.sequence, new)
}

func (fib *Fib) last() int{
    return fib.sequence[len(fib.sequence)-1]
}