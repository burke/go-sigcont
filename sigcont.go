package sigcont

/*
extern int RegisterSIGCONTHandler();
extern int WaitForSIGCONT();
*/
import "C"

import (
	"os"
	"syscall"
)

func Notify(sigch chan os.Signal) {
	if ret := int(C.RegisterSIGCONTHandler()); ret < 0 {
		panic("already registered SIGCONT handler. You can only call this once.")
	}

	go func() {
		for {
			ret := C.WaitForSIGCONT()
			if ret < 0 {
				// error
				continue
			}
			sigch <- syscall.SIGCONT
		}
	}()

}
