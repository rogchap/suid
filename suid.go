package suid

import (
	"crypto/rand"
	"math"
	"math/big"
	"strings"
)

// Suid represent a short unique identifier generator.
type SUID struct {
	opts options

	dict    []rune
	counter int
}

// New creates a new short unique identifier generator, with various options.
func New(opt ...Option) *SUID {
	opts := defaultOptions
	for _, o := range opt {
		o.apply(&opts)
	}

	s := &SUID{
		opts: opts,
	}
	s.setDict(dictByName(s.opts.namedDict))
	s.counter = s.opts.counter

	return s
}

// Rnd generates a random unique identifier.
// The ID can be configured by setting various options on the SUID generator.
// For example the dictionary/character set, or the generated length of the ID
// (defaults to lenght 6)
func (s *SUID) Rnd() string {
	var id strings.Builder

	dictLen := len(s.dict)
	if dictLen <= 1 {
		panic("suid: dict is <= 1")
	}

	for i := 0; i < s.opts.length; i++ {

		n, err := rand.Int(rand.Reader, big.NewInt(int64(dictLen)))
		if err != nil {
			// The error is EOF only if no bytes were read, or ErrUnexpectedEOF
			// happens after reading some but not all the bytes. We already panic
			// if there is a dict <= 1, so we should always be able to generate a rand.Int.
			panic(err)
		}
		// Euclidean modulus
		idx := n.Mod(n, big.NewInt(int64(dictLen)))
		id.WriteRune(s.dict[idx.Int64()]) // always returns a nil error, so ignoring
	}

	return id.String()
}

// Seq generates an ID based on the set dictionanry/character set and a counter,
// which increments on every Seq ID generated.
// You can set the starting counter in the options of the Suid generator.
func (s *SUID) Seq() string {
	var id strings.Builder

	dictLen := len(s.dict)
	c := s.counter

	for c != 0 {
		r := c % dictLen
		id.WriteRune(s.dict[r])
		c = int(math.Trunc(float64(c) / float64(dictLen)))
	}
	s.counter += 1

	return id.String()
}

func (s *SUID) setDict(d dict) {
	var dict []rune
	for _, rng := range d {
		for i := rng[0]; i <= rng[1]; i++ {
			dict = append(dict, i)
		}
	}

	shuffle(len(dict), func(i, j int) {
		dict[i], dict[j] = dict[j], dict[i]
	})

	s.dict = dict
}

func shuffle(n int, swap func(i, j int)) {
	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	if n < 0 {
		panic("suid: invalid argument: n < 0")
	}

	for i := n - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		swap(i, int(j.Uint64()))
	}
}
