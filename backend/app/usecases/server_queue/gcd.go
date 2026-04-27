package server_queue

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = gcd(result, arr[i])
		if result == 1 {
			return 1
		}
	}
	return result
}
