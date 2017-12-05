package main 

import (
"fmt"
)
//关于map
func main(){
	//声明一个map key=>value 都是string的
	var userMap map[string]string
	userMap = make(map[string]string)
	userMap["username"] = "wanwenjie"
	userMap["password"] = "123123"
	fmt.Println(userMap)

	for k,v := range userMap {
		fmt.Println(k,v)
	}

	//删除map元素 key
	delete(userMap, "username")

	fmt.Println(userMap)
	fmt.Println(jiecheng(5))

	ll := 8
	fmt.Println(string(ll))
	fmt.Println(string(ll)," hello world")

}


func jiecheng(n int) int {
	if n > 0 {
		return n * jiecheng(n-1)
	}
	return 1
}