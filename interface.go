package main 

import (
	"fmt"
	"math"
)

//定义一个形状接口 
//定义一个面积计算方法
type Shap interface {
	area() float64
}

type Rectangle struct {
	width float64
	height float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	r float64
}

//圆形实现接口
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}



//实现一个计算面积的方法
func calArea(a ...Shap) float64 {
	total := 0.00
	for _,v := range a {
		total += v.area()
	}
	return total
}




//重写接口方法
//引入的方法一般都是开头首字符大写
func main(){
	// num := -math.Sqrt2
	// fmt.Println(num)
	// product := map[string]Circle {
	// 	"c1" : {r:3},
	// 	"c2" : {r:4},
	// 	"c3" : {r:5},
	// }
	c1 := Circle{r:1}
	c2 := Circle{r:2}
	c3 := Circle{r:3}
	fmt.Println(calArea(&c1, &c2, &c3))

	// for _,v := range product {
	// 	fmt.Println(calArea(v))
	// }
}