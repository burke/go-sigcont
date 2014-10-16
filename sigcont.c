#include "_cgo_export.h"

#include <signal.h>
#include <unistd.h>

static int received_sigcont = 0;

int FetchSIGCONTStatus()
{
  int cnt = received_sigcont;
  if (cnt > 0)
    received_sigcont = 0;
  return cnt;
}

static void sigcont_handler(int a)
{
  received_sigcont = 1;
}

void RegisterSIGCONTHandler()
{
	signal(SIGCONT, sigcont_handler);
}


