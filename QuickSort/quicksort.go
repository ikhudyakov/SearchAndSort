package quicksort

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	split := partition(arr)

	QuickSort(arr[:split]) // выполняем QuickSort для массива от 0-го элемента до возращенного из функции partion элемента
	QuickSort(arr[split:]) // выполняем QuickSort для массива от возращенного из функции partion элемента до последнего элемента
}

func partition(arr []int) int {
	pivot := arr[len(arr)/2] // получаем элемент в середине массива

	left := 0             // нулевой элемент массива
	right := len(arr) - 1 // последний элемент массива

	for {
		for ; arr[left] < pivot; left++ { // перебираем элементы слева от pivot, пока не найдем такой, который будет больше pivot
		}

		for ; arr[right] > pivot; right-- { // перебираем элементы справа от pivot, пока не найдем такой, который будет меньше pivot
		}

		if left >= right { // если индекс левого элемента получился больше или равен индексу правого, возвращаем правый и повторяем QuickSort
			return right
		}

		arr[left], arr[right] = arr[right], arr[left] // меняем местами удовлетворяющие условиям выше элементы
	}
}
