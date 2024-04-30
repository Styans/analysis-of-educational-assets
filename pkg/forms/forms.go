package forms

import "errors"

func CheckForm(args ...interface{}) error {
	for _, v := range args {
		if v == "" {
			return errors.New("Заполните Форму")
		}
	}
	return nil
}
