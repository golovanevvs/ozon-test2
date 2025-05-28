package task31

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type sepStr struct {
	str1  string
	str2  string
	index int
}

func Task() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("../tests/1")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %s\n", err.Error())
	}
	in := bufio.NewReader(file)

	// Раскомментить при отправке на платформу, закомментить при запуске
	// in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Основное решение задачи
	Run(in, out)
}

// Run - основное решение задачи
// Менять какие-либо аргументы или возвращаемые значения не требуется
func Run(in *bufio.Reader, out *bufio.Writer) {
	// t - количество входных данных (подзадач)
	var t int
	fmt.Fscanln(in, &t)

	// Решение t подзадач
	for range t {
		// Примеры чтения параметров подзадачи i
		var n int
		fmt.Fscanln(in, &n)

		datas := make([]string, n)
		for i := range n {
			// Чтение из одной строки нескольких значений, разделённых пробелом
			// var s string
			// fmt.Fscan(in, &s)
			// fmt.Println(s)

			// Чтение целой строки до переноса
			str1, _ := in.ReadString('\n')
			// Удаление символа \n (на некоторых машинах надо удалить два символа \r\n)
			// str := strings.Trim(str1, "\n")
			str := strings.Trim(str1, "\r\n")
			// Для избежания ошибки выведем str (закомментить при использовании шаблона)
			datas[i] = str
		}
		// Запуск и вывод в out решения подзадачи t
		// В зависимости от условия задачи алгоритм вывода может потребовать доработки
		// fmt.Println("")
		// fmt.Println("Задача", tt+1)
		fmt.Fprintln(out, tTaskSolving(datas))
	}
}

// tTaskSolving - функция для решения подзадачи t задачи
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving(datas []string) (count int) {
	count = 0
	sliceStr := make([]sepStr, len(datas))
	visitedStr := make(map[int]bool)
	for i, data := range datas {
		str1 := make([]rune, 0, len(data)/2+1)
		str2 := make([]rune, 0, len(data)/2)
		for j, v := range data {
			if (j+1)%2 == 0 {
				str2 = append(str2, v)
			} else {
				str1 = append(str1, v)
			}
		}
		sliceStr[i].index = i
		sliceStr[i].str1 = string(str1)
		sliceStr[i].str2 = string(str2)
	}
	var dfs func(currentSlice []sepStr)
	dfs = func(currentSlice []sepStr) {
		for idx := 1; idx < len(currentSlice); idx++ {
			if !visitedStr[currentSlice[0].index] && ((len(currentSlice[0].str1) > 0 && currentSlice[0].str1 == currentSlice[idx].str1) || (len(currentSlice[0].str2) > 0 && currentSlice[0].str2 == currentSlice[idx].str2)) {
				count++
			}
			dfs(currentSlice[idx:])
		}
		visitedStr[currentSlice[0].index] = true
	}
	dfs(sliceStr)
	return count
}
