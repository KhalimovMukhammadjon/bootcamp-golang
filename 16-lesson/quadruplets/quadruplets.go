package quadruplets

import "fmt"

func FindSum() {
	nums := []int{1, 2, 3, 6}
	//count := 0

	//nums[d] = nums[a] + nums[b] + nums[c]
	for a := 0; a < len(nums); a++ {
		fmt.Println(nums[a])
		// for b := 0; b < len(nums); b++ {
		// 	b = nums[a]+1
		// }
	}
}
