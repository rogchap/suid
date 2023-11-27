package suid

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"
)

type Suid struct {
	opts options

	dict    []rune
	counter int
}

func New(opt ...Option) *Suid {
	opts := defaultOptions
	for _, o := range opt {
		o.apply(&opts)
	}
	s := &Suid{
		opts: opts,
	}

	s.setDict(dictByName(s.opts.namedDict))

	fmt.Printf("%+q\n", s.dict)

	return s
}

func (s *Suid) Rnd() (string, error) {
	var id strings.Builder

	dictLen := len(s.dict)
	if dictLen <= 0 {
		panic("suid: dict is <= 0")
	}

	for i := 0; i < s.opts.length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(dictLen-1)))
		if err != nil {
			return "", err
		}
		if _, err := id.WriteRune(s.dict[n.Int64()]); err != nil {
			return "", err
		}
	}

	return id.String(), nil
}

func (s *Suid) Seq() string {
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

func (s *Suid) setDict(d dict) {
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
