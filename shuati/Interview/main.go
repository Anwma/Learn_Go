package main

//
//import "unsafe"
//
//func main() {
//
//}
//
//type StringHeader struct {
//	Data uintptr
//	Len  int
//}
//
//type SliceHeader struct {
//	Data uintptr
//	Len  int
//	Cap  int
//}
//
//// Stack describes a Go execution stack.
//// The bounds of the stack are exactly [lo, hi),
//// with no implicit data structures on either side.
//// 用于记录goroutine使用的栈的起始和结束位置
//type stack struct {
//	lo uintptr // 栈顶，指向内存低地址
//	hi uintptr // 栈底，指向内存高地址
//}
//
//type gobuf struct {
//	// The offsets of sp, pc, and g are known to (hard-coded in) libmach.
//	//
//	// ctxt is unusual with respect to GC: it may be a
//	// heap-allocated funcval, so GC needs to track it, but it
//	// needs to be set and cleared from assembly, where it's
//	// difficult to have write barriers. However, ctxt is really a
//	// saved, live register, and we only ever exchange it between
//	// the real register and the gobuf. Hence, we treat it as a
//	// root during stack scanning, which means assembly that saves
//	// and restores it doesn't need write barriers. It's still
//	// typed as a pointer so that any other writes from Go get
//	// write barriers.
//	sp   uintptr  // 保存CPU的rsp寄存器的值
//	pc   uintptr  // 保存CPU的rip寄存器的值
//	g    guintptr // 记录当前这个gobuf对象属于哪个goroutine
//	ctxt unsafe.Pointer
//
//	// 保存系统调用的返回值，因为从系统调用返回之后如果p被其它工作线程抢占，
//	// 则这个goroutine会被放入全局运行队列被其它工作线程调度，其它线程需要知道系统
//	//调用的返回值。
//
//	ret sys.Uintreg
//	lr  uintptr
//
//	// 保存CPU的rip寄存器的值
//	bp uintptr // for GOEXPERIMENT=framepointer
//}
//
//
//// 前文所说的g结构体，它代表了一个goroutine
//type g struct {
//	// Stack parameters.
//	// stack describes the actual stack memory: [stack.lo, stack.hi).
//	// stackguard0 is the stack pointer compared in the Go stack growth
//	//prologue.
//
//	// It is stack.lo+StackGuard normally, but can be StackPreempt to
//	//trigger a preemption.
//
//	// stackguard1 is the stack pointer compared in the C stack growth
//	//prologue.
//
//	// It is stack.lo+StackGuard on g0 and gsignal stacks.
//	// It is ~0 on other goroutine stacks, to trigger a call to morestackc
//	//(and crash).
//
//	// 记录该goroutine使用的栈
//	stack stack // offset known to runtime/cgo
//	// 下面两个成员用于栈溢出检查，实现栈的自动伸缩，抢占调度也会用到stackguard0
//	stackguard0 uintptr // offset known to liblink
//	stackguard1 uintptr // offset known to liblink
//	......
//
//	// 此goroutine正在被哪个工作线程执行
//	m *m // current m; offset known to arm liblink
//	// 保存调度信息，主要是几个寄存器的值
//	sched gobuf
//	......
//	// schedlink字段指向全局运行队列中的下一个g，
//	//所有位于全局运行队列中的g形成一个链表
//	schedlink guintptr
//	......
//	// 抢占调度标志，如果需要抢占调度，设置preempt为true
//	preempt bool // preemption signal, duplicates stackguard0
//	= stackpreempt
//	......
//}
