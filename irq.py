import os
import sys
import psutil
import subprocess

def getNIC():
	nic = psutil.net_if_stats()
	for i in nic.items():
		subprocess.Popen(["ethtool",i[0]])

def getCPU():
	cpu = psutil.cpu_percent()
	print cpu

def getInterrupts():
	with open('/proc/interrupts') as interrupts:
		for i in interrupts.readlines():
			print i

if __name__ == '__main__':
	getNIC()
	getCPU()
	getInterrupts()

