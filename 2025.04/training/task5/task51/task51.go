package task51

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type tis struct {
	iID int
	sID int
	t   int
}

func Task() {
	file, err := os.Open("./tests/12")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v", err)
		return
	}
	defer file.Close()

	in := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)

	for range t {
		var n int
		fmt.Fscanln(in, &n)

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

		fmt.Fprintln(out, tTaskSolving(n, ss, m, is))
	}
}

func tTaskSolving(n int, ss []string, m int, is []string) string {
	if len(is) == 1 {
		return ""
	}

	servers := make([]int, n)
	for i, v := range ss {
		servers[i], _ = strconv.Atoi(v)
	}

	images := make([]int, m)
	for i, v := range is {
		images[i], _ = strconv.Atoi(v)
	}

	tiss := make([][]tis, m)

	for ii := range tiss {
		tiss[ii] = make([]tis, n)
		for is := range tiss[ii] {
			tiss[ii][is].iID = ii
			tiss[ii][is].sID = is
			tiss[ii][is].t = getTime(images[ii], servers[is])
		}
	}

	allCombs := [][]tis{{}}
	for _, column := range tiss {
		var newCombs [][]tis
		for _, value := range column {
			for _, comb := range allCombs {
				newComb := append([]tis{}, comb...)
				newComb = append(newComb, value)
				newCombs = append(newCombs, newComb)
			}
		}
		allCombs = newCombs
	}

	diffs := make([]int, 0)

	for _, v := range allCombs {
		max := slices.MaxFunc(v, func(a, b tis) int {
			return a.t - b.t
		})
		min := slices.MinFunc(v, func(a, b tis) int {
			return a.t - b.t
		})
		diff := max.t - min.t
		diffs = append(diffs, diff)
	}

	minDiff := slices.Min(diffs)
	minDiffIndex := slices.Index(diffs, minDiff)
	resMinDiff := allCombs[minDiffIndex]
	var resSe string
	for _, v := range resMinDiff {
		resSe += strconv.Itoa(v.sID+1) + " "
	}
	res := fmt.Sprintf("%d\n%s", minDiff, resSe)
	return res
}

func getTime(i int, s int) int {
	if i%s == 0 {
		return i / s
	} else {
		return i/s + 1
	}
}
