package sim

import "fmt"

// Config represents the configuration for a single simulation run
type Config struct {
	DurationMinutes float64
	OrdersPerHour   float64
	Cooks           int
	PrepTimeMinutes float64
}

func (c Config) Validate() error {
	if c.DurationMinutes <= 0 {
		return fmt.Errorf("durationMinutes must be greater than 0, got %g", c.DurationMinutes)
	}
	if c.OrdersPerHour <= 0 {
		return fmt.Errorf("ordersPerHour must be greater than 0, got %g", c.OrdersPerHour)
	}
	if c.Cooks <= 0 {
		return fmt.Errorf("cooks must be greater than 0, got %d", c.Cooks)
	}
	if c.PrepTimeMinutes <= 0 {
		return fmt.Errorf("prepTimeMinutes must be greater than 0, got %g", c.PrepTimeMinutes)
	}
	return nil
}
