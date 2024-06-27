package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v
func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	// }

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
	assert.NotEqual(t, AddOne(2), 2)

	// assert for nil (good for errors)
	assert.Nil(t, nil, nil)
}

func TestAddTwo(t *testing.T) {

	// assert.Equal(t, AddTwo(1), 3, "AddOne(2) should be 3")
	// assert.NotEqual(t, AddTwo(1), 2)

	// // assert for nil (good for errors)
	// assert.Nil(t, nil, nil)

	actual := AddTwo(4)
	if actual != 2 {
		t.Errorf("AddOne(%d), input %d, actual = %d", 1, 2, actual)
	}
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3) // Nếu failed thì đoạn code dưới sẽ không được thực thi
	fmt.Println("Not executed")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3) // Nếu failed thì đoạn code dưới sẽ được thực thi
	fmt.Println("Executing")
}
