package main

import (
	"fmt"
)

type Pcb struct {
	Pid    int
	Ppid   int
	Prio   int
	status int
}
type Pnode struct {
	Node    *Pcb
	Sub     *Pnode
	Brother *Pnode
	Next    *Pnode
}

var Proot *Pnode
var Plink *Pnode

func Createpc(para [3]int) int {
	var p, p1, pp *Pnode
	pflag := 0
	for p = Plink; p != nil; p = p.Next {
		if p.Node.Pid == para[0] {
			fmt.Printf("Pid %d 已经存在", para[0])
			return -1
		}
		if p.Node.Pid == para[1] {
			pflag = 1
			pp = p
		}
	}
	if pflag != 1 {
		fmt.Printf("父节点Ppid %d 不存在!", para[1])
		return -2
	}
	p1 = new(Pnode)
	p1.Node = new(Pcb)
	p1.Node.Pid = para[0]
	p1.Node.Ppid = para[1]
	p1.Node.Prio = para[2]
	p1.Sub = nil
	p1.Brother = nil
	p1.Next = nil
	if pp.Sub == nil {
		pp.Sub = p1
	} else {
		for p = pp.Sub; p.Brother != nil; p = p.Brother {
		}
		p.Brother = p1
	}
	for p = Plink; p.Next != nil; p = p.Next {
	}
	p.Next = p1
	return 0
}

func Showdetail() {
	var p, p1 *Pnode
	p = Plink
	for p != nil {
		fmt.Printf("%d (Prio %d):", p.Node.Pid, p.Node.Prio)
		p1 = p.Sub
		for p1 != nil {
			fmt.Printf(" %d (Prio %d)", p1.Node.Pid, p1.Node.Prio)
			p1 = p1.Brother
		}
		fmt.Printf("\n")
		p = p.Next
	}
	fmt.Printf("\n")
}

func main() {
	var cflag, pflag int
	var chose int
	Proot := new(Pnode)
	Proot.Node = new(Pcb)
	Proot.Node.Pid = 0
	Proot.Node.Ppid = 1
	Proot.Node.Prio = 0
	Proot.Next = nil
	Proot.Sub = nil
	Proot.Brother = nil
	Plink = Proot
	fmt.Println("创建进程：1    显示所有进程：2     退出：0")

LABEL:
	for {
		cflag = 0
		pflag = 0
		fmt.Printf("请输入命令:")
		fmt.Scanf("%d", &chose)
		switch chose {
		case 0:
			break LABEL
		case 2:
			Showdetail()
			cflag = 1
			pflag = 1
		case 1:
			cflag = 1
			var temp [3]int
			for i := 0; i < 3; i++ {
				fmt.Scanf("%d", &temp[i])
			}
			Createpc(temp)
			pflag = 1
		}
		if cflag == 0 {
			fmt.Println("输入错误命令！")
		} else {
			if pflag == 0 {
				fmt.Println("参数有误！")
			}
		}
	}
}
