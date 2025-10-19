package pkg

import (
	"errors"
	"strings"
)

func CheckLink(link string) error {
	if strings.Contains(link, "https://biggeek.ru/products/") {
		return nil
	}
	return errors.New("не подходящая ссылка")
}
