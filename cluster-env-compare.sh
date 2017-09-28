#/bin/bash
#This script will compare some server configuration and hardware , and the first server will be template

filename="compare-"`date | awk '{print $4}' | tr -s ':' '-'`".txt"

R='\033[31m' # red
G='\033[32m' # green
Y='\033[33m' # yellow
NC='\033[0m' # no color

function info() {
        echo -e "$G[Info] $@ $NC"
}

function warn() {
        echo -e "$Y[Warn] $@ $NC"
}

function error() {
        echo -e "$R[Fatal] $@ $NC"
}

host=($*)

function get_kernel(){
    ker=`ssh root@$1 uname -r`
    echo -n $ker
}

function get_os_version(){
    osv=`ssh root@$1 cat /etc/issue.net`
    echo -n $osv
}

function get_docker_version(){
    dockerv=`ssh root@$1 docker -v | awk '{print $3}' | tr -s ',' ' '`
    echo -n $dockerv
}

function get_cpu_member(){
    c=`ssh root@$1 lscpu  | grep "CPU(s):" | grep -v NUMA | awk '{print $2}'`
    echo -n $c
}

function get_mem(){
    m=`ssh root@$1 free -h | grep Mem | awk '{print $2}'`
    echo -n $m
}

function get_nic(){
    nic=`ssh root@$1 cat  /proc/net/dev | awk '{print $1}' | grep : | grep -vE 'lo|tunl' | tr -s ':' ' '`
    for i in $nic
    do
        if ssh root@$1 ethtool -i $i | grep igb > /dev/null
        then
            echo -n $i:1G ' '
        elif ssh root@$1 ethtool -i $i | grep ixgbe > /dev/null
        then
            echo -n $i:10G ' '
        fi
    done
}

function get_disk(){
    a=`ssh root@$1 fdisk -l | grep "Disk /dev/sd" | awk '{print $3}'`
    sum=`echo $a | awk '{print NF}'`
    echo -n $sum:$a
}

function get_net_info(){
    ip=`ssh root@$1 ip addr | grep -w inet | awk '{print $2}' | grep 100.64 | head -n 1`
    net=`echo $ip | awk -F '.' '{print $1"."$2"."$3".0"}'`
    mask=`echo $ip | awk -F '/' '{print $2}'`
    echo -n $net/$mask
}

echo host"  | "kernel_version"   | "os_version"     |     "docker_version" | "network"      |  "cpu" | "mem" | "disk_count:size"    |    "nic > $filename 
echo ----"  | "--------------"   | "----------"     |     "------" | "-------"      |  "---" | "---" | "----"    |    "--- >> $filename 
for i in $*
do
echo $i " | " $(get_kernel $i) " | " $(get_os_version $i) " | " $(get_docker_version $i) " | " $(get_net_info $i) " | " $(get_cpu_member $i) " | " $(get_mem $i) " | " $(get_disk $i) " | " $(get_nic $i) >> $filename 
done

function compare(){
num=2
while (( num <= 9 ))
do
    IFS=$'\x0A'
    list=(`grep -vE 'host|---' $filename | awk -v num="$num" -F '|' '{print $num}'`)
    IFS=' '
    i=1
    while (( i < ${#list[*]}))
    do
        IFS=$'\x0A'
        if [ ${list[0]} != ${list[$i]} ]
        then
            warn ${host[$i]} `grep host $filename | awk -v num="$num" -F '|' '{print $num}'` mismatch ${host[0]}"(template)"
        fi
        IFS=' '
            let "i++"
    done
    let "num++"
done
}

compare
echo ' '
cat $filename 
