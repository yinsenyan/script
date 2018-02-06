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
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type nic struct {
	name         string
	irqIndex     string
	irqs         []string
	oldCPUNumber []string
	newcPUNumber []int
}

func getNicList() (niclist []string) {
	nics, _ := netlink.LinkList()
	for _, nic := range nics {
		if nic.Type() == "device" {
			if nic.Attrs().Name != "lo" && nic.Attrs().OperState.String() == "up" {
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

func getCPUNumber() (cpuNumber []string) {
	smp, err := os.OpenFile("/proc/irq/19/smp_affinity", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("can't open file /proc/irq/xx/smp_affinity")
	}
	defer smp.Close()
	data, err := ioutil.ReadAll(smp)
	if err != nil {
		fmt.Println("can't read file /proc/irq/xx/smp_affinity")
	}
	cpuNumber = append(cpuNumber, strings.Replace(string(data), "\n", "", -1))
	return
}

func getIrq(name string) (irqIndex string, irqs []string) {
	reg := regexp.MustCompile(name + "?.+")
	var irq []string
	irq = make([]string, 10)
	interrupts, _ := os.OpenFile("/proc/interrupts", os.O_RDONLY, 0444)
	defer interrupts.Close()
	buf := bufio.NewReader(interrupts)
	for {
		i, _, j := buf.ReadLine()
		if j == io.EOF {
			break
		}
		str := string(i)
		if isok, _ := regexp.MatchString(name, str); isok {
			queue := reg.FindString(str)
			if queue == name {
				irq[0] = strings.Split(str, ":")[0]
			} else {
				irq = append(irq, strings.Split(str, ":")[0])
			}
		}
	}
	return irq[0], irq[1:]
}

func main() {
	//var nics []nic
	for _, i := range getNicList() {
		var n nic
		n.name = i
		n.irqIndex, n.irqs = getIrq(i)
		n.oldCPUNumber = getCPUNumber()
		fmt.Println(n)
	}
	fmt.Println(getCPUCount())
}
