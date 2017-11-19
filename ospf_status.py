import os,sys,time

pid = os.fork()
if pid>0:
	sys.exit(0)

os.chdir("/tmp")
os.setsid()
os.umask(0)

pid = os.fork()
if pid>0:
	sys.exit(0)

sys.stdout.flush()
sys.stderr.flush()
si = file("/dev/null",'r')
so = file("/dev/null",'a+')
se = file("/dev/null",'a+',0)
os.dup2(si.fileno(),sys.stdin.fileno())
os.dup2(so.fileno(),sys.stdout.fileno())
os.dup2(se.fileno(),sys.stderr.fileno())

while True:
	time.sleep(10)
	f = open('/tmp/test.txt','a')
	f.write('hello')