package main

import "fmt"

func bubble(l []int)  {
	le := len(l)
	for i,_ := range l {
		for j := 0; j < le - 1 -i; j++ {
			if l[j] > l[j+1] {
				l[j],l[j+1] = l[j+1],l[j]
			}
		}
	}
}

func part(l []int, low int, high int) (p int) {
	p = low
	v := l[p]	// 比较基准值
	pos := low+1	// 低区间存放位置
	for i:=pos; i<=high; i++ {
		if v > l[i] {
			l[pos],l[i] = l[i],l[pos]
			pos++
		}
	}
	// pos目前指向高区间的第一个坑
	// 找准p的位置
	l[pos-1],l[p] = l[p],l[pos-1]
	p = pos-1
	return p
}
func quick(l []int, low int, high int) {
	if low < high {
		p := part(l, low, high)
		quick(l, low, p-1)
		quick(l, p+1, high)
	}
}

func main() {
	l := []int{1,4,0,32432,2,5,2,6,87,23,343,22,309045,343,222}
	fmt.Println(l)
	//bubble(l)
	quick(l, 0, len(l)-1)
	fmt.Println(l)

}
