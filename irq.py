import re
import os
import sys
import psutil
import subprocess

def getNIC():
	eth = []
	nic = psutil.net_if_stats()
	for i in nic.items():
		p = subprocess.Popen(["ethtool","-i",i[0]], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
		out, err = p.communicate()
		if err != '':
			pass
		elif re.match("driver: e1000", out):
			eth.append(i[0])
	return eth

def getCPU():
	cpu = psutil.cpu_percent()
	print cpu

def getInterrupts(nic):
	#nic = []
	cpu_id = []
	interrupts = open('/proc/interrupts')
	#print interrupts
	for i in range(len(nic)):
		for j in interrupts:
			print j
			if re.match(nic[i], j):
				cpu_id.append(j.split(":")[0])
			else:
				pass
	interrupts.close()
	return cpu_id


if __name__ == '__main__':
	nic = []
	nic = getNIC()
	#print nic
	print getInterrupts(nic)

