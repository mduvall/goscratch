package algorithms

import (
	"fmt"
	"testing"
)

func Test_BinomialCoefficientsWorksCorrectly(t *testing.T) {
	coeff := BinomialCoefficients(10, 8)
	fmt.Println(coeff)
}
