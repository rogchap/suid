package suid_test

import (
	"testing"

	"rogchap.com/suid"
)

func TestSuidRnd(t *testing.T) {
	s := suid.New(suid.Dict(suid.DictNum))
	println(s.Rnd())
}

func TestSuidSeq(t *testing.T) {
	s := suid.New()
	println(s.Seq())
}
