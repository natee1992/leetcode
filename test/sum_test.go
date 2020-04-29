package test

import (
	"leetcode/acc"
	"math/rand"
	"testing"
)

const N = 1000

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		acc.TwoSum1(nums, 9)
	}
}

func BenchmarkFindSumNumIndexOp(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		acc.FindSumNumIndexOp(nums, 9)
	}
}
