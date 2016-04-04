package closures

import (
	"bytes"
	"os"
	"testing"
)

func ExampleGenerate() {
	spec := Generate(func(s string) {
		os.Stdout.Write([]byte(s))
	})

	spec.Object(func(obj Object) {
		obj.Key("hello").Array(func(arr Array) {
			arr.Number(1)
			arr.Null()
			arr.False()
		})
		obj.Key("lol").Number(1)
	})

	// Output:
	// {"hello":[1,null,false],"lol":1}
}

func TestSpec(t *testing.T) {
	tests := []struct {
		want string
		call func(Spec)
	}{

		{want: `1234.56789`,
			call: func(spec Spec) {
				spec.Number(1234.56789)
			}},
		{want: `"lol"`,
			call: func(spec Spec) {
				spec.String("lol")
			}},

		// object
		{want: `{"hello":"world","bonjour":"le monde","derp":{"lol":null}}`,
			call: func(spec Spec) {
				spec.Object(func(obj Object) {
					obj.Key("hello").String("world")
					obj.Key("bonjour").String("le monde")
					obj.Key("derp").Object(func(sub Object) {
						sub.Key("lol").Null()
					})
				})
			}},

		// array
		{want: `[]`,
			call: func(spec Spec) {
				spec.Array(func(Array) {})
			}},
		{want: `[{},[],[[]]]`,
			call: func(spec Spec) {
				spec.Array(func(arr Array) {
					arr.Object(func(Object) {})
					arr.Array(func(Array) {})
					arr.Array(func(sub Array) {
						sub.Array(func(Array) {})
					})
				})
			}},

		// value
		{want: `1234.56789`,
			call: func(spec Spec) {
				spec.Number(1234.56789)
			}},
		{want: `"lol"`,
			call: func(spec Spec) {
				spec.String("lol")
			}},
		{want: `null`,
			call: func(spec Spec) {
				spec.Null()
			}},
		{want: `true`,
			call: func(spec Spec) {
				spec.True()
			}},
		{want: `false`,
			call: func(spec Spec) {
				spec.False()
			}},
	}

	for i, tt := range tests {
		t.Logf("test %d", i)

		w := new(bytes.Buffer)
		dst := func(in string) {
			_, _ = w.WriteString(in)
		}
		tt.call(Generate(dst))

		if want, got := tt.want, w.String(); want != got {
			t.Errorf("want=%q", want)
			t.Errorf(" got=%q", got)
		}
	}
}
