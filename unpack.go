package main

import (
	"errors"
	"unicode"
)

func Unpack(s string) (string, error) {
	var res []rune
	runes := []rune(s)
	var n int
	var mem rune

	for i, c := range runes {
		if unicode.IsNumber(c) {
			if n == 0 {
				if i == 0 {
					return "", errors.New("string is incorrect")
				}
				mem = runes[i-1]
			}
			n = n*10 + int(c-'0')
		} else {
			if n != 0 {
				for i := 0; i < n-1; i++ {
					res = append(res, mem)
				}

				n = 0
			}
			res = append(res, c)

		}
	}

	for i := 0; i < n-1; i++ {
		res = append(res, mem)
	}

	return string(res), nil
}
