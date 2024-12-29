package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}

func GetBigger(str1 string, str2 string) (bigger string) {
	var comp int = strings.Compare(str1, str2)

	switch comp {
	case 1:
		bigger = str1
	case -1:
		bigger = str2
	case 0:
		bigger = "Same"
	}

	return
}