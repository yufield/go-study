package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Println(a)
	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	fmt.Println(ss)
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)
	sort.Stable(sort.IntSlice(a))
	fmt.Println(a)
	fmt.Println(sort.IsSorted(sort.IntSlice(a)))
	search := sort.Search(len(a), func(i int) bool { return a[i] >= 3 })
	fmt.Println(search)
	fmt.Println(sort.SearchInts(a, 3))
}
