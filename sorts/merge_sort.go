package main

import "fmt"

/**
题目：归并排序算法
使用归并排序算法排序数组a
 */
func merge_sort(a []int,n int)  {
	merge_sort_c(a,0,n-1)
}
//递归函数,递归排序
func merge_sort_c( a []int, p,r int) {
	//递归结束条件
	if p >= r {
		return
	}
	//获取中间变量，分而治之
	q := (p+r)/2
	//递归左侧和右侧子数值
	merge_sort_c(a,p,q)
	merge_sort_c(a,q+1,r)
	//合并已排好序的左右两侧子数组
	merge(a ,p,q,r)
}
//合并数组
func merge(a []int,p,q,r int) {
	i := p
	j:=q+1
	k:=0
	var tmp = make([]int,r+1)
	for ;i<=q && j<=r; {
		if a[i]<=a[j]{
			tmp[k] = a[i]
			k++
			i++
		}else{
			tmp[k] = a[j]
			k++
			j++
		}
	}
	//判断哪个自切片有剩余数据
	start := i
	end := q
	if j<=r {
		start=j
		end = r
	}
	for ;start<=end;{
		tmp[k]=a[start]
		k++
		start++
	}
	// 将tmp中的数组拷贝回A[p...r]  for i:=0 to r-p do {    A[p+i] = tmp[i]  }
	for i=0;i<= r-p;i++{
		a[p+i] = tmp[i]
	}
}

func main() {
	var a = []int{11,23,1,2,3,5,6,4,9,10}
	fmt.Println(a)
	merge_sort(a, len(a))
	fmt.Println(a)
}
