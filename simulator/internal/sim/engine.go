package sim

// makeCooks creates the kitchen staff. Everyone is free at minute 0.
func makeCooks(count int) []Cook {
	cooks := []Cook{}
	for i := 0; i < count; i++ {
		cooks = append(cooks, Cook{
			ID:              i + 1,
			AvailableAt:     0,
			OrdersCompleted: 0,
		})
	}
	return cooks
}

// earliestCook returns the INDEX of the cook who becomes free soonest.
// It returns an index, not a Cook, because the caller needs to modify
// the real cook in the slice — a returned Cook would be a copy.
func earliestCook(cooks []Cook) int {
	best := 0
	for i := range cooks {
		if cooks[i].AvailableAt < cooks[best].AvailableAt {
			best = i
		}
	}
	return best
}

// Run executes one simulation and returns every order with its
// preparation times filled in.
func Run(cfg Config) ([]Order, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	orders := GenerateArrivals(cfg)
	cooks := makeCooks(cfg.Cooks)

	// Note: `for i := range orders`, not `for _, o := range orders`.
	// Range gives copies. We need to modify the real orders.
	for i := range orders {
		c := earliestCook(cooks)

		// Prep starts at whichever is later: the order arriving,
		// or the cook finishing their previous order.
		start := max(orders[i].ArrivalTime, cooks[c].AvailableAt)
		end := start + cfg.PrepTimeMinutes

		orders[i].PrepStartTime = start
		orders[i].PrepEndTime = end
		orders[i].Status = Completed

		cooks[c].AvailableAt = end
		cooks[c].OrdersCompleted++
	}

	return orders, nil
}
