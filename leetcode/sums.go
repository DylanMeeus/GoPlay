
package main

import "fmt"

func main() {
    values := []int{1,2,4}
    fmt.Println(twoSums(values, 6))
}


func twoSums(nums []int, target int) []int {
    indices := make([]int, 2)
    indices[0] = 1
    for i, v := range(nums) {
        for j, v2 := range(nums) { 
            if i != j {
                if v + v2 == target{
                    return []int{i, j}
                }
            }
        }
    }
    return nil
}
