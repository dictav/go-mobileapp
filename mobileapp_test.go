package mobileapp

import (
	"testing"
)

var tests = []struct {
	in   string
	ok   bool
	kind Kind
	id   string
}{
	{"mobileapp::0", false, Unknown, ""},
	{"mobileapp::0-1234567", false, Unknown, ""},
	{"mobileapp::0-com.example.app", false, Unknown, ""},
	{"mobileapp:1-1234567", false, Unknown, ""},
	{"mobileapp:2-com.example.app", false, Unknown, ""},
	{"mobileapp::1-1234567", true, IOS, "1234567"},
	{"mobileapp::2-com.example.app", true, Android, "com.example.app"},
	{"mobileapp::3-windows-phone?", true, WindowsPhone, "windows-phone?"},
}

func Test_parseRegexp(t *testing.T) {
	for _, v := range tests {
		t.Run(v.in, func(t *testing.T) {
			app, ok := parseRegexp(v.in)
			if v.ok != ok {
				t.Errorf("want %t, got %t", v.ok, ok)
				return
			}

			if !v.ok {
				return
			}

			if v.kind != app.Kind {
				t.Errorf("want %q, got %q", v.kind, app.Kind)
			}

			if v.id != app.ID {
				t.Errorf("want %q, got %q", v.id, app.ID)
			}
		})
	}
}

func Test_parseSplit(t *testing.T) {
	for _, v := range tests {
		t.Run(v.in, func(t *testing.T) {
			app, ok := parseSplit(v.in)
			if v.ok != ok {
				t.Errorf("want %t, got %t, %#v, %#v", v.ok, ok, app, IOS)
				return
			}

			if !v.ok {
				return
			}

			if v.kind != app.Kind {
				t.Errorf("want %q, got %q", v.kind, app.Kind)
			}

			if v.id != app.ID {
				t.Errorf("want %q, got %q", v.id, app.ID)
			}
		})
	}
}

func Benchmark_paraseRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			_, _ = parseRegexp(v.in)
		}
	}
}

func Benchmark_parseSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			_, _ = parseSplit(v.in)
		}
	}
}

var tests2 = []struct {
	in   string
	kind Kind
}{
	// FIXME: Unknown case
	{"123456z", Android},
	// FIXME: Unknown case
	{"a123456", Android},
	{"123456", IOS},
	{"a12345.example.app", Android},
}

func Benchmark_detectAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range tests2 {
			if v.kind != detectAtoi(v.in) {
				b.Errorf("not match")
			}
		}
	}
}

func Benchmark_detectScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range tests2 {
			if v.kind != detectScan(v.in) {
				b.Errorf("not match")
			}
		}
	}
}

func Benchmark_detectBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range tests2 {
			if v.kind != detectBytes(v.in) {
				b.Errorf("not match")
			}
		}
	}
}
