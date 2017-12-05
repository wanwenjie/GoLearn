package main 

import (
	"fmt"
	"strings"
)

func main(){
	str := "hello"
	fmt.Println(strings.Contains(str, "ll"))

	fmt.Println(strings.Split("1#sss#333", "#"))

	fmt.Println(strings.Join([]string{"222","aaaa","4444"}, "#"))

	fmt.Println(strings.ToUpper(str))
	
	fmt.Println(strings.ToLower(str))

	fmt.Println(strings.Count(str, "l"))

	fmt.Println(strings.Index(str, "o"))

	fmt.Println(strings.HasSuffix(str, "ll"))//是否有后缀

	fmt.Println(strings.HasPrefix(str, "el"))//是否有前缀

	fmt.Println(strings.Replace(str, "l", "0", 2))
	
}