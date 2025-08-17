package main

import (
	"fmt"
	"github.com/allivka/slurpy/pkg/formatter"
)

func main() {
	
	src := `1 +  2 + 
	3 + 4 \
	
	+ 4    + 5 +6
	`
	
	result, err := formatter.Format(src)
	if err == nil {
		fmt.Println(result)
		return
	}

	fmt.Println(err)

}
