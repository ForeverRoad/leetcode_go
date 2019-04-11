package easy

func main() {
	A:= []int {1,19,4,5,6,3}
	println(peakIndexInMountainArray(A))
}

/**
Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].
*/
func peakIndexInMountainArray(A []int) int {
	//抓住了 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]的特性
	lo := 0
	hi := len(A) - 1
	for lo < hi {
		mid := (lo + hi) / 2
		if A[mid] < A[mid + 1] { // peak index is after mid.
			lo = mid + 1
		}else if A[mid -1] > A[mid] { // peak index is before mid.
			hi = mid
		}else { // peak index is mid.
			return mid
		}
	}
	return -1 // no peak.
	//for i, v := range A {
	//	if i > 0 && i < len(A)-1 && v > A[i-1] && v > A[i+1] {
	//		return i
	//	}
	//}
	//return -1
}
