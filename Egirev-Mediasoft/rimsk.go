package main

import (
	"fmt"
)

func romanToArabic(romanNum string) int {
	// Создаем карту соответствий между римскими и арабскими цифрами
	romanArabicMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	arabicNum := 0
	for i := 0; i < len(romanNum); i++ {
		// Получаем текущую и следующую цифры в римском числе
		current := romanArabicMap[rune(romanNum[i])]
		if i+1 < len(romanNum) {
			next := romanArabicMap[rune(romanNum[i+1])]

			// Если следующая цифра больше текущей, то вычитаем текущую из следующей и увеличиваем счетчик i на 1
			if next > current {
				arabicNum += next - current
				i++
			} else {
				arabicNum += current
			}
		} else {
			arabicNum += current
		}
	}

	return arabicNum
}

func main() {
	// Тестируем конвертер римских чисел в арабские
	romanNums := []string{"I", "IV", "V", "IX", "XII", "XVII", "XX", "XLV", "LXXXVIII", "CXXVII", "DCLXXVI", "MCMXCIV"}
	for _, romanNum := range romanNums {
		arabicNum := romanToArabic(romanNum)
		fmt.Printf("%s в арабской записи: %d\n", romanNum, arabicNum)
	}
}
