package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type AASTORE struct {
	base.NoOperandsInstruction
}

type BASTORE struct {
	base.NoOperandsInstruction
}

type CASTORE struct {
	base.NoOperandsInstruction
}

type DASTORE struct {
	base.NoOperandsInstruction
}

type FASTORE struct {
	base.NoOperandsInstruction
}

type IASTORE struct {
	base.NoOperandsInstruction
}

type LASTORE struct {
	base.NoOperandsInstruction
}

type SASTORE struct {
	base.NoOperandsInstruction
}


func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}

//todo