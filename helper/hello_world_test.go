package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	var result string = HelloWorld("andre");

	assert.Equal(t, "Hello andre", result, "Result must be 'Hello andre'")

	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	var result string = HelloWorld("andre");

	require.Equal(t, "Hello andre", result, "Result must be 'Hello andre'")

	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldPertama(t *testing.T) {
	var result string = HelloWorld("andre");

	if result != "Hello andre" {
		// error
		t.Error("Result Must Be Hello andre")
	}

	fmt.Println("TestHelloWorldPertama done")
}

func TestHelloWorldKedua(t *testing.T) {
	var result string = HelloWorld("maesha");

	if result != "Hello maesha" {
		// error
		t.Fatal("Result Must Be Hello maesha")
	}

	fmt.Println("TestHelloWorldKedua done")
}