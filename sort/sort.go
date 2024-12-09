package sort

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func Shellsort(arr *[]int, outputLabel *widget.Label) {
	n := len(*arr)
	output := strings.Builder{}
	output.WriteString(fmt.Sprintf("Изначальный массив: %v\n", *arr))

	// Начинаем с большего шага и уменьшаем его
	for gap := n / 2; gap > 0; gap /= 2 {
		output.WriteString(fmt.Sprintf("Шаг с gap = %d:\n", gap))
		insertionSort(arr, gap, &output)
		output.WriteString(fmt.Sprintf("  Массив после шага %d: %v\n", gap, *arr))
	}

	output.WriteString("Сортировка завершена :)\n")
	outputLabel.SetText(output.String())
}

func insertionSort(arr *[]int, gap int, output *strings.Builder) {
	n := len(*arr)
	for i := gap; i < n; i++ {
		// Сохраняем текущий элемент
		temp := (*arr)[i]
		j := i

		// Перемещаем элементы, которые больше temp
		for j >= gap {
			output.WriteString(fmt.Sprintf("Сравнение: %d и %d\n", (*arr)[j-gap], temp))
			if (*arr)[j-gap] > temp {
				(*arr)[j] = (*arr)[j-gap]
				j -= gap
			} else {
				break
			}
		}
		// Ставим temp на правильное место
		(*arr)[j] = temp

		// Добавляем вывод состояния массива после каждой вставки
		output.WriteString(fmt.Sprintf("  Вставка: arr[%d] = %d\n  Текущий массив: %v\n", j, temp, *arr))
	}
}
