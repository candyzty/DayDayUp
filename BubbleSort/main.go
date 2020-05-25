package main

import "fmt"

func BubbleSort(arr []int)[]int{

	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1-i;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	return arr

}




func main(){
	arr := []int{44, 38, 65, 97, 76, 14, 27, 48}
	fmt.Println(BubbleSort(arr))
}
