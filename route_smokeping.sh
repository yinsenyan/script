#!/bin/bash
#to look up the ip's isp and route add a best gateway
#auther: yanshne
#date:2015-11-26
smokeping_master=127.0.0.1
wget http://$smokeping_master:8080/Targets
wget http://$smokeping_master:8080/host_gateway
ip_list=`grep host Targets  | grep CUCC | awk '{print $3}'`
for i in $ip_list
do
    gateway=`grep CUCC host_gateway | grep $HOSTNAME | awk '{print $3}'`
    network_dev=`grep CUCC host_gateway | grep $HOSTNAME | awk '{print $4}'`
    route add -host $i netmask 0.0.0.0 gw $gateway dev $network_dev
done
rm Targets
rm host_gateway
