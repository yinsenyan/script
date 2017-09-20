#-*- c
import requests
import re
#http://www.juemei.com/renwu/201609/5122_4.html
#<div class="wrap"><img src="http://img.juemei.com/album/2016-09-20/57e092c9a9007.jpg">
#<div class="add_album_tips">next page</div>
#</div>

def GetHtml(url):
	pattern = '<div class="wrap"><img src="(.*?)">'
	html = requests.get(url)
	if html.text == '404':
	 	img_url = 'Url not found'
	else:
		img_url = re.search(pattern, html.content, re.S).group(1)
	return img_url

if __name__ == '__main__':
	for i in range(1,10):
		url = 'http://www.juemei.com/renwu/201609/5122_%d.html' % (i)
		print GetHtml(url)
	# url = 'http://www.juemei.com/renwu/201609/5122_2.html'
	# print GetHtml(url)