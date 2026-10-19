# environ

Load environment variables with optional defaults

## Usage

```go
// create loader with wanted lookup func:
load := NewLoader(os.LookupEnv)

// get value or default:
value := load.Str("KEY", "default_value")

// get value or error:
value := load.ReqStr("KEY")
if err := load.Err(); err != nil {
    return err
}

// use in a config struct assignment:
cfg := Config{
 Host: load.Str("HOST", "localhost"),
 Port: load.Str("PORT", "8080"),
 DBURL:  load.ReqStr("DATABASE_URL"),
}
if err := l.Err(); err != nil {
 return err // lists all missing vars at once
}
```
