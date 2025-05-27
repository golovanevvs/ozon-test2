package task21

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	for i := 1; i <= 12; i++ {
		file, err := os.Open(fmt.Sprintf("./tests/%d", i))
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %d\n", i)
			continue
		}
		defer file.Close()

		t.Run(fmt.Sprintf("Test: %d", i), func(t *testing.T) {
			fmt.Printf("Тест %d\n", i)
			in := bufio.NewReader(file)

			expecteds, err := os.ReadFile(fmt.Sprintf("./tests/%d.a", i))
			require.Nil(t, err)

			var buffer bytes.Buffer
			out := bufio.NewWriter(&buffer)
			Run(in, out)
			out.Flush()
			results, err := io.ReadAll(bufio.NewReader(&buffer))
			require.Nil(t, err)

			sliceExpecteds := strings.Split(string(expecteds), "\r\n")
			sliceExpecteds1 := sliceExpecteds[:len(sliceExpecteds)-1]
			sliceResults := strings.Split(string(results), "\r\n")
			sliceResults1 := sliceResults[:len(sliceResults)-1]
			for j, expected := range sliceExpecteds1 {
				exp, _ := strconv.ParseFloat(expected, 64)
				res, _ := strconv.ParseFloat(sliceResults1[j], 64)
				a := exp - res
				assert.LessOrEqual(t, math.Abs(a), 0.000001)
			}
		})
	}
}

// func BenchmarkTask3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Task3()
// 	}
// }
