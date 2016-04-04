package funcopts

import (
	"io"
	"strconv"
)

func Fprint(w io.Writer, v value) { v(w) }

type value func(io.Writer)

func Null() value            { return func(w io.Writer) { w.Write([]byte("null")) } }
func True() value            { return func(w io.Writer) { w.Write([]byte("true")) } }
func False() value           { return func(w io.Writer) { w.Write([]byte("false")) } }
func String(s string) value  { return func(w io.Writer) { w.Write([]byte(asString(s))) } }
func Number(v float64) value { return func(w io.Writer) { w.Write([]byte(asNumber(v))) } }

func Array(elems ...value) value {
	return func(w io.Writer) {
		w.Write([]byte("["))
		for i, elem := range elems {
			if i != 0 {
				w.Write([]byte(","))
			}
			elem(w)
		}
		w.Write([]byte("]"))
	}
}

func Object(keys ...keyOpt) value {
	return func(w io.Writer) {
		w.Write([]byte("{"))
		for i, key := range keys {
			if i != 0 {
				w.Write([]byte(","))
			}
			key(w)
		}
		w.Write([]byte("}"))
	}
}

type keyOpt func(io.Writer)

func Key(key string, value value) keyOpt {
	return func(w io.Writer) {
		w.Write([]byte(asString(key)))
		w.Write([]byte(":"))
		value(w)
	}
}

func asString(s string) string  { return strconv.Quote(s) }
func asNumber(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }
