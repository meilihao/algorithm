package demo

import (
	"algo/helper"
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseZpoolStatusConfig(t *testing.T) {
	// zpool status输出的config内容, 已按换行切分, 并去除行首的'\t'
	raw := `NAME        STATE     READ WRITE CKSUM
tank        ONLINE       0     0     0
  mirror-0  ONLINE       0     0     0
    sde     ONLINE       0     0     0
    sdf     ONLINE       0     0     0
  mirror-1  ONLINE       0     0     0
    sdg     ONLINE       0     0     0
    sdh     ONLINE       0     0     0
  mirror-2  ONLINE       0     0     0
    sdi     ONLINE       0     0     0
    sdj     ONLINE       0     0     0
cache
  c0t0d0    ONLINE       0     0     0
logs
  mirror    ONLINE       0     0     0
    c0t2d0  ONLINE       0     0     0
    c0t4d0  ONLINE       0     0     0`

	var lines []string
	sc := bufio.NewScanner(strings.NewReader(raw))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	spew.Dump(lines)

	devs := ParseZpoolStatusConfig(lines)
	PrintDevices(0, devs)
}

func PrintDevices(indent int, devices []*Vdev) {
	for _, dev := range devices {
		fmt.Printf("%s%+v\n", strings.Repeat(" ", indent), dev.Name)
		if len(dev.Children) > 0 {
			PrintDevices(indent+2, dev.Children)
		}
	}
}

type Vdev struct {
	Name     string
	Children []*Vdev
}

/*
解析zpool vdev, 思路:
1. zpool status输出的config的一个字符是"\t", 之后按两个空格来缩进表示层次
*/
func ParseZpoolStatusConfig(lines []string) []*Vdev {
	devs := []*Vdev{}
	stack := helper.NewStack[*Vdev]()
	lastLevel := 0 // 上一行的缩进
	curLevel := 0  //  当前行的缩进
	for _, line := range lines {
		curLevel = GetIndentSpace(line)
		fields := strings.Fields(line)
		if len(fields) < 1 {
			fmt.Println("invalid line:", line)
			continue
		}
		if fields[0] == "NAME" {
			continue
		}

		dev := &Vdev{
			Name: fields[0],
		}

		if curLevel == 0 { // root vdev
			devs = append(devs, dev)
			stack.Push(dev)
			lastLevel = 0
			continue
		}

		if curLevel > lastLevel { // child vdev
			p, ok := stack.Peek()
			if !ok {
				panic(line)
			}
			p.Children = append(p.Children, dev)
			stack.Push(dev)
			lastLevel = curLevel
		} else if curLevel == lastLevel { // 兄弟dev
			stack.Pop()
			p, ok := stack.Peek()
			if !ok {
				panic(line)
			}
			p.Children = append(p.Children, dev)
			stack.Push(dev)
		} else if curLevel < lastLevel {
			for curLevel < lastLevel { // 找到同级vdev
				stack.Pop()
				lastLevel -= 1
			}

			// 找到同级vdev的parent
			stack.Pop()
			p, ok := stack.Peek()
			if !ok {
				panic(line)
			}
			p.Children = append(p.Children, dev)
			stack.Push(dev)
		}
	}
	return devs
}

func GetIndentSpace(line string) int {
	var n int
	for _, v := range line {
		if v == ' ' {
			n++
		} else {
			return n / 2
		}
	}
	return n / 2
}
