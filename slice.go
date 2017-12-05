package main 

import "fmt"
//slice

func  main(){
	list := make([]int, 10)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}

	arr := [5]string{"aa","bb","cc","dd","ee"}
	//根据数组创建切片
	larr := arr[1:4]

	//声明一个切片
	slice := []float64{1.22, 4.22, 5.33}

	var lslice []int

	if lslice == nil {
		fmt.Println("lslice nil")
	}

	fmt.Println(slice)
	fmt.Println(list)
	fmt.Println(list[1:3])
	fmt.Println(larr)

	list = append(list, 12)

	for k,v := range list {
		fmt.Println(k,v)
	}
	
	fmt.Println(cap(list))
	//newlist := copy(list,)
	fmt.Println(list)
	//fmt.Println(newlist)
}