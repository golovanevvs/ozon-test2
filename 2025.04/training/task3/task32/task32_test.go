package task32

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	openAnyFile := false
	for i := 21; i <= 21; i++ {
		file, err := os.Open(fmt.Sprintf("../tests/%d", i))
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %d\n", i)
			continue
		}
		defer file.Close()
		openAnyFile = true

		t.Run(fmt.Sprintf("Test: %d", i), func(t *testing.T) {
			fmt.Printf("Тест %d\n", i)
			in := bufio.NewReader(file)

			expecteds, err := os.ReadFile(fmt.Sprintf("../tests/%d.a", i))
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
				exp, _ := strconv.Atoi(expected)
				res, _ := strconv.Atoi(sliceResults1[j])
				assert.Equal(t, exp, res)
			}
		})
	}
	require.True(t, openAnyFile)
}

// func BenchmarkTask3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Task3()
// 	}
// }
