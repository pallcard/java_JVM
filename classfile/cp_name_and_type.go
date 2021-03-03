package classfile

/**
类型描述符
byte        B
short       S
char        C
int         I
long        J
float       F
double      D
引用         L+类完全限定名+分号
数组         [+数组元素类型描述符

字段描述符    字段类型描述符                              例如：S   Ljava.lang.Object;  [[D
方法描述符    (分号分割参数类型描述符)+返回值类型描述符            ()V   (Ljava.lang.Object;)F

*/

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16 //名称
	descriptorIndex uint16 //描述符
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
