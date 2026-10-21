# environ

Load environment variables with optional defaults

## Supported types

- `string`
- `bool`
- `time.Duration`
- `string` as seconds to `time.Duration`, i.e. "10" gives 10 seconds.
- WIP: `url.URL`

## Usage

```go
// create loader with wanted lookup func:
load := NewLoader(os.LookupEnv)

// get value or (optional) default:
value := load.Str("KEY", "default_value")

// get value or error:
value := load.Str("KEY")
if err := load.Err(); err != nil {
    return err
}

// use in a config struct assignment:
cfg := Config{
 UseFeature: load.Bool("FEATURE_FLAG"),
 Threads: load.Int("THREADS", "5"),
 DBURL:  load.URL("DATABASE_URL"),
}
if err := l.Err(); err != nil {
 return err // lists all missing vars at once
}
```

```go
// in testing, use a map:
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

load := NewLoader(MapLookup(testMap))
```

## TODO

- [ ] Provide convenience function to lookup from env files
