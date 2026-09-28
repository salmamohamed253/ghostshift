package sim

import "math/rand/v2"

// Streams holds one independent random number generator per source of
// randomness in the simulation. Each source gets its own generator so that
// drawing from one (e.g. prep times) never shifts the sequence of another
// (e.g. arrivals). This keeps demand identical when only staffing changes.
type Streams struct {
	Arrivals *rand.Rand
	PrepTime *rand.Rand
	Delivery *rand.Rand
}

// Fixed stream identifiers. Never renumber these: changing one changes
// every result ever produced for a given seed.
const (
	streamArrivals uint64 = 1
	streamPrepTime uint64 = 2
	streamDelivery uint64 = 3
)

// NewStreams derives all generators from a single master seed, so the user
// only ever supplies one number (--seed 42) to reproduce a whole run.
func NewStreams(seed uint64) Streams {
	return Streams{
		Arrivals: newStream(seed, streamArrivals),
		PrepTime: newStream(seed, streamPrepTime),
		Delivery: newStream(seed, streamDelivery),
	}
}

func newStream(seed, id uint64) *rand.Rand {
	return rand.New(rand.NewPCG(seed, id))
}