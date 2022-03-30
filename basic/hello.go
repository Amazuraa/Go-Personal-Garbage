package main

import (
	"fmt"
	// "strconv"
)

var (
	home   = "-"
	user   = "-"
	gopath = "z"
)

func main() {
	var ii = 0
	
	for i := 0; i <= 6; i++ {
		if i < 4 {
			ii+=2
			for j := 0; j < ii; j++ {
				fmt.Print(home)
			}

			test := (map[bool]string{true: " Go Looping Test + Ternary Op", false: ""})[ii == 8]
			fmt.Println(test)
		} else {
			ii-=2
			for j := 0; j < ii; j++ {
				fmt.Print(user)
			}
			fmt.Println("")
		}
	}
}