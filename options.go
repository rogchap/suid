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

func Dict(d namedDict) Option {
	return newFuncOption(func(o *options) {
		o.namedDict = d
	})
}
