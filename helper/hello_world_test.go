package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Andre", func(t *testing.T) {
		var result string = HelloWorld("Andre")

		require.Equal(t, "Hello Andre", result, "Result must be 'Hello Andre'")
	})

	t.Run("Maesha", func(*testing.T) {
		var result string = HelloWorld("Maesha")

		require.Equal(t, "Hai Maesha", result, "Result must be 'Hello Maesha'")
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
