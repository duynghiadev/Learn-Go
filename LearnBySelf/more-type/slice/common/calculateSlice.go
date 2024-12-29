package common

import "fmt"

// Hàm thêm phần tử vào slice
func AddElement(slice []int, element int) []int {
	return append(slice, element)
}

// Hàm sửa phần tử tại vị trí index
func UpdateElement(slice []int, index int, newElement int) []int {
	if index >= 0 && index < len(slice) {
		fmt.Printf("len(slice): %d\n", len(slice))
		slice[index] = newElement
	} else {
		fmt.Println("Index out of range")
	}
	return slice
}

// Hàm xóa phần tử tại vị trí index
func RemoveElement(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		return append(slice[:index], slice[index+1:]...)
	} else {
		fmt.Println("Index out of range")
	}
	return slice
}

// Hàm in slice
func PrintSlice(slice []int) {
	fmt.Println("Current Slice:", slice)
}
