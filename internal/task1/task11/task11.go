package task11

import (
	"bufio"
	"fmt"
	"os"
)

func Task() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("../internal/task1/tests/2")
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

		// Чтение из одной строки нескольких значений, разделённых пробелом
		var s string
		fmt.Fscanln(in, &s)

		// Чтение целой строки
		// str1, _ := in.ReadString('\n')
		// Удаление символа \n (на некоторых машинах надо удалить два символа \r\n)
		// str := strings.Trim(str1, "\n")
		// str := strings.Trim(str1, "\r\n")
		// Для избежания ошибки выведем str (закомментить при использовании шаблона)
		// fmt.Println(str)

		// Запуск и вывод в out решения подзадачи t
		// В зависимости от условия задачи алгоритм вывода может потребовать доработки
		fmt.Fprintln(out, tTaskSolving(s))
	}
}

// tTaskSolving## - функция для решения подзадачи t задачи # (вариант #)
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving(s string) string {
	if len(s) == 1 {
		return "YES"
	}

	if s[0] != s[len(s)-1] {
		return "NO"
	}

	if len(s) < 3 && s[0] == s[1] {
		return "YES"
	}

	for i := 0; i <= len(s)-3; {
		if s[i] == s[i+1] {
			i++
		} else {
			if s[i] != s[i+1] && s[i] != s[i+2] {
				return "NO"
			} else if s[i] != s[i+1] && s[i] == s[i+2] {
				i += 2
			}
		}
	}

	return "YES"
}
