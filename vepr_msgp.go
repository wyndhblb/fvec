package vepr

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Resolution) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zxvk int32
		zxvk, err = dc.ReadInt32()
		(*z) = Resolution(zxvk)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Resolution) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Resolution) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Resolution) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zbzg int32
		zbzg, bts, err = msgp.ReadInt32Bytes(bts)
		(*z) = Resolution(zbzg)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Resolution) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Tag) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Value":
			z.Value, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Tag) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Tag) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tag) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Tag) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.StringPrefixSize + len(z.Value)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zhct uint32
			zhct, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zhct) {
				z.Vec = (z.Vec)[:zhct]
			} else {
				z.Vec = make([]float64, zhct)
			}
			for zajw := range z.Vec {
				z.Vec[zajw], err = dc.ReadFloat64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zajw := range z.Vec {
		err = en.WriteFloat64(z.Vec[zajw])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zajw := range z.Vec {
		o = msgp.AppendFloat64(o, z.Vec[zajw])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zxhx uint32
			zxhx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zxhx) {
				z.Vec = (z.Vec)[:zxhx]
			} else {
				z.Vec = make([]float64, zxhx)
			}
			for zajw := range z.Vec {
				z.Vec[zajw], bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Float64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zpks uint32
			zpks, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zpks) {
				z.Vec = (z.Vec)[:zpks]
			} else {
				z.Vec = make([]*VTDblDbl, zpks)
			}
			for zlqf := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zlqf] = nil
				} else {
					if z.Vec[zlqf] == nil {
						z.Vec[zlqf] = new(VTDblDbl)
					}
					var zjfb uint32
					zjfb, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zjfb > 0 {
						zjfb--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zlqf].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zlqf].Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLDblDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zlqf := range z.Vec {
		if z.Vec[zlqf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zlqf].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zlqf].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLDblDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zlqf := range z.Vec {
		if z.Vec[zlqf] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zlqf].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zlqf].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zeff uint32
			zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zeff) {
				z.Vec = (z.Vec)[:zeff]
			} else {
				z.Vec = make([]*VTDblDbl, zeff)
			}
			for zlqf := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zlqf] = nil
				} else {
					if z.Vec[zlqf] == nil {
						z.Vec[zlqf] = new(VTDblDbl)
					}
					var zrsw uint32
					zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zrsw > 0 {
						zrsw--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zlqf].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zlqf].Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLDblDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zlqf := range z.Vec {
		if z.Vec[zlqf] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLDblInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdnj uint32
	zdnj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdnj > 0 {
		zdnj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zobc uint32
			zobc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zobc) {
				z.Vec = (z.Vec)[:zobc]
			} else {
				z.Vec = make([]*VTDblInt, zobc)
			}
			for zxpk := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zxpk] = nil
				} else {
					if z.Vec[zxpk] == nil {
						z.Vec[zxpk] = new(VTDblInt)
					}
					var zsnv uint32
					zsnv, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zsnv > 0 {
						zsnv--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zxpk].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zxpk].Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLDblInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zxpk := range z.Vec {
		if z.Vec[zxpk] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zxpk].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zxpk].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLDblInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zxpk := range z.Vec {
		if z.Vec[zxpk] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zxpk].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zxpk].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zkgt uint32
	zkgt, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zkgt > 0 {
		zkgt--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zema uint32
			zema, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zema) {
				z.Vec = (z.Vec)[:zema]
			} else {
				z.Vec = make([]*VTDblInt, zema)
			}
			for zxpk := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zxpk] = nil
				} else {
					if z.Vec[zxpk] == nil {
						z.Vec[zxpk] = new(VTDblInt)
					}
					var zpez uint32
					zpez, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zpez > 0 {
						zpez--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zxpk].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zxpk].Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLDblInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zxpk := range z.Vec {
		if z.Vec[zxpk] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLDblStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqyh uint32
	zqyh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zyzr uint32
			zyzr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zyzr) {
				z.Vec = (z.Vec)[:zyzr]
			} else {
				z.Vec = make([]*VTDblStr, zyzr)
			}
			for zqke := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zqke] = nil
				} else {
					if z.Vec[zqke] == nil {
						z.Vec[zqke] = new(VTDblStr)
					}
					var zywj uint32
					zywj, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zywj > 0 {
						zywj--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zqke].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zqke].Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLDblStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zqke := range z.Vec {
		if z.Vec[zqke] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zqke].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zqke].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLDblStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zqke := range z.Vec {
		if z.Vec[zqke] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zqke].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zqke].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zzpf uint32
			zzpf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zzpf) {
				z.Vec = (z.Vec)[:zzpf]
			} else {
				z.Vec = make([]*VTDblStr, zzpf)
			}
			for zqke := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zqke] = nil
				} else {
					if z.Vec[zqke] == nil {
						z.Vec[zqke] = new(VTDblStr)
					}
					var zrfe uint32
					zrfe, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zrfe > 0 {
						zrfe--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zqke].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zqke].Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLDblStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zqke := range z.Vec {
		if z.Vec[zqke] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zqke].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztaf uint32
	ztaf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztaf > 0 {
		ztaf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zeth uint32
			zeth, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zeth) {
				z.Vec = (z.Vec)[:zeth]
			} else {
				z.Vec = make([]int64, zeth)
			}
			for zgmo := range z.Vec {
				z.Vec[zgmo], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zgmo := range z.Vec {
		err = en.WriteInt64(z.Vec[zgmo])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zgmo := range z.Vec {
		o = msgp.AppendInt64(o, z.Vec[zgmo])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsbz uint32
	zsbz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsbz > 0 {
		zsbz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zrjx uint32
			zrjx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zrjx) {
				z.Vec = (z.Vec)[:zrjx]
			} else {
				z.Vec = make([]int64, zrjx)
			}
			for zgmo := range z.Vec {
				z.Vec[zgmo], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwel uint32
	zwel, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwel > 0 {
		zwel--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zrbe uint32
			zrbe, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zrbe) {
				z.Vec = (z.Vec)[:zrbe]
			} else {
				z.Vec = make([]*VTIntDbl, zrbe)
			}
			for zawn := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zawn] = nil
				} else {
					if z.Vec[zawn] == nil {
						z.Vec[zawn] = new(VTIntDbl)
					}
					var zmfd uint32
					zmfd, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zmfd > 0 {
						zmfd--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zawn].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zawn].Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLIntDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zawn := range z.Vec {
		if z.Vec[zawn] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zawn].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zawn].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLIntDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zawn := range z.Vec {
		if z.Vec[zawn] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zawn].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zawn].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzdc uint32
	zzdc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzdc > 0 {
		zzdc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zelx uint32
			zelx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zelx) {
				z.Vec = (z.Vec)[:zelx]
			} else {
				z.Vec = make([]*VTIntDbl, zelx)
			}
			for zawn := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zawn] = nil
				} else {
					if z.Vec[zawn] == nil {
						z.Vec[zawn] = new(VTIntDbl)
					}
					var zbal uint32
					zbal, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zbal > 0 {
						zbal--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zawn].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zawn].Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLIntDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zawn := range z.Vec {
		if z.Vec[zawn] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLIntInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zkct uint32
	zkct, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zkct > 0 {
		zkct--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var ztmt uint32
			ztmt, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(ztmt) {
				z.Vec = (z.Vec)[:ztmt]
			} else {
				z.Vec = make([]*VTIntInt, ztmt)
			}
			for zjqz := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zjqz] = nil
				} else {
					if z.Vec[zjqz] == nil {
						z.Vec[zjqz] = new(VTIntInt)
					}
					var ztco uint32
					ztco, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for ztco > 0 {
						ztco--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zjqz].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zjqz].Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLIntInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zjqz := range z.Vec {
		if z.Vec[zjqz] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zjqz].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zjqz].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLIntInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zjqz := range z.Vec {
		if z.Vec[zjqz] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zjqz].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zjqz].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zana uint32
	zana, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zana > 0 {
		zana--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var ztyy uint32
			ztyy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(ztyy) {
				z.Vec = (z.Vec)[:ztyy]
			} else {
				z.Vec = make([]*VTIntInt, ztyy)
			}
			for zjqz := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zjqz] = nil
				} else {
					if z.Vec[zjqz] == nil {
						z.Vec[zjqz] = new(VTIntInt)
					}
					var zinl uint32
					zinl, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zinl > 0 {
						zinl--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zjqz].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zjqz].Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLIntInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zjqz := range z.Vec {
		if z.Vec[zjqz] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLIntStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zljy uint32
	zljy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zljy > 0 {
		zljy--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zixj uint32
			zixj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zixj) {
				z.Vec = (z.Vec)[:zixj]
			} else {
				z.Vec = make([]*VTIntStr, zixj)
			}
			for zare := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zare] = nil
				} else {
					if z.Vec[zare] == nil {
						z.Vec[zare] = new(VTIntStr)
					}
					var zrsc uint32
					zrsc, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zrsc > 0 {
						zrsc--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zare].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zare].Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLIntStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zare := range z.Vec {
		if z.Vec[zare] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zare].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zare].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLIntStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zare := range z.Vec {
		if z.Vec[zare] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zare].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zare].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zctn uint32
	zctn, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zctn > 0 {
		zctn--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zswy uint32
			zswy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zswy) {
				z.Vec = (z.Vec)[:zswy]
			} else {
				z.Vec = make([]*VTIntStr, zswy)
			}
			for zare := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zare] = nil
				} else {
					if z.Vec[zare] == nil {
						z.Vec[zare] = new(VTIntStr)
					}
					var znsg uint32
					znsg, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for znsg > 0 {
						znsg--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zare].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zare].Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLIntStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zare := range z.Vec {
		if z.Vec[zare] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zare].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsvm uint32
	zsvm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsvm > 0 {
		zsvm--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zaoz uint32
			zaoz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zaoz) {
				z.Vec = (z.Vec)[:zaoz]
			} else {
				z.Vec = make([]string, zaoz)
			}
			for zrus := range z.Vec {
				z.Vec[zrus], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zrus := range z.Vec {
		err = en.WriteString(z.Vec[zrus])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zrus := range z.Vec {
		o = msgp.AppendString(o, z.Vec[zrus])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zfzb uint32
	zfzb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zfzb > 0 {
		zfzb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zsbo uint32
			zsbo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zsbo) {
				z.Vec = (z.Vec)[:zsbo]
			} else {
				z.Vec = make([]string, zsbo)
			}
			for zrus := range z.Vec {
				z.Vec[zrus], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zrus := range z.Vec {
		s += msgp.StringPrefixSize + len(z.Vec[zrus])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqgz uint32
	zqgz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqgz > 0 {
		zqgz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zsnw uint32
			zsnw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zsnw) {
				z.Vec = (z.Vec)[:zsnw]
			} else {
				z.Vec = make([]*VTStrDbl, zsnw)
			}
			for zjif := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zjif] = nil
				} else {
					if z.Vec[zjif] == nil {
						z.Vec[zjif] = new(VTStrDbl)
					}
					var ztls uint32
					ztls, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for ztls > 0 {
						ztls--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zjif].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zjif].Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLStrDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zjif := range z.Vec {
		if z.Vec[zjif] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zjif].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zjif].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zjif := range z.Vec {
		if z.Vec[zjif] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zjif].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zjif].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmvo uint32
	zmvo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmvo > 0 {
		zmvo--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zigk uint32
			zigk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zigk) {
				z.Vec = (z.Vec)[:zigk]
			} else {
				z.Vec = make([]*VTStrDbl, zigk)
			}
			for zjif := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zjif] = nil
				} else {
					if z.Vec[zjif] == nil {
						z.Vec[zjif] = new(VTStrDbl)
					}
					var zopb uint32
					zopb, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zopb > 0 {
						zopb--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zjif].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zjif].Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLStrDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zjif := range z.Vec {
		if z.Vec[zjif] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zjif].Key) + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zedl uint32
	zedl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zedl > 0 {
		zedl--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zupd uint32
			zupd, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zupd) {
				z.Vec = (z.Vec)[:zupd]
			} else {
				z.Vec = make([]*VTStrInt, zupd)
			}
			for zuop := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zuop] = nil
				} else {
					if z.Vec[zuop] == nil {
						z.Vec[zuop] = new(VTStrInt)
					}
					var zome uint32
					zome, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zome > 0 {
						zome--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zuop].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zuop].Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLStrInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zuop := range z.Vec {
		if z.Vec[zuop] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zuop].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zuop].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zuop := range z.Vec {
		if z.Vec[zuop] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zuop].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zuop].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrvj uint32
	zrvj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrvj > 0 {
		zrvj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zarz uint32
			zarz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zarz) {
				z.Vec = (z.Vec)[:zarz]
			} else {
				z.Vec = make([]*VTStrInt, zarz)
			}
			for zuop := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zuop] = nil
				} else {
					if z.Vec[zuop] == nil {
						z.Vec[zuop] = new(VTStrInt)
					}
					var zknt uint32
					zknt, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zknt > 0 {
						zknt--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zuop].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zuop].Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLStrInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zuop := range z.Vec {
		if z.Vec[zuop] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zuop].Key) + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zucw uint32
	zucw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zucw > 0 {
		zucw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zlsx uint32
			zlsx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zlsx) {
				z.Vec = (z.Vec)[:zlsx]
			} else {
				z.Vec = make([]*VTStrStr, zlsx)
			}
			for zxye := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zxye] = nil
				} else {
					if z.Vec[zxye] == nil {
						z.Vec[zxye] = new(VTStrStr)
					}
					var zbgy uint32
					zbgy, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zbgy > 0 {
						zbgy--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zxye].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zxye].Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VLStrStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zxye := range z.Vec {
		if z.Vec[zxye] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zxye].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zxye].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VLStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zxye := range z.Vec {
		if z.Vec[zxye] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zxye].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zxye].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrao uint32
	zrao, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrao > 0 {
		zrao--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zmbt uint32
			zmbt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zmbt) {
				z.Vec = (z.Vec)[:zmbt]
			} else {
				z.Vec = make([]*VTStrStr, zmbt)
			}
			for zxye := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zxye] = nil
				} else {
					if z.Vec[zxye] == nil {
						z.Vec[zxye] = new(VTStrStr)
					}
					var zvls uint32
					zvls, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zvls > 0 {
						zvls--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zxye].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zxye].Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VLStrStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zxye := range z.Vec {
		if z.Vec[zxye] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zxye].Key) + 6 + msgp.StringPrefixSize + len(z.Vec[zxye].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjfj uint32
	zjfj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjfj > 0 {
		zjfj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "name"
	err = en.Append(0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "name"
	o = append(o, 0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzak uint32
	zzak, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzak > 0 {
		zzak--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbtz uint32
	zbtz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbtz > 0 {
		zbtz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "name"
	err = en.Append(0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "name"
	o = append(o, 0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsym uint32
	zsym, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsym > 0 {
		zsym--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgeu uint32
	zgeu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgeu > 0 {
		zgeu--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "name"
	err = en.Append(0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "name"
	o = append(o, 0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdtr uint32
	zdtr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdtr > 0 {
		zdtr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zzqm uint32
	zzqm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zzqm > 0 {
		zzqm--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPDblDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPDblDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdqi uint32
	zdqi, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdqi > 0 {
		zdqi--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPDblDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPDblInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zyco uint32
	zyco, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zyco > 0 {
		zyco--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPDblInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPDblInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhgh uint32
	zhgh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhgh > 0 {
		zhgh--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPDblInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPDblStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zovg uint32
	zovg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zovg > 0 {
		zovg--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPDblStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPDblStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsey uint32
	zsey, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsey > 0 {
		zsey--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPDblStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcjp uint32
	zcjp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcjp > 0 {
		zcjp--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPIntDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPIntDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjhy uint32
	zjhy, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjhy > 0 {
		zjhy--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPIntDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPIntInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var znuf uint32
	znuf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for znuf > 0 {
		znuf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPIntInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPIntInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var znjj uint32
	znjj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for znjj > 0 {
		znjj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPIntInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPIntStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zhhj uint32
	zhhj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhhj > 0 {
		zhhj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPIntStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPIntStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zuvr uint32
	zuvr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zuvr > 0 {
		zuvr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPIntStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zusq uint32
	zusq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zusq > 0 {
		zusq--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPStrDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zfgq uint32
	zfgq, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zfgq > 0 {
		zfgq--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPStrDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zvml uint32
	zvml, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zvml > 0 {
		zvml--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPStrInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpyv uint32
	zpyv, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpyv > 0 {
		zpyv--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPStrInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntTPStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zlur uint32
	zlur, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zlur > 0 {
		zlur--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMIntTPStrStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Name"
	err = en.Append(0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMIntTPStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Name"
	o = append(o, 0x81, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMIntTPStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zupi uint32
	zupi, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zupi > 0 {
		zupi--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMIntTPStrStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbmy uint32
	zbmy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbmy > 0 {
		zbmy--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zarl uint32
			zarl, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zarl > 0 {
				z.Vec = make(map[string]float64, zarl)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zarl > 0 {
				zarl--
				var zfvi string
				var zzrg float64
				zfvi, err = dc.ReadString()
				if err != nil {
					return
				}
				zzrg, err = dc.ReadFloat64()
				if err != nil {
					return
				}
				z.Vec[zfvi] = zzrg
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zfvi, zzrg := range z.Vec {
		err = en.WriteString(zfvi)
		if err != nil {
			return
		}
		err = en.WriteFloat64(zzrg)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zfvi, zzrg := range z.Vec {
		o = msgp.AppendString(o, zfvi)
		o = msgp.AppendFloat64(o, zzrg)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zctz uint32
	zctz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zctz > 0 {
		zctz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zljl uint32
			zljl, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zljl > 0 {
				z.Vec = make(map[string]float64, zljl)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zljl > 0 {
				var zfvi string
				var zzrg float64
				zljl--
				zfvi, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zzrg, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
				z.Vec[zfvi] = zzrg
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zfvi, zzrg := range z.Vec {
			_ = zzrg
			s += msgp.StringPrefixSize + len(zfvi) + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zmlx uint32
	zmlx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zmlx > 0 {
		zmlx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zvbw uint32
			zvbw, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zvbw > 0 {
				z.Vec = make(map[string]int64, zvbw)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zvbw > 0 {
				zvbw--
				var zziv string
				var zabj int64
				zziv, err = dc.ReadString()
				if err != nil {
					return
				}
				zabj, err = dc.ReadInt64()
				if err != nil {
					return
				}
				z.Vec[zziv] = zabj
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zziv, zabj := range z.Vec {
		err = en.WriteString(zziv)
		if err != nil {
			return
		}
		err = en.WriteInt64(zabj)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zziv, zabj := range z.Vec {
		o = msgp.AppendString(o, zziv)
		o = msgp.AppendInt64(o, zabj)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zgvb uint32
	zgvb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zgvb > 0 {
		zgvb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zqzg uint32
			zqzg, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zqzg > 0 {
				z.Vec = make(map[string]int64, zqzg)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zqzg > 0 {
				var zziv string
				var zabj int64
				zqzg--
				zziv, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zabj, bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
				z.Vec[zziv] = zabj
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zziv, zabj := range z.Vec {
			_ = zabj
			s += msgp.StringPrefixSize + len(zziv) + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsdj uint32
	zsdj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsdj > 0 {
		zsdj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zsgp uint32
			zsgp, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zsgp > 0 {
				z.Vec = make(map[string]string, zsgp)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zsgp > 0 {
				zsgp--
				var zexy string
				var zakb string
				zexy, err = dc.ReadString()
				if err != nil {
					return
				}
				zakb, err = dc.ReadString()
				if err != nil {
					return
				}
				z.Vec[zexy] = zakb
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zexy, zakb := range z.Vec {
		err = en.WriteString(zexy)
		if err != nil {
			return
		}
		err = en.WriteString(zakb)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zexy, zakb := range z.Vec {
		o = msgp.AppendString(o, zexy)
		o = msgp.AppendString(o, zakb)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zngc uint32
	zngc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zngc > 0 {
		zngc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zwfl uint32
			zwfl, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zwfl > 0 {
				z.Vec = make(map[string]string, zwfl)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zwfl > 0 {
				var zexy string
				var zakb string
				zwfl--
				zexy, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zakb, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Vec[zexy] = zakb
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zexy, zakb := range z.Vec {
			_ = zakb
			s += msgp.StringPrefixSize + len(zexy) + msgp.StringPrefixSize + len(zakb)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zuff uint32
	zuff, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zuff > 0 {
		zuff--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zmow uint32
			zmow, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zmow > 0 {
				z.Vec = make(map[string]*VTDblDbl, zmow)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmow > 0 {
				zmow--
				var zdif string
				var zibu *VTDblDbl
				zdif, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zibu = nil
				} else {
					if zibu == nil {
						zibu = new(VTDblDbl)
					}
					var zdit uint32
					zdit, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zdit > 0 {
						zdit--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zibu.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							zibu.Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zdif] = zibu
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPDblDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zdif, zibu := range z.Vec {
		err = en.WriteString(zdif)
		if err != nil {
			return
		}
		if zibu == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zibu.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zibu.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPDblDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zdif, zibu := range z.Vec {
		o = msgp.AppendString(o, zdif)
		if zibu == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, zibu.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zibu.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zslz uint32
	zslz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zslz > 0 {
		zslz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zoqj uint32
			zoqj, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zoqj > 0 {
				z.Vec = make(map[string]*VTDblDbl, zoqj)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zoqj > 0 {
				var zdif string
				var zibu *VTDblDbl
				zoqj--
				zdif, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zibu = nil
				} else {
					if zibu == nil {
						zibu = new(VTDblDbl)
					}
					var zmqr uint32
					zmqr, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zmqr > 0 {
						zmqr--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zibu.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zibu.Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zdif] = zibu
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPDblDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zdif, zibu := range z.Vec {
			_ = zibu
			s += msgp.StringPrefixSize + len(zdif)
			if zibu == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Float64Size + 6 + msgp.Float64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPDblInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ziyx uint32
	ziyx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ziyx > 0 {
		ziyx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zyes uint32
			zyes, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zyes > 0 {
				z.Vec = make(map[string]*VTDblInt, zyes)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zyes > 0 {
				zyes--
				var ztic string
				var ztoj *VTDblInt
				ztic, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					ztoj = nil
				} else {
					if ztoj == nil {
						ztoj = new(VTDblInt)
					}
					var zxzy uint32
					zxzy, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zxzy > 0 {
						zxzy--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							ztoj.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							ztoj.Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[ztic] = ztoj
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPDblInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for ztic, ztoj := range z.Vec {
		err = en.WriteString(ztic)
		if err != nil {
			return
		}
		if ztoj == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(ztoj.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(ztoj.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPDblInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for ztic, ztoj := range z.Vec {
		o = msgp.AppendString(o, ztic)
		if ztoj == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, ztoj.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, ztoj.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zfro uint32
	zfro, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zfro > 0 {
		zfro--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zrod uint32
			zrod, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zrod > 0 {
				z.Vec = make(map[string]*VTDblInt, zrod)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zrod > 0 {
				var ztic string
				var ztoj *VTDblInt
				zrod--
				ztic, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					ztoj = nil
				} else {
					if ztoj == nil {
						ztoj = new(VTDblInt)
					}
					var zmbn uint32
					zmbn, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zmbn > 0 {
						zmbn--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							ztoj.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							ztoj.Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[ztic] = ztoj
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPDblInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for ztic, ztoj := range z.Vec {
			_ = ztoj
			s += msgp.StringPrefixSize + len(ztic)
			if ztoj == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Float64Size + 6 + msgp.Int64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPDblStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrwc uint32
	zrwc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrwc > 0 {
		zrwc--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zjpm uint32
			zjpm, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zjpm > 0 {
				z.Vec = make(map[string]*VTDblStr, zjpm)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zjpm > 0 {
				zjpm--
				var zdrz string
				var znpn *VTDblStr
				zdrz, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					znpn = nil
				} else {
					if znpn == nil {
						znpn = new(VTDblStr)
					}
					var zhdt uint32
					zhdt, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zhdt > 0 {
						zhdt--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							znpn.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							znpn.Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zdrz] = znpn
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPDblStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zdrz, znpn := range z.Vec {
		err = en.WriteString(zdrz)
		if err != nil {
			return
		}
		if znpn == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(znpn.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(znpn.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPDblStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zdrz, znpn := range z.Vec {
		o = msgp.AppendString(o, zdrz)
		if znpn == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, znpn.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, znpn.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjmh uint32
	zjmh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjmh > 0 {
		zjmh--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zayo uint32
			zayo, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zayo > 0 {
				z.Vec = make(map[string]*VTDblStr, zayo)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zayo > 0 {
				var zdrz string
				var znpn *VTDblStr
				zayo--
				zdrz, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					znpn = nil
				} else {
					if znpn == nil {
						znpn = new(VTDblStr)
					}
					var zrsu uint32
					zrsu, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zrsu > 0 {
						zrsu--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							znpn.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							znpn.Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zdrz] = znpn
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPDblStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zdrz, znpn := range z.Vec {
			_ = znpn
			s += msgp.StringPrefixSize + len(zdrz)
			if znpn == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(znpn.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zvgz uint32
	zvgz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zvgz > 0 {
		zvgz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zhbk uint32
			zhbk, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zhbk > 0 {
				z.Vec = make(map[string]*VTIntDbl, zhbk)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zhbk > 0 {
				zhbk--
				var zfum string
				var zaps *VTIntDbl
				zfum, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zaps = nil
				} else {
					if zaps == nil {
						zaps = new(VTIntDbl)
					}
					var zmyy uint32
					zmyy, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zmyy > 0 {
						zmyy--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zaps.Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							zaps.Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zfum] = zaps
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPIntDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zfum, zaps := range z.Vec {
		err = en.WriteString(zfum)
		if err != nil {
			return
		}
		if zaps == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zaps.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zaps.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPIntDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zfum, zaps := range z.Vec {
		o = msgp.AppendString(o, zfum)
		if zaps == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zaps.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zaps.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztej uint32
	ztej, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztej > 0 {
		ztej--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zvgw uint32
			zvgw, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zvgw > 0 {
				z.Vec = make(map[string]*VTIntDbl, zvgw)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zvgw > 0 {
				var zfum string
				var zaps *VTIntDbl
				zvgw--
				zfum, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zaps = nil
				} else {
					if zaps == nil {
						zaps = new(VTIntDbl)
					}
					var zffb uint32
					zffb, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zffb > 0 {
						zffb--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zaps.Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zaps.Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zfum] = zaps
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPIntDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zfum, zaps := range z.Vec {
			_ = zaps
			s += msgp.StringPrefixSize + len(zfum)
			if zaps == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Int64Size + 6 + msgp.Float64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPIntInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbae uint32
	zbae, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbae > 0 {
		zbae--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zreu uint32
			zreu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zreu > 0 {
				z.Vec = make(map[string]*VTIntInt, zreu)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zreu > 0 {
				zreu--
				var zbgg string
				var zcnq *VTIntInt
				zbgg, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zcnq = nil
				} else {
					if zcnq == nil {
						zcnq = new(VTIntInt)
					}
					var znuz uint32
					znuz, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for znuz > 0 {
						znuz--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zcnq.Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							zcnq.Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zbgg] = zcnq
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPIntInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zbgg, zcnq := range z.Vec {
		err = en.WriteString(zbgg)
		if err != nil {
			return
		}
		if zcnq == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zcnq.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zcnq.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPIntInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zbgg, zcnq := range z.Vec {
		o = msgp.AppendString(o, zbgg)
		if zcnq == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zcnq.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zcnq.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjqx uint32
	zjqx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjqx > 0 {
		zjqx--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zmzo uint32
			zmzo, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zmzo > 0 {
				z.Vec = make(map[string]*VTIntInt, zmzo)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmzo > 0 {
				var zbgg string
				var zcnq *VTIntInt
				zmzo--
				zbgg, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zcnq = nil
				} else {
					if zcnq == nil {
						zcnq = new(VTIntInt)
					}
					var ztar uint32
					ztar, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for ztar > 0 {
						ztar--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zcnq.Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zcnq.Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zbgg] = zcnq
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPIntInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zbgg, zcnq := range z.Vec {
			_ = zcnq
			s += msgp.StringPrefixSize + len(zbgg)
			if zcnq == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Int64Size + 6 + msgp.Int64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPIntStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zmsv uint32
	zmsv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zmsv > 0 {
		zmsv--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zyba uint32
			zyba, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zyba > 0 {
				z.Vec = make(map[string]*VTIntStr, zyba)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zyba > 0 {
				zyba--
				var zkut string
				var zmyg *VTIntStr
				zkut, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zmyg = nil
				} else {
					if zmyg == nil {
						zmyg = new(VTIntStr)
					}
					var zddv uint32
					zddv, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zddv > 0 {
						zddv--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zmyg.Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							zmyg.Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zkut] = zmyg
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPIntStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zkut, zmyg := range z.Vec {
		err = en.WriteString(zkut)
		if err != nil {
			return
		}
		if zmyg == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zmyg.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(zmyg.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPIntStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zkut, zmyg := range z.Vec {
		o = msgp.AppendString(o, zkut)
		if zmyg == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zmyg.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, zmyg.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zoxi uint32
	zoxi, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zoxi > 0 {
		zoxi--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zfsf uint32
			zfsf, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zfsf > 0 {
				z.Vec = make(map[string]*VTIntStr, zfsf)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zfsf > 0 {
				var zkut string
				var zmyg *VTIntStr
				zfsf--
				zkut, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zmyg = nil
				} else {
					if zmyg == nil {
						zmyg = new(VTIntStr)
					}
					var zgpy uint32
					zgpy, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zgpy > 0 {
						zgpy--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zmyg.Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zmyg.Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zkut] = zmyg
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPIntStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zkut, zmyg := range z.Vec {
			_ = zmyg
			s += msgp.StringPrefixSize + len(zkut)
			if zmyg == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(zmyg.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbey uint32
	zbey, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbey > 0 {
		zbey--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var ztnd uint32
			ztnd, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && ztnd > 0 {
				z.Vec = make(map[string]*VTStrDbl, ztnd)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for ztnd > 0 {
				ztnd--
				var zclm string
				var zxiu *VTStrDbl
				zclm, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zxiu = nil
				} else {
					if zxiu == nil {
						zxiu = new(VTStrDbl)
					}
					var ztja uint32
					ztja, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for ztja > 0 {
						ztja--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zxiu.Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							zxiu.Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zclm] = zxiu
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPStrDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zclm, zxiu := range z.Vec {
		err = en.WriteString(zclm)
		if err != nil {
			return
		}
		if zxiu == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(zxiu.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zxiu.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zclm, zxiu := range z.Vec {
		o = msgp.AppendString(o, zclm)
		if zxiu == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zxiu.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zxiu.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zyrr uint32
	zyrr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyrr > 0 {
		zyrr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zded uint32
			zded, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zded > 0 {
				z.Vec = make(map[string]*VTStrDbl, zded)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zded > 0 {
				var zclm string
				var zxiu *VTStrDbl
				zded--
				zclm, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zxiu = nil
				} else {
					if zxiu == nil {
						zxiu = new(VTStrDbl)
					}
					var zmiy uint32
					zmiy, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zmiy > 0 {
						zmiy--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zxiu.Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							zxiu.Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zclm] = zxiu
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPStrDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zclm, zxiu := range z.Vec {
			_ = zxiu
			s += msgp.StringPrefixSize + len(zclm)
			if zxiu == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zxiu.Key) + 6 + msgp.Float64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ziog uint32
	ziog, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ziog > 0 {
		ziog--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zjhy uint32
			zjhy, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zjhy > 0 {
				z.Vec = make(map[string]*VTStrInt, zjhy)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zjhy > 0 {
				zjhy--
				var zkpr string
				var zucj *VTStrInt
				zkpr, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zucj = nil
				} else {
					if zucj == nil {
						zucj = new(VTStrInt)
					}
					var zevw uint32
					zevw, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zevw > 0 {
						zevw--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zucj.Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							zucj.Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zkpr] = zucj
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPStrInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zkpr, zucj := range z.Vec {
		err = en.WriteString(zkpr)
		if err != nil {
			return
		}
		if zucj == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(zucj.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zucj.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zkpr, zucj := range z.Vec {
		o = msgp.AppendString(o, zkpr)
		if zucj == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zucj.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zucj.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbtc uint32
	zbtc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbtc > 0 {
		zbtc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zmkf uint32
			zmkf, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zmkf > 0 {
				z.Vec = make(map[string]*VTStrInt, zmkf)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmkf > 0 {
				var zkpr string
				var zucj *VTStrInt
				zmkf--
				zkpr, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zucj = nil
				} else {
					if zucj == nil {
						zucj = new(VTStrInt)
					}
					var zrdg uint32
					zrdg, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zrdg > 0 {
						zrdg--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zucj.Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							zucj.Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zkpr] = zucj
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPStrInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zkpr, zucj := range z.Vec {
			_ = zucj
			s += msgp.StringPrefixSize + len(zkpr)
			if zucj == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zucj.Key) + 6 + msgp.Int64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zvgq uint32
	zvgq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zvgq > 0 {
		zvgq--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zmzc uint32
			zmzc, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zmzc > 0 {
				z.Vec = make(map[string]*VTStrStr, zmzc)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmzc > 0 {
				zmzc--
				var zxqw string
				var zpzw *VTStrStr
				zxqw, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zpzw = nil
				} else {
					if zpzw == nil {
						zpzw = new(VTStrStr)
					}
					var zkvc uint32
					zkvc, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zkvc > 0 {
						zkvc--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zpzw.Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							zpzw.Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zxqw] = zpzw
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VMStrTPStrStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Vec"
	err = en.Append(0xa3, 0x56, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zxqw, zpzw := range z.Vec {
		err = en.WriteString(zxqw)
		if err != nil {
			return
		}
		if zpzw == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(zpzw.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(zpzw.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrTPStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Vec"
	o = append(o, 0xa3, 0x56, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zxqw, zpzw := range z.Vec {
		o = msgp.AppendString(o, zxqw)
		if zpzw == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zpzw.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, zpzw.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxas uint32
	zxas, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxas > 0 {
		zxas--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Vec":
			var zjkd uint32
			zjkd, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zjkd > 0 {
				z.Vec = make(map[string]*VTStrStr, zjkd)
			} else if len(z.Vec) > 0 {
				for key, _ := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zjkd > 0 {
				var zxqw string
				var zpzw *VTStrStr
				zjkd--
				zxqw, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zpzw = nil
				} else {
					if zpzw == nil {
						zpzw = new(VTStrStr)
					}
					var zsye uint32
					zsye, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zsye > 0 {
						zsye--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zpzw.Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							zpzw.Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
				z.Vec[zxqw] = zpzw
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VMStrTPStrStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zxqw, zpzw := range z.Vec {
			_ = zpzw
			s += msgp.StringPrefixSize + len(zxqw)
			if zpzw == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zpzw.Key) + 6 + msgp.StringPrefixSize + len(zpzw.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VName) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zley uint32
	zley, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zley > 0 {
		zley--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "res":
			{
				var zeql int32
				zeql, err = dc.ReadInt32()
				z.Resolution = Resolution(zeql)
			}
			if err != nil {
				return
			}
		case "ttl":
			z.Ttl, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "uid":
			z.XUniqueStr, err = dc.ReadString()
			if err != nil {
				return
			}
		case "tags":
			var zlhq uint32
			zlhq, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(zlhq) {
				z.Tags = (z.Tags)[:zlhq]
			} else {
				z.Tags = make([]*Tag, zlhq)
			}
			for zofl := range z.Tags {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Tags[zofl] = nil
				} else {
					if z.Tags[zofl] == nil {
						z.Tags[zofl] = new(Tag)
					}
					var zgbp uint32
					zgbp, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zgbp > 0 {
						zgbp--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Name":
							z.Tags[zofl].Name, err = dc.ReadString()
							if err != nil {
								return
							}
						case "Value":
							z.Tags[zofl].Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VName) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "key"
	err = en.Append(0x85, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "res"
	err = en.Append(0xa3, 0x72, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Resolution))
	if err != nil {
		return
	}
	// write "ttl"
	err = en.Append(0xa3, 0x74, 0x74, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.Ttl)
	if err != nil {
		return
	}
	// write "uid"
	err = en.Append(0xa3, 0x75, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.XUniqueStr)
	if err != nil {
		return
	}
	// write "tags"
	err = en.Append(0xa4, 0x74, 0x61, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Tags)))
	if err != nil {
		return
	}
	for zofl := range z.Tags {
		if z.Tags[zofl] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "Name"
			err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Tags[zofl].Name)
			if err != nil {
				return
			}
			// write "Value"
			err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Tags[zofl].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VName) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "key"
	o = append(o, 0x85, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "res"
	o = append(o, 0xa3, 0x72, 0x65, 0x73)
	o = msgp.AppendInt32(o, int32(z.Resolution))
	// string "ttl"
	o = append(o, 0xa3, 0x74, 0x74, 0x6c)
	o = msgp.AppendUint32(o, z.Ttl)
	// string "uid"
	o = append(o, 0xa3, 0x75, 0x69, 0x64)
	o = msgp.AppendString(o, z.XUniqueStr)
	// string "tags"
	o = append(o, 0xa4, 0x74, 0x61, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Tags)))
	for zofl := range z.Tags {
		if z.Tags[zofl] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "Name"
			o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Tags[zofl].Name)
			// string "Value"
			o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Tags[zofl].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VName) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var znbp uint32
	znbp, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for znbp > 0 {
		znbp--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "res":
			{
				var zbgh int32
				zbgh, bts, err = msgp.ReadInt32Bytes(bts)
				z.Resolution = Resolution(zbgh)
			}
			if err != nil {
				return
			}
		case "ttl":
			z.Ttl, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "uid":
			z.XUniqueStr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "tags":
			var zmku uint32
			zmku, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(zmku) {
				z.Tags = (z.Tags)[:zmku]
			} else {
				z.Tags = make([]*Tag, zmku)
			}
			for zofl := range z.Tags {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Tags[zofl] = nil
				} else {
					if z.Tags[zofl] == nil {
						z.Tags[zofl] = new(Tag)
					}
					var zidj uint32
					zidj, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zidj > 0 {
						zidj--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Name":
							z.Tags[zofl].Name, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "Value":
							z.Tags[zofl].Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VName) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 4 + msgp.Int32Size + 4 + msgp.Uint32Size + 4 + msgp.StringPrefixSize + len(z.XUniqueStr) + 5 + msgp.ArrayHeaderSize
	for zofl := range z.Tags {
		if z.Tags[zofl] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 5 + msgp.StringPrefixSize + len(z.Tags[zofl].Name) + 6 + msgp.StringPrefixSize + len(z.Tags[zofl].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpbh uint32
	zpbh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpbh > 0 {
		zpbh--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zcsj uint32
			zcsj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zcsj) {
				z.Vec = (z.Vec)[:zcsj]
			} else {
				z.Vec = make([]float64, zcsj)
			}
			for zumm := range z.Vec {
				z.Vec[zumm], err = dc.ReadFloat64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zumm := range z.Vec {
		err = en.WriteFloat64(z.Vec[zumm])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zumm := range z.Vec {
		o = msgp.AppendFloat64(o, z.Vec[zumm])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmjj uint32
	zmjj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmjj > 0 {
		zmjj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zxzu uint32
			zxzu, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zxzu) {
				z.Vec = (z.Vec)[:zxzu]
			} else {
				z.Vec = make([]float64, zxzu)
			}
			for zumm := range z.Vec {
				z.Vec[zumm], bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Float64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsnb uint32
	zsnb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsnb > 0 {
		zsnb--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zalq uint32
			zalq, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zalq) {
				z.Vec = (z.Vec)[:zalq]
			} else {
				z.Vec = make([]*VTDblDbl, zalq)
			}
			for zaii := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zaii] = nil
				} else {
					if z.Vec[zaii] == nil {
						z.Vec[zaii] = new(VTDblDbl)
					}
					var zswq uint32
					zswq, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zswq > 0 {
						zswq--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zaii].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zaii].Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSDblDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zaii := range z.Vec {
		if z.Vec[zaii] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zaii].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zaii].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSDblDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zaii := range z.Vec {
		if z.Vec[zaii] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zaii].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zaii].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpoq uint32
	zpoq, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpoq > 0 {
		zpoq--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zgnc uint32
			zgnc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zgnc) {
				z.Vec = (z.Vec)[:zgnc]
			} else {
				z.Vec = make([]*VTDblDbl, zgnc)
			}
			for zaii := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zaii] = nil
				} else {
					if z.Vec[zaii] == nil {
						z.Vec[zaii] = new(VTDblDbl)
					}
					var zzga uint32
					zzga, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zzga > 0 {
						zzga--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zaii].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zaii].Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSDblDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zaii := range z.Vec {
		if z.Vec[zaii] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSDblInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zink uint32
	zink, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zink > 0 {
		zink--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zqki uint32
			zqki, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zqki) {
				z.Vec = (z.Vec)[:zqki]
			} else {
				z.Vec = make([]*VTDblInt, zqki)
			}
			for zcza := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zcza] = nil
				} else {
					if z.Vec[zcza] == nil {
						z.Vec[zcza] = new(VTDblInt)
					}
					var zbaa uint32
					zbaa, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zbaa > 0 {
						zbaa--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zcza].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zcza].Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSDblInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zcza := range z.Vec {
		if z.Vec[zcza] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zcza].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zcza].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSDblInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zcza := range z.Vec {
		if z.Vec[zcza] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zcza].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zcza].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztky uint32
	ztky, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztky > 0 {
		ztky--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zhda uint32
			zhda, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zhda) {
				z.Vec = (z.Vec)[:zhda]
			} else {
				z.Vec = make([]*VTDblInt, zhda)
			}
			for zcza := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zcza] = nil
				} else {
					if z.Vec[zcza] == nil {
						z.Vec[zcza] = new(VTDblInt)
					}
					var zopo uint32
					zopo, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zopo > 0 {
						zopo--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zcza].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zcza].Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSDblInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zcza := range z.Vec {
		if z.Vec[zcza] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSDblStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zlqi uint32
	zlqi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zlqi > 0 {
		zlqi--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zexs uint32
			zexs, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zexs) {
				z.Vec = (z.Vec)[:zexs]
			} else {
				z.Vec = make([]*VTDblStr, zexs)
			}
			for zvfo := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zvfo] = nil
				} else {
					if z.Vec[zvfo] == nil {
						z.Vec[zvfo] = new(VTDblStr)
					}
					var zfzx uint32
					zfzx, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zfzx > 0 {
						zfzx--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zvfo].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zvfo].Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSDblStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zvfo := range z.Vec {
		if z.Vec[zvfo] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zvfo].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zvfo].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSDblStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zvfo := range z.Vec {
		if z.Vec[zvfo] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zvfo].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zvfo].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzrk uint32
	zzrk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzrk > 0 {
		zzrk--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zczt uint32
			zczt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zczt) {
				z.Vec = (z.Vec)[:zczt]
			} else {
				z.Vec = make([]*VTDblStr, zczt)
			}
			for zvfo := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zvfo] = nil
				} else {
					if z.Vec[zvfo] == nil {
						z.Vec[zvfo] = new(VTDblStr)
					}
					var zxcd uint32
					zxcd, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zxcd > 0 {
						zxcd--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zvfo].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zvfo].Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSDblStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zvfo := range z.Vec {
		if z.Vec[zvfo] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zvfo].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zuyz uint32
	zuyz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zuyz > 0 {
		zuyz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zhrc uint32
			zhrc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zhrc) {
				z.Vec = (z.Vec)[:zhrc]
			} else {
				z.Vec = make([]int64, zhrc)
			}
			for zjjf := range z.Vec {
				z.Vec[zjjf], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zjjf := range z.Vec {
		err = en.WriteInt64(z.Vec[zjjf])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zjjf := range z.Vec {
		o = msgp.AppendInt64(o, z.Vec[zjjf])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zovg uint32
	zovg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zovg > 0 {
		zovg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zpvv uint32
			zpvv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zpvv) {
				z.Vec = (z.Vec)[:zpvv]
			} else {
				z.Vec = make([]int64, zpvv)
			}
			for zjjf := range z.Vec {
				z.Vec[zjjf], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxak uint32
	zxak, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxak > 0 {
		zxak--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zgqa uint32
			zgqa, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zgqa) {
				z.Vec = (z.Vec)[:zgqa]
			} else {
				z.Vec = make([]*VTIntDbl, zgqa)
			}
			for zkgs := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zkgs] = nil
				} else {
					if z.Vec[zkgs] == nil {
						z.Vec[zkgs] = new(VTIntDbl)
					}
					var zrmn uint32
					zrmn, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zrmn > 0 {
						zrmn--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zkgs].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zkgs].Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSIntDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zkgs := range z.Vec {
		if z.Vec[zkgs] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zkgs].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zkgs].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSIntDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zkgs := range z.Vec {
		if z.Vec[zkgs] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zkgs].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zkgs].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zebz uint32
	zebz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zebz > 0 {
		zebz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zbfe uint32
			zbfe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zbfe) {
				z.Vec = (z.Vec)[:zbfe]
			} else {
				z.Vec = make([]*VTIntDbl, zbfe)
			}
			for zkgs := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zkgs] = nil
				} else {
					if z.Vec[zkgs] == nil {
						z.Vec[zkgs] = new(VTIntDbl)
					}
					var zkhx uint32
					zkhx, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zkhx > 0 {
						zkhx--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zkgs].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zkgs].Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSIntDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zkgs := range z.Vec {
		if z.Vec[zkgs] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSIntInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zfna uint32
	zfna, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zfna > 0 {
		zfna--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zvta uint32
			zvta, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zvta) {
				z.Vec = (z.Vec)[:zvta]
			} else {
				z.Vec = make([]*VTIntInt, zvta)
			}
			for zlzz := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zlzz] = nil
				} else {
					if z.Vec[zlzz] == nil {
						z.Vec[zlzz] = new(VTIntInt)
					}
					var zyyq uint32
					zyyq, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zyyq > 0 {
						zyyq--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zlzz].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zlzz].Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSIntInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zlzz := range z.Vec {
		if z.Vec[zlzz] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zlzz].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zlzz].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSIntInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zlzz := range z.Vec {
		if z.Vec[zlzz] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zlzz].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zlzz].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwzl uint32
	zwzl, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwzl > 0 {
		zwzl--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zqqf uint32
			zqqf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zqqf) {
				z.Vec = (z.Vec)[:zqqf]
			} else {
				z.Vec = make([]*VTIntInt, zqqf)
			}
			for zlzz := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zlzz] = nil
				} else {
					if z.Vec[zlzz] == nil {
						z.Vec[zlzz] = new(VTIntInt)
					}
					var zbuc uint32
					zbuc, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zbuc > 0 {
						zbuc--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zlzz].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zlzz].Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSIntInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zlzz := range z.Vec {
		if z.Vec[zlzz] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSIntStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zywr uint32
	zywr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zywr > 0 {
		zywr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zncg uint32
			zncg, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zncg) {
				z.Vec = (z.Vec)[:zncg]
			} else {
				z.Vec = make([]*VTIntStr, zncg)
			}
			for zqnj := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zqnj] = nil
				} else {
					if z.Vec[zqnj] == nil {
						z.Vec[zqnj] = new(VTIntStr)
					}
					var zllk uint32
					zllk, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zllk > 0 {
						zllk--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zqnj].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zqnj].Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSIntStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zqnj := range z.Vec {
		if z.Vec[zqnj] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zqnj].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zqnj].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSIntStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zqnj := range z.Vec {
		if z.Vec[zqnj] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zqnj].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zqnj].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdtl uint32
	zdtl, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdtl > 0 {
		zdtl--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var znyo uint32
			znyo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(znyo) {
				z.Vec = (z.Vec)[:znyo]
			} else {
				z.Vec = make([]*VTIntStr, znyo)
			}
			for zqnj := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zqnj] = nil
				} else {
					if z.Vec[zqnj] == nil {
						z.Vec[zqnj] = new(VTIntStr)
					}
					var zcsf uint32
					zcsf, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zcsf > 0 {
						zcsf--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zqnj].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zqnj].Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSIntStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zqnj := range z.Vec {
		if z.Vec[zqnj] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zqnj].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zhsv uint32
	zhsv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhsv > 0 {
		zhsv--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zvxs uint32
			zvxs, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zvxs) {
				z.Vec = (z.Vec)[:zvxs]
			} else {
				z.Vec = make([]string, zvxs)
			}
			for zlfo := range z.Vec {
				z.Vec[zlfo], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zlfo := range z.Vec {
		err = en.WriteString(z.Vec[zlfo])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zlfo := range z.Vec {
		o = msgp.AppendString(o, z.Vec[zlfo])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zazw uint32
	zazw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zazw > 0 {
		zazw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zexe uint32
			zexe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zexe) {
				z.Vec = (z.Vec)[:zexe]
			} else {
				z.Vec = make([]string, zexe)
			}
			for zlfo := range z.Vec {
				z.Vec[zlfo], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zlfo := range z.Vec {
		s += msgp.StringPrefixSize + len(z.Vec[zlfo])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqux uint32
	zqux, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqux > 0 {
		zqux--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zdaa uint32
			zdaa, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zdaa) {
				z.Vec = (z.Vec)[:zdaa]
			} else {
				z.Vec = make([]*VTStrDbl, zdaa)
			}
			for zjha := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zjha] = nil
				} else {
					if z.Vec[zjha] == nil {
						z.Vec[zjha] = new(VTStrDbl)
					}
					var zazk uint32
					zazk, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zazk > 0 {
						zazk--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zjha].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zjha].Value, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSStrDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zjha := range z.Vec {
		if z.Vec[zjha] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zjha].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zjha].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zjha := range z.Vec {
		if z.Vec[zjha] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zjha].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zjha].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrho uint32
	zrho, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrho > 0 {
		zrho--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var znxv uint32
			znxv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(znxv) {
				z.Vec = (z.Vec)[:znxv]
			} else {
				z.Vec = make([]*VTStrDbl, znxv)
			}
			for zjha := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zjha] = nil
				} else {
					if z.Vec[zjha] == nil {
						z.Vec[zjha] = new(VTStrDbl)
					}
					var zpay uint32
					zpay, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zpay > 0 {
						zpay--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zjha].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zjha].Value, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSStrDbl) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zjha := range z.Vec {
		if z.Vec[zjha] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zjha].Key) + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqcn uint32
	zqcn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqcn > 0 {
		zqcn--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zctu uint32
			zctu, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zctu) {
				z.Vec = (z.Vec)[:zctu]
			} else {
				z.Vec = make([]*VTStrInt, zctu)
			}
			for zoss := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zoss] = nil
				} else {
					if z.Vec[zoss] == nil {
						z.Vec[zoss] = new(VTStrInt)
					}
					var zgza uint32
					zgza, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zgza > 0 {
						zgza--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zoss].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zoss].Value, err = dc.ReadInt64()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSStrInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zoss := range z.Vec {
		if z.Vec[zoss] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zoss].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zoss].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zoss := range z.Vec {
		if z.Vec[zoss] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zoss].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zoss].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmct uint32
	zmct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmct > 0 {
		zmct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var zozv uint32
			zozv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zozv) {
				z.Vec = (z.Vec)[:zozv]
			} else {
				z.Vec = make([]*VTStrInt, zozv)
			}
			for zoss := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zoss] = nil
				} else {
					if z.Vec[zoss] == nil {
						z.Vec[zoss] = new(VTStrInt)
					}
					var zpyn uint32
					zpyn, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zpyn > 0 {
						zpyn--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zoss].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zoss].Value, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSStrInt) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zoss := range z.Vec {
		if z.Vec[zoss] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zoss].Key) + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zhid uint32
	zhid, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhid > 0 {
		zhid--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				err = z.Name.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "vec":
			var zxal uint32
			zxal, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zxal) {
				z.Vec = (z.Vec)[:zxal]
			} else {
				z.Vec = make([]*VTStrStr, zxal)
			}
			for zaep := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zaep] = nil
				} else {
					if z.Vec[zaep] == nil {
						z.Vec[zaep] = new(VTStrStr)
					}
					var zuaq uint32
					zuaq, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zuaq > 0 {
						zuaq--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zaep].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zaep].Value, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VSStrStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "name"
	err = en.Append(0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	if z.Name == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Name.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "vec"
	err = en.Append(0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zaep := range z.Vec {
		if z.Vec[zaep] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "key"
			err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zaep].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zaep].Value)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "name"
	o = append(o, 0x82, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if z.Name == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Name.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "vec"
	o = append(o, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zaep := range z.Vec {
		if z.Vec[zaep] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zaep].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zaep].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmbd uint32
	zmbd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmbd > 0 {
		zmbd--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Name = nil
			} else {
				if z.Name == nil {
					z.Name = new(VName)
				}
				bts, err = z.Name.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "vec":
			var ztzt uint32
			ztzt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(ztzt) {
				z.Vec = (z.Vec)[:ztzt]
			} else {
				z.Vec = make([]*VTStrStr, ztzt)
			}
			for zaep := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zaep] = nil
				} else {
					if z.Vec[zaep] == nil {
						z.Vec[zaep] = new(VTStrStr)
					}
					var zcof uint32
					zcof, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zcof > 0 {
						zcof--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zaep].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zaep].Value, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VSStrStr) Msgsize() (s int) {
	s = 1 + 5
	if z.Name == nil {
		s += msgp.NilSize
	} else {
		s += z.Name.Msgsize()
	}
	s += 4 + msgp.ArrayHeaderSize
	for zaep := range z.Vec {
		if z.Vec[zaep] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zaep].Key) + 6 + msgp.StringPrefixSize + len(z.Vec[zaep].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zfsp uint32
	zfsp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zfsp > 0 {
		zfsp--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTDblDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTDblDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendFloat64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqlx uint32
	zqlx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqlx > 0 {
		zqlx--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTDblDbl) Msgsize() (s int) {
	s = 1 + 4 + msgp.Float64Size + 6 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTDblInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zske uint32
	zske, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zske > 0 {
		zske--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTDblInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTDblInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendInt64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zfza uint32
	zfza, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zfza > 0 {
		zfza--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTDblInt) Msgsize() (s int) {
	s = 1 + 4 + msgp.Float64Size + 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTDblStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpaj uint32
	zpaj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpaj > 0 {
		zpaj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTDblStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTDblStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzkd uint32
	zzkd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzkd > 0 {
		zzkd--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTDblStr) Msgsize() (s int) {
	s = 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(z.Value)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zaue uint32
	zaue, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zaue > 0 {
		zaue--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTIntDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTIntDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendInt64(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendFloat64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdhi uint32
	zdhi, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdhi > 0 {
		zdhi--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTIntDbl) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 6 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTIntInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztgh uint32
	ztgh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztgh > 0 {
		ztgh--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTIntInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTIntInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendInt64(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendInt64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zvbr uint32
	zvbr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zvbr > 0 {
		zvbr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTIntInt) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTIntStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqpq uint32
	zqpq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqpq > 0 {
		zqpq--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTIntStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTIntStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendInt64(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwar uint32
	zwar, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwar > 0 {
		zwar--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTIntStr) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Value)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpxp uint32
	zpxp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpxp > 0 {
		zpxp--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTStrDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendFloat64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztpv uint32
	ztpv, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztpv > 0 {
		ztpv--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTStrDbl) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 6 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgnp uint32
	zgnp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgnp > 0 {
		zgnp--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTStrInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendInt64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdge uint32
	zdge, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdge > 0 {
		zdge--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTStrInt) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrwv uint32
	zrwv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrwv > 0 {
		zrwv--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z VTStrStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "key"
	err = en.Append(0x82, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z VTStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "key"
	o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VTStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhgc uint32
	zhgc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhgc > 0 {
		zhgc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z VTStrStr) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 6 + msgp.StringPrefixSize + len(z.Value)
	return
}
