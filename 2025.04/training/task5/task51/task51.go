package task51

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type image struct {
	idImg int
	s     map[int]int
}

func Task() {
	file, err := os.Open("./tests/1")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v", err)
		return
	}
	defer file.Close()

	in := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	run(in, out)
}

func run(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)
	fmt.Println(t)

	for range t {
		var n int
		fmt.Fscanln(in, &n)
		fmt.Println(n)

		var ssStr string
		ssStr1, _ := in.ReadString('\n')
		if strings.HasSuffix(ssStr1, "\r\n") {
			ssStr = strings.TrimRight(ssStr1, "\r\n")
		} else if strings.HasSuffix(ssStr1, "\n") {
			ssStr = strings.TrimRight(ssStr1, "\n")
		} else {
			ssStr = ssStr1
		}

		var m int
		fmt.Fscanln(in, &m)

		var isStr string
		isStr1, _ := in.ReadString('\n')
		if strings.HasSuffix(isStr1, "\r\n") {
			isStr = strings.TrimRight(isStr1, "\r\n")
		} else if strings.HasSuffix(isStr1, "\n") {
			isStr = strings.TrimRight(isStr1, "\n")
		} else {
			isStr = isStr1
		}

		ss := strings.Fields(ssStr)
		is := strings.Fields(isStr)

		fmt.Fprintln(out, tTaskSolving(ss, is))
	}
}

func tTaskSolving(ss []string, is []string) []image {
	res := make([]image, 0)
	for i, v := range is {
		ii := image{
			idImg: i,
			s:     make(map[int]int),
		}
		iInt, _ := strconv.Atoi(v)
		for j, s := range ss {
			sInt, _ := strconv.Atoi(s)
			ii.s[j] = getTime(iInt, sInt)
		}
		res = append(res, ii)
	}

	return res
}

func getTime(i int, s int) int {
	if i%s == 0 {
		return i / s
	} else {
		return i/s + 1
	}
}
