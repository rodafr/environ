// Package environ provides the value-or-default and value-or-error loaders and their methods.
package environ

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type LookupFunc func(string) (string, bool)

type Loader struct {
	lf   LookupFunc
	errs []error
}

func NewLoader(lf LookupFunc) *Loader {
	return &Loader{lf: lf}
}

func (l *Loader) Err() error { return errors.Join(l.errs...) }

func (l *Loader) load[T any](key, kind string, def []T, parse func(string) (T, error)) T {
	var zero T
	if s, ok := l.lf(key); ok && s != "" {
		v, err := parse(s)
		if err != nil {
			l.errs = append(l.errs, fmt.Errorf("parse %s env var %q: %w", kind, key, err))
			return zero
		}
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	l.errs = append(l.errs, fmt.Errorf("required %s env var %q not set", kind, key))
	return zero
}

func (l *Loader) Str(key string, def ...string) string {
	return l.load(key, "string", def, func(s string) (string, error) { return s, nil })
}

func (l *Loader) Bool(key string, def ...bool) bool {
	return l.load(key, "bool", def, strconv.ParseBool)
}

func (l *Loader) Dur(key string, def ...time.Duration) time.Duration {
	return l.load(key, "duration", def, time.ParseDuration)
}

func (l *Loader) Secs(key string, def ...time.Duration) time.Duration {
	return l.load(key, "duration as seconds", def, func(s string) (time.Duration, error) { return time.ParseDuration(s + "s") })
}

// TODO: make this work
func (l *Loader) URL(key string, def ...*url.URL) *url.URL {
	return l.load(key, "url", def, url.Parse)
}

func MapLookup(m map[string]string) LookupFunc {
	return func(k string) (string, bool) {
		v, ok := m[k]
		return v, ok
	}
}
