package fvec

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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Float64Size))
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Int64Size))
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
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
func (z *VMStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zzak uint32
			zzak, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zzak > 0 {
				z.Vec = make(map[string]float64, zzak)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zzak > 0 {
				zzak--
				var zmbt string
				var zvls float64
				zmbt, err = dc.ReadString()
				if err != nil {
					return
				}
				zvls, err = dc.ReadFloat64()
				if err != nil {
					return
				}
				z.Vec[zmbt] = zvls
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zmbt, zvls := range z.Vec {
		err = en.WriteString(zmbt)
		if err != nil {
			return
		}
		err = en.WriteFloat64(zvls)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zmbt, zvls := range z.Vec {
		o = msgp.AppendString(o, zmbt)
		o = msgp.AppendFloat64(o, zvls)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbtz uint32
	zbtz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbtz > 0 {
		zbtz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zsym uint32
			zsym, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zsym > 0 {
				z.Vec = make(map[string]float64, zsym)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zsym > 0 {
				var zmbt string
				var zvls float64
				zsym--
				zmbt, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zvls, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
				z.Vec[zmbt] = zvls
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
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zmbt, zvls := range z.Vec {
			_ = zvls
			s += msgp.StringPrefixSize + len(zmbt) + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zdqi uint32
			zdqi, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zdqi > 0 {
				z.Vec = make(map[string]int64, zdqi)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zdqi > 0 {
				zdqi--
				var zgeu string
				var zdtr int64
				zgeu, err = dc.ReadString()
				if err != nil {
					return
				}
				zdtr, err = dc.ReadInt64()
				if err != nil {
					return
				}
				z.Vec[zgeu] = zdtr
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zgeu, zdtr := range z.Vec {
		err = en.WriteString(zgeu)
		if err != nil {
			return
		}
		err = en.WriteInt64(zdtr)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zgeu, zdtr := range z.Vec {
		o = msgp.AppendString(o, zgeu)
		o = msgp.AppendInt64(o, zdtr)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zyco uint32
	zyco, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zyco > 0 {
		zyco--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zhgh uint32
			zhgh, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zhgh > 0 {
				z.Vec = make(map[string]int64, zhgh)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zhgh > 0 {
				var zgeu string
				var zdtr int64
				zhgh--
				zgeu, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zdtr, bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
				z.Vec[zgeu] = zdtr
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
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zgeu, zdtr := range z.Vec {
			_ = zdtr
			s += msgp.StringPrefixSize + len(zgeu) + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zjhy uint32
			zjhy, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zjhy > 0 {
				z.Vec = make(map[string]string, zjhy)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zjhy > 0 {
				zjhy--
				var zovg string
				var zsey string
				zovg, err = dc.ReadString()
				if err != nil {
					return
				}
				zsey, err = dc.ReadString()
				if err != nil {
					return
				}
				z.Vec[zovg] = zsey
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zovg, zsey := range z.Vec {
		err = en.WriteString(zovg)
		if err != nil {
			return
		}
		err = en.WriteString(zsey)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VMStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zovg, zsey := range z.Vec {
		o = msgp.AppendString(o, zovg)
		o = msgp.AppendString(o, zsey)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var znuf uint32
	znuf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for znuf > 0 {
		znuf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var znjj uint32
			znjj, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && znjj > 0 {
				z.Vec = make(map[string]string, znjj)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for znjj > 0 {
				var zovg string
				var zsey string
				znjj--
				zovg, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zsey, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Vec[zovg] = zsey
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
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zovg, zsey := range z.Vec {
			_ = zsey
			s += msgp.StringPrefixSize + len(zovg) + msgp.StringPrefixSize + len(zsey)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zfgq uint32
			zfgq, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zfgq > 0 {
				z.Vec = make(map[string]*VTDblDbl, zfgq)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zfgq > 0 {
				zfgq--
				var zhhj string
				var zuvr *VTDblDbl
				zhhj, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zuvr = nil
				} else {
					if zuvr == nil {
						zuvr = new(VTDblDbl)
					}
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
						case "key":
							zuvr.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							zuvr.Value, err = dc.ReadFloat64()
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
				z.Vec[zhhj] = zuvr
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zhhj, zuvr := range z.Vec {
		err = en.WriteString(zhhj)
		if err != nil {
			return
		}
		if zuvr == nil {
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
			err = en.WriteFloat64(zuvr.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(zuvr.Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zhhj, zuvr := range z.Vec {
		o = msgp.AppendString(o, zhhj)
		if zuvr == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, zuvr.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zuvr.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var zlur uint32
			zlur, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zlur > 0 {
				z.Vec = make(map[string]*VTDblDbl, zlur)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zlur > 0 {
				var zhhj string
				var zuvr *VTDblDbl
				zlur--
				zhhj, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zuvr = nil
				} else {
					if zuvr == nil {
						zuvr = new(VTDblDbl)
					}
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
						case "key":
							zuvr.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zuvr.Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
				z.Vec[zhhj] = zuvr
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
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zhhj, zuvr := range z.Vec {
			_ = zuvr
			s += msgp.StringPrefixSize + len(zhhj)
			if zuvr == nil {
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
		case "vec":
			var zarl uint32
			zarl, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zarl > 0 {
				z.Vec = make(map[string]*VTDblInt, zarl)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zarl > 0 {
				zarl--
				var zfvi string
				var zzrg *VTDblInt
				zfvi, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zzrg = nil
				} else {
					if zzrg == nil {
						zzrg = new(VTDblInt)
					}
					var zctz uint32
					zctz, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zctz > 0 {
						zctz--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zzrg.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							zzrg.Value, err = dc.ReadInt64()
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
func (z *VMStrTPDblInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
		if zzrg == nil {
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
			err = en.WriteFloat64(zzrg.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(zzrg.Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zfvi, zzrg := range z.Vec {
		o = msgp.AppendString(o, zfvi)
		if zzrg == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, zzrg.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zzrg.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zljl uint32
	zljl, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zljl > 0 {
		zljl--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zziv uint32
			zziv, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zziv > 0 {
				z.Vec = make(map[string]*VTDblInt, zziv)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zziv > 0 {
				var zfvi string
				var zzrg *VTDblInt
				zziv--
				zfvi, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zzrg = nil
				} else {
					if zzrg == nil {
						zzrg = new(VTDblInt)
					}
					var zabj uint32
					zabj, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zabj > 0 {
						zabj--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zzrg.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zzrg.Value, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *VMStrTPDblInt) Msgsize() (s int) {
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zfvi, zzrg := range z.Vec {
			_ = zzrg
			s += msgp.StringPrefixSize + len(zfvi)
			if zzrg == nil {
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
	var zgvb uint32
	zgvb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgvb > 0 {
		zgvb--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zqzg uint32
			zqzg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zqzg > 0 {
				z.Vec = make(map[string]*VTDblStr, zqzg)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zqzg > 0 {
				zqzg--
				var zmlx string
				var zvbw *VTDblStr
				zmlx, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					zvbw = nil
				} else {
					if zvbw == nil {
						zvbw = new(VTDblStr)
					}
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
						case "key":
							zvbw.Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							zvbw.Value, err = dc.ReadString()
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
				z.Vec[zmlx] = zvbw
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zmlx, zvbw := range z.Vec {
		err = en.WriteString(zmlx)
		if err != nil {
			return
		}
		if zvbw == nil {
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
			err = en.WriteFloat64(zvbw.Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(zvbw.Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zmlx, zvbw := range z.Vec {
		o = msgp.AppendString(o, zmlx)
		if zvbw == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, zvbw.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, zvbw.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zakb uint32
	zakb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zakb > 0 {
		zakb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zsdj uint32
			zsdj, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zsdj > 0 {
				z.Vec = make(map[string]*VTDblStr, zsdj)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zsdj > 0 {
				var zmlx string
				var zvbw *VTDblStr
				zsdj--
				zmlx, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					zvbw = nil
				} else {
					if zvbw == nil {
						zvbw = new(VTDblStr)
					}
					var zsgp uint32
					zsgp, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zsgp > 0 {
						zsgp--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							zvbw.Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							zvbw.Value, bts, err = msgp.ReadStringBytes(bts)
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
				z.Vec[zmlx] = zvbw
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
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zmlx, zvbw := range z.Vec {
			_ = zvbw
			s += msgp.StringPrefixSize + len(zmlx)
			if zvbw == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(zvbw.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zibu uint32
			zibu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zibu > 0 {
				z.Vec = make(map[string]*VTIntDbl, zibu)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zibu > 0 {
				zibu--
				var zngc string
				var zwfl *VTIntDbl
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
						zwfl = new(VTIntDbl)
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
							zwfl.Key, err = dc.ReadInt64()
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
func (z *VMStrTPIntDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
			err = en.WriteInt64(zwfl.Key)
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
func (z *VMStrTPIntDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zngc, zwfl := range z.Vec {
		o = msgp.AppendString(o, zngc)
		if zwfl == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zwfl.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zwfl.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var zdit uint32
			zdit, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zdit > 0 {
				z.Vec = make(map[string]*VTIntDbl, zdit)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zdit > 0 {
				var zngc string
				var zwfl *VTIntDbl
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
						zwfl = new(VTIntDbl)
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
							zwfl.Key, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *VMStrTPIntDbl) Msgsize() (s int) {
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zngc, zwfl := range z.Vec {
			_ = zwfl
			s += msgp.StringPrefixSize + len(zngc)
			if zwfl == nil {
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
		case "vec":
			var ztoj uint32
			ztoj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && ztoj > 0 {
				z.Vec = make(map[string]*VTIntInt, ztoj)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for ztoj > 0 {
				ztoj--
				var zoqj string
				var zmqr *VTIntInt
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
						zmqr = new(VTIntInt)
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
							zmqr.Key, err = dc.ReadInt64()
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
func (z *VMStrTPIntInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
			err = en.WriteInt64(zmqr.Key)
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
func (z *VMStrTPIntInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zoqj, zmqr := range z.Vec {
		o = msgp.AppendString(o, zoqj)
		if zmqr == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zmqr.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zmqr.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var zxzy uint32
			zxzy, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zxzy > 0 {
				z.Vec = make(map[string]*VTIntInt, zxzy)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zxzy > 0 {
				var zoqj string
				var zmqr *VTIntInt
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
						zmqr = new(VTIntInt)
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
							zmqr.Key, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *VMStrTPIntInt) Msgsize() (s int) {
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zoqj, zmqr := range z.Vec {
			_ = zmqr
			s += msgp.StringPrefixSize + len(zoqj)
			if zmqr == nil {
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
		case "vec":
			var znpn uint32
			znpn, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && znpn > 0 {
				z.Vec = make(map[string]*VTIntStr, znpn)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for znpn > 0 {
				znpn--
				var zrod string
				var zmbn *VTIntStr
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
						zmbn = new(VTIntStr)
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
							zmbn.Key, err = dc.ReadInt64()
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
func (z *VMStrTPIntStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
			err = en.WriteInt64(zmbn.Key)
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
func (z *VMStrTPIntStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zrod, zmbn := range z.Vec {
		o = msgp.AppendString(o, zrod)
		if zmbn == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, zmbn.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, zmbn.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var zhdt uint32
			zhdt, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zhdt > 0 {
				z.Vec = make(map[string]*VTIntStr, zhdt)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zhdt > 0 {
				var zrod string
				var zmbn *VTIntStr
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
						zmbn = new(VTIntStr)
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
							zmbn.Key, bts, err = msgp.ReadInt64Bytes(bts)
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
func (z *VMStrTPIntStr) Msgsize() (s int) {
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zrod, zmbn := range z.Vec {
			_ = zmbn
			s += msgp.StringPrefixSize + len(zrod)
			if zmbn == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(zmbn.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zaps uint32
			zaps, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zaps > 0 {
				z.Vec = make(map[string]*VTStrDbl, zaps)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zaps > 0 {
				zaps--
				var zayo string
				var zrsu *VTStrDbl
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
						zrsu = new(VTStrDbl)
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
							zrsu.Key, err = dc.ReadString()
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
func (z *VMStrTPStrDbl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
			err = en.WriteString(zrsu.Key)
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
func (z *VMStrTPStrDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zayo, zrsu := range z.Vec {
		o = msgp.AppendString(o, zayo)
		if zrsu == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zrsu.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, zrsu.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var zmyy uint32
			zmyy, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zmyy > 0 {
				z.Vec = make(map[string]*VTStrDbl, zmyy)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmyy > 0 {
				var zayo string
				var zrsu *VTStrDbl
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
						zrsu = new(VTStrDbl)
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
							zrsu.Key, bts, err = msgp.ReadStringBytes(bts)
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
func (z *VMStrTPStrDbl) Msgsize() (s int) {
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zayo, zrsu := range z.Vec {
			_ = zrsu
			s += msgp.StringPrefixSize + len(zayo)
			if zrsu == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zrsu.Key) + 6 + msgp.Float64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zcnq uint32
			zcnq, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zcnq > 0 {
				z.Vec = make(map[string]*VTStrInt, zcnq)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zcnq > 0 {
				zcnq--
				var zvgw string
				var zffb *VTStrInt
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
						zffb = new(VTStrInt)
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
							zffb.Key, err = dc.ReadString()
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
func (z *VMStrTPStrInt) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
			err = en.WriteString(zffb.Key)
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
func (z *VMStrTPStrInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zvgw, zffb := range z.Vec {
		o = msgp.AppendString(o, zvgw)
		if zffb == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, zffb.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, zffb.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var znuz uint32
			znuz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && znuz > 0 {
				z.Vec = make(map[string]*VTStrInt, znuz)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for znuz > 0 {
				var zvgw string
				var zffb *VTStrInt
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
						zffb = new(VTStrInt)
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
							zffb.Key, bts, err = msgp.ReadStringBytes(bts)
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
func (z *VMStrTPStrInt) Msgsize() (s int) {
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zvgw, zffb := range z.Vec {
			_ = zffb
			s += msgp.StringPrefixSize + len(zvgw)
			if zffb == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(zffb.Key) + 6 + msgp.Int64Size
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VMStrTPStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "vec":
			var zmyg uint32
			zmyg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Vec == nil && zmyg > 0 {
				z.Vec = make(map[string]*VTStrStr, zmyg)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zmyg > 0 {
				zmyg--
				var zmzo string
				var ztar *VTStrStr
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
						ztar = new(VTStrStr)
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
							ztar.Key, err = dc.ReadString()
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
func (z *VMStrTPStrStr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
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
			err = en.WriteString(ztar.Key)
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
func (z *VMStrTPStrStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendMapHeader(o, uint32(len(z.Vec)))
	for zmzo, ztar := range z.Vec {
		o = msgp.AppendString(o, zmzo)
		if ztar == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, ztar.Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, ztar.Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VMStrTPStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var zddv uint32
			zddv, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Vec == nil && zddv > 0 {
				z.Vec = make(map[string]*VTStrStr, zddv)
			} else if len(z.Vec) > 0 {
				for key := range z.Vec {
					delete(z.Vec, key)
				}
			}
			for zddv > 0 {
				var zmzo string
				var ztar *VTStrStr
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
						ztar = new(VTStrStr)
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
							ztar.Key, bts, err = msgp.ReadStringBytes(bts)
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
func (z *VMStrTPStrStr) Msgsize() (s int) {
	s = 1 + 4 + msgp.MapHeaderSize
	if z.Vec != nil {
		for zmzo, ztar := range z.Vec {
			_ = ztar
			s += msgp.StringPrefixSize + len(zmzo)
			if ztar == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 4 + msgp.StringPrefixSize + len(ztar.Key) + 6 + msgp.StringPrefixSize + len(ztar.Value)
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VName) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgpy uint32
	zgpy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgpy > 0 {
		zgpy--
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
			var zclm uint32
			zclm, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(zclm) {
				z.Tags = (z.Tags)[:zclm]
			} else {
				z.Tags = make([]*Tag, zclm)
			}
			for zfsf := range z.Tags {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Tags[zfsf] = nil
				} else {
					if z.Tags[zfsf] == nil {
						z.Tags[zfsf] = new(Tag)
					}
					var zxiu uint32
					zxiu, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zxiu > 0 {
						zxiu--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Name":
							z.Tags[zfsf].Name, err = dc.ReadString()
							if err != nil {
								return
							}
						case "Value":
							z.Tags[zfsf].Value, err = dc.ReadString()
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
	for zfsf := range z.Tags {
		if z.Tags[zfsf] == nil {
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
			err = en.WriteString(z.Tags[zfsf].Name)
			if err != nil {
				return
			}
			// write "Value"
			err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Tags[zfsf].Value)
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
	for zfsf := range z.Tags {
		if z.Tags[zfsf] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "Name"
			o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Tags[zfsf].Name)
			// string "Value"
			o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Tags[zfsf].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VName) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbey uint32
	zbey, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbey > 0 {
		zbey--
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
			var ztnd uint32
			ztnd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(ztnd) {
				z.Tags = (z.Tags)[:ztnd]
			} else {
				z.Tags = make([]*Tag, ztnd)
			}
			for zfsf := range z.Tags {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Tags[zfsf] = nil
				} else {
					if z.Tags[zfsf] == nil {
						z.Tags[zfsf] = new(Tag)
					}
					var ztja uint32
					ztja, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for ztja > 0 {
						ztja--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Name":
							z.Tags[zfsf].Name, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "Value":
							z.Tags[zfsf].Value, bts, err = msgp.ReadStringBytes(bts)
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
	for zfsf := range z.Tags {
		if z.Tags[zfsf] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 5 + msgp.StringPrefixSize + len(z.Tags[zfsf].Name) + 6 + msgp.StringPrefixSize + len(z.Tags[zfsf].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zded uint32
	zded, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zded > 0 {
		zded--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zmiy uint32
			zmiy, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zmiy) {
				z.Vec = (z.Vec)[:zmiy]
			} else {
				z.Vec = make([]float64, zmiy)
			}
			for zyrr := range z.Vec {
				z.Vec[zyrr], err = dc.ReadFloat64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zyrr := range z.Vec {
		err = en.WriteFloat64(z.Vec[zyrr])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSDbl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zyrr := range z.Vec {
		o = msgp.AppendFloat64(o, z.Vec[zyrr])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zkpr uint32
	zkpr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zkpr > 0 {
		zkpr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zucj uint32
			zucj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zucj) {
				z.Vec = (z.Vec)[:zucj]
			} else {
				z.Vec = make([]float64, zucj)
			}
			for zyrr := range z.Vec {
				z.Vec[zyrr], bts, err = msgp.ReadFloat64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Float64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjhy uint32
	zjhy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjhy > 0 {
		zjhy--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zevw uint32
			zevw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zevw) {
				z.Vec = (z.Vec)[:zevw]
			} else {
				z.Vec = make([]*VTDblDbl, zevw)
			}
			for ziog := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[ziog] = nil
				} else {
					if z.Vec[ziog] == nil {
						z.Vec[ziog] = new(VTDblDbl)
					}
					var zbtc uint32
					zbtc, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zbtc > 0 {
						zbtc--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[ziog].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[ziog].Value, err = dc.ReadFloat64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for ziog := range z.Vec {
		if z.Vec[ziog] == nil {
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
			err = en.WriteFloat64(z.Vec[ziog].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[ziog].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for ziog := range z.Vec {
		if z.Vec[ziog] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[ziog].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[ziog].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmkf uint32
	zmkf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmkf > 0 {
		zmkf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zrdg uint32
			zrdg, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zrdg) {
				z.Vec = (z.Vec)[:zrdg]
			} else {
				z.Vec = make([]*VTDblDbl, zrdg)
			}
			for ziog := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[ziog] = nil
				} else {
					if z.Vec[ziog] == nil {
						z.Vec[ziog] = new(VTDblDbl)
					}
					var zxqw uint32
					zxqw, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zxqw > 0 {
						zxqw--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[ziog].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[ziog].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for ziog := range z.Vec {
		if z.Vec[ziog] == nil {
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
		case "vec":
			var zmzc uint32
			zmzc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zmzc) {
				z.Vec = (z.Vec)[:zmzc]
			} else {
				z.Vec = make([]*VTDblInt, zmzc)
			}
			for zpzw := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zpzw] = nil
				} else {
					if z.Vec[zpzw] == nil {
						z.Vec[zpzw] = new(VTDblInt)
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
							z.Vec[zpzw].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zpzw].Value, err = dc.ReadInt64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zpzw := range z.Vec {
		if z.Vec[zpzw] == nil {
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
			err = en.WriteFloat64(z.Vec[zpzw].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zpzw].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zpzw := range z.Vec {
		if z.Vec[zpzw] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zpzw].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zpzw].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "vec":
			var zjkd uint32
			zjkd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zjkd) {
				z.Vec = (z.Vec)[:zjkd]
			} else {
				z.Vec = make([]*VTDblInt, zjkd)
			}
			for zpzw := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zpzw] = nil
				} else {
					if z.Vec[zpzw] == nil {
						z.Vec[zpzw] = new(VTDblInt)
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
							z.Vec[zpzw].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zpzw].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zpzw := range z.Vec {
		if z.Vec[zpzw] == nil {
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
		case "vec":
			var zeql uint32
			zeql, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zeql) {
				z.Vec = (z.Vec)[:zeql]
			} else {
				z.Vec = make([]*VTDblStr, zeql)
			}
			for zofl := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zofl] = nil
				} else {
					if z.Vec[zofl] == nil {
						z.Vec[zofl] = new(VTDblStr)
					}
					var zlhq uint32
					zlhq, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zlhq > 0 {
						zlhq--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zofl].Key, err = dc.ReadFloat64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zofl].Value, err = dc.ReadString()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zofl := range z.Vec {
		if z.Vec[zofl] == nil {
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
			err = en.WriteFloat64(z.Vec[zofl].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zofl].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zofl := range z.Vec {
		if z.Vec[zofl] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendFloat64(o, z.Vec[zofl].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zofl].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSDblStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
		case "vec":
			var znbp uint32
			znbp, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(znbp) {
				z.Vec = (z.Vec)[:znbp]
			} else {
				z.Vec = make([]*VTDblStr, znbp)
			}
			for zofl := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zofl] = nil
				} else {
					if z.Vec[zofl] == nil {
						z.Vec[zofl] = new(VTDblStr)
					}
					var zbgh uint32
					zbgh, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zbgh > 0 {
						zbgh--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "key":
							z.Vec[zofl].Key, bts, err = msgp.ReadFloat64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zofl].Value, bts, err = msgp.ReadStringBytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zofl := range z.Vec {
		if z.Vec[zofl] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Float64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zofl].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zidj uint32
	zidj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zidj > 0 {
		zidj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zumm uint32
			zumm, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zumm) {
				z.Vec = (z.Vec)[:zumm]
			} else {
				z.Vec = make([]int64, zumm)
			}
			for zmku := range z.Vec {
				z.Vec[zmku], err = dc.ReadInt64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zmku := range z.Vec {
		err = en.WriteInt64(z.Vec[zmku])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zmku := range z.Vec {
		o = msgp.AppendInt64(o, z.Vec[zmku])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpbh uint32
	zpbh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpbh > 0 {
		zpbh--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zcsj uint32
			zcsj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zcsj) {
				z.Vec = (z.Vec)[:zcsj]
			} else {
				z.Vec = make([]int64, zcsj)
			}
			for zmku := range z.Vec {
				z.Vec[zmku], bts, err = msgp.ReadInt64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize + (len(z.Vec) * (msgp.Int64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSIntDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zaii uint32
			zaii, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zaii) {
				z.Vec = (z.Vec)[:zaii]
			} else {
				z.Vec = make([]*VTIntDbl, zaii)
			}
			for zmjj := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zmjj] = nil
				} else {
					if z.Vec[zmjj] == nil {
						z.Vec[zmjj] = new(VTIntDbl)
					}
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
						case "key":
							z.Vec[zmjj].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zmjj].Value, err = dc.ReadFloat64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zmjj := range z.Vec {
		if z.Vec[zmjj] == nil {
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
			err = en.WriteInt64(z.Vec[zmjj].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zmjj].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zmjj := range z.Vec {
		if z.Vec[zmjj] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zmjj].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zmjj].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zswq uint32
			zswq, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zswq) {
				z.Vec = (z.Vec)[:zswq]
			} else {
				z.Vec = make([]*VTIntDbl, zswq)
			}
			for zmjj := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zmjj] = nil
				} else {
					if z.Vec[zmjj] == nil {
						z.Vec[zmjj] = new(VTIntDbl)
					}
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
						case "key":
							z.Vec[zmjj].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zmjj].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zmjj := range z.Vec {
		if z.Vec[zmjj] == nil {
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
		case "vec":
			var zcza uint32
			zcza, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zcza) {
				z.Vec = (z.Vec)[:zcza]
			} else {
				z.Vec = make([]*VTIntInt, zcza)
			}
			for zgnc := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zgnc] = nil
				} else {
					if z.Vec[zgnc] == nil {
						z.Vec[zgnc] = new(VTIntInt)
					}
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
						case "key":
							z.Vec[zgnc].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zgnc].Value, err = dc.ReadInt64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zgnc := range z.Vec {
		if z.Vec[zgnc] == nil {
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
			err = en.WriteInt64(z.Vec[zgnc].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zgnc].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zgnc := range z.Vec {
		if z.Vec[zgnc] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zgnc].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zgnc].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zbaa uint32
			zbaa, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zbaa) {
				z.Vec = (z.Vec)[:zbaa]
			} else {
				z.Vec = make([]*VTIntInt, zbaa)
			}
			for zgnc := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zgnc] = nil
				} else {
					if z.Vec[zgnc] == nil {
						z.Vec[zgnc] = new(VTIntInt)
					}
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
						case "key":
							z.Vec[zgnc].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zgnc].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zgnc := range z.Vec {
		if z.Vec[zgnc] == nil {
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
		case "vec":
			var zvfo uint32
			zvfo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zvfo) {
				z.Vec = (z.Vec)[:zvfo]
			} else {
				z.Vec = make([]*VTIntStr, zvfo)
			}
			for zhda := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zhda] = nil
				} else {
					if z.Vec[zhda] == nil {
						z.Vec[zhda] = new(VTIntStr)
					}
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
						case "key":
							z.Vec[zhda].Key, err = dc.ReadInt64()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zhda].Value, err = dc.ReadString()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zhda := range z.Vec {
		if z.Vec[zhda] == nil {
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
			err = en.WriteInt64(z.Vec[zhda].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zhda].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zhda := range z.Vec {
		if z.Vec[zhda] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendInt64(o, z.Vec[zhda].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zhda].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSIntStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zfzx uint32
			zfzx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zfzx) {
				z.Vec = (z.Vec)[:zfzx]
			} else {
				z.Vec = make([]*VTIntStr, zfzx)
			}
			for zhda := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zhda] = nil
				} else {
					if z.Vec[zhda] == nil {
						z.Vec[zhda] = new(VTIntStr)
					}
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
						case "key":
							z.Vec[zhda].Key, bts, err = msgp.ReadInt64Bytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zhda].Value, bts, err = msgp.ReadStringBytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zhda := range z.Vec {
		if z.Vec[zhda] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Vec[zhda].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxcd uint32
	zxcd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxcd > 0 {
		zxcd--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zjjf uint32
			zjjf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zjjf) {
				z.Vec = (z.Vec)[:zjjf]
			} else {
				z.Vec = make([]string, zjjf)
			}
			for zczt := range z.Vec {
				z.Vec[zczt], err = dc.ReadString()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zczt := range z.Vec {
		err = en.WriteString(z.Vec[zczt])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VSStr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zczt := range z.Vec {
		o = msgp.AppendString(o, z.Vec[zczt])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zuyz uint32
	zuyz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zuyz > 0 {
		zuyz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "vec":
			var zhrc uint32
			zhrc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zhrc) {
				z.Vec = (z.Vec)[:zhrc]
			} else {
				z.Vec = make([]string, zhrc)
			}
			for zczt := range z.Vec {
				z.Vec[zczt], bts, err = msgp.ReadStringBytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zczt := range z.Vec {
		s += msgp.StringPrefixSize + len(z.Vec[zczt])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zkgs uint32
			zkgs, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zkgs) {
				z.Vec = (z.Vec)[:zkgs]
			} else {
				z.Vec = make([]*VTStrDbl, zkgs)
			}
			for zovg := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zovg] = nil
				} else {
					if z.Vec[zovg] == nil {
						z.Vec[zovg] = new(VTStrDbl)
					}
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
						case "key":
							z.Vec[zovg].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zovg].Value, err = dc.ReadFloat64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zovg := range z.Vec {
		if z.Vec[zovg] == nil {
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
			err = en.WriteString(z.Vec[zovg].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteFloat64(z.Vec[zovg].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zovg := range z.Vec {
		if z.Vec[zovg] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zovg].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendFloat64(o, z.Vec[zovg].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrDbl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zrmn uint32
			zrmn, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zrmn) {
				z.Vec = (z.Vec)[:zrmn]
			} else {
				z.Vec = make([]*VTStrDbl, zrmn)
			}
			for zovg := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zovg] = nil
				} else {
					if z.Vec[zovg] == nil {
						z.Vec[zovg] = new(VTStrDbl)
					}
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
						case "key":
							z.Vec[zovg].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zovg].Value, bts, err = msgp.ReadFloat64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zovg := range z.Vec {
		if z.Vec[zovg] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zovg].Key) + 6 + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrInt) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zlzz uint32
			zlzz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zlzz) {
				z.Vec = (z.Vec)[:zlzz]
			} else {
				z.Vec = make([]*VTStrInt, zlzz)
			}
			for zbfe := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zbfe] = nil
				} else {
					if z.Vec[zbfe] == nil {
						z.Vec[zbfe] = new(VTStrInt)
					}
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
						case "key":
							z.Vec[zbfe].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zbfe].Value, err = dc.ReadInt64()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zbfe := range z.Vec {
		if z.Vec[zbfe] == nil {
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
			err = en.WriteString(z.Vec[zbfe].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteInt64(z.Vec[zbfe].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zbfe := range z.Vec {
		if z.Vec[zbfe] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zbfe].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendInt64(o, z.Vec[zbfe].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zyyq uint32
			zyyq, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zyyq) {
				z.Vec = (z.Vec)[:zyyq]
			} else {
				z.Vec = make([]*VTStrInt, zyyq)
			}
			for zbfe := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zbfe] = nil
				} else {
					if z.Vec[zbfe] == nil {
						z.Vec[zbfe] = new(VTStrInt)
					}
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
						case "key":
							z.Vec[zbfe].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zbfe].Value, bts, err = msgp.ReadInt64Bytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zbfe := range z.Vec {
		if z.Vec[zbfe] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zbfe].Key) + 6 + msgp.Int64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VSStrStr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zqnj uint32
			zqnj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zqnj) {
				z.Vec = (z.Vec)[:zqnj]
			} else {
				z.Vec = make([]*VTStrStr, zqnj)
			}
			for zqqf := range z.Vec {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Vec[zqqf] = nil
				} else {
					if z.Vec[zqqf] == nil {
						z.Vec[zqqf] = new(VTStrStr)
					}
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
						case "key":
							z.Vec[zqqf].Key, err = dc.ReadString()
							if err != nil {
								return
							}
						case "value":
							z.Vec[zqqf].Value, err = dc.ReadString()
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
	// map header, size 1
	// write "vec"
	err = en.Append(0x81, 0xa3, 0x76, 0x65, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Vec)))
	if err != nil {
		return
	}
	for zqqf := range z.Vec {
		if z.Vec[zqqf] == nil {
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
			err = en.WriteString(z.Vec[zqqf].Key)
			if err != nil {
				return
			}
			// write "value"
			err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Vec[zqqf].Value)
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
	// map header, size 1
	// string "vec"
	o = append(o, 0x81, 0xa3, 0x76, 0x65, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Vec)))
	for zqqf := range z.Vec {
		if z.Vec[zqqf] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "key"
			o = append(o, 0x82, 0xa3, 0x6b, 0x65, 0x79)
			o = msgp.AppendString(o, z.Vec[zqqf].Key)
			// string "value"
			o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
			o = msgp.AppendString(o, z.Vec[zqqf].Value)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VSStrStr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
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
		case "vec":
			var zllk uint32
			zllk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Vec) >= int(zllk) {
				z.Vec = (z.Vec)[:zllk]
			} else {
				z.Vec = make([]*VTStrStr, zllk)
			}
			for zqqf := range z.Vec {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Vec[zqqf] = nil
				} else {
					if z.Vec[zqqf] == nil {
						z.Vec[zqqf] = new(VTStrStr)
					}
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
						case "key":
							z.Vec[zqqf].Key, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "value":
							z.Vec[zqqf].Value, bts, err = msgp.ReadStringBytes(bts)
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
	s = 1 + 4 + msgp.ArrayHeaderSize
	for zqqf := range z.Vec {
		if z.Vec[zqqf] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.StringPrefixSize + len(z.Vec[zqqf].Key) + 6 + msgp.StringPrefixSize + len(z.Vec[zqqf].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VTDblDbl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var znyo uint32
	znyo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for znyo > 0 {
		znyo--
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
	var zlfo uint32
	zlfo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zlfo > 0 {
		zlfo--
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
	var zhsv uint32
	zhsv, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhsv > 0 {
		zhsv--
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
	var znxv uint32
	znxv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for znxv > 0 {
		znxv--
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
	var zoss uint32
	zoss, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zoss > 0 {
		zoss--
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
	var zqcn uint32
	zqcn, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqcn > 0 {
		zqcn--
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
	var zctu uint32
	zctu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zctu > 0 {
		zctu--
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
	var zgza uint32
	zgza, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zgza > 0 {
		zgza--
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
