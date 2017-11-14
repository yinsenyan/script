import os
import sys
import psutil

def getNIC():
	nic = psutil.net_if_addrs()
	for i in nic.items():
		print i[0]

def getCPU():
	cpu = psutil.cpu_percent()
	print cpu

if __name__ == '__main__':
	getNIC()
	getCPU()

