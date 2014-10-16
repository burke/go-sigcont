# go-sigcont

Go, for reasons unbeknownst to me, can't trap `SIGCONT`. This library allows
you to trap `SIGCONT`. It uses `cgo`.

Right now it polls every 50 milliseconds for stupid reasons. I should change this.

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
