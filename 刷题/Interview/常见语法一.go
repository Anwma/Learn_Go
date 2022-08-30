package main

import (
	"fmt"
	"runtime"
)

//------------------------------------------------------------------

//1、下⾯代码能运⾏吗？为什么。
/*
type Param map[string]interface{}
type Show struct {
	Param
}

// error1 : main后不能加数字
// error2 : new关键字无法初始化 Show 结构体中的 Param 属性，所以直接对 s.Param 操作会出错
func main1() {
	s := new(Show)
	s.Param["RMB"] = 10000
}

*/

//------------------------------------------------------------------

//2.请说出下⾯代码存在什么问题。
/*
type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student:
		msg.Name = "hello"
	}

	//case *student,student:
	//	msg.Name = "hello"
	//}
}

*/

/*
golang中有规定， switch type 的 case T1 ，类型列表只有⼀个，那么 v := m.(type)
中的 v 的类型就是T1类型。
如果是 case T1, T2 ，类型列表中有多个，那 v 的类型还是对应接⼝的类型，也就
是 m 的类型。
所以这⾥ msg 的类型还是 interface{} ，所以他没有 Name 这个字段，编译阶段就会
报错。具体解释⻅： https://golang.org/ref/spec#Type_switches
*/

//------------------------------------------------------------------

//3、写出打印的结果。

/*
type People struct {
	//私有属性 无法被导出
	name string `json:"name"`
}

func main() {
	js := `{
 		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
*/

//Output: people:  {}

//------------------------------------------------------------------

//4、下⾯的代码是有问题的，请说明原因。
/*
type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}
func main() {
	p := &People{}
	p.String()
}

*/

//在golang中 String() string ⽅法实际上是实现了 String 的接⼝的，该接⼝定义在
//fmt/print.go 中：

//type Stringer interface {
//	String() string
//}

//在使⽤  fmt 包中的打印⽅法时，如果类型实现了这个接⼝，会直接调⽤。⽽题⽬中打
//印  p 的时候会直接调⽤  p 实现的  String() ⽅法，然后就产⽣了循环调⽤。

//------------------------------------------------------------------

// 5、请找出下⾯代码的问题所在。
/*
func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

*/
/*
在 golang 中  goroutine 的调度时间是不确定的，在题⽬中，第⼀个
写  channel 的  goroutine 可能还未调⽤，或已调⽤但没有写完时直接  close 管道，
可能导致写失败，既然出现  panic 错误。
*/

//------------------------------------------------------------------
//6、请说明下⾯代码书写是否正确。
/*
var value int32

func main() {
	SetValue(12)
}
func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
}

*/

//atomic.CompareAndSwapInt32 函数不需要循环调⽤。

// ------------------------------------------------------------------

// 7、下⾯的程序运⾏后为什么会爆异常。
/*
type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}
func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}
func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError() //1
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}
func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 100000000000000) //2
}
func main() {
	p := new(Project)
	p.Main()
}

*/

//有⼀下⼏个问题：
//1. time.Sleep 的参数数值太⼤，超过了  1<<63 - 1 的限制。
//2. defer p.deferError() 需要在协程开始出调⽤，否则⽆法捕获  panic 。

// ------------------------------------------------------------------

// 8、请说出下⾯代码哪⾥写错了
/*
func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for a := range abc {
			fmt.Println("a: ", a)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}

*/

//协程可能还未启动，管道就关闭了。

// ------------------------------------------------------------------

// 9、请说出下⾯代码，执⾏时为什么会报错
/*
type Student struct {
	name string
}

func main() {
	//m := map[string]Student{"people": {"zhoujielun"}}
	m := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
	fmt.Printf("%v", m)
}

*/

//map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在
//map改变时会变得⽆效。故如果需要修改map值，可以将 map 中的⾮指针类型
//value ，修改为指针类型，⽐如使⽤ map[string]*Student .

// ------------------------------------------------------------------

// 10、请说出下⾯的代码存在什么问题？
/*
type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}
func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}

*/

//依据4个goroutine的启动后执⾏效率，很可能打印111func4，但其他的111func*也
//可能先执⾏，exec只会返回⼀条信息

// ------------------------------------------------------------------

//11、下⾯这段代码为什么会卡死？

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}
/*

Golang 中，byte 其实被 alias 到 uint8 上了。所以上⾯的 for 循环会始终成⽴，因为
i++ 到 i=255 的时候会溢出，i <= 255 ⼀定成⽴。
也即是， for 循环永远⽆法退出，所以上⾯的代码其实可以等价于这样：

go func () {
	for {}
}

正在被执⾏的 goroutine 发⽣以下情况时让出当前 goroutine 的执⾏权，并调度后⾯的
goroutine 执⾏：
1.  IO 操作
2.  Channel 阻塞
3.  system call
4.  运⾏较⻓时间
如果⼀个 goroutine 执⾏时间太⻓，scheduler 会在其 G 对象上打上⼀个标志（
preempt），当这个 goroutine 内部发⽣函数调⽤的时候，会先主动检查这个标志，如
果为 true 则会让出执⾏权。
main 函数⾥启动的 goroutine 其实是⼀个没有 IO 阻塞、没有 Channel 阻塞、没有
system call、没有函数调⽤的死循环。
也就是，它⽆法主动让出⾃⼰的执⾏权，即使已经执⾏很⻓时间，scheduler 已经标志
了 preempt。
⽽ golang 的 GC 动作是需要所有正在运⾏  goroutine 都停⽌后进⾏的。因此，程序
会卡在  runtime.GC() 等待所有协程退出。
 */
// ------------------------------------------------------------------
