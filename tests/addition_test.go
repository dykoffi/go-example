package utils

import (
	"lab/exp1/src/utils"
	"testing"
)

func TestAddition(t *testing.T) {
	got := utils.Addtion(2, 5)
	if got != 7 {
		t.Errorf("Addtion(2, 5) = %d; want 7", got)
	}
}

func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.Addtion(i, i+1)
	}
}
