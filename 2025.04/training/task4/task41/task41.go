package task41

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Task() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("./tests/18")
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
	i, j := 0, 0
	n := len(datas)
	m := len(datas[0])
	sliceS := make([]string, 0, n+m-1)
	for u := range datas {
		s := ""
		for v := range datas[u] {
			s += datas[u][v]
		}
		sliceS = append(sliceS, s)
	}
	for u := range datas[0] {
		s := ""
		for v := range datas {
			s += datas[v][u]
		}
		sliceS = append(sliceS, s)
	}
	for range n + m - 1 {
		s := ""
		for u, v := i, j; u >= 0 && v <= m-1; u, v = u-1, v+1 {
			s += datas[u][v]
		}
		if len(s) >= k {
			sliceS = append(sliceS, s)
		}
		if i < n-1 {
			i++
		} else {
			j++
		}
	}
	i, j = 0, m-1
	for range n + m - 1 {
		s := ""
		for u, v := i, j; u >= 0 && v >= 0; u, v = u-1, v-1 {
			s += datas[u][v]
		}
		if len(s) >= k {
			sliceS = append(sliceS, s)
		}
		if i < n-1 {
			i++
		} else {
			j--
		}
	}
	flagCheckmate := false
	for _, v := range sliceS {
		if findCheckmate(v, k) {
			return "NO"
		}
		if !flagCheckmate && findShah(v, k) {
			flagCheckmate = true
		}
	}
	if flagCheckmate {
		return "YES"
	}
	return "NO"
}

func findCheckmate(str string, k int) bool {
	reg1 := regexp.MustCompile(fmt.Sprintf("X{%d}", k))
	reg2 := regexp.MustCompile(fmt.Sprintf("O{%d}", k))
	switch {
	case reg1.MatchString(str):
		return true
	case reg2.MatchString(str):
		return true
	}
	return false
}

func findShah(str string, k int) bool {
	reg1 := regexp.MustCompile(fmt.Sprintf(`\.X{%d}`, k-1))
	reg2 := regexp.MustCompile(fmt.Sprintf(`X{%d}\.`, k-1))
	reg3 := regexp.MustCompile(`X+\.X+`)
	switch {
	case reg1.MatchString(str):
		return true
	case reg2.MatchString(str):
		return true
	case reg3.MatchString(str):
		res := reg3.FindString(str)
		if len(res) >= k {
			return true
		}
	}
	return false
}
