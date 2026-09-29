package sim
import "math/rand/v2"

func orderFromTimes(times []float64) []Order {
	orders := make([]Order, len(times))
	for i, t := range times {
		orders[i] = Order{
			ID:            i + 1,
			ArrivalTime:   t,
			PrepStartTime: TimeUnset,
			PrepEndTime:   TimeUnset,
			Status:        Waiting,
		}
	}
	return orders
}
// arrivalTimes generates customer arrival times, in minutes since opening,
// for a Poisson arrival process. Gaps between consecutive arrivals are
// exponentially distributed with mean 60/ordersPerHour minutes.
// Only arrivals strictly before durationMinutes are returned.
func arrivalTimes(ordersPerHour, durationMinutes float64, r *rand.Rand) []float64 {
	if ordersPerHour <= 0 || durationMinutes <= 0 {
		return nil
	}
	lambda := ordersPerHour / 60.0 // arrivals per minute
	var times []float64
	t := 0.0
	for {
		t += r.ExpFloat64() / lambda
		if t >= durationMinutes {
			break
		}
		times = append(times, t)
	}
	return times
}
