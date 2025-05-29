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
	d := []string{"X", ".", "X", "B"}
	s := findVictory(d)
	if s {
		return "YES"
	}
	return "NO"
}

func findVictory(strSlice []string) bool {
	reg := regexp.MustCompile(`X{3}`)
	str := strings.Join(strSlice, "")
	res := reg.MatchString(str)
	fmt.Println(str)
	if res {
		return res
	}
	return res
}

func findPotVictory(strSlice []string) bool {
	reg1 := regexp.MustCompile(`\.X+`)
	reg2 := regexp.MustCompile(`X+\.X+`)
	reg3 := regexp.MustCompile(`X+\.`)
	str := strings.Join(strSlice, "")
	res1 := reg1.MatchString(str)
	res2 := reg2.MatchString(str)
	res3 := reg3.MatchString(str)
	fmt.Println(str)
	if res1 || res2 || res3 {
		return true
	}
	return false
}
