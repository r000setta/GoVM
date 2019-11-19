package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT__MethodRef         = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Intger             = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameandType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Intger:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{Cp: cp}}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT__MethodRef:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{Cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceRefInfo{ConstantMemberRefInfo{Cp: cp}}
	case CONSTANT_NameandType:
		return &ConstantNameAndType{}
	//case CONSTANT_MethodType:
	//	return &ConstTypeI
	//case CONSTANT_MethodHandle:
	//	return &
	//case CONSTANT_InvokeDynamic:
	//	return
	default:
		panic("Constant pool error!")
	}
}
