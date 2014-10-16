# go-sigcont

Go, for reasons unbeknownst to me, can't trap `SIGCONT`. This library allows
you to trap `SIGCONT`. It uses `cgo`.

Go gets unhappy if you call into a Go function from a signal handler, so the
implementation creates a pipe, writes to it from the handler, and spawns a
goroutine to push a `syscall.SIGCONT` to the channel each time data is received
on the pipe.

## Usage:

```go
import "github.com/burke/go-sigcont"
import "os"

// ...
ch := make(chan os.Signal)
sigcont.Notify(ch)
```

You can combine this with a normal `signal.Notify` too.

```go
ch := make(chan os.Signal)
signal.Notify(ch, syscall.SIGTERM)
sigcont.Notify(ch)
```
