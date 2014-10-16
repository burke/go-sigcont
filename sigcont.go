package sigcont

/*
extern void RegisterSIGCONTHandler();
extern int FetchSIGCONTStatus();
*/
import "C"

import (
	"os"
	"syscall"
	"time"
)

func Notify(sigch chan os.Signal) {
	C.RegisterSIGCONTHandler()

	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			status := int(C.FetchSIGCONTStatus())
			if status > 0 {
				sigch <- syscall.SIGCONT
			}
		}
	}()

}
