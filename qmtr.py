#!/user/bin/python
#a tool for mtr , search the ip's isp of every ttl in the traceroute 

import os
import sys
import json
import socket
import urllib2

ipaddress = sys.argv[1]
osname = socket.gethostname()
st1='mtr -n -i 0.3 -c 10 -r -w'+' '+ ipaddress+' '+'> ipaddress'

def GetIP(ip):
    respone = urllib2.urlopen('http://ip.taobao.com/service/getIpInfo.php?ip='+ip)
    d = respone.read()
    data = json.loads(d)
    return data['data']['country'],data['data']['city'],data['data']['isp']

os.system(st1)
fp = open('ipaddress')
print 'HOST: ' + osname + '           Loss%   Snt  Last   Avg   Best  Wrst StDev        ISP'
for i in fp.readlines()[1:]:
    j=i.strip()
    for ip in i.strip().split()[1:2]:
        if ip in i.strip().split()[1:2]:
            if ip == '???':
                ip = 'Unkown IP address'
                print j + '     ' + ip
            elif  ip == '`|--':
                del ip
            elif ip == '|--':
                del ip
            else:
                x,y,z = GetIP()
                if x == 'IANA':
                    z='Internet IP'
                print  unicode(j + '    ' + x + y + z).encode('utf-8')
fp.close()
os.remove('ipaddress')

