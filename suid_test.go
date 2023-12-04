package suid_test

import (
	"fmt"
	"testing"

	"rogchap.com/suid"
)

func TestRnd(t *testing.T) {
	tests := [...]struct {
		desc   string
		uid    *suid.SUID
		count  int
		assert func(t *testing.T, uid *suid.SUID, res []string)
	}{
		{
			desc:  "an identifer should be generated",
			uid:   suid.New(),
			count: 1,
			assert: func(t *testing.T, uid *suid.SUID, res []string) {
				if len(res) == 0 {
					t.Fatal("unexpected number of results, exp 1, got 0")
				}
				if len(res[0]) == 0 {
					t.Fatal("Generated ID is of zero length")
				}
			},
		},
		{
			desc:  "default identifier should be 6 characters in length",
			uid:   suid.New(),
			count: 4,
			assert: func(t *testing.T, uid *suid.SUID, res []string) {
				for _, r := range res {
					if len(r) != 6 {
						t.Fatalf("identifier is an unexpected length: exp 6, got %d", len(r))
					}
				}
			},
		},
		{
			desc:  "identifers should be unique",
			uid:   suid.New(),
			count: 25_000, // see suid.MaxBeforeCollision()
			assert: func(t *testing.T, uid *suid.SUID, res []string) {
				uniques := make(map[string]struct{})
				var duplicates []string
				for _, v := range res {
					if _, ok := uniques[v]; ok {
						duplicates = append(duplicates, v)
					}
					uniques[v] = struct{}{}
				}

				if len(uniques) != len(res) {
					t.Fatalf("an identifier that was generated was not unique: %+q", duplicates)
				}
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			var res []string
			for i := 0; i < tt.count; i++ {
				res = append(res, tt.uid.Rnd())
			}

			tt.assert(t, tt.uid, res)
		})
	}
}

func TestSuid(t *testing.T) {
	t.Parallel()

	// s := suid.New(suid.Len(16))
	s := suid.New()

	fmt.Printf("%.20f\n", s.AvailableIDs())
	fmt.Printf("%.20f\n", s.MaxBeforeCollision())
	fmt.Printf("%.20f\n", s.CollisionProbability())
}
