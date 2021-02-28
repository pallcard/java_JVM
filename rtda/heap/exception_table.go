package heap

import "jvmgo/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlePc  int
	catchType *ClassRef
}



func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc: int(entry.StartPc()),
			endPc: int(entry.EndPc()),
			handlePc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}
	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handle := range self {
		if pc >= handle.startPc && pc < handle.endPc {
			return handle
		}
		catchClass := handle.catchType.ResolveClass()
		if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
			return handle
		}
	}
	return nil
}