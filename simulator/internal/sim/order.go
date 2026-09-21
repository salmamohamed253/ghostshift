package sim

// marks timestamp that ahs not yet occured to avoid future bugs and misconseptions
const TimeUnset float64 = -1

// Order represents a single customer order moving through the kitchen
type Order struct {
	ID            int
	ArrivalTime   float64
	PrepStartTime float64
	PrepEndTime   float64
	Status        Status
}
type Status string

const (
	Preparing Status = "PREPARING"
	Waiting   Status = "WAITING"
	Completed Status = "COMPLETED"
)
