package task32

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type sepStr struct {
	str1  string
	str2  string
	color string
	index int
}

func Task() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("./tests/1")
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
	sliceStr := make([]sepStr, len(datas))
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
		sliceStr[i].color = "white"
	}

	// fmt.Println("sliceStr", sliceStr)

	count = 0
	stack := stack{}
	stack.push(0)

	for !stack.isEmpty() {
		i, _ := stack.pop()
		// fmt.Println("Взяли из стека i:", i, "color:", sliceStr[i].color)
		if sliceStr[i].color == "white" {
			// fmt.Println("Покрасили в серый, положили в стек")
			sliceStr[i].color = "gray"
			stack.push(i)
			for j := i + 1; j < len(sliceStr); j++ {
				// fmt.Println("Смотрим вершину j исходящего ребра:", j, "color:", sliceStr[j].color)
				if (len(sliceStr[i].str1) > 0 && sliceStr[i].str1 == sliceStr[j].str1) || (len(sliceStr[i].str2) > 0 && sliceStr[i].str2 == sliceStr[j].str2) {
					count++
				}
				if sliceStr[j].color == "white" {
					stack.push(j)
					// fmt.Println("Положили в стек, т.к. цвет белый")
				} else {
					// fmt.Println("Не положили в стек, т.к. цвет не белый")
				}
			}
		} else if sliceStr[i].color == "gray" {
			// fmt.Println("Покрасили в чёрный")
			sliceStr[i].color = "black"
		}
	}
	return count
}
