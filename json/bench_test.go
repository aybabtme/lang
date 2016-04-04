package jsonbench

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/aybabtme/lang/json/closures"
	"github.com/aybabtme/lang/json/funcopts"
)

func BenchmarkFuncOpts(b *testing.B) {
	buf := bytes.NewBuffer(nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		funcopts.Fprint(buf,
			funcopts.Object(
				funcopts.Key("hello", funcopts.Null()),
				funcopts.Key("world", funcopts.Number(1)),
				funcopts.Key("lol", funcopts.Array(
					funcopts.String("derp"),
					funcopts.Number(2),
				)),
			),
		)
	}
}

func BenchmarkClosures(b *testing.B) {
	buf := bytes.NewBuffer(nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()

		closures.Generate(func(s string) {
			buf.WriteString(s)
		}).Object(func(obj closures.Object) {
			obj.Key("hello").Null()
			obj.Key("world").Number(1)
			obj.Key("lol").Array(func(arr closures.Array) {
				arr.String("derp")
				arr.Number(2)
			})
		})
	}
}

func BenchmarkEncodingJSON(b *testing.B) {
	type obj struct {
		Hello *string       `json:"hello"`
		World float64       `json:"world"`
		Lol   []interface{} `json:"lol"`
	}
	buf := bytes.NewBuffer(nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		json.NewEncoder(buf).Encode(&obj{
			Hello: nil,
			World: 1,
			Lol: []interface{}{
				"derp", 2,
			},
		})
	}
}
