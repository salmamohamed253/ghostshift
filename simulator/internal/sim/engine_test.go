package sim

import (
	"testing"
	"reflect"
)


func TestRun_SameSeedSameResult(t *testing.T) {
	cfg := Config{DurationMinutes: 480, OrdersPerHour: 30, Cooks: 3, PrepTimeMinutes: 5, Seed: 42}
	a, err := Run(cfg)
	if err != nil {
		t.Fatal(err)
	}
	b, err := Run(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("same seed gave different results:\n%+v\n%+v", a, b)
	}
}

func TestRun_DifferentSeedDifferentResult(t *testing.T) {
	cfg := Config{DurationMinutes: 480, OrdersPerHour: 30, Cooks: 3, PrepTimeMinutes: 5, Seed: 42}
	a, err := Run(cfg)
	if err != nil {
		t.Fatal(err)
	}
	cfg.Seed = 43
	b, err := Run(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(a, b) {
		t.Fatal("seeds 42 and 43 gave identical results; is the seed being used?")
	}
}