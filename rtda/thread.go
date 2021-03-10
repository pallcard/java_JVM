package rtda

import "jvmgo/rtda/heap"

/**
线程
*/
type Thread struct {
	pc    int    //程序计数器
	stack *Stack //java虚拟机栈
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}

func (self *Thread) ClearStack() {
	self.stack.clear()
}

func (self *Thread) GetFrames() []*Frame {
	return self.stack.GetFrames()
}
