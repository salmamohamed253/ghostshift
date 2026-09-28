package sim

import "testing"

func TestNewStreams_SameSeedSameSequence(t *testing.T) {
	a := NewStreams(42)
	b := NewStreams(42)
	for i := 0; i < 100; i++ {
		x, y := a.Arrivals.Float64(), b.Arrivals.Float64()
		if x != y {
			t.Fatalf("draw %d: got %v and %v, want identical", i, x, y)
		}
	}
}

func TestNewStreams_DifferentSeedDifferentSequence(t *testing.T) {
	a := NewStreams(42)
	b := NewStreams(43)
	if a.Arrivals.Float64() == b.Arrivals.Float64() {
		t.Fatal("seeds 42 and 43 produced the same first draw")
	}
}

func TestNewStreams_DrawingOneStreamDoesNotShiftAnother(t *testing.T) {
	quiet := NewStreams(42)
	busy := NewStreams(42)
	for i := 0; i < 50; i++ {
		busy.PrepTime.Float64()
	}
	for i := 0; i < 100; i++ {
		x, y := quiet.Arrivals.Float64(), busy.Arrivals.Float64()
		if x != y {
			t.Fatalf("draw %d: arrivals shifted after prep draws: %v vs %v", i, x, y)
		}
	}
}