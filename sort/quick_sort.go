package sort

func QuickSort(nums []int, l, r int) {
	if l > r {
		return
	}
	i, j, x := l, r, nums[l]
	for i < j {
		for i < j && nums[j] > x {
			j--
		}
		for i < j && nums[i] <= x {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[l] = nums[l], nums[i]
	QuickSort(nums, l, i-1)
	QuickSort(nums, i+1, r)
}