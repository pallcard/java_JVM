package classfile

import "encoding/binary"

/**
用于读取class文件，并将其作为字节流进行处理
 */
type ClassReader struct {
	data []byte // byte是uint8别名
}

// 读u1
// 这里使用指针，故会修改掉ClassReader的值
func (self *ClassReader) readUint8() uint8 {
	// 读取第一个字节
	val := self.data[0]
	// 剔除读取的第一个字节
	self.data = self.data[1:]
	return val
}

// 读u2
// java虚拟机采用的大端序进行存储
// 大端序，低地址端存放高位字节
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 读u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
