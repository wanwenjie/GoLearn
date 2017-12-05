package main 

import "fmt"

type Author struct {
	name string
	age int
	desc string
}

func (a *Author) call() {
	fmt.Println("hhhhh")
}

//嵌入式编程
type Book struct {
	name string
	price float64
	Author
}



// func main(){
// 	book1 := Book{"明朝的那些事儿", 10.32, "不知道谁是作者"}
// 	book := &book1
// 	printr(book)
// }

// func printr(book *Book) {
// 	println(book.name)
// 	println(book.price)
// 	println(book.author)
// }

func main(){
	a := Author{name:"wanwenjie", age:21, desc:"handsame"}
	a.call()
	b := Book{name:"c++入门到放弃", price:23.33, Author:Author{name:"金庸", age:100, desc:"著名，武侠小说家"}}
	fmt.Println(b)
	param := []int{1,3,5,7,9}
	//sum 应该传int的类型
	fmt.Println(sum(param ...))
	fmt.Println(sum(1,3,5,7,9))

}

func sum(argc ...int) int{
	sum := 0
	for _, v := range argc {
		sum += v
	}
	return sum
}