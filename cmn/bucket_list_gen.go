// Package provides common low-level types and utilities for all aistore projects
/*
 * Copyright (c) 2018-2020, NVIDIA CORPORATION. All rights reserved.
 */
package cmn

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BucketEntry) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "n":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "s":
			z.Size, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Size")
				return
			}
		case "cs":
			z.Checksum, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Checksum")
				return
			}
		case "a":
			z.Atime, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Atime")
				return
			}
		case "v":
			z.Version, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "t":
			z.TargetURL, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "TargetURL")
				return
			}
		case "c":
			z.Copies, err = dc.ReadInt16()
			if err != nil {
				err = msgp.WrapError(err, "Copies")
				return
			}
		case "f":
			z.Flags, err = dc.ReadUint16()
			if err != nil {
				err = msgp.WrapError(err, "Flags")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketEntry) EncodeMsg(en *msgp.Writer) (err error) {
	// omitempty: check for empty values
	zb0001Len := uint32(8)
	var zb0001Mask uint8 /* 8 bits */
	if z.Size == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if z.Checksum == "" {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if z.Atime == "" {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if z.Version == "" {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	if z.TargetURL == "" {
		zb0001Len--
		zb0001Mask |= 0x20
	}
	if z.Copies == 0 {
		zb0001Len--
		zb0001Mask |= 0x40
	}
	if z.Flags == 0 {
		zb0001Len--
		zb0001Mask |= 0x80
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "n"
	err = en.Append(0xa1, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	if (zb0001Mask & 0x2) == 0 { // if not empty
		// write "s"
		err = en.Append(0xa1, 0x73)
		if err != nil {
			return
		}
		err = en.WriteInt64(z.Size)
		if err != nil {
			err = msgp.WrapError(err, "Size")
			return
		}
	}
	if (zb0001Mask & 0x4) == 0 { // if not empty
		// write "cs"
		err = en.Append(0xa2, 0x63, 0x73)
		if err != nil {
			return
		}
		err = en.WriteString(z.Checksum)
		if err != nil {
			err = msgp.WrapError(err, "Checksum")
			return
		}
	}
	if (zb0001Mask & 0x8) == 0 { // if not empty
		// write "a"
		err = en.Append(0xa1, 0x61)
		if err != nil {
			return
		}
		err = en.WriteString(z.Atime)
		if err != nil {
			err = msgp.WrapError(err, "Atime")
			return
		}
	}
	if (zb0001Mask & 0x10) == 0 { // if not empty
		// write "v"
		err = en.Append(0xa1, 0x76)
		if err != nil {
			return
		}
		err = en.WriteString(z.Version)
		if err != nil {
			err = msgp.WrapError(err, "Version")
			return
		}
	}
	if (zb0001Mask & 0x20) == 0 { // if not empty
		// write "t"
		err = en.Append(0xa1, 0x74)
		if err != nil {
			return
		}
		err = en.WriteString(z.TargetURL)
		if err != nil {
			err = msgp.WrapError(err, "TargetURL")
			return
		}
	}
	if (zb0001Mask & 0x40) == 0 { // if not empty
		// write "c"
		err = en.Append(0xa1, 0x63)
		if err != nil {
			return
		}
		err = en.WriteInt16(z.Copies)
		if err != nil {
			err = msgp.WrapError(err, "Copies")
			return
		}
	}
	if (zb0001Mask & 0x80) == 0 { // if not empty
		// write "f"
		err = en.Append(0xa1, 0x66)
		if err != nil {
			return
		}
		err = en.WriteUint16(z.Flags)
		if err != nil {
			err = msgp.WrapError(err, "Flags")
			return
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketEntry) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Name) + 2 + msgp.Int64Size + 3 + msgp.StringPrefixSize + len(z.Checksum) + 2 + msgp.StringPrefixSize + len(z.Atime) + 2 + msgp.StringPrefixSize + len(z.Version) + 2 + msgp.StringPrefixSize + len(z.TargetURL) + 2 + msgp.Int16Size + 2 + msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BucketList) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "UUID":
			z.UUID, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "UUID")
				return
			}
		case "Entries":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Entries")
				return
			}
			if cap(z.Entries) >= int(zb0002) {
				z.Entries = (z.Entries)[:zb0002]
			} else {
				z.Entries = make([]*BucketEntry, zb0002)
			}
			for za0001 := range z.Entries {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Entries", za0001)
						return
					}
					z.Entries[za0001] = nil
				} else {
					if z.Entries[za0001] == nil {
						z.Entries[za0001] = new(BucketEntry)
					}
					err = z.Entries[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Entries", za0001)
						return
					}
				}
			}
		case "ContinuationToken":
			z.ContinuationToken, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ContinuationToken")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketList) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "UUID"
	err = en.Append(0x83, 0xa4, 0x55, 0x55, 0x49, 0x44)
	if err != nil {
		return
	}
	err = en.WriteString(z.UUID)
	if err != nil {
		err = msgp.WrapError(err, "UUID")
		return
	}
	// write "Entries"
	err = en.Append(0xa7, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Entries)))
	if err != nil {
		err = msgp.WrapError(err, "Entries")
		return
	}
	for za0001 := range z.Entries {
		if z.Entries[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Entries[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Entries", za0001)
				return
			}
		}
	}
	// write "ContinuationToken"
	err = en.Append(0xb1, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.ContinuationToken)
	if err != nil {
		err = msgp.WrapError(err, "ContinuationToken")
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketList) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.UUID) + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Entries {
		if z.Entries[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Entries[za0001].Msgsize()
		}
	}
	s += 18 + msgp.StringPrefixSize + len(z.ContinuationToken)
	return
}
