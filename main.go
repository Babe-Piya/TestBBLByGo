package main

import (
	"fmt"
	"strconv"
)


func main() {
	var day string
	var month string
	var year string 
	fmt.Println("Enter day mounth year : ")
	fmt.Scan(&day,&month,&year)
	dayChanged := changeToNum(day)
	if dayChanged == 0 {
		fmt.Println("wrong Day")
		return
	}
	mouthChanged := changeToNum(month)
	if mouthChanged == 0 {
		fmt.Println("wrong Month")
		return
	}
	yearChanged := changeToNum(year)
	if yearChanged == 0 {
		fmt.Println("wrong Year")
		return
	}
	fmt.Printf("%v/%v/%v",mouthChanged,dayChanged,yearChanged)
}

func changeToNum(str string) (int64){
	if str == "January" {
		return 1
	}else if str == "February" {
		return 2
	}else if str == "March" {
		return 3
	}else if str == "April" {
		return 4
	}else if str == "May" {
		return 5
	}else if str == "June" {
		return 6
	}else if str == "July" {
		return 7
	}else if str == "August" {
		return 8
	}else if str == "September" {
		return 9
	}else if str == "October" {
		return 10
	}else if str == "November" {
		return 11
	}else if str == "December" {
		return 12
	} else {
		r,err := strconv.ParseInt(str,10,64)

		if err != nil {
			return 0
		}

		return r
	}
}