// Package environ provides the value-or-default and value-or-error loaders and their methods.
package environ

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type LookupFunc func(string) (string, bool)

type Loader struct {
	lf   LookupFunc
	errs []error
}

func NewLoader(lf LookupFunc) Loader {
	return Loader{lf: lf}
}

func (l *Loader) Error() error { return errors.Join(l.errs...) }

func (l *Loader) Str(key, defVal string) string {
	if val, ok := l.lf(key); ok && val != "" {
		return val
	}
	return defVal
}

func (l *Loader) ReqStr(key string) string {
	if val, ok := l.lf(key); ok && val != "" {
		return val
	}
	l.errs = append(l.errs, fmt.Errorf("required string env var %q not set", key))
	return ""
}

func (l *Loader) Bool(key string, defVal bool) bool {
	if val, ok := l.lf(key); ok && val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defVal
}

func (l *Loader) ReqBool(key string) bool {
	if val, ok := l.lf(key); ok && val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	l.errs = append(l.errs, fmt.Errorf("required boolean env var %q not set", key))
	return false
}

func (l *Loader) Dur(key string, defVal time.Duration) time.Duration {
	if val, ok := l.lf(key); ok && val != "" {
		if b, err := time.ParseDuration(val); err == nil {
			return b
		}
	}
	return defVal
}

func (l *Loader) ReqDur(key string) time.Duration {
	if val, ok := l.lf(key); ok && val != "" {
		if b, err := time.ParseDuration(val); err == nil {
			return b
		}
	}
	l.errs = append(l.errs, fmt.Errorf("required time duration env var %q not set", key))
	return 0
}
