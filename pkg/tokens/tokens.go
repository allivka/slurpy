package tokens

import (
	"strconv"
	"fmt"
	wp "github.com/allivka/slurpy/pkg/words"
)



type Identificator struct {
	value string
}

func(s *Identificator) NewFromWord(word string) error {
	
	wt, err := wp.GetWordType(word);
	
	if err == nil && wt == wp.Identificator {
		s.value = word
		return nil
	}
	
	return fmt.Errorf("Failed creating new Identificator token from word '%s': %w", word, err)
}

type IntLiteral struct {
	value int64
}

func(i *IntLiteral) NewFromWord(word string) error {
	temp, err := strconv.ParseInt(word, 10, 64)
	if err != nil {
		return fmt.Errorf("Failed creating new Identificator token from word '%s': %w", word, err)
	}
	i.value = temp
	
	return nil
}
