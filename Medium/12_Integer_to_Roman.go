/*
https://leetcode.com/problems/integer-to-roman/

Seven different symbols represent Roman numerals with the different values.

Input: num = 3749
Output: "MMMDCCXLIX"

I = 1
V = 5
X = 10
L = 50
C = 100
D = 500
M = 1000
*/

package Medium

func intToRoman(num int) string {
	num1 := []string{"", "M", "MM", "MMM", "MMMM", "MMMMM"}
	num2 := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	num3 := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	num4 := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return num1[num/1000] + num2[(num%1000)/100] + num3[(num%100)/10] + num4[num%10]
}
