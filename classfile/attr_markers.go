package classfile

type MarkerAttribute struct {
}

// 标记已过时
type DeprecatedAttribute struct {
	MarkerAttribute
}

// 标记源文件不存在、由编译器生成的类成员
type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
