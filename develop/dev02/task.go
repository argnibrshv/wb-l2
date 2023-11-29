package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const backSlash = '\\'

var ErrInvalidString = errors.New("некорректная строка")

func SbNumConvError(prevErr error) error {
	return fmt.Errorf("ошибка конвертации числа из sbNum: %w", prevErr)
}

func main() {
	s := "qwe\\45"
	res, err := Unpack(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func Unpack(s string) (string, error) {
	strLen := utf8.RuneCountInString(s)
	runes := []rune(s)
	if strLen < 2 {
		switch {
		case strLen == 0:
			return "", nil
		case strLen == 1:
			if !unicode.IsNumber(runes[0]) && runes[0] != backSlash {
				return s, nil
			} else {
				return "", ErrInvalidString
			}
		}
	}
	var sb strings.Builder
	lstIdx := strLen - 1
	var sbNum strings.Builder
	var lstRn rune
	var esc bool
	for i := 0; i < lstIdx; i++ {
		switch {
		case runes[i] == backSlash:
			if esc {
				lstRn = runes[i]
				esc = false
			} else {
				sNum := sbNum.String()
				sbNum.Reset()
				if len(sNum) == 0 {
					if lstRn == 0 {
						esc = true
					} else {
						sb.WriteRune(lstRn)
						lstRn = 0
						esc = true
					}
				} else {
					num, err := strconv.Atoi(sNum)
					if err != nil {
						return "", SbNumConvError(err)
					}
					for j := 0; j < num; j++ {
						sb.WriteRune(lstRn)
					}
					lstRn = 0
					esc = true
				}
			}
		case unicode.IsNumber(runes[i]):
			if !esc {
				if lstRn == 0 {
					return "", ErrInvalidString
				}
				sbNum.WriteRune(runes[i])
			} else {
				lstRn = runes[i]
				esc = false
			}
		default:
			if esc {
				switch {
				case runes[i] == 'n':
					sb.WriteRune('\n')
					esc = false
				case runes[i] == 't':
					sb.WriteRune('\t')
					esc = false
				default:
					return "", ErrInvalidString
				}
			} else {
				sNum := sbNum.String()
				sbNum.Reset()
				if len(sNum) == 0 {
					if lstRn == 0 {
						lstRn = runes[i]
					} else {
						sb.WriteRune(lstRn)
						lstRn = runes[i]
					}
				} else {
					num, err := strconv.Atoi(sNum)
					if err != nil {
						return "", SbNumConvError(err)
					}
					for j := 0; j < num; j++ {
						sb.WriteRune(lstRn)
					}
					lstRn = runes[i]
				}
			}
		}
	}
	// вышли из цикла
	if esc {
		switch {
		case unicode.IsNumber(runes[lstIdx]):
			sb.WriteRune(runes[lstIdx])
		case runes[lstIdx] == backSlash:
			sb.WriteRune(runes[lstIdx])
		case runes[lstIdx] == 'n':
			sb.WriteRune('\n')
		case runes[lstIdx] == 't':
			sb.WriteRune('\t')
		default:
			return "", ErrInvalidString
		}
	} else {
		switch {
		case runes[lstIdx] == backSlash:
			return "", ErrInvalidString
		case unicode.IsNumber(runes[lstIdx]):
			switch {
			case lstRn == 0:
				return "", ErrInvalidString
			default:
				sbNum.WriteRune(runes[lstIdx])
				num, err := strconv.Atoi(sbNum.String())
				if err != nil {
					return "", SbNumConvError(err)
				}
				for j := 0; j < num; j++ {
					sb.WriteRune(lstRn)
				}
			}
		default:
			sNum := sbNum.String()
			sbNum.Reset()
			if len(sNum) == 0 {
				sb.WriteRune(lstRn)
			} else {
				num, err := strconv.Atoi(sNum)
				if err != nil {
					return "", SbNumConvError(err)
				}
				for j := 0; j < num; j++ {
					sb.WriteRune(lstRn)
				}
			}
			sb.WriteRune(runes[lstIdx])
		}
	}
	return sb.String(), nil
}
