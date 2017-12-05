package main 

type User struct {
	username string
	password string
}

var user1 = User{"wanwenjie", "123"}
var x = "hello world"
var bb int
var user3 User

var arr [5]int
//用冒号
//此种声明方式不允许定义在函数体外 不要忘记 不要忘记 不要忘记
//user2 := User{username : "hhhh"}

func main(){
	bb = 77
	user2 := User{username : "hhhh"}
	user3 = User{"hello", "world"}
	println(x)
	//arr = [5] int {1, 2, 4, 5, 6}
	arr = [...]int{3,4,5,6,7}

	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	var zhengfang [5][5]int
	format(zhengfang, 5, 5)
	//println(zhengfang[1][1])
	println(user1.username)
	println(user2.username)
	println(user3.password)
}

//数组用的我心好累
func format(arr [5][5]int, v1, v2 int) {
	//数组不能存放一个变量
	x := 0
	y := 0
	for i := 0; i < v1; i++ {
		for j := 0; j < v2; j++ {
			x = i + 1
			y = j + 1
			arr[i][j] = x*y
			print(arr[i][j],"\t")
		}
		print("\r\n")
	}

}