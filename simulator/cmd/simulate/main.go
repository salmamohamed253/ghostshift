package main

import (
	"fmt"

	"github.com/salmamohamed253/ghostshift/simulator/internal/sim"
)

func main() {
	orders := sim.GenerateArrivals(sim.Config{
		DurationMinutes: 50,
		OrdersPerHour:   7,
		Cooks:           2,
		PrepTimeMinutes: 10,
	})

	fmt.Println("total orders:", len(orders))
	for _, o := range orders[:5] {
		fmt.Println(o.ID, o.ArrivalTime, o.Status)
	}
}
