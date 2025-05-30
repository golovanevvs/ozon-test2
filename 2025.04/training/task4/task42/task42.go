package task42

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("./tests/1")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %s\n", err.Error())
	}
	defer file.Close()
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
		var k int
		fmt.Fscanln(in, &k)
		// Чтение целой строки до переноса
		str1, _ := in.ReadString('\n')
		str := strings.Trim(str1, "\r\n")
		s := strings.Fields(str)
		n, _ := strconv.Atoi(s[0])
		m, _ := strconv.Atoi(s[1])
		datas := make([][]string, n)
		for i := range n {
			datas[i] = make([]string, m)
			strData1, _ := in.ReadString('\n')
			strData := strings.Trim(strData1, "\r\n")
			datas[i] = strings.Split(strData, "")
		}

		// Запуск и вывод в out решения подзадачи t
		// В зависимости от условия задачи алгоритм вывода может потребовать доработки
		// fmt.Println("")
		// fmt.Println("Задача", tt+1)
		fmt.Fprintln(out, tTaskSolving(datas, k))
	}
}

// tTaskSolving - функция для решения подзадачи t задачи
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving(datas [][]string, k int) (result string) {
	if findWin(datas, "X", k) || findWin(datas, "O", k) {
		return "NO"
	}
	for u := range datas {
		for v := range datas[u] {
			if datas[u][v] == "." {
				datasTemp := make([][]string, len(datas))
				for a := range datas {
					datasTemp[a] = make([]string, len(datas[a]))
					copy(datasTemp[a], datas[a])
				}
				datasTemp[u][v] = "X"
				if findWin(datasTemp, "X", k) {
					return "YES"
				}
			}
		}
	}
	return "NO"
}

func findWin(datas [][]string, target string, countForWin int) bool {
	n := len(datas)
	m := len(datas[0])
	for u := range datas {
		count := 0
		for v := range datas[u] {
			if datas[u][v] == target {
				count++
				if count >= countForWin {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	for u := range datas[0] {
		count := 0
		for v := range datas {
			if datas[v][u] == target {
				count++
				if count >= countForWin {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	i, j := 0, 0
	for range n + m - 1 {
		count := 0
		for u, v := i, j; u >= 0 && v <= m-1; u, v = u-1, v+1 {
			if datas[u][v] == target {
				count++
				if count >= countForWin {
					return true
				}
			} else {
				count = 0
			}
		}
		if i < n-1 {
			i++
		} else {
			j++
		}
	}
	i, j = 0, m-1
	for range n + m - 1 {
		count := 0
		for u, v := i, j; u >= 0 && v >= 0; u, v = u-1, v-1 {
			if datas[u][v] == target {
				count++
				if count >= countForWin {
					return true
				}
			} else {
				count = 0
			}
		}
		if i < n-1 {
			i++
		} else {
			j--
		}
	}
	return false
}
