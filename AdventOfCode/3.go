package main

import(
    "fmt"
    "math"
    )

func main(){
    distance := getDistance(277678)
    fmt.Println(int(distance))
}

func getDistance(n int) float64 {
    x := getX(n)
    y := getY(n)

    // we need to find the path down to X = 0, Y = 0 (sum of coordinates)
    fmt.Print("x: ")
    fmt.Println(x)
    fmt.Print("y: ")
    fmt.Println(y)
    return x + y
}

func getX(n int) float64{
    if n == 1{
    	return 0.0
    }

    result := getX(n-1)
    return result + math.Sin(float64((int64(math.Floor(math.Sqrt(float64(4*(n-2)+1)))) % 4)) * (math.Pi/2))

}

func getY(n int) float64{
    if n == 1{
    	return 0.0
    }

    result := getY(n-1)
    return result + math.Cos(float64((int64(math.Floor(math.Sqrt(float64(4*(n-2)+1)))) % 4)) * (math.Pi/2))
}
