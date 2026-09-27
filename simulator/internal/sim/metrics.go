package sim

// Result holds the aggregate outcome of a single simulation run.
type Result struct {
	OrdersReceived     int
	OrdersCompleted    int
	AverageWaitMin     float64
	MaxWaitMin         float64
	KitchenUtilization float64
}

// computeMetrics folds a finished order slice down into summary numbers.
func computeMetrics(cfg Config, orders []Order) Result {
	result := Result{
		OrdersReceived: len(orders),
	}

	totalWait := 0.0

	// Value copy is fine here: we only READ. Compare with the engine
	// loop, which needs `for i := range` because it writes.
	for _, o := range orders {
		if o.Status != Completed {
			continue
		}

		result.OrdersCompleted++

		wait := o.PrepStartTime - o.ArrivalTime
		totalWait += wait

		if wait > result.MaxWaitMin {
			result.MaxWaitMin = wait
		}
	}

	// Guard every division. A 5-minute shift can produce zero orders,
	// and dividing by zero here gives NaN — which then silently spreads
	// through every downstream number without ever raising an error.
	if result.OrdersCompleted > 0 {
		result.AverageWaitMin = totalWait / float64(result.OrdersCompleted)
	}

	// Utilization = time spent cooking / total cook-time available.
	// This CAN exceed 1.0, and that is not a bug: it means cooks were
	// still working past the end of the shift, clearing a backlog.
	// Leave it uncapped — the overflow is the signal that the kitchen
	// is over capacity.
	capacity := float64(cfg.Cooks) * cfg.DurationMinutes
	if capacity > 0 {
		busy := float64(result.OrdersCompleted) * cfg.PrepTimeMinutes
		result.KitchenUtilization = busy / capacity
	}

	return result
}
