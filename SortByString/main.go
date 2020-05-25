package main

import (
	"fmt"
)

func SelectSort(arr []string)[]string{

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

func main(){
	arr :=[]string {"c","q","a","b","h","y","z","u"}
	fmt.Println(SelectSort(arr))
}

