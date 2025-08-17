package main

import (
	"fmt"
	"github.com/allivka/slurpy/pkg/formatter"
)

func main() {

	src := `11++  22 + 
	3 + 4 \
	
	+ 4    + 55 +66
	`

	result, err := formatter.WordsFromString(src)
	if err == nil {
		for _, v := range result {
			fmt.Println(v)
		}
		return
	}

	fmt.Println(err)

}
