package suid

type options struct {
	length    int
	counter   int
	namedDict namedDict
}

var defaultOptions = options{
	length:    6,
	counter:   0,
	namedDict: DictAlphaNum,
}

// Option sets an option on the SUID, such as ID length and the dictionary/characters used.
type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (f *funcOption) apply(o *options) {
	f.f(o)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// Dict is an option to sets the dictionary (aka characters) used in generating the ID,
// to one of the named dictionaries; default: DictAlphaNum
func Dict(d namedDict) Option {
	return newFuncOption(func(o *options) {
		o.namedDict = d
	})
}

// Len is an option to set the length of the ID to generate; default: 6
// 6 was chosen as the default ID length as it will provide provide millions of IDs
// with a very low probability of producing a duplicate.
func Len(l int) Option {
	return newFuncOption(func(o *options) {
		o.length = l
	})
}

// Counter is an option to set the starting counter for sequential IDs; default: 0
func Counter(c int) Option {
	return newFuncOption(func(o *options) {
		o.counter = c
	})
}
