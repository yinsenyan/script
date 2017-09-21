#-*- c
import requests
import re


def GetHtml(url):
#<a href="networking-concepts-HOWTO-1.html">Introduction</a>
	pattern = '<A HREF=".*?">(.*?)</A>'
	html = requests.get(url)
	print html.status_code
	if html.text == '404':
	 	img_url = 'Url not found'
	else:
		img_url = re.findall(pattern, html.content, re.S)#.groups()
	return img_url

if __name__ == '__main__':
	for i in range(1):
		url = 'http://www.netfilter.org/documentation/HOWTO//networking-concepts-HOWTO.html'
		l = GetHtml(url)
		for i in  range(1,len(l)):
			print i, l[i]
