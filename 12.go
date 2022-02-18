package main

var LIST = []struct {
	val  int
	char string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	ret := []byte{}
	for _, vs := range LIST {
		for num > vs.val {
			num -= vs.val
			ret = append(ret, vs.char...)
		}
	}
	return string(ret)
}
