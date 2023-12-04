package suid

type namedDict int32

const (
	DictNum namedDict = iota
	DictAlpha
	DictAlphaLower
	DictAlphaUpper
	DictAlphaNum
	DictAlphaNumLower
	DictAlphaNumUpper
	DictHex
)

type dict [][2]rune

func dictByName(d namedDict) dict {
	digits := [2]rune{'0', '9'}
	lower := [2]rune{'a', 'z'}
	upper := [2]rune{'A', 'Z'}
	hex := [2]rune{'a', 'f'}

	switch d {
	default:
		fallthrough
	case DictNum:
		return dict{digits}
	case DictAlpha:
		return dict{lower, upper}
	case DictAlphaLower:
		return dict{lower}
	case DictAlphaUpper:
		return dict{upper}
	case DictAlphaNum:
		return dict{digits, lower, upper}
	case DictAlphaNumLower:
		return dict{digits, lower}
	case DictAlphaNumUpper:
		return dict{digits, upper}
	case DictHex:
		return dict{digits, hex}
	}
}
