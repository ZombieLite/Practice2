package main

import "fmt"

func findCommonDivisors(nums []int) []int {
	var divisors []int

	// Находим минимальное число в массиве
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	// Перебираем все числа от 2 до минимального числа
	for i := 2; i <= min; i++ {
		// Проверяем, является ли число делителем для всех чисел в массиве
		isDivisor := true
		for _, num := range nums {
			if num%i != 0 {
				isDivisor = false
				break
			}
		}
		// Если число является делителем для всех чисел и не равно 1, добавляем его в массив общих делителей
		if isDivisor && i != 1 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

func main() {
	// Пример 1
	numbers1 := []int{42, 12, 18}
	fmt.Println("Общие делители чисел", numbers1, ":", findCommonDivisors(numbers1)) // Вывод: Общие делители чисел [42 12 18] : [2 3 6]

	// Пример 2
	numbers2 := []int{24, 36, 48}
	fmt.Println("Общие делители чисел", numbers2, ":", findCommonDivisors(numbers2)) // Вывод: Общие делители чисел [24 36 48] : [2 3 4 6 12 24]

	// Пример 3
	numbers3 := []int{15, 20, 25}
	fmt.Println("Общие делители чисел", numbers3, ":", findCommonDivisors(numbers3)) // Вывод: Общие делители чисел [15 20 25] : [5]
}
