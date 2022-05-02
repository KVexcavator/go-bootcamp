//Ф-я должна возвращать большее из двух переданных ей целых чисел (содержит ошибку)
package compare

func Larger(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
