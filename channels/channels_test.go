package channels

import "testing"

func TestSimulateChannels(t *testing.T) {
	t.Run(
		"SimulateChannels",
		func(t *testing.T) {
			SimulateChannels(1000, 1000000)
		},
	)
}

func BenchmarkSimulateChannels(b *testing.B) {
	for b.Loop() {
		SimulateChannels(1000, 1000000)
	}
}
