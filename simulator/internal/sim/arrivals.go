package sim

func GenerateArrivals(cfg Config) []Order {
	order := []Order{}
	// Calculate the total number of orders based on the simulation duration and order rate
	for i := 0; ; i++ {
		interval := 60 / cfg.OrdersPerHour
		arrivalTime := float64(i) * interval
		if arrivalTime >= cfg.DurationMinutes {
			break
		}
		order = append(order, Order{
			ID:            i + 1,
			ArrivalTime:   arrivalTime,
			PrepStartTime: TimeUnset,
			PrepEndTime:   TimeUnset,
			Status:        Waiting,
		})
	}
	return order
}
