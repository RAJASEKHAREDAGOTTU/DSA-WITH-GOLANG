// 1. permutation in go

package main

import "fmt"

func main() {
    sample := "AbS@"
    sampleRune := []rune(sample)
    generatePermutation(sampleRune, 0, len(sampleRune)-1)
}

func generatePermutation(sampleRune []rune, left, right int) {
    if left == right {
        fmt.Println(string(sampleRune))
    } else {
        for i := left; i <= right; i++ {
            sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
            generatePermutation(sampleRune, left+1, right)
            sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
        }
    }
}

$go run main.go
AbS@
Ab@S
ASb@
AS@b
A@Sb
A@bS
bAS@
bA@S
bSA@
bS@A
b@SA
b@AS
SbA@
Sb@A
SAb@
SA@b
S@Ab
S@bA
@bSA
@bAS
@SbA
@SAb
@ASb
@AbS