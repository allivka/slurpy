package main

import (
	"fmt"
	"github.com/allivka/slurpy/pkg/formatter"
)

func main() {

	result, err := formatter.BasicFormat(`1 +  2 + 
	3 + 4 \
	+ 4    + 5 +6
	`)
	if err == nil {
		fmt.Println(result)
		return
	}

	fmt.Println(err)

}
