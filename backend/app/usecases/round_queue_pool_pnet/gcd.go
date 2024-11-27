package round_queue_pool_pnet

// Функция для нахождения НОД двух чисел с использованием алгоритма Евклида
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Функция для нахождения НОД массива чисел
func gcdArray(arr []int) int {
	if len(arr) == 0 {
		return 0 // Если массив пустой, возвращаем 0
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = gcd(result, arr[i])
		if result == 1 {
			return 1 // Если НОД равен 1, дальнейшие вычисления не имеют смысла
		}
	}
	return result
}
