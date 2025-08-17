package main

import (
	"fmt"
	"github.com/allivka/slurpy/pkg/formatter"
)

func main() {

	src := `11++  22 + 
	3 + 4 \
	//3 -3 323 322234234 32423423 123123 24
	+ 4    + 55 +66
	`

	result, err := formatter.WordsFromSrcString(src)
	if err == nil {
		for _, v := range result {
			fmt.Println(v)
		}
		return
	}

	fmt.Println(err)

}
