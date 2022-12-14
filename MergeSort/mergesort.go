package mergesort

func MergeSort(arr []int) []int { // рекурсивная функция, чтобы разбить массив на две равные части и вызвать эту жу функцию снова для двух этих частей
	if len(arr) < 2 { // если длина меньше 2, возвращаем массив с 1 элементом
		return arr
	}
	first := MergeSort(arr[:len(arr)/2])  // вызываем функцию снова, пока len(arr) не будет меньще 2
	second := MergeSort(arr[len(arr)/2:]) // вызываем функцию снова, пока len(arr) не будет меньще 2
	return merge(first, second)           // вызываем функицю merge() для объединения двух массивов
}

func merge(a []int, b []int) []int { // функция сравниввает элементы в двух массивах и объединяет их в один отсортированный
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] { // берем элемент из массива a и b и сравниваем, если a[i] < b[j] то в результирующий массив кладем
			final = append(final, a[i]) // то в результирующий массив кладем a[i]
			i++
		} else {
			final = append(final, b[j]) // если наоборот, то b[j]
			j++
		}
	}
	for ; i < len(a); i++ { // если массив b закончился раньше, то оставшиеся элементы из массива a кладем в результирующий массив
		final = append(final, a[i])
	}
	for ; j < len(b); j++ { // если массив a закончился раньше, то оставшиеся элементы из массива b кладем в результирующий массив
		final = append(final, b[j])
	}
	return final //возвращаем отсортированный массив
}
