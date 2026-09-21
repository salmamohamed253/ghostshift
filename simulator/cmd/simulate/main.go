package main

import (
	"fmt"

	"github.com/salmamohamed253/ghostshift/simulator/internal/sim"
)

func main() {
	orders, err := sim.Run(sim.Config{
		DurationMinutes: 60,
		OrdersPerHour:   6,
		Cooks:           2,
		PrepTimeMinutes: 15,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, o := range orders {
		fmt.Printf("order %d: arrived %.1f, started %.1f, wait %.1f\n",
			o.ID, o.ArrivalTime, o.PrepStartTime, o.PrepStartTime-o.ArrivalTime)
	}
}
