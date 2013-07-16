package algorithms

import (
	"fmt"
	"testing"
)

func Test_FibWorks(t *testing.T) {
	f := Fib(9)

	fmt.Println(f)
}
