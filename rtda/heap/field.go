package heap

import "jvmgo/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfFields := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfFields)
		fields[i].copyAttributes(cfFields)
	}
	return fields
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags & ACC_VOLATILE
}

func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags & ACC_TRANSIENT
}

func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueIndex(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}