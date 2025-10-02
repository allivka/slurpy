package main

import (
	"fmt"
	"github.com/allivka/slurpy/pkg/formatter"
	"github.com/allivka/slurpy/pkg/lexer"
)

func main() {

	src := `11++  22 + --a++b *d1a
	3 + 4 \
	//3 -3 323 322234234 32423423 123123 24
	+ 4    + 55 ++66
	`

	formatted, err := formatter.WordsFromSrcString(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(formatted)

	lexed, err := lexer.Lex(formatted)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range lexed {
		fmt.Printf("%+v\n", token)
	}

}
