package Array

func TwoSum(nums []int, target int) []int {
	dic := map[int]int{}
	for k, v := range nums {
		_, ok := dic[target-v]
		if ok {
			return []int{dic[target-v], k}
		}
		dic[v] = k
	}
	return []int{}
}
