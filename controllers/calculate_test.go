package controllers_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-testing-11/controllers"
	"testing"
)

func TestSum(t *testing.T) {
	expected 	:= 10
	actual 		:= controllers.Sum(5,5)

	if actual != expected {
		t.Errorf("expected %d, get %d", expected, actual)
	}
}

func TestSumTS(t *testing.T) {
	expected 	:= 10
	actual 		:= controllers.Sum(5,5)

	equalMsg := fmt.Sprintf("expected %d, get %d", expected, actual)
	assert.Equal(t, expected, actual, equalMsg)
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		controllers.Sum(5,5)
	}
}
