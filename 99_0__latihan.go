package main

import "fmt"

func main() {

	var num1 int8=30;
	var num2 int8=30;
	result:= num1+num2;
	fmt.Println(result);

	
	indexMap := make(map[int]int)
	indexMap[0]=0;
	indexMap[1]=10;
	indexMap[2]=20;
	indexMap[3]=30;
	indexMap[30]=300;
	fmt.Println(indexMap)
	
	nums := []int{2,7,11,15}
	fmt.Println(nums)

}