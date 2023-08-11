package four_sum

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	memo := make(map[string]any)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			left, right := 0, len(nums)-1

			for left < right {
				if i == left || j == left {
					left++
					continue
				} else if i == right || j == right {
					right--
					continue
				}

				// c + d = target - a - b
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					// we need to make it smaller
					right--
				} else if sum < target {
					// we need to make it bigger
					left++
				} else {
					r := []int{nums[i], nums[j], nums[left], nums[right]}
					sort.Ints(r)
					key := fmt.Sprintf("%v", r)
					if _, ok := memo[key]; !ok {
						result = append(result, r)
						memo[key] = nil
					}
                    left++
				}
			}
		}
	}

	return result
}
