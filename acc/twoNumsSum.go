package acc

import (
	"fmt"
	"time"
)

var Nums = []int{1, 3, 9, 6, 7, 4}

const Target = 5

/*
	给定一个整数数组 nums 和一个目标值 target，在数组中找出和为目标值的两个数在数组中的下标
*/


func FindSumNumIndexOp(nums []int, target int) []int {
	targetMap := make(map[int]int, len(nums))
	for index, num := range nums {
		v := target - num
		if j,ok := targetMap[v];ok{
			return []int{j,index}
		}
		targetMap[num] = index
	}
	return nil
}

func TwoSum1(nums []int, target int) []int {

	n := len(nums)

	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	startTime := time.Now()
	fmt.Println(FindSumNumIndexOp(Nums, Target))
	//fmt.Println(TwoSum1(Nums, Target))
	fmt.Println("time cost = ", time.Since(startTime))
}
