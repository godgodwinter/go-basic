// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
package main

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	arr_hay:=strings.Split(haystack,"");
  // arr_needle:=strings.Split(needle,"");
  result:=0;
   var isExists = strings.Contains(haystack, needle)
   if isExists {
	  for i,_ :=range arr_hay{
		newStr := haystack[i:]
		  if strings.HasPrefix(newStr,needle) {
			  return i;
		  }
	  }
	  return result;
   }else{
	 return -1;
   }
}


func main() {
	haystack:="sadbutsad";
	needle:="sad";

	strStr(haystack,needle);
}
