#!/usr/bin/python
import re
import sys
import pyping

def calcIPlist(minhost,maxhost):
	ipaddress = []
	for ip in range(minhost[3], maxhost[3]+1):
		ipaddress.append(str(minhost[0]) + '.' + str(minhost[1]) + '.' + str(minhost[2]) + '.' + str(ip))
	return ipaddress

def BintoInt(list2):
	ip = [0,0,0,0]
	for i in range(0,4):
		for n in range(8*i, 8*i + 8):
			if list2[n] == '1':
				ip[i] += 2**( 7 - n%8)
	return ip

def calcIP(net):
	ip = net.split('/')[0]
	mask = net.split('/')[1]
	ip2 = []
	mask2 = []
	netnumber2 = []
	unmask2 = []
	broadcast2 = []
	minhost2 = []
	maxhost2 = []
	for i in ip.split('.'):
		need = '0' * (10 - len(bin(int(i))))
		for j in range(0, len(str(bin(int(i))).replace('0b',need))):
			ip2.append(str(bin(int(i))).replace('0b', need)[j:j+1])
	for k in range(0, 32):
		mask2.append(str(bin(2**32 - 2**(32 - int(mask)))).replace('0b', '')[k:k+1])
	for n in range(0, 32):
		if mask2[n] == '0' :
			netnumber2.append('0')
			unmask2.append('1')
			broadcast2.append('1')
		else:
			netnumber2.append(ip2[n])
			unmask2.append('0')
			broadcast2.append(ip2[n])
		minhost2.append(netnumber2[n])
		maxhost2.append(broadcast2[n])
	minhost2[31] = '1'
	maxhost2[31] = '0'
	# print BintoInt(ip2)
	# print BintoInt(mask2)
	# print BintoInt(unmask2)
	# print BintoInt(netnumber2)
	# print BintoInt(minhost2)
	# print BintoInt(maxhost2)
	# print BintoInt(broadcast2)
	return calcIPlist(BintoInt(minhost2), BintoInt(maxhost2))

def GetIP():
	net = sys.argv[1]
	if re.findall('/', net):
		return calcIP(net)
	else:
		return net

def GetIPTTL(ip):
	p = pyping.ping(str(ip) ,count=1)
	info = p.output[1]
	print info
	if re.search(r"timed out", info, flags=0):
		pass
	else:
		return ip, re.findall(r"ttl=(\d*)", info, re.I)[0]

if __name__ == '__main__':
	ip_list = GetIP()
	ttl_list = []
	linux_list = []
	windows_list = []
	switch_list = []
	try:
		p = pyping.ping('114.114.114.114' ,count=1)
	except Exception as e:
		print e
	else:
		for i in ip_list:
			if GetIPTTL(str(i)) == None:
				pass
			else:
				ttl_list.append(GetIPTTL(str(i)))
	for i in range(len(ttl_list)):
		if 1 <= int(ttl_list[i][1]) <= 64:
			linux_list.append(ttl_list[i][0])
		elif 65 <= int(ttl_list[i][1]) <= 128:
			windows_list.append(ttl_list[i][0])
		elif 129 <= int(ttl_list[i][1]) <= 256:
			switch_list.append(ttl_list[i][0])
		else:
			pass
	print ip_list,ttl_list
	print switch_list,windows_list,linux_list
