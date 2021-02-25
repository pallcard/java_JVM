package heap

import (
	"fmt"
	"jvmgo/classfile"
)

type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			intInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantLongInfo:
			intInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = intInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			intInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = intInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			intInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = intInfo.String()
		case *classfile.ConstantClassInfo:
			//intInfo := cpInfo.(*classfile.ConstantClassInfo)
			//consts[i] = intInfo.String()
			//todo
		case *classfile.ConstantFieldrefInfo:
			//intInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			//consts[i] = intInfo.String()
			//todo
		case *classfile.ConstantMethodrefInfo:
			//intInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			//consts[i] = intInfo.String()
			//todo
		case *classfile.ConstantInterfaceMethodrefInfo:
			//intInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			//consts[i] = intInfo.String()
			//todo
		default:
			//todo
		}
	}
	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
