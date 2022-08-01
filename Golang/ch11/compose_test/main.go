package main

import "fmt"

type Teacher struct {
	Name  string
	Age   int
	Title string
}

func (t Teacher) teacherInfo() {
	fmt.Printf("姓名:%s, 年龄:%d, 职称:%s", t.Name, t.Age, t.Title)
}

type Course struct {
	Teacher //如果讲师的信息比较多怎么办?将另一个结构体的变量放进来
	Name    string
	Price   int
	Url     string
}

//匿名嵌套 这种做法其实就是我们的语法糖
func (c Course) courseInfo() {
	fmt.Printf("课程名：%s, 价格：%d, 讲师信息：%s %d %s", c.Name, c.Price,
		c.Teacher.Name, c.Age, c.Title)
}

//这种继承的效果有点说服不了人

func main() {
	//go语言的继承 组合
	t := Teacher{
		Name:  "anwma",
		Age:   18,
		Title: "开发",
	}
	c := Course{
		Teacher: t,
		Price:   100,
		Url:     "",
		Name:    "django",
	}
	c.courseInfo()
	//c := Course{
	//	Teacher: Teacher{
	//		Name:  "anwma",
	//		Age:   18,
	//		Title: "rd",
	//	},
	//}

}
