package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32 //默认情况下执行跳转所需的字节码偏移量
	low           int32 //记录case的范围
	high          int32 //记录case的范围
	jumpOffsets   []int32 //索引表，存放high-low+1个int值
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding() //用于保证defaultOffset是4的倍数
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
