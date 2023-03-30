// &&
// ||
// !
package main

import "fmt"

func operasi_bool() {
	var nilai1 bool = true
	var nilai2 bool = true
	var nilai3 bool = false
	var result = nilai1 && nilai2
	var result2 = nilai1 && nilai3
	
	fmt.Println(result,result2)
}
