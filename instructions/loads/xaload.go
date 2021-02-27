package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

/**
 * 加载数组类型
 */

type XALOAD struct {
	base.Index8Instruction
}

//todo

type AALOAD struct {
	base.NoOperandsInstruction
}
type BALOAD struct {
	base.NoOperandsInstruction
}
type CALOAD struct {
	base.NoOperandsInstruction
}
type DALOAD struct {
	base.NoOperandsInstruction
}
type FALOAD struct {
	base.NoOperandsInstruction
}
type IALOAD struct {
	base.NoOperandsInstruction
}
type LALOAD struct {
	base.NoOperandsInstruction
}
type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
