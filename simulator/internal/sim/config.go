package sim

// Config represents the configuration for a single simulation run
type Config struct {
	DurationMinutes float64
	OrderPerHour    float64
	Cooks           int
	PrepTimeMinutes float64
}
