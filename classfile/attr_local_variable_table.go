package classfile


type LocalVariableTableEntry struct {
	startPc uint16
	lineNumber uint16
}

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}
//todo
func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}