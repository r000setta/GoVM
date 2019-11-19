package classfile

type MemberInfo struct {
	Cp              ConstantPool
	AccessFlag      uint16
	NameIndex       uint16
	DescriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		Cp:              cp,
		AccessFlag:      reader.readUint16(),
		NameIndex:       reader.readUint16(),
		DescriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) Name() string {
	return self.Cp.getUtf8(self.NameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.Cp.getUtf8(self.DescriptorIndex)
}
