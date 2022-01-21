package sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoNumbers(t *testing.T) {
	var result int = SumNumbers(2, 2)

	assert.Equal(t, 4, result, "result should be 4")

	fmt.Println("done")
}

func TestTwoNumbersSubtraction(t *testing.T) {
	var result int = Subtraction(5, 1)

	assert.Equal(t, 4, result, "Result should return 4 because 5 - 1 equal 4")
}
