package easy

func main() {
	result := selfDividingNumbers(1, 128)
	for _, v := range result {
		println(v)
	}
}
func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for ; left <= right; left++ {
		temp := left
		fla := true
		for ; temp > 0; temp /= 10 {
			if temp%10 == 0 || left%(temp%10) != 0 {
				fla = false
				break
			}
		}
		if fla {
			result = append(result, left)
		}

	}
	return result
}
