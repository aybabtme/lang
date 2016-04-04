package funcopts

import (
	"bytes"
	"os"
	"testing"
)

func ExampleObject() {
	Fprint(os.Stdout,
		Object(
			Key("hello", Null()),
			Key("world", Number(1)),
			Key("lol", Array(
				String("derp"),
				Number(2),
			)),
		),
	)
	// Output:
	// {"hello":null,"world":1,"lol":["derp",2]}
}

func BenchmarkObject(b *testing.B) {
	buf := bytes.NewBuffer(nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		Fprint(buf,
			Object(
				Key("hello", Null()),
				Key("world", Number(1)),
				Key("lol", Array(
					String("derp"),
					Number(2),
				)),
			),
		)
	}
}

func TestSpec(t *testing.T) {
	tests := []struct {
		want string
		call value
	}{

		{want: `1234.56789`,
			call: Number(1234.56789)},

		{want: `"lol"`,
			call: String("lol")},

		// object
		{want: `{"hello":"world","bonjour":"le monde","derp":{"lol":null}}`,
			call: Object(
				Key("hello", String("world")),
				Key("bonjour", String("le monde")),
				Key("derp", Object(Key("lol", Null()))),
			),
		},

		// array
		{want: `[]`,
			call: Array()},

		{want: `[{},[],[[]]]`,
			call: Array(
				Object(),
				Array(),
				Array(Array()),
			)},

		// value
		{want: `1234.56789`,
			call: Number(1234.56789)},
		{want: `"lol"`,
			call: String("lol")},
		{want: `null`,
			call: Null()},
		{want: `true`,
			call: True()},
		{want: `false`,
			call: False()},
	}

	for i, tt := range tests {
		t.Logf("test %d", i)

		w := new(bytes.Buffer)
		tt.call(w)

		if want, got := tt.want, w.String(); want != got {
			t.Errorf("want=%q", want)
			t.Errorf(" got=%q", got)
		}
	}
}
