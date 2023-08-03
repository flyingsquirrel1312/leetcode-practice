package main

// https://leetcode.com/problems/3sum/
import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	m := make(map[string]struct{})
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				key := fmt.Sprintf("%d:%d:%d", nums[i], nums[j], nums[k])
				if _, ok := m[key]; !ok {
					m[key] = struct{}{}
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			}
		}
	}
	return res
}

func main() {
	var nums = []int{0, 0, 0}
	res := threeSum(nums)
	for _, probe := range res {
		fmt.Printf("%d %d %d\n", probe[0], probe[1], probe[2])
	}
}
