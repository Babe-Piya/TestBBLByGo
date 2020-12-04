package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)


func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Target : ")
	scanner.Scan()
	target := scanner.Text()


	fmt.Print("Mask : ")
	scanner.Scan()
	mask := scanner.Text()

	if len(target)>50 || len(mask)>50 {
		fmt.Print("text longer 50 ")
		return
	}

	if len(target) == len(mask) {

		checkRune(target ,mask )

	} else if len(target) > len(mask){
		for i:=0;i<len(target)-len(mask);i++{
			mask += " "
		}

		checkRune(target ,mask )

	}else if len(target) < len(mask){
		for i:=0;i<len(mask)-len(target);i++{
			target += " "
		}

		checkRune(target ,mask )

	}
	

}

func checkRune(target string,mask string) {
	for k,r := range target {
		for kM,rM := range mask {
			if k == kM {
				if unicode.IsUpper(r) && unicode.IsUpper(rM){
					fmt.Print(string(r))
					}else if unicode.IsLower(r) && unicode.IsLower(rM) {
						fmt.Print(string(rM))
					}else if unicode.IsLower(r) && unicode.IsUpper(rM) {
						fmt.Print("$")
					}else if unicode.IsUpper(r) && unicode.IsLower(rM) {
						fmt.Print("$")
					} else {
						fmt.Print("#")
					}
			}
		}
		
	}
}