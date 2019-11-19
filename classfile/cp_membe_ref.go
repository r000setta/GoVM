package classfile

type ConstantMemberRefInfo struct {
	Cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberRefInfo) ClassName() string {
	return self.Cp.getClassName(self.nameAndTypeIndex)
}

func (self *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return self.Cp.getNameAndType(self.classIndex)
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfaceRefInfo struct {
	ConstantMemberRefInfo
}
