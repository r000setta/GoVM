package classfile

type ClassFile struct {
	Magic        uint32
	MinorVersion uint16
	MajorVersion uint16
	//ConstantPool
	AccessFlag uint16
	ThisClass  uint16
	SuperClass uint16
	Interfaces []uint16
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
