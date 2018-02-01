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

type nic struct {
	name         string
	irqIndex     string
	irqs         []string
	oldCPUNumber []int
	newcPUNumber []int
}

func getNicList() (niclist []string) {
	nics, _ := netlink.LinkList()
	for _, nic := range nics {
		if nic.Type() == "device" {
			if nic.Attrs().Name != "lo" {
				niclist = append(niclist, nic.Attrs().Name)
			}
		}
	}
	return
}

func getCPUCount() (cpuCount int) {
	cpuCount = 0
	stat, err := os.OpenFile("/proc/stat", os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("Error open file /proc/stat")
	}
	defer stat.Close()
	buf := bufio.NewReader(stat)
	for {
		i, _, j := buf.ReadLine()
		if j == io.EOF {
			break
		}
		if isOk, _ := regexp.MatchString("cpu[0-9]+", string(i)); isOk {
			cpuCount++
		}
	}
	return
}

func getCPUNumber() {
	smp, err := os.OpenFile("/proc/irq/19/smp_affinity", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("can't open file /proc/irq/xx/smp_affinity")
	}
	defer smp.Close()
	fmt.Println(smp)
}

func getIrq(name string) (irqIndex string, irqs []string) {
	var irq []string
	interrupts, _ := os.OpenFile("/proc/interrupts", os.O_RDONLY, 0444)
	defer interrupts.Close()
	buf := bufio.NewReader(interrupts)
	for {
		i, _, j := buf.ReadLine()
		if j == io.EOF {
			break
		}
		str := string(i)
		if isOk, _ := regexp.MatchString(name, str); isOk {
			irq = append(irq, strings.Split(str, ":")[0])
		}
	}
	return irq[0], irq[1:]
}

func main() {
	for _, i := range getNicList() {
		var n nic
		n.name = i
		n.irqIndex, n.irqs = getIrq(i)
		fmt.Println(n)
	}
	fmt.Println(getCPUCount())
	getCPUNumber()
}
