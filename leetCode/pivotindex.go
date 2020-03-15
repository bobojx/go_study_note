package main

import "fmt"

// 寻找数组的中心索引
func pivotIndex(nums []int) int {
	leftSum, rightSum := 0, 0
	for _, v := range nums {
		rightSum += v
	}
	for i, v := range nums {
		rightSum -= v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}
	return -1
}

func main() {
	//nums := []int{1, 3, 7, 6, 5, 6}
	nums := []int{-1, -1, -1, -1, 0, 1}
	i := pivotIndex(nums)
	fmt.Println(i)
}
