package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/**
索引从2个字节变成4个字节
*/
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
