package main 

import (
	"fmt"
	"os"
)

func main(){
	//输出test.go文件的内容
	fh, err := os.Open("data.log")
	fmt.Println(fh)//fh是文件指针
	if( err != nil ){
		//如果报错
		fmt.Println(err)
		return 
	}
	//当程序结束时候，关闭文件资源
	defer fh.Close()
	//统计文件大小
	stat, err := fh.Stat()
	fmt.Println(stat)
	if( err != nil ){
		fmt.Println(err)
		return
	}

	fmt.Println(stat.Size())

	content := make([]byte, stat.Size())

	//此处应该用 = 而不是 :=
	_, err = fh.Read(content)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)
	str := string(content)
	fmt.Println(str)



}