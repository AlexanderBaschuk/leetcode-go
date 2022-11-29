package twoSumPkg

func TwoSum(nums []int, target int) []int {
	knownValues := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		pairValue := target - nums[i]
		pairIndex, ok := knownValues[pairValue]
		if ok {
			result := []int{i, pairIndex}
			return result
		}
		knownValues[nums[i]] = i
	}
	return []int{}
}
