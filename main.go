package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} //78
	numGorut := 3

	var totalSum int 

	ch := make(chan int, numGorut)

	partSize := len(arr) / numGorut

	for i := 0; i < numGorut; i++ {

		startIdx := i*partSize
		endIdx := startIdx + partSize

		go getSum(ch, arr[startIdx:endIdx])
	}

	for i:=0; i < numGorut; i++{
		totalSum += <-ch
	}

	fmt.Println("Сумма:", totalSum)
}

func getSum(sumCh chan int, arr []int) {
	var res int
	for _, data := range arr {
		res += data
	}
	sumCh <- res
}