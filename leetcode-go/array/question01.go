package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		tmp := target - v
		if j, ok := m[tmp]; ok {
			return []int{i, j}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	x := 0
	for _, v := range nums {

		if nums[x] != v {
			x++
			nums[x] = v
		}
	}
	return x + 1
}

func removeElement(nums []int, val int) int {
	slow := 0
	for _, v := range nums {
		if v != val {
			nums[slow] = v
			slow++
		}
	}
	return slow
}
func plusOne(digits []int) []int {
	n := len(digits)
	ten := false
	ans := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		tmp := digits[i]
		if i == n-1 {
			tmp++
		}
		if ten {
			tmp++
		}
		if tmp >= 10 {
			ans[i+1] = tmp - 10
			ten = true
		} else {
			ten = false
			ans[i+1] = tmp
		}
	}
	if ten {
		ans[0] = 1
		return ans
	}
	return ans[1:]
}
