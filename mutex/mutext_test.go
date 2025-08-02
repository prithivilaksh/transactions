package mutex

import (
	"testing"
)

func TestSimulateMutex(t *testing.T) {
	t.Run(
		"SimulateMutex",
		func(t *testing.T) {
			SimulateMutex(1000, 1000000)
		},
	)
}

func BenchmarkSimulateMutex(b *testing.B) {
	for b.Loop() {
		SimulateMutex(1000, 1000000)
	}
}
