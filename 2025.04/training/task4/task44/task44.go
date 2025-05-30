package task44

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
		datas := make([]string, 0, n)
		for range n {
			strData1, _ := in.ReadString('\n')
			strData := strings.Trim(strData1, "\r\n")
			datas = append(datas, strData)
		}

		// Запуск и вывод в out решения подзадачи t
		// В зависимости от условия задачи алгоритм вывода может потребовать доработки
		// fmt.Println("")
		// fmt.Println("Задача", tt+1)
		fmt.Fprintln(out, tTaskSolving(datas, n, m, k))
	}
}

// tTaskSolving - функция для решения подзадачи t задачи
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving(datas []string, n, m, k int) (result string) {
	if findWinForFullDatas(datas, 'X', k, n, m) || findWinForFullDatas(datas, 'O', k, n, m) {
		return "NO"
	}
	for u := range n {
		for v := range m {
			if datas[u][v] == '.' {
				if findWinForOnePoint(datas, k, n, m, u, v) {
					return "YES"
				}
			}
		}
	}
	return "NO"
}

func findWinForOnePoint(datas []string, k, n, m, u, v int) bool {
	dirs := [][]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		count := 1

		for x, y := u+dx, v+dy; x >= 0 && x < n && y >= 0 && y < m && datas[x][y] == 'X'; x, y = x+dx, y+dy {
			count++
			if count >= k {
				return true
			}
		}

		for x, y := u-dx, v-dy; x >= 0 && x < n && y >= 0 && y < m && datas[x][y] == 'X'; x, y = x-dx, y-dy {
			count++
			if count >= k {
				return true
			}
		}
	}

	return false
}

func findWinForFullDatas(datas []string, target byte, k, n, m int) bool {
	for u := range n {
		count := 0
		for v := range m {
			if datas[u][v] == target {
				count++
				if count >= k {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	for u := range m {
		count := 0
		for v := range n {
			if datas[v][u] == target {
				count++
				if count >= k {
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
				if count >= k {
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
				if count >= k {
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
