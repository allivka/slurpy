package tokens

import (
	"strconv"
)

type UncreatableToken struct {}
func(UncreatableToken) NewFromWord(string) error{return nil}

type StringName struct {
	value string
}

func(s *StringName) NewFromWord(word string) error {
	s.value = word
	
	return nil
}

type IntLiteral struct {
	value int64
}

func(i *IntLiteral) NewFromWord(word string) error {
	temp, err := strconv.ParseInt(word, 10, 64)
	if err != nil {
		return err
	}
	i.value = temp
	
	return nil
}
