package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sub benchmark testing
func BenchmarkSub(b *testing.B) {
	b.Run("Andre", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Andre")
		}
	})

	b.Run("Maesha", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Maesha")
		}
	})
}

// BenchmarkHelloWorld
func BenchmarkHelloWorldAndre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Andre")
	}
}

func BenchmarkHelloWorldMaesha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Maesha")
	}
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Andre",
			Request:  "Andre",
			Expected: "Hello Andre",
		},
		{
			Name:     "Maesha",
			Request:  "Maesha",
			Expected: "Hello Maesha",
		},
		{
			Name:     "Khannedy",
			Request:  "Khannedy",
			Expected: "Hello Khannedy",
		},
		{
			Name:     "Budi",
			Request:  "Budi",
			Expected: "Hello Budi",
		},
		{
			Name:     "Joko",
			Request:  "Joko",
			Expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			var result string = HelloWorld(test.Request)

			require.Equal(t, test.Expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Andre", func(t *testing.T) {
		var result string = HelloWorld("Andre")

		require.Equal(t, "Hello Andre", result, "Result must be 'Hello Andre'")
	})

	t.Run("Maesha", func(*testing.T) {
		var result string = HelloWorld("Maesha")

		require.Equal(t, "Hello Maesha", result, "Result must be 'Hello Maesha'")
	})
}

// this unit test always run but this unit test only once running at package
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "window" {
		t.Skip("Cannot run on window")
	}

	var result string = HelloWorld("andre")

	assert.Equal(t, "Hello andre", result, "Result must be 'Hello andre'")

	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldAssert(t *testing.T) {
	var result string = HelloWorld("andre")

	assert.Equal(t, "Hello andre", result, "Result must be 'Hello andre'")

	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	var result string = HelloWorld("andre")

	require.Equal(t, "Hello andre", result, "Result must be 'Hello andre'")

	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldPertama(t *testing.T) {
	var result string = HelloWorld("andre")

	if result != "Hello andre" {
		// error
		t.Error("Result Must Be Hello andre")
	}

	fmt.Println("TestHelloWorldPertama done")
}

func TestHelloWorldKedua(t *testing.T) {
	var result string = HelloWorld("maesha")

	if result != "Hello maesha" {
		// error
		t.Fatal("Result Must Be Hello maesha")
	}

	fmt.Println("TestHelloWorldKedua done")
}
