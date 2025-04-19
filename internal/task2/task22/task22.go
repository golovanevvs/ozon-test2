package task22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Exchange struct {
	from string
	to   string
	n    int
	m    int
}

type Bank struct {
	exchanges []Exchange
}

func Task22() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	t := readInt(scanner)

	for i := 0; i < t; i++ {
		n := readInt(scanner) // Количество банков
		m := readInt(scanner) // Количество курсов обмена в каждом банке

		banks := make([]Bank, n)
		for j := 0; j < n; j++ {
			exchanges := make([]Exchange, m)
			for k := 0; k < m; k++ {
				from := readString(scanner)
				to := readString(scanner)
				nVal := readInt(scanner)
				mVal := readInt(scanner)
				exchanges[k] = Exchange{from, to, nVal, mVal}
			}
			banks[j] = Bank{exchanges}
		}

		maxDollars := 0.0
		memo := make(map[string]float64)

		var dfs func(currency string, amount float64, usedBanks map[int]bool)
		dfs = func(currency string, amount float64, usedBanks map[int]bool) {
			if currency == "USD" {
				if amount > maxDollars {
					maxDollars = amount
				}
				return
			}

			key := fmt.Sprintf("%s_%.10f", currency, amount)
			if val, ok := memo[key]; ok && val >= amount {
				return
			}
			memo[key] = amount

			for bankIdx, bank := range banks {
				if usedBanks[bankIdx] {
					continue
				}
				usedBanks[bankIdx] = true
				for _, exchange := range bank.exchanges {
					if exchange.from == currency {
						newAmount := amount / float64(exchange.n) * float64(exchange.m)
						dfs(exchange.to, newAmount, usedBanks)
					}
				}
				usedBanks[bankIdx] = false
			}
		}

		usedBanks := make(map[int]bool)
		dfs("RUB", 1.0, usedBanks)

		fmt.Printf("%.10f\n", maxDollars)
	}
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
