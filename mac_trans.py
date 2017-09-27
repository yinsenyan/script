import sys
import re

mac_address = sys.argv[1]
mac = []
macaddress = ''

if re.search(r':',mac_address):
    for i in mac_address.split(":"):
        mac.append(i)
    for j in mac:
        macaddress += j

else:
    for i in mac_address.split("-"):
        mac.append(i)
    for j in mac:
        macaddress += j

print "Switch Mac Address:",macaddress[0:4]+'-'+macaddress[4:8]+'-'+macaddress[8:12]
print "Server Mac Address:",macaddress[0:2]+':'+macaddress[2:4]+':'+macaddress[4:6]+':'+macaddress[6:8]+':'+macaddress[8:10]+':'+macaddress[10:12]

