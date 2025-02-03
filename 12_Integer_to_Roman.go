// Seven different symbols represent Roman numerals with the different values.

// Input: num = 3749
// Output: "MMMDCCXLIX"

//I = 1
//V = 5
//X = 10
//L = 50
//C = 100
//D = 500
//M = 1000

package main

import (
	"strings"
)

func intToRoman(num int) string {
	num1 := num / 1000
	num2 := num / 100 % 10
	num3 := num / 10 % 10
	num4 := num % 10

	s1 := strings.Repeat("M", num1)
	s2, s3, s4 := "", "", ""

	switch num2 {
	case 4:
		s2 = "CD"
	case 9:
		s2 = "CM"
	default:
		if num2 > 4 {
			s2 = "D"
			num2 = num2 - 5
		}
		s2 = s2 + strings.Repeat("C", num2)
	}

	switch num3 {
	case 4:
		s3 = "XL"
	case 9:
		s3 = "XC"
	default:
		if num3 > 4 {
			s3 = "L"
			num3 = num3 - 5
		}
		s3 = s3 + strings.Repeat("X", num3)
	}

	switch num4 {
	case 4:
		s4 = "IV"
	case 9:
		s4 = "IX"
	default:
		if num4 > 4 {
			s4 = "V"
			num4 = num4 - 5
		}
		s4 = s4 + strings.Repeat("I", num4)
	}
	return s1 + s2 + s3 + s4
}
