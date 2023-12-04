package suid

import (
	"math"
)

func (s *SUID) AvailableIDs() float64 {
	return math.Trunc(math.Pow(float64(len(s.dict)), float64(s.opts.length)))
}

func (s *SUID) MaxBeforeCollision() float64 {
	return math.Sqrt((math.Pi / 2) * float64(s.AvailableIDs()))
}

func (s *SUID) CollisionProbability() float64 {
	return s.MaxBeforeCollision() / s.AvailableIDs()
}
