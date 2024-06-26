package main
import (
	"fmt"
)

func plusOne(digits []int) []int {
    lenDigit := len(digits)


    for i:=lenDigit-1; i>=0; i-- {
        if digits[i] < 9 {
            digits[i] += 1
            return digits
        } else {
            digits[i] = 0
        }
    }

    digits = append([]int{1}, digits...)
    return digits

}

func main() {
    digits := []int{9, 1, 9}
    fmt.Println(plusOne(digits))
}
