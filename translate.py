#/usr/bin/env python
#call baidu api to translate word
#coding=utf8

import httplib
import md5
import urllib
import random
import argparse
import json
import re

def analysis_lang():
	pattern_e = '[a-z]+'
	parser = argparse.ArgumentParser()
	parser.add_argument('word', type=str, help='Need a word to translate')
	word = parser.parse_args().word
	if re.match(pattern_e, word, re.I):
		fromLang = 'en'
		toLang = 'zh'
	else:
		fromLang = 'zh'
		toLang = 'en'
	return word, fromLang, toLang

def translate(option):
	appid = 'id'
	secretKey = 'password'
	word = option[0]
	fromLang = option[1]
	toLang = option[2]
	httpClient = None
	myurl = '/api/trans/vip/translate'
	salt = random.randint(32768, 65536)

	sign = appid+word+str(salt)+secretKey
	m1 = md5.new()
	m1.update(sign)
	sign = m1.hexdigest()
	myurl = myurl+'?appid='+appid+'&q='+urllib.quote(word)+'&from='+fromLang+'&to='+toLang+'&salt='+str(salt)+'&sign='+sign

	try:
	    httpClient = httplib.HTTPConnection('api.fanyi.baidu.com')
	    httpClient.request('GET', myurl)

	    response = httpClient.getresponse()
	    js = json.load(response)
	    return js['trans_result'][0]['dst'], word
	except Exception, e:
	    print e
	finally:
	    if httpClient:
	        httpClient.close()

if __name__ == '__main__':
	result = translate(analysis_lang())
	print result[1], '--->', result[0]
