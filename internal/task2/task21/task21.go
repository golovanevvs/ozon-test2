package task21

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type exchangeRate struct {
	from   string
	to     string
	bankID int
	rate   float64 // m/n - коэффициент обмена
}

func Task() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("../internal/task2/tests/5")
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
	for range t {
		// Примеры чтения параметров подзадачи i
		bankCount := 3

		var allRates []exchangeRate

		for bankID := range bankCount {
			for j := range 6 {
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
				values := strings.Split(str, " ")
				n, _ := strconv.Atoi(values[0])
				m, _ := strconv.Atoi(values[1])

				var from, to string
				switch j {
				case 0:
					from, to = "RUB", "USD"
				case 1:
					from, to = "RUB", "EUR"
				case 2:
					from, to = "USD", "RUB"
				case 3:
					from, to = "USD", "EUR"
				case 4:
					from, to = "EUR", "RUB"
				case 5:
					from, to = "EUR", "USD"
				}

				allRates = append(allRates, exchangeRate{
					from:   from,
					to:     to,
					bankID: bankID,
					rate:   float64(m) / float64(n),
				})
			}
		}
		// fmt.Println(allRates)
		// Запуск и вывод в out решения подзадачи t
		// В зависимости от условия задачи алгоритм вывода может потребовать доработки
		fmt.Fprintf(out, "%g\r\n", tTaskSolving(allRates, bankCount, "RUB", "USD", 1))
	}
}

// tTaskSolving## - функция для решения подзадачи t задачи # (вариант #)
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving(rates []exchangeRate, bankCount int, from, to string, initialAmount float64) (maxTo float64) {
	maxAmount := 0.0
	visitedBanks := make([]bool, bankCount)

	var dfs func(currentCurrency string, currentAmount float64, depth int)
	dfs = func(currentCurrency string, currentAmount float64, depth int) {
		if currentCurrency == to {
			if currentAmount > maxAmount {
				maxAmount = currentAmount
			}
		}

		if depth >= bankCount {
			return
		}

		for _, rate := range rates {
			if rate.from == currentCurrency && !visitedBanks[rate.bankID] {
				visitedBanks[rate.bankID] = true
				newAmount := currentAmount * rate.rate
				dfs(rate.to, newAmount, depth+1)
				visitedBanks[rate.bankID] = false
			}
		}
	}

	dfs(from, initialAmount, 0)
	return maxAmount
}
