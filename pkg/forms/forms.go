package forms

import (
	"errors"
	"strconv"
)

func CheckForm(args ...interface{}) error {
	for _, v := range args {
		if v == "" {
			return errors.New("Заполните Форму")
		}
	}
	return nil
}

func IsInt(s string) (int, error) {
	return strconv.Atoi(s)
}
