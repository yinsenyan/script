#!/bin/bash
#to deploy a test server for upload/download/smokeping services.
smokeping_master=127.0.0.1
if grep nameserver /etc/resolv.conf
then
    echo nameserver is OK
else
    echo 'nameserver 114.114.114.114' >> /etc/resolv.conf
fi
wget http://$smokeping_master:8080/idc.tar
tar xvf idc.tar
./recv_v_port 8000 &
apt-get update
apt-get -y install nginx
sed -i 's/listen 80/listen 8888/g' /etc/nginx/sites-enabled/default
apt-get -y install smokeping libio-socket-inet6-perl sendmail
service nginx restart

function init_12.04
{
    echo This is OK
}
function init_14.04
{
    apt-get -y install apache2
    cp ./000-default.conf /etc/apache2/sites-enabled/
    a2enmod cgi
}

if grep 14.04 /etc/issue.net
then
    init_14.04
    nginx_path=/usr/share/nginx/html/
else
    init_12.04
    nginx_path=/usr/share/nginx/www/
fi
    cp ./Targets /etc/smokeping/config.d/
    cp ./100k.jpg $nginx_path
    service smokeping restart
    service apache2 restart
    echo upload: `ip addr | grep -w inet | awk '{print $2}' | awk -F '/' '{print $1}' | grep -v 127` 8000
    echo download: http://`ip addr | grep -w inet | awk '{print $2}' | awk -F '/' '{print $1}' | grep -v 127`:8888/100k.jpg
    echo -e '\t' md5sum: `md5sum $nginx_path/100k.jpg`
    echo smokeping: http://`ip addr| grep -w inet| awk '{print $2}' | awk -F '/' '{print $1}' | grep -v 127`/cgi-bin/smokeping.cgi
