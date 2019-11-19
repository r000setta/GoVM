package classfile

import "encoding/binary"

type ClassReader struct {
	Data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.Data[0]
	self.Data = self.Data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.Data)
	self.Data = self.Data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.Data)
	self.Data = self.Data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.Data)
	self.Data = self.Data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.Data[:n]
	self.Data = self.Data[n:]
	return bytes
}
