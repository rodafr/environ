package environ

import (
	"testing"
	"time"
)

var testMap = map[string]string{
	"SET":     "true",
	"INVALID": "ture",
	"EMPTY":   "",
}

func TestBoolLoader(t *testing.T) {
	testCases := []struct {
		desc    string
		key     string
		defVal  bool
		want    bool
		wantErr bool
	}{
		{
			desc:   "value found with default",
			key:    "SET",
			defVal: false,
			want:   true,
		},
		{
			desc: "value found without default",
			key:  "SET",
			want: true,
		},
		{
			desc:    "value found but invalid with default",
			key:     "INVALID",
			defVal:  false,
			wantErr: true,
		},
		{
			desc:    "value found but invalid without default",
			key:     "INVALID",
			wantErr: true,
		},
		{
			desc:   "value found but empty with default",
			key:    "EMPTY",
			defVal: true,
			want:   true,
		},
		{
			desc:    "value found but empty without default",
			key:     "EMPTY",
			wantErr: true,
		},
		{
			desc:   "value not found with default",
			key:    "VOID",
			defVal: true,
			want:   true,
		},
		{
			desc:    "value not found without default",
			key:     "VOID",
			wantErr: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			load := NewLoader(MapLookup(testMap))
			var got bool

			if tC.defVal == false {
				got = load.Bool(tC.key)
			} else {
				got = load.Bool(tC.key, tC.defVal)
			}

			if tC.wantErr {
				if err := load.Err(); err != nil {
					return
				} else {
					t.Fatalf("got no error, wanted one")
				}
			}

			if got != tC.want {
				t.Fatalf("got %v, want: %v", got, tC.want)
			}

			if err := load.Err(); err != nil {
				t.Fatalf("got error %s, wanted no error ", err.Error())
			}
		})
	}
}

func TestStringLoader(t *testing.T) {
	testCases := []struct {
		desc    string
		key     string
		defVal  string
		want    string
		wantErr bool
	}{
		{
			desc:   "value found with default",
			key:    "SET",
			defVal: "def",
			want:   "true",
		},
		{
			desc: "value found without default",
			key:  "SET",
			want: "true",
		},
		{
			desc:   "value found but empty with default",
			key:    "EMPTY",
			defVal: "def",
			want:   "def",
		},
		{
			desc:    "value found but empty without default",
			key:     "EMPTY",
			defVal:  "",
			want:    "",
			wantErr: true,
		},
		{
			desc:   "value not found with default",
			key:    "VOID",
			defVal: "def",
			want:   "def",
		},
		{
			desc:    "value not found without default",
			key:     "VOID",
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			load := NewLoader(MapLookup(testMap))
			var got string

			if tC.defVal == "" {
				got = load.Str(tC.key)
			} else {
				got = load.Str(tC.key, tC.defVal)
			}

			if tC.wantErr {
				if err := load.Err(); err != nil {
					return
				} else {
					t.Fatalf("got no error, wanted one")
				}
			}

			if got != tC.want {
				t.Fatalf("got %s, want: %s", got, tC.want)
			}

			if err := load.Err(); err != nil {
				t.Fatalf("got error %s, wanted no error ", err.Error())
			}
		})
	}
}

func TestDurLoader(t *testing.T) {
	durMap := map[string]string{
		"SET":     "33s",
		"INVALID": "bullshit",
		"EMPTY":   "",
	}

	testCases := []struct {
		desc    string
		key     string
		defVal  time.Duration
		want    time.Duration
		wantErr bool
	}{
		{
			desc:   "value found with default",
			key:    "SET",
			defVal: MustParseDuration(t, "10s"),
			want:   MustParseDuration(t, "33s"),
		},
		{
			desc: "value found without default",
			key:  "SET",
			want: MustParseDuration(t, "33s"),
		},
		{
			desc:    "value found but invalid with default",
			key:     "INVALID",
			defVal:  MustParseDuration(t, "77s"),
			wantErr: true,
		},
		{
			desc:    "value found but invalid without default",
			key:     "INVALID",
			wantErr: true,
		},
		{
			desc:   "value found but empty with default",
			key:    "EMPTY",
			defVal: MustParseDuration(t, "1s"),
			want:   MustParseDuration(t, "1s"),
		},
		{
			desc:    "value found but empty without default",
			key:     "EMPTY",
			wantErr: true,
		},
		{
			desc:   "value not found with default",
			key:    "VOID",
			defVal: MustParseDuration(t, "83s"),
			want:   MustParseDuration(t, "83s"),
		},
		{
			desc:    "value not found without default",
			key:     "VOID",
			wantErr: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			load := NewLoader(MapLookup(durMap))
			var got time.Duration

			if tC.defVal == 0 {
				got = load.Dur(tC.key)
			} else {
				got = load.Dur(tC.key, tC.defVal)
			}

			if tC.wantErr {
				if err := load.Err(); err != nil {
					return
				} else {
					t.Fatalf("got no error, wanted one")
				}
			}

			if got != tC.want {
				t.Fatalf("got %v, want: %v", got, tC.want)
			}

			if err := load.Err(); err != nil {
				t.Fatalf("got error %s, wanted no error ", err.Error())
			}
		})
	}
}

func MapLookup(m map[string]string) LookupFunc {
	return func(k string) (string, bool) {
		v, ok := m[k]
		return v, ok
	}
}

func MustParseDuration(t *testing.T, s string) time.Duration {
	t.Helper()
	dur, _ := time.ParseDuration(s)
	return dur
}
