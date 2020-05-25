package main

import (
	"fmt"
)

//选择排序
func SelectSortByMax(arr []int)int{
	//给定一个数组,判断其长度
	if len(arr)<=1{
		return arr[0]
	}else {
		max :=arr[0]

		for i:=1;i<len(arr);i++{
			if arr[i]>max{
				max=arr[i]
			}
		}
		return max
	}
}
//选择排序
func SelectSortByMin(arr []int)int{
	//给定一个数组,判断其长度
	if len(arr)<=1{
		return arr[0]
	}else {
		min :=arr[0]

		for i:=1;i<len(arr);i++{
			if arr[i]<min{
				min=arr[i]
			}
		}
		return min
	}
}

func SelectSort(arr []int)[]int{

	for i:=0;i<len(arr)-1;i++{
		min :=i
		for j:=i+1;j<len(arr);j++{
			if arr[min]>arr[j]{
				min =j //记录最小值的索引
			}
		}
		if min!=i{
			arr[i],arr[min] = arr[min],arr[i]
		}

	}
	return arr
}

func main() {
	arr := []int{44, 38, 65, 97, 76, 14, 27, 48}
	fmt.Println(SelectSort(arr))
	fmt.Println(SelectSortByMin(arr))
	fmt.Println(SelectSortByMax(arr))
}