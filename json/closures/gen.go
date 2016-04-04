package closures

import "strconv"

type Input func(string)

func Generate(dst Input) Spec {
	return genValue(dst)
}

type Spec Value

type Object interface {
	Key(string) Value
}

type Array Value

type Value interface {
	String(string)
	Number(float64)
	Object(func(Object))
	Array(func(Array))
	True()
	False()
	Null()
}

type generator Input

func (gen generator) Value(fn func(Value)) { fn(genValue(gen)) }
func (gen generator) String(s string)      { gen(asString(s)) }
func (gen generator) Number(v float64)     { gen(asNumber(v)) }
func (gen generator) Object(fn func(Object)) {
	gen("{")
	fn(&genObject{input: Input(gen)})
	gen("}")
}
func (gen generator) Array(fn func(Array)) {
	gen("[")
	fn(&genArray{input: genValue(gen)})
	gen("]")
}

type genObject struct {
	input    Input
	notFirst bool
}

func (gen *genObject) Key(key string) Value {
	if gen.notFirst {
		gen.input(",")
	} else {
		gen.notFirst = true
	}
	gen.input(strconv.Quote(key))
	gen.input(":")
	return genValue(gen.input)
}

type genArray struct {
	input    genValue
	notFirst bool
}

func (gen *genArray) pre() {
	if gen.notFirst {
		gen.input(",")
	} else {
		gen.notFirst = true
	}
}

func (gen *genArray) String(s string)        { gen.pre(); gen.input.String(s) }
func (gen *genArray) Number(v float64)       { gen.pre(); gen.input.Number(v) }
func (gen *genArray) True()                  { gen.pre(); gen.input.True() }
func (gen *genArray) False()                 { gen.pre(); gen.input.False() }
func (gen *genArray) Null()                  { gen.pre(); gen.input.Null() }
func (gen *genArray) Object(fn func(Object)) { gen.pre(); gen.input.Object(fn) }
func (gen *genArray) Array(fn func(Array))   { gen.pre(); gen.input.Array(fn) }

type genValue Input

func (gen genValue) String(s string)  { gen(asString(s)) }
func (gen genValue) Number(v float64) { gen(asNumber(v)) }
func (gen genValue) True()            { gen("true") }
func (gen genValue) False()           { gen("false") }
func (gen genValue) Null()            { gen("null") }
func (gen genValue) Object(fn func(Object)) {
	gen("{")
	fn(&genObject{input: Input(gen)})
	gen("}")
}
func (gen genValue) Array(fn func(Array)) {
	gen("[")
	fn(&genArray{input: genValue(gen)})
	gen("]")
}

func asString(s string) string  { return strconv.Quote(s) }
func asNumber(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }
