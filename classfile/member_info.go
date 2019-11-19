package classfile

type MemberInfo struct {
	//Cp
	AccessFlag      uint16
	NameIndex       uint16
	DescriptorIndex uint16
	//Attributes []
}
