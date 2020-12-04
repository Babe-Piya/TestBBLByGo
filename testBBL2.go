package main

import "fmt"

func main() {
	 ary :=[6]int{5,18,2,14,10,9}
	var newary [6]int
	 for i:=0;i<len(ary);i++{
		fmt.Print(" ",ary[i])
		newary[i] = ary[len(ary)-i-1] 	
		
	 }
	 fmt.Println()
	 for j:=0;j<len(newary);j++{
		ary[j]=newary[j]
		fmt.Print(" ",ary[j])
	 }
	
}