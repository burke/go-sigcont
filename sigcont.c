#include "_cgo_export.h"

#include <signal.h>
#include <unistd.h>
#include <sys/types.h>

static int registered = 0;
static int fd[2];
static char whatever = 'x';

static void sigcont_handler(int a)
{
  write(fd[1], &whatever, 1);
}

int RegisterSIGCONTHandler()
{
  if (registered > 0) {
    return -1;
  }
  registered = 1;
  pipe(fd);
	signal(SIGCONT, sigcont_handler);
  return 0;
}

int WaitForSIGCONT()
{
  fd_set set;
  int ret;

  /* Initialize the file descriptor set. */
  FD_ZERO(&set);
  FD_SET(fd[0], &set);

  if ((ret = select(fd[0]+1, &set, NULL, NULL, NULL)) < 0) {
    return -1; // error
  } else {
    read(fd[0], &whatever, 1);
    return 0; //sigcont received
  }
}
