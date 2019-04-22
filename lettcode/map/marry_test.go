package action

import (
	"testing"
)

func TestMoveZeroParam(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 0, 3, 0, 12, 9)
	moveZeroes(nums)
}
func TestFindDoubleParam(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 0, 3, 2, 3, 9)
	resp := findDuplicate(nums)
	t.Log(resp)
}
func TestPattern(t *testing.T) {
	b := wordPattern("abba", "c v v c")
	t.Log(b)
}
func BenchmarkLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
func TestStreamMid(t *testing.T) {
	list := Constructor()
	list.AddNum(1)
	list.AddNum(2)
	list.AddNum(3)
	mid := list.FindMedian()
	t.Log(mid)
}
func TestArry(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 0, 2, 2, 3)
	reslut := removeElement(nums, 2)
	t.Log(reslut)
}
func TestStrLen(t *testing.T){

	r:=lengthOfLongestSubstring("abcdab")
	t.Log(r)
}