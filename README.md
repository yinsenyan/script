# My python Scrapy/Script

author: yanshen
date:2017年 9月12日 星期二 10时26分06秒 CST

## mtr_isp.py
对mtr的封装，打印出mtr路径上的IP的运营商信息

## cluster-env-compare.sh
比较多台服务器配置(系统版本、内核版本、CPU、磁盘、网卡等)是否相同，以第一个参数为hostname的服务器的配置为准

## translate.py
调用百度的免费API实现一个命令行的翻译工具，用户id和密码可以自己申请，嫌麻烦有需要的邮件：01deyishu@gmail.com
可以这么用：
```
echo 'python /usr/local/bin/translate.py $1' > tl
chmod a+x tl
mv tl `echo $PATH | awk -F ':' '{print $1}'`
```