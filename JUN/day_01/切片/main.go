package main

import "fmt"

// slice 可以向后扩展，不可以向前扩展
// s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)

func updateSlices(s []int) {
	s[0] = 100

}

func slices() {
	arr := [...]int{2, 3, 4, 5, 6, 7, 6, 7, 4, 3}
	s := arr[2:6]
	fmt.Println("arr[2:6]=", s[2:6])
	fmt.Println("arr[:6]=", s[:6])
	fmt.Println("arr[2:]=", s[2:])
	fmt.Println("arr[:]=", s[:])
	s1 := arr[2:]
	fmt.Println("s1", s1)
	s2 := arr[:]
	fmt.Println("arr[:]", s2)
	fmt.Println("After updateSlice(s1)")
	updateSlices(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("After updateSlice(s2)")
	updateSlices(s2)
	fmt.Println(s2)
	fmt.Println(arr)
	fmt.Println("Extending slice")
	arr[0], arr[2] = 2, 4
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Println(s2[3:5])
}
func main() {
	slices()
}
