package vepr

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Tag) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
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
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
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
	var zcmr uint32
	zcmr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
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
			var zajw uint32
			zajw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zajw) {
				z.Vec = (z.Vec)[:zajw]
			} else {
				z.Vec = make([]float64, zajw)
			}
			for zbai := range z.Vec {
				z.Vec[zbai], err = dc.ReadFloat64()
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
	for zbai := range z.Vec {
		err = en.WriteFloat64(z.Vec[zbai])
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
	for zbai := range z.Vec {
		o = msgp.AppendFloat64(o, z.Vec[zbai])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
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
			var zhct uint32
			zhct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zhct) {
				z.Vec = (z.Vec)[:zhct]
			} else {
				z.Vec = make([]float64, zhct)
			}
			for zbai := range z.Vec {
				z.Vec[zbai], bts, err = msgp.ReadFloat64Bytes(bts)
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
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
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
			var zlqf uint32
			zlqf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zlqf) {
				z.Vec = (z.Vec)[:zlqf]
			} else {
				z.Vec = make([]*VTDblDbl, zlqf)
			}
			for zcua := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zcua] = nil
				} else {
					if z.Vec[zcua] == nil {
						z.Vec[zcua] = new(VTDblDbl)
					}
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
						case "key":
							z.Vec[zcua].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zcua].Value, err = dc.ReadFloat64()
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
	for zcua := range z.Vec {
		if z.Vec[zcua] == nil {
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
			err = en.WriteFloat64(z.Vec[zcua].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zcua].Value)
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
	for zcua := range z.Vec {
		if z.Vec[zcua] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zcua].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zcua].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
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
			var zjfb uint32
			zjfb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zjfb) {
				z.Vec = (z.Vec)[:zjfb]
			} else {
				z.Vec = make([]*VTDblDbl, zjfb)
			}
			for zcua := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zcua] = nil
				} else {
					if z.Vec[zcua] == nil {
						z.Vec[zcua] = new(VTDblDbl)
					}
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
						case "key":
							z.Vec[zcua].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zcua].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	for zcua := range z.Vec {
		if z.Vec[zcua] == nil {
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
	var zrsw uint32
	zrsw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
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
			var zxpk uint32
			zxpk, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zxpk) {
				z.Vec = (z.Vec)[:zxpk]
			} else {
				z.Vec = make([]*VTDblInt, zxpk)
			}
			for zeff := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zeff] = nil
				} else {
					if z.Vec[zeff] == nil {
						z.Vec[zeff] = new(VTDblInt)
					}
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
						case "key":
							z.Vec[zeff].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zeff].Value, err = dc.ReadInt64()
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
	for zeff := range z.Vec {
		if z.Vec[zeff] == nil {
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
			err = en.WriteFloat64(z.Vec[zeff].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zeff].Value)
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
	for zeff := range z.Vec {
		if z.Vec[zeff] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zeff].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zeff].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zobc uint32
	zobc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zobc > 0 {
		zobc--
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
			var zsnv uint32
			zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zsnv) {
				z.Vec = (z.Vec)[:zsnv]
			} else {
				z.Vec = make([]*VTDblInt, zsnv)
			}
			for zeff := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zeff] = nil
				} else {
					if z.Vec[zeff] == nil {
						z.Vec[zeff] = new(VTDblInt)
					}
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
						case "key":
							z.Vec[zeff].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zeff].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	for zeff := range z.Vec {
		if z.Vec[zeff] == nil {
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
	var zpez uint32
	zpez, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpez > 0 {
		zpez--
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
			var zqke uint32
			zqke, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zqke) {
				z.Vec = (z.Vec)[:zqke]
			} else {
				z.Vec = make([]*VTDblStr, zqke)
			}
			for zema := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zema] = nil
				} else {
					if z.Vec[zema] == nil {
						z.Vec[zema] = new(VTDblStr)
					}
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
						case "key":
							z.Vec[zema].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zema].Value, err = dc.ReadString()
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
	for zema := range z.Vec {
		if z.Vec[zema] == nil {
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
			err = en.WriteFloat64(z.Vec[zema].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zema].Value)
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
	for zema := range z.Vec {
		if z.Vec[zema] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zema].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zema].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zyzr uint32
	zyzr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
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
			var zywj uint32
			zywj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zywj) {
				z.Vec = (z.Vec)[:zywj]
			} else {
				z.Vec = make([]*VTDblStr, zywj)
			}
			for zema := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zema] = nil
				} else {
					if z.Vec[zema] == nil {
						z.Vec[zema] = new(VTDblStr)
					}
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
						case "key":
							z.Vec[zema].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zema].Value, bts, err = msgp.ReadStringBytes(bts)
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
	for zema := range z.Vec {
		if z.Vec[zema] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zema].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrfe uint32
	zrfe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrfe > 0 {
		zrfe--
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
			var zgmo uint32
			zgmo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zgmo) {
				z.Vec = (z.Vec)[:zgmo]
			} else {
				z.Vec = make([]int64, zgmo)
			}
			for zzpf := range z.Vec {
				z.Vec[zzpf], err = dc.ReadInt64()
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
	for zzpf := range z.Vec {
		err = en.WriteInt64(z.Vec[zzpf])
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
	for zzpf := range z.Vec {
		o = msgp.AppendInt64(o, z.Vec[zzpf])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztaf uint32
	ztaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztaf > 0 {
		ztaf--
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
			var zeth uint32
			zeth, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zeth) {
				z.Vec = (z.Vec)[:zeth]
			} else {
				z.Vec = make([]int64, zeth)
			}
			for zzpf := range z.Vec {
				z.Vec[zzpf], bts, err = msgp.ReadInt64Bytes(bts)
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
	var zrjx uint32
	zrjx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrjx > 0 {
		zrjx--
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
			var zawn uint32
			zawn, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zawn) {
				z.Vec = (z.Vec)[:zawn]
			} else {
				z.Vec = make([]*VTIntDbl, zawn)
			}
			for zsbz := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zsbz] = nil
				} else {
					if z.Vec[zsbz] == nil {
						z.Vec[zsbz] = new(VTIntDbl)
					}
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
						case "key":
							z.Vec[zsbz].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zsbz].Value, err = dc.ReadFloat64()
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
	for zsbz := range z.Vec {
		if z.Vec[zsbz] == nil {
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
			err = en.WriteInt64(z.Vec[zsbz].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zsbz].Value)
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
	for zsbz := range z.Vec {
		if z.Vec[zsbz] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zsbz].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zsbz].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrbe uint32
	zrbe, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
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
			var zmfd uint32
			zmfd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zmfd) {
				z.Vec = (z.Vec)[:zmfd]
			} else {
				z.Vec = make([]*VTIntDbl, zmfd)
			}
			for zsbz := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zsbz] = nil
				} else {
					if z.Vec[zsbz] == nil {
						z.Vec[zsbz] = new(VTIntDbl)
					}
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
						case "key":
							z.Vec[zsbz].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zsbz].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	for zsbz := range z.Vec {
		if z.Vec[zsbz] == nil {
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
	var zbal uint32
	zbal, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbal > 0 {
		zbal--
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
			var zjqz uint32
			zjqz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zjqz) {
				z.Vec = (z.Vec)[:zjqz]
			} else {
				z.Vec = make([]*VTIntInt, zjqz)
			}
			for zelx := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zelx] = nil
				} else {
					if z.Vec[zelx] == nil {
						z.Vec[zelx] = new(VTIntInt)
					}
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
						case "key":
							z.Vec[zelx].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zelx].Value, err = dc.ReadInt64()
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
	for zelx := range z.Vec {
		if z.Vec[zelx] == nil {
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
			err = en.WriteInt64(z.Vec[zelx].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zelx].Value)
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
	for zelx := range z.Vec {
		if z.Vec[zelx] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zelx].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zelx].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztmt uint32
	ztmt, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztmt > 0 {
		ztmt--
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
			var ztco uint32
			ztco, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(ztco) {
				z.Vec = (z.Vec)[:ztco]
			} else {
				z.Vec = make([]*VTIntInt, ztco)
			}
			for zelx := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zelx] = nil
				} else {
					if z.Vec[zelx] == nil {
						z.Vec[zelx] = new(VTIntInt)
					}
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
						case "key":
							z.Vec[zelx].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zelx].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	for zelx := range z.Vec {
		if z.Vec[zelx] == nil {
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
	var zinl uint32
	zinl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zinl > 0 {
		zinl--
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
			var zare uint32
			zare, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zare) {
				z.Vec = (z.Vec)[:zare]
			} else {
				z.Vec = make([]*VTIntStr, zare)
			}
			for ztyy := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[ztyy] = nil
				} else {
					if z.Vec[ztyy] == nil {
						z.Vec[ztyy] = new(VTIntStr)
					}
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
						case "key":
							z.Vec[ztyy].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[ztyy].Value, err = dc.ReadString()
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
	for ztyy := range z.Vec {
		if z.Vec[ztyy] == nil {
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
			err = en.WriteInt64(z.Vec[ztyy].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[ztyy].Value)
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
	for ztyy := range z.Vec {
		if z.Vec[ztyy] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[ztyy].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[ztyy].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zixj uint32
	zixj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zixj > 0 {
		zixj--
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
			var zrsc uint32
			zrsc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zrsc) {
				z.Vec = (z.Vec)[:zrsc]
			} else {
				z.Vec = make([]*VTIntStr, zrsc)
			}
			for ztyy := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[ztyy] = nil
				} else {
					if z.Vec[ztyy] == nil {
						z.Vec[ztyy] = new(VTIntStr)
					}
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
						case "key":
							z.Vec[ztyy].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[ztyy].Value, bts, err = msgp.ReadStringBytes(bts)
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
	for ztyy := range z.Vec {
		if z.Vec[ztyy] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Vec[ztyy].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var znsg uint32
	znsg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for znsg > 0 {
		znsg--
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
			var zrus uint32
			zrus, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zrus) {
				z.Vec = (z.Vec)[:zrus]
			} else {
				z.Vec = make([]string, zrus)
			}
			for zswy := range z.Vec {
				z.Vec[zswy], err = dc.ReadString()
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
	for zswy := range z.Vec {
		err = en.WriteString(z.Vec[zswy])
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
	for zswy := range z.Vec {
		o = msgp.AppendString(o, z.Vec[zswy])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsvm uint32
	zsvm, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsvm > 0 {
		zsvm--
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
			var zaoz uint32
			zaoz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zaoz) {
				z.Vec = (z.Vec)[:zaoz]
			} else {
				z.Vec = make([]string, zaoz)
			}
			for zswy := range z.Vec {
				z.Vec[zswy], bts, err = msgp.ReadStringBytes(bts)
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
	for zswy := range z.Vec {
		s += msgp.StringPrefixSize + len(z.Vec[zswy])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsbo uint32
	zsbo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsbo > 0 {
		zsbo--
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
			var zjif uint32
			zjif, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zjif) {
				z.Vec = (z.Vec)[:zjif]
			} else {
				z.Vec = make([]*VTStrDbl, zjif)
			}
			for zfzb := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zfzb] = nil
				} else {
					if z.Vec[zfzb] == nil {
						z.Vec[zfzb] = new(VTStrDbl)
					}
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
						case "key":
							z.Vec[zfzb].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zfzb].Value, err = dc.ReadFloat64()
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
	for zfzb := range z.Vec {
		if z.Vec[zfzb] == nil {
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
			err = en.WriteString(z.Vec[zfzb].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zfzb].Value)
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
	for zfzb := range z.Vec {
		if z.Vec[zfzb] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zfzb].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zfzb].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsnw uint32
	zsnw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsnw > 0 {
		zsnw--
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
			var ztls uint32
			ztls, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(ztls) {
				z.Vec = (z.Vec)[:ztls]
			} else {
				z.Vec = make([]*VTStrDbl, ztls)
			}
			for zfzb := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zfzb] = nil
				} else {
					if z.Vec[zfzb] == nil {
						z.Vec[zfzb] = new(VTStrDbl)
					}
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
						case "key":
							z.Vec[zfzb].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zfzb].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	for zfzb := range z.Vec {
		if z.Vec[zfzb] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zfzb].Key) + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zopb uint32
	zopb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zopb > 0 {
		zopb--
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
			var zuop uint32
			zuop, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zuop) {
				z.Vec = (z.Vec)[:zuop]
			} else {
				z.Vec = make([]*VTStrInt, zuop)
			}
			for zigk := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zigk] = nil
				} else {
					if z.Vec[zigk] == nil {
						z.Vec[zigk] = new(VTStrInt)
					}
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
						case "key":
							z.Vec[zigk].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zigk].Value, err = dc.ReadInt64()
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
	for zigk := range z.Vec {
		if z.Vec[zigk] == nil {
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
			err = en.WriteString(z.Vec[zigk].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zigk].Value)
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
	for zigk := range z.Vec {
		if z.Vec[zigk] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zigk].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zigk].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zupd uint32
	zupd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zupd > 0 {
		zupd--
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
			var zome uint32
			zome, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zome) {
				z.Vec = (z.Vec)[:zome]
			} else {
				z.Vec = make([]*VTStrInt, zome)
			}
			for zigk := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zigk] = nil
				} else {
					if z.Vec[zigk] == nil {
						z.Vec[zigk] = new(VTStrInt)
					}
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
						case "key":
							z.Vec[zigk].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zigk].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	for zigk := range z.Vec {
		if z.Vec[zigk] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zigk].Key) + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VLStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zknt uint32
	zknt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zknt > 0 {
		zknt--
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
			var zxye uint32
			zxye, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zxye) {
				z.Vec = (z.Vec)[:zxye]
			} else {
				z.Vec = make([]*VTStrStr, zxye)
			}
			for zarz := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zarz] = nil
				} else {
					if z.Vec[zarz] == nil {
						z.Vec[zarz] = new(VTStrStr)
					}
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
						case "key":
							z.Vec[zarz].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zarz].Value, err = dc.ReadString()
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
	for zarz := range z.Vec {
		if z.Vec[zarz] == nil {
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
			err = en.WriteString(z.Vec[zarz].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zarz].Value)
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
	for zarz := range z.Vec {
		if z.Vec[zarz] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zarz].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zarz].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VLStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zlsx uint32
	zlsx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zlsx > 0 {
		zlsx--
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
			var zbgy uint32
			zbgy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zbgy) {
				z.Vec = (z.Vec)[:zbgy]
			} else {
				z.Vec = make([]*VTStrStr, zbgy)
			}
			for zarz := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zarz] = nil
				} else {
					if z.Vec[zarz] == nil {
						z.Vec[zarz] = new(VTStrStr)
					}
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
						case "key":
							z.Vec[zarz].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zarz].Value, bts, err = msgp.ReadStringBytes(bts)
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
	for zarz := range z.Vec {
		if z.Vec[zarz] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zarz].Key) + 6 + msgp.StringPrefixSize + len(z.Vec[zarz].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zmbt uint32
	zmbt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zmbt > 0 {
		zmbt--
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
	var zfvi uint32
	zfvi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zfvi > 0 {
		zfvi--
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
			var zzrg uint32
			zzrg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zzrg > 0 {
				z.Vec = make(map[string]float64, zzrg)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zzrg > 0 {
				zzrg--
				var zlur string
				var zupi float64
				zlur, err = dc.ReadString()
				if err != nil {
					return
				}
				zupi, err = dc.ReadFloat64()
				if err != nil {
					return
				}
				z.Vec[zlur] = zupi
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
	for zlur, zupi := range z.Vec {
		err = en.WriteString(zlur)
		if err != nil {
			return
		}
		err = en.WriteFloat64(zupi)
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
	for zlur, zupi := range z.Vec {
		o = msgp.AppendString(o, zlur)
		o = msgp.AppendFloat64(o, zupi)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbmy uint32
	zbmy, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbmy > 0 {
		zbmy--
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
			var zarl uint32
			zarl, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zarl > 0 {
				z.Vec = make(map[string]float64, zarl)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zarl > 0 {
				var zlur string
				var zupi float64
				zarl--
				zlur, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zupi, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
				z.Vec[zlur] = zupi
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
		for zlur, zupi := range z.Vec {
			_ = zupi
			s += msgp.StringPrefixSize + len(zlur) + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zziv uint32
	zziv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zziv > 0 {
		zziv--
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
			var zabj uint32
			zabj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zabj > 0 {
				z.Vec = make(map[string]int64, zabj)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zabj > 0 {
				zabj--
				var zctz string
				var zljl int64
				zctz, err = dc.ReadString()
				if err != nil {
					return
				}
				zljl, err = dc.ReadInt64()
				if err != nil {
					return
				}
				z.Vec[zctz] = zljl
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
	for zctz, zljl := range z.Vec {
		err = en.WriteString(zctz)
		if err != nil {
			return
		}
		err = en.WriteInt64(zljl)
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
	for zctz, zljl := range z.Vec {
		o = msgp.AppendString(o, zctz)
		o = msgp.AppendInt64(o, zljl)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmlx uint32
	zmlx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmlx > 0 {
		zmlx--
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
			var zvbw uint32
			zvbw, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zvbw > 0 {
				z.Vec = make(map[string]int64, zvbw)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zvbw > 0 {
				var zctz string
				var zljl int64
				zvbw--
				zctz, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zljl, bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
				z.Vec[zctz] = zljl
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
		for zctz, zljl := range z.Vec {
			_ = zljl
			s += msgp.StringPrefixSize + len(zctz) + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zexy uint32
	zexy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zexy > 0 {
		zexy--
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
			var zakb uint32
			zakb, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zakb > 0 {
				z.Vec = make(map[string]string, zakb)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zakb > 0 {
				zakb--
				var zgvb string
				var zqzg string
				zgvb, err = dc.ReadString()
				if err != nil {
					return
				}
				zqzg, err = dc.ReadString()
				if err != nil {
					return
				}
				z.Vec[zgvb] = zqzg
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
	for zgvb, zqzg := range z.Vec {
		err = en.WriteString(zgvb)
		if err != nil {
			return
		}
		err = en.WriteString(zqzg)
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
	for zgvb, zqzg := range z.Vec {
		o = msgp.AppendString(o, zgvb)
		o = msgp.AppendString(o, zqzg)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsdj uint32
	zsdj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsdj > 0 {
		zsdj--
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
			var zsgp uint32
			zsgp, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zsgp > 0 {
				z.Vec = make(map[string]string, zsgp)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zsgp > 0 {
				var zgvb string
				var zqzg string
				zsgp--
				zgvb, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zqzg, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Vec[zgvb] = zqzg
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
		for zgvb, zqzg := range z.Vec {
			_ = zqzg
			s += msgp.StringPrefixSize + len(zgvb) + msgp.StringPrefixSize + len(zqzg)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdif uint32
	zdif, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdif > 0 {
		zdif--
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
			var zibu uint32
			zibu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zibu > 0 {
				z.Vec = make(map[string]*VTDblDbl, zibu)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zibu > 0 {
				zibu--
				var zngc string
				var zwfl *VTDblDbl
				zngc, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zwfl = nil
				} else {
					if zwfl == nil {
						zwfl = new(VTDblDbl)
					}
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
						case "key":
							zwfl.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							zwfl.Value, err = dc.ReadFloat64()
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
				z.Vec[zngc] = zwfl
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
	for zngc, zwfl := range z.Vec {
		err = en.WriteString(zngc)
		if err != nil {
			return
		}
		if zwfl == nil {
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
			err = en.WriteFloat64(zwfl.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zwfl.Value)
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
	for zngc, zwfl := range z.Vec {
		o = msgp.AppendString(o, zngc)
		if zwfl == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, zwfl.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zwfl.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmow uint32
	zmow, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmow > 0 {
		zmow--
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
			var zdit uint32
			zdit, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zdit > 0 {
				z.Vec = make(map[string]*VTDblDbl, zdit)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zdit > 0 {
				var zngc string
				var zwfl *VTDblDbl
				zdit--
				zngc, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zwfl = nil
				} else {
					if zwfl == nil {
						zwfl = new(VTDblDbl)
					}
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
						case "key":
							zwfl.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zwfl.Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
				z.Vec[zngc] = zwfl
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
		for zngc, zwfl := range z.Vec {
			_ = zwfl
			s += msgp.StringPrefixSize + len(zngc)
			if zwfl == nil {
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
	var ztic uint32
	ztic, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztic > 0 {
		ztic--
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
			var ztoj uint32
			ztoj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && ztoj > 0 {
				z.Vec = make(map[string]*VTDblInt, ztoj)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for ztoj > 0 {
				ztoj--
				var zoqj string
				var zmqr *VTDblInt
				zoqj, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zmqr = nil
				} else {
					if zmqr == nil {
						zmqr = new(VTDblInt)
					}
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
						case "key":
							zmqr.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							zmqr.Value, err = dc.ReadInt64()
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
				z.Vec[zoqj] = zmqr
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
	for zoqj, zmqr := range z.Vec {
		err = en.WriteString(zoqj)
		if err != nil {
			return
		}
		if zmqr == nil {
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
			err = en.WriteFloat64(zmqr.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zmqr.Value)
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
	for zoqj, zmqr := range z.Vec {
		o = msgp.AppendString(o, zoqj)
		if zmqr == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, zmqr.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zmqr.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zyes uint32
	zyes, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyes > 0 {
		zyes--
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
			var zxzy uint32
			zxzy, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zxzy > 0 {
				z.Vec = make(map[string]*VTDblInt, zxzy)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zxzy > 0 {
				var zoqj string
				var zmqr *VTDblInt
				zxzy--
				zoqj, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zmqr = nil
				} else {
					if zmqr == nil {
						zmqr = new(VTDblInt)
					}
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
						case "key":
							zmqr.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zmqr.Value, bts, err = msgp.ReadInt64Bytes(bts)
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
				z.Vec[zoqj] = zmqr
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
		for zoqj, zmqr := range z.Vec {
			_ = zmqr
			s += msgp.StringPrefixSize + len(zoqj)
			if zmqr == nil {
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
	var zdrz uint32
	zdrz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdrz > 0 {
		zdrz--
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
			var znpn uint32
			znpn, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && znpn > 0 {
				z.Vec = make(map[string]*VTDblStr, znpn)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for znpn > 0 {
				znpn--
				var zrod string
				var zmbn *VTDblStr
				zrod, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zmbn = nil
				} else {
					if zmbn == nil {
						zmbn = new(VTDblStr)
					}
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
						case "key":
							zmbn.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							zmbn.Value, err = dc.ReadString()
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
				z.Vec[zrod] = zmbn
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
	for zrod, zmbn := range z.Vec {
		err = en.WriteString(zrod)
		if err != nil {
			return
		}
		if zmbn == nil {
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
			err = en.WriteFloat64(zmbn.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(zmbn.Value)
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
	for zrod, zmbn := range z.Vec {
		o = msgp.AppendString(o, zrod)
		if zmbn == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, zmbn.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, zmbn.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjpm uint32
	zjpm, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjpm > 0 {
		zjpm--
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
			var zhdt uint32
			zhdt, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zhdt > 0 {
				z.Vec = make(map[string]*VTDblStr, zhdt)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zhdt > 0 {
				var zrod string
				var zmbn *VTDblStr
				zhdt--
				zrod, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zmbn = nil
				} else {
					if zmbn == nil {
						zmbn = new(VTDblStr)
					}
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
						case "key":
							zmbn.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zmbn.Value, bts, err = msgp.ReadStringBytes(bts)
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
				z.Vec[zrod] = zmbn
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
		for zrod, zmbn := range z.Vec {
			_ = zmbn
			s += msgp.StringPrefixSize + len(zrod)
			if zmbn == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(zmbn.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zfum uint32
	zfum, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zfum > 0 {
		zfum--
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
			var zaps uint32
			zaps, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zaps > 0 {
				z.Vec = make(map[string]*VTIntDbl, zaps)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zaps > 0 {
				zaps--
				var zayo string
				var zrsu *VTIntDbl
				zayo, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zrsu = nil
				} else {
					if zrsu == nil {
						zrsu = new(VTIntDbl)
					}
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
						case "key":
							zrsu.Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							zrsu.Value, err = dc.ReadFloat64()
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
				z.Vec[zayo] = zrsu
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
	for zayo, zrsu := range z.Vec {
		err = en.WriteString(zayo)
		if err != nil {
			return
		}
		if zrsu == nil {
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
			err = en.WriteInt64(zrsu.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zrsu.Value)
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
	for zayo, zrsu := range z.Vec {
		o = msgp.AppendString(o, zayo)
		if zrsu == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zrsu.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zrsu.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhbk uint32
	zhbk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhbk > 0 {
		zhbk--
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
			var zmyy uint32
			zmyy, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zmyy > 0 {
				z.Vec = make(map[string]*VTIntDbl, zmyy)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmyy > 0 {
				var zayo string
				var zrsu *VTIntDbl
				zmyy--
				zayo, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zrsu = nil
				} else {
					if zrsu == nil {
						zrsu = new(VTIntDbl)
					}
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
						case "key":
							zrsu.Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zrsu.Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
				z.Vec[zayo] = zrsu
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
		for zayo, zrsu := range z.Vec {
			_ = zrsu
			s += msgp.StringPrefixSize + len(zayo)
			if zrsu == nil {
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
	var zbgg uint32
	zbgg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbgg > 0 {
		zbgg--
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
			var zcnq uint32
			zcnq, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zcnq > 0 {
				z.Vec = make(map[string]*VTIntInt, zcnq)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zcnq > 0 {
				zcnq--
				var zvgw string
				var zffb *VTIntInt
				zvgw, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zffb = nil
				} else {
					if zffb == nil {
						zffb = new(VTIntInt)
					}
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
						case "key":
							zffb.Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							zffb.Value, err = dc.ReadInt64()
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
				z.Vec[zvgw] = zffb
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
	for zvgw, zffb := range z.Vec {
		err = en.WriteString(zvgw)
		if err != nil {
			return
		}
		if zffb == nil {
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
			err = en.WriteInt64(zffb.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zffb.Value)
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
	for zvgw, zffb := range z.Vec {
		o = msgp.AppendString(o, zvgw)
		if zffb == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zffb.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zffb.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zreu uint32
	zreu, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zreu > 0 {
		zreu--
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
			var znuz uint32
			znuz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && znuz > 0 {
				z.Vec = make(map[string]*VTIntInt, znuz)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for znuz > 0 {
				var zvgw string
				var zffb *VTIntInt
				znuz--
				zvgw, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zffb = nil
				} else {
					if zffb == nil {
						zffb = new(VTIntInt)
					}
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
						case "key":
							zffb.Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zffb.Value, bts, err = msgp.ReadInt64Bytes(bts)
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
				z.Vec[zvgw] = zffb
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
		for zvgw, zffb := range z.Vec {
			_ = zffb
			s += msgp.StringPrefixSize + len(zvgw)
			if zffb == nil {
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
	var zkut uint32
	zkut, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zkut > 0 {
		zkut--
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
			var zmyg uint32
			zmyg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zmyg > 0 {
				z.Vec = make(map[string]*VTIntStr, zmyg)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmyg > 0 {
				zmyg--
				var zmzo string
				var ztar *VTIntStr
				zmzo, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					ztar = nil
				} else {
					if ztar == nil {
						ztar = new(VTIntStr)
					}
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
						case "key":
							ztar.Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							ztar.Value, err = dc.ReadString()
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
				z.Vec[zmzo] = ztar
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
	for zmzo, ztar := range z.Vec {
		err = en.WriteString(zmzo)
		if err != nil {
			return
		}
		if ztar == nil {
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
			err = en.WriteInt64(ztar.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(ztar.Value)
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
	for zmzo, ztar := range z.Vec {
		o = msgp.AppendString(o, zmzo)
		if ztar == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, ztar.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, ztar.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zyba uint32
	zyba, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyba > 0 {
		zyba--
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
			var zddv uint32
			zddv, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zddv > 0 {
				z.Vec = make(map[string]*VTIntStr, zddv)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zddv > 0 {
				var zmzo string
				var ztar *VTIntStr
				zddv--
				zmzo, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					ztar = nil
				} else {
					if ztar == nil {
						ztar = new(VTIntStr)
					}
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
						case "key":
							ztar.Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							ztar.Value, bts, err = msgp.ReadStringBytes(bts)
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
				z.Vec[zmzo] = ztar
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
		for zmzo, ztar := range z.Vec {
			_ = ztar
			s += msgp.StringPrefixSize + len(zmzo)
			if ztar == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(ztar.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zclm uint32
	zclm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zclm > 0 {
		zclm--
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
			var zxiu uint32
			zxiu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zxiu > 0 {
				z.Vec = make(map[string]*VTStrDbl, zxiu)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zxiu > 0 {
				zxiu--
				var zfsf string
				var zgpy *VTStrDbl
				zfsf, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zgpy = nil
				} else {
					if zgpy == nil {
						zgpy = new(VTStrDbl)
					}
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
						case "key":
							zgpy.Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							zgpy.Value, err = dc.ReadFloat64()
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
				z.Vec[zfsf] = zgpy
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
	for zfsf, zgpy := range z.Vec {
		err = en.WriteString(zfsf)
		if err != nil {
			return
		}
		if zgpy == nil {
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
			err = en.WriteString(zgpy.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zgpy.Value)
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
	for zfsf, zgpy := range z.Vec {
		o = msgp.AppendString(o, zfsf)
		if zgpy == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zgpy.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zgpy.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztnd uint32
	ztnd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztnd > 0 {
		ztnd--
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
			var ztja uint32
			ztja, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && ztja > 0 {
				z.Vec = make(map[string]*VTStrDbl, ztja)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for ztja > 0 {
				var zfsf string
				var zgpy *VTStrDbl
				ztja--
				zfsf, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zgpy = nil
				} else {
					if zgpy == nil {
						zgpy = new(VTStrDbl)
					}
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
						case "key":
							zgpy.Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							zgpy.Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
				z.Vec[zfsf] = zgpy
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
		for zfsf, zgpy := range z.Vec {
			_ = zgpy
			s += msgp.StringPrefixSize + len(zfsf)
			if zgpy == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zgpy.Key) + 6 + msgp.Float64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zkpr uint32
	zkpr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zkpr > 0 {
		zkpr--
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
			var zucj uint32
			zucj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zucj > 0 {
				z.Vec = make(map[string]*VTStrInt, zucj)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zucj > 0 {
				zucj--
				var zded string
				var zmiy *VTStrInt
				zded, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zmiy = nil
				} else {
					if zmiy == nil {
						zmiy = new(VTStrInt)
					}
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
						case "key":
							zmiy.Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							zmiy.Value, err = dc.ReadInt64()
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
				z.Vec[zded] = zmiy
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
	for zded, zmiy := range z.Vec {
		err = en.WriteString(zded)
		if err != nil {
			return
		}
		if zmiy == nil {
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
			err = en.WriteString(zmiy.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zmiy.Value)
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
	for zded, zmiy := range z.Vec {
		o = msgp.AppendString(o, zded)
		if zmiy == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zmiy.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zmiy.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Vec":
			var zevw uint32
			zevw, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zevw > 0 {
				z.Vec = make(map[string]*VTStrInt, zevw)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zevw > 0 {
				var zded string
				var zmiy *VTStrInt
				zevw--
				zded, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zmiy = nil
				} else {
					if zmiy == nil {
						zmiy = new(VTStrInt)
					}
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
						case "key":
							zmiy.Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							zmiy.Value, bts, err = msgp.ReadInt64Bytes(bts)
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
				z.Vec[zded] = zmiy
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
		for zded, zmiy := range z.Vec {
			_ = zmiy
			s += msgp.StringPrefixSize + len(zded)
			if zmiy == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zmiy.Key) + 6 + msgp.Int64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxqw uint32
	zxqw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxqw > 0 {
		zxqw--
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
			var zpzw uint32
			zpzw, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zpzw > 0 {
				z.Vec = make(map[string]*VTStrStr, zpzw)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zpzw > 0 {
				zpzw--
				var zmkf string
				var zrdg *VTStrStr
				zmkf, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zrdg = nil
				} else {
					if zrdg == nil {
						zrdg = new(VTStrStr)
					}
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
						case "key":
							zrdg.Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							zrdg.Value, err = dc.ReadString()
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
				z.Vec[zmkf] = zrdg
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
	for zmkf, zrdg := range z.Vec {
		err = en.WriteString(zmkf)
		if err != nil {
			return
		}
		if zrdg == nil {
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
			err = en.WriteString(zrdg.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(zrdg.Value)
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
	for zmkf, zrdg := range z.Vec {
		o = msgp.AppendString(o, zmkf)
		if zrdg == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zrdg.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, zrdg.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmzc uint32
	zmzc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmzc > 0 {
		zmzc--
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
			var zkvc uint32
			zkvc, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zkvc > 0 {
				z.Vec = make(map[string]*VTStrStr, zkvc)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zkvc > 0 {
				var zmkf string
				var zrdg *VTStrStr
				zkvc--
				zmkf, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zrdg = nil
				} else {
					if zrdg == nil {
						zrdg = new(VTStrStr)
					}
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
						case "key":
							zrdg.Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							zrdg.Value, bts, err = msgp.ReadStringBytes(bts)
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
				z.Vec[zmkf] = zrdg
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
		for zmkf, zrdg := range z.Vec {
			_ = zrdg
			s += msgp.StringPrefixSize + len(zmkf)
			if zrdg == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zrdg.Key) + 6 + msgp.StringPrefixSize + len(zrdg.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VName) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsye uint32
	zsye, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsye > 0 {
		zsye--
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
			err = z.Resolution.DecodeMsg(dc)
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
			var zofl uint32
			zofl, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(zofl) {
				z.Tags = (z.Tags)[:zofl]
			} else {
				z.Tags = make([]*Tag, zofl)
			}
			for zjkd := range z.Tags {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Tags[zjkd] = nil
				} else {
					if z.Tags[zjkd] == nil {
						z.Tags[zjkd] = new(Tag)
					}
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
						case "Name":
							z.Tags[zjkd].Name, err = dc.ReadString()
							if err != nil {
								return
							}
						case "Value":
							z.Tags[zjkd].Value, err = dc.ReadString()
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
	err = z.Resolution.EncodeMsg(en)
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
	for zjkd := range z.Tags {
		if z.Tags[zjkd] == nil {
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
			err = en.WriteString(z.Tags[zjkd].Name)
			if err != nil {
				return
			}
			// write "Value"
			err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Tags[zjkd].Value)
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
	o, err = z.Resolution.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "ttl"
	o = append(o, 0xa3, 0x74, 0x74, 0x6c)
	o = msgp.AppendUint32(o, z.Ttl)
	// string "uid"
	o = append(o, 0xa3, 0x75, 0x69, 0x64)
	o = msgp.AppendString(o, z.XUniqueStr)
	// string "tags"
	o = append(o, 0xa4, 0x74, 0x61, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Tags)))
	for zjkd := range z.Tags {
		if z.Tags[zjkd] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "Name"
			o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Tags[zjkd].Name)
			// string "Value"
			o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Tags[zjkd].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VName) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zeql uint32
	zeql, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zeql > 0 {
		zeql--
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
			bts, err = z.Resolution.UnmarshalMsg(bts)
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
			var zlhq uint32
			zlhq, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(zlhq) {
				z.Tags = (z.Tags)[:zlhq]
			} else {
				z.Tags = make([]*Tag, zlhq)
			}
			for zjkd := range z.Tags {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Tags[zjkd] = nil
				} else {
					if z.Tags[zjkd] == nil {
						z.Tags[zjkd] = new(Tag)
					}
					var zgbp uint32
					zgbp, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zgbp > 0 {
						zgbp--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Name":
							z.Tags[zjkd].Name, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "Value":
							z.Tags[zjkd].Value, bts, err = msgp.ReadStringBytes(bts)
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
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 4 + z.Resolution.Msgsize() + 4 + msgp.Uint32Size + 4 + msgp.StringPrefixSize + len(z.XUniqueStr) + 5 + msgp.ArrayHeaderSize
	for zjkd := range z.Tags {
		if z.Tags[zjkd] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 5 + msgp.StringPrefixSize + len(z.Tags[zjkd].Name) + 6 + msgp.StringPrefixSize + len(z.Tags[zjkd].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbgh uint32
	zbgh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbgh > 0 {
		zbgh--
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
			var zmku uint32
			zmku, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zmku) {
				z.Vec = (z.Vec)[:zmku]
			} else {
				z.Vec = make([]float64, zmku)
			}
			for znbp := range z.Vec {
				z.Vec[znbp], err = dc.ReadFloat64()
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
	for znbp := range z.Vec {
		err = en.WriteFloat64(z.Vec[znbp])
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
	for znbp := range z.Vec {
		o = msgp.AppendFloat64(o, z.Vec[znbp])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
			var zumm uint32
			zumm, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zumm) {
				z.Vec = (z.Vec)[:zumm]
			} else {
				z.Vec = make([]float64, zumm)
			}
			for znbp := range z.Vec {
				z.Vec[znbp], bts, err = msgp.ReadFloat64Bytes(bts)
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
	var zcsj uint32
	zcsj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcsj > 0 {
		zcsj--
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
			var zmjj uint32
			zmjj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zmjj) {
				z.Vec = (z.Vec)[:zmjj]
			} else {
				z.Vec = make([]*VTDblDbl, zmjj)
			}
			for zpbh := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zpbh] = nil
				} else {
					if z.Vec[zpbh] == nil {
						z.Vec[zpbh] = new(VTDblDbl)
					}
					var zxzu uint32
					zxzu, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zxzu > 0 {
						zxzu--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zpbh].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zpbh].Value, err = dc.ReadFloat64()
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
	for zpbh := range z.Vec {
		if z.Vec[zpbh] == nil {
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
			err = en.WriteFloat64(z.Vec[zpbh].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zpbh].Value)
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
	for zpbh := range z.Vec {
		if z.Vec[zpbh] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zpbh].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zpbh].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zaii uint32
	zaii, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zaii > 0 {
		zaii--
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
			var zsnb uint32
			zsnb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zsnb) {
				z.Vec = (z.Vec)[:zsnb]
			} else {
				z.Vec = make([]*VTDblDbl, zsnb)
			}
			for zpbh := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zpbh] = nil
				} else {
					if z.Vec[zpbh] == nil {
						z.Vec[zpbh] = new(VTDblDbl)
					}
					var zalq uint32
					zalq, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zalq > 0 {
						zalq--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zpbh].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zpbh].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	for zpbh := range z.Vec {
		if z.Vec[zpbh] == nil {
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
	var zpoq uint32
	zpoq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpoq > 0 {
		zpoq--
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
			var zgnc uint32
			zgnc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zgnc) {
				z.Vec = (z.Vec)[:zgnc]
			} else {
				z.Vec = make([]*VTDblInt, zgnc)
			}
			for zswq := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zswq] = nil
				} else {
					if z.Vec[zswq] == nil {
						z.Vec[zswq] = new(VTDblInt)
					}
					var zzga uint32
					zzga, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zzga > 0 {
						zzga--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zswq].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zswq].Value, err = dc.ReadInt64()
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
	for zswq := range z.Vec {
		if z.Vec[zswq] == nil {
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
			err = en.WriteFloat64(z.Vec[zswq].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zswq].Value)
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
	for zswq := range z.Vec {
		if z.Vec[zswq] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zswq].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zswq].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcza uint32
	zcza, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcza > 0 {
		zcza--
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
			var zink uint32
			zink, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zink) {
				z.Vec = (z.Vec)[:zink]
			} else {
				z.Vec = make([]*VTDblInt, zink)
			}
			for zswq := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zswq] = nil
				} else {
					if z.Vec[zswq] == nil {
						z.Vec[zswq] = new(VTDblInt)
					}
					var zqki uint32
					zqki, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zqki > 0 {
						zqki--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zswq].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zswq].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	for zswq := range z.Vec {
		if z.Vec[zswq] == nil {
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
	var ztky uint32
	ztky, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztky > 0 {
		ztky--
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
			var zhda uint32
			zhda, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zhda) {
				z.Vec = (z.Vec)[:zhda]
			} else {
				z.Vec = make([]*VTDblStr, zhda)
			}
			for zbaa := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zbaa] = nil
				} else {
					if z.Vec[zbaa] == nil {
						z.Vec[zbaa] = new(VTDblStr)
					}
					var zopo uint32
					zopo, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zopo > 0 {
						zopo--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zbaa].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zbaa].Value, err = dc.ReadString()
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
	for zbaa := range z.Vec {
		if z.Vec[zbaa] == nil {
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
			err = en.WriteFloat64(z.Vec[zbaa].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zbaa].Value)
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
	for zbaa := range z.Vec {
		if z.Vec[zbaa] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zbaa].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zbaa].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zvfo uint32
	zvfo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zvfo > 0 {
		zvfo--
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
			var zlqi uint32
			zlqi, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zlqi) {
				z.Vec = (z.Vec)[:zlqi]
			} else {
				z.Vec = make([]*VTDblStr, zlqi)
			}
			for zbaa := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zbaa] = nil
				} else {
					if z.Vec[zbaa] == nil {
						z.Vec[zbaa] = new(VTDblStr)
					}
					var zexs uint32
					zexs, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zexs > 0 {
						zexs--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zbaa].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zbaa].Value, bts, err = msgp.ReadStringBytes(bts)
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
	for zbaa := range z.Vec {
		if z.Vec[zbaa] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zbaa].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zzrk uint32
	zzrk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zzrk > 0 {
		zzrk--
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
			var zczt uint32
			zczt, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zczt) {
				z.Vec = (z.Vec)[:zczt]
			} else {
				z.Vec = make([]int64, zczt)
			}
			for zfzx := range z.Vec {
				z.Vec[zfzx], err = dc.ReadInt64()
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
	for zfzx := range z.Vec {
		err = en.WriteInt64(z.Vec[zfzx])
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
	for zfzx := range z.Vec {
		o = msgp.AppendInt64(o, z.Vec[zfzx])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
			var zjjf uint32
			zjjf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zjjf) {
				z.Vec = (z.Vec)[:zjjf]
			} else {
				z.Vec = make([]int64, zjjf)
			}
			for zfzx := range z.Vec {
				z.Vec[zfzx], bts, err = msgp.ReadInt64Bytes(bts)
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
	var zhrc uint32
	zhrc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhrc > 0 {
		zhrc--
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
			var zovg uint32
			zovg, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zovg) {
				z.Vec = (z.Vec)[:zovg]
			} else {
				z.Vec = make([]*VTIntDbl, zovg)
			}
			for zuyz := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zuyz] = nil
				} else {
					if z.Vec[zuyz] == nil {
						z.Vec[zuyz] = new(VTIntDbl)
					}
					var zpvv uint32
					zpvv, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zpvv > 0 {
						zpvv--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zuyz].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zuyz].Value, err = dc.ReadFloat64()
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
	for zuyz := range z.Vec {
		if z.Vec[zuyz] == nil {
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
			err = en.WriteInt64(z.Vec[zuyz].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zuyz].Value)
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
	for zuyz := range z.Vec {
		if z.Vec[zuyz] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zuyz].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zuyz].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zkgs uint32
	zkgs, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zkgs > 0 {
		zkgs--
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
			var zxak uint32
			zxak, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zxak) {
				z.Vec = (z.Vec)[:zxak]
			} else {
				z.Vec = make([]*VTIntDbl, zxak)
			}
			for zuyz := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zuyz] = nil
				} else {
					if z.Vec[zuyz] == nil {
						z.Vec[zuyz] = new(VTIntDbl)
					}
					var zgqa uint32
					zgqa, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zgqa > 0 {
						zgqa--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zuyz].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zuyz].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	for zuyz := range z.Vec {
		if z.Vec[zuyz] == nil {
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
	var zebz uint32
	zebz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zebz > 0 {
		zebz--
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
			var zbfe uint32
			zbfe, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zbfe) {
				z.Vec = (z.Vec)[:zbfe]
			} else {
				z.Vec = make([]*VTIntInt, zbfe)
			}
			for zrmn := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zrmn] = nil
				} else {
					if z.Vec[zrmn] == nil {
						z.Vec[zrmn] = new(VTIntInt)
					}
					var zkhx uint32
					zkhx, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zkhx > 0 {
						zkhx--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zrmn].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zrmn].Value, err = dc.ReadInt64()
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
	for zrmn := range z.Vec {
		if z.Vec[zrmn] == nil {
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
			err = en.WriteInt64(z.Vec[zrmn].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zrmn].Value)
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
	for zrmn := range z.Vec {
		if z.Vec[zrmn] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zrmn].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zrmn].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zlzz uint32
	zlzz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zlzz > 0 {
		zlzz--
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
			var zfna uint32
			zfna, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zfna) {
				z.Vec = (z.Vec)[:zfna]
			} else {
				z.Vec = make([]*VTIntInt, zfna)
			}
			for zrmn := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zrmn] = nil
				} else {
					if z.Vec[zrmn] == nil {
						z.Vec[zrmn] = new(VTIntInt)
					}
					var zvta uint32
					zvta, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zvta > 0 {
						zvta--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zrmn].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zrmn].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	for zrmn := range z.Vec {
		if z.Vec[zrmn] == nil {
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
	var zwzl uint32
	zwzl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwzl > 0 {
		zwzl--
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
			var zqqf uint32
			zqqf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zqqf) {
				z.Vec = (z.Vec)[:zqqf]
			} else {
				z.Vec = make([]*VTIntStr, zqqf)
			}
			for zyyq := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zyyq] = nil
				} else {
					if z.Vec[zyyq] == nil {
						z.Vec[zyyq] = new(VTIntStr)
					}
					var zbuc uint32
					zbuc, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zbuc > 0 {
						zbuc--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zyyq].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zyyq].Value, err = dc.ReadString()
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
	for zyyq := range z.Vec {
		if z.Vec[zyyq] == nil {
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
			err = en.WriteInt64(z.Vec[zyyq].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zyyq].Value)
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
	for zyyq := range z.Vec {
		if z.Vec[zyyq] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zyyq].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zyyq].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqnj uint32
	zqnj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqnj > 0 {
		zqnj--
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
			var zywr uint32
			zywr, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zywr) {
				z.Vec = (z.Vec)[:zywr]
			} else {
				z.Vec = make([]*VTIntStr, zywr)
			}
			for zyyq := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zyyq] = nil
				} else {
					if z.Vec[zyyq] == nil {
						z.Vec[zyyq] = new(VTIntStr)
					}
					var zncg uint32
					zncg, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zncg > 0 {
						zncg--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zyyq].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zyyq].Value, bts, err = msgp.ReadStringBytes(bts)
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
	for zyyq := range z.Vec {
		if z.Vec[zyyq] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zyyq].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdtl uint32
	zdtl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdtl > 0 {
		zdtl--
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
			var znyo uint32
			znyo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(znyo) {
				z.Vec = (z.Vec)[:znyo]
			} else {
				z.Vec = make([]string, znyo)
			}
			for zllk := range z.Vec {
				z.Vec[zllk], err = dc.ReadString()
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
	for zllk := range z.Vec {
		err = en.WriteString(z.Vec[zllk])
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
	for zllk := range z.Vec {
		o = msgp.AppendString(o, z.Vec[zllk])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
			var zlfo uint32
			zlfo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zlfo) {
				z.Vec = (z.Vec)[:zlfo]
			} else {
				z.Vec = make([]string, zlfo)
			}
			for zllk := range z.Vec {
				z.Vec[zllk], bts, err = msgp.ReadStringBytes(bts)
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
	for zllk := range z.Vec {
		s += msgp.StringPrefixSize + len(z.Vec[zllk])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zvxs uint32
	zvxs, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zvxs > 0 {
		zvxs--
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
			var zazw uint32
			zazw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zazw) {
				z.Vec = (z.Vec)[:zazw]
			} else {
				z.Vec = make([]*VTStrDbl, zazw)
			}
			for zhsv := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zhsv] = nil
				} else {
					if z.Vec[zhsv] == nil {
						z.Vec[zhsv] = new(VTStrDbl)
					}
					var zexe uint32
					zexe, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zexe > 0 {
						zexe--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zhsv].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zhsv].Value, err = dc.ReadFloat64()
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
	for zhsv := range z.Vec {
		if z.Vec[zhsv] == nil {
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
			err = en.WriteString(z.Vec[zhsv].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zhsv].Value)
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
	for zhsv := range z.Vec {
		if z.Vec[zhsv] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zhsv].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zhsv].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjha uint32
	zjha, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjha > 0 {
		zjha--
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
			var zqux uint32
			zqux, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zqux) {
				z.Vec = (z.Vec)[:zqux]
			} else {
				z.Vec = make([]*VTStrDbl, zqux)
			}
			for zhsv := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zhsv] = nil
				} else {
					if z.Vec[zhsv] == nil {
						z.Vec[zhsv] = new(VTStrDbl)
					}
					var zdaa uint32
					zdaa, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zdaa > 0 {
						zdaa--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zhsv].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zhsv].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	for zhsv := range z.Vec {
		if z.Vec[zhsv] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zhsv].Key) + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrho uint32
	zrho, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrho > 0 {
		zrho--
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
			var znxv uint32
			znxv, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(znxv) {
				z.Vec = (z.Vec)[:znxv]
			} else {
				z.Vec = make([]*VTStrInt, znxv)
			}
			for zazk := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zazk] = nil
				} else {
					if z.Vec[zazk] == nil {
						z.Vec[zazk] = new(VTStrInt)
					}
					var zpay uint32
					zpay, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zpay > 0 {
						zpay--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zazk].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zazk].Value, err = dc.ReadInt64()
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
	for zazk := range z.Vec {
		if z.Vec[zazk] == nil {
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
			err = en.WriteString(z.Vec[zazk].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zazk].Value)
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
	for zazk := range z.Vec {
		if z.Vec[zazk] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zazk].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zazk].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zoss uint32
	zoss, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zoss > 0 {
		zoss--
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
			var zqcn uint32
			zqcn, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zqcn) {
				z.Vec = (z.Vec)[:zqcn]
			} else {
				z.Vec = make([]*VTStrInt, zqcn)
			}
			for zazk := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zazk] = nil
				} else {
					if z.Vec[zazk] == nil {
						z.Vec[zazk] = new(VTStrInt)
					}
					var zctu uint32
					zctu, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zctu > 0 {
						zctu--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zazk].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zazk].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	for zazk := range z.Vec {
		if z.Vec[zazk] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zazk].Key) + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zmct uint32
	zmct, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zmct > 0 {
		zmct--
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
			var zozv uint32
			zozv, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zozv) {
				z.Vec = (z.Vec)[:zozv]
			} else {
				z.Vec = make([]*VTStrStr, zozv)
			}
			for zgza := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zgza] = nil
				} else {
					if z.Vec[zgza] == nil {
						z.Vec[zgza] = new(VTStrStr)
					}
					var zpyn uint32
					zpyn, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zpyn > 0 {
						zpyn--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zgza].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zgza].Value, err = dc.ReadString()
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
	for zgza := range z.Vec {
		if z.Vec[zgza] == nil {
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
			err = en.WriteString(z.Vec[zgza].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zgza].Value)
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
	for zgza := range z.Vec {
		if z.Vec[zgza] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zgza].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zgza].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zaep uint32
	zaep, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zaep > 0 {
		zaep--
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
			var zhid uint32
			zhid, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zhid) {
				z.Vec = (z.Vec)[:zhid]
			} else {
				z.Vec = make([]*VTStrStr, zhid)
			}
			for zgza := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zgza] = nil
				} else {
					if z.Vec[zgza] == nil {
						z.Vec[zgza] = new(VTStrStr)
					}
					var zxal uint32
					zxal, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zxal > 0 {
						zxal--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zgza].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zgza].Value, bts, err = msgp.ReadStringBytes(bts)
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
	for zgza := range z.Vec {
		if z.Vec[zgza] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zgza].Key) + 6 + msgp.StringPrefixSize + len(z.Vec[zgza].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
	var ztzt uint32
	ztzt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztzt > 0 {
		ztzt--
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
