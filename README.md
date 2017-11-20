# My python/shell Script

author: yanshen
date:2017年 9月12日 星期二 10时26分06秒 CST

## mtr_isp.py
对mtr的封装，打印出mtr路径上的IP的运营商信息

## cluster-env-compare.sh
比较多台服务器配置(系统版本、内核版本、CPU、磁盘、网卡等)是否相同，以第一个参数为hostname的服务器的配置为准
注意：测试了下，当前仅仅适用于Ubuntu 16.04

## translate.py
调用百度的免费API实现一个命令行的翻译工具，用户id和密码可以自己申请，嫌麻烦有需要的邮件：01deyishu@gmail.com
可以这么用：
```
echo 'python /usr/local/bin/translate.py $1' > tl
chmod a+x tl
mv tl `echo $PATH | awk -F ':' '{print $1}'`
```

## mac_trans.py
转换mac地址格式(解决服务器和交换机mac地址格式不同，每次手动改很烦的问题)

## irq.py
网卡绑定软中断到不同CPU

## seek_switch.py
通过ping的ttl判断远端IP是否为交换机，目前仅支持CIDR大于等于24位的扫描

## ospf_status.py
监控交换机ospf状态的脚本

## pachong_v1.py
练手的爬虫，v0.1版本

## apue.h error.c
学习unix高级环境编程的头文件,在学习的例子代码前面添加
```
#include "error.c"
```
来使用作者定义的一些函数和统一导入头文件。
PS:error.c中已经include了apue.h
