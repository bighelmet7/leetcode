package median_sorted_arrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	sort.Ints(nums)
	median := len(nums) / 2
	if len(nums)%2 != 0 {
		return float64(nums[median])
	}
	return (float64(nums[median-1]) + float64(nums[median])) / 2.0
}
