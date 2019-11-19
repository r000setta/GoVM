package classfile

import "fmt"

type ClassFile struct {
	Magic        uint32
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool ConstantPool
	AccessFlag   uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []uint16
	Fields       []*MemberInfo
	Methods      []*MemberInfo
	Attributes   []AttributeInfo
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("Magic number error!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.MajorVersion = reader.readUint16()
	self.MinorVersion = reader.readUint16()
	switch self.MajorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.MinorVersion == 0 {
			return
		}
	}
	panic("Version error!")
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.ConstantPool = readConstantPool(reader)
	self.AccessFlag = reader.readUint16()
	self.ThisClass = reader.readUint16()
	self.SuperClass = reader.readUint16()
	self.Interfaces = reader.readUint16s()
	self.Fields = readMembers(reader, self.ConstantPool)
	self.Methods = readMembers(reader, self.ConstantPool)
	self.Attributes = readAttributes(reader, self.ConstantPool)
}

func (self *ClassFile) getMajorVersion() uint16 {
	return self.MajorVersion
}

func (self *ClassFile) getMinorVersion() uint16 {
	return self.MinorVersion
}

func (self *ClassFile) getAccessFlag() uint16 {
	return self.AccessFlag
}

func (self *ClassFile) getFields() []*MemberInfo {
	return self.Fields
}

func (self *ClassFile) getMethods() []*MemberInfo {
	return self.Methods
}

func (self *ClassFile) ClassName() string {
	return self.ConstantPool.getClassName(self.ThisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.SuperClass > 0 {
		return self.ConstantPool.getClassName(self.SuperClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfacesName := make([]string, len(self.Interfaces))
	for i, cpIndex := range self.Interfaces {
		interfacesName[i] = self.ConstantPool.getClassName(cpIndex)
	}
	return interfacesName
}
