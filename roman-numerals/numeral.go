package main

import "strings"

func ConvertToArabic(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

type romanNumeral struct {
	Value  uint16
	Symbol string
}

type romanNumerals []romanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

var allRomanNumerals = romanNumerals{
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

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

func ConvertToRomanOld(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	// for arabic > 0 {
	// 	switch {
	// 	case arabic > 9:
	// 		result.WriteString("X")
	// 		arabic -= 10
	// 	case arabic > 8:
	// 		result.WriteString("IX")
	// 		arabic -= 9
	// 	case arabic > 4:
	// 		result.WriteString("V")
	// 		arabic -= 5
	// 	case arabic > 3:
	// 		result.WriteString("IV")
	// 		arabic -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		arabic--
	// 	}
	// }

	// for i := arabic; i > 0; i-- {
	// 	if i == 5 {
	// 		result.WriteString("V")
	// 		break
	// 	}
	// 	if i == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}

	// 	result.WriteString("I")
	// }

	return result.String()
}

// func ConvertToArabicOld(roman string) uint16 {
// 	total := 0
// 	for i := 0; i < len(roman); i++ {
// 		symbol := roman[i]

// 		// look ahead to next symbol, if we can, and the current symbol is base 10 (only valid subtractors)
// 		if couldBeSubtractive(i, symbol, roman) {
// 			// nextSymbol := roman[i+1]

// 			// build the two character string
// 			// potentialNumber := string([]byte{symbol, nextSymbol})

// 			// get the value of the two character string
// 			// value := romanNumerals.ValueOf(potentialNumber)

// 			if value := allRomanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
// 				total += value
// 				i++ // move past this character too for the next loop
// 			} else {
// 				// total++
// 				total += allRomanNumerals.ValueOf(symbol)
// 			}
// 		} else {
// 			total += allRomanNumerals.ValueOf(symbol)
// 		}
// 	}
// 	return total
// }

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	// look ahead to next symbol, if we can, and the current symbol is base 10 (only valid subtractors)
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
