package task51

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/goforj/godump"
)

type image struct {
	id       int
	v        int
	mapIDSeT map[int]int
}

type server struct {
	id int
	v  int
}

// type result struct {
// 	mapIDImIDSe map[int]int
// 	diff        int
// }

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

	for i, v := range images {
		images[i].mapIDSeT = make(map[int]int)
		for _, w := range servers {
			images[i].mapIDSeT[w.id] = getTime(v.v, w.v)
		}
	}

	godump.Dump(images)
	return ""
}

func getTime(i int, s int) int {
	if i%s == 0 {
		return i / s
	} else {
		return i/s + 1
	}
}
