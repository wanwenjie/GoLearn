package main 

//指针的学习

func main(){
	//var ptr *int
	//a := 10
	//ptr = &a
	//println(ptr)

	//数组的指针 
	arr := [5]int {4,5,6,7,8}
	atr := &arr[0]
	println(atr)
	println(*atr)
	//指针变量加上* 表示访问该内存地址的存储值

	//指针数组
	var mptr [5]*int
	for i := 0; i < len(arr); i++ {
		mptr[i] = &arr[i]
		print(mptr[i],"\t")
	}

	a := 5
	b := 10
	swap(&a, &b)
	print(a,b)

}

//传入的是地址，交换的是变量的值
func swap(x *int, y *int) {
	tmp := 0
	tmp = *x
	*x = *y
	*y = tmp
}

//传入的是值，并不会真正改变变量
func swp(x int, y int) {
	tmp := 0
	tmp = x
	x = y
	y = tmp
}