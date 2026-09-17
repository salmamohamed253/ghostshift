package sim

// Cook represents a single cook in the kitchen to prepare the order
type Cook struct {
	ID              int
	AvailableAt     float64
	OrdeerCompleted int
}
