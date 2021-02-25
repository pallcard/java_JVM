package heap

type SymRef struct {
	cp        *ConstantPool //运行时常量池指针
	className string        //完全限定名
	class     *Class        //解析后类结构指针
}

func (self *SymRef) ResolveClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}

