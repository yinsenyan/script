import re
import os
import sys
import psutil
import subprocess

def checkIRQ(irq_number):
	irq_cpu = []
	for i in range(len(irq_number)):
		with open('/proc/irq/%d/smp_affinity' %int(irq_number[i]), 'r') as number:
			irq_cpu.append(number.read())
	if len(set(irq_cpu)) == 1:
		print "irq not optimization"
		return False
	else:
		print "irq optimization done"
		return True


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
	cpu_count = psutil.count()
	print cpu_count

def getInterrupts(nic):
	cpu_id = []
	with  open('/proc/interrupts') as interrupts:
		for i in interrupts:
			for j in range(len(nic)):
				if re.search(nic[j], i):
					cpu_id.append(i.split(":")[0])
				else:
					pass
	return cpu_id


if __name__ == '__main__':
	#nic = getNIC()
	#irq_number = getInterrupts(nic)
	#checkIRQ(irq_number)
	getCPU()

