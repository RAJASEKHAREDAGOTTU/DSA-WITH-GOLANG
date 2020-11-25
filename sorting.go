// 1. Bubble Sort in Golang

//example 1 :

package main

import "fmt"

func main() {
    sample := []int{3, 4, 5, 2, 1}
    bubbleSort(sample)
    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    bubbleSort(sample)
}

func bubbleSort(arr []int) {
    len := len(arr)
    for i := 0; i < len-1; i++ {
        for j := 0; j < len-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println("\nAfter Bubble Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}

$go run main.go
After Bubble Sorting
1
2
3
4
5

After Bubble Sorting
-3
-1
1
2
3
4
5
7
8

// example 2:

package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {
 
    slice := generateSlice(10)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    bubblesort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
 func generateSlice(size int) []int {
 
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(100) - rand.Intn(100)
    }
    return slice
}
  
func bubblesort(items []int) {
    var (
        n = len(items)
        sorted = false
    )
    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
                items[i+1], items[i] = items[i], items[i+1]
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }
}

$go run main.go

--- Unsorted --- 

 [4 54 22 56 -45 72 -33 -40 29 26]

--- Sorted ---

 [-45 -40 -33 4 22 26 29 54 56 72] 

// example 3:

package main

import (
    "fmt"
)

var toBeSorted [10]int = [10]int{1,3,2,4,8,6,7,2,3,0}

func bubbleSort(input [10]int) {
    
    n := 10
   
    swapped := true
 
    for swapped {
        
        swapped = false
        
        for i := 1; i < n; i++ {
            
            if input[i-1] > input[i] {
                
                fmt.Println("Swapping")
               
                input[i], input[i-1] = input[i-1], input[i]
               
                swapped = true
            }
        }
    }
 fmt.Println(input)
}

func main() {
    fmt.Println("Hello World")
    bubbleSort(toBeSorted)
}

$go run main.go

Hello World
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
Swapping
[0 1 2 2 3 3 4 6 7 8]

// 2.Insertion sort

// example 1:
package main

import "fmt"

func main() {
    sample := []int{3, 4, 5, 2, 1}
    insertionSort(sample)

    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    insertionSort(sample)
}

func insertionSort(arr []int) {
    len := len(arr)
    for i := 1; i < len; i++ {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    
    fmt.Println("After Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}

$go run main.go
After Sorting
1
2
3
4
5
After Sorting
-3
-1
1
2
3
4
5
7
8

example 2 :

package main

import "fmt"

func main() {
    sample := []int{3, 4, 5, 2, 1,9,0,-9,-6,7}
    insertionSort(sample)

    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    insertionSort(sample)
}

func insertionSort(arr []int) {
    len := len(arr)
    for i := 1; i < len; i++ {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    
    fmt.Println("After Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}

$go run main.go
After Sorting
-9
-6
0
1
2
3
4
5
7
9
After Sorting
-3
-1
1
2
3
4
5
7
8

// 3. Selection sort :

// example 1 :
package main

import "fmt"

func main() {
    sample := []int{3, 4, 5, 2, 1}
    selectionSort(sample)
    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    selectionSort(sample)
}

func selectionSort(arr []int) {
    len := len(arr)
    for i := 0; i < len-1; i++ {
        minIndex := i
        for j := i + 1; j < len; j++ {
            if arr[j] < arr[minIndex] {
                arr[j], arr[minIndex] = arr[minIndex], arr[j]
            }
        }
    }
    fmt.Println("\nAfter SelectionSort")
    for _, val := range arr {
        fmt.Println(val)
    }
}

$go run main.go
After Sorting
1
2
3
4
5
After Sorting
-3
-1
1
2
3
4
5
7
8

example 2:

package main

import "fmt"

func main() {
    sample := []int{33,21,10,49,50}
    selectionSort(sample)
    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3,90,78,567,89}
    selectionSort(sample)
}

func selectionSort(arr []int) {
    len := len(arr)
    for i := 0; i < len-1; i++ {
        minIndex := i
        for j := i + 1; j < len; j++ {
            if arr[j] < arr[minIndex] {
                arr[j], arr[minIndex] = arr[minIndex], arr[j]
            }
        }
    }
    fmt.Println("\nAfter SelectionSort")
    for _, val := range arr {
        fmt.Println(val)
    }
}
$go run main.go

After SelectionSort
10
21
33
49
50

After SelectionSort
-3
-1
1
2
3
4
5
7
8
78
89
90
567

// 4.  Quick sort 

// example 1 :

package main

import (
	"fmt"
	"strconv"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		var pivot = partition(arr, low, high)
		quickSort(arr, low, pivot)
		quickSort(arr, pivot + 1, high)
	}
}

func partition(arr []int, low, high int) int {
	var pivot = arr[low]
	var i = low
	var j = high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++;
		}
		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot

	return j
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println("")
}

func main() {
	var arr = []int { 18, 3, 11, 6, -9, 1, 0 , 5 , 90 , 56 , 7 , 8 }

	fmt.Print("Before Sorting: ")
	printArray(arr)

	quickSort(arr, 0, len(arr) - 1)
	fmt.Print("After Sorting: ")
	printArray(arr)
}

$go run main.go
Before Sorting: 18 3 11 6 -9 1 0 5 90 56 7 8 
After Sorting: -9 0 1 3 5 6 7 8 11 18 56 90 

// example 2 :

package main

import (
	"fmt"
)

func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 41}
	
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

$go run main.go
[16 20 41 43 46 55 59 74 90]

// 5. merge sort

// example 1 :

package main

import (
    "fmt"
)

func main() {
    A := []int{3, 5, 1, 6, 1, 7, 2, 4, 5,9,10,-9,-7}
    fmt.Println(sort(A))
}
func sort(A []int) []int {
    if len(A) <= 1 {
        return A
    }

    left, right := split(A)
    left = sort(left)
    right = sort(right)
    return merge(left, right)
}

// split array into two
func split(A []int) ([]int, []int) {
    return A[0:len(A) / 2], A[len(A) / 2:]
}

// assumes that A and B are sorted
func merge(A, B []int) []int {
    arr := make([]int, len(A) + len(B))

    // index j for A, k for B
    j, k := 0, 0

    for i := 0; i < len(arr); i++ {
        
        if j >= len(A) {
            arr[i] = B[k]
            k++
            continue
        } else if k >= len(B) {
            arr[i] = A[j]
            j++
            continue
        }
        
        if A[j] > B[k] {
            arr[i] = B[k]
            k++
        } else {
            arr[i] = A[j]
            j++
        }
    }

    return arr
}

$go run main.go
[-9 -7 1 1 2 3 4 5 5 6 7 9 10]

// example 2 :

package main

import (
    "fmt"
)

func main() {
    A := []int{20,30,10,90,4,7,3,9,10,-9,-7}
    fmt.Println(sort(A))
}


func sort(A []int) []int {
    if len(A) <= 1 {
        return A
    }

    left, right := split(A)
    left = sort(left)
    right = sort(right)
    return merge(left, right)
}

// split array into two
func split(A []int) ([]int, []int) {
    return A[0:len(A) / 2], A[len(A) / 2:]
}


func merge(A, B []int) []int {
    arr := make([]int, len(A) + len(B))

    // index j for A, k for B
    j, k := 0, 0

    for i := 0; i < len(arr); i++ {
        
        if j >= len(A) {
            arr[i] = B[k]
            k++
            continue
        } else if k >= len(B) {
            arr[i] = A[j]
            j++
            continue
        }
       
        if A[j] > B[k] {
            arr[i] = B[k]
            k++
        } else {
            arr[i] = A[j]
            j++
        }
    }

    return arr
}

$go run main.go
[-9 -7 3 4 7 9 10 10 20 30 90]

// 6. Heap sort

// example 1 :

package main

import "fmt"

type minheap struct {
    arr []int
}

func newMinHeap(arr []int) *minheap {
    minheap := &minheap{
        arr: arr,
    }
    return minheap
}

func (m *minheap) leftchildIndex(index int) int {
    return 2*index + 1
}

func (m *minheap) rightchildIndex(index int) int {
    return 2*index + 2
}

func (m *minheap) swap(first, second int) {
    temp := m.arr[first]
    m.arr[first] = m.arr[second]
    m.arr[second] = temp
}

func (m *minheap) leaf(index int, size int) bool {
    if index >= (size/2) && index <= size {
        return true
    }
    return false
}

func (m *minheap) downHeapify(current int, size int) {
    if m.leaf(current, size) {
        return
    }
    smallest := current
    leftChildIndex := m.leftchildIndex(current)
    rightRightIndex := m.rightchildIndex(current)
    if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
        smallest = leftChildIndex
    }
    if rightRightIndex < size && m.arr[rightRightIndex] < m.arr[smallest] {
        smallest = rightRightIndex
    }
    if smallest != current {
        m.swap(current, smallest)
        m.downHeapify(smallest, size)
    }
    return
}

func (m *minheap) buildMinHeap(size int) {
    for index := ((size / 2) - 1); index >= 0; index-- {
        m.downHeapify(index, size)
    }
}

func (m *minheap) sort(size int) {
    m.buildMinHeap(size)
    for i := size - 1; i > 0; i-- {
        // Move current root to end
        m.swap(0, i)
        m.downHeapify(0, i)
    }
}

func (m *minheap) print() {
    for _, val := range m.arr {
        fmt.Println(val)
    }
}

func main() {
    inputArray := []int{6, 5, 3, 7, 2, 8, -1}
    minHeap := newMinHeap(inputArray)
    minHeap.sort(len(inputArray))
    minHeap.print()
    fmt.Scanln()
}

$go run main.go
8
7
6
5
3
2
-1

// example 2 :

package main

import "fmt"

type minheap struct {
    arr []int
}

func newMinHeap(arr []int) *minheap {
    minheap := &minheap{
        arr: arr,
    }
    return minheap
}

func (m *minheap) leftchildIndex(index int) int {
    return 2*index + 1
}

func (m *minheap) rightchildIndex(index int) int {
    return 2*index + 2
}

func (m *minheap) swap(first, second int) {
    temp := m.arr[first]
    m.arr[first] = m.arr[second]
    m.arr[second] = temp
}

func (m *minheap) leaf(index int, size int) bool {
    if index >= (size/2) && index <= size {
        return true
    }
    return false
}

func (m *minheap) downHeapify(current int, size int) {
    if m.leaf(current, size) {
        return
    }
    smallest := current
    leftChildIndex := m.leftchildIndex(current)
    rightRightIndex := m.rightchildIndex(current)
    if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
        smallest = leftChildIndex
    }
    if rightRightIndex < size && m.arr[rightRightIndex] < m.arr[smallest] {
        smallest = rightRightIndex
    }
    if smallest != current {
        m.swap(current, smallest)
        m.downHeapify(smallest, size)
    }
    return
}

func (m *minheap) buildMinHeap(size int) {
    for index := ((size / 2) - 1); index >= 0; index-- {
        m.downHeapify(index, size)
    }
}

func (m *minheap) sort(size int) {
    m.buildMinHeap(size)
    for i := size - 1; i > 0; i-- {
        // Move current root to end
        m.swap(0, i)
        m.downHeapify(0, i)
    }
}

func (m *minheap) print() {
    for _, val := range m.arr {
        fmt.Println(val)
    }
}

func main() {
    inputArray := []int{2,6,7,5,0,-2}
    minHeap := newMinHeap(inputArray)
    minHeap.sort(len(inputArray))
    minHeap.print()
    fmt.Scanln()
}

$go run main.go
7
6
5
2
0
-2
