package main

type size struct {
	ID      int
	ValueTr string
	ValueEn string
}

var sizeMap = map[int]size{
	1:  {ID: 1, ValueTr: "XXXS", ValueEn: "XXXS"},
	2:  {ID: 2, ValueTr: "XXS", ValueEn: "XXS"},
	3:  {ID: 3, ValueTr: "XS", ValueEn: "XS"},
	4:  {ID: 4, ValueTr: "S", ValueEn: "S"},
	5:  {ID: 5, ValueTr: "M", ValueEn: "M"},
	6:  {ID: 6, ValueTr: "L", ValueEn: "L"},
	7:  {ID: 7, ValueTr: "XL", ValueEn: "XL"},
	8:  {ID: 8, ValueTr: "XXL", ValueEn: "XXL"},
	9:  {ID: 9, ValueTr: "3XL", ValueEn: "3XL"},
	10: {ID: 10, ValueTr: "XS/S", ValueEn: "XS/S"},
	11: {ID: 11, ValueTr: "XXS/XS", ValueEn: "XXS/XS"},
	12: {ID: 12, ValueTr: "S/M", ValueEn: "S/M"},
	13: {ID: 13, ValueTr: "M/L", ValueEn: "M/L"},
	14: {ID: 14, ValueTr: "L/XL", ValueEn: "L/XL"},
	15: {ID: 15, ValueTr: "SM", ValueEn: "SM"},
	16: {ID: 16, ValueTr: "ML", ValueEn: "ML"},
	17: {ID: 17, ValueTr: "0", ValueEn: "0"},
	18: {ID: 18, ValueTr: "01", ValueEn: "01"},
	19: {ID: 19, ValueTr: "1", ValueEn: "1"},
	20: {ID: 20, ValueTr: "2", ValueEn: "2"},
	21: {ID: 21, ValueTr: "3", ValueEn: "3"},
	22: {ID: 22, ValueTr: "4", ValueEn: "4"},
	23: {ID: 23, ValueTr: "5", ValueEn: "5"},
	24: {ID: 24, ValueTr: "6", ValueEn: "6"},
	25: {ID: 25, ValueTr: "8", ValueEn: "8"},
	26: {ID: 26, ValueTr: "8.5", ValueEn: "8.5"},
	27: {ID: 27, ValueTr: "10", ValueEn: "10"},
	28: {ID: 28, ValueTr: "12", ValueEn: "12"},
	29: {ID: 29, ValueTr: "14", ValueEn: "14"},
	30: {ID: 30, ValueTr: "16", ValueEn: "16"},
	31: {ID: 31, ValueTr: "23", ValueEn: "23"},
	32: {ID: 32, ValueTr: "24", ValueEn: "24"},
	33: {ID: 33, ValueTr: "25", ValueEn: "25"},
	34: {ID: 34, ValueTr: "26", ValueEn: "26"},
	35: {ID: 35, ValueTr: "27", ValueEn: "27"},
	36: {ID: 36, ValueTr: "28", ValueEn: "28"},
	37: {ID: 37, ValueTr: "29", ValueEn: "29"},
	38: {ID: 38, ValueTr: "30", ValueEn: "30"},
	39: {ID: 39, ValueTr: "31", ValueEn: "31"},
	40: {ID: 40, ValueTr: "32", ValueEn: "32"},
	41: {ID: 41, ValueTr: "33.5", ValueEn: "33.5"},
	42: {ID: 42, ValueTr: "34", ValueEn: "34"},
	43: {ID: 43, ValueTr: "36", ValueEn: "36"},
	44: {ID: 44, ValueTr: "37", ValueEn: "37"},
	45: {ID: 45, ValueTr: "38", ValueEn: "38"},
	46: {ID: 46, ValueTr: "38.5", ValueEn: "38.5"},
	47: {ID: 47, ValueTr: "39", ValueEn: "39"},
	48: {ID: 48, ValueTr: "40", ValueEn: "40"},
	49: {ID: 49, ValueTr: "41", ValueEn: "41"},
	50: {ID: 50, ValueTr: "42", ValueEn: "42"},
	51: {ID: 51, ValueTr: "44", ValueEn: "44"},
	52: {ID: 52, ValueTr: "46", ValueEn: "46"},
	53: {ID: 53, ValueTr: "48", ValueEn: "48"},
	54: {ID: 54, ValueTr: "50", ValueEn: "50"},
	55: {ID: 55, ValueTr: "52", ValueEn: "52"},
	56: {ID: 56, ValueTr: "56", ValueEn: "56"},
	57: {ID: 57, ValueTr: "58", ValueEn: "58"},
	58: {ID: 58, ValueTr: "Standart", ValueEn: "Standard"},
}

func sizeAI(userPrompt string) (string, error) {
	var systemPrompt = sizeSystemPrompt
	
	output, err := askAI(userPrompt, systemPrompt)
	if err != nil {
		return "", err
	}
	
	return output, nil
}