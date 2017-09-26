#/usr/bin/python
#author yanshen(01deyishu@gmail.com)
#date:Tue Sep 26 14:43:48 UTC 2017
#This script prints out all the IP ISP information on the mtr path
import os
import re
import sys
import time
import requests

ipaddr = sys.argv[1]
mtr = 'mtr -n -i 0.3 -c 10 -r -w'+' '+ ipaddr+' '+'> mtr-txt'
os.system(mtr)

def GetIPInfo(ip):
    info = requests.get('http://freeapi.ipip.net/' + ip)
    return info.content.replace('"','').replace(',,,',',')

with open('mtr-txt') as mtr_info:
    info = ''
    pattern = '(?<![\.\d])(?:\d{1,3}\.){3}\d{1,3}(?![\.\d])'
    for i in mtr_info.readlines():
        ip = i.split(' ')[3]
        if re.match(pattern, ip, flags=0):
            info = GetIPInfo(ip)
            time.sleep(1)
        print i.replace('\n',''), info

os.remove('mtr-txt')
