package task21

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type bank struct {
	rd float64
	re float64
	dr float64
	de float64
	er float64
	ed float64
}

func Task() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("../internal/task2/tests/1")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %s\n", err.Error())
	}
	in := bufio.NewReader(file)

	// Раскомментить при отправке на платформу, закомментить при запуске
	// in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Основное решение задачи # (вариант #)
	Run(in, out)
}

// Run## - основное решение задачи # (вариант #)
// Менять какие-либо аргументы или возвращаемые значения не требуется
func Run(in *bufio.Reader, out *bufio.Writer) {
	// t - количество входных данных (подзадач)
	var t int
	fmt.Fscanln(in, &t)

	// Решение t подзадач
	for i := 0; i < t; i++ {
		// Примеры чтения параметров подзадачи i
		banks := make([]bank, 3)

		for j := 0; j < 3; j++ {
			for k := 0; k < 6; k++ {

				// Чтение из одной строки нескольких значений, разделённых пробелом
				// var s string
				// fmt.Fscan(in, &s)
				// fmt.Println(s)

				// Чтение целой строки до переноса
				str1, _ := in.ReadString('\n')
				// Удаление символа \n (на некоторых машинах надо удалить два символа \r\n)
				str := strings.Trim(str1, "\n")
				// str := strings.Trim(str1, "\r\n")
				// Для избежания ошибки выведем str (закомментить при использовании шаблона)
				valueStrSlice := strings.Split(str, " ")
				value1, _ := strconv.ParseFloat(valueStrSlice[0], 64)
				value2, _ := strconv.ParseFloat(valueStrSlice[1], 64)
				value := value2 / value1
				switch k {
				case 0:
					banks[j].rd = value
				case 1:
					banks[j].re = value
				case 2:
					banks[j].dr = value
				case 3:
					banks[j].de = value
				case 4:
					banks[j].er = value
				case 5:
					banks[j].ed = value
				}

				// Запуск и вывод в out решения подзадачи t
				// В зависимости от условия задачи алгоритм вывода может потребовать доработки
			}
		}
		fmt.Fprintln(out, tTaskSolving(banks))
	}
}

// tTaskSolving## - функция для решения подзадачи t задачи # (вариант #)
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving(banks []bank) (maxD float64) {
	maxDs := make([]float64, 0)
	for i := range banks {
		var a0, a1, a2 int
		switch i {
		case 0:
			a0 = 0
			a1 = 1
			a2 = 2
		case 1:
			a1 = 0
			a0 = 1
			a2 = 2
		case 2:
			a2 = 0
			a1 = 1
			a0 = 2
		}
		d1 := banks[a2].rd + banks[a1].dr + banks[a0].rd
		maxDs = append(maxDs, d1)
		d2 := banks[a2].rd + banks[a1].er + banks[a0].rd
		maxDs = append(maxDs, d2)
		d3 := banks[a2].rd + banks[a1].de + banks[a0].ed
		maxDs = append(maxDs, d3)
		d4 := banks[a1].rd + banks[a2].dr + banks[a0].rd
		maxDs = append(maxDs, d4)
		d5 := banks[a1].rd + banks[a2].er + banks[a0].rd
		maxDs = append(maxDs, d5)
		d6 := banks[a1].rd + banks[a2].de + banks[a0].ed
		maxDs = append(maxDs, d6)
		d7 := banks[a1].re + banks[a0].ed
		maxDs = append(maxDs, d7)
		d8 := banks[a2].re + banks[a0].ed
		maxDs = append(maxDs, d8)
		d9 := banks[a0].rd
		maxDs = append(maxDs, d9)
	}

	return slices.Max(maxDs)
}
