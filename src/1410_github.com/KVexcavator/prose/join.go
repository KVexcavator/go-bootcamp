// Функция, которая объединяет несколько строк в одну строку
package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		// разделяют все строки, кромеполедней запятыми
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		// вставляее "and" перед последным пересисляемым словом
		result += ", and "
		// добавляет последнюю строку
		result += phrases[len(phrases)-1]
		return result
	}
}
