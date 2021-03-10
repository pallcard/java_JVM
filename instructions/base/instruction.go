package base

import (
	"jvmgo/rtda"
)

/**
基本指令
1. 无操作数指令
2. 跳转指令
3. 单字节操作数指令
4. 双字节操作数指令
 */
type Instruction interface {
	// 提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令
	Execute(frame *rtda.Frame)
}

// 无操作数指令
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadUint16())
}

// 存储和加载类指令
// 索引由单字节操作数给出
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// 访问运行时常量池
// 索引由双字节操作数给出
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
