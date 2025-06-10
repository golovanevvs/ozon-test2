package task51

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type image struct {
	id int
	v  int
}

type server struct {
	id int
	v  int
}

type result struct {
	mapIDImIDSe map[int]int
	diff        int
}

func Task() {
	file, err := os.Open("./tests/22")
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

	servers := make([]server, n)
	for i, v := range ss {
		vInt, _ := strconv.Atoi(v)
		server := server{
			id: i,
			v:  vInt,
		}
		servers[i] = server
	}

	sort.Slice(servers, func(i, j int) bool {
		return servers[i].v > servers[j].v
	})

	images := make([]image, m)
	for i, v := range is {
		vInt, _ := strconv.Atoi(v)
		image := image{
			id: i,
			v:  vInt,
		}
		images[i] = image
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].v > images[j].v
	})

	results := make([]result, 0)

	for i := range servers {
		result := result{
			mapIDImIDSe: make(map[int]int),
		}
		t0 := getTime(images[0].v, servers[i].v)
		result.mapIDImIDSe[images[0].id] = servers[i].id
		resDiffi := 0

		for j := 1; j < len(images); j++ {
			diffMin := math.MaxInt64
			var idSe int

			for k := range servers {
				t := getTime(images[j].v, servers[k].v)
				var diff int
				if t == t0 {
					diffMin = 0
					idSe = servers[k].id
					break
				}
				if t < t0 {
					diff = t0 - t
				} else {
					diff = t - t0
				}
				if diff < diffMin {
					diffMin = diff
					idSe = servers[k].id
				}
			}

			result.mapIDImIDSe[images[j].id] = idSe
			if diffMin > resDiffi {
				resDiffi = diffMin
			}
		}
		result.diff = resDiffi
		results = append(results, result)
		if resDiffi == 0 {
			break
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].diff < results[j].diff
	})
	resDiff := results[0].diff
	resSlice := make([]string, m)
	for i := range resSlice {
		resSlice[i] = strconv.Itoa(results[0].mapIDImIDSe[i] + 1)
	}

	resString := strings.Join(resSlice, " ")
	ress := fmt.Sprintf("%d\n%s ", resDiff, resString)
	return ress
}

func getTime(i int, s int) int {
	if i%s == 0 {
		return i / s
	} else {
		return i/s + 1
	}
}
