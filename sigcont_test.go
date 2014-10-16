package sigcont

import (
	"os"
	"syscall"
	"testing"
	"time"
)

func TestSIGCONTNotify(t *testing.T) {

	ch := make(chan os.Signal)
	Notify(ch)

	select {
	case <-ch:
		t.Error("didn't receive sigcont yet but library thinks we did.")
	default:
	}

	syscall.Kill(0, syscall.SIGCONT)

	select {
	case <-ch:
	case <-time.After(100 * time.Millisecond):
		t.Error("should have got SIGCONT.")
	}

}
