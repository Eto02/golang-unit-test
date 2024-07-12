package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Tahta",
			request: "Tahta",
		},
		{
			name:    "Mubarak",
			request: "Mubarak",
		},
		{
			name:    "Budi",
			request: "BudiSaya",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Tahta", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Tahta")
		}
	})
	b.Run("Failah", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Failah")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Tahta")
	}
}

func BenchmarkHelloWorldFailah(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Failah")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Tahta",
			request:  "Tahta",
			expected: "Hello Tahta",
		}, {
			name:     "Mubarak",
			request:  "Mubarak",
			expected: "Hello Mubarak",
		}, {
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HelloWorld(tt.request)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}
func TestSubTest(t *testing.T) {
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result must be 'Budi'")
		fmt.Println("Dieksekusi")
	})
	t.Run("Failah", func(t *testing.T) {
		result := HelloWorld("Failah")
		require.Equal(t, "Hello Failah", result, "Result must be 'Failah'")
		fmt.Println("Dieksekusi")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on windows ")
	}
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Tahta")
	require.Equal(t, "Hello Tahta", result, "Result must be 'Tahta'")
	fmt.Println("Dieksekusi")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Tahta")
	assert.Equal(t, "Hello Tahta", result, "Result must be 'Tahta'")
	fmt.Println("Dieksekusi")
}

func TestHelloWorldMubarak(t *testing.T) {
	result := HelloWorld("Mubarak")
	if result != "Hello Mubarak" {
		t.Error("Result must be 'Mubarak'")
	}
	fmt.Println("TestHelloWorldMubarak Done")
}
