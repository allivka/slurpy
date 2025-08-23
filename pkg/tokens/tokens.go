package tokens

import (
	"strconv"
	"fmt"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	wp "github.com/allivka/slurpy/pkg/words"
)


type Identificator struct {
	bts.BasicToken
}

func(s Identificator) NewFromWord(word string) (bts.Token, error) {
	
	wt, err := wp.GetWordType(word);
	
	if err == nil && wt == wp.Identificator {
		t, _ := bts.BasicToken{}.NewFromWord(word)
		return Identificator{ BasicToken: t.(bts.BasicToken) }, nil
	}
	
	return Identificator{}, fmt.Errorf("Failed creating new identificator token from word '%s': %w", word, err)
}

type Integer struct {
	bts.BasicToken
	value int64
}

func(i Integer) NewFromWord(word string) (bts.Token, error) {
	temp, err := strconv.ParseInt(word, 10, 64)
	if err != nil {
		return Integer{}, fmt.Errorf("Failed creating new integer token from word '%s': %w", word, err)
	}
	
	t, _ := bts.BasicToken{}.NewFromWord(word)

	return Integer {
		value: temp,
		BasicToken: t.(bts.BasicToken),
	}, nil
}

type Float struct {
	bts.BasicToken
	value float64
}

func(f Float) NewFromWord(word string) (bts.Token, error) {
	temp, err := strconv.ParseFloat(word, 64)
	if err != nil {
		return Float{}, fmt.Errorf("Failed creating new float token from word '%s': %w", word, err)
	}
	
	t, _ := bts.BasicToken{}.NewFromWord(word)
	
	return Float {
		value: temp,
		BasicToken: t.(bts.BasicToken),
	}, nil
}


