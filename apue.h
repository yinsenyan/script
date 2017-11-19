#ifndef _APUE_H
#define _APUE_H

#define _XOPEN_SOURCE 600

#include <sys/types.h>
#include <sys/stat.h>
#include <sys/termios.h>
#ifndef TIOCGWINSZ
#include <sys/ioctl.h>
#endif
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include <errno.h>
#include <stdarg.h>
#include <syslog.h>

#define MAXLINE 4096

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)
#define DIR_MODE (FILE_MODE | S_IXUSR | S_IXGRP | S_IXOTH)

typedef void Sigfunc(int);

#if defined(SIG_IGN) && !defined(SIG_ERR)
#define SIG_ERR ((Sigfunc *)-1)
#endif

#define min(a,b) ((a) < (b)) ? ((a) : (b))
#define max(a,b) ((a) > (b)) ? ((a) : (b))

char *path_alloc(int *);
long open_max(void);
void chr_fl(int, int);
void set_fl(int, int);
void pr_exit(int);
void pr_mask(const char *);
Sigfunc *Signal_intr(int, Sigfunc *);

int tty_cbreak(int);
int tty_raw(int);
int tty_reset(int);
void tty_atexit(void);
#ifdef ECHO
struct termios *tty_termios(void);
#endif

void sleep_us(unsigned int);
ssize_t readn(int, void *, ssize_t);
ssize_t writen(int, const void *, ssize_t);
void daemonize(const char *);

int s_pipe(int *);
int recv_fd(int, ssize_t (*func)(int, const void *, size_t));
int send_fd(int, int);
int send_err(int, int, const char *);
int serv_listen(const char *);
int serv_accept(int, uid_t *);
int cli_conn(const char *);
int buf_args(char *, int (*func)(int, char **));

int ptym_open(char *, int);
int ptys_open(char *);
#ifdef TIOCGWINSZ
pid_t pty_fork(int *, char *, const struct termios *, const struct winsize *);
#endif

int lock_reg(int, int, int, off_t, int, off_t);
#define read_lock(fd, offset, whence, len) lock_reg((fd), F_SETLK, F_RDLCK, (offset), (whence), (len))
#define readw_lock(fd, offset, whence, len) lock_reg((fd), F_SETLKW, F_RDLCK, (offset), (whence), (len))
#define write_lock(fd, offset, whence, len) lock_reg((fd), F_SETLK, F_WRLCK, (offset), (whence), (len))
#define writew_lock(fd, offset, whence, len) lock_reg((fd), F_SETLKW, F_WRLCK, (offset), (whence), (len))
#define un_lock(fd, offset, whence, len) lock_reg((fd), F_SETLK, F_UNLCK, (offset), (whence), (len))

pid_t lock_test(int, int, off_t, int, off_t);

#define is_read_lockable(fd, offset, whence, len) (lock_test((fd), F_RDLCK, (offset), (whence), (len)) == 0)
#define is_write_lockable(fd, offset, whence, len) (lock_test((fd), F_WRLCK, (offset), (whence), (len)) == 0)

void err_dump(const char *, ...);
void err_msg(const char *, ...);
void err_quit(const char *, ...);
void err_exit(int, const char *, ...);
void err_ret(const char *, ...);
void err_sys(const char *, ...);

void log_msg(const char *, ...);
void log_open(const char *, int, int);
void log_quit(const char *, ...);
void log_ret(const char *, ...);
void log_sys(const char *, ...);

void TELL_WAIT(void);
void TELL_PARENT(pid_t);
void TELL_CHILD(pid_t);
void WAIT_PARENT(void);
void WAIT_CHILD(void);
#endif

static void err_doit(int, int, const char *, va_list);

void err_ret(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	err_doit(1, errno, fmt, ap);
	va_end(ap);
}

void err_sys(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	err_doit(1, errno, fmt, ap);
	va_end(ap);
	exit(1);
}

void err_exit(int error, const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	err_doit(1, errno, fmt, ap);
	va_end(ap);
	exit(1);
}

void err_dump(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	err_doit(1, errno, fmt, ap);
	va_end(ap);
	abort();
	exit(1);
}

void err_msg(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	err_doit(0, 0, fmt, ap);
	va_end(ap);
}

void err_quit(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	err_doit(0, 0, fmt, ap);
	va_end(ap);
	exit(1);
}

static void err_doit(int errnoflag, int error, const char *fmt, va_list ap)
{
	char buf[MAXLINE];
	vsprintf(buf, fmt, ap);
	if (errnoflag)
		snprintf(buf + strlen(buf), MAXLINE-strlen(buf), ": %s", strerror(errno));
	strcat(buf, "\n");
	fflush(stdout);
	fputs(buf, stderr);
	fflush(NULL);
}

static void log_doit(int, int, const char *, va_list ap);
extern int log_to_stderr;

void log_open(const char *ident, int option, int facility)
{
	if (log_to_stderr == 0)
		openlog(ident, option, facility);
}

void log_ret(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	log_doit(1, LOG_ERR, fmt, ap);
	va_end(ap);
	return;
}

void log_sys(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	log_doit(1, LOG_ERR, fmt, ap);
	va_end(ap);
	exit(2);
}

void log_msg(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	log_doit(0, LOG_ERR, fmt, ap);
	va_end(ap);
}

void log_quit(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	log_doit(0, LOG_ERR, fmt, ap);
	va_end(ap);
	exit(2);
}

static void log_doit(int errnoflag, int priority, const char *fmt, va_list ap)
{
	int errno_save;
	char buf[MAXLINE];

	errno_save = errno;
	vsprintf(buf, fmt, ap);
	if(errnoflag)
		snprintf(buf + strlen(buf), MAXLINE-strlen(buf), ": %s", strerror(errno_save));
	strcat(buf, "\n");
	if (log_to_stderr) {
		fflush(stdout);
		fputs(buf, stderr);
		fflush(stderr);
	} else {
		syslog(priority, buf);
	}
}