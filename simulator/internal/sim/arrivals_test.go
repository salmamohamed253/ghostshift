package sim

import (
	"math"
	"testing"
)

func TestArrivalTimes_SameSeedSameArrivals(t *testing.T) {
	a := arrivalTimes(30, 480, NewStreams(42).Arrivals)
	b := arrivalTimes(30, 480, NewStreams(42).Arrivals)
	if len(a) != len(b) {
		t.Fatalf("lengths differ: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("arrival %d differs: %v vs %v", i, a[i], b[i])
		}
	}
}

func TestArrivalTimes_IncreasingAndInsideHorizon(t *testing.T) {
	const horizon = 480.0
	times := arrivalTimes(30, horizon, NewStreams(7).Arrivals)
	prev := 0.0
	for i, tm := range times {
		if tm <= prev || tm >= horizon {
			t.Fatalf("arrival %d at %v: want %v < t < %v", i, tm, prev, horizon)
		}
		prev = tm
	}
}

func TestArrivalTimes_MeanGapMatchesRate(t *testing.T) {
	const ordersPerHour = 30.0 // expected mean gap = 60/30 = 2 minutes
	times := arrivalTimes(ordersPerHour, 100_000, NewStreams(1).Arrivals)
	meanGap := times[len(times)-1] / float64(len(times))
	want := 60.0 / ordersPerHour
	if math.Abs(meanGap-want)/want > 0.02 {
		t.Fatalf("mean gap = %.4f min, want %.4f within 2%%", meanGap, want)
	}
}

func TestArrivalTimes_ZeroRateOrDuration(t *testing.T) {
	r := NewStreams(1).Arrivals
	if got := arrivalTimes(0, 60, r); len(got) != 0 {
		t.Fatalf("zero rate: got %d arrivals, want 0", len(got))
	}
	if got := arrivalTimes(30, 0, r); len(got) != 0 {
		t.Fatalf("zero duration: got %d arrivals, want 0", len(got))
	}
}