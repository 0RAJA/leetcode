package p0001

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for index, num := range nums {
		if matchIndex, ok := seen[target-num]; ok {
			return []int{matchIndex, index}
		}
		seen[num] = index
	}
	return nil
}
