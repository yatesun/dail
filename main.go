package main

import "fmt"

func main() {
	dial := newDial()
	dial.convertDigits(2, 3)
	// [ad ae af bd be bf cd ce cf]
	dial.convertDigits(7, 9)
	// [pw px py pz qw qx qy qz rw rx ry rz sw sx sy sz]
	dial.convertDigits(2, 0, 3)
	// [ad ae af bd be bf cd ce cf]
	dial.convertDigits(0, 3)
	// [d e f]
	dial.convertDigits(2, 3, 4)
	// [adg adh adi aeg aeh aei afg afh afi bdg bdh bdi beg beh bei bfg bfh bfi cdg cdh cdi ceg ceh cei cfg cfh cfi]
}

// Stage 1
type Dial struct {
	m map[int]string
}

func newDial() *Dial {
	m := make(map[int]string)
	m = map[int]string{2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"}
	return &Dial{m: m}
}

func (d Dial) convertDigits(arr ...int) (list []string) {
	// empty arr
	if len(arr) == 0 {
		return
	}
	// how many string in return
	var num int
	for _, n := range arr {
		i := len(d.m[n])
		if i == 0 {
			continue
		}
		if num == 0 {
			num = i
		} else {
			num = num * i
		}
	}

	// empty string list
	if num == 0 {
		return
	}

	// list
	list = make([]string, num)

	// [2, 3, 4]
	for _, n := range arr {
		if len(d.m[n]) == 0 {
			continue
		}
		num = num / len(d.m[n])
		for i := range list {
			for j, s := range d.m[n] {
				if j == i/num%len(d.m[n]) {
					list[i] += string(s)
				}
			}
		}
	}
	fmt.Println(list)
	return list
}
