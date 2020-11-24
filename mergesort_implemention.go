// 1. merge sort with go routine

package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {
 
    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice), "\n")
}
 
// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
 
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}
  
func mergeSort(items []int) []int {
    var num = len(items)
      
    if num == 1 {
        return items
    }
      
    middle := int(num / 2)
    var (
        left = make([]int, middle)
        right = make([]int, num-middle)
    )
    for i := 0; i < num; i++ {
        if i < middle {
            left[i] = items[i]
        } else {
            right[i-middle] = items[i]
        }
    }
      
    return merge(mergeSort(left), mergeSort(right))
}
  
func merge(left, right []int) (result []int) {
    result = make([]int, len(left) + len(right))
      
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
      
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
      
    return
}

$go run main.go

--- Unsorted --- 

 [-86 -816 -205 86 -106 -624 357 189 -263 149 419 125 -277 -72 235 117 -361 -321 528 -170]

--- Sorted ---

 [-816 -624 -361 -321 -277 -263 -205 -170 -106 -86 -72 86 117 125 149 189 235 357 419 528] 