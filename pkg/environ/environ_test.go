package environ

import (
	"testing"
)

func MapLookup(m map[string]string) LookupFunc {
	return func(k string) (string, bool) {
		v, ok := m[k]
		return v, ok
	}
}

var testMap = map[string]string{
	"SET":   "true",
	"EMPTY": "",
}

func TestBoolLoader(t *testing.T) {
	testCases := []struct {
		desc   string
		key    string
		defVal bool
		want   bool
	}{
		{
			desc:   "value found",
			key:    "SET",
			defVal: false,
			want:   true,
		},
		{
			desc:   "value found but empty with default",
			key:    "EMPTY",
			defVal: true,
			want:   true,
		},
		{
			desc:   "value found but empty without default",
			key:    "EMPTY",
			defVal: false,
			want:   false,
		},
		{
			desc:   "value not found with default",
			key:    "VOID",
			defVal: true,
			want:   true,
		},
		{
			desc:   "value not found without default",
			key:    "VOID",
			defVal: false,
			want:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			load := NewLoader(MapLookup(testMap))
			got := load.Bool(tC.key, tC.defVal)

			if got != tC.want {
				t.Fatalf("got %v, want: %v", got, tC.want)
			}
		})
	}
}

func TestReqBoolLoader(t *testing.T) {
	testCases := []struct {
		desc    string
		key     string
		want    bool
		wantErr bool
	}{
		{
			desc: "value found",
			key:  "SET",
			want: true,
		},
		{
			desc:    "value not found",
			key:     "VOID",
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			load := NewLoader(MapLookup(testMap))

			got := load.ReqBool(tC.key)

			if tC.wantErr {
				if err := load.Error(); err == nil {
					t.Fatalf("got no error, wanted one")
				}
			}

			if got != tC.want {
				t.Fatalf("got %v, want: %v", got, tC.want)
			}
		})
	}
}

func TestStringLoader(t *testing.T) {
	testCases := []struct {
		desc   string
		key    string
		defVal string
		want   string
	}{
		{
			desc:   "value found and not empty",
			key:    "SET",
			defVal: "def",
			want:   "true",
		},
		{
			desc:   "value found but empty",
			key:    "EMPTY",
			defVal: "def",
			want:   "def",
		},
		{
			desc:   "value not found with default",
			key:    "VOID",
			defVal: "def",
			want:   "def",
		},
		{
			desc:   "value not found without default",
			key:    "VOID",
			defVal: "",
			want:   "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			load := NewLoader(MapLookup(testMap))
			got := load.Str(tC.key, tC.defVal)

			if got != tC.want {
				t.Fatalf("got %s, want: %s", got, tC.want)
			}
		})
	}
}
