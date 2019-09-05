package main

import (
	"errors"
	"fmt"
)

func Unpack(s string) (string, error) {
	var res []rune
	runes := []rune(s)
	var n int
	var mem rune

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if c > 48 && c < 58 {
			if n == 0 {
				if i == 0 {
					return "", errors.New("String is incorrect")
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

func main() {
	fmt.Println(Unpack("a4bc2d11e"))
}
