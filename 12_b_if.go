package main

import "fmt"
//! https://leetcode.com/problems/two-sum/description/
// !two-sum leetcode
func main() {
	nums := []int{2,7,11,15}
	var result []int=twoSum(nums,22);
	fmt.Println(result);

}


func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for index, data := range nums {
		if requiredIdx, isPresent := indexMap[target-data]; isPresent {
			fmt.Println(requiredIdx,index,isPresent,target,data)
			fmt.Println("map:",indexMap[target-data])
			return []int{requiredIdx, index} //!return data ditemukan dan break perulangan.
		}
		indexMap[data] = index
	}
	return []int{} //return map kosong ketika data tidak ditemukan
}