package main

//修改过了
import (
	"fmt"
	"math"
)

//1.创建一个表示错误的结构类型
type error interface {
	Error() string
}
type areaError struct {
	err    string
	radius float64
}

//2.实现error接口
//使用接收者*areaError，实现error接口的Error()string方法
func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f : %s\n", e.radius, e.err)
}

//返回的是float64, 和error接口
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}
func main() {
	/*
			f, err := os.Open("/test.txt")
			if err != nil { //有错误
				fmt.Println(err)
				fmt.Printf("%T", err)
				return
			}


		fmt.Println(f.Name(), "opened successfully")

	*/
	radius := -20.0
	area, err := circleArea(radius)
	fmt.Printf("%T\n", err)
	fmt.Println(err)
	fmt.Println(err.Error())
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)

}
