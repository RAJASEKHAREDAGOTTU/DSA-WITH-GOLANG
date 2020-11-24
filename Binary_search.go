// 1. Binary Search

package main
import "fmt"
 
func binarySearch(needle int, haystack []int) bool {
 
    low := 0
    high := len(haystack) - 1
 
    for low <= high{
        median := (low + high) / 2
 
        if haystack[median] < needle {
            low = median + 1
        }else{
            high = median - 1
        }
    }
 
    if low == len(haystack) || haystack[low] != needle {
        return false
    }
 
    return true
}
 
 
func main(){
    items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
    fmt.Println(binarySearch(70, items))
}
$go run main.go
true

// 2. Binary search

package main

import "fmt"

func main() {
    searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
    searchNumber := 23

    fmt.Println("Running Program")
    fmt.Println("Searching list of numbers: ", searchField)
    fmt.Println("Searching for number: ", searchNumber)

    numFound := false
    
    result, _ := binarySearch2(searchField, len(searchField), searchNumber, numFound)
    fmt.Println("Found! Your number is found in position: ", result)
    
}

func binarySearch2(a []int, field int, search int, numFound bool) (result int, searchCount int) {
    
    searchCount, i := 0, 0
    for !numFound {
        searchCount++
        mid := i + (field-i)/2
        if search == a[mid] {
            numFound = true
            result = mid
            return result, searchCount
        } else if search > a[mid] {
            field++
            //i = mid + 1 causes a stack overflow
            return binarySearch2(a, field, search, numFound)
        }
        field = mid
        return binarySearch2(a, field, search, numFound)
    }
    return result, searchCount
}

$go run main.go
Running Program
Searching list of numbers:  [2 5 8 12 16 23 38 56 72 91]
Searching for number:  16
Found! Your number is found in position:  4
