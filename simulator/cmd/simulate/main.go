package main

import (
	"flag"
	"fmt"

	"github.com/salmamohamed253/ghostshift/simulator/internal/sim"
)

func main() {
	duration := flag.Float64("duration", 240, "Duration of the simulation in minutes")
	ordersPerHour := flag.Float64("orders", 30, "Number of orders per hour")
	cooks := flag.Int("cooks", 3, "Number of cooks in the kitchen")
	prepTime := flag.Float64("prep", 12, "Preparation time for each order in minutes")
	flag.Parse()
	result, err := sim.Run(sim.Config{
		DurationMinutes: *duration,
		OrdersPerHour:   *ordersPerHour,
		Cooks:           *cooks,
		PrepTimeMinutes: *prepTime,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%+v\n", result)
}
