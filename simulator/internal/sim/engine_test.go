package sim

import (
	"math"
	"testing"
)

func TestEmptyOrderSlice(t *testing.T) {
	cfg := Config{DurationMinutes: 60, OrdersPerHour: 6, Cooks: 1, PrepTimeMinutes: 10}

	result := computeMetrics(cfg, []Order{})

	if result.OrdersReceived != 0 {
		t.Errorf("received: got %d, want 0", result.OrdersReceived)
	}
	if math.IsNaN(result.AverageWaitMin) {
		t.Error("average wait is NaN — a division was not guarded")
	}
}
