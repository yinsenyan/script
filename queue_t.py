#!/usr/bin/env python
#coding:utf8

import random
import threading
import time

from Queue import Queue

class Producer(threading.Thread):
	def __init__(self,t_name,queue):
		threading.Thread.__init__(self,name=t_name)
		self.data=queue
	def run(self):
		for i in rang(10):
			randomnum = random.ranndint(1,99)
			print "%s:%s is producing %d to the queue!" % (time.ctime(),self.getName(),randomnum)
			self.data.put(randomnum)
			time.sleep(1)
		print "%s:%s finished!" %(time.ctime(),self.getName())

class Consumer_even(threading.Thread):
	def __init__(self,t_name,queue):
		threading.Thread.__init__(self,name=t_name)
		self.data=queue
	def run(self):
		while 1:
			try:
				val_even=self.data.get(1,5)
				if val_even%2 == 0:
					print "%s:%s is consuming. %d in the queue is consumed!" % (time.ctime(),self.getName(,val_even))
					time.sleep(2)
				else:
					self.data.put(val_even)
					time.sleep(2)

			except:
				print "%s:%s finished!" %(time.ctime(),self.getName())
				break

class Consumer_odd(threading.Thread):
	def __init__(self,t_name,queue):
		threading.Thread.__init__(self,name=t_name)
		self.data=queue
	def run(self):
		while 1:
			try:
				val_odd = self.data.get(1,5)
				if val_odd%2!=0:
					print "%s:%s is consuming. %d in the queue is consumed!" %(time.ctime(),self.getName(),val_odd)
					time.sleep(2)
				else:
					self.data.put(val_odd)
					time.sleep(2)
			except:
				print "%s:%s finished!"% (time.ctime(),self.getName())
				break

def main():
	queue=Queue()
	producer = Producer('Pro.',queue)
	consumer_even = Consumer_even('Con_even.',queue)
	consumer_odd = Consumer_odd('Con_odd.',queue)
	producer.start()
	consumer_even.start()
	consumer_odd.start()
	producer.join()
	consumer_even.join()
	consumer_odd.join()
	print 'All threads terminate!'

if __name__ == '__main__':
	main()

		