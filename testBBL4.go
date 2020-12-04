package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)


func main() {
	f,err:=os.Open("product.dat")

	if err!= nil {
		fmt.Print(err)
	}

	s := bufio.NewScanner(f)

    for s.Scan() {
        line := s.Text()

		fmt.Println(line)

    }
}