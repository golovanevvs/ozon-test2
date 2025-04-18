package task11

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	for i := 1; i <= 31; i++ {
		file, err := os.Open(fmt.Sprintf("C:\\KKO11\\Golang\\ozon-test2\\internal\\task1\\tests\\%d", i))
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %d\n", i)
			continue
		}
		defer file.Close()

		t.Run(fmt.Sprintf("Test: %d", i), func(t *testing.T) {
			in := bufio.NewReader(file)

			expected, err := os.ReadFile(fmt.Sprintf("C:\\KKO11\\Golang\\ozon-test2\\internal\\task1\\tests\\%d.a", i))
			require.Nil(t, err)

			var buffer bytes.Buffer
			out := bufio.NewWriter(&buffer)

			Run(in, out)

			out.Flush()

			result, err := io.ReadAll(bufio.NewReader(&buffer))
			require.Nil(t, err)

			require.Equal(t, string(expected), string(result))
		})
	}
}

// func BenchmarkTask3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Task3()
// 	}
// }
