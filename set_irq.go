// 网卡软中断优化
// 0 检查网卡软中断是否优化
// 1 获取CPU核数、网卡队列数、网卡中断编号
// 2 计算模型
// 3 写入文件
// 4 确认中断生效

package main

import (
	//"flag"
	"bufio"
	"fmt"
	"github.com/vishvananda/netlink"
	"io"
	"os"
	"regexp"
	"strings"
)

func getNicList() (niclist []string) {
	nics, _ := netlink.LinkList()
	for _, nic := range nics {
		if nic.Type() == "device" {
			fmt.Println(nic.Attrs().Name)
			niclist = append(niclist, nic.Attrs().Name)
		}
	}
	return
}

func getInterruptsNumber(nic string) {
	interrupts, _ := os.OpenFile("/proc/interrupts", os.O_RDONLY, 0444)
	defer interrupts.Close()
	buf := bufio.NewReader(interrupts)
	for {
		i, _, j := buf.ReadLine()
		if j == io.EOF {
			break
		}
		str := string(i)
		if isOk, _ := regexp.MatchString(nic, str); isOk {
			fmt.Println(strings.Split(str, ":")[0], nic)
		}
	}
}

func main() {
	for _, n := range getNicList() {
		getInterruptsNumber(n)
	}
	fmt.Println("...")
}
