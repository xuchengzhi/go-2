// +build !notfastpath

// Code generated - DO NOT EDIT.

package codec

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

const (
	// ----- content types ----
	codecSelferC_UTF819781 = 1
	codecSelferC_RAW19781  = 0
	// ----- value types used ----
	codecSelferValueTypeArray19781 = 10
	codecSelferValueTypeMap19781   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey19781    = 2
	codecSelfer_containerMapValue19781  = 3
	codecSelfer_containerMapEnd19781    = 4
	codecSelfer_containerArrayElem19781 = 6
	codecSelfer_containerArrayEnd19781  = 7
)

var (
	codecSelferBitsize19781                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr19781 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer19781 struct{}

func init() {
	if GenVersion != 8 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			8, GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
	}
}

func (x *TestMammoth2) CodecEncodeSelf(e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(574)
			} else {
				r.WriteMapStart(574)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FIntf == nil {
					r.EncodeNil()
				} else {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						z.EncFallback(x.FIntf)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FIntf"))
				r.WriteMapElemValue()
				if x.FIntf == nil {
					r.EncodeNil()
				} else {
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						z.EncFallback(x.FIntf)
					}
				}
			}
			var yyn6 bool
			if x.FptrIntf == nil {
				yyn6 = true
				goto LABEL6
			}
		LABEL6:
			if yyr2 || yy2arr2 {
				if yyn6 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrIntf == nil {
						r.EncodeNil()
					} else {
						yy7 := *x.FptrIntf
						yym8 := z.EncBinary()
						_ = yym8
						if false {
						} else {
							z.EncFallback(yy7)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrIntf"))
				r.WriteMapElemValue()
				if yyn6 {
					r.EncodeNil()
				} else {
					if x.FptrIntf == nil {
						r.EncodeNil()
					} else {
						yy9 := *x.FptrIntf
						yym10 := z.EncBinary()
						_ = yym10
						if false {
						} else {
							z.EncFallback(yy9)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym12 := z.EncBinary()
				_ = yym12
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF819781, string(x.FString))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FString"))
				r.WriteMapElemValue()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF819781, string(x.FString))
				}
			}
			var yyn14 bool
			if x.FptrString == nil {
				yyn14 = true
				goto LABEL14
			}
		LABEL14:
			if yyr2 || yy2arr2 {
				if yyn14 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrString == nil {
						r.EncodeNil()
					} else {
						yy15 := *x.FptrString
						yym16 := z.EncBinary()
						_ = yym16
						if false {
						} else {
							r.EncodeString(codecSelferC_UTF819781, string(yy15))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrString"))
				r.WriteMapElemValue()
				if yyn14 {
					r.EncodeNil()
				} else {
					if x.FptrString == nil {
						r.EncodeNil()
					} else {
						yy17 := *x.FptrString
						yym18 := z.EncBinary()
						_ = yym18
						if false {
						} else {
							r.EncodeString(codecSelferC_UTF819781, string(yy17))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym20 := z.EncBinary()
				_ = yym20
				if false {
				} else {
					r.EncodeFloat32(float32(x.FFloat32))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FFloat32"))
				r.WriteMapElemValue()
				yym21 := z.EncBinary()
				_ = yym21
				if false {
				} else {
					r.EncodeFloat32(float32(x.FFloat32))
				}
			}
			var yyn22 bool
			if x.FptrFloat32 == nil {
				yyn22 = true
				goto LABEL22
			}
		LABEL22:
			if yyr2 || yy2arr2 {
				if yyn22 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrFloat32 == nil {
						r.EncodeNil()
					} else {
						yy23 := *x.FptrFloat32
						yym24 := z.EncBinary()
						_ = yym24
						if false {
						} else {
							r.EncodeFloat32(float32(yy23))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrFloat32"))
				r.WriteMapElemValue()
				if yyn22 {
					r.EncodeNil()
				} else {
					if x.FptrFloat32 == nil {
						r.EncodeNil()
					} else {
						yy25 := *x.FptrFloat32
						yym26 := z.EncBinary()
						_ = yym26
						if false {
						} else {
							r.EncodeFloat32(float32(yy25))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym28 := z.EncBinary()
				_ = yym28
				if false {
				} else {
					r.EncodeFloat64(float64(x.FFloat64))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FFloat64"))
				r.WriteMapElemValue()
				yym29 := z.EncBinary()
				_ = yym29
				if false {
				} else {
					r.EncodeFloat64(float64(x.FFloat64))
				}
			}
			var yyn30 bool
			if x.FptrFloat64 == nil {
				yyn30 = true
				goto LABEL30
			}
		LABEL30:
			if yyr2 || yy2arr2 {
				if yyn30 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrFloat64 == nil {
						r.EncodeNil()
					} else {
						yy31 := *x.FptrFloat64
						yym32 := z.EncBinary()
						_ = yym32
						if false {
						} else {
							r.EncodeFloat64(float64(yy31))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrFloat64"))
				r.WriteMapElemValue()
				if yyn30 {
					r.EncodeNil()
				} else {
					if x.FptrFloat64 == nil {
						r.EncodeNil()
					} else {
						yy33 := *x.FptrFloat64
						yym34 := z.EncBinary()
						_ = yym34
						if false {
						} else {
							r.EncodeFloat64(float64(yy33))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym36 := z.EncBinary()
				_ = yym36
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FUint"))
				r.WriteMapElemValue()
				yym37 := z.EncBinary()
				_ = yym37
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint))
				}
			}
			var yyn38 bool
			if x.FptrUint == nil {
				yyn38 = true
				goto LABEL38
			}
		LABEL38:
			if yyr2 || yy2arr2 {
				if yyn38 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrUint == nil {
						r.EncodeNil()
					} else {
						yy39 := *x.FptrUint
						yym40 := z.EncBinary()
						_ = yym40
						if false {
						} else {
							r.EncodeUint(uint64(yy39))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrUint"))
				r.WriteMapElemValue()
				if yyn38 {
					r.EncodeNil()
				} else {
					if x.FptrUint == nil {
						r.EncodeNil()
					} else {
						yy41 := *x.FptrUint
						yym42 := z.EncBinary()
						_ = yym42
						if false {
						} else {
							r.EncodeUint(uint64(yy41))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym44 := z.EncBinary()
				_ = yym44
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint8))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FUint8"))
				r.WriteMapElemValue()
				yym45 := z.EncBinary()
				_ = yym45
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint8))
				}
			}
			var yyn46 bool
			if x.FptrUint8 == nil {
				yyn46 = true
				goto LABEL46
			}
		LABEL46:
			if yyr2 || yy2arr2 {
				if yyn46 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrUint8 == nil {
						r.EncodeNil()
					} else {
						yy47 := *x.FptrUint8
						yym48 := z.EncBinary()
						_ = yym48
						if false {
						} else {
							r.EncodeUint(uint64(yy47))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrUint8"))
				r.WriteMapElemValue()
				if yyn46 {
					r.EncodeNil()
				} else {
					if x.FptrUint8 == nil {
						r.EncodeNil()
					} else {
						yy49 := *x.FptrUint8
						yym50 := z.EncBinary()
						_ = yym50
						if false {
						} else {
							r.EncodeUint(uint64(yy49))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym52 := z.EncBinary()
				_ = yym52
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint16))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FUint16"))
				r.WriteMapElemValue()
				yym53 := z.EncBinary()
				_ = yym53
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint16))
				}
			}
			var yyn54 bool
			if x.FptrUint16 == nil {
				yyn54 = true
				goto LABEL54
			}
		LABEL54:
			if yyr2 || yy2arr2 {
				if yyn54 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrUint16 == nil {
						r.EncodeNil()
					} else {
						yy55 := *x.FptrUint16
						yym56 := z.EncBinary()
						_ = yym56
						if false {
						} else {
							r.EncodeUint(uint64(yy55))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrUint16"))
				r.WriteMapElemValue()
				if yyn54 {
					r.EncodeNil()
				} else {
					if x.FptrUint16 == nil {
						r.EncodeNil()
					} else {
						yy57 := *x.FptrUint16
						yym58 := z.EncBinary()
						_ = yym58
						if false {
						} else {
							r.EncodeUint(uint64(yy57))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym60 := z.EncBinary()
				_ = yym60
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint32))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FUint32"))
				r.WriteMapElemValue()
				yym61 := z.EncBinary()
				_ = yym61
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint32))
				}
			}
			var yyn62 bool
			if x.FptrUint32 == nil {
				yyn62 = true
				goto LABEL62
			}
		LABEL62:
			if yyr2 || yy2arr2 {
				if yyn62 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrUint32 == nil {
						r.EncodeNil()
					} else {
						yy63 := *x.FptrUint32
						yym64 := z.EncBinary()
						_ = yym64
						if false {
						} else {
							r.EncodeUint(uint64(yy63))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrUint32"))
				r.WriteMapElemValue()
				if yyn62 {
					r.EncodeNil()
				} else {
					if x.FptrUint32 == nil {
						r.EncodeNil()
					} else {
						yy65 := *x.FptrUint32
						yym66 := z.EncBinary()
						_ = yym66
						if false {
						} else {
							r.EncodeUint(uint64(yy65))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym68 := z.EncBinary()
				_ = yym68
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint64))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FUint64"))
				r.WriteMapElemValue()
				yym69 := z.EncBinary()
				_ = yym69
				if false {
				} else {
					r.EncodeUint(uint64(x.FUint64))
				}
			}
			var yyn70 bool
			if x.FptrUint64 == nil {
				yyn70 = true
				goto LABEL70
			}
		LABEL70:
			if yyr2 || yy2arr2 {
				if yyn70 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrUint64 == nil {
						r.EncodeNil()
					} else {
						yy71 := *x.FptrUint64
						yym72 := z.EncBinary()
						_ = yym72
						if false {
						} else {
							r.EncodeUint(uint64(yy71))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrUint64"))
				r.WriteMapElemValue()
				if yyn70 {
					r.EncodeNil()
				} else {
					if x.FptrUint64 == nil {
						r.EncodeNil()
					} else {
						yy73 := *x.FptrUint64
						yym74 := z.EncBinary()
						_ = yym74
						if false {
						} else {
							r.EncodeUint(uint64(yy73))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym76 := z.EncBinary()
				_ = yym76
				if false {
				} else {
					r.EncodeUint(uint64(x.FUintptr))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FUintptr"))
				r.WriteMapElemValue()
				yym77 := z.EncBinary()
				_ = yym77
				if false {
				} else {
					r.EncodeUint(uint64(x.FUintptr))
				}
			}
			var yyn78 bool
			if x.FptrUintptr == nil {
				yyn78 = true
				goto LABEL78
			}
		LABEL78:
			if yyr2 || yy2arr2 {
				if yyn78 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrUintptr == nil {
						r.EncodeNil()
					} else {
						yy79 := *x.FptrUintptr
						yym80 := z.EncBinary()
						_ = yym80
						if false {
						} else {
							r.EncodeUint(uint64(yy79))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrUintptr"))
				r.WriteMapElemValue()
				if yyn78 {
					r.EncodeNil()
				} else {
					if x.FptrUintptr == nil {
						r.EncodeNil()
					} else {
						yy81 := *x.FptrUintptr
						yym82 := z.EncBinary()
						_ = yym82
						if false {
						} else {
							r.EncodeUint(uint64(yy81))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym84 := z.EncBinary()
				_ = yym84
				if false {
				} else {
					r.EncodeInt(int64(x.FInt))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FInt"))
				r.WriteMapElemValue()
				yym85 := z.EncBinary()
				_ = yym85
				if false {
				} else {
					r.EncodeInt(int64(x.FInt))
				}
			}
			var yyn86 bool
			if x.FptrInt == nil {
				yyn86 = true
				goto LABEL86
			}
		LABEL86:
			if yyr2 || yy2arr2 {
				if yyn86 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrInt == nil {
						r.EncodeNil()
					} else {
						yy87 := *x.FptrInt
						yym88 := z.EncBinary()
						_ = yym88
						if false {
						} else {
							r.EncodeInt(int64(yy87))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrInt"))
				r.WriteMapElemValue()
				if yyn86 {
					r.EncodeNil()
				} else {
					if x.FptrInt == nil {
						r.EncodeNil()
					} else {
						yy89 := *x.FptrInt
						yym90 := z.EncBinary()
						_ = yym90
						if false {
						} else {
							r.EncodeInt(int64(yy89))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym92 := z.EncBinary()
				_ = yym92
				if false {
				} else {
					r.EncodeInt(int64(x.FInt8))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FInt8"))
				r.WriteMapElemValue()
				yym93 := z.EncBinary()
				_ = yym93
				if false {
				} else {
					r.EncodeInt(int64(x.FInt8))
				}
			}
			var yyn94 bool
			if x.FptrInt8 == nil {
				yyn94 = true
				goto LABEL94
			}
		LABEL94:
			if yyr2 || yy2arr2 {
				if yyn94 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrInt8 == nil {
						r.EncodeNil()
					} else {
						yy95 := *x.FptrInt8
						yym96 := z.EncBinary()
						_ = yym96
						if false {
						} else {
							r.EncodeInt(int64(yy95))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrInt8"))
				r.WriteMapElemValue()
				if yyn94 {
					r.EncodeNil()
				} else {
					if x.FptrInt8 == nil {
						r.EncodeNil()
					} else {
						yy97 := *x.FptrInt8
						yym98 := z.EncBinary()
						_ = yym98
						if false {
						} else {
							r.EncodeInt(int64(yy97))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym100 := z.EncBinary()
				_ = yym100
				if false {
				} else {
					r.EncodeInt(int64(x.FInt16))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FInt16"))
				r.WriteMapElemValue()
				yym101 := z.EncBinary()
				_ = yym101
				if false {
				} else {
					r.EncodeInt(int64(x.FInt16))
				}
			}
			var yyn102 bool
			if x.FptrInt16 == nil {
				yyn102 = true
				goto LABEL102
			}
		LABEL102:
			if yyr2 || yy2arr2 {
				if yyn102 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrInt16 == nil {
						r.EncodeNil()
					} else {
						yy103 := *x.FptrInt16
						yym104 := z.EncBinary()
						_ = yym104
						if false {
						} else {
							r.EncodeInt(int64(yy103))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrInt16"))
				r.WriteMapElemValue()
				if yyn102 {
					r.EncodeNil()
				} else {
					if x.FptrInt16 == nil {
						r.EncodeNil()
					} else {
						yy105 := *x.FptrInt16
						yym106 := z.EncBinary()
						_ = yym106
						if false {
						} else {
							r.EncodeInt(int64(yy105))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym108 := z.EncBinary()
				_ = yym108
				if false {
				} else {
					r.EncodeInt(int64(x.FInt32))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FInt32"))
				r.WriteMapElemValue()
				yym109 := z.EncBinary()
				_ = yym109
				if false {
				} else {
					r.EncodeInt(int64(x.FInt32))
				}
			}
			var yyn110 bool
			if x.FptrInt32 == nil {
				yyn110 = true
				goto LABEL110
			}
		LABEL110:
			if yyr2 || yy2arr2 {
				if yyn110 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrInt32 == nil {
						r.EncodeNil()
					} else {
						yy111 := *x.FptrInt32
						yym112 := z.EncBinary()
						_ = yym112
						if false {
						} else {
							r.EncodeInt(int64(yy111))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrInt32"))
				r.WriteMapElemValue()
				if yyn110 {
					r.EncodeNil()
				} else {
					if x.FptrInt32 == nil {
						r.EncodeNil()
					} else {
						yy113 := *x.FptrInt32
						yym114 := z.EncBinary()
						_ = yym114
						if false {
						} else {
							r.EncodeInt(int64(yy113))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym116 := z.EncBinary()
				_ = yym116
				if false {
				} else {
					r.EncodeInt(int64(x.FInt64))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FInt64"))
				r.WriteMapElemValue()
				yym117 := z.EncBinary()
				_ = yym117
				if false {
				} else {
					r.EncodeInt(int64(x.FInt64))
				}
			}
			var yyn118 bool
			if x.FptrInt64 == nil {
				yyn118 = true
				goto LABEL118
			}
		LABEL118:
			if yyr2 || yy2arr2 {
				if yyn118 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrInt64 == nil {
						r.EncodeNil()
					} else {
						yy119 := *x.FptrInt64
						yym120 := z.EncBinary()
						_ = yym120
						if false {
						} else {
							r.EncodeInt(int64(yy119))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrInt64"))
				r.WriteMapElemValue()
				if yyn118 {
					r.EncodeNil()
				} else {
					if x.FptrInt64 == nil {
						r.EncodeNil()
					} else {
						yy121 := *x.FptrInt64
						yym122 := z.EncBinary()
						_ = yym122
						if false {
						} else {
							r.EncodeInt(int64(yy121))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym124 := z.EncBinary()
				_ = yym124
				if false {
				} else {
					r.EncodeBool(bool(x.FBool))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FBool"))
				r.WriteMapElemValue()
				yym125 := z.EncBinary()
				_ = yym125
				if false {
				} else {
					r.EncodeBool(bool(x.FBool))
				}
			}
			var yyn126 bool
			if x.FptrBool == nil {
				yyn126 = true
				goto LABEL126
			}
		LABEL126:
			if yyr2 || yy2arr2 {
				if yyn126 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrBool == nil {
						r.EncodeNil()
					} else {
						yy127 := *x.FptrBool
						yym128 := z.EncBinary()
						_ = yym128
						if false {
						} else {
							r.EncodeBool(bool(yy127))
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrBool"))
				r.WriteMapElemValue()
				if yyn126 {
					r.EncodeNil()
				} else {
					if x.FptrBool == nil {
						r.EncodeNil()
					} else {
						yy129 := *x.FptrBool
						yym130 := z.EncBinary()
						_ = yym130
						if false {
						} else {
							r.EncodeBool(bool(yy129))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceIntf == nil {
					r.EncodeNil()
				} else {
					yym132 := z.EncBinary()
					_ = yym132
					if false {
					} else {
						z.F.EncSliceIntfV(x.FSliceIntf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceIntf"))
				r.WriteMapElemValue()
				if x.FSliceIntf == nil {
					r.EncodeNil()
				} else {
					yym133 := z.EncBinary()
					_ = yym133
					if false {
					} else {
						z.F.EncSliceIntfV(x.FSliceIntf, e)
					}
				}
			}
			var yyn134 bool
			if x.FptrSliceIntf == nil {
				yyn134 = true
				goto LABEL134
			}
		LABEL134:
			if yyr2 || yy2arr2 {
				if yyn134 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceIntf == nil {
						r.EncodeNil()
					} else {
						yy135 := *x.FptrSliceIntf
						yym136 := z.EncBinary()
						_ = yym136
						if false {
						} else {
							z.F.EncSliceIntfV(yy135, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceIntf"))
				r.WriteMapElemValue()
				if yyn134 {
					r.EncodeNil()
				} else {
					if x.FptrSliceIntf == nil {
						r.EncodeNil()
					} else {
						yy137 := *x.FptrSliceIntf
						yym138 := z.EncBinary()
						_ = yym138
						if false {
						} else {
							z.F.EncSliceIntfV(yy137, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceString == nil {
					r.EncodeNil()
				} else {
					yym140 := z.EncBinary()
					_ = yym140
					if false {
					} else {
						z.F.EncSliceStringV(x.FSliceString, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceString"))
				r.WriteMapElemValue()
				if x.FSliceString == nil {
					r.EncodeNil()
				} else {
					yym141 := z.EncBinary()
					_ = yym141
					if false {
					} else {
						z.F.EncSliceStringV(x.FSliceString, e)
					}
				}
			}
			var yyn142 bool
			if x.FptrSliceString == nil {
				yyn142 = true
				goto LABEL142
			}
		LABEL142:
			if yyr2 || yy2arr2 {
				if yyn142 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceString == nil {
						r.EncodeNil()
					} else {
						yy143 := *x.FptrSliceString
						yym144 := z.EncBinary()
						_ = yym144
						if false {
						} else {
							z.F.EncSliceStringV(yy143, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceString"))
				r.WriteMapElemValue()
				if yyn142 {
					r.EncodeNil()
				} else {
					if x.FptrSliceString == nil {
						r.EncodeNil()
					} else {
						yy145 := *x.FptrSliceString
						yym146 := z.EncBinary()
						_ = yym146
						if false {
						} else {
							z.F.EncSliceStringV(yy145, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceFloat32 == nil {
					r.EncodeNil()
				} else {
					yym148 := z.EncBinary()
					_ = yym148
					if false {
					} else {
						z.F.EncSliceFloat32V(x.FSliceFloat32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceFloat32"))
				r.WriteMapElemValue()
				if x.FSliceFloat32 == nil {
					r.EncodeNil()
				} else {
					yym149 := z.EncBinary()
					_ = yym149
					if false {
					} else {
						z.F.EncSliceFloat32V(x.FSliceFloat32, e)
					}
				}
			}
			var yyn150 bool
			if x.FptrSliceFloat32 == nil {
				yyn150 = true
				goto LABEL150
			}
		LABEL150:
			if yyr2 || yy2arr2 {
				if yyn150 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceFloat32 == nil {
						r.EncodeNil()
					} else {
						yy151 := *x.FptrSliceFloat32
						yym152 := z.EncBinary()
						_ = yym152
						if false {
						} else {
							z.F.EncSliceFloat32V(yy151, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceFloat32"))
				r.WriteMapElemValue()
				if yyn150 {
					r.EncodeNil()
				} else {
					if x.FptrSliceFloat32 == nil {
						r.EncodeNil()
					} else {
						yy153 := *x.FptrSliceFloat32
						yym154 := z.EncBinary()
						_ = yym154
						if false {
						} else {
							z.F.EncSliceFloat32V(yy153, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceFloat64 == nil {
					r.EncodeNil()
				} else {
					yym156 := z.EncBinary()
					_ = yym156
					if false {
					} else {
						z.F.EncSliceFloat64V(x.FSliceFloat64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceFloat64"))
				r.WriteMapElemValue()
				if x.FSliceFloat64 == nil {
					r.EncodeNil()
				} else {
					yym157 := z.EncBinary()
					_ = yym157
					if false {
					} else {
						z.F.EncSliceFloat64V(x.FSliceFloat64, e)
					}
				}
			}
			var yyn158 bool
			if x.FptrSliceFloat64 == nil {
				yyn158 = true
				goto LABEL158
			}
		LABEL158:
			if yyr2 || yy2arr2 {
				if yyn158 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceFloat64 == nil {
						r.EncodeNil()
					} else {
						yy159 := *x.FptrSliceFloat64
						yym160 := z.EncBinary()
						_ = yym160
						if false {
						} else {
							z.F.EncSliceFloat64V(yy159, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceFloat64"))
				r.WriteMapElemValue()
				if yyn158 {
					r.EncodeNil()
				} else {
					if x.FptrSliceFloat64 == nil {
						r.EncodeNil()
					} else {
						yy161 := *x.FptrSliceFloat64
						yym162 := z.EncBinary()
						_ = yym162
						if false {
						} else {
							z.F.EncSliceFloat64V(yy161, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceUint == nil {
					r.EncodeNil()
				} else {
					yym164 := z.EncBinary()
					_ = yym164
					if false {
					} else {
						z.F.EncSliceUintV(x.FSliceUint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceUint"))
				r.WriteMapElemValue()
				if x.FSliceUint == nil {
					r.EncodeNil()
				} else {
					yym165 := z.EncBinary()
					_ = yym165
					if false {
					} else {
						z.F.EncSliceUintV(x.FSliceUint, e)
					}
				}
			}
			var yyn166 bool
			if x.FptrSliceUint == nil {
				yyn166 = true
				goto LABEL166
			}
		LABEL166:
			if yyr2 || yy2arr2 {
				if yyn166 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceUint == nil {
						r.EncodeNil()
					} else {
						yy167 := *x.FptrSliceUint
						yym168 := z.EncBinary()
						_ = yym168
						if false {
						} else {
							z.F.EncSliceUintV(yy167, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceUint"))
				r.WriteMapElemValue()
				if yyn166 {
					r.EncodeNil()
				} else {
					if x.FptrSliceUint == nil {
						r.EncodeNil()
					} else {
						yy169 := *x.FptrSliceUint
						yym170 := z.EncBinary()
						_ = yym170
						if false {
						} else {
							z.F.EncSliceUintV(yy169, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceUint16 == nil {
					r.EncodeNil()
				} else {
					yym172 := z.EncBinary()
					_ = yym172
					if false {
					} else {
						z.F.EncSliceUint16V(x.FSliceUint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceUint16"))
				r.WriteMapElemValue()
				if x.FSliceUint16 == nil {
					r.EncodeNil()
				} else {
					yym173 := z.EncBinary()
					_ = yym173
					if false {
					} else {
						z.F.EncSliceUint16V(x.FSliceUint16, e)
					}
				}
			}
			var yyn174 bool
			if x.FptrSliceUint16 == nil {
				yyn174 = true
				goto LABEL174
			}
		LABEL174:
			if yyr2 || yy2arr2 {
				if yyn174 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceUint16 == nil {
						r.EncodeNil()
					} else {
						yy175 := *x.FptrSliceUint16
						yym176 := z.EncBinary()
						_ = yym176
						if false {
						} else {
							z.F.EncSliceUint16V(yy175, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceUint16"))
				r.WriteMapElemValue()
				if yyn174 {
					r.EncodeNil()
				} else {
					if x.FptrSliceUint16 == nil {
						r.EncodeNil()
					} else {
						yy177 := *x.FptrSliceUint16
						yym178 := z.EncBinary()
						_ = yym178
						if false {
						} else {
							z.F.EncSliceUint16V(yy177, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceUint32 == nil {
					r.EncodeNil()
				} else {
					yym180 := z.EncBinary()
					_ = yym180
					if false {
					} else {
						z.F.EncSliceUint32V(x.FSliceUint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceUint32"))
				r.WriteMapElemValue()
				if x.FSliceUint32 == nil {
					r.EncodeNil()
				} else {
					yym181 := z.EncBinary()
					_ = yym181
					if false {
					} else {
						z.F.EncSliceUint32V(x.FSliceUint32, e)
					}
				}
			}
			var yyn182 bool
			if x.FptrSliceUint32 == nil {
				yyn182 = true
				goto LABEL182
			}
		LABEL182:
			if yyr2 || yy2arr2 {
				if yyn182 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceUint32 == nil {
						r.EncodeNil()
					} else {
						yy183 := *x.FptrSliceUint32
						yym184 := z.EncBinary()
						_ = yym184
						if false {
						} else {
							z.F.EncSliceUint32V(yy183, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceUint32"))
				r.WriteMapElemValue()
				if yyn182 {
					r.EncodeNil()
				} else {
					if x.FptrSliceUint32 == nil {
						r.EncodeNil()
					} else {
						yy185 := *x.FptrSliceUint32
						yym186 := z.EncBinary()
						_ = yym186
						if false {
						} else {
							z.F.EncSliceUint32V(yy185, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceUint64 == nil {
					r.EncodeNil()
				} else {
					yym188 := z.EncBinary()
					_ = yym188
					if false {
					} else {
						z.F.EncSliceUint64V(x.FSliceUint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceUint64"))
				r.WriteMapElemValue()
				if x.FSliceUint64 == nil {
					r.EncodeNil()
				} else {
					yym189 := z.EncBinary()
					_ = yym189
					if false {
					} else {
						z.F.EncSliceUint64V(x.FSliceUint64, e)
					}
				}
			}
			var yyn190 bool
			if x.FptrSliceUint64 == nil {
				yyn190 = true
				goto LABEL190
			}
		LABEL190:
			if yyr2 || yy2arr2 {
				if yyn190 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceUint64 == nil {
						r.EncodeNil()
					} else {
						yy191 := *x.FptrSliceUint64
						yym192 := z.EncBinary()
						_ = yym192
						if false {
						} else {
							z.F.EncSliceUint64V(yy191, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceUint64"))
				r.WriteMapElemValue()
				if yyn190 {
					r.EncodeNil()
				} else {
					if x.FptrSliceUint64 == nil {
						r.EncodeNil()
					} else {
						yy193 := *x.FptrSliceUint64
						yym194 := z.EncBinary()
						_ = yym194
						if false {
						} else {
							z.F.EncSliceUint64V(yy193, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceUintptr == nil {
					r.EncodeNil()
				} else {
					yym196 := z.EncBinary()
					_ = yym196
					if false {
					} else {
						z.F.EncSliceUintptrV(x.FSliceUintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceUintptr"))
				r.WriteMapElemValue()
				if x.FSliceUintptr == nil {
					r.EncodeNil()
				} else {
					yym197 := z.EncBinary()
					_ = yym197
					if false {
					} else {
						z.F.EncSliceUintptrV(x.FSliceUintptr, e)
					}
				}
			}
			var yyn198 bool
			if x.FptrSliceUintptr == nil {
				yyn198 = true
				goto LABEL198
			}
		LABEL198:
			if yyr2 || yy2arr2 {
				if yyn198 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceUintptr == nil {
						r.EncodeNil()
					} else {
						yy199 := *x.FptrSliceUintptr
						yym200 := z.EncBinary()
						_ = yym200
						if false {
						} else {
							z.F.EncSliceUintptrV(yy199, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceUintptr"))
				r.WriteMapElemValue()
				if yyn198 {
					r.EncodeNil()
				} else {
					if x.FptrSliceUintptr == nil {
						r.EncodeNil()
					} else {
						yy201 := *x.FptrSliceUintptr
						yym202 := z.EncBinary()
						_ = yym202
						if false {
						} else {
							z.F.EncSliceUintptrV(yy201, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceInt == nil {
					r.EncodeNil()
				} else {
					yym204 := z.EncBinary()
					_ = yym204
					if false {
					} else {
						z.F.EncSliceIntV(x.FSliceInt, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceInt"))
				r.WriteMapElemValue()
				if x.FSliceInt == nil {
					r.EncodeNil()
				} else {
					yym205 := z.EncBinary()
					_ = yym205
					if false {
					} else {
						z.F.EncSliceIntV(x.FSliceInt, e)
					}
				}
			}
			var yyn206 bool
			if x.FptrSliceInt == nil {
				yyn206 = true
				goto LABEL206
			}
		LABEL206:
			if yyr2 || yy2arr2 {
				if yyn206 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceInt == nil {
						r.EncodeNil()
					} else {
						yy207 := *x.FptrSliceInt
						yym208 := z.EncBinary()
						_ = yym208
						if false {
						} else {
							z.F.EncSliceIntV(yy207, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceInt"))
				r.WriteMapElemValue()
				if yyn206 {
					r.EncodeNil()
				} else {
					if x.FptrSliceInt == nil {
						r.EncodeNil()
					} else {
						yy209 := *x.FptrSliceInt
						yym210 := z.EncBinary()
						_ = yym210
						if false {
						} else {
							z.F.EncSliceIntV(yy209, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceInt8 == nil {
					r.EncodeNil()
				} else {
					yym212 := z.EncBinary()
					_ = yym212
					if false {
					} else {
						z.F.EncSliceInt8V(x.FSliceInt8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceInt8"))
				r.WriteMapElemValue()
				if x.FSliceInt8 == nil {
					r.EncodeNil()
				} else {
					yym213 := z.EncBinary()
					_ = yym213
					if false {
					} else {
						z.F.EncSliceInt8V(x.FSliceInt8, e)
					}
				}
			}
			var yyn214 bool
			if x.FptrSliceInt8 == nil {
				yyn214 = true
				goto LABEL214
			}
		LABEL214:
			if yyr2 || yy2arr2 {
				if yyn214 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceInt8 == nil {
						r.EncodeNil()
					} else {
						yy215 := *x.FptrSliceInt8
						yym216 := z.EncBinary()
						_ = yym216
						if false {
						} else {
							z.F.EncSliceInt8V(yy215, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceInt8"))
				r.WriteMapElemValue()
				if yyn214 {
					r.EncodeNil()
				} else {
					if x.FptrSliceInt8 == nil {
						r.EncodeNil()
					} else {
						yy217 := *x.FptrSliceInt8
						yym218 := z.EncBinary()
						_ = yym218
						if false {
						} else {
							z.F.EncSliceInt8V(yy217, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceInt16 == nil {
					r.EncodeNil()
				} else {
					yym220 := z.EncBinary()
					_ = yym220
					if false {
					} else {
						z.F.EncSliceInt16V(x.FSliceInt16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceInt16"))
				r.WriteMapElemValue()
				if x.FSliceInt16 == nil {
					r.EncodeNil()
				} else {
					yym221 := z.EncBinary()
					_ = yym221
					if false {
					} else {
						z.F.EncSliceInt16V(x.FSliceInt16, e)
					}
				}
			}
			var yyn222 bool
			if x.FptrSliceInt16 == nil {
				yyn222 = true
				goto LABEL222
			}
		LABEL222:
			if yyr2 || yy2arr2 {
				if yyn222 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceInt16 == nil {
						r.EncodeNil()
					} else {
						yy223 := *x.FptrSliceInt16
						yym224 := z.EncBinary()
						_ = yym224
						if false {
						} else {
							z.F.EncSliceInt16V(yy223, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceInt16"))
				r.WriteMapElemValue()
				if yyn222 {
					r.EncodeNil()
				} else {
					if x.FptrSliceInt16 == nil {
						r.EncodeNil()
					} else {
						yy225 := *x.FptrSliceInt16
						yym226 := z.EncBinary()
						_ = yym226
						if false {
						} else {
							z.F.EncSliceInt16V(yy225, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceInt32 == nil {
					r.EncodeNil()
				} else {
					yym228 := z.EncBinary()
					_ = yym228
					if false {
					} else {
						z.F.EncSliceInt32V(x.FSliceInt32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceInt32"))
				r.WriteMapElemValue()
				if x.FSliceInt32 == nil {
					r.EncodeNil()
				} else {
					yym229 := z.EncBinary()
					_ = yym229
					if false {
					} else {
						z.F.EncSliceInt32V(x.FSliceInt32, e)
					}
				}
			}
			var yyn230 bool
			if x.FptrSliceInt32 == nil {
				yyn230 = true
				goto LABEL230
			}
		LABEL230:
			if yyr2 || yy2arr2 {
				if yyn230 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceInt32 == nil {
						r.EncodeNil()
					} else {
						yy231 := *x.FptrSliceInt32
						yym232 := z.EncBinary()
						_ = yym232
						if false {
						} else {
							z.F.EncSliceInt32V(yy231, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceInt32"))
				r.WriteMapElemValue()
				if yyn230 {
					r.EncodeNil()
				} else {
					if x.FptrSliceInt32 == nil {
						r.EncodeNil()
					} else {
						yy233 := *x.FptrSliceInt32
						yym234 := z.EncBinary()
						_ = yym234
						if false {
						} else {
							z.F.EncSliceInt32V(yy233, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceInt64 == nil {
					r.EncodeNil()
				} else {
					yym236 := z.EncBinary()
					_ = yym236
					if false {
					} else {
						z.F.EncSliceInt64V(x.FSliceInt64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceInt64"))
				r.WriteMapElemValue()
				if x.FSliceInt64 == nil {
					r.EncodeNil()
				} else {
					yym237 := z.EncBinary()
					_ = yym237
					if false {
					} else {
						z.F.EncSliceInt64V(x.FSliceInt64, e)
					}
				}
			}
			var yyn238 bool
			if x.FptrSliceInt64 == nil {
				yyn238 = true
				goto LABEL238
			}
		LABEL238:
			if yyr2 || yy2arr2 {
				if yyn238 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceInt64 == nil {
						r.EncodeNil()
					} else {
						yy239 := *x.FptrSliceInt64
						yym240 := z.EncBinary()
						_ = yym240
						if false {
						} else {
							z.F.EncSliceInt64V(yy239, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceInt64"))
				r.WriteMapElemValue()
				if yyn238 {
					r.EncodeNil()
				} else {
					if x.FptrSliceInt64 == nil {
						r.EncodeNil()
					} else {
						yy241 := *x.FptrSliceInt64
						yym242 := z.EncBinary()
						_ = yym242
						if false {
						} else {
							z.F.EncSliceInt64V(yy241, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FSliceBool == nil {
					r.EncodeNil()
				} else {
					yym244 := z.EncBinary()
					_ = yym244
					if false {
					} else {
						z.F.EncSliceBoolV(x.FSliceBool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FSliceBool"))
				r.WriteMapElemValue()
				if x.FSliceBool == nil {
					r.EncodeNil()
				} else {
					yym245 := z.EncBinary()
					_ = yym245
					if false {
					} else {
						z.F.EncSliceBoolV(x.FSliceBool, e)
					}
				}
			}
			var yyn246 bool
			if x.FptrSliceBool == nil {
				yyn246 = true
				goto LABEL246
			}
		LABEL246:
			if yyr2 || yy2arr2 {
				if yyn246 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrSliceBool == nil {
						r.EncodeNil()
					} else {
						yy247 := *x.FptrSliceBool
						yym248 := z.EncBinary()
						_ = yym248
						if false {
						} else {
							z.F.EncSliceBoolV(yy247, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrSliceBool"))
				r.WriteMapElemValue()
				if yyn246 {
					r.EncodeNil()
				} else {
					if x.FptrSliceBool == nil {
						r.EncodeNil()
					} else {
						yy249 := *x.FptrSliceBool
						yym250 := z.EncBinary()
						_ = yym250
						if false {
						} else {
							z.F.EncSliceBoolV(yy249, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfIntf == nil {
					r.EncodeNil()
				} else {
					yym252 := z.EncBinary()
					_ = yym252
					if false {
					} else {
						z.F.EncMapIntfIntfV(x.FMapIntfIntf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfIntf"))
				r.WriteMapElemValue()
				if x.FMapIntfIntf == nil {
					r.EncodeNil()
				} else {
					yym253 := z.EncBinary()
					_ = yym253
					if false {
					} else {
						z.F.EncMapIntfIntfV(x.FMapIntfIntf, e)
					}
				}
			}
			var yyn254 bool
			if x.FptrMapIntfIntf == nil {
				yyn254 = true
				goto LABEL254
			}
		LABEL254:
			if yyr2 || yy2arr2 {
				if yyn254 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfIntf == nil {
						r.EncodeNil()
					} else {
						yy255 := *x.FptrMapIntfIntf
						yym256 := z.EncBinary()
						_ = yym256
						if false {
						} else {
							z.F.EncMapIntfIntfV(yy255, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfIntf"))
				r.WriteMapElemValue()
				if yyn254 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfIntf == nil {
						r.EncodeNil()
					} else {
						yy257 := *x.FptrMapIntfIntf
						yym258 := z.EncBinary()
						_ = yym258
						if false {
						} else {
							z.F.EncMapIntfIntfV(yy257, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfString == nil {
					r.EncodeNil()
				} else {
					yym260 := z.EncBinary()
					_ = yym260
					if false {
					} else {
						z.F.EncMapIntfStringV(x.FMapIntfString, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfString"))
				r.WriteMapElemValue()
				if x.FMapIntfString == nil {
					r.EncodeNil()
				} else {
					yym261 := z.EncBinary()
					_ = yym261
					if false {
					} else {
						z.F.EncMapIntfStringV(x.FMapIntfString, e)
					}
				}
			}
			var yyn262 bool
			if x.FptrMapIntfString == nil {
				yyn262 = true
				goto LABEL262
			}
		LABEL262:
			if yyr2 || yy2arr2 {
				if yyn262 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfString == nil {
						r.EncodeNil()
					} else {
						yy263 := *x.FptrMapIntfString
						yym264 := z.EncBinary()
						_ = yym264
						if false {
						} else {
							z.F.EncMapIntfStringV(yy263, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfString"))
				r.WriteMapElemValue()
				if yyn262 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfString == nil {
						r.EncodeNil()
					} else {
						yy265 := *x.FptrMapIntfString
						yym266 := z.EncBinary()
						_ = yym266
						if false {
						} else {
							z.F.EncMapIntfStringV(yy265, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfUint == nil {
					r.EncodeNil()
				} else {
					yym268 := z.EncBinary()
					_ = yym268
					if false {
					} else {
						z.F.EncMapIntfUintV(x.FMapIntfUint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfUint"))
				r.WriteMapElemValue()
				if x.FMapIntfUint == nil {
					r.EncodeNil()
				} else {
					yym269 := z.EncBinary()
					_ = yym269
					if false {
					} else {
						z.F.EncMapIntfUintV(x.FMapIntfUint, e)
					}
				}
			}
			var yyn270 bool
			if x.FptrMapIntfUint == nil {
				yyn270 = true
				goto LABEL270
			}
		LABEL270:
			if yyr2 || yy2arr2 {
				if yyn270 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfUint == nil {
						r.EncodeNil()
					} else {
						yy271 := *x.FptrMapIntfUint
						yym272 := z.EncBinary()
						_ = yym272
						if false {
						} else {
							z.F.EncMapIntfUintV(yy271, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfUint"))
				r.WriteMapElemValue()
				if yyn270 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfUint == nil {
						r.EncodeNil()
					} else {
						yy273 := *x.FptrMapIntfUint
						yym274 := z.EncBinary()
						_ = yym274
						if false {
						} else {
							z.F.EncMapIntfUintV(yy273, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfUint8 == nil {
					r.EncodeNil()
				} else {
					yym276 := z.EncBinary()
					_ = yym276
					if false {
					} else {
						z.F.EncMapIntfUint8V(x.FMapIntfUint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfUint8"))
				r.WriteMapElemValue()
				if x.FMapIntfUint8 == nil {
					r.EncodeNil()
				} else {
					yym277 := z.EncBinary()
					_ = yym277
					if false {
					} else {
						z.F.EncMapIntfUint8V(x.FMapIntfUint8, e)
					}
				}
			}
			var yyn278 bool
			if x.FptrMapIntfUint8 == nil {
				yyn278 = true
				goto LABEL278
			}
		LABEL278:
			if yyr2 || yy2arr2 {
				if yyn278 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfUint8 == nil {
						r.EncodeNil()
					} else {
						yy279 := *x.FptrMapIntfUint8
						yym280 := z.EncBinary()
						_ = yym280
						if false {
						} else {
							z.F.EncMapIntfUint8V(yy279, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfUint8"))
				r.WriteMapElemValue()
				if yyn278 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfUint8 == nil {
						r.EncodeNil()
					} else {
						yy281 := *x.FptrMapIntfUint8
						yym282 := z.EncBinary()
						_ = yym282
						if false {
						} else {
							z.F.EncMapIntfUint8V(yy281, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfUint16 == nil {
					r.EncodeNil()
				} else {
					yym284 := z.EncBinary()
					_ = yym284
					if false {
					} else {
						z.F.EncMapIntfUint16V(x.FMapIntfUint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfUint16"))
				r.WriteMapElemValue()
				if x.FMapIntfUint16 == nil {
					r.EncodeNil()
				} else {
					yym285 := z.EncBinary()
					_ = yym285
					if false {
					} else {
						z.F.EncMapIntfUint16V(x.FMapIntfUint16, e)
					}
				}
			}
			var yyn286 bool
			if x.FptrMapIntfUint16 == nil {
				yyn286 = true
				goto LABEL286
			}
		LABEL286:
			if yyr2 || yy2arr2 {
				if yyn286 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfUint16 == nil {
						r.EncodeNil()
					} else {
						yy287 := *x.FptrMapIntfUint16
						yym288 := z.EncBinary()
						_ = yym288
						if false {
						} else {
							z.F.EncMapIntfUint16V(yy287, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfUint16"))
				r.WriteMapElemValue()
				if yyn286 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfUint16 == nil {
						r.EncodeNil()
					} else {
						yy289 := *x.FptrMapIntfUint16
						yym290 := z.EncBinary()
						_ = yym290
						if false {
						} else {
							z.F.EncMapIntfUint16V(yy289, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfUint32 == nil {
					r.EncodeNil()
				} else {
					yym292 := z.EncBinary()
					_ = yym292
					if false {
					} else {
						z.F.EncMapIntfUint32V(x.FMapIntfUint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfUint32"))
				r.WriteMapElemValue()
				if x.FMapIntfUint32 == nil {
					r.EncodeNil()
				} else {
					yym293 := z.EncBinary()
					_ = yym293
					if false {
					} else {
						z.F.EncMapIntfUint32V(x.FMapIntfUint32, e)
					}
				}
			}
			var yyn294 bool
			if x.FptrMapIntfUint32 == nil {
				yyn294 = true
				goto LABEL294
			}
		LABEL294:
			if yyr2 || yy2arr2 {
				if yyn294 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfUint32 == nil {
						r.EncodeNil()
					} else {
						yy295 := *x.FptrMapIntfUint32
						yym296 := z.EncBinary()
						_ = yym296
						if false {
						} else {
							z.F.EncMapIntfUint32V(yy295, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfUint32"))
				r.WriteMapElemValue()
				if yyn294 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfUint32 == nil {
						r.EncodeNil()
					} else {
						yy297 := *x.FptrMapIntfUint32
						yym298 := z.EncBinary()
						_ = yym298
						if false {
						} else {
							z.F.EncMapIntfUint32V(yy297, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfUint64 == nil {
					r.EncodeNil()
				} else {
					yym300 := z.EncBinary()
					_ = yym300
					if false {
					} else {
						z.F.EncMapIntfUint64V(x.FMapIntfUint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfUint64"))
				r.WriteMapElemValue()
				if x.FMapIntfUint64 == nil {
					r.EncodeNil()
				} else {
					yym301 := z.EncBinary()
					_ = yym301
					if false {
					} else {
						z.F.EncMapIntfUint64V(x.FMapIntfUint64, e)
					}
				}
			}
			var yyn302 bool
			if x.FptrMapIntfUint64 == nil {
				yyn302 = true
				goto LABEL302
			}
		LABEL302:
			if yyr2 || yy2arr2 {
				if yyn302 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfUint64 == nil {
						r.EncodeNil()
					} else {
						yy303 := *x.FptrMapIntfUint64
						yym304 := z.EncBinary()
						_ = yym304
						if false {
						} else {
							z.F.EncMapIntfUint64V(yy303, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfUint64"))
				r.WriteMapElemValue()
				if yyn302 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfUint64 == nil {
						r.EncodeNil()
					} else {
						yy305 := *x.FptrMapIntfUint64
						yym306 := z.EncBinary()
						_ = yym306
						if false {
						} else {
							z.F.EncMapIntfUint64V(yy305, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfUintptr == nil {
					r.EncodeNil()
				} else {
					yym308 := z.EncBinary()
					_ = yym308
					if false {
					} else {
						z.F.EncMapIntfUintptrV(x.FMapIntfUintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfUintptr"))
				r.WriteMapElemValue()
				if x.FMapIntfUintptr == nil {
					r.EncodeNil()
				} else {
					yym309 := z.EncBinary()
					_ = yym309
					if false {
					} else {
						z.F.EncMapIntfUintptrV(x.FMapIntfUintptr, e)
					}
				}
			}
			var yyn310 bool
			if x.FptrMapIntfUintptr == nil {
				yyn310 = true
				goto LABEL310
			}
		LABEL310:
			if yyr2 || yy2arr2 {
				if yyn310 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfUintptr == nil {
						r.EncodeNil()
					} else {
						yy311 := *x.FptrMapIntfUintptr
						yym312 := z.EncBinary()
						_ = yym312
						if false {
						} else {
							z.F.EncMapIntfUintptrV(yy311, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfUintptr"))
				r.WriteMapElemValue()
				if yyn310 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfUintptr == nil {
						r.EncodeNil()
					} else {
						yy313 := *x.FptrMapIntfUintptr
						yym314 := z.EncBinary()
						_ = yym314
						if false {
						} else {
							z.F.EncMapIntfUintptrV(yy313, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfInt == nil {
					r.EncodeNil()
				} else {
					yym316 := z.EncBinary()
					_ = yym316
					if false {
					} else {
						z.F.EncMapIntfIntV(x.FMapIntfInt, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfInt"))
				r.WriteMapElemValue()
				if x.FMapIntfInt == nil {
					r.EncodeNil()
				} else {
					yym317 := z.EncBinary()
					_ = yym317
					if false {
					} else {
						z.F.EncMapIntfIntV(x.FMapIntfInt, e)
					}
				}
			}
			var yyn318 bool
			if x.FptrMapIntfInt == nil {
				yyn318 = true
				goto LABEL318
			}
		LABEL318:
			if yyr2 || yy2arr2 {
				if yyn318 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfInt == nil {
						r.EncodeNil()
					} else {
						yy319 := *x.FptrMapIntfInt
						yym320 := z.EncBinary()
						_ = yym320
						if false {
						} else {
							z.F.EncMapIntfIntV(yy319, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfInt"))
				r.WriteMapElemValue()
				if yyn318 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfInt == nil {
						r.EncodeNil()
					} else {
						yy321 := *x.FptrMapIntfInt
						yym322 := z.EncBinary()
						_ = yym322
						if false {
						} else {
							z.F.EncMapIntfIntV(yy321, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfInt8 == nil {
					r.EncodeNil()
				} else {
					yym324 := z.EncBinary()
					_ = yym324
					if false {
					} else {
						z.F.EncMapIntfInt8V(x.FMapIntfInt8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfInt8"))
				r.WriteMapElemValue()
				if x.FMapIntfInt8 == nil {
					r.EncodeNil()
				} else {
					yym325 := z.EncBinary()
					_ = yym325
					if false {
					} else {
						z.F.EncMapIntfInt8V(x.FMapIntfInt8, e)
					}
				}
			}
			var yyn326 bool
			if x.FptrMapIntfInt8 == nil {
				yyn326 = true
				goto LABEL326
			}
		LABEL326:
			if yyr2 || yy2arr2 {
				if yyn326 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfInt8 == nil {
						r.EncodeNil()
					} else {
						yy327 := *x.FptrMapIntfInt8
						yym328 := z.EncBinary()
						_ = yym328
						if false {
						} else {
							z.F.EncMapIntfInt8V(yy327, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfInt8"))
				r.WriteMapElemValue()
				if yyn326 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfInt8 == nil {
						r.EncodeNil()
					} else {
						yy329 := *x.FptrMapIntfInt8
						yym330 := z.EncBinary()
						_ = yym330
						if false {
						} else {
							z.F.EncMapIntfInt8V(yy329, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfInt16 == nil {
					r.EncodeNil()
				} else {
					yym332 := z.EncBinary()
					_ = yym332
					if false {
					} else {
						z.F.EncMapIntfInt16V(x.FMapIntfInt16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfInt16"))
				r.WriteMapElemValue()
				if x.FMapIntfInt16 == nil {
					r.EncodeNil()
				} else {
					yym333 := z.EncBinary()
					_ = yym333
					if false {
					} else {
						z.F.EncMapIntfInt16V(x.FMapIntfInt16, e)
					}
				}
			}
			var yyn334 bool
			if x.FptrMapIntfInt16 == nil {
				yyn334 = true
				goto LABEL334
			}
		LABEL334:
			if yyr2 || yy2arr2 {
				if yyn334 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfInt16 == nil {
						r.EncodeNil()
					} else {
						yy335 := *x.FptrMapIntfInt16
						yym336 := z.EncBinary()
						_ = yym336
						if false {
						} else {
							z.F.EncMapIntfInt16V(yy335, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfInt16"))
				r.WriteMapElemValue()
				if yyn334 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfInt16 == nil {
						r.EncodeNil()
					} else {
						yy337 := *x.FptrMapIntfInt16
						yym338 := z.EncBinary()
						_ = yym338
						if false {
						} else {
							z.F.EncMapIntfInt16V(yy337, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfInt32 == nil {
					r.EncodeNil()
				} else {
					yym340 := z.EncBinary()
					_ = yym340
					if false {
					} else {
						z.F.EncMapIntfInt32V(x.FMapIntfInt32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfInt32"))
				r.WriteMapElemValue()
				if x.FMapIntfInt32 == nil {
					r.EncodeNil()
				} else {
					yym341 := z.EncBinary()
					_ = yym341
					if false {
					} else {
						z.F.EncMapIntfInt32V(x.FMapIntfInt32, e)
					}
				}
			}
			var yyn342 bool
			if x.FptrMapIntfInt32 == nil {
				yyn342 = true
				goto LABEL342
			}
		LABEL342:
			if yyr2 || yy2arr2 {
				if yyn342 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfInt32 == nil {
						r.EncodeNil()
					} else {
						yy343 := *x.FptrMapIntfInt32
						yym344 := z.EncBinary()
						_ = yym344
						if false {
						} else {
							z.F.EncMapIntfInt32V(yy343, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfInt32"))
				r.WriteMapElemValue()
				if yyn342 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfInt32 == nil {
						r.EncodeNil()
					} else {
						yy345 := *x.FptrMapIntfInt32
						yym346 := z.EncBinary()
						_ = yym346
						if false {
						} else {
							z.F.EncMapIntfInt32V(yy345, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfInt64 == nil {
					r.EncodeNil()
				} else {
					yym348 := z.EncBinary()
					_ = yym348
					if false {
					} else {
						z.F.EncMapIntfInt64V(x.FMapIntfInt64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfInt64"))
				r.WriteMapElemValue()
				if x.FMapIntfInt64 == nil {
					r.EncodeNil()
				} else {
					yym349 := z.EncBinary()
					_ = yym349
					if false {
					} else {
						z.F.EncMapIntfInt64V(x.FMapIntfInt64, e)
					}
				}
			}
			var yyn350 bool
			if x.FptrMapIntfInt64 == nil {
				yyn350 = true
				goto LABEL350
			}
		LABEL350:
			if yyr2 || yy2arr2 {
				if yyn350 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfInt64 == nil {
						r.EncodeNil()
					} else {
						yy351 := *x.FptrMapIntfInt64
						yym352 := z.EncBinary()
						_ = yym352
						if false {
						} else {
							z.F.EncMapIntfInt64V(yy351, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfInt64"))
				r.WriteMapElemValue()
				if yyn350 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfInt64 == nil {
						r.EncodeNil()
					} else {
						yy353 := *x.FptrMapIntfInt64
						yym354 := z.EncBinary()
						_ = yym354
						if false {
						} else {
							z.F.EncMapIntfInt64V(yy353, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfFloat32 == nil {
					r.EncodeNil()
				} else {
					yym356 := z.EncBinary()
					_ = yym356
					if false {
					} else {
						z.F.EncMapIntfFloat32V(x.FMapIntfFloat32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfFloat32"))
				r.WriteMapElemValue()
				if x.FMapIntfFloat32 == nil {
					r.EncodeNil()
				} else {
					yym357 := z.EncBinary()
					_ = yym357
					if false {
					} else {
						z.F.EncMapIntfFloat32V(x.FMapIntfFloat32, e)
					}
				}
			}
			var yyn358 bool
			if x.FptrMapIntfFloat32 == nil {
				yyn358 = true
				goto LABEL358
			}
		LABEL358:
			if yyr2 || yy2arr2 {
				if yyn358 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfFloat32 == nil {
						r.EncodeNil()
					} else {
						yy359 := *x.FptrMapIntfFloat32
						yym360 := z.EncBinary()
						_ = yym360
						if false {
						} else {
							z.F.EncMapIntfFloat32V(yy359, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfFloat32"))
				r.WriteMapElemValue()
				if yyn358 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfFloat32 == nil {
						r.EncodeNil()
					} else {
						yy361 := *x.FptrMapIntfFloat32
						yym362 := z.EncBinary()
						_ = yym362
						if false {
						} else {
							z.F.EncMapIntfFloat32V(yy361, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfFloat64 == nil {
					r.EncodeNil()
				} else {
					yym364 := z.EncBinary()
					_ = yym364
					if false {
					} else {
						z.F.EncMapIntfFloat64V(x.FMapIntfFloat64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfFloat64"))
				r.WriteMapElemValue()
				if x.FMapIntfFloat64 == nil {
					r.EncodeNil()
				} else {
					yym365 := z.EncBinary()
					_ = yym365
					if false {
					} else {
						z.F.EncMapIntfFloat64V(x.FMapIntfFloat64, e)
					}
				}
			}
			var yyn366 bool
			if x.FptrMapIntfFloat64 == nil {
				yyn366 = true
				goto LABEL366
			}
		LABEL366:
			if yyr2 || yy2arr2 {
				if yyn366 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfFloat64 == nil {
						r.EncodeNil()
					} else {
						yy367 := *x.FptrMapIntfFloat64
						yym368 := z.EncBinary()
						_ = yym368
						if false {
						} else {
							z.F.EncMapIntfFloat64V(yy367, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfFloat64"))
				r.WriteMapElemValue()
				if yyn366 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfFloat64 == nil {
						r.EncodeNil()
					} else {
						yy369 := *x.FptrMapIntfFloat64
						yym370 := z.EncBinary()
						_ = yym370
						if false {
						} else {
							z.F.EncMapIntfFloat64V(yy369, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntfBool == nil {
					r.EncodeNil()
				} else {
					yym372 := z.EncBinary()
					_ = yym372
					if false {
					} else {
						z.F.EncMapIntfBoolV(x.FMapIntfBool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntfBool"))
				r.WriteMapElemValue()
				if x.FMapIntfBool == nil {
					r.EncodeNil()
				} else {
					yym373 := z.EncBinary()
					_ = yym373
					if false {
					} else {
						z.F.EncMapIntfBoolV(x.FMapIntfBool, e)
					}
				}
			}
			var yyn374 bool
			if x.FptrMapIntfBool == nil {
				yyn374 = true
				goto LABEL374
			}
		LABEL374:
			if yyr2 || yy2arr2 {
				if yyn374 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntfBool == nil {
						r.EncodeNil()
					} else {
						yy375 := *x.FptrMapIntfBool
						yym376 := z.EncBinary()
						_ = yym376
						if false {
						} else {
							z.F.EncMapIntfBoolV(yy375, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntfBool"))
				r.WriteMapElemValue()
				if yyn374 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntfBool == nil {
						r.EncodeNil()
					} else {
						yy377 := *x.FptrMapIntfBool
						yym378 := z.EncBinary()
						_ = yym378
						if false {
						} else {
							z.F.EncMapIntfBoolV(yy377, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringIntf == nil {
					r.EncodeNil()
				} else {
					yym380 := z.EncBinary()
					_ = yym380
					if false {
					} else {
						z.F.EncMapStringIntfV(x.FMapStringIntf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringIntf"))
				r.WriteMapElemValue()
				if x.FMapStringIntf == nil {
					r.EncodeNil()
				} else {
					yym381 := z.EncBinary()
					_ = yym381
					if false {
					} else {
						z.F.EncMapStringIntfV(x.FMapStringIntf, e)
					}
				}
			}
			var yyn382 bool
			if x.FptrMapStringIntf == nil {
				yyn382 = true
				goto LABEL382
			}
		LABEL382:
			if yyr2 || yy2arr2 {
				if yyn382 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringIntf == nil {
						r.EncodeNil()
					} else {
						yy383 := *x.FptrMapStringIntf
						yym384 := z.EncBinary()
						_ = yym384
						if false {
						} else {
							z.F.EncMapStringIntfV(yy383, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringIntf"))
				r.WriteMapElemValue()
				if yyn382 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringIntf == nil {
						r.EncodeNil()
					} else {
						yy385 := *x.FptrMapStringIntf
						yym386 := z.EncBinary()
						_ = yym386
						if false {
						} else {
							z.F.EncMapStringIntfV(yy385, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringString == nil {
					r.EncodeNil()
				} else {
					yym388 := z.EncBinary()
					_ = yym388
					if false {
					} else {
						z.F.EncMapStringStringV(x.FMapStringString, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringString"))
				r.WriteMapElemValue()
				if x.FMapStringString == nil {
					r.EncodeNil()
				} else {
					yym389 := z.EncBinary()
					_ = yym389
					if false {
					} else {
						z.F.EncMapStringStringV(x.FMapStringString, e)
					}
				}
			}
			var yyn390 bool
			if x.FptrMapStringString == nil {
				yyn390 = true
				goto LABEL390
			}
		LABEL390:
			if yyr2 || yy2arr2 {
				if yyn390 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringString == nil {
						r.EncodeNil()
					} else {
						yy391 := *x.FptrMapStringString
						yym392 := z.EncBinary()
						_ = yym392
						if false {
						} else {
							z.F.EncMapStringStringV(yy391, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringString"))
				r.WriteMapElemValue()
				if yyn390 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringString == nil {
						r.EncodeNil()
					} else {
						yy393 := *x.FptrMapStringString
						yym394 := z.EncBinary()
						_ = yym394
						if false {
						} else {
							z.F.EncMapStringStringV(yy393, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringUint == nil {
					r.EncodeNil()
				} else {
					yym396 := z.EncBinary()
					_ = yym396
					if false {
					} else {
						z.F.EncMapStringUintV(x.FMapStringUint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringUint"))
				r.WriteMapElemValue()
				if x.FMapStringUint == nil {
					r.EncodeNil()
				} else {
					yym397 := z.EncBinary()
					_ = yym397
					if false {
					} else {
						z.F.EncMapStringUintV(x.FMapStringUint, e)
					}
				}
			}
			var yyn398 bool
			if x.FptrMapStringUint == nil {
				yyn398 = true
				goto LABEL398
			}
		LABEL398:
			if yyr2 || yy2arr2 {
				if yyn398 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringUint == nil {
						r.EncodeNil()
					} else {
						yy399 := *x.FptrMapStringUint
						yym400 := z.EncBinary()
						_ = yym400
						if false {
						} else {
							z.F.EncMapStringUintV(yy399, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringUint"))
				r.WriteMapElemValue()
				if yyn398 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringUint == nil {
						r.EncodeNil()
					} else {
						yy401 := *x.FptrMapStringUint
						yym402 := z.EncBinary()
						_ = yym402
						if false {
						} else {
							z.F.EncMapStringUintV(yy401, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringUint8 == nil {
					r.EncodeNil()
				} else {
					yym404 := z.EncBinary()
					_ = yym404
					if false {
					} else {
						z.F.EncMapStringUint8V(x.FMapStringUint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringUint8"))
				r.WriteMapElemValue()
				if x.FMapStringUint8 == nil {
					r.EncodeNil()
				} else {
					yym405 := z.EncBinary()
					_ = yym405
					if false {
					} else {
						z.F.EncMapStringUint8V(x.FMapStringUint8, e)
					}
				}
			}
			var yyn406 bool
			if x.FptrMapStringUint8 == nil {
				yyn406 = true
				goto LABEL406
			}
		LABEL406:
			if yyr2 || yy2arr2 {
				if yyn406 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringUint8 == nil {
						r.EncodeNil()
					} else {
						yy407 := *x.FptrMapStringUint8
						yym408 := z.EncBinary()
						_ = yym408
						if false {
						} else {
							z.F.EncMapStringUint8V(yy407, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringUint8"))
				r.WriteMapElemValue()
				if yyn406 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringUint8 == nil {
						r.EncodeNil()
					} else {
						yy409 := *x.FptrMapStringUint8
						yym410 := z.EncBinary()
						_ = yym410
						if false {
						} else {
							z.F.EncMapStringUint8V(yy409, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringUint16 == nil {
					r.EncodeNil()
				} else {
					yym412 := z.EncBinary()
					_ = yym412
					if false {
					} else {
						z.F.EncMapStringUint16V(x.FMapStringUint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringUint16"))
				r.WriteMapElemValue()
				if x.FMapStringUint16 == nil {
					r.EncodeNil()
				} else {
					yym413 := z.EncBinary()
					_ = yym413
					if false {
					} else {
						z.F.EncMapStringUint16V(x.FMapStringUint16, e)
					}
				}
			}
			var yyn414 bool
			if x.FptrMapStringUint16 == nil {
				yyn414 = true
				goto LABEL414
			}
		LABEL414:
			if yyr2 || yy2arr2 {
				if yyn414 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringUint16 == nil {
						r.EncodeNil()
					} else {
						yy415 := *x.FptrMapStringUint16
						yym416 := z.EncBinary()
						_ = yym416
						if false {
						} else {
							z.F.EncMapStringUint16V(yy415, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringUint16"))
				r.WriteMapElemValue()
				if yyn414 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringUint16 == nil {
						r.EncodeNil()
					} else {
						yy417 := *x.FptrMapStringUint16
						yym418 := z.EncBinary()
						_ = yym418
						if false {
						} else {
							z.F.EncMapStringUint16V(yy417, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringUint32 == nil {
					r.EncodeNil()
				} else {
					yym420 := z.EncBinary()
					_ = yym420
					if false {
					} else {
						z.F.EncMapStringUint32V(x.FMapStringUint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringUint32"))
				r.WriteMapElemValue()
				if x.FMapStringUint32 == nil {
					r.EncodeNil()
				} else {
					yym421 := z.EncBinary()
					_ = yym421
					if false {
					} else {
						z.F.EncMapStringUint32V(x.FMapStringUint32, e)
					}
				}
			}
			var yyn422 bool
			if x.FptrMapStringUint32 == nil {
				yyn422 = true
				goto LABEL422
			}
		LABEL422:
			if yyr2 || yy2arr2 {
				if yyn422 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringUint32 == nil {
						r.EncodeNil()
					} else {
						yy423 := *x.FptrMapStringUint32
						yym424 := z.EncBinary()
						_ = yym424
						if false {
						} else {
							z.F.EncMapStringUint32V(yy423, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringUint32"))
				r.WriteMapElemValue()
				if yyn422 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringUint32 == nil {
						r.EncodeNil()
					} else {
						yy425 := *x.FptrMapStringUint32
						yym426 := z.EncBinary()
						_ = yym426
						if false {
						} else {
							z.F.EncMapStringUint32V(yy425, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringUint64 == nil {
					r.EncodeNil()
				} else {
					yym428 := z.EncBinary()
					_ = yym428
					if false {
					} else {
						z.F.EncMapStringUint64V(x.FMapStringUint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringUint64"))
				r.WriteMapElemValue()
				if x.FMapStringUint64 == nil {
					r.EncodeNil()
				} else {
					yym429 := z.EncBinary()
					_ = yym429
					if false {
					} else {
						z.F.EncMapStringUint64V(x.FMapStringUint64, e)
					}
				}
			}
			var yyn430 bool
			if x.FptrMapStringUint64 == nil {
				yyn430 = true
				goto LABEL430
			}
		LABEL430:
			if yyr2 || yy2arr2 {
				if yyn430 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringUint64 == nil {
						r.EncodeNil()
					} else {
						yy431 := *x.FptrMapStringUint64
						yym432 := z.EncBinary()
						_ = yym432
						if false {
						} else {
							z.F.EncMapStringUint64V(yy431, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringUint64"))
				r.WriteMapElemValue()
				if yyn430 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringUint64 == nil {
						r.EncodeNil()
					} else {
						yy433 := *x.FptrMapStringUint64
						yym434 := z.EncBinary()
						_ = yym434
						if false {
						} else {
							z.F.EncMapStringUint64V(yy433, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringUintptr == nil {
					r.EncodeNil()
				} else {
					yym436 := z.EncBinary()
					_ = yym436
					if false {
					} else {
						z.F.EncMapStringUintptrV(x.FMapStringUintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringUintptr"))
				r.WriteMapElemValue()
				if x.FMapStringUintptr == nil {
					r.EncodeNil()
				} else {
					yym437 := z.EncBinary()
					_ = yym437
					if false {
					} else {
						z.F.EncMapStringUintptrV(x.FMapStringUintptr, e)
					}
				}
			}
			var yyn438 bool
			if x.FptrMapStringUintptr == nil {
				yyn438 = true
				goto LABEL438
			}
		LABEL438:
			if yyr2 || yy2arr2 {
				if yyn438 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringUintptr == nil {
						r.EncodeNil()
					} else {
						yy439 := *x.FptrMapStringUintptr
						yym440 := z.EncBinary()
						_ = yym440
						if false {
						} else {
							z.F.EncMapStringUintptrV(yy439, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringUintptr"))
				r.WriteMapElemValue()
				if yyn438 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringUintptr == nil {
						r.EncodeNil()
					} else {
						yy441 := *x.FptrMapStringUintptr
						yym442 := z.EncBinary()
						_ = yym442
						if false {
						} else {
							z.F.EncMapStringUintptrV(yy441, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringInt == nil {
					r.EncodeNil()
				} else {
					yym444 := z.EncBinary()
					_ = yym444
					if false {
					} else {
						z.F.EncMapStringIntV(x.FMapStringInt, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringInt"))
				r.WriteMapElemValue()
				if x.FMapStringInt == nil {
					r.EncodeNil()
				} else {
					yym445 := z.EncBinary()
					_ = yym445
					if false {
					} else {
						z.F.EncMapStringIntV(x.FMapStringInt, e)
					}
				}
			}
			var yyn446 bool
			if x.FptrMapStringInt == nil {
				yyn446 = true
				goto LABEL446
			}
		LABEL446:
			if yyr2 || yy2arr2 {
				if yyn446 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringInt == nil {
						r.EncodeNil()
					} else {
						yy447 := *x.FptrMapStringInt
						yym448 := z.EncBinary()
						_ = yym448
						if false {
						} else {
							z.F.EncMapStringIntV(yy447, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringInt"))
				r.WriteMapElemValue()
				if yyn446 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringInt == nil {
						r.EncodeNil()
					} else {
						yy449 := *x.FptrMapStringInt
						yym450 := z.EncBinary()
						_ = yym450
						if false {
						} else {
							z.F.EncMapStringIntV(yy449, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringInt8 == nil {
					r.EncodeNil()
				} else {
					yym452 := z.EncBinary()
					_ = yym452
					if false {
					} else {
						z.F.EncMapStringInt8V(x.FMapStringInt8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringInt8"))
				r.WriteMapElemValue()
				if x.FMapStringInt8 == nil {
					r.EncodeNil()
				} else {
					yym453 := z.EncBinary()
					_ = yym453
					if false {
					} else {
						z.F.EncMapStringInt8V(x.FMapStringInt8, e)
					}
				}
			}
			var yyn454 bool
			if x.FptrMapStringInt8 == nil {
				yyn454 = true
				goto LABEL454
			}
		LABEL454:
			if yyr2 || yy2arr2 {
				if yyn454 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringInt8 == nil {
						r.EncodeNil()
					} else {
						yy455 := *x.FptrMapStringInt8
						yym456 := z.EncBinary()
						_ = yym456
						if false {
						} else {
							z.F.EncMapStringInt8V(yy455, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringInt8"))
				r.WriteMapElemValue()
				if yyn454 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringInt8 == nil {
						r.EncodeNil()
					} else {
						yy457 := *x.FptrMapStringInt8
						yym458 := z.EncBinary()
						_ = yym458
						if false {
						} else {
							z.F.EncMapStringInt8V(yy457, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringInt16 == nil {
					r.EncodeNil()
				} else {
					yym460 := z.EncBinary()
					_ = yym460
					if false {
					} else {
						z.F.EncMapStringInt16V(x.FMapStringInt16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringInt16"))
				r.WriteMapElemValue()
				if x.FMapStringInt16 == nil {
					r.EncodeNil()
				} else {
					yym461 := z.EncBinary()
					_ = yym461
					if false {
					} else {
						z.F.EncMapStringInt16V(x.FMapStringInt16, e)
					}
				}
			}
			var yyn462 bool
			if x.FptrMapStringInt16 == nil {
				yyn462 = true
				goto LABEL462
			}
		LABEL462:
			if yyr2 || yy2arr2 {
				if yyn462 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringInt16 == nil {
						r.EncodeNil()
					} else {
						yy463 := *x.FptrMapStringInt16
						yym464 := z.EncBinary()
						_ = yym464
						if false {
						} else {
							z.F.EncMapStringInt16V(yy463, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringInt16"))
				r.WriteMapElemValue()
				if yyn462 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringInt16 == nil {
						r.EncodeNil()
					} else {
						yy465 := *x.FptrMapStringInt16
						yym466 := z.EncBinary()
						_ = yym466
						if false {
						} else {
							z.F.EncMapStringInt16V(yy465, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringInt32 == nil {
					r.EncodeNil()
				} else {
					yym468 := z.EncBinary()
					_ = yym468
					if false {
					} else {
						z.F.EncMapStringInt32V(x.FMapStringInt32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringInt32"))
				r.WriteMapElemValue()
				if x.FMapStringInt32 == nil {
					r.EncodeNil()
				} else {
					yym469 := z.EncBinary()
					_ = yym469
					if false {
					} else {
						z.F.EncMapStringInt32V(x.FMapStringInt32, e)
					}
				}
			}
			var yyn470 bool
			if x.FptrMapStringInt32 == nil {
				yyn470 = true
				goto LABEL470
			}
		LABEL470:
			if yyr2 || yy2arr2 {
				if yyn470 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringInt32 == nil {
						r.EncodeNil()
					} else {
						yy471 := *x.FptrMapStringInt32
						yym472 := z.EncBinary()
						_ = yym472
						if false {
						} else {
							z.F.EncMapStringInt32V(yy471, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringInt32"))
				r.WriteMapElemValue()
				if yyn470 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringInt32 == nil {
						r.EncodeNil()
					} else {
						yy473 := *x.FptrMapStringInt32
						yym474 := z.EncBinary()
						_ = yym474
						if false {
						} else {
							z.F.EncMapStringInt32V(yy473, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringInt64 == nil {
					r.EncodeNil()
				} else {
					yym476 := z.EncBinary()
					_ = yym476
					if false {
					} else {
						z.F.EncMapStringInt64V(x.FMapStringInt64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringInt64"))
				r.WriteMapElemValue()
				if x.FMapStringInt64 == nil {
					r.EncodeNil()
				} else {
					yym477 := z.EncBinary()
					_ = yym477
					if false {
					} else {
						z.F.EncMapStringInt64V(x.FMapStringInt64, e)
					}
				}
			}
			var yyn478 bool
			if x.FptrMapStringInt64 == nil {
				yyn478 = true
				goto LABEL478
			}
		LABEL478:
			if yyr2 || yy2arr2 {
				if yyn478 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringInt64 == nil {
						r.EncodeNil()
					} else {
						yy479 := *x.FptrMapStringInt64
						yym480 := z.EncBinary()
						_ = yym480
						if false {
						} else {
							z.F.EncMapStringInt64V(yy479, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringInt64"))
				r.WriteMapElemValue()
				if yyn478 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringInt64 == nil {
						r.EncodeNil()
					} else {
						yy481 := *x.FptrMapStringInt64
						yym482 := z.EncBinary()
						_ = yym482
						if false {
						} else {
							z.F.EncMapStringInt64V(yy481, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringFloat32 == nil {
					r.EncodeNil()
				} else {
					yym484 := z.EncBinary()
					_ = yym484
					if false {
					} else {
						z.F.EncMapStringFloat32V(x.FMapStringFloat32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringFloat32"))
				r.WriteMapElemValue()
				if x.FMapStringFloat32 == nil {
					r.EncodeNil()
				} else {
					yym485 := z.EncBinary()
					_ = yym485
					if false {
					} else {
						z.F.EncMapStringFloat32V(x.FMapStringFloat32, e)
					}
				}
			}
			var yyn486 bool
			if x.FptrMapStringFloat32 == nil {
				yyn486 = true
				goto LABEL486
			}
		LABEL486:
			if yyr2 || yy2arr2 {
				if yyn486 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringFloat32 == nil {
						r.EncodeNil()
					} else {
						yy487 := *x.FptrMapStringFloat32
						yym488 := z.EncBinary()
						_ = yym488
						if false {
						} else {
							z.F.EncMapStringFloat32V(yy487, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringFloat32"))
				r.WriteMapElemValue()
				if yyn486 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringFloat32 == nil {
						r.EncodeNil()
					} else {
						yy489 := *x.FptrMapStringFloat32
						yym490 := z.EncBinary()
						_ = yym490
						if false {
						} else {
							z.F.EncMapStringFloat32V(yy489, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringFloat64 == nil {
					r.EncodeNil()
				} else {
					yym492 := z.EncBinary()
					_ = yym492
					if false {
					} else {
						z.F.EncMapStringFloat64V(x.FMapStringFloat64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringFloat64"))
				r.WriteMapElemValue()
				if x.FMapStringFloat64 == nil {
					r.EncodeNil()
				} else {
					yym493 := z.EncBinary()
					_ = yym493
					if false {
					} else {
						z.F.EncMapStringFloat64V(x.FMapStringFloat64, e)
					}
				}
			}
			var yyn494 bool
			if x.FptrMapStringFloat64 == nil {
				yyn494 = true
				goto LABEL494
			}
		LABEL494:
			if yyr2 || yy2arr2 {
				if yyn494 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringFloat64 == nil {
						r.EncodeNil()
					} else {
						yy495 := *x.FptrMapStringFloat64
						yym496 := z.EncBinary()
						_ = yym496
						if false {
						} else {
							z.F.EncMapStringFloat64V(yy495, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringFloat64"))
				r.WriteMapElemValue()
				if yyn494 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringFloat64 == nil {
						r.EncodeNil()
					} else {
						yy497 := *x.FptrMapStringFloat64
						yym498 := z.EncBinary()
						_ = yym498
						if false {
						} else {
							z.F.EncMapStringFloat64V(yy497, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapStringBool == nil {
					r.EncodeNil()
				} else {
					yym500 := z.EncBinary()
					_ = yym500
					if false {
					} else {
						z.F.EncMapStringBoolV(x.FMapStringBool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapStringBool"))
				r.WriteMapElemValue()
				if x.FMapStringBool == nil {
					r.EncodeNil()
				} else {
					yym501 := z.EncBinary()
					_ = yym501
					if false {
					} else {
						z.F.EncMapStringBoolV(x.FMapStringBool, e)
					}
				}
			}
			var yyn502 bool
			if x.FptrMapStringBool == nil {
				yyn502 = true
				goto LABEL502
			}
		LABEL502:
			if yyr2 || yy2arr2 {
				if yyn502 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapStringBool == nil {
						r.EncodeNil()
					} else {
						yy503 := *x.FptrMapStringBool
						yym504 := z.EncBinary()
						_ = yym504
						if false {
						} else {
							z.F.EncMapStringBoolV(yy503, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapStringBool"))
				r.WriteMapElemValue()
				if yyn502 {
					r.EncodeNil()
				} else {
					if x.FptrMapStringBool == nil {
						r.EncodeNil()
					} else {
						yy505 := *x.FptrMapStringBool
						yym506 := z.EncBinary()
						_ = yym506
						if false {
						} else {
							z.F.EncMapStringBoolV(yy505, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Intf == nil {
					r.EncodeNil()
				} else {
					yym508 := z.EncBinary()
					_ = yym508
					if false {
					} else {
						z.F.EncMapFloat32IntfV(x.FMapFloat32Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Intf"))
				r.WriteMapElemValue()
				if x.FMapFloat32Intf == nil {
					r.EncodeNil()
				} else {
					yym509 := z.EncBinary()
					_ = yym509
					if false {
					} else {
						z.F.EncMapFloat32IntfV(x.FMapFloat32Intf, e)
					}
				}
			}
			var yyn510 bool
			if x.FptrMapFloat32Intf == nil {
				yyn510 = true
				goto LABEL510
			}
		LABEL510:
			if yyr2 || yy2arr2 {
				if yyn510 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Intf == nil {
						r.EncodeNil()
					} else {
						yy511 := *x.FptrMapFloat32Intf
						yym512 := z.EncBinary()
						_ = yym512
						if false {
						} else {
							z.F.EncMapFloat32IntfV(yy511, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Intf"))
				r.WriteMapElemValue()
				if yyn510 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Intf == nil {
						r.EncodeNil()
					} else {
						yy513 := *x.FptrMapFloat32Intf
						yym514 := z.EncBinary()
						_ = yym514
						if false {
						} else {
							z.F.EncMapFloat32IntfV(yy513, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32String == nil {
					r.EncodeNil()
				} else {
					yym516 := z.EncBinary()
					_ = yym516
					if false {
					} else {
						z.F.EncMapFloat32StringV(x.FMapFloat32String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32String"))
				r.WriteMapElemValue()
				if x.FMapFloat32String == nil {
					r.EncodeNil()
				} else {
					yym517 := z.EncBinary()
					_ = yym517
					if false {
					} else {
						z.F.EncMapFloat32StringV(x.FMapFloat32String, e)
					}
				}
			}
			var yyn518 bool
			if x.FptrMapFloat32String == nil {
				yyn518 = true
				goto LABEL518
			}
		LABEL518:
			if yyr2 || yy2arr2 {
				if yyn518 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32String == nil {
						r.EncodeNil()
					} else {
						yy519 := *x.FptrMapFloat32String
						yym520 := z.EncBinary()
						_ = yym520
						if false {
						} else {
							z.F.EncMapFloat32StringV(yy519, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32String"))
				r.WriteMapElemValue()
				if yyn518 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32String == nil {
						r.EncodeNil()
					} else {
						yy521 := *x.FptrMapFloat32String
						yym522 := z.EncBinary()
						_ = yym522
						if false {
						} else {
							z.F.EncMapFloat32StringV(yy521, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Uint == nil {
					r.EncodeNil()
				} else {
					yym524 := z.EncBinary()
					_ = yym524
					if false {
					} else {
						z.F.EncMapFloat32UintV(x.FMapFloat32Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Uint"))
				r.WriteMapElemValue()
				if x.FMapFloat32Uint == nil {
					r.EncodeNil()
				} else {
					yym525 := z.EncBinary()
					_ = yym525
					if false {
					} else {
						z.F.EncMapFloat32UintV(x.FMapFloat32Uint, e)
					}
				}
			}
			var yyn526 bool
			if x.FptrMapFloat32Uint == nil {
				yyn526 = true
				goto LABEL526
			}
		LABEL526:
			if yyr2 || yy2arr2 {
				if yyn526 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Uint == nil {
						r.EncodeNil()
					} else {
						yy527 := *x.FptrMapFloat32Uint
						yym528 := z.EncBinary()
						_ = yym528
						if false {
						} else {
							z.F.EncMapFloat32UintV(yy527, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Uint"))
				r.WriteMapElemValue()
				if yyn526 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Uint == nil {
						r.EncodeNil()
					} else {
						yy529 := *x.FptrMapFloat32Uint
						yym530 := z.EncBinary()
						_ = yym530
						if false {
						} else {
							z.F.EncMapFloat32UintV(yy529, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Uint8 == nil {
					r.EncodeNil()
				} else {
					yym532 := z.EncBinary()
					_ = yym532
					if false {
					} else {
						z.F.EncMapFloat32Uint8V(x.FMapFloat32Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Uint8"))
				r.WriteMapElemValue()
				if x.FMapFloat32Uint8 == nil {
					r.EncodeNil()
				} else {
					yym533 := z.EncBinary()
					_ = yym533
					if false {
					} else {
						z.F.EncMapFloat32Uint8V(x.FMapFloat32Uint8, e)
					}
				}
			}
			var yyn534 bool
			if x.FptrMapFloat32Uint8 == nil {
				yyn534 = true
				goto LABEL534
			}
		LABEL534:
			if yyr2 || yy2arr2 {
				if yyn534 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Uint8 == nil {
						r.EncodeNil()
					} else {
						yy535 := *x.FptrMapFloat32Uint8
						yym536 := z.EncBinary()
						_ = yym536
						if false {
						} else {
							z.F.EncMapFloat32Uint8V(yy535, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Uint8"))
				r.WriteMapElemValue()
				if yyn534 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Uint8 == nil {
						r.EncodeNil()
					} else {
						yy537 := *x.FptrMapFloat32Uint8
						yym538 := z.EncBinary()
						_ = yym538
						if false {
						} else {
							z.F.EncMapFloat32Uint8V(yy537, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Uint16 == nil {
					r.EncodeNil()
				} else {
					yym540 := z.EncBinary()
					_ = yym540
					if false {
					} else {
						z.F.EncMapFloat32Uint16V(x.FMapFloat32Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Uint16"))
				r.WriteMapElemValue()
				if x.FMapFloat32Uint16 == nil {
					r.EncodeNil()
				} else {
					yym541 := z.EncBinary()
					_ = yym541
					if false {
					} else {
						z.F.EncMapFloat32Uint16V(x.FMapFloat32Uint16, e)
					}
				}
			}
			var yyn542 bool
			if x.FptrMapFloat32Uint16 == nil {
				yyn542 = true
				goto LABEL542
			}
		LABEL542:
			if yyr2 || yy2arr2 {
				if yyn542 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Uint16 == nil {
						r.EncodeNil()
					} else {
						yy543 := *x.FptrMapFloat32Uint16
						yym544 := z.EncBinary()
						_ = yym544
						if false {
						} else {
							z.F.EncMapFloat32Uint16V(yy543, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Uint16"))
				r.WriteMapElemValue()
				if yyn542 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Uint16 == nil {
						r.EncodeNil()
					} else {
						yy545 := *x.FptrMapFloat32Uint16
						yym546 := z.EncBinary()
						_ = yym546
						if false {
						} else {
							z.F.EncMapFloat32Uint16V(yy545, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Uint32 == nil {
					r.EncodeNil()
				} else {
					yym548 := z.EncBinary()
					_ = yym548
					if false {
					} else {
						z.F.EncMapFloat32Uint32V(x.FMapFloat32Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Uint32"))
				r.WriteMapElemValue()
				if x.FMapFloat32Uint32 == nil {
					r.EncodeNil()
				} else {
					yym549 := z.EncBinary()
					_ = yym549
					if false {
					} else {
						z.F.EncMapFloat32Uint32V(x.FMapFloat32Uint32, e)
					}
				}
			}
			var yyn550 bool
			if x.FptrMapFloat32Uint32 == nil {
				yyn550 = true
				goto LABEL550
			}
		LABEL550:
			if yyr2 || yy2arr2 {
				if yyn550 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Uint32 == nil {
						r.EncodeNil()
					} else {
						yy551 := *x.FptrMapFloat32Uint32
						yym552 := z.EncBinary()
						_ = yym552
						if false {
						} else {
							z.F.EncMapFloat32Uint32V(yy551, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Uint32"))
				r.WriteMapElemValue()
				if yyn550 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Uint32 == nil {
						r.EncodeNil()
					} else {
						yy553 := *x.FptrMapFloat32Uint32
						yym554 := z.EncBinary()
						_ = yym554
						if false {
						} else {
							z.F.EncMapFloat32Uint32V(yy553, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Uint64 == nil {
					r.EncodeNil()
				} else {
					yym556 := z.EncBinary()
					_ = yym556
					if false {
					} else {
						z.F.EncMapFloat32Uint64V(x.FMapFloat32Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Uint64"))
				r.WriteMapElemValue()
				if x.FMapFloat32Uint64 == nil {
					r.EncodeNil()
				} else {
					yym557 := z.EncBinary()
					_ = yym557
					if false {
					} else {
						z.F.EncMapFloat32Uint64V(x.FMapFloat32Uint64, e)
					}
				}
			}
			var yyn558 bool
			if x.FptrMapFloat32Uint64 == nil {
				yyn558 = true
				goto LABEL558
			}
		LABEL558:
			if yyr2 || yy2arr2 {
				if yyn558 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Uint64 == nil {
						r.EncodeNil()
					} else {
						yy559 := *x.FptrMapFloat32Uint64
						yym560 := z.EncBinary()
						_ = yym560
						if false {
						} else {
							z.F.EncMapFloat32Uint64V(yy559, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Uint64"))
				r.WriteMapElemValue()
				if yyn558 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Uint64 == nil {
						r.EncodeNil()
					} else {
						yy561 := *x.FptrMapFloat32Uint64
						yym562 := z.EncBinary()
						_ = yym562
						if false {
						} else {
							z.F.EncMapFloat32Uint64V(yy561, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Uintptr == nil {
					r.EncodeNil()
				} else {
					yym564 := z.EncBinary()
					_ = yym564
					if false {
					} else {
						z.F.EncMapFloat32UintptrV(x.FMapFloat32Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Uintptr"))
				r.WriteMapElemValue()
				if x.FMapFloat32Uintptr == nil {
					r.EncodeNil()
				} else {
					yym565 := z.EncBinary()
					_ = yym565
					if false {
					} else {
						z.F.EncMapFloat32UintptrV(x.FMapFloat32Uintptr, e)
					}
				}
			}
			var yyn566 bool
			if x.FptrMapFloat32Uintptr == nil {
				yyn566 = true
				goto LABEL566
			}
		LABEL566:
			if yyr2 || yy2arr2 {
				if yyn566 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Uintptr == nil {
						r.EncodeNil()
					} else {
						yy567 := *x.FptrMapFloat32Uintptr
						yym568 := z.EncBinary()
						_ = yym568
						if false {
						} else {
							z.F.EncMapFloat32UintptrV(yy567, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Uintptr"))
				r.WriteMapElemValue()
				if yyn566 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Uintptr == nil {
						r.EncodeNil()
					} else {
						yy569 := *x.FptrMapFloat32Uintptr
						yym570 := z.EncBinary()
						_ = yym570
						if false {
						} else {
							z.F.EncMapFloat32UintptrV(yy569, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Int == nil {
					r.EncodeNil()
				} else {
					yym572 := z.EncBinary()
					_ = yym572
					if false {
					} else {
						z.F.EncMapFloat32IntV(x.FMapFloat32Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Int"))
				r.WriteMapElemValue()
				if x.FMapFloat32Int == nil {
					r.EncodeNil()
				} else {
					yym573 := z.EncBinary()
					_ = yym573
					if false {
					} else {
						z.F.EncMapFloat32IntV(x.FMapFloat32Int, e)
					}
				}
			}
			var yyn574 bool
			if x.FptrMapFloat32Int == nil {
				yyn574 = true
				goto LABEL574
			}
		LABEL574:
			if yyr2 || yy2arr2 {
				if yyn574 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Int == nil {
						r.EncodeNil()
					} else {
						yy575 := *x.FptrMapFloat32Int
						yym576 := z.EncBinary()
						_ = yym576
						if false {
						} else {
							z.F.EncMapFloat32IntV(yy575, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Int"))
				r.WriteMapElemValue()
				if yyn574 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Int == nil {
						r.EncodeNil()
					} else {
						yy577 := *x.FptrMapFloat32Int
						yym578 := z.EncBinary()
						_ = yym578
						if false {
						} else {
							z.F.EncMapFloat32IntV(yy577, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Int8 == nil {
					r.EncodeNil()
				} else {
					yym580 := z.EncBinary()
					_ = yym580
					if false {
					} else {
						z.F.EncMapFloat32Int8V(x.FMapFloat32Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Int8"))
				r.WriteMapElemValue()
				if x.FMapFloat32Int8 == nil {
					r.EncodeNil()
				} else {
					yym581 := z.EncBinary()
					_ = yym581
					if false {
					} else {
						z.F.EncMapFloat32Int8V(x.FMapFloat32Int8, e)
					}
				}
			}
			var yyn582 bool
			if x.FptrMapFloat32Int8 == nil {
				yyn582 = true
				goto LABEL582
			}
		LABEL582:
			if yyr2 || yy2arr2 {
				if yyn582 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Int8 == nil {
						r.EncodeNil()
					} else {
						yy583 := *x.FptrMapFloat32Int8
						yym584 := z.EncBinary()
						_ = yym584
						if false {
						} else {
							z.F.EncMapFloat32Int8V(yy583, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Int8"))
				r.WriteMapElemValue()
				if yyn582 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Int8 == nil {
						r.EncodeNil()
					} else {
						yy585 := *x.FptrMapFloat32Int8
						yym586 := z.EncBinary()
						_ = yym586
						if false {
						} else {
							z.F.EncMapFloat32Int8V(yy585, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Int16 == nil {
					r.EncodeNil()
				} else {
					yym588 := z.EncBinary()
					_ = yym588
					if false {
					} else {
						z.F.EncMapFloat32Int16V(x.FMapFloat32Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Int16"))
				r.WriteMapElemValue()
				if x.FMapFloat32Int16 == nil {
					r.EncodeNil()
				} else {
					yym589 := z.EncBinary()
					_ = yym589
					if false {
					} else {
						z.F.EncMapFloat32Int16V(x.FMapFloat32Int16, e)
					}
				}
			}
			var yyn590 bool
			if x.FptrMapFloat32Int16 == nil {
				yyn590 = true
				goto LABEL590
			}
		LABEL590:
			if yyr2 || yy2arr2 {
				if yyn590 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Int16 == nil {
						r.EncodeNil()
					} else {
						yy591 := *x.FptrMapFloat32Int16
						yym592 := z.EncBinary()
						_ = yym592
						if false {
						} else {
							z.F.EncMapFloat32Int16V(yy591, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Int16"))
				r.WriteMapElemValue()
				if yyn590 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Int16 == nil {
						r.EncodeNil()
					} else {
						yy593 := *x.FptrMapFloat32Int16
						yym594 := z.EncBinary()
						_ = yym594
						if false {
						} else {
							z.F.EncMapFloat32Int16V(yy593, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Int32 == nil {
					r.EncodeNil()
				} else {
					yym596 := z.EncBinary()
					_ = yym596
					if false {
					} else {
						z.F.EncMapFloat32Int32V(x.FMapFloat32Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Int32"))
				r.WriteMapElemValue()
				if x.FMapFloat32Int32 == nil {
					r.EncodeNil()
				} else {
					yym597 := z.EncBinary()
					_ = yym597
					if false {
					} else {
						z.F.EncMapFloat32Int32V(x.FMapFloat32Int32, e)
					}
				}
			}
			var yyn598 bool
			if x.FptrMapFloat32Int32 == nil {
				yyn598 = true
				goto LABEL598
			}
		LABEL598:
			if yyr2 || yy2arr2 {
				if yyn598 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Int32 == nil {
						r.EncodeNil()
					} else {
						yy599 := *x.FptrMapFloat32Int32
						yym600 := z.EncBinary()
						_ = yym600
						if false {
						} else {
							z.F.EncMapFloat32Int32V(yy599, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Int32"))
				r.WriteMapElemValue()
				if yyn598 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Int32 == nil {
						r.EncodeNil()
					} else {
						yy601 := *x.FptrMapFloat32Int32
						yym602 := z.EncBinary()
						_ = yym602
						if false {
						} else {
							z.F.EncMapFloat32Int32V(yy601, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Int64 == nil {
					r.EncodeNil()
				} else {
					yym604 := z.EncBinary()
					_ = yym604
					if false {
					} else {
						z.F.EncMapFloat32Int64V(x.FMapFloat32Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Int64"))
				r.WriteMapElemValue()
				if x.FMapFloat32Int64 == nil {
					r.EncodeNil()
				} else {
					yym605 := z.EncBinary()
					_ = yym605
					if false {
					} else {
						z.F.EncMapFloat32Int64V(x.FMapFloat32Int64, e)
					}
				}
			}
			var yyn606 bool
			if x.FptrMapFloat32Int64 == nil {
				yyn606 = true
				goto LABEL606
			}
		LABEL606:
			if yyr2 || yy2arr2 {
				if yyn606 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Int64 == nil {
						r.EncodeNil()
					} else {
						yy607 := *x.FptrMapFloat32Int64
						yym608 := z.EncBinary()
						_ = yym608
						if false {
						} else {
							z.F.EncMapFloat32Int64V(yy607, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Int64"))
				r.WriteMapElemValue()
				if yyn606 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Int64 == nil {
						r.EncodeNil()
					} else {
						yy609 := *x.FptrMapFloat32Int64
						yym610 := z.EncBinary()
						_ = yym610
						if false {
						} else {
							z.F.EncMapFloat32Int64V(yy609, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Float32 == nil {
					r.EncodeNil()
				} else {
					yym612 := z.EncBinary()
					_ = yym612
					if false {
					} else {
						z.F.EncMapFloat32Float32V(x.FMapFloat32Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Float32"))
				r.WriteMapElemValue()
				if x.FMapFloat32Float32 == nil {
					r.EncodeNil()
				} else {
					yym613 := z.EncBinary()
					_ = yym613
					if false {
					} else {
						z.F.EncMapFloat32Float32V(x.FMapFloat32Float32, e)
					}
				}
			}
			var yyn614 bool
			if x.FptrMapFloat32Float32 == nil {
				yyn614 = true
				goto LABEL614
			}
		LABEL614:
			if yyr2 || yy2arr2 {
				if yyn614 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Float32 == nil {
						r.EncodeNil()
					} else {
						yy615 := *x.FptrMapFloat32Float32
						yym616 := z.EncBinary()
						_ = yym616
						if false {
						} else {
							z.F.EncMapFloat32Float32V(yy615, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Float32"))
				r.WriteMapElemValue()
				if yyn614 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Float32 == nil {
						r.EncodeNil()
					} else {
						yy617 := *x.FptrMapFloat32Float32
						yym618 := z.EncBinary()
						_ = yym618
						if false {
						} else {
							z.F.EncMapFloat32Float32V(yy617, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Float64 == nil {
					r.EncodeNil()
				} else {
					yym620 := z.EncBinary()
					_ = yym620
					if false {
					} else {
						z.F.EncMapFloat32Float64V(x.FMapFloat32Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Float64"))
				r.WriteMapElemValue()
				if x.FMapFloat32Float64 == nil {
					r.EncodeNil()
				} else {
					yym621 := z.EncBinary()
					_ = yym621
					if false {
					} else {
						z.F.EncMapFloat32Float64V(x.FMapFloat32Float64, e)
					}
				}
			}
			var yyn622 bool
			if x.FptrMapFloat32Float64 == nil {
				yyn622 = true
				goto LABEL622
			}
		LABEL622:
			if yyr2 || yy2arr2 {
				if yyn622 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Float64 == nil {
						r.EncodeNil()
					} else {
						yy623 := *x.FptrMapFloat32Float64
						yym624 := z.EncBinary()
						_ = yym624
						if false {
						} else {
							z.F.EncMapFloat32Float64V(yy623, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Float64"))
				r.WriteMapElemValue()
				if yyn622 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Float64 == nil {
						r.EncodeNil()
					} else {
						yy625 := *x.FptrMapFloat32Float64
						yym626 := z.EncBinary()
						_ = yym626
						if false {
						} else {
							z.F.EncMapFloat32Float64V(yy625, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat32Bool == nil {
					r.EncodeNil()
				} else {
					yym628 := z.EncBinary()
					_ = yym628
					if false {
					} else {
						z.F.EncMapFloat32BoolV(x.FMapFloat32Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat32Bool"))
				r.WriteMapElemValue()
				if x.FMapFloat32Bool == nil {
					r.EncodeNil()
				} else {
					yym629 := z.EncBinary()
					_ = yym629
					if false {
					} else {
						z.F.EncMapFloat32BoolV(x.FMapFloat32Bool, e)
					}
				}
			}
			var yyn630 bool
			if x.FptrMapFloat32Bool == nil {
				yyn630 = true
				goto LABEL630
			}
		LABEL630:
			if yyr2 || yy2arr2 {
				if yyn630 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat32Bool == nil {
						r.EncodeNil()
					} else {
						yy631 := *x.FptrMapFloat32Bool
						yym632 := z.EncBinary()
						_ = yym632
						if false {
						} else {
							z.F.EncMapFloat32BoolV(yy631, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat32Bool"))
				r.WriteMapElemValue()
				if yyn630 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat32Bool == nil {
						r.EncodeNil()
					} else {
						yy633 := *x.FptrMapFloat32Bool
						yym634 := z.EncBinary()
						_ = yym634
						if false {
						} else {
							z.F.EncMapFloat32BoolV(yy633, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Intf == nil {
					r.EncodeNil()
				} else {
					yym636 := z.EncBinary()
					_ = yym636
					if false {
					} else {
						z.F.EncMapFloat64IntfV(x.FMapFloat64Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Intf"))
				r.WriteMapElemValue()
				if x.FMapFloat64Intf == nil {
					r.EncodeNil()
				} else {
					yym637 := z.EncBinary()
					_ = yym637
					if false {
					} else {
						z.F.EncMapFloat64IntfV(x.FMapFloat64Intf, e)
					}
				}
			}
			var yyn638 bool
			if x.FptrMapFloat64Intf == nil {
				yyn638 = true
				goto LABEL638
			}
		LABEL638:
			if yyr2 || yy2arr2 {
				if yyn638 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Intf == nil {
						r.EncodeNil()
					} else {
						yy639 := *x.FptrMapFloat64Intf
						yym640 := z.EncBinary()
						_ = yym640
						if false {
						} else {
							z.F.EncMapFloat64IntfV(yy639, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Intf"))
				r.WriteMapElemValue()
				if yyn638 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Intf == nil {
						r.EncodeNil()
					} else {
						yy641 := *x.FptrMapFloat64Intf
						yym642 := z.EncBinary()
						_ = yym642
						if false {
						} else {
							z.F.EncMapFloat64IntfV(yy641, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64String == nil {
					r.EncodeNil()
				} else {
					yym644 := z.EncBinary()
					_ = yym644
					if false {
					} else {
						z.F.EncMapFloat64StringV(x.FMapFloat64String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64String"))
				r.WriteMapElemValue()
				if x.FMapFloat64String == nil {
					r.EncodeNil()
				} else {
					yym645 := z.EncBinary()
					_ = yym645
					if false {
					} else {
						z.F.EncMapFloat64StringV(x.FMapFloat64String, e)
					}
				}
			}
			var yyn646 bool
			if x.FptrMapFloat64String == nil {
				yyn646 = true
				goto LABEL646
			}
		LABEL646:
			if yyr2 || yy2arr2 {
				if yyn646 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64String == nil {
						r.EncodeNil()
					} else {
						yy647 := *x.FptrMapFloat64String
						yym648 := z.EncBinary()
						_ = yym648
						if false {
						} else {
							z.F.EncMapFloat64StringV(yy647, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64String"))
				r.WriteMapElemValue()
				if yyn646 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64String == nil {
						r.EncodeNil()
					} else {
						yy649 := *x.FptrMapFloat64String
						yym650 := z.EncBinary()
						_ = yym650
						if false {
						} else {
							z.F.EncMapFloat64StringV(yy649, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Uint == nil {
					r.EncodeNil()
				} else {
					yym652 := z.EncBinary()
					_ = yym652
					if false {
					} else {
						z.F.EncMapFloat64UintV(x.FMapFloat64Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Uint"))
				r.WriteMapElemValue()
				if x.FMapFloat64Uint == nil {
					r.EncodeNil()
				} else {
					yym653 := z.EncBinary()
					_ = yym653
					if false {
					} else {
						z.F.EncMapFloat64UintV(x.FMapFloat64Uint, e)
					}
				}
			}
			var yyn654 bool
			if x.FptrMapFloat64Uint == nil {
				yyn654 = true
				goto LABEL654
			}
		LABEL654:
			if yyr2 || yy2arr2 {
				if yyn654 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Uint == nil {
						r.EncodeNil()
					} else {
						yy655 := *x.FptrMapFloat64Uint
						yym656 := z.EncBinary()
						_ = yym656
						if false {
						} else {
							z.F.EncMapFloat64UintV(yy655, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Uint"))
				r.WriteMapElemValue()
				if yyn654 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Uint == nil {
						r.EncodeNil()
					} else {
						yy657 := *x.FptrMapFloat64Uint
						yym658 := z.EncBinary()
						_ = yym658
						if false {
						} else {
							z.F.EncMapFloat64UintV(yy657, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Uint8 == nil {
					r.EncodeNil()
				} else {
					yym660 := z.EncBinary()
					_ = yym660
					if false {
					} else {
						z.F.EncMapFloat64Uint8V(x.FMapFloat64Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Uint8"))
				r.WriteMapElemValue()
				if x.FMapFloat64Uint8 == nil {
					r.EncodeNil()
				} else {
					yym661 := z.EncBinary()
					_ = yym661
					if false {
					} else {
						z.F.EncMapFloat64Uint8V(x.FMapFloat64Uint8, e)
					}
				}
			}
			var yyn662 bool
			if x.FptrMapFloat64Uint8 == nil {
				yyn662 = true
				goto LABEL662
			}
		LABEL662:
			if yyr2 || yy2arr2 {
				if yyn662 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Uint8 == nil {
						r.EncodeNil()
					} else {
						yy663 := *x.FptrMapFloat64Uint8
						yym664 := z.EncBinary()
						_ = yym664
						if false {
						} else {
							z.F.EncMapFloat64Uint8V(yy663, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Uint8"))
				r.WriteMapElemValue()
				if yyn662 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Uint8 == nil {
						r.EncodeNil()
					} else {
						yy665 := *x.FptrMapFloat64Uint8
						yym666 := z.EncBinary()
						_ = yym666
						if false {
						} else {
							z.F.EncMapFloat64Uint8V(yy665, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Uint16 == nil {
					r.EncodeNil()
				} else {
					yym668 := z.EncBinary()
					_ = yym668
					if false {
					} else {
						z.F.EncMapFloat64Uint16V(x.FMapFloat64Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Uint16"))
				r.WriteMapElemValue()
				if x.FMapFloat64Uint16 == nil {
					r.EncodeNil()
				} else {
					yym669 := z.EncBinary()
					_ = yym669
					if false {
					} else {
						z.F.EncMapFloat64Uint16V(x.FMapFloat64Uint16, e)
					}
				}
			}
			var yyn670 bool
			if x.FptrMapFloat64Uint16 == nil {
				yyn670 = true
				goto LABEL670
			}
		LABEL670:
			if yyr2 || yy2arr2 {
				if yyn670 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Uint16 == nil {
						r.EncodeNil()
					} else {
						yy671 := *x.FptrMapFloat64Uint16
						yym672 := z.EncBinary()
						_ = yym672
						if false {
						} else {
							z.F.EncMapFloat64Uint16V(yy671, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Uint16"))
				r.WriteMapElemValue()
				if yyn670 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Uint16 == nil {
						r.EncodeNil()
					} else {
						yy673 := *x.FptrMapFloat64Uint16
						yym674 := z.EncBinary()
						_ = yym674
						if false {
						} else {
							z.F.EncMapFloat64Uint16V(yy673, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Uint32 == nil {
					r.EncodeNil()
				} else {
					yym676 := z.EncBinary()
					_ = yym676
					if false {
					} else {
						z.F.EncMapFloat64Uint32V(x.FMapFloat64Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Uint32"))
				r.WriteMapElemValue()
				if x.FMapFloat64Uint32 == nil {
					r.EncodeNil()
				} else {
					yym677 := z.EncBinary()
					_ = yym677
					if false {
					} else {
						z.F.EncMapFloat64Uint32V(x.FMapFloat64Uint32, e)
					}
				}
			}
			var yyn678 bool
			if x.FptrMapFloat64Uint32 == nil {
				yyn678 = true
				goto LABEL678
			}
		LABEL678:
			if yyr2 || yy2arr2 {
				if yyn678 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Uint32 == nil {
						r.EncodeNil()
					} else {
						yy679 := *x.FptrMapFloat64Uint32
						yym680 := z.EncBinary()
						_ = yym680
						if false {
						} else {
							z.F.EncMapFloat64Uint32V(yy679, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Uint32"))
				r.WriteMapElemValue()
				if yyn678 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Uint32 == nil {
						r.EncodeNil()
					} else {
						yy681 := *x.FptrMapFloat64Uint32
						yym682 := z.EncBinary()
						_ = yym682
						if false {
						} else {
							z.F.EncMapFloat64Uint32V(yy681, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Uint64 == nil {
					r.EncodeNil()
				} else {
					yym684 := z.EncBinary()
					_ = yym684
					if false {
					} else {
						z.F.EncMapFloat64Uint64V(x.FMapFloat64Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Uint64"))
				r.WriteMapElemValue()
				if x.FMapFloat64Uint64 == nil {
					r.EncodeNil()
				} else {
					yym685 := z.EncBinary()
					_ = yym685
					if false {
					} else {
						z.F.EncMapFloat64Uint64V(x.FMapFloat64Uint64, e)
					}
				}
			}
			var yyn686 bool
			if x.FptrMapFloat64Uint64 == nil {
				yyn686 = true
				goto LABEL686
			}
		LABEL686:
			if yyr2 || yy2arr2 {
				if yyn686 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Uint64 == nil {
						r.EncodeNil()
					} else {
						yy687 := *x.FptrMapFloat64Uint64
						yym688 := z.EncBinary()
						_ = yym688
						if false {
						} else {
							z.F.EncMapFloat64Uint64V(yy687, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Uint64"))
				r.WriteMapElemValue()
				if yyn686 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Uint64 == nil {
						r.EncodeNil()
					} else {
						yy689 := *x.FptrMapFloat64Uint64
						yym690 := z.EncBinary()
						_ = yym690
						if false {
						} else {
							z.F.EncMapFloat64Uint64V(yy689, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Uintptr == nil {
					r.EncodeNil()
				} else {
					yym692 := z.EncBinary()
					_ = yym692
					if false {
					} else {
						z.F.EncMapFloat64UintptrV(x.FMapFloat64Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Uintptr"))
				r.WriteMapElemValue()
				if x.FMapFloat64Uintptr == nil {
					r.EncodeNil()
				} else {
					yym693 := z.EncBinary()
					_ = yym693
					if false {
					} else {
						z.F.EncMapFloat64UintptrV(x.FMapFloat64Uintptr, e)
					}
				}
			}
			var yyn694 bool
			if x.FptrMapFloat64Uintptr == nil {
				yyn694 = true
				goto LABEL694
			}
		LABEL694:
			if yyr2 || yy2arr2 {
				if yyn694 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Uintptr == nil {
						r.EncodeNil()
					} else {
						yy695 := *x.FptrMapFloat64Uintptr
						yym696 := z.EncBinary()
						_ = yym696
						if false {
						} else {
							z.F.EncMapFloat64UintptrV(yy695, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Uintptr"))
				r.WriteMapElemValue()
				if yyn694 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Uintptr == nil {
						r.EncodeNil()
					} else {
						yy697 := *x.FptrMapFloat64Uintptr
						yym698 := z.EncBinary()
						_ = yym698
						if false {
						} else {
							z.F.EncMapFloat64UintptrV(yy697, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Int == nil {
					r.EncodeNil()
				} else {
					yym700 := z.EncBinary()
					_ = yym700
					if false {
					} else {
						z.F.EncMapFloat64IntV(x.FMapFloat64Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Int"))
				r.WriteMapElemValue()
				if x.FMapFloat64Int == nil {
					r.EncodeNil()
				} else {
					yym701 := z.EncBinary()
					_ = yym701
					if false {
					} else {
						z.F.EncMapFloat64IntV(x.FMapFloat64Int, e)
					}
				}
			}
			var yyn702 bool
			if x.FptrMapFloat64Int == nil {
				yyn702 = true
				goto LABEL702
			}
		LABEL702:
			if yyr2 || yy2arr2 {
				if yyn702 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Int == nil {
						r.EncodeNil()
					} else {
						yy703 := *x.FptrMapFloat64Int
						yym704 := z.EncBinary()
						_ = yym704
						if false {
						} else {
							z.F.EncMapFloat64IntV(yy703, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Int"))
				r.WriteMapElemValue()
				if yyn702 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Int == nil {
						r.EncodeNil()
					} else {
						yy705 := *x.FptrMapFloat64Int
						yym706 := z.EncBinary()
						_ = yym706
						if false {
						} else {
							z.F.EncMapFloat64IntV(yy705, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Int8 == nil {
					r.EncodeNil()
				} else {
					yym708 := z.EncBinary()
					_ = yym708
					if false {
					} else {
						z.F.EncMapFloat64Int8V(x.FMapFloat64Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Int8"))
				r.WriteMapElemValue()
				if x.FMapFloat64Int8 == nil {
					r.EncodeNil()
				} else {
					yym709 := z.EncBinary()
					_ = yym709
					if false {
					} else {
						z.F.EncMapFloat64Int8V(x.FMapFloat64Int8, e)
					}
				}
			}
			var yyn710 bool
			if x.FptrMapFloat64Int8 == nil {
				yyn710 = true
				goto LABEL710
			}
		LABEL710:
			if yyr2 || yy2arr2 {
				if yyn710 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Int8 == nil {
						r.EncodeNil()
					} else {
						yy711 := *x.FptrMapFloat64Int8
						yym712 := z.EncBinary()
						_ = yym712
						if false {
						} else {
							z.F.EncMapFloat64Int8V(yy711, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Int8"))
				r.WriteMapElemValue()
				if yyn710 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Int8 == nil {
						r.EncodeNil()
					} else {
						yy713 := *x.FptrMapFloat64Int8
						yym714 := z.EncBinary()
						_ = yym714
						if false {
						} else {
							z.F.EncMapFloat64Int8V(yy713, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Int16 == nil {
					r.EncodeNil()
				} else {
					yym716 := z.EncBinary()
					_ = yym716
					if false {
					} else {
						z.F.EncMapFloat64Int16V(x.FMapFloat64Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Int16"))
				r.WriteMapElemValue()
				if x.FMapFloat64Int16 == nil {
					r.EncodeNil()
				} else {
					yym717 := z.EncBinary()
					_ = yym717
					if false {
					} else {
						z.F.EncMapFloat64Int16V(x.FMapFloat64Int16, e)
					}
				}
			}
			var yyn718 bool
			if x.FptrMapFloat64Int16 == nil {
				yyn718 = true
				goto LABEL718
			}
		LABEL718:
			if yyr2 || yy2arr2 {
				if yyn718 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Int16 == nil {
						r.EncodeNil()
					} else {
						yy719 := *x.FptrMapFloat64Int16
						yym720 := z.EncBinary()
						_ = yym720
						if false {
						} else {
							z.F.EncMapFloat64Int16V(yy719, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Int16"))
				r.WriteMapElemValue()
				if yyn718 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Int16 == nil {
						r.EncodeNil()
					} else {
						yy721 := *x.FptrMapFloat64Int16
						yym722 := z.EncBinary()
						_ = yym722
						if false {
						} else {
							z.F.EncMapFloat64Int16V(yy721, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Int32 == nil {
					r.EncodeNil()
				} else {
					yym724 := z.EncBinary()
					_ = yym724
					if false {
					} else {
						z.F.EncMapFloat64Int32V(x.FMapFloat64Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Int32"))
				r.WriteMapElemValue()
				if x.FMapFloat64Int32 == nil {
					r.EncodeNil()
				} else {
					yym725 := z.EncBinary()
					_ = yym725
					if false {
					} else {
						z.F.EncMapFloat64Int32V(x.FMapFloat64Int32, e)
					}
				}
			}
			var yyn726 bool
			if x.FptrMapFloat64Int32 == nil {
				yyn726 = true
				goto LABEL726
			}
		LABEL726:
			if yyr2 || yy2arr2 {
				if yyn726 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Int32 == nil {
						r.EncodeNil()
					} else {
						yy727 := *x.FptrMapFloat64Int32
						yym728 := z.EncBinary()
						_ = yym728
						if false {
						} else {
							z.F.EncMapFloat64Int32V(yy727, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Int32"))
				r.WriteMapElemValue()
				if yyn726 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Int32 == nil {
						r.EncodeNil()
					} else {
						yy729 := *x.FptrMapFloat64Int32
						yym730 := z.EncBinary()
						_ = yym730
						if false {
						} else {
							z.F.EncMapFloat64Int32V(yy729, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Int64 == nil {
					r.EncodeNil()
				} else {
					yym732 := z.EncBinary()
					_ = yym732
					if false {
					} else {
						z.F.EncMapFloat64Int64V(x.FMapFloat64Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Int64"))
				r.WriteMapElemValue()
				if x.FMapFloat64Int64 == nil {
					r.EncodeNil()
				} else {
					yym733 := z.EncBinary()
					_ = yym733
					if false {
					} else {
						z.F.EncMapFloat64Int64V(x.FMapFloat64Int64, e)
					}
				}
			}
			var yyn734 bool
			if x.FptrMapFloat64Int64 == nil {
				yyn734 = true
				goto LABEL734
			}
		LABEL734:
			if yyr2 || yy2arr2 {
				if yyn734 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Int64 == nil {
						r.EncodeNil()
					} else {
						yy735 := *x.FptrMapFloat64Int64
						yym736 := z.EncBinary()
						_ = yym736
						if false {
						} else {
							z.F.EncMapFloat64Int64V(yy735, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Int64"))
				r.WriteMapElemValue()
				if yyn734 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Int64 == nil {
						r.EncodeNil()
					} else {
						yy737 := *x.FptrMapFloat64Int64
						yym738 := z.EncBinary()
						_ = yym738
						if false {
						} else {
							z.F.EncMapFloat64Int64V(yy737, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Float32 == nil {
					r.EncodeNil()
				} else {
					yym740 := z.EncBinary()
					_ = yym740
					if false {
					} else {
						z.F.EncMapFloat64Float32V(x.FMapFloat64Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Float32"))
				r.WriteMapElemValue()
				if x.FMapFloat64Float32 == nil {
					r.EncodeNil()
				} else {
					yym741 := z.EncBinary()
					_ = yym741
					if false {
					} else {
						z.F.EncMapFloat64Float32V(x.FMapFloat64Float32, e)
					}
				}
			}
			var yyn742 bool
			if x.FptrMapFloat64Float32 == nil {
				yyn742 = true
				goto LABEL742
			}
		LABEL742:
			if yyr2 || yy2arr2 {
				if yyn742 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Float32 == nil {
						r.EncodeNil()
					} else {
						yy743 := *x.FptrMapFloat64Float32
						yym744 := z.EncBinary()
						_ = yym744
						if false {
						} else {
							z.F.EncMapFloat64Float32V(yy743, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Float32"))
				r.WriteMapElemValue()
				if yyn742 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Float32 == nil {
						r.EncodeNil()
					} else {
						yy745 := *x.FptrMapFloat64Float32
						yym746 := z.EncBinary()
						_ = yym746
						if false {
						} else {
							z.F.EncMapFloat64Float32V(yy745, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Float64 == nil {
					r.EncodeNil()
				} else {
					yym748 := z.EncBinary()
					_ = yym748
					if false {
					} else {
						z.F.EncMapFloat64Float64V(x.FMapFloat64Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Float64"))
				r.WriteMapElemValue()
				if x.FMapFloat64Float64 == nil {
					r.EncodeNil()
				} else {
					yym749 := z.EncBinary()
					_ = yym749
					if false {
					} else {
						z.F.EncMapFloat64Float64V(x.FMapFloat64Float64, e)
					}
				}
			}
			var yyn750 bool
			if x.FptrMapFloat64Float64 == nil {
				yyn750 = true
				goto LABEL750
			}
		LABEL750:
			if yyr2 || yy2arr2 {
				if yyn750 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Float64 == nil {
						r.EncodeNil()
					} else {
						yy751 := *x.FptrMapFloat64Float64
						yym752 := z.EncBinary()
						_ = yym752
						if false {
						} else {
							z.F.EncMapFloat64Float64V(yy751, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Float64"))
				r.WriteMapElemValue()
				if yyn750 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Float64 == nil {
						r.EncodeNil()
					} else {
						yy753 := *x.FptrMapFloat64Float64
						yym754 := z.EncBinary()
						_ = yym754
						if false {
						} else {
							z.F.EncMapFloat64Float64V(yy753, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapFloat64Bool == nil {
					r.EncodeNil()
				} else {
					yym756 := z.EncBinary()
					_ = yym756
					if false {
					} else {
						z.F.EncMapFloat64BoolV(x.FMapFloat64Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapFloat64Bool"))
				r.WriteMapElemValue()
				if x.FMapFloat64Bool == nil {
					r.EncodeNil()
				} else {
					yym757 := z.EncBinary()
					_ = yym757
					if false {
					} else {
						z.F.EncMapFloat64BoolV(x.FMapFloat64Bool, e)
					}
				}
			}
			var yyn758 bool
			if x.FptrMapFloat64Bool == nil {
				yyn758 = true
				goto LABEL758
			}
		LABEL758:
			if yyr2 || yy2arr2 {
				if yyn758 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapFloat64Bool == nil {
						r.EncodeNil()
					} else {
						yy759 := *x.FptrMapFloat64Bool
						yym760 := z.EncBinary()
						_ = yym760
						if false {
						} else {
							z.F.EncMapFloat64BoolV(yy759, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapFloat64Bool"))
				r.WriteMapElemValue()
				if yyn758 {
					r.EncodeNil()
				} else {
					if x.FptrMapFloat64Bool == nil {
						r.EncodeNil()
					} else {
						yy761 := *x.FptrMapFloat64Bool
						yym762 := z.EncBinary()
						_ = yym762
						if false {
						} else {
							z.F.EncMapFloat64BoolV(yy761, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintIntf == nil {
					r.EncodeNil()
				} else {
					yym764 := z.EncBinary()
					_ = yym764
					if false {
					} else {
						z.F.EncMapUintIntfV(x.FMapUintIntf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintIntf"))
				r.WriteMapElemValue()
				if x.FMapUintIntf == nil {
					r.EncodeNil()
				} else {
					yym765 := z.EncBinary()
					_ = yym765
					if false {
					} else {
						z.F.EncMapUintIntfV(x.FMapUintIntf, e)
					}
				}
			}
			var yyn766 bool
			if x.FptrMapUintIntf == nil {
				yyn766 = true
				goto LABEL766
			}
		LABEL766:
			if yyr2 || yy2arr2 {
				if yyn766 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintIntf == nil {
						r.EncodeNil()
					} else {
						yy767 := *x.FptrMapUintIntf
						yym768 := z.EncBinary()
						_ = yym768
						if false {
						} else {
							z.F.EncMapUintIntfV(yy767, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintIntf"))
				r.WriteMapElemValue()
				if yyn766 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintIntf == nil {
						r.EncodeNil()
					} else {
						yy769 := *x.FptrMapUintIntf
						yym770 := z.EncBinary()
						_ = yym770
						if false {
						} else {
							z.F.EncMapUintIntfV(yy769, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintString == nil {
					r.EncodeNil()
				} else {
					yym772 := z.EncBinary()
					_ = yym772
					if false {
					} else {
						z.F.EncMapUintStringV(x.FMapUintString, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintString"))
				r.WriteMapElemValue()
				if x.FMapUintString == nil {
					r.EncodeNil()
				} else {
					yym773 := z.EncBinary()
					_ = yym773
					if false {
					} else {
						z.F.EncMapUintStringV(x.FMapUintString, e)
					}
				}
			}
			var yyn774 bool
			if x.FptrMapUintString == nil {
				yyn774 = true
				goto LABEL774
			}
		LABEL774:
			if yyr2 || yy2arr2 {
				if yyn774 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintString == nil {
						r.EncodeNil()
					} else {
						yy775 := *x.FptrMapUintString
						yym776 := z.EncBinary()
						_ = yym776
						if false {
						} else {
							z.F.EncMapUintStringV(yy775, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintString"))
				r.WriteMapElemValue()
				if yyn774 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintString == nil {
						r.EncodeNil()
					} else {
						yy777 := *x.FptrMapUintString
						yym778 := z.EncBinary()
						_ = yym778
						if false {
						} else {
							z.F.EncMapUintStringV(yy777, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintUint == nil {
					r.EncodeNil()
				} else {
					yym780 := z.EncBinary()
					_ = yym780
					if false {
					} else {
						z.F.EncMapUintUintV(x.FMapUintUint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintUint"))
				r.WriteMapElemValue()
				if x.FMapUintUint == nil {
					r.EncodeNil()
				} else {
					yym781 := z.EncBinary()
					_ = yym781
					if false {
					} else {
						z.F.EncMapUintUintV(x.FMapUintUint, e)
					}
				}
			}
			var yyn782 bool
			if x.FptrMapUintUint == nil {
				yyn782 = true
				goto LABEL782
			}
		LABEL782:
			if yyr2 || yy2arr2 {
				if yyn782 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintUint == nil {
						r.EncodeNil()
					} else {
						yy783 := *x.FptrMapUintUint
						yym784 := z.EncBinary()
						_ = yym784
						if false {
						} else {
							z.F.EncMapUintUintV(yy783, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintUint"))
				r.WriteMapElemValue()
				if yyn782 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintUint == nil {
						r.EncodeNil()
					} else {
						yy785 := *x.FptrMapUintUint
						yym786 := z.EncBinary()
						_ = yym786
						if false {
						} else {
							z.F.EncMapUintUintV(yy785, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintUint8 == nil {
					r.EncodeNil()
				} else {
					yym788 := z.EncBinary()
					_ = yym788
					if false {
					} else {
						z.F.EncMapUintUint8V(x.FMapUintUint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintUint8"))
				r.WriteMapElemValue()
				if x.FMapUintUint8 == nil {
					r.EncodeNil()
				} else {
					yym789 := z.EncBinary()
					_ = yym789
					if false {
					} else {
						z.F.EncMapUintUint8V(x.FMapUintUint8, e)
					}
				}
			}
			var yyn790 bool
			if x.FptrMapUintUint8 == nil {
				yyn790 = true
				goto LABEL790
			}
		LABEL790:
			if yyr2 || yy2arr2 {
				if yyn790 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintUint8 == nil {
						r.EncodeNil()
					} else {
						yy791 := *x.FptrMapUintUint8
						yym792 := z.EncBinary()
						_ = yym792
						if false {
						} else {
							z.F.EncMapUintUint8V(yy791, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintUint8"))
				r.WriteMapElemValue()
				if yyn790 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintUint8 == nil {
						r.EncodeNil()
					} else {
						yy793 := *x.FptrMapUintUint8
						yym794 := z.EncBinary()
						_ = yym794
						if false {
						} else {
							z.F.EncMapUintUint8V(yy793, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintUint16 == nil {
					r.EncodeNil()
				} else {
					yym796 := z.EncBinary()
					_ = yym796
					if false {
					} else {
						z.F.EncMapUintUint16V(x.FMapUintUint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintUint16"))
				r.WriteMapElemValue()
				if x.FMapUintUint16 == nil {
					r.EncodeNil()
				} else {
					yym797 := z.EncBinary()
					_ = yym797
					if false {
					} else {
						z.F.EncMapUintUint16V(x.FMapUintUint16, e)
					}
				}
			}
			var yyn798 bool
			if x.FptrMapUintUint16 == nil {
				yyn798 = true
				goto LABEL798
			}
		LABEL798:
			if yyr2 || yy2arr2 {
				if yyn798 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintUint16 == nil {
						r.EncodeNil()
					} else {
						yy799 := *x.FptrMapUintUint16
						yym800 := z.EncBinary()
						_ = yym800
						if false {
						} else {
							z.F.EncMapUintUint16V(yy799, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintUint16"))
				r.WriteMapElemValue()
				if yyn798 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintUint16 == nil {
						r.EncodeNil()
					} else {
						yy801 := *x.FptrMapUintUint16
						yym802 := z.EncBinary()
						_ = yym802
						if false {
						} else {
							z.F.EncMapUintUint16V(yy801, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintUint32 == nil {
					r.EncodeNil()
				} else {
					yym804 := z.EncBinary()
					_ = yym804
					if false {
					} else {
						z.F.EncMapUintUint32V(x.FMapUintUint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintUint32"))
				r.WriteMapElemValue()
				if x.FMapUintUint32 == nil {
					r.EncodeNil()
				} else {
					yym805 := z.EncBinary()
					_ = yym805
					if false {
					} else {
						z.F.EncMapUintUint32V(x.FMapUintUint32, e)
					}
				}
			}
			var yyn806 bool
			if x.FptrMapUintUint32 == nil {
				yyn806 = true
				goto LABEL806
			}
		LABEL806:
			if yyr2 || yy2arr2 {
				if yyn806 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintUint32 == nil {
						r.EncodeNil()
					} else {
						yy807 := *x.FptrMapUintUint32
						yym808 := z.EncBinary()
						_ = yym808
						if false {
						} else {
							z.F.EncMapUintUint32V(yy807, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintUint32"))
				r.WriteMapElemValue()
				if yyn806 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintUint32 == nil {
						r.EncodeNil()
					} else {
						yy809 := *x.FptrMapUintUint32
						yym810 := z.EncBinary()
						_ = yym810
						if false {
						} else {
							z.F.EncMapUintUint32V(yy809, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintUint64 == nil {
					r.EncodeNil()
				} else {
					yym812 := z.EncBinary()
					_ = yym812
					if false {
					} else {
						z.F.EncMapUintUint64V(x.FMapUintUint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintUint64"))
				r.WriteMapElemValue()
				if x.FMapUintUint64 == nil {
					r.EncodeNil()
				} else {
					yym813 := z.EncBinary()
					_ = yym813
					if false {
					} else {
						z.F.EncMapUintUint64V(x.FMapUintUint64, e)
					}
				}
			}
			var yyn814 bool
			if x.FptrMapUintUint64 == nil {
				yyn814 = true
				goto LABEL814
			}
		LABEL814:
			if yyr2 || yy2arr2 {
				if yyn814 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintUint64 == nil {
						r.EncodeNil()
					} else {
						yy815 := *x.FptrMapUintUint64
						yym816 := z.EncBinary()
						_ = yym816
						if false {
						} else {
							z.F.EncMapUintUint64V(yy815, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintUint64"))
				r.WriteMapElemValue()
				if yyn814 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintUint64 == nil {
						r.EncodeNil()
					} else {
						yy817 := *x.FptrMapUintUint64
						yym818 := z.EncBinary()
						_ = yym818
						if false {
						} else {
							z.F.EncMapUintUint64V(yy817, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintUintptr == nil {
					r.EncodeNil()
				} else {
					yym820 := z.EncBinary()
					_ = yym820
					if false {
					} else {
						z.F.EncMapUintUintptrV(x.FMapUintUintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintUintptr"))
				r.WriteMapElemValue()
				if x.FMapUintUintptr == nil {
					r.EncodeNil()
				} else {
					yym821 := z.EncBinary()
					_ = yym821
					if false {
					} else {
						z.F.EncMapUintUintptrV(x.FMapUintUintptr, e)
					}
				}
			}
			var yyn822 bool
			if x.FptrMapUintUintptr == nil {
				yyn822 = true
				goto LABEL822
			}
		LABEL822:
			if yyr2 || yy2arr2 {
				if yyn822 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintUintptr == nil {
						r.EncodeNil()
					} else {
						yy823 := *x.FptrMapUintUintptr
						yym824 := z.EncBinary()
						_ = yym824
						if false {
						} else {
							z.F.EncMapUintUintptrV(yy823, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintUintptr"))
				r.WriteMapElemValue()
				if yyn822 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintUintptr == nil {
						r.EncodeNil()
					} else {
						yy825 := *x.FptrMapUintUintptr
						yym826 := z.EncBinary()
						_ = yym826
						if false {
						} else {
							z.F.EncMapUintUintptrV(yy825, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintInt == nil {
					r.EncodeNil()
				} else {
					yym828 := z.EncBinary()
					_ = yym828
					if false {
					} else {
						z.F.EncMapUintIntV(x.FMapUintInt, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintInt"))
				r.WriteMapElemValue()
				if x.FMapUintInt == nil {
					r.EncodeNil()
				} else {
					yym829 := z.EncBinary()
					_ = yym829
					if false {
					} else {
						z.F.EncMapUintIntV(x.FMapUintInt, e)
					}
				}
			}
			var yyn830 bool
			if x.FptrMapUintInt == nil {
				yyn830 = true
				goto LABEL830
			}
		LABEL830:
			if yyr2 || yy2arr2 {
				if yyn830 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintInt == nil {
						r.EncodeNil()
					} else {
						yy831 := *x.FptrMapUintInt
						yym832 := z.EncBinary()
						_ = yym832
						if false {
						} else {
							z.F.EncMapUintIntV(yy831, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintInt"))
				r.WriteMapElemValue()
				if yyn830 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintInt == nil {
						r.EncodeNil()
					} else {
						yy833 := *x.FptrMapUintInt
						yym834 := z.EncBinary()
						_ = yym834
						if false {
						} else {
							z.F.EncMapUintIntV(yy833, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintInt8 == nil {
					r.EncodeNil()
				} else {
					yym836 := z.EncBinary()
					_ = yym836
					if false {
					} else {
						z.F.EncMapUintInt8V(x.FMapUintInt8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintInt8"))
				r.WriteMapElemValue()
				if x.FMapUintInt8 == nil {
					r.EncodeNil()
				} else {
					yym837 := z.EncBinary()
					_ = yym837
					if false {
					} else {
						z.F.EncMapUintInt8V(x.FMapUintInt8, e)
					}
				}
			}
			var yyn838 bool
			if x.FptrMapUintInt8 == nil {
				yyn838 = true
				goto LABEL838
			}
		LABEL838:
			if yyr2 || yy2arr2 {
				if yyn838 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintInt8 == nil {
						r.EncodeNil()
					} else {
						yy839 := *x.FptrMapUintInt8
						yym840 := z.EncBinary()
						_ = yym840
						if false {
						} else {
							z.F.EncMapUintInt8V(yy839, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintInt8"))
				r.WriteMapElemValue()
				if yyn838 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintInt8 == nil {
						r.EncodeNil()
					} else {
						yy841 := *x.FptrMapUintInt8
						yym842 := z.EncBinary()
						_ = yym842
						if false {
						} else {
							z.F.EncMapUintInt8V(yy841, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintInt16 == nil {
					r.EncodeNil()
				} else {
					yym844 := z.EncBinary()
					_ = yym844
					if false {
					} else {
						z.F.EncMapUintInt16V(x.FMapUintInt16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintInt16"))
				r.WriteMapElemValue()
				if x.FMapUintInt16 == nil {
					r.EncodeNil()
				} else {
					yym845 := z.EncBinary()
					_ = yym845
					if false {
					} else {
						z.F.EncMapUintInt16V(x.FMapUintInt16, e)
					}
				}
			}
			var yyn846 bool
			if x.FptrMapUintInt16 == nil {
				yyn846 = true
				goto LABEL846
			}
		LABEL846:
			if yyr2 || yy2arr2 {
				if yyn846 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintInt16 == nil {
						r.EncodeNil()
					} else {
						yy847 := *x.FptrMapUintInt16
						yym848 := z.EncBinary()
						_ = yym848
						if false {
						} else {
							z.F.EncMapUintInt16V(yy847, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintInt16"))
				r.WriteMapElemValue()
				if yyn846 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintInt16 == nil {
						r.EncodeNil()
					} else {
						yy849 := *x.FptrMapUintInt16
						yym850 := z.EncBinary()
						_ = yym850
						if false {
						} else {
							z.F.EncMapUintInt16V(yy849, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintInt32 == nil {
					r.EncodeNil()
				} else {
					yym852 := z.EncBinary()
					_ = yym852
					if false {
					} else {
						z.F.EncMapUintInt32V(x.FMapUintInt32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintInt32"))
				r.WriteMapElemValue()
				if x.FMapUintInt32 == nil {
					r.EncodeNil()
				} else {
					yym853 := z.EncBinary()
					_ = yym853
					if false {
					} else {
						z.F.EncMapUintInt32V(x.FMapUintInt32, e)
					}
				}
			}
			var yyn854 bool
			if x.FptrMapUintInt32 == nil {
				yyn854 = true
				goto LABEL854
			}
		LABEL854:
			if yyr2 || yy2arr2 {
				if yyn854 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintInt32 == nil {
						r.EncodeNil()
					} else {
						yy855 := *x.FptrMapUintInt32
						yym856 := z.EncBinary()
						_ = yym856
						if false {
						} else {
							z.F.EncMapUintInt32V(yy855, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintInt32"))
				r.WriteMapElemValue()
				if yyn854 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintInt32 == nil {
						r.EncodeNil()
					} else {
						yy857 := *x.FptrMapUintInt32
						yym858 := z.EncBinary()
						_ = yym858
						if false {
						} else {
							z.F.EncMapUintInt32V(yy857, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintInt64 == nil {
					r.EncodeNil()
				} else {
					yym860 := z.EncBinary()
					_ = yym860
					if false {
					} else {
						z.F.EncMapUintInt64V(x.FMapUintInt64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintInt64"))
				r.WriteMapElemValue()
				if x.FMapUintInt64 == nil {
					r.EncodeNil()
				} else {
					yym861 := z.EncBinary()
					_ = yym861
					if false {
					} else {
						z.F.EncMapUintInt64V(x.FMapUintInt64, e)
					}
				}
			}
			var yyn862 bool
			if x.FptrMapUintInt64 == nil {
				yyn862 = true
				goto LABEL862
			}
		LABEL862:
			if yyr2 || yy2arr2 {
				if yyn862 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintInt64 == nil {
						r.EncodeNil()
					} else {
						yy863 := *x.FptrMapUintInt64
						yym864 := z.EncBinary()
						_ = yym864
						if false {
						} else {
							z.F.EncMapUintInt64V(yy863, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintInt64"))
				r.WriteMapElemValue()
				if yyn862 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintInt64 == nil {
						r.EncodeNil()
					} else {
						yy865 := *x.FptrMapUintInt64
						yym866 := z.EncBinary()
						_ = yym866
						if false {
						} else {
							z.F.EncMapUintInt64V(yy865, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintFloat32 == nil {
					r.EncodeNil()
				} else {
					yym868 := z.EncBinary()
					_ = yym868
					if false {
					} else {
						z.F.EncMapUintFloat32V(x.FMapUintFloat32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintFloat32"))
				r.WriteMapElemValue()
				if x.FMapUintFloat32 == nil {
					r.EncodeNil()
				} else {
					yym869 := z.EncBinary()
					_ = yym869
					if false {
					} else {
						z.F.EncMapUintFloat32V(x.FMapUintFloat32, e)
					}
				}
			}
			var yyn870 bool
			if x.FptrMapUintFloat32 == nil {
				yyn870 = true
				goto LABEL870
			}
		LABEL870:
			if yyr2 || yy2arr2 {
				if yyn870 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintFloat32 == nil {
						r.EncodeNil()
					} else {
						yy871 := *x.FptrMapUintFloat32
						yym872 := z.EncBinary()
						_ = yym872
						if false {
						} else {
							z.F.EncMapUintFloat32V(yy871, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintFloat32"))
				r.WriteMapElemValue()
				if yyn870 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintFloat32 == nil {
						r.EncodeNil()
					} else {
						yy873 := *x.FptrMapUintFloat32
						yym874 := z.EncBinary()
						_ = yym874
						if false {
						} else {
							z.F.EncMapUintFloat32V(yy873, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintFloat64 == nil {
					r.EncodeNil()
				} else {
					yym876 := z.EncBinary()
					_ = yym876
					if false {
					} else {
						z.F.EncMapUintFloat64V(x.FMapUintFloat64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintFloat64"))
				r.WriteMapElemValue()
				if x.FMapUintFloat64 == nil {
					r.EncodeNil()
				} else {
					yym877 := z.EncBinary()
					_ = yym877
					if false {
					} else {
						z.F.EncMapUintFloat64V(x.FMapUintFloat64, e)
					}
				}
			}
			var yyn878 bool
			if x.FptrMapUintFloat64 == nil {
				yyn878 = true
				goto LABEL878
			}
		LABEL878:
			if yyr2 || yy2arr2 {
				if yyn878 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintFloat64 == nil {
						r.EncodeNil()
					} else {
						yy879 := *x.FptrMapUintFloat64
						yym880 := z.EncBinary()
						_ = yym880
						if false {
						} else {
							z.F.EncMapUintFloat64V(yy879, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintFloat64"))
				r.WriteMapElemValue()
				if yyn878 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintFloat64 == nil {
						r.EncodeNil()
					} else {
						yy881 := *x.FptrMapUintFloat64
						yym882 := z.EncBinary()
						_ = yym882
						if false {
						} else {
							z.F.EncMapUintFloat64V(yy881, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintBool == nil {
					r.EncodeNil()
				} else {
					yym884 := z.EncBinary()
					_ = yym884
					if false {
					} else {
						z.F.EncMapUintBoolV(x.FMapUintBool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintBool"))
				r.WriteMapElemValue()
				if x.FMapUintBool == nil {
					r.EncodeNil()
				} else {
					yym885 := z.EncBinary()
					_ = yym885
					if false {
					} else {
						z.F.EncMapUintBoolV(x.FMapUintBool, e)
					}
				}
			}
			var yyn886 bool
			if x.FptrMapUintBool == nil {
				yyn886 = true
				goto LABEL886
			}
		LABEL886:
			if yyr2 || yy2arr2 {
				if yyn886 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintBool == nil {
						r.EncodeNil()
					} else {
						yy887 := *x.FptrMapUintBool
						yym888 := z.EncBinary()
						_ = yym888
						if false {
						} else {
							z.F.EncMapUintBoolV(yy887, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintBool"))
				r.WriteMapElemValue()
				if yyn886 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintBool == nil {
						r.EncodeNil()
					} else {
						yy889 := *x.FptrMapUintBool
						yym890 := z.EncBinary()
						_ = yym890
						if false {
						} else {
							z.F.EncMapUintBoolV(yy889, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Intf == nil {
					r.EncodeNil()
				} else {
					yym892 := z.EncBinary()
					_ = yym892
					if false {
					} else {
						z.F.EncMapUint8IntfV(x.FMapUint8Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Intf"))
				r.WriteMapElemValue()
				if x.FMapUint8Intf == nil {
					r.EncodeNil()
				} else {
					yym893 := z.EncBinary()
					_ = yym893
					if false {
					} else {
						z.F.EncMapUint8IntfV(x.FMapUint8Intf, e)
					}
				}
			}
			var yyn894 bool
			if x.FptrMapUint8Intf == nil {
				yyn894 = true
				goto LABEL894
			}
		LABEL894:
			if yyr2 || yy2arr2 {
				if yyn894 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Intf == nil {
						r.EncodeNil()
					} else {
						yy895 := *x.FptrMapUint8Intf
						yym896 := z.EncBinary()
						_ = yym896
						if false {
						} else {
							z.F.EncMapUint8IntfV(yy895, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Intf"))
				r.WriteMapElemValue()
				if yyn894 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Intf == nil {
						r.EncodeNil()
					} else {
						yy897 := *x.FptrMapUint8Intf
						yym898 := z.EncBinary()
						_ = yym898
						if false {
						} else {
							z.F.EncMapUint8IntfV(yy897, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8String == nil {
					r.EncodeNil()
				} else {
					yym900 := z.EncBinary()
					_ = yym900
					if false {
					} else {
						z.F.EncMapUint8StringV(x.FMapUint8String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8String"))
				r.WriteMapElemValue()
				if x.FMapUint8String == nil {
					r.EncodeNil()
				} else {
					yym901 := z.EncBinary()
					_ = yym901
					if false {
					} else {
						z.F.EncMapUint8StringV(x.FMapUint8String, e)
					}
				}
			}
			var yyn902 bool
			if x.FptrMapUint8String == nil {
				yyn902 = true
				goto LABEL902
			}
		LABEL902:
			if yyr2 || yy2arr2 {
				if yyn902 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8String == nil {
						r.EncodeNil()
					} else {
						yy903 := *x.FptrMapUint8String
						yym904 := z.EncBinary()
						_ = yym904
						if false {
						} else {
							z.F.EncMapUint8StringV(yy903, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8String"))
				r.WriteMapElemValue()
				if yyn902 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8String == nil {
						r.EncodeNil()
					} else {
						yy905 := *x.FptrMapUint8String
						yym906 := z.EncBinary()
						_ = yym906
						if false {
						} else {
							z.F.EncMapUint8StringV(yy905, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Uint == nil {
					r.EncodeNil()
				} else {
					yym908 := z.EncBinary()
					_ = yym908
					if false {
					} else {
						z.F.EncMapUint8UintV(x.FMapUint8Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Uint"))
				r.WriteMapElemValue()
				if x.FMapUint8Uint == nil {
					r.EncodeNil()
				} else {
					yym909 := z.EncBinary()
					_ = yym909
					if false {
					} else {
						z.F.EncMapUint8UintV(x.FMapUint8Uint, e)
					}
				}
			}
			var yyn910 bool
			if x.FptrMapUint8Uint == nil {
				yyn910 = true
				goto LABEL910
			}
		LABEL910:
			if yyr2 || yy2arr2 {
				if yyn910 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Uint == nil {
						r.EncodeNil()
					} else {
						yy911 := *x.FptrMapUint8Uint
						yym912 := z.EncBinary()
						_ = yym912
						if false {
						} else {
							z.F.EncMapUint8UintV(yy911, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Uint"))
				r.WriteMapElemValue()
				if yyn910 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Uint == nil {
						r.EncodeNil()
					} else {
						yy913 := *x.FptrMapUint8Uint
						yym914 := z.EncBinary()
						_ = yym914
						if false {
						} else {
							z.F.EncMapUint8UintV(yy913, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Uint8 == nil {
					r.EncodeNil()
				} else {
					yym916 := z.EncBinary()
					_ = yym916
					if false {
					} else {
						z.F.EncMapUint8Uint8V(x.FMapUint8Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Uint8"))
				r.WriteMapElemValue()
				if x.FMapUint8Uint8 == nil {
					r.EncodeNil()
				} else {
					yym917 := z.EncBinary()
					_ = yym917
					if false {
					} else {
						z.F.EncMapUint8Uint8V(x.FMapUint8Uint8, e)
					}
				}
			}
			var yyn918 bool
			if x.FptrMapUint8Uint8 == nil {
				yyn918 = true
				goto LABEL918
			}
		LABEL918:
			if yyr2 || yy2arr2 {
				if yyn918 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Uint8 == nil {
						r.EncodeNil()
					} else {
						yy919 := *x.FptrMapUint8Uint8
						yym920 := z.EncBinary()
						_ = yym920
						if false {
						} else {
							z.F.EncMapUint8Uint8V(yy919, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Uint8"))
				r.WriteMapElemValue()
				if yyn918 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Uint8 == nil {
						r.EncodeNil()
					} else {
						yy921 := *x.FptrMapUint8Uint8
						yym922 := z.EncBinary()
						_ = yym922
						if false {
						} else {
							z.F.EncMapUint8Uint8V(yy921, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Uint16 == nil {
					r.EncodeNil()
				} else {
					yym924 := z.EncBinary()
					_ = yym924
					if false {
					} else {
						z.F.EncMapUint8Uint16V(x.FMapUint8Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Uint16"))
				r.WriteMapElemValue()
				if x.FMapUint8Uint16 == nil {
					r.EncodeNil()
				} else {
					yym925 := z.EncBinary()
					_ = yym925
					if false {
					} else {
						z.F.EncMapUint8Uint16V(x.FMapUint8Uint16, e)
					}
				}
			}
			var yyn926 bool
			if x.FptrMapUint8Uint16 == nil {
				yyn926 = true
				goto LABEL926
			}
		LABEL926:
			if yyr2 || yy2arr2 {
				if yyn926 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Uint16 == nil {
						r.EncodeNil()
					} else {
						yy927 := *x.FptrMapUint8Uint16
						yym928 := z.EncBinary()
						_ = yym928
						if false {
						} else {
							z.F.EncMapUint8Uint16V(yy927, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Uint16"))
				r.WriteMapElemValue()
				if yyn926 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Uint16 == nil {
						r.EncodeNil()
					} else {
						yy929 := *x.FptrMapUint8Uint16
						yym930 := z.EncBinary()
						_ = yym930
						if false {
						} else {
							z.F.EncMapUint8Uint16V(yy929, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Uint32 == nil {
					r.EncodeNil()
				} else {
					yym932 := z.EncBinary()
					_ = yym932
					if false {
					} else {
						z.F.EncMapUint8Uint32V(x.FMapUint8Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Uint32"))
				r.WriteMapElemValue()
				if x.FMapUint8Uint32 == nil {
					r.EncodeNil()
				} else {
					yym933 := z.EncBinary()
					_ = yym933
					if false {
					} else {
						z.F.EncMapUint8Uint32V(x.FMapUint8Uint32, e)
					}
				}
			}
			var yyn934 bool
			if x.FptrMapUint8Uint32 == nil {
				yyn934 = true
				goto LABEL934
			}
		LABEL934:
			if yyr2 || yy2arr2 {
				if yyn934 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Uint32 == nil {
						r.EncodeNil()
					} else {
						yy935 := *x.FptrMapUint8Uint32
						yym936 := z.EncBinary()
						_ = yym936
						if false {
						} else {
							z.F.EncMapUint8Uint32V(yy935, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Uint32"))
				r.WriteMapElemValue()
				if yyn934 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Uint32 == nil {
						r.EncodeNil()
					} else {
						yy937 := *x.FptrMapUint8Uint32
						yym938 := z.EncBinary()
						_ = yym938
						if false {
						} else {
							z.F.EncMapUint8Uint32V(yy937, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Uint64 == nil {
					r.EncodeNil()
				} else {
					yym940 := z.EncBinary()
					_ = yym940
					if false {
					} else {
						z.F.EncMapUint8Uint64V(x.FMapUint8Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Uint64"))
				r.WriteMapElemValue()
				if x.FMapUint8Uint64 == nil {
					r.EncodeNil()
				} else {
					yym941 := z.EncBinary()
					_ = yym941
					if false {
					} else {
						z.F.EncMapUint8Uint64V(x.FMapUint8Uint64, e)
					}
				}
			}
			var yyn942 bool
			if x.FptrMapUint8Uint64 == nil {
				yyn942 = true
				goto LABEL942
			}
		LABEL942:
			if yyr2 || yy2arr2 {
				if yyn942 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Uint64 == nil {
						r.EncodeNil()
					} else {
						yy943 := *x.FptrMapUint8Uint64
						yym944 := z.EncBinary()
						_ = yym944
						if false {
						} else {
							z.F.EncMapUint8Uint64V(yy943, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Uint64"))
				r.WriteMapElemValue()
				if yyn942 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Uint64 == nil {
						r.EncodeNil()
					} else {
						yy945 := *x.FptrMapUint8Uint64
						yym946 := z.EncBinary()
						_ = yym946
						if false {
						} else {
							z.F.EncMapUint8Uint64V(yy945, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Uintptr == nil {
					r.EncodeNil()
				} else {
					yym948 := z.EncBinary()
					_ = yym948
					if false {
					} else {
						z.F.EncMapUint8UintptrV(x.FMapUint8Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Uintptr"))
				r.WriteMapElemValue()
				if x.FMapUint8Uintptr == nil {
					r.EncodeNil()
				} else {
					yym949 := z.EncBinary()
					_ = yym949
					if false {
					} else {
						z.F.EncMapUint8UintptrV(x.FMapUint8Uintptr, e)
					}
				}
			}
			var yyn950 bool
			if x.FptrMapUint8Uintptr == nil {
				yyn950 = true
				goto LABEL950
			}
		LABEL950:
			if yyr2 || yy2arr2 {
				if yyn950 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Uintptr == nil {
						r.EncodeNil()
					} else {
						yy951 := *x.FptrMapUint8Uintptr
						yym952 := z.EncBinary()
						_ = yym952
						if false {
						} else {
							z.F.EncMapUint8UintptrV(yy951, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Uintptr"))
				r.WriteMapElemValue()
				if yyn950 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Uintptr == nil {
						r.EncodeNil()
					} else {
						yy953 := *x.FptrMapUint8Uintptr
						yym954 := z.EncBinary()
						_ = yym954
						if false {
						} else {
							z.F.EncMapUint8UintptrV(yy953, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Int == nil {
					r.EncodeNil()
				} else {
					yym956 := z.EncBinary()
					_ = yym956
					if false {
					} else {
						z.F.EncMapUint8IntV(x.FMapUint8Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Int"))
				r.WriteMapElemValue()
				if x.FMapUint8Int == nil {
					r.EncodeNil()
				} else {
					yym957 := z.EncBinary()
					_ = yym957
					if false {
					} else {
						z.F.EncMapUint8IntV(x.FMapUint8Int, e)
					}
				}
			}
			var yyn958 bool
			if x.FptrMapUint8Int == nil {
				yyn958 = true
				goto LABEL958
			}
		LABEL958:
			if yyr2 || yy2arr2 {
				if yyn958 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Int == nil {
						r.EncodeNil()
					} else {
						yy959 := *x.FptrMapUint8Int
						yym960 := z.EncBinary()
						_ = yym960
						if false {
						} else {
							z.F.EncMapUint8IntV(yy959, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Int"))
				r.WriteMapElemValue()
				if yyn958 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Int == nil {
						r.EncodeNil()
					} else {
						yy961 := *x.FptrMapUint8Int
						yym962 := z.EncBinary()
						_ = yym962
						if false {
						} else {
							z.F.EncMapUint8IntV(yy961, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Int8 == nil {
					r.EncodeNil()
				} else {
					yym964 := z.EncBinary()
					_ = yym964
					if false {
					} else {
						z.F.EncMapUint8Int8V(x.FMapUint8Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Int8"))
				r.WriteMapElemValue()
				if x.FMapUint8Int8 == nil {
					r.EncodeNil()
				} else {
					yym965 := z.EncBinary()
					_ = yym965
					if false {
					} else {
						z.F.EncMapUint8Int8V(x.FMapUint8Int8, e)
					}
				}
			}
			var yyn966 bool
			if x.FptrMapUint8Int8 == nil {
				yyn966 = true
				goto LABEL966
			}
		LABEL966:
			if yyr2 || yy2arr2 {
				if yyn966 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Int8 == nil {
						r.EncodeNil()
					} else {
						yy967 := *x.FptrMapUint8Int8
						yym968 := z.EncBinary()
						_ = yym968
						if false {
						} else {
							z.F.EncMapUint8Int8V(yy967, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Int8"))
				r.WriteMapElemValue()
				if yyn966 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Int8 == nil {
						r.EncodeNil()
					} else {
						yy969 := *x.FptrMapUint8Int8
						yym970 := z.EncBinary()
						_ = yym970
						if false {
						} else {
							z.F.EncMapUint8Int8V(yy969, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Int16 == nil {
					r.EncodeNil()
				} else {
					yym972 := z.EncBinary()
					_ = yym972
					if false {
					} else {
						z.F.EncMapUint8Int16V(x.FMapUint8Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Int16"))
				r.WriteMapElemValue()
				if x.FMapUint8Int16 == nil {
					r.EncodeNil()
				} else {
					yym973 := z.EncBinary()
					_ = yym973
					if false {
					} else {
						z.F.EncMapUint8Int16V(x.FMapUint8Int16, e)
					}
				}
			}
			var yyn974 bool
			if x.FptrMapUint8Int16 == nil {
				yyn974 = true
				goto LABEL974
			}
		LABEL974:
			if yyr2 || yy2arr2 {
				if yyn974 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Int16 == nil {
						r.EncodeNil()
					} else {
						yy975 := *x.FptrMapUint8Int16
						yym976 := z.EncBinary()
						_ = yym976
						if false {
						} else {
							z.F.EncMapUint8Int16V(yy975, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Int16"))
				r.WriteMapElemValue()
				if yyn974 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Int16 == nil {
						r.EncodeNil()
					} else {
						yy977 := *x.FptrMapUint8Int16
						yym978 := z.EncBinary()
						_ = yym978
						if false {
						} else {
							z.F.EncMapUint8Int16V(yy977, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Int32 == nil {
					r.EncodeNil()
				} else {
					yym980 := z.EncBinary()
					_ = yym980
					if false {
					} else {
						z.F.EncMapUint8Int32V(x.FMapUint8Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Int32"))
				r.WriteMapElemValue()
				if x.FMapUint8Int32 == nil {
					r.EncodeNil()
				} else {
					yym981 := z.EncBinary()
					_ = yym981
					if false {
					} else {
						z.F.EncMapUint8Int32V(x.FMapUint8Int32, e)
					}
				}
			}
			var yyn982 bool
			if x.FptrMapUint8Int32 == nil {
				yyn982 = true
				goto LABEL982
			}
		LABEL982:
			if yyr2 || yy2arr2 {
				if yyn982 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Int32 == nil {
						r.EncodeNil()
					} else {
						yy983 := *x.FptrMapUint8Int32
						yym984 := z.EncBinary()
						_ = yym984
						if false {
						} else {
							z.F.EncMapUint8Int32V(yy983, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Int32"))
				r.WriteMapElemValue()
				if yyn982 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Int32 == nil {
						r.EncodeNil()
					} else {
						yy985 := *x.FptrMapUint8Int32
						yym986 := z.EncBinary()
						_ = yym986
						if false {
						} else {
							z.F.EncMapUint8Int32V(yy985, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Int64 == nil {
					r.EncodeNil()
				} else {
					yym988 := z.EncBinary()
					_ = yym988
					if false {
					} else {
						z.F.EncMapUint8Int64V(x.FMapUint8Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Int64"))
				r.WriteMapElemValue()
				if x.FMapUint8Int64 == nil {
					r.EncodeNil()
				} else {
					yym989 := z.EncBinary()
					_ = yym989
					if false {
					} else {
						z.F.EncMapUint8Int64V(x.FMapUint8Int64, e)
					}
				}
			}
			var yyn990 bool
			if x.FptrMapUint8Int64 == nil {
				yyn990 = true
				goto LABEL990
			}
		LABEL990:
			if yyr2 || yy2arr2 {
				if yyn990 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Int64 == nil {
						r.EncodeNil()
					} else {
						yy991 := *x.FptrMapUint8Int64
						yym992 := z.EncBinary()
						_ = yym992
						if false {
						} else {
							z.F.EncMapUint8Int64V(yy991, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Int64"))
				r.WriteMapElemValue()
				if yyn990 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Int64 == nil {
						r.EncodeNil()
					} else {
						yy993 := *x.FptrMapUint8Int64
						yym994 := z.EncBinary()
						_ = yym994
						if false {
						} else {
							z.F.EncMapUint8Int64V(yy993, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Float32 == nil {
					r.EncodeNil()
				} else {
					yym996 := z.EncBinary()
					_ = yym996
					if false {
					} else {
						z.F.EncMapUint8Float32V(x.FMapUint8Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Float32"))
				r.WriteMapElemValue()
				if x.FMapUint8Float32 == nil {
					r.EncodeNil()
				} else {
					yym997 := z.EncBinary()
					_ = yym997
					if false {
					} else {
						z.F.EncMapUint8Float32V(x.FMapUint8Float32, e)
					}
				}
			}
			var yyn998 bool
			if x.FptrMapUint8Float32 == nil {
				yyn998 = true
				goto LABEL998
			}
		LABEL998:
			if yyr2 || yy2arr2 {
				if yyn998 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Float32 == nil {
						r.EncodeNil()
					} else {
						yy999 := *x.FptrMapUint8Float32
						yym1000 := z.EncBinary()
						_ = yym1000
						if false {
						} else {
							z.F.EncMapUint8Float32V(yy999, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Float32"))
				r.WriteMapElemValue()
				if yyn998 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Float32 == nil {
						r.EncodeNil()
					} else {
						yy1001 := *x.FptrMapUint8Float32
						yym1002 := z.EncBinary()
						_ = yym1002
						if false {
						} else {
							z.F.EncMapUint8Float32V(yy1001, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Float64 == nil {
					r.EncodeNil()
				} else {
					yym1004 := z.EncBinary()
					_ = yym1004
					if false {
					} else {
						z.F.EncMapUint8Float64V(x.FMapUint8Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Float64"))
				r.WriteMapElemValue()
				if x.FMapUint8Float64 == nil {
					r.EncodeNil()
				} else {
					yym1005 := z.EncBinary()
					_ = yym1005
					if false {
					} else {
						z.F.EncMapUint8Float64V(x.FMapUint8Float64, e)
					}
				}
			}
			var yyn1006 bool
			if x.FptrMapUint8Float64 == nil {
				yyn1006 = true
				goto LABEL1006
			}
		LABEL1006:
			if yyr2 || yy2arr2 {
				if yyn1006 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Float64 == nil {
						r.EncodeNil()
					} else {
						yy1007 := *x.FptrMapUint8Float64
						yym1008 := z.EncBinary()
						_ = yym1008
						if false {
						} else {
							z.F.EncMapUint8Float64V(yy1007, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Float64"))
				r.WriteMapElemValue()
				if yyn1006 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Float64 == nil {
						r.EncodeNil()
					} else {
						yy1009 := *x.FptrMapUint8Float64
						yym1010 := z.EncBinary()
						_ = yym1010
						if false {
						} else {
							z.F.EncMapUint8Float64V(yy1009, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint8Bool == nil {
					r.EncodeNil()
				} else {
					yym1012 := z.EncBinary()
					_ = yym1012
					if false {
					} else {
						z.F.EncMapUint8BoolV(x.FMapUint8Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint8Bool"))
				r.WriteMapElemValue()
				if x.FMapUint8Bool == nil {
					r.EncodeNil()
				} else {
					yym1013 := z.EncBinary()
					_ = yym1013
					if false {
					} else {
						z.F.EncMapUint8BoolV(x.FMapUint8Bool, e)
					}
				}
			}
			var yyn1014 bool
			if x.FptrMapUint8Bool == nil {
				yyn1014 = true
				goto LABEL1014
			}
		LABEL1014:
			if yyr2 || yy2arr2 {
				if yyn1014 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint8Bool == nil {
						r.EncodeNil()
					} else {
						yy1015 := *x.FptrMapUint8Bool
						yym1016 := z.EncBinary()
						_ = yym1016
						if false {
						} else {
							z.F.EncMapUint8BoolV(yy1015, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint8Bool"))
				r.WriteMapElemValue()
				if yyn1014 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint8Bool == nil {
						r.EncodeNil()
					} else {
						yy1017 := *x.FptrMapUint8Bool
						yym1018 := z.EncBinary()
						_ = yym1018
						if false {
						} else {
							z.F.EncMapUint8BoolV(yy1017, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Intf == nil {
					r.EncodeNil()
				} else {
					yym1020 := z.EncBinary()
					_ = yym1020
					if false {
					} else {
						z.F.EncMapUint16IntfV(x.FMapUint16Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Intf"))
				r.WriteMapElemValue()
				if x.FMapUint16Intf == nil {
					r.EncodeNil()
				} else {
					yym1021 := z.EncBinary()
					_ = yym1021
					if false {
					} else {
						z.F.EncMapUint16IntfV(x.FMapUint16Intf, e)
					}
				}
			}
			var yyn1022 bool
			if x.FptrMapUint16Intf == nil {
				yyn1022 = true
				goto LABEL1022
			}
		LABEL1022:
			if yyr2 || yy2arr2 {
				if yyn1022 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Intf == nil {
						r.EncodeNil()
					} else {
						yy1023 := *x.FptrMapUint16Intf
						yym1024 := z.EncBinary()
						_ = yym1024
						if false {
						} else {
							z.F.EncMapUint16IntfV(yy1023, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Intf"))
				r.WriteMapElemValue()
				if yyn1022 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Intf == nil {
						r.EncodeNil()
					} else {
						yy1025 := *x.FptrMapUint16Intf
						yym1026 := z.EncBinary()
						_ = yym1026
						if false {
						} else {
							z.F.EncMapUint16IntfV(yy1025, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16String == nil {
					r.EncodeNil()
				} else {
					yym1028 := z.EncBinary()
					_ = yym1028
					if false {
					} else {
						z.F.EncMapUint16StringV(x.FMapUint16String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16String"))
				r.WriteMapElemValue()
				if x.FMapUint16String == nil {
					r.EncodeNil()
				} else {
					yym1029 := z.EncBinary()
					_ = yym1029
					if false {
					} else {
						z.F.EncMapUint16StringV(x.FMapUint16String, e)
					}
				}
			}
			var yyn1030 bool
			if x.FptrMapUint16String == nil {
				yyn1030 = true
				goto LABEL1030
			}
		LABEL1030:
			if yyr2 || yy2arr2 {
				if yyn1030 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16String == nil {
						r.EncodeNil()
					} else {
						yy1031 := *x.FptrMapUint16String
						yym1032 := z.EncBinary()
						_ = yym1032
						if false {
						} else {
							z.F.EncMapUint16StringV(yy1031, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16String"))
				r.WriteMapElemValue()
				if yyn1030 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16String == nil {
						r.EncodeNil()
					} else {
						yy1033 := *x.FptrMapUint16String
						yym1034 := z.EncBinary()
						_ = yym1034
						if false {
						} else {
							z.F.EncMapUint16StringV(yy1033, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Uint == nil {
					r.EncodeNil()
				} else {
					yym1036 := z.EncBinary()
					_ = yym1036
					if false {
					} else {
						z.F.EncMapUint16UintV(x.FMapUint16Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Uint"))
				r.WriteMapElemValue()
				if x.FMapUint16Uint == nil {
					r.EncodeNil()
				} else {
					yym1037 := z.EncBinary()
					_ = yym1037
					if false {
					} else {
						z.F.EncMapUint16UintV(x.FMapUint16Uint, e)
					}
				}
			}
			var yyn1038 bool
			if x.FptrMapUint16Uint == nil {
				yyn1038 = true
				goto LABEL1038
			}
		LABEL1038:
			if yyr2 || yy2arr2 {
				if yyn1038 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Uint == nil {
						r.EncodeNil()
					} else {
						yy1039 := *x.FptrMapUint16Uint
						yym1040 := z.EncBinary()
						_ = yym1040
						if false {
						} else {
							z.F.EncMapUint16UintV(yy1039, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Uint"))
				r.WriteMapElemValue()
				if yyn1038 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Uint == nil {
						r.EncodeNil()
					} else {
						yy1041 := *x.FptrMapUint16Uint
						yym1042 := z.EncBinary()
						_ = yym1042
						if false {
						} else {
							z.F.EncMapUint16UintV(yy1041, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1044 := z.EncBinary()
					_ = yym1044
					if false {
					} else {
						z.F.EncMapUint16Uint8V(x.FMapUint16Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Uint8"))
				r.WriteMapElemValue()
				if x.FMapUint16Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1045 := z.EncBinary()
					_ = yym1045
					if false {
					} else {
						z.F.EncMapUint16Uint8V(x.FMapUint16Uint8, e)
					}
				}
			}
			var yyn1046 bool
			if x.FptrMapUint16Uint8 == nil {
				yyn1046 = true
				goto LABEL1046
			}
		LABEL1046:
			if yyr2 || yy2arr2 {
				if yyn1046 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1047 := *x.FptrMapUint16Uint8
						yym1048 := z.EncBinary()
						_ = yym1048
						if false {
						} else {
							z.F.EncMapUint16Uint8V(yy1047, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Uint8"))
				r.WriteMapElemValue()
				if yyn1046 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1049 := *x.FptrMapUint16Uint8
						yym1050 := z.EncBinary()
						_ = yym1050
						if false {
						} else {
							z.F.EncMapUint16Uint8V(yy1049, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1052 := z.EncBinary()
					_ = yym1052
					if false {
					} else {
						z.F.EncMapUint16Uint16V(x.FMapUint16Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Uint16"))
				r.WriteMapElemValue()
				if x.FMapUint16Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1053 := z.EncBinary()
					_ = yym1053
					if false {
					} else {
						z.F.EncMapUint16Uint16V(x.FMapUint16Uint16, e)
					}
				}
			}
			var yyn1054 bool
			if x.FptrMapUint16Uint16 == nil {
				yyn1054 = true
				goto LABEL1054
			}
		LABEL1054:
			if yyr2 || yy2arr2 {
				if yyn1054 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1055 := *x.FptrMapUint16Uint16
						yym1056 := z.EncBinary()
						_ = yym1056
						if false {
						} else {
							z.F.EncMapUint16Uint16V(yy1055, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Uint16"))
				r.WriteMapElemValue()
				if yyn1054 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1057 := *x.FptrMapUint16Uint16
						yym1058 := z.EncBinary()
						_ = yym1058
						if false {
						} else {
							z.F.EncMapUint16Uint16V(yy1057, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1060 := z.EncBinary()
					_ = yym1060
					if false {
					} else {
						z.F.EncMapUint16Uint32V(x.FMapUint16Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Uint32"))
				r.WriteMapElemValue()
				if x.FMapUint16Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1061 := z.EncBinary()
					_ = yym1061
					if false {
					} else {
						z.F.EncMapUint16Uint32V(x.FMapUint16Uint32, e)
					}
				}
			}
			var yyn1062 bool
			if x.FptrMapUint16Uint32 == nil {
				yyn1062 = true
				goto LABEL1062
			}
		LABEL1062:
			if yyr2 || yy2arr2 {
				if yyn1062 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1063 := *x.FptrMapUint16Uint32
						yym1064 := z.EncBinary()
						_ = yym1064
						if false {
						} else {
							z.F.EncMapUint16Uint32V(yy1063, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Uint32"))
				r.WriteMapElemValue()
				if yyn1062 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1065 := *x.FptrMapUint16Uint32
						yym1066 := z.EncBinary()
						_ = yym1066
						if false {
						} else {
							z.F.EncMapUint16Uint32V(yy1065, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1068 := z.EncBinary()
					_ = yym1068
					if false {
					} else {
						z.F.EncMapUint16Uint64V(x.FMapUint16Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Uint64"))
				r.WriteMapElemValue()
				if x.FMapUint16Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1069 := z.EncBinary()
					_ = yym1069
					if false {
					} else {
						z.F.EncMapUint16Uint64V(x.FMapUint16Uint64, e)
					}
				}
			}
			var yyn1070 bool
			if x.FptrMapUint16Uint64 == nil {
				yyn1070 = true
				goto LABEL1070
			}
		LABEL1070:
			if yyr2 || yy2arr2 {
				if yyn1070 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1071 := *x.FptrMapUint16Uint64
						yym1072 := z.EncBinary()
						_ = yym1072
						if false {
						} else {
							z.F.EncMapUint16Uint64V(yy1071, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Uint64"))
				r.WriteMapElemValue()
				if yyn1070 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1073 := *x.FptrMapUint16Uint64
						yym1074 := z.EncBinary()
						_ = yym1074
						if false {
						} else {
							z.F.EncMapUint16Uint64V(yy1073, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1076 := z.EncBinary()
					_ = yym1076
					if false {
					} else {
						z.F.EncMapUint16UintptrV(x.FMapUint16Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Uintptr"))
				r.WriteMapElemValue()
				if x.FMapUint16Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1077 := z.EncBinary()
					_ = yym1077
					if false {
					} else {
						z.F.EncMapUint16UintptrV(x.FMapUint16Uintptr, e)
					}
				}
			}
			var yyn1078 bool
			if x.FptrMapUint16Uintptr == nil {
				yyn1078 = true
				goto LABEL1078
			}
		LABEL1078:
			if yyr2 || yy2arr2 {
				if yyn1078 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1079 := *x.FptrMapUint16Uintptr
						yym1080 := z.EncBinary()
						_ = yym1080
						if false {
						} else {
							z.F.EncMapUint16UintptrV(yy1079, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Uintptr"))
				r.WriteMapElemValue()
				if yyn1078 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1081 := *x.FptrMapUint16Uintptr
						yym1082 := z.EncBinary()
						_ = yym1082
						if false {
						} else {
							z.F.EncMapUint16UintptrV(yy1081, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Int == nil {
					r.EncodeNil()
				} else {
					yym1084 := z.EncBinary()
					_ = yym1084
					if false {
					} else {
						z.F.EncMapUint16IntV(x.FMapUint16Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Int"))
				r.WriteMapElemValue()
				if x.FMapUint16Int == nil {
					r.EncodeNil()
				} else {
					yym1085 := z.EncBinary()
					_ = yym1085
					if false {
					} else {
						z.F.EncMapUint16IntV(x.FMapUint16Int, e)
					}
				}
			}
			var yyn1086 bool
			if x.FptrMapUint16Int == nil {
				yyn1086 = true
				goto LABEL1086
			}
		LABEL1086:
			if yyr2 || yy2arr2 {
				if yyn1086 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Int == nil {
						r.EncodeNil()
					} else {
						yy1087 := *x.FptrMapUint16Int
						yym1088 := z.EncBinary()
						_ = yym1088
						if false {
						} else {
							z.F.EncMapUint16IntV(yy1087, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Int"))
				r.WriteMapElemValue()
				if yyn1086 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Int == nil {
						r.EncodeNil()
					} else {
						yy1089 := *x.FptrMapUint16Int
						yym1090 := z.EncBinary()
						_ = yym1090
						if false {
						} else {
							z.F.EncMapUint16IntV(yy1089, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Int8 == nil {
					r.EncodeNil()
				} else {
					yym1092 := z.EncBinary()
					_ = yym1092
					if false {
					} else {
						z.F.EncMapUint16Int8V(x.FMapUint16Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Int8"))
				r.WriteMapElemValue()
				if x.FMapUint16Int8 == nil {
					r.EncodeNil()
				} else {
					yym1093 := z.EncBinary()
					_ = yym1093
					if false {
					} else {
						z.F.EncMapUint16Int8V(x.FMapUint16Int8, e)
					}
				}
			}
			var yyn1094 bool
			if x.FptrMapUint16Int8 == nil {
				yyn1094 = true
				goto LABEL1094
			}
		LABEL1094:
			if yyr2 || yy2arr2 {
				if yyn1094 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Int8 == nil {
						r.EncodeNil()
					} else {
						yy1095 := *x.FptrMapUint16Int8
						yym1096 := z.EncBinary()
						_ = yym1096
						if false {
						} else {
							z.F.EncMapUint16Int8V(yy1095, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Int8"))
				r.WriteMapElemValue()
				if yyn1094 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Int8 == nil {
						r.EncodeNil()
					} else {
						yy1097 := *x.FptrMapUint16Int8
						yym1098 := z.EncBinary()
						_ = yym1098
						if false {
						} else {
							z.F.EncMapUint16Int8V(yy1097, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Int16 == nil {
					r.EncodeNil()
				} else {
					yym1100 := z.EncBinary()
					_ = yym1100
					if false {
					} else {
						z.F.EncMapUint16Int16V(x.FMapUint16Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Int16"))
				r.WriteMapElemValue()
				if x.FMapUint16Int16 == nil {
					r.EncodeNil()
				} else {
					yym1101 := z.EncBinary()
					_ = yym1101
					if false {
					} else {
						z.F.EncMapUint16Int16V(x.FMapUint16Int16, e)
					}
				}
			}
			var yyn1102 bool
			if x.FptrMapUint16Int16 == nil {
				yyn1102 = true
				goto LABEL1102
			}
		LABEL1102:
			if yyr2 || yy2arr2 {
				if yyn1102 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Int16 == nil {
						r.EncodeNil()
					} else {
						yy1103 := *x.FptrMapUint16Int16
						yym1104 := z.EncBinary()
						_ = yym1104
						if false {
						} else {
							z.F.EncMapUint16Int16V(yy1103, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Int16"))
				r.WriteMapElemValue()
				if yyn1102 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Int16 == nil {
						r.EncodeNil()
					} else {
						yy1105 := *x.FptrMapUint16Int16
						yym1106 := z.EncBinary()
						_ = yym1106
						if false {
						} else {
							z.F.EncMapUint16Int16V(yy1105, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Int32 == nil {
					r.EncodeNil()
				} else {
					yym1108 := z.EncBinary()
					_ = yym1108
					if false {
					} else {
						z.F.EncMapUint16Int32V(x.FMapUint16Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Int32"))
				r.WriteMapElemValue()
				if x.FMapUint16Int32 == nil {
					r.EncodeNil()
				} else {
					yym1109 := z.EncBinary()
					_ = yym1109
					if false {
					} else {
						z.F.EncMapUint16Int32V(x.FMapUint16Int32, e)
					}
				}
			}
			var yyn1110 bool
			if x.FptrMapUint16Int32 == nil {
				yyn1110 = true
				goto LABEL1110
			}
		LABEL1110:
			if yyr2 || yy2arr2 {
				if yyn1110 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Int32 == nil {
						r.EncodeNil()
					} else {
						yy1111 := *x.FptrMapUint16Int32
						yym1112 := z.EncBinary()
						_ = yym1112
						if false {
						} else {
							z.F.EncMapUint16Int32V(yy1111, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Int32"))
				r.WriteMapElemValue()
				if yyn1110 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Int32 == nil {
						r.EncodeNil()
					} else {
						yy1113 := *x.FptrMapUint16Int32
						yym1114 := z.EncBinary()
						_ = yym1114
						if false {
						} else {
							z.F.EncMapUint16Int32V(yy1113, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Int64 == nil {
					r.EncodeNil()
				} else {
					yym1116 := z.EncBinary()
					_ = yym1116
					if false {
					} else {
						z.F.EncMapUint16Int64V(x.FMapUint16Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Int64"))
				r.WriteMapElemValue()
				if x.FMapUint16Int64 == nil {
					r.EncodeNil()
				} else {
					yym1117 := z.EncBinary()
					_ = yym1117
					if false {
					} else {
						z.F.EncMapUint16Int64V(x.FMapUint16Int64, e)
					}
				}
			}
			var yyn1118 bool
			if x.FptrMapUint16Int64 == nil {
				yyn1118 = true
				goto LABEL1118
			}
		LABEL1118:
			if yyr2 || yy2arr2 {
				if yyn1118 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Int64 == nil {
						r.EncodeNil()
					} else {
						yy1119 := *x.FptrMapUint16Int64
						yym1120 := z.EncBinary()
						_ = yym1120
						if false {
						} else {
							z.F.EncMapUint16Int64V(yy1119, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Int64"))
				r.WriteMapElemValue()
				if yyn1118 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Int64 == nil {
						r.EncodeNil()
					} else {
						yy1121 := *x.FptrMapUint16Int64
						yym1122 := z.EncBinary()
						_ = yym1122
						if false {
						} else {
							z.F.EncMapUint16Int64V(yy1121, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Float32 == nil {
					r.EncodeNil()
				} else {
					yym1124 := z.EncBinary()
					_ = yym1124
					if false {
					} else {
						z.F.EncMapUint16Float32V(x.FMapUint16Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Float32"))
				r.WriteMapElemValue()
				if x.FMapUint16Float32 == nil {
					r.EncodeNil()
				} else {
					yym1125 := z.EncBinary()
					_ = yym1125
					if false {
					} else {
						z.F.EncMapUint16Float32V(x.FMapUint16Float32, e)
					}
				}
			}
			var yyn1126 bool
			if x.FptrMapUint16Float32 == nil {
				yyn1126 = true
				goto LABEL1126
			}
		LABEL1126:
			if yyr2 || yy2arr2 {
				if yyn1126 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Float32 == nil {
						r.EncodeNil()
					} else {
						yy1127 := *x.FptrMapUint16Float32
						yym1128 := z.EncBinary()
						_ = yym1128
						if false {
						} else {
							z.F.EncMapUint16Float32V(yy1127, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Float32"))
				r.WriteMapElemValue()
				if yyn1126 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Float32 == nil {
						r.EncodeNil()
					} else {
						yy1129 := *x.FptrMapUint16Float32
						yym1130 := z.EncBinary()
						_ = yym1130
						if false {
						} else {
							z.F.EncMapUint16Float32V(yy1129, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Float64 == nil {
					r.EncodeNil()
				} else {
					yym1132 := z.EncBinary()
					_ = yym1132
					if false {
					} else {
						z.F.EncMapUint16Float64V(x.FMapUint16Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Float64"))
				r.WriteMapElemValue()
				if x.FMapUint16Float64 == nil {
					r.EncodeNil()
				} else {
					yym1133 := z.EncBinary()
					_ = yym1133
					if false {
					} else {
						z.F.EncMapUint16Float64V(x.FMapUint16Float64, e)
					}
				}
			}
			var yyn1134 bool
			if x.FptrMapUint16Float64 == nil {
				yyn1134 = true
				goto LABEL1134
			}
		LABEL1134:
			if yyr2 || yy2arr2 {
				if yyn1134 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Float64 == nil {
						r.EncodeNil()
					} else {
						yy1135 := *x.FptrMapUint16Float64
						yym1136 := z.EncBinary()
						_ = yym1136
						if false {
						} else {
							z.F.EncMapUint16Float64V(yy1135, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Float64"))
				r.WriteMapElemValue()
				if yyn1134 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Float64 == nil {
						r.EncodeNil()
					} else {
						yy1137 := *x.FptrMapUint16Float64
						yym1138 := z.EncBinary()
						_ = yym1138
						if false {
						} else {
							z.F.EncMapUint16Float64V(yy1137, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint16Bool == nil {
					r.EncodeNil()
				} else {
					yym1140 := z.EncBinary()
					_ = yym1140
					if false {
					} else {
						z.F.EncMapUint16BoolV(x.FMapUint16Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint16Bool"))
				r.WriteMapElemValue()
				if x.FMapUint16Bool == nil {
					r.EncodeNil()
				} else {
					yym1141 := z.EncBinary()
					_ = yym1141
					if false {
					} else {
						z.F.EncMapUint16BoolV(x.FMapUint16Bool, e)
					}
				}
			}
			var yyn1142 bool
			if x.FptrMapUint16Bool == nil {
				yyn1142 = true
				goto LABEL1142
			}
		LABEL1142:
			if yyr2 || yy2arr2 {
				if yyn1142 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint16Bool == nil {
						r.EncodeNil()
					} else {
						yy1143 := *x.FptrMapUint16Bool
						yym1144 := z.EncBinary()
						_ = yym1144
						if false {
						} else {
							z.F.EncMapUint16BoolV(yy1143, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint16Bool"))
				r.WriteMapElemValue()
				if yyn1142 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint16Bool == nil {
						r.EncodeNil()
					} else {
						yy1145 := *x.FptrMapUint16Bool
						yym1146 := z.EncBinary()
						_ = yym1146
						if false {
						} else {
							z.F.EncMapUint16BoolV(yy1145, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Intf == nil {
					r.EncodeNil()
				} else {
					yym1148 := z.EncBinary()
					_ = yym1148
					if false {
					} else {
						z.F.EncMapUint32IntfV(x.FMapUint32Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Intf"))
				r.WriteMapElemValue()
				if x.FMapUint32Intf == nil {
					r.EncodeNil()
				} else {
					yym1149 := z.EncBinary()
					_ = yym1149
					if false {
					} else {
						z.F.EncMapUint32IntfV(x.FMapUint32Intf, e)
					}
				}
			}
			var yyn1150 bool
			if x.FptrMapUint32Intf == nil {
				yyn1150 = true
				goto LABEL1150
			}
		LABEL1150:
			if yyr2 || yy2arr2 {
				if yyn1150 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Intf == nil {
						r.EncodeNil()
					} else {
						yy1151 := *x.FptrMapUint32Intf
						yym1152 := z.EncBinary()
						_ = yym1152
						if false {
						} else {
							z.F.EncMapUint32IntfV(yy1151, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Intf"))
				r.WriteMapElemValue()
				if yyn1150 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Intf == nil {
						r.EncodeNil()
					} else {
						yy1153 := *x.FptrMapUint32Intf
						yym1154 := z.EncBinary()
						_ = yym1154
						if false {
						} else {
							z.F.EncMapUint32IntfV(yy1153, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32String == nil {
					r.EncodeNil()
				} else {
					yym1156 := z.EncBinary()
					_ = yym1156
					if false {
					} else {
						z.F.EncMapUint32StringV(x.FMapUint32String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32String"))
				r.WriteMapElemValue()
				if x.FMapUint32String == nil {
					r.EncodeNil()
				} else {
					yym1157 := z.EncBinary()
					_ = yym1157
					if false {
					} else {
						z.F.EncMapUint32StringV(x.FMapUint32String, e)
					}
				}
			}
			var yyn1158 bool
			if x.FptrMapUint32String == nil {
				yyn1158 = true
				goto LABEL1158
			}
		LABEL1158:
			if yyr2 || yy2arr2 {
				if yyn1158 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32String == nil {
						r.EncodeNil()
					} else {
						yy1159 := *x.FptrMapUint32String
						yym1160 := z.EncBinary()
						_ = yym1160
						if false {
						} else {
							z.F.EncMapUint32StringV(yy1159, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32String"))
				r.WriteMapElemValue()
				if yyn1158 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32String == nil {
						r.EncodeNil()
					} else {
						yy1161 := *x.FptrMapUint32String
						yym1162 := z.EncBinary()
						_ = yym1162
						if false {
						} else {
							z.F.EncMapUint32StringV(yy1161, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Uint == nil {
					r.EncodeNil()
				} else {
					yym1164 := z.EncBinary()
					_ = yym1164
					if false {
					} else {
						z.F.EncMapUint32UintV(x.FMapUint32Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Uint"))
				r.WriteMapElemValue()
				if x.FMapUint32Uint == nil {
					r.EncodeNil()
				} else {
					yym1165 := z.EncBinary()
					_ = yym1165
					if false {
					} else {
						z.F.EncMapUint32UintV(x.FMapUint32Uint, e)
					}
				}
			}
			var yyn1166 bool
			if x.FptrMapUint32Uint == nil {
				yyn1166 = true
				goto LABEL1166
			}
		LABEL1166:
			if yyr2 || yy2arr2 {
				if yyn1166 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Uint == nil {
						r.EncodeNil()
					} else {
						yy1167 := *x.FptrMapUint32Uint
						yym1168 := z.EncBinary()
						_ = yym1168
						if false {
						} else {
							z.F.EncMapUint32UintV(yy1167, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Uint"))
				r.WriteMapElemValue()
				if yyn1166 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Uint == nil {
						r.EncodeNil()
					} else {
						yy1169 := *x.FptrMapUint32Uint
						yym1170 := z.EncBinary()
						_ = yym1170
						if false {
						} else {
							z.F.EncMapUint32UintV(yy1169, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1172 := z.EncBinary()
					_ = yym1172
					if false {
					} else {
						z.F.EncMapUint32Uint8V(x.FMapUint32Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Uint8"))
				r.WriteMapElemValue()
				if x.FMapUint32Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1173 := z.EncBinary()
					_ = yym1173
					if false {
					} else {
						z.F.EncMapUint32Uint8V(x.FMapUint32Uint8, e)
					}
				}
			}
			var yyn1174 bool
			if x.FptrMapUint32Uint8 == nil {
				yyn1174 = true
				goto LABEL1174
			}
		LABEL1174:
			if yyr2 || yy2arr2 {
				if yyn1174 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1175 := *x.FptrMapUint32Uint8
						yym1176 := z.EncBinary()
						_ = yym1176
						if false {
						} else {
							z.F.EncMapUint32Uint8V(yy1175, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Uint8"))
				r.WriteMapElemValue()
				if yyn1174 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1177 := *x.FptrMapUint32Uint8
						yym1178 := z.EncBinary()
						_ = yym1178
						if false {
						} else {
							z.F.EncMapUint32Uint8V(yy1177, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1180 := z.EncBinary()
					_ = yym1180
					if false {
					} else {
						z.F.EncMapUint32Uint16V(x.FMapUint32Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Uint16"))
				r.WriteMapElemValue()
				if x.FMapUint32Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1181 := z.EncBinary()
					_ = yym1181
					if false {
					} else {
						z.F.EncMapUint32Uint16V(x.FMapUint32Uint16, e)
					}
				}
			}
			var yyn1182 bool
			if x.FptrMapUint32Uint16 == nil {
				yyn1182 = true
				goto LABEL1182
			}
		LABEL1182:
			if yyr2 || yy2arr2 {
				if yyn1182 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1183 := *x.FptrMapUint32Uint16
						yym1184 := z.EncBinary()
						_ = yym1184
						if false {
						} else {
							z.F.EncMapUint32Uint16V(yy1183, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Uint16"))
				r.WriteMapElemValue()
				if yyn1182 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1185 := *x.FptrMapUint32Uint16
						yym1186 := z.EncBinary()
						_ = yym1186
						if false {
						} else {
							z.F.EncMapUint32Uint16V(yy1185, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1188 := z.EncBinary()
					_ = yym1188
					if false {
					} else {
						z.F.EncMapUint32Uint32V(x.FMapUint32Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Uint32"))
				r.WriteMapElemValue()
				if x.FMapUint32Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1189 := z.EncBinary()
					_ = yym1189
					if false {
					} else {
						z.F.EncMapUint32Uint32V(x.FMapUint32Uint32, e)
					}
				}
			}
			var yyn1190 bool
			if x.FptrMapUint32Uint32 == nil {
				yyn1190 = true
				goto LABEL1190
			}
		LABEL1190:
			if yyr2 || yy2arr2 {
				if yyn1190 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1191 := *x.FptrMapUint32Uint32
						yym1192 := z.EncBinary()
						_ = yym1192
						if false {
						} else {
							z.F.EncMapUint32Uint32V(yy1191, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Uint32"))
				r.WriteMapElemValue()
				if yyn1190 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1193 := *x.FptrMapUint32Uint32
						yym1194 := z.EncBinary()
						_ = yym1194
						if false {
						} else {
							z.F.EncMapUint32Uint32V(yy1193, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1196 := z.EncBinary()
					_ = yym1196
					if false {
					} else {
						z.F.EncMapUint32Uint64V(x.FMapUint32Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Uint64"))
				r.WriteMapElemValue()
				if x.FMapUint32Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1197 := z.EncBinary()
					_ = yym1197
					if false {
					} else {
						z.F.EncMapUint32Uint64V(x.FMapUint32Uint64, e)
					}
				}
			}
			var yyn1198 bool
			if x.FptrMapUint32Uint64 == nil {
				yyn1198 = true
				goto LABEL1198
			}
		LABEL1198:
			if yyr2 || yy2arr2 {
				if yyn1198 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1199 := *x.FptrMapUint32Uint64
						yym1200 := z.EncBinary()
						_ = yym1200
						if false {
						} else {
							z.F.EncMapUint32Uint64V(yy1199, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Uint64"))
				r.WriteMapElemValue()
				if yyn1198 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1201 := *x.FptrMapUint32Uint64
						yym1202 := z.EncBinary()
						_ = yym1202
						if false {
						} else {
							z.F.EncMapUint32Uint64V(yy1201, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1204 := z.EncBinary()
					_ = yym1204
					if false {
					} else {
						z.F.EncMapUint32UintptrV(x.FMapUint32Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Uintptr"))
				r.WriteMapElemValue()
				if x.FMapUint32Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1205 := z.EncBinary()
					_ = yym1205
					if false {
					} else {
						z.F.EncMapUint32UintptrV(x.FMapUint32Uintptr, e)
					}
				}
			}
			var yyn1206 bool
			if x.FptrMapUint32Uintptr == nil {
				yyn1206 = true
				goto LABEL1206
			}
		LABEL1206:
			if yyr2 || yy2arr2 {
				if yyn1206 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1207 := *x.FptrMapUint32Uintptr
						yym1208 := z.EncBinary()
						_ = yym1208
						if false {
						} else {
							z.F.EncMapUint32UintptrV(yy1207, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Uintptr"))
				r.WriteMapElemValue()
				if yyn1206 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1209 := *x.FptrMapUint32Uintptr
						yym1210 := z.EncBinary()
						_ = yym1210
						if false {
						} else {
							z.F.EncMapUint32UintptrV(yy1209, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Int == nil {
					r.EncodeNil()
				} else {
					yym1212 := z.EncBinary()
					_ = yym1212
					if false {
					} else {
						z.F.EncMapUint32IntV(x.FMapUint32Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Int"))
				r.WriteMapElemValue()
				if x.FMapUint32Int == nil {
					r.EncodeNil()
				} else {
					yym1213 := z.EncBinary()
					_ = yym1213
					if false {
					} else {
						z.F.EncMapUint32IntV(x.FMapUint32Int, e)
					}
				}
			}
			var yyn1214 bool
			if x.FptrMapUint32Int == nil {
				yyn1214 = true
				goto LABEL1214
			}
		LABEL1214:
			if yyr2 || yy2arr2 {
				if yyn1214 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Int == nil {
						r.EncodeNil()
					} else {
						yy1215 := *x.FptrMapUint32Int
						yym1216 := z.EncBinary()
						_ = yym1216
						if false {
						} else {
							z.F.EncMapUint32IntV(yy1215, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Int"))
				r.WriteMapElemValue()
				if yyn1214 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Int == nil {
						r.EncodeNil()
					} else {
						yy1217 := *x.FptrMapUint32Int
						yym1218 := z.EncBinary()
						_ = yym1218
						if false {
						} else {
							z.F.EncMapUint32IntV(yy1217, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Int8 == nil {
					r.EncodeNil()
				} else {
					yym1220 := z.EncBinary()
					_ = yym1220
					if false {
					} else {
						z.F.EncMapUint32Int8V(x.FMapUint32Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Int8"))
				r.WriteMapElemValue()
				if x.FMapUint32Int8 == nil {
					r.EncodeNil()
				} else {
					yym1221 := z.EncBinary()
					_ = yym1221
					if false {
					} else {
						z.F.EncMapUint32Int8V(x.FMapUint32Int8, e)
					}
				}
			}
			var yyn1222 bool
			if x.FptrMapUint32Int8 == nil {
				yyn1222 = true
				goto LABEL1222
			}
		LABEL1222:
			if yyr2 || yy2arr2 {
				if yyn1222 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Int8 == nil {
						r.EncodeNil()
					} else {
						yy1223 := *x.FptrMapUint32Int8
						yym1224 := z.EncBinary()
						_ = yym1224
						if false {
						} else {
							z.F.EncMapUint32Int8V(yy1223, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Int8"))
				r.WriteMapElemValue()
				if yyn1222 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Int8 == nil {
						r.EncodeNil()
					} else {
						yy1225 := *x.FptrMapUint32Int8
						yym1226 := z.EncBinary()
						_ = yym1226
						if false {
						} else {
							z.F.EncMapUint32Int8V(yy1225, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Int16 == nil {
					r.EncodeNil()
				} else {
					yym1228 := z.EncBinary()
					_ = yym1228
					if false {
					} else {
						z.F.EncMapUint32Int16V(x.FMapUint32Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Int16"))
				r.WriteMapElemValue()
				if x.FMapUint32Int16 == nil {
					r.EncodeNil()
				} else {
					yym1229 := z.EncBinary()
					_ = yym1229
					if false {
					} else {
						z.F.EncMapUint32Int16V(x.FMapUint32Int16, e)
					}
				}
			}
			var yyn1230 bool
			if x.FptrMapUint32Int16 == nil {
				yyn1230 = true
				goto LABEL1230
			}
		LABEL1230:
			if yyr2 || yy2arr2 {
				if yyn1230 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Int16 == nil {
						r.EncodeNil()
					} else {
						yy1231 := *x.FptrMapUint32Int16
						yym1232 := z.EncBinary()
						_ = yym1232
						if false {
						} else {
							z.F.EncMapUint32Int16V(yy1231, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Int16"))
				r.WriteMapElemValue()
				if yyn1230 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Int16 == nil {
						r.EncodeNil()
					} else {
						yy1233 := *x.FptrMapUint32Int16
						yym1234 := z.EncBinary()
						_ = yym1234
						if false {
						} else {
							z.F.EncMapUint32Int16V(yy1233, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Int32 == nil {
					r.EncodeNil()
				} else {
					yym1236 := z.EncBinary()
					_ = yym1236
					if false {
					} else {
						z.F.EncMapUint32Int32V(x.FMapUint32Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Int32"))
				r.WriteMapElemValue()
				if x.FMapUint32Int32 == nil {
					r.EncodeNil()
				} else {
					yym1237 := z.EncBinary()
					_ = yym1237
					if false {
					} else {
						z.F.EncMapUint32Int32V(x.FMapUint32Int32, e)
					}
				}
			}
			var yyn1238 bool
			if x.FptrMapUint32Int32 == nil {
				yyn1238 = true
				goto LABEL1238
			}
		LABEL1238:
			if yyr2 || yy2arr2 {
				if yyn1238 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Int32 == nil {
						r.EncodeNil()
					} else {
						yy1239 := *x.FptrMapUint32Int32
						yym1240 := z.EncBinary()
						_ = yym1240
						if false {
						} else {
							z.F.EncMapUint32Int32V(yy1239, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Int32"))
				r.WriteMapElemValue()
				if yyn1238 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Int32 == nil {
						r.EncodeNil()
					} else {
						yy1241 := *x.FptrMapUint32Int32
						yym1242 := z.EncBinary()
						_ = yym1242
						if false {
						} else {
							z.F.EncMapUint32Int32V(yy1241, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Int64 == nil {
					r.EncodeNil()
				} else {
					yym1244 := z.EncBinary()
					_ = yym1244
					if false {
					} else {
						z.F.EncMapUint32Int64V(x.FMapUint32Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Int64"))
				r.WriteMapElemValue()
				if x.FMapUint32Int64 == nil {
					r.EncodeNil()
				} else {
					yym1245 := z.EncBinary()
					_ = yym1245
					if false {
					} else {
						z.F.EncMapUint32Int64V(x.FMapUint32Int64, e)
					}
				}
			}
			var yyn1246 bool
			if x.FptrMapUint32Int64 == nil {
				yyn1246 = true
				goto LABEL1246
			}
		LABEL1246:
			if yyr2 || yy2arr2 {
				if yyn1246 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Int64 == nil {
						r.EncodeNil()
					} else {
						yy1247 := *x.FptrMapUint32Int64
						yym1248 := z.EncBinary()
						_ = yym1248
						if false {
						} else {
							z.F.EncMapUint32Int64V(yy1247, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Int64"))
				r.WriteMapElemValue()
				if yyn1246 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Int64 == nil {
						r.EncodeNil()
					} else {
						yy1249 := *x.FptrMapUint32Int64
						yym1250 := z.EncBinary()
						_ = yym1250
						if false {
						} else {
							z.F.EncMapUint32Int64V(yy1249, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Float32 == nil {
					r.EncodeNil()
				} else {
					yym1252 := z.EncBinary()
					_ = yym1252
					if false {
					} else {
						z.F.EncMapUint32Float32V(x.FMapUint32Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Float32"))
				r.WriteMapElemValue()
				if x.FMapUint32Float32 == nil {
					r.EncodeNil()
				} else {
					yym1253 := z.EncBinary()
					_ = yym1253
					if false {
					} else {
						z.F.EncMapUint32Float32V(x.FMapUint32Float32, e)
					}
				}
			}
			var yyn1254 bool
			if x.FptrMapUint32Float32 == nil {
				yyn1254 = true
				goto LABEL1254
			}
		LABEL1254:
			if yyr2 || yy2arr2 {
				if yyn1254 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Float32 == nil {
						r.EncodeNil()
					} else {
						yy1255 := *x.FptrMapUint32Float32
						yym1256 := z.EncBinary()
						_ = yym1256
						if false {
						} else {
							z.F.EncMapUint32Float32V(yy1255, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Float32"))
				r.WriteMapElemValue()
				if yyn1254 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Float32 == nil {
						r.EncodeNil()
					} else {
						yy1257 := *x.FptrMapUint32Float32
						yym1258 := z.EncBinary()
						_ = yym1258
						if false {
						} else {
							z.F.EncMapUint32Float32V(yy1257, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Float64 == nil {
					r.EncodeNil()
				} else {
					yym1260 := z.EncBinary()
					_ = yym1260
					if false {
					} else {
						z.F.EncMapUint32Float64V(x.FMapUint32Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Float64"))
				r.WriteMapElemValue()
				if x.FMapUint32Float64 == nil {
					r.EncodeNil()
				} else {
					yym1261 := z.EncBinary()
					_ = yym1261
					if false {
					} else {
						z.F.EncMapUint32Float64V(x.FMapUint32Float64, e)
					}
				}
			}
			var yyn1262 bool
			if x.FptrMapUint32Float64 == nil {
				yyn1262 = true
				goto LABEL1262
			}
		LABEL1262:
			if yyr2 || yy2arr2 {
				if yyn1262 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Float64 == nil {
						r.EncodeNil()
					} else {
						yy1263 := *x.FptrMapUint32Float64
						yym1264 := z.EncBinary()
						_ = yym1264
						if false {
						} else {
							z.F.EncMapUint32Float64V(yy1263, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Float64"))
				r.WriteMapElemValue()
				if yyn1262 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Float64 == nil {
						r.EncodeNil()
					} else {
						yy1265 := *x.FptrMapUint32Float64
						yym1266 := z.EncBinary()
						_ = yym1266
						if false {
						} else {
							z.F.EncMapUint32Float64V(yy1265, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint32Bool == nil {
					r.EncodeNil()
				} else {
					yym1268 := z.EncBinary()
					_ = yym1268
					if false {
					} else {
						z.F.EncMapUint32BoolV(x.FMapUint32Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint32Bool"))
				r.WriteMapElemValue()
				if x.FMapUint32Bool == nil {
					r.EncodeNil()
				} else {
					yym1269 := z.EncBinary()
					_ = yym1269
					if false {
					} else {
						z.F.EncMapUint32BoolV(x.FMapUint32Bool, e)
					}
				}
			}
			var yyn1270 bool
			if x.FptrMapUint32Bool == nil {
				yyn1270 = true
				goto LABEL1270
			}
		LABEL1270:
			if yyr2 || yy2arr2 {
				if yyn1270 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint32Bool == nil {
						r.EncodeNil()
					} else {
						yy1271 := *x.FptrMapUint32Bool
						yym1272 := z.EncBinary()
						_ = yym1272
						if false {
						} else {
							z.F.EncMapUint32BoolV(yy1271, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint32Bool"))
				r.WriteMapElemValue()
				if yyn1270 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint32Bool == nil {
						r.EncodeNil()
					} else {
						yy1273 := *x.FptrMapUint32Bool
						yym1274 := z.EncBinary()
						_ = yym1274
						if false {
						} else {
							z.F.EncMapUint32BoolV(yy1273, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Intf == nil {
					r.EncodeNil()
				} else {
					yym1276 := z.EncBinary()
					_ = yym1276
					if false {
					} else {
						z.F.EncMapUint64IntfV(x.FMapUint64Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Intf"))
				r.WriteMapElemValue()
				if x.FMapUint64Intf == nil {
					r.EncodeNil()
				} else {
					yym1277 := z.EncBinary()
					_ = yym1277
					if false {
					} else {
						z.F.EncMapUint64IntfV(x.FMapUint64Intf, e)
					}
				}
			}
			var yyn1278 bool
			if x.FptrMapUint64Intf == nil {
				yyn1278 = true
				goto LABEL1278
			}
		LABEL1278:
			if yyr2 || yy2arr2 {
				if yyn1278 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Intf == nil {
						r.EncodeNil()
					} else {
						yy1279 := *x.FptrMapUint64Intf
						yym1280 := z.EncBinary()
						_ = yym1280
						if false {
						} else {
							z.F.EncMapUint64IntfV(yy1279, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Intf"))
				r.WriteMapElemValue()
				if yyn1278 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Intf == nil {
						r.EncodeNil()
					} else {
						yy1281 := *x.FptrMapUint64Intf
						yym1282 := z.EncBinary()
						_ = yym1282
						if false {
						} else {
							z.F.EncMapUint64IntfV(yy1281, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64String == nil {
					r.EncodeNil()
				} else {
					yym1284 := z.EncBinary()
					_ = yym1284
					if false {
					} else {
						z.F.EncMapUint64StringV(x.FMapUint64String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64String"))
				r.WriteMapElemValue()
				if x.FMapUint64String == nil {
					r.EncodeNil()
				} else {
					yym1285 := z.EncBinary()
					_ = yym1285
					if false {
					} else {
						z.F.EncMapUint64StringV(x.FMapUint64String, e)
					}
				}
			}
			var yyn1286 bool
			if x.FptrMapUint64String == nil {
				yyn1286 = true
				goto LABEL1286
			}
		LABEL1286:
			if yyr2 || yy2arr2 {
				if yyn1286 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64String == nil {
						r.EncodeNil()
					} else {
						yy1287 := *x.FptrMapUint64String
						yym1288 := z.EncBinary()
						_ = yym1288
						if false {
						} else {
							z.F.EncMapUint64StringV(yy1287, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64String"))
				r.WriteMapElemValue()
				if yyn1286 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64String == nil {
						r.EncodeNil()
					} else {
						yy1289 := *x.FptrMapUint64String
						yym1290 := z.EncBinary()
						_ = yym1290
						if false {
						} else {
							z.F.EncMapUint64StringV(yy1289, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Uint == nil {
					r.EncodeNil()
				} else {
					yym1292 := z.EncBinary()
					_ = yym1292
					if false {
					} else {
						z.F.EncMapUint64UintV(x.FMapUint64Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Uint"))
				r.WriteMapElemValue()
				if x.FMapUint64Uint == nil {
					r.EncodeNil()
				} else {
					yym1293 := z.EncBinary()
					_ = yym1293
					if false {
					} else {
						z.F.EncMapUint64UintV(x.FMapUint64Uint, e)
					}
				}
			}
			var yyn1294 bool
			if x.FptrMapUint64Uint == nil {
				yyn1294 = true
				goto LABEL1294
			}
		LABEL1294:
			if yyr2 || yy2arr2 {
				if yyn1294 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Uint == nil {
						r.EncodeNil()
					} else {
						yy1295 := *x.FptrMapUint64Uint
						yym1296 := z.EncBinary()
						_ = yym1296
						if false {
						} else {
							z.F.EncMapUint64UintV(yy1295, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Uint"))
				r.WriteMapElemValue()
				if yyn1294 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Uint == nil {
						r.EncodeNil()
					} else {
						yy1297 := *x.FptrMapUint64Uint
						yym1298 := z.EncBinary()
						_ = yym1298
						if false {
						} else {
							z.F.EncMapUint64UintV(yy1297, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1300 := z.EncBinary()
					_ = yym1300
					if false {
					} else {
						z.F.EncMapUint64Uint8V(x.FMapUint64Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Uint8"))
				r.WriteMapElemValue()
				if x.FMapUint64Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1301 := z.EncBinary()
					_ = yym1301
					if false {
					} else {
						z.F.EncMapUint64Uint8V(x.FMapUint64Uint8, e)
					}
				}
			}
			var yyn1302 bool
			if x.FptrMapUint64Uint8 == nil {
				yyn1302 = true
				goto LABEL1302
			}
		LABEL1302:
			if yyr2 || yy2arr2 {
				if yyn1302 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1303 := *x.FptrMapUint64Uint8
						yym1304 := z.EncBinary()
						_ = yym1304
						if false {
						} else {
							z.F.EncMapUint64Uint8V(yy1303, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Uint8"))
				r.WriteMapElemValue()
				if yyn1302 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1305 := *x.FptrMapUint64Uint8
						yym1306 := z.EncBinary()
						_ = yym1306
						if false {
						} else {
							z.F.EncMapUint64Uint8V(yy1305, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1308 := z.EncBinary()
					_ = yym1308
					if false {
					} else {
						z.F.EncMapUint64Uint16V(x.FMapUint64Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Uint16"))
				r.WriteMapElemValue()
				if x.FMapUint64Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1309 := z.EncBinary()
					_ = yym1309
					if false {
					} else {
						z.F.EncMapUint64Uint16V(x.FMapUint64Uint16, e)
					}
				}
			}
			var yyn1310 bool
			if x.FptrMapUint64Uint16 == nil {
				yyn1310 = true
				goto LABEL1310
			}
		LABEL1310:
			if yyr2 || yy2arr2 {
				if yyn1310 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1311 := *x.FptrMapUint64Uint16
						yym1312 := z.EncBinary()
						_ = yym1312
						if false {
						} else {
							z.F.EncMapUint64Uint16V(yy1311, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Uint16"))
				r.WriteMapElemValue()
				if yyn1310 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1313 := *x.FptrMapUint64Uint16
						yym1314 := z.EncBinary()
						_ = yym1314
						if false {
						} else {
							z.F.EncMapUint64Uint16V(yy1313, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1316 := z.EncBinary()
					_ = yym1316
					if false {
					} else {
						z.F.EncMapUint64Uint32V(x.FMapUint64Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Uint32"))
				r.WriteMapElemValue()
				if x.FMapUint64Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1317 := z.EncBinary()
					_ = yym1317
					if false {
					} else {
						z.F.EncMapUint64Uint32V(x.FMapUint64Uint32, e)
					}
				}
			}
			var yyn1318 bool
			if x.FptrMapUint64Uint32 == nil {
				yyn1318 = true
				goto LABEL1318
			}
		LABEL1318:
			if yyr2 || yy2arr2 {
				if yyn1318 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1319 := *x.FptrMapUint64Uint32
						yym1320 := z.EncBinary()
						_ = yym1320
						if false {
						} else {
							z.F.EncMapUint64Uint32V(yy1319, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Uint32"))
				r.WriteMapElemValue()
				if yyn1318 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1321 := *x.FptrMapUint64Uint32
						yym1322 := z.EncBinary()
						_ = yym1322
						if false {
						} else {
							z.F.EncMapUint64Uint32V(yy1321, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1324 := z.EncBinary()
					_ = yym1324
					if false {
					} else {
						z.F.EncMapUint64Uint64V(x.FMapUint64Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Uint64"))
				r.WriteMapElemValue()
				if x.FMapUint64Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1325 := z.EncBinary()
					_ = yym1325
					if false {
					} else {
						z.F.EncMapUint64Uint64V(x.FMapUint64Uint64, e)
					}
				}
			}
			var yyn1326 bool
			if x.FptrMapUint64Uint64 == nil {
				yyn1326 = true
				goto LABEL1326
			}
		LABEL1326:
			if yyr2 || yy2arr2 {
				if yyn1326 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1327 := *x.FptrMapUint64Uint64
						yym1328 := z.EncBinary()
						_ = yym1328
						if false {
						} else {
							z.F.EncMapUint64Uint64V(yy1327, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Uint64"))
				r.WriteMapElemValue()
				if yyn1326 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1329 := *x.FptrMapUint64Uint64
						yym1330 := z.EncBinary()
						_ = yym1330
						if false {
						} else {
							z.F.EncMapUint64Uint64V(yy1329, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1332 := z.EncBinary()
					_ = yym1332
					if false {
					} else {
						z.F.EncMapUint64UintptrV(x.FMapUint64Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Uintptr"))
				r.WriteMapElemValue()
				if x.FMapUint64Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1333 := z.EncBinary()
					_ = yym1333
					if false {
					} else {
						z.F.EncMapUint64UintptrV(x.FMapUint64Uintptr, e)
					}
				}
			}
			var yyn1334 bool
			if x.FptrMapUint64Uintptr == nil {
				yyn1334 = true
				goto LABEL1334
			}
		LABEL1334:
			if yyr2 || yy2arr2 {
				if yyn1334 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1335 := *x.FptrMapUint64Uintptr
						yym1336 := z.EncBinary()
						_ = yym1336
						if false {
						} else {
							z.F.EncMapUint64UintptrV(yy1335, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Uintptr"))
				r.WriteMapElemValue()
				if yyn1334 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1337 := *x.FptrMapUint64Uintptr
						yym1338 := z.EncBinary()
						_ = yym1338
						if false {
						} else {
							z.F.EncMapUint64UintptrV(yy1337, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Int == nil {
					r.EncodeNil()
				} else {
					yym1340 := z.EncBinary()
					_ = yym1340
					if false {
					} else {
						z.F.EncMapUint64IntV(x.FMapUint64Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Int"))
				r.WriteMapElemValue()
				if x.FMapUint64Int == nil {
					r.EncodeNil()
				} else {
					yym1341 := z.EncBinary()
					_ = yym1341
					if false {
					} else {
						z.F.EncMapUint64IntV(x.FMapUint64Int, e)
					}
				}
			}
			var yyn1342 bool
			if x.FptrMapUint64Int == nil {
				yyn1342 = true
				goto LABEL1342
			}
		LABEL1342:
			if yyr2 || yy2arr2 {
				if yyn1342 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Int == nil {
						r.EncodeNil()
					} else {
						yy1343 := *x.FptrMapUint64Int
						yym1344 := z.EncBinary()
						_ = yym1344
						if false {
						} else {
							z.F.EncMapUint64IntV(yy1343, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Int"))
				r.WriteMapElemValue()
				if yyn1342 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Int == nil {
						r.EncodeNil()
					} else {
						yy1345 := *x.FptrMapUint64Int
						yym1346 := z.EncBinary()
						_ = yym1346
						if false {
						} else {
							z.F.EncMapUint64IntV(yy1345, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Int8 == nil {
					r.EncodeNil()
				} else {
					yym1348 := z.EncBinary()
					_ = yym1348
					if false {
					} else {
						z.F.EncMapUint64Int8V(x.FMapUint64Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Int8"))
				r.WriteMapElemValue()
				if x.FMapUint64Int8 == nil {
					r.EncodeNil()
				} else {
					yym1349 := z.EncBinary()
					_ = yym1349
					if false {
					} else {
						z.F.EncMapUint64Int8V(x.FMapUint64Int8, e)
					}
				}
			}
			var yyn1350 bool
			if x.FptrMapUint64Int8 == nil {
				yyn1350 = true
				goto LABEL1350
			}
		LABEL1350:
			if yyr2 || yy2arr2 {
				if yyn1350 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Int8 == nil {
						r.EncodeNil()
					} else {
						yy1351 := *x.FptrMapUint64Int8
						yym1352 := z.EncBinary()
						_ = yym1352
						if false {
						} else {
							z.F.EncMapUint64Int8V(yy1351, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Int8"))
				r.WriteMapElemValue()
				if yyn1350 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Int8 == nil {
						r.EncodeNil()
					} else {
						yy1353 := *x.FptrMapUint64Int8
						yym1354 := z.EncBinary()
						_ = yym1354
						if false {
						} else {
							z.F.EncMapUint64Int8V(yy1353, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Int16 == nil {
					r.EncodeNil()
				} else {
					yym1356 := z.EncBinary()
					_ = yym1356
					if false {
					} else {
						z.F.EncMapUint64Int16V(x.FMapUint64Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Int16"))
				r.WriteMapElemValue()
				if x.FMapUint64Int16 == nil {
					r.EncodeNil()
				} else {
					yym1357 := z.EncBinary()
					_ = yym1357
					if false {
					} else {
						z.F.EncMapUint64Int16V(x.FMapUint64Int16, e)
					}
				}
			}
			var yyn1358 bool
			if x.FptrMapUint64Int16 == nil {
				yyn1358 = true
				goto LABEL1358
			}
		LABEL1358:
			if yyr2 || yy2arr2 {
				if yyn1358 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Int16 == nil {
						r.EncodeNil()
					} else {
						yy1359 := *x.FptrMapUint64Int16
						yym1360 := z.EncBinary()
						_ = yym1360
						if false {
						} else {
							z.F.EncMapUint64Int16V(yy1359, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Int16"))
				r.WriteMapElemValue()
				if yyn1358 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Int16 == nil {
						r.EncodeNil()
					} else {
						yy1361 := *x.FptrMapUint64Int16
						yym1362 := z.EncBinary()
						_ = yym1362
						if false {
						} else {
							z.F.EncMapUint64Int16V(yy1361, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Int32 == nil {
					r.EncodeNil()
				} else {
					yym1364 := z.EncBinary()
					_ = yym1364
					if false {
					} else {
						z.F.EncMapUint64Int32V(x.FMapUint64Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Int32"))
				r.WriteMapElemValue()
				if x.FMapUint64Int32 == nil {
					r.EncodeNil()
				} else {
					yym1365 := z.EncBinary()
					_ = yym1365
					if false {
					} else {
						z.F.EncMapUint64Int32V(x.FMapUint64Int32, e)
					}
				}
			}
			var yyn1366 bool
			if x.FptrMapUint64Int32 == nil {
				yyn1366 = true
				goto LABEL1366
			}
		LABEL1366:
			if yyr2 || yy2arr2 {
				if yyn1366 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Int32 == nil {
						r.EncodeNil()
					} else {
						yy1367 := *x.FptrMapUint64Int32
						yym1368 := z.EncBinary()
						_ = yym1368
						if false {
						} else {
							z.F.EncMapUint64Int32V(yy1367, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Int32"))
				r.WriteMapElemValue()
				if yyn1366 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Int32 == nil {
						r.EncodeNil()
					} else {
						yy1369 := *x.FptrMapUint64Int32
						yym1370 := z.EncBinary()
						_ = yym1370
						if false {
						} else {
							z.F.EncMapUint64Int32V(yy1369, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Int64 == nil {
					r.EncodeNil()
				} else {
					yym1372 := z.EncBinary()
					_ = yym1372
					if false {
					} else {
						z.F.EncMapUint64Int64V(x.FMapUint64Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Int64"))
				r.WriteMapElemValue()
				if x.FMapUint64Int64 == nil {
					r.EncodeNil()
				} else {
					yym1373 := z.EncBinary()
					_ = yym1373
					if false {
					} else {
						z.F.EncMapUint64Int64V(x.FMapUint64Int64, e)
					}
				}
			}
			var yyn1374 bool
			if x.FptrMapUint64Int64 == nil {
				yyn1374 = true
				goto LABEL1374
			}
		LABEL1374:
			if yyr2 || yy2arr2 {
				if yyn1374 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Int64 == nil {
						r.EncodeNil()
					} else {
						yy1375 := *x.FptrMapUint64Int64
						yym1376 := z.EncBinary()
						_ = yym1376
						if false {
						} else {
							z.F.EncMapUint64Int64V(yy1375, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Int64"))
				r.WriteMapElemValue()
				if yyn1374 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Int64 == nil {
						r.EncodeNil()
					} else {
						yy1377 := *x.FptrMapUint64Int64
						yym1378 := z.EncBinary()
						_ = yym1378
						if false {
						} else {
							z.F.EncMapUint64Int64V(yy1377, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Float32 == nil {
					r.EncodeNil()
				} else {
					yym1380 := z.EncBinary()
					_ = yym1380
					if false {
					} else {
						z.F.EncMapUint64Float32V(x.FMapUint64Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Float32"))
				r.WriteMapElemValue()
				if x.FMapUint64Float32 == nil {
					r.EncodeNil()
				} else {
					yym1381 := z.EncBinary()
					_ = yym1381
					if false {
					} else {
						z.F.EncMapUint64Float32V(x.FMapUint64Float32, e)
					}
				}
			}
			var yyn1382 bool
			if x.FptrMapUint64Float32 == nil {
				yyn1382 = true
				goto LABEL1382
			}
		LABEL1382:
			if yyr2 || yy2arr2 {
				if yyn1382 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Float32 == nil {
						r.EncodeNil()
					} else {
						yy1383 := *x.FptrMapUint64Float32
						yym1384 := z.EncBinary()
						_ = yym1384
						if false {
						} else {
							z.F.EncMapUint64Float32V(yy1383, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Float32"))
				r.WriteMapElemValue()
				if yyn1382 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Float32 == nil {
						r.EncodeNil()
					} else {
						yy1385 := *x.FptrMapUint64Float32
						yym1386 := z.EncBinary()
						_ = yym1386
						if false {
						} else {
							z.F.EncMapUint64Float32V(yy1385, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Float64 == nil {
					r.EncodeNil()
				} else {
					yym1388 := z.EncBinary()
					_ = yym1388
					if false {
					} else {
						z.F.EncMapUint64Float64V(x.FMapUint64Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Float64"))
				r.WriteMapElemValue()
				if x.FMapUint64Float64 == nil {
					r.EncodeNil()
				} else {
					yym1389 := z.EncBinary()
					_ = yym1389
					if false {
					} else {
						z.F.EncMapUint64Float64V(x.FMapUint64Float64, e)
					}
				}
			}
			var yyn1390 bool
			if x.FptrMapUint64Float64 == nil {
				yyn1390 = true
				goto LABEL1390
			}
		LABEL1390:
			if yyr2 || yy2arr2 {
				if yyn1390 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Float64 == nil {
						r.EncodeNil()
					} else {
						yy1391 := *x.FptrMapUint64Float64
						yym1392 := z.EncBinary()
						_ = yym1392
						if false {
						} else {
							z.F.EncMapUint64Float64V(yy1391, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Float64"))
				r.WriteMapElemValue()
				if yyn1390 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Float64 == nil {
						r.EncodeNil()
					} else {
						yy1393 := *x.FptrMapUint64Float64
						yym1394 := z.EncBinary()
						_ = yym1394
						if false {
						} else {
							z.F.EncMapUint64Float64V(yy1393, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUint64Bool == nil {
					r.EncodeNil()
				} else {
					yym1396 := z.EncBinary()
					_ = yym1396
					if false {
					} else {
						z.F.EncMapUint64BoolV(x.FMapUint64Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUint64Bool"))
				r.WriteMapElemValue()
				if x.FMapUint64Bool == nil {
					r.EncodeNil()
				} else {
					yym1397 := z.EncBinary()
					_ = yym1397
					if false {
					} else {
						z.F.EncMapUint64BoolV(x.FMapUint64Bool, e)
					}
				}
			}
			var yyn1398 bool
			if x.FptrMapUint64Bool == nil {
				yyn1398 = true
				goto LABEL1398
			}
		LABEL1398:
			if yyr2 || yy2arr2 {
				if yyn1398 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUint64Bool == nil {
						r.EncodeNil()
					} else {
						yy1399 := *x.FptrMapUint64Bool
						yym1400 := z.EncBinary()
						_ = yym1400
						if false {
						} else {
							z.F.EncMapUint64BoolV(yy1399, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUint64Bool"))
				r.WriteMapElemValue()
				if yyn1398 {
					r.EncodeNil()
				} else {
					if x.FptrMapUint64Bool == nil {
						r.EncodeNil()
					} else {
						yy1401 := *x.FptrMapUint64Bool
						yym1402 := z.EncBinary()
						_ = yym1402
						if false {
						} else {
							z.F.EncMapUint64BoolV(yy1401, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrIntf == nil {
					r.EncodeNil()
				} else {
					yym1404 := z.EncBinary()
					_ = yym1404
					if false {
					} else {
						z.F.EncMapUintptrIntfV(x.FMapUintptrIntf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrIntf"))
				r.WriteMapElemValue()
				if x.FMapUintptrIntf == nil {
					r.EncodeNil()
				} else {
					yym1405 := z.EncBinary()
					_ = yym1405
					if false {
					} else {
						z.F.EncMapUintptrIntfV(x.FMapUintptrIntf, e)
					}
				}
			}
			var yyn1406 bool
			if x.FptrMapUintptrIntf == nil {
				yyn1406 = true
				goto LABEL1406
			}
		LABEL1406:
			if yyr2 || yy2arr2 {
				if yyn1406 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrIntf == nil {
						r.EncodeNil()
					} else {
						yy1407 := *x.FptrMapUintptrIntf
						yym1408 := z.EncBinary()
						_ = yym1408
						if false {
						} else {
							z.F.EncMapUintptrIntfV(yy1407, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrIntf"))
				r.WriteMapElemValue()
				if yyn1406 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrIntf == nil {
						r.EncodeNil()
					} else {
						yy1409 := *x.FptrMapUintptrIntf
						yym1410 := z.EncBinary()
						_ = yym1410
						if false {
						} else {
							z.F.EncMapUintptrIntfV(yy1409, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrString == nil {
					r.EncodeNil()
				} else {
					yym1412 := z.EncBinary()
					_ = yym1412
					if false {
					} else {
						z.F.EncMapUintptrStringV(x.FMapUintptrString, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrString"))
				r.WriteMapElemValue()
				if x.FMapUintptrString == nil {
					r.EncodeNil()
				} else {
					yym1413 := z.EncBinary()
					_ = yym1413
					if false {
					} else {
						z.F.EncMapUintptrStringV(x.FMapUintptrString, e)
					}
				}
			}
			var yyn1414 bool
			if x.FptrMapUintptrString == nil {
				yyn1414 = true
				goto LABEL1414
			}
		LABEL1414:
			if yyr2 || yy2arr2 {
				if yyn1414 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrString == nil {
						r.EncodeNil()
					} else {
						yy1415 := *x.FptrMapUintptrString
						yym1416 := z.EncBinary()
						_ = yym1416
						if false {
						} else {
							z.F.EncMapUintptrStringV(yy1415, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrString"))
				r.WriteMapElemValue()
				if yyn1414 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrString == nil {
						r.EncodeNil()
					} else {
						yy1417 := *x.FptrMapUintptrString
						yym1418 := z.EncBinary()
						_ = yym1418
						if false {
						} else {
							z.F.EncMapUintptrStringV(yy1417, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrUint == nil {
					r.EncodeNil()
				} else {
					yym1420 := z.EncBinary()
					_ = yym1420
					if false {
					} else {
						z.F.EncMapUintptrUintV(x.FMapUintptrUint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrUint"))
				r.WriteMapElemValue()
				if x.FMapUintptrUint == nil {
					r.EncodeNil()
				} else {
					yym1421 := z.EncBinary()
					_ = yym1421
					if false {
					} else {
						z.F.EncMapUintptrUintV(x.FMapUintptrUint, e)
					}
				}
			}
			var yyn1422 bool
			if x.FptrMapUintptrUint == nil {
				yyn1422 = true
				goto LABEL1422
			}
		LABEL1422:
			if yyr2 || yy2arr2 {
				if yyn1422 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrUint == nil {
						r.EncodeNil()
					} else {
						yy1423 := *x.FptrMapUintptrUint
						yym1424 := z.EncBinary()
						_ = yym1424
						if false {
						} else {
							z.F.EncMapUintptrUintV(yy1423, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrUint"))
				r.WriteMapElemValue()
				if yyn1422 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrUint == nil {
						r.EncodeNil()
					} else {
						yy1425 := *x.FptrMapUintptrUint
						yym1426 := z.EncBinary()
						_ = yym1426
						if false {
						} else {
							z.F.EncMapUintptrUintV(yy1425, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrUint8 == nil {
					r.EncodeNil()
				} else {
					yym1428 := z.EncBinary()
					_ = yym1428
					if false {
					} else {
						z.F.EncMapUintptrUint8V(x.FMapUintptrUint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrUint8"))
				r.WriteMapElemValue()
				if x.FMapUintptrUint8 == nil {
					r.EncodeNil()
				} else {
					yym1429 := z.EncBinary()
					_ = yym1429
					if false {
					} else {
						z.F.EncMapUintptrUint8V(x.FMapUintptrUint8, e)
					}
				}
			}
			var yyn1430 bool
			if x.FptrMapUintptrUint8 == nil {
				yyn1430 = true
				goto LABEL1430
			}
		LABEL1430:
			if yyr2 || yy2arr2 {
				if yyn1430 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrUint8 == nil {
						r.EncodeNil()
					} else {
						yy1431 := *x.FptrMapUintptrUint8
						yym1432 := z.EncBinary()
						_ = yym1432
						if false {
						} else {
							z.F.EncMapUintptrUint8V(yy1431, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrUint8"))
				r.WriteMapElemValue()
				if yyn1430 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrUint8 == nil {
						r.EncodeNil()
					} else {
						yy1433 := *x.FptrMapUintptrUint8
						yym1434 := z.EncBinary()
						_ = yym1434
						if false {
						} else {
							z.F.EncMapUintptrUint8V(yy1433, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrUint16 == nil {
					r.EncodeNil()
				} else {
					yym1436 := z.EncBinary()
					_ = yym1436
					if false {
					} else {
						z.F.EncMapUintptrUint16V(x.FMapUintptrUint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrUint16"))
				r.WriteMapElemValue()
				if x.FMapUintptrUint16 == nil {
					r.EncodeNil()
				} else {
					yym1437 := z.EncBinary()
					_ = yym1437
					if false {
					} else {
						z.F.EncMapUintptrUint16V(x.FMapUintptrUint16, e)
					}
				}
			}
			var yyn1438 bool
			if x.FptrMapUintptrUint16 == nil {
				yyn1438 = true
				goto LABEL1438
			}
		LABEL1438:
			if yyr2 || yy2arr2 {
				if yyn1438 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrUint16 == nil {
						r.EncodeNil()
					} else {
						yy1439 := *x.FptrMapUintptrUint16
						yym1440 := z.EncBinary()
						_ = yym1440
						if false {
						} else {
							z.F.EncMapUintptrUint16V(yy1439, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrUint16"))
				r.WriteMapElemValue()
				if yyn1438 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrUint16 == nil {
						r.EncodeNil()
					} else {
						yy1441 := *x.FptrMapUintptrUint16
						yym1442 := z.EncBinary()
						_ = yym1442
						if false {
						} else {
							z.F.EncMapUintptrUint16V(yy1441, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrUint32 == nil {
					r.EncodeNil()
				} else {
					yym1444 := z.EncBinary()
					_ = yym1444
					if false {
					} else {
						z.F.EncMapUintptrUint32V(x.FMapUintptrUint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrUint32"))
				r.WriteMapElemValue()
				if x.FMapUintptrUint32 == nil {
					r.EncodeNil()
				} else {
					yym1445 := z.EncBinary()
					_ = yym1445
					if false {
					} else {
						z.F.EncMapUintptrUint32V(x.FMapUintptrUint32, e)
					}
				}
			}
			var yyn1446 bool
			if x.FptrMapUintptrUint32 == nil {
				yyn1446 = true
				goto LABEL1446
			}
		LABEL1446:
			if yyr2 || yy2arr2 {
				if yyn1446 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrUint32 == nil {
						r.EncodeNil()
					} else {
						yy1447 := *x.FptrMapUintptrUint32
						yym1448 := z.EncBinary()
						_ = yym1448
						if false {
						} else {
							z.F.EncMapUintptrUint32V(yy1447, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrUint32"))
				r.WriteMapElemValue()
				if yyn1446 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrUint32 == nil {
						r.EncodeNil()
					} else {
						yy1449 := *x.FptrMapUintptrUint32
						yym1450 := z.EncBinary()
						_ = yym1450
						if false {
						} else {
							z.F.EncMapUintptrUint32V(yy1449, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrUint64 == nil {
					r.EncodeNil()
				} else {
					yym1452 := z.EncBinary()
					_ = yym1452
					if false {
					} else {
						z.F.EncMapUintptrUint64V(x.FMapUintptrUint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrUint64"))
				r.WriteMapElemValue()
				if x.FMapUintptrUint64 == nil {
					r.EncodeNil()
				} else {
					yym1453 := z.EncBinary()
					_ = yym1453
					if false {
					} else {
						z.F.EncMapUintptrUint64V(x.FMapUintptrUint64, e)
					}
				}
			}
			var yyn1454 bool
			if x.FptrMapUintptrUint64 == nil {
				yyn1454 = true
				goto LABEL1454
			}
		LABEL1454:
			if yyr2 || yy2arr2 {
				if yyn1454 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrUint64 == nil {
						r.EncodeNil()
					} else {
						yy1455 := *x.FptrMapUintptrUint64
						yym1456 := z.EncBinary()
						_ = yym1456
						if false {
						} else {
							z.F.EncMapUintptrUint64V(yy1455, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrUint64"))
				r.WriteMapElemValue()
				if yyn1454 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrUint64 == nil {
						r.EncodeNil()
					} else {
						yy1457 := *x.FptrMapUintptrUint64
						yym1458 := z.EncBinary()
						_ = yym1458
						if false {
						} else {
							z.F.EncMapUintptrUint64V(yy1457, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrUintptr == nil {
					r.EncodeNil()
				} else {
					yym1460 := z.EncBinary()
					_ = yym1460
					if false {
					} else {
						z.F.EncMapUintptrUintptrV(x.FMapUintptrUintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrUintptr"))
				r.WriteMapElemValue()
				if x.FMapUintptrUintptr == nil {
					r.EncodeNil()
				} else {
					yym1461 := z.EncBinary()
					_ = yym1461
					if false {
					} else {
						z.F.EncMapUintptrUintptrV(x.FMapUintptrUintptr, e)
					}
				}
			}
			var yyn1462 bool
			if x.FptrMapUintptrUintptr == nil {
				yyn1462 = true
				goto LABEL1462
			}
		LABEL1462:
			if yyr2 || yy2arr2 {
				if yyn1462 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrUintptr == nil {
						r.EncodeNil()
					} else {
						yy1463 := *x.FptrMapUintptrUintptr
						yym1464 := z.EncBinary()
						_ = yym1464
						if false {
						} else {
							z.F.EncMapUintptrUintptrV(yy1463, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrUintptr"))
				r.WriteMapElemValue()
				if yyn1462 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrUintptr == nil {
						r.EncodeNil()
					} else {
						yy1465 := *x.FptrMapUintptrUintptr
						yym1466 := z.EncBinary()
						_ = yym1466
						if false {
						} else {
							z.F.EncMapUintptrUintptrV(yy1465, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrInt == nil {
					r.EncodeNil()
				} else {
					yym1468 := z.EncBinary()
					_ = yym1468
					if false {
					} else {
						z.F.EncMapUintptrIntV(x.FMapUintptrInt, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrInt"))
				r.WriteMapElemValue()
				if x.FMapUintptrInt == nil {
					r.EncodeNil()
				} else {
					yym1469 := z.EncBinary()
					_ = yym1469
					if false {
					} else {
						z.F.EncMapUintptrIntV(x.FMapUintptrInt, e)
					}
				}
			}
			var yyn1470 bool
			if x.FptrMapUintptrInt == nil {
				yyn1470 = true
				goto LABEL1470
			}
		LABEL1470:
			if yyr2 || yy2arr2 {
				if yyn1470 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrInt == nil {
						r.EncodeNil()
					} else {
						yy1471 := *x.FptrMapUintptrInt
						yym1472 := z.EncBinary()
						_ = yym1472
						if false {
						} else {
							z.F.EncMapUintptrIntV(yy1471, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrInt"))
				r.WriteMapElemValue()
				if yyn1470 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrInt == nil {
						r.EncodeNil()
					} else {
						yy1473 := *x.FptrMapUintptrInt
						yym1474 := z.EncBinary()
						_ = yym1474
						if false {
						} else {
							z.F.EncMapUintptrIntV(yy1473, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrInt8 == nil {
					r.EncodeNil()
				} else {
					yym1476 := z.EncBinary()
					_ = yym1476
					if false {
					} else {
						z.F.EncMapUintptrInt8V(x.FMapUintptrInt8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrInt8"))
				r.WriteMapElemValue()
				if x.FMapUintptrInt8 == nil {
					r.EncodeNil()
				} else {
					yym1477 := z.EncBinary()
					_ = yym1477
					if false {
					} else {
						z.F.EncMapUintptrInt8V(x.FMapUintptrInt8, e)
					}
				}
			}
			var yyn1478 bool
			if x.FptrMapUintptrInt8 == nil {
				yyn1478 = true
				goto LABEL1478
			}
		LABEL1478:
			if yyr2 || yy2arr2 {
				if yyn1478 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrInt8 == nil {
						r.EncodeNil()
					} else {
						yy1479 := *x.FptrMapUintptrInt8
						yym1480 := z.EncBinary()
						_ = yym1480
						if false {
						} else {
							z.F.EncMapUintptrInt8V(yy1479, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrInt8"))
				r.WriteMapElemValue()
				if yyn1478 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrInt8 == nil {
						r.EncodeNil()
					} else {
						yy1481 := *x.FptrMapUintptrInt8
						yym1482 := z.EncBinary()
						_ = yym1482
						if false {
						} else {
							z.F.EncMapUintptrInt8V(yy1481, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrInt16 == nil {
					r.EncodeNil()
				} else {
					yym1484 := z.EncBinary()
					_ = yym1484
					if false {
					} else {
						z.F.EncMapUintptrInt16V(x.FMapUintptrInt16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrInt16"))
				r.WriteMapElemValue()
				if x.FMapUintptrInt16 == nil {
					r.EncodeNil()
				} else {
					yym1485 := z.EncBinary()
					_ = yym1485
					if false {
					} else {
						z.F.EncMapUintptrInt16V(x.FMapUintptrInt16, e)
					}
				}
			}
			var yyn1486 bool
			if x.FptrMapUintptrInt16 == nil {
				yyn1486 = true
				goto LABEL1486
			}
		LABEL1486:
			if yyr2 || yy2arr2 {
				if yyn1486 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrInt16 == nil {
						r.EncodeNil()
					} else {
						yy1487 := *x.FptrMapUintptrInt16
						yym1488 := z.EncBinary()
						_ = yym1488
						if false {
						} else {
							z.F.EncMapUintptrInt16V(yy1487, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrInt16"))
				r.WriteMapElemValue()
				if yyn1486 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrInt16 == nil {
						r.EncodeNil()
					} else {
						yy1489 := *x.FptrMapUintptrInt16
						yym1490 := z.EncBinary()
						_ = yym1490
						if false {
						} else {
							z.F.EncMapUintptrInt16V(yy1489, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrInt32 == nil {
					r.EncodeNil()
				} else {
					yym1492 := z.EncBinary()
					_ = yym1492
					if false {
					} else {
						z.F.EncMapUintptrInt32V(x.FMapUintptrInt32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrInt32"))
				r.WriteMapElemValue()
				if x.FMapUintptrInt32 == nil {
					r.EncodeNil()
				} else {
					yym1493 := z.EncBinary()
					_ = yym1493
					if false {
					} else {
						z.F.EncMapUintptrInt32V(x.FMapUintptrInt32, e)
					}
				}
			}
			var yyn1494 bool
			if x.FptrMapUintptrInt32 == nil {
				yyn1494 = true
				goto LABEL1494
			}
		LABEL1494:
			if yyr2 || yy2arr2 {
				if yyn1494 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrInt32 == nil {
						r.EncodeNil()
					} else {
						yy1495 := *x.FptrMapUintptrInt32
						yym1496 := z.EncBinary()
						_ = yym1496
						if false {
						} else {
							z.F.EncMapUintptrInt32V(yy1495, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrInt32"))
				r.WriteMapElemValue()
				if yyn1494 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrInt32 == nil {
						r.EncodeNil()
					} else {
						yy1497 := *x.FptrMapUintptrInt32
						yym1498 := z.EncBinary()
						_ = yym1498
						if false {
						} else {
							z.F.EncMapUintptrInt32V(yy1497, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrInt64 == nil {
					r.EncodeNil()
				} else {
					yym1500 := z.EncBinary()
					_ = yym1500
					if false {
					} else {
						z.F.EncMapUintptrInt64V(x.FMapUintptrInt64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrInt64"))
				r.WriteMapElemValue()
				if x.FMapUintptrInt64 == nil {
					r.EncodeNil()
				} else {
					yym1501 := z.EncBinary()
					_ = yym1501
					if false {
					} else {
						z.F.EncMapUintptrInt64V(x.FMapUintptrInt64, e)
					}
				}
			}
			var yyn1502 bool
			if x.FptrMapUintptrInt64 == nil {
				yyn1502 = true
				goto LABEL1502
			}
		LABEL1502:
			if yyr2 || yy2arr2 {
				if yyn1502 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrInt64 == nil {
						r.EncodeNil()
					} else {
						yy1503 := *x.FptrMapUintptrInt64
						yym1504 := z.EncBinary()
						_ = yym1504
						if false {
						} else {
							z.F.EncMapUintptrInt64V(yy1503, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrInt64"))
				r.WriteMapElemValue()
				if yyn1502 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrInt64 == nil {
						r.EncodeNil()
					} else {
						yy1505 := *x.FptrMapUintptrInt64
						yym1506 := z.EncBinary()
						_ = yym1506
						if false {
						} else {
							z.F.EncMapUintptrInt64V(yy1505, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrFloat32 == nil {
					r.EncodeNil()
				} else {
					yym1508 := z.EncBinary()
					_ = yym1508
					if false {
					} else {
						z.F.EncMapUintptrFloat32V(x.FMapUintptrFloat32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrFloat32"))
				r.WriteMapElemValue()
				if x.FMapUintptrFloat32 == nil {
					r.EncodeNil()
				} else {
					yym1509 := z.EncBinary()
					_ = yym1509
					if false {
					} else {
						z.F.EncMapUintptrFloat32V(x.FMapUintptrFloat32, e)
					}
				}
			}
			var yyn1510 bool
			if x.FptrMapUintptrFloat32 == nil {
				yyn1510 = true
				goto LABEL1510
			}
		LABEL1510:
			if yyr2 || yy2arr2 {
				if yyn1510 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrFloat32 == nil {
						r.EncodeNil()
					} else {
						yy1511 := *x.FptrMapUintptrFloat32
						yym1512 := z.EncBinary()
						_ = yym1512
						if false {
						} else {
							z.F.EncMapUintptrFloat32V(yy1511, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrFloat32"))
				r.WriteMapElemValue()
				if yyn1510 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrFloat32 == nil {
						r.EncodeNil()
					} else {
						yy1513 := *x.FptrMapUintptrFloat32
						yym1514 := z.EncBinary()
						_ = yym1514
						if false {
						} else {
							z.F.EncMapUintptrFloat32V(yy1513, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrFloat64 == nil {
					r.EncodeNil()
				} else {
					yym1516 := z.EncBinary()
					_ = yym1516
					if false {
					} else {
						z.F.EncMapUintptrFloat64V(x.FMapUintptrFloat64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrFloat64"))
				r.WriteMapElemValue()
				if x.FMapUintptrFloat64 == nil {
					r.EncodeNil()
				} else {
					yym1517 := z.EncBinary()
					_ = yym1517
					if false {
					} else {
						z.F.EncMapUintptrFloat64V(x.FMapUintptrFloat64, e)
					}
				}
			}
			var yyn1518 bool
			if x.FptrMapUintptrFloat64 == nil {
				yyn1518 = true
				goto LABEL1518
			}
		LABEL1518:
			if yyr2 || yy2arr2 {
				if yyn1518 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrFloat64 == nil {
						r.EncodeNil()
					} else {
						yy1519 := *x.FptrMapUintptrFloat64
						yym1520 := z.EncBinary()
						_ = yym1520
						if false {
						} else {
							z.F.EncMapUintptrFloat64V(yy1519, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrFloat64"))
				r.WriteMapElemValue()
				if yyn1518 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrFloat64 == nil {
						r.EncodeNil()
					} else {
						yy1521 := *x.FptrMapUintptrFloat64
						yym1522 := z.EncBinary()
						_ = yym1522
						if false {
						} else {
							z.F.EncMapUintptrFloat64V(yy1521, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapUintptrBool == nil {
					r.EncodeNil()
				} else {
					yym1524 := z.EncBinary()
					_ = yym1524
					if false {
					} else {
						z.F.EncMapUintptrBoolV(x.FMapUintptrBool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapUintptrBool"))
				r.WriteMapElemValue()
				if x.FMapUintptrBool == nil {
					r.EncodeNil()
				} else {
					yym1525 := z.EncBinary()
					_ = yym1525
					if false {
					} else {
						z.F.EncMapUintptrBoolV(x.FMapUintptrBool, e)
					}
				}
			}
			var yyn1526 bool
			if x.FptrMapUintptrBool == nil {
				yyn1526 = true
				goto LABEL1526
			}
		LABEL1526:
			if yyr2 || yy2arr2 {
				if yyn1526 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapUintptrBool == nil {
						r.EncodeNil()
					} else {
						yy1527 := *x.FptrMapUintptrBool
						yym1528 := z.EncBinary()
						_ = yym1528
						if false {
						} else {
							z.F.EncMapUintptrBoolV(yy1527, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapUintptrBool"))
				r.WriteMapElemValue()
				if yyn1526 {
					r.EncodeNil()
				} else {
					if x.FptrMapUintptrBool == nil {
						r.EncodeNil()
					} else {
						yy1529 := *x.FptrMapUintptrBool
						yym1530 := z.EncBinary()
						_ = yym1530
						if false {
						} else {
							z.F.EncMapUintptrBoolV(yy1529, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntIntf == nil {
					r.EncodeNil()
				} else {
					yym1532 := z.EncBinary()
					_ = yym1532
					if false {
					} else {
						z.F.EncMapIntIntfV(x.FMapIntIntf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntIntf"))
				r.WriteMapElemValue()
				if x.FMapIntIntf == nil {
					r.EncodeNil()
				} else {
					yym1533 := z.EncBinary()
					_ = yym1533
					if false {
					} else {
						z.F.EncMapIntIntfV(x.FMapIntIntf, e)
					}
				}
			}
			var yyn1534 bool
			if x.FptrMapIntIntf == nil {
				yyn1534 = true
				goto LABEL1534
			}
		LABEL1534:
			if yyr2 || yy2arr2 {
				if yyn1534 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntIntf == nil {
						r.EncodeNil()
					} else {
						yy1535 := *x.FptrMapIntIntf
						yym1536 := z.EncBinary()
						_ = yym1536
						if false {
						} else {
							z.F.EncMapIntIntfV(yy1535, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntIntf"))
				r.WriteMapElemValue()
				if yyn1534 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntIntf == nil {
						r.EncodeNil()
					} else {
						yy1537 := *x.FptrMapIntIntf
						yym1538 := z.EncBinary()
						_ = yym1538
						if false {
						} else {
							z.F.EncMapIntIntfV(yy1537, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntString == nil {
					r.EncodeNil()
				} else {
					yym1540 := z.EncBinary()
					_ = yym1540
					if false {
					} else {
						z.F.EncMapIntStringV(x.FMapIntString, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntString"))
				r.WriteMapElemValue()
				if x.FMapIntString == nil {
					r.EncodeNil()
				} else {
					yym1541 := z.EncBinary()
					_ = yym1541
					if false {
					} else {
						z.F.EncMapIntStringV(x.FMapIntString, e)
					}
				}
			}
			var yyn1542 bool
			if x.FptrMapIntString == nil {
				yyn1542 = true
				goto LABEL1542
			}
		LABEL1542:
			if yyr2 || yy2arr2 {
				if yyn1542 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntString == nil {
						r.EncodeNil()
					} else {
						yy1543 := *x.FptrMapIntString
						yym1544 := z.EncBinary()
						_ = yym1544
						if false {
						} else {
							z.F.EncMapIntStringV(yy1543, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntString"))
				r.WriteMapElemValue()
				if yyn1542 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntString == nil {
						r.EncodeNil()
					} else {
						yy1545 := *x.FptrMapIntString
						yym1546 := z.EncBinary()
						_ = yym1546
						if false {
						} else {
							z.F.EncMapIntStringV(yy1545, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntUint == nil {
					r.EncodeNil()
				} else {
					yym1548 := z.EncBinary()
					_ = yym1548
					if false {
					} else {
						z.F.EncMapIntUintV(x.FMapIntUint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntUint"))
				r.WriteMapElemValue()
				if x.FMapIntUint == nil {
					r.EncodeNil()
				} else {
					yym1549 := z.EncBinary()
					_ = yym1549
					if false {
					} else {
						z.F.EncMapIntUintV(x.FMapIntUint, e)
					}
				}
			}
			var yyn1550 bool
			if x.FptrMapIntUint == nil {
				yyn1550 = true
				goto LABEL1550
			}
		LABEL1550:
			if yyr2 || yy2arr2 {
				if yyn1550 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntUint == nil {
						r.EncodeNil()
					} else {
						yy1551 := *x.FptrMapIntUint
						yym1552 := z.EncBinary()
						_ = yym1552
						if false {
						} else {
							z.F.EncMapIntUintV(yy1551, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntUint"))
				r.WriteMapElemValue()
				if yyn1550 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntUint == nil {
						r.EncodeNil()
					} else {
						yy1553 := *x.FptrMapIntUint
						yym1554 := z.EncBinary()
						_ = yym1554
						if false {
						} else {
							z.F.EncMapIntUintV(yy1553, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntUint8 == nil {
					r.EncodeNil()
				} else {
					yym1556 := z.EncBinary()
					_ = yym1556
					if false {
					} else {
						z.F.EncMapIntUint8V(x.FMapIntUint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntUint8"))
				r.WriteMapElemValue()
				if x.FMapIntUint8 == nil {
					r.EncodeNil()
				} else {
					yym1557 := z.EncBinary()
					_ = yym1557
					if false {
					} else {
						z.F.EncMapIntUint8V(x.FMapIntUint8, e)
					}
				}
			}
			var yyn1558 bool
			if x.FptrMapIntUint8 == nil {
				yyn1558 = true
				goto LABEL1558
			}
		LABEL1558:
			if yyr2 || yy2arr2 {
				if yyn1558 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntUint8 == nil {
						r.EncodeNil()
					} else {
						yy1559 := *x.FptrMapIntUint8
						yym1560 := z.EncBinary()
						_ = yym1560
						if false {
						} else {
							z.F.EncMapIntUint8V(yy1559, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntUint8"))
				r.WriteMapElemValue()
				if yyn1558 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntUint8 == nil {
						r.EncodeNil()
					} else {
						yy1561 := *x.FptrMapIntUint8
						yym1562 := z.EncBinary()
						_ = yym1562
						if false {
						} else {
							z.F.EncMapIntUint8V(yy1561, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntUint16 == nil {
					r.EncodeNil()
				} else {
					yym1564 := z.EncBinary()
					_ = yym1564
					if false {
					} else {
						z.F.EncMapIntUint16V(x.FMapIntUint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntUint16"))
				r.WriteMapElemValue()
				if x.FMapIntUint16 == nil {
					r.EncodeNil()
				} else {
					yym1565 := z.EncBinary()
					_ = yym1565
					if false {
					} else {
						z.F.EncMapIntUint16V(x.FMapIntUint16, e)
					}
				}
			}
			var yyn1566 bool
			if x.FptrMapIntUint16 == nil {
				yyn1566 = true
				goto LABEL1566
			}
		LABEL1566:
			if yyr2 || yy2arr2 {
				if yyn1566 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntUint16 == nil {
						r.EncodeNil()
					} else {
						yy1567 := *x.FptrMapIntUint16
						yym1568 := z.EncBinary()
						_ = yym1568
						if false {
						} else {
							z.F.EncMapIntUint16V(yy1567, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntUint16"))
				r.WriteMapElemValue()
				if yyn1566 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntUint16 == nil {
						r.EncodeNil()
					} else {
						yy1569 := *x.FptrMapIntUint16
						yym1570 := z.EncBinary()
						_ = yym1570
						if false {
						} else {
							z.F.EncMapIntUint16V(yy1569, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntUint32 == nil {
					r.EncodeNil()
				} else {
					yym1572 := z.EncBinary()
					_ = yym1572
					if false {
					} else {
						z.F.EncMapIntUint32V(x.FMapIntUint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntUint32"))
				r.WriteMapElemValue()
				if x.FMapIntUint32 == nil {
					r.EncodeNil()
				} else {
					yym1573 := z.EncBinary()
					_ = yym1573
					if false {
					} else {
						z.F.EncMapIntUint32V(x.FMapIntUint32, e)
					}
				}
			}
			var yyn1574 bool
			if x.FptrMapIntUint32 == nil {
				yyn1574 = true
				goto LABEL1574
			}
		LABEL1574:
			if yyr2 || yy2arr2 {
				if yyn1574 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntUint32 == nil {
						r.EncodeNil()
					} else {
						yy1575 := *x.FptrMapIntUint32
						yym1576 := z.EncBinary()
						_ = yym1576
						if false {
						} else {
							z.F.EncMapIntUint32V(yy1575, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntUint32"))
				r.WriteMapElemValue()
				if yyn1574 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntUint32 == nil {
						r.EncodeNil()
					} else {
						yy1577 := *x.FptrMapIntUint32
						yym1578 := z.EncBinary()
						_ = yym1578
						if false {
						} else {
							z.F.EncMapIntUint32V(yy1577, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntUint64 == nil {
					r.EncodeNil()
				} else {
					yym1580 := z.EncBinary()
					_ = yym1580
					if false {
					} else {
						z.F.EncMapIntUint64V(x.FMapIntUint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntUint64"))
				r.WriteMapElemValue()
				if x.FMapIntUint64 == nil {
					r.EncodeNil()
				} else {
					yym1581 := z.EncBinary()
					_ = yym1581
					if false {
					} else {
						z.F.EncMapIntUint64V(x.FMapIntUint64, e)
					}
				}
			}
			var yyn1582 bool
			if x.FptrMapIntUint64 == nil {
				yyn1582 = true
				goto LABEL1582
			}
		LABEL1582:
			if yyr2 || yy2arr2 {
				if yyn1582 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntUint64 == nil {
						r.EncodeNil()
					} else {
						yy1583 := *x.FptrMapIntUint64
						yym1584 := z.EncBinary()
						_ = yym1584
						if false {
						} else {
							z.F.EncMapIntUint64V(yy1583, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntUint64"))
				r.WriteMapElemValue()
				if yyn1582 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntUint64 == nil {
						r.EncodeNil()
					} else {
						yy1585 := *x.FptrMapIntUint64
						yym1586 := z.EncBinary()
						_ = yym1586
						if false {
						} else {
							z.F.EncMapIntUint64V(yy1585, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntUintptr == nil {
					r.EncodeNil()
				} else {
					yym1588 := z.EncBinary()
					_ = yym1588
					if false {
					} else {
						z.F.EncMapIntUintptrV(x.FMapIntUintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntUintptr"))
				r.WriteMapElemValue()
				if x.FMapIntUintptr == nil {
					r.EncodeNil()
				} else {
					yym1589 := z.EncBinary()
					_ = yym1589
					if false {
					} else {
						z.F.EncMapIntUintptrV(x.FMapIntUintptr, e)
					}
				}
			}
			var yyn1590 bool
			if x.FptrMapIntUintptr == nil {
				yyn1590 = true
				goto LABEL1590
			}
		LABEL1590:
			if yyr2 || yy2arr2 {
				if yyn1590 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntUintptr == nil {
						r.EncodeNil()
					} else {
						yy1591 := *x.FptrMapIntUintptr
						yym1592 := z.EncBinary()
						_ = yym1592
						if false {
						} else {
							z.F.EncMapIntUintptrV(yy1591, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntUintptr"))
				r.WriteMapElemValue()
				if yyn1590 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntUintptr == nil {
						r.EncodeNil()
					} else {
						yy1593 := *x.FptrMapIntUintptr
						yym1594 := z.EncBinary()
						_ = yym1594
						if false {
						} else {
							z.F.EncMapIntUintptrV(yy1593, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntInt == nil {
					r.EncodeNil()
				} else {
					yym1596 := z.EncBinary()
					_ = yym1596
					if false {
					} else {
						z.F.EncMapIntIntV(x.FMapIntInt, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntInt"))
				r.WriteMapElemValue()
				if x.FMapIntInt == nil {
					r.EncodeNil()
				} else {
					yym1597 := z.EncBinary()
					_ = yym1597
					if false {
					} else {
						z.F.EncMapIntIntV(x.FMapIntInt, e)
					}
				}
			}
			var yyn1598 bool
			if x.FptrMapIntInt == nil {
				yyn1598 = true
				goto LABEL1598
			}
		LABEL1598:
			if yyr2 || yy2arr2 {
				if yyn1598 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntInt == nil {
						r.EncodeNil()
					} else {
						yy1599 := *x.FptrMapIntInt
						yym1600 := z.EncBinary()
						_ = yym1600
						if false {
						} else {
							z.F.EncMapIntIntV(yy1599, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntInt"))
				r.WriteMapElemValue()
				if yyn1598 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntInt == nil {
						r.EncodeNil()
					} else {
						yy1601 := *x.FptrMapIntInt
						yym1602 := z.EncBinary()
						_ = yym1602
						if false {
						} else {
							z.F.EncMapIntIntV(yy1601, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntInt8 == nil {
					r.EncodeNil()
				} else {
					yym1604 := z.EncBinary()
					_ = yym1604
					if false {
					} else {
						z.F.EncMapIntInt8V(x.FMapIntInt8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntInt8"))
				r.WriteMapElemValue()
				if x.FMapIntInt8 == nil {
					r.EncodeNil()
				} else {
					yym1605 := z.EncBinary()
					_ = yym1605
					if false {
					} else {
						z.F.EncMapIntInt8V(x.FMapIntInt8, e)
					}
				}
			}
			var yyn1606 bool
			if x.FptrMapIntInt8 == nil {
				yyn1606 = true
				goto LABEL1606
			}
		LABEL1606:
			if yyr2 || yy2arr2 {
				if yyn1606 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntInt8 == nil {
						r.EncodeNil()
					} else {
						yy1607 := *x.FptrMapIntInt8
						yym1608 := z.EncBinary()
						_ = yym1608
						if false {
						} else {
							z.F.EncMapIntInt8V(yy1607, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntInt8"))
				r.WriteMapElemValue()
				if yyn1606 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntInt8 == nil {
						r.EncodeNil()
					} else {
						yy1609 := *x.FptrMapIntInt8
						yym1610 := z.EncBinary()
						_ = yym1610
						if false {
						} else {
							z.F.EncMapIntInt8V(yy1609, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntInt16 == nil {
					r.EncodeNil()
				} else {
					yym1612 := z.EncBinary()
					_ = yym1612
					if false {
					} else {
						z.F.EncMapIntInt16V(x.FMapIntInt16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntInt16"))
				r.WriteMapElemValue()
				if x.FMapIntInt16 == nil {
					r.EncodeNil()
				} else {
					yym1613 := z.EncBinary()
					_ = yym1613
					if false {
					} else {
						z.F.EncMapIntInt16V(x.FMapIntInt16, e)
					}
				}
			}
			var yyn1614 bool
			if x.FptrMapIntInt16 == nil {
				yyn1614 = true
				goto LABEL1614
			}
		LABEL1614:
			if yyr2 || yy2arr2 {
				if yyn1614 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntInt16 == nil {
						r.EncodeNil()
					} else {
						yy1615 := *x.FptrMapIntInt16
						yym1616 := z.EncBinary()
						_ = yym1616
						if false {
						} else {
							z.F.EncMapIntInt16V(yy1615, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntInt16"))
				r.WriteMapElemValue()
				if yyn1614 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntInt16 == nil {
						r.EncodeNil()
					} else {
						yy1617 := *x.FptrMapIntInt16
						yym1618 := z.EncBinary()
						_ = yym1618
						if false {
						} else {
							z.F.EncMapIntInt16V(yy1617, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntInt32 == nil {
					r.EncodeNil()
				} else {
					yym1620 := z.EncBinary()
					_ = yym1620
					if false {
					} else {
						z.F.EncMapIntInt32V(x.FMapIntInt32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntInt32"))
				r.WriteMapElemValue()
				if x.FMapIntInt32 == nil {
					r.EncodeNil()
				} else {
					yym1621 := z.EncBinary()
					_ = yym1621
					if false {
					} else {
						z.F.EncMapIntInt32V(x.FMapIntInt32, e)
					}
				}
			}
			var yyn1622 bool
			if x.FptrMapIntInt32 == nil {
				yyn1622 = true
				goto LABEL1622
			}
		LABEL1622:
			if yyr2 || yy2arr2 {
				if yyn1622 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntInt32 == nil {
						r.EncodeNil()
					} else {
						yy1623 := *x.FptrMapIntInt32
						yym1624 := z.EncBinary()
						_ = yym1624
						if false {
						} else {
							z.F.EncMapIntInt32V(yy1623, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntInt32"))
				r.WriteMapElemValue()
				if yyn1622 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntInt32 == nil {
						r.EncodeNil()
					} else {
						yy1625 := *x.FptrMapIntInt32
						yym1626 := z.EncBinary()
						_ = yym1626
						if false {
						} else {
							z.F.EncMapIntInt32V(yy1625, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntInt64 == nil {
					r.EncodeNil()
				} else {
					yym1628 := z.EncBinary()
					_ = yym1628
					if false {
					} else {
						z.F.EncMapIntInt64V(x.FMapIntInt64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntInt64"))
				r.WriteMapElemValue()
				if x.FMapIntInt64 == nil {
					r.EncodeNil()
				} else {
					yym1629 := z.EncBinary()
					_ = yym1629
					if false {
					} else {
						z.F.EncMapIntInt64V(x.FMapIntInt64, e)
					}
				}
			}
			var yyn1630 bool
			if x.FptrMapIntInt64 == nil {
				yyn1630 = true
				goto LABEL1630
			}
		LABEL1630:
			if yyr2 || yy2arr2 {
				if yyn1630 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntInt64 == nil {
						r.EncodeNil()
					} else {
						yy1631 := *x.FptrMapIntInt64
						yym1632 := z.EncBinary()
						_ = yym1632
						if false {
						} else {
							z.F.EncMapIntInt64V(yy1631, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntInt64"))
				r.WriteMapElemValue()
				if yyn1630 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntInt64 == nil {
						r.EncodeNil()
					} else {
						yy1633 := *x.FptrMapIntInt64
						yym1634 := z.EncBinary()
						_ = yym1634
						if false {
						} else {
							z.F.EncMapIntInt64V(yy1633, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntFloat32 == nil {
					r.EncodeNil()
				} else {
					yym1636 := z.EncBinary()
					_ = yym1636
					if false {
					} else {
						z.F.EncMapIntFloat32V(x.FMapIntFloat32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntFloat32"))
				r.WriteMapElemValue()
				if x.FMapIntFloat32 == nil {
					r.EncodeNil()
				} else {
					yym1637 := z.EncBinary()
					_ = yym1637
					if false {
					} else {
						z.F.EncMapIntFloat32V(x.FMapIntFloat32, e)
					}
				}
			}
			var yyn1638 bool
			if x.FptrMapIntFloat32 == nil {
				yyn1638 = true
				goto LABEL1638
			}
		LABEL1638:
			if yyr2 || yy2arr2 {
				if yyn1638 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntFloat32 == nil {
						r.EncodeNil()
					} else {
						yy1639 := *x.FptrMapIntFloat32
						yym1640 := z.EncBinary()
						_ = yym1640
						if false {
						} else {
							z.F.EncMapIntFloat32V(yy1639, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntFloat32"))
				r.WriteMapElemValue()
				if yyn1638 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntFloat32 == nil {
						r.EncodeNil()
					} else {
						yy1641 := *x.FptrMapIntFloat32
						yym1642 := z.EncBinary()
						_ = yym1642
						if false {
						} else {
							z.F.EncMapIntFloat32V(yy1641, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntFloat64 == nil {
					r.EncodeNil()
				} else {
					yym1644 := z.EncBinary()
					_ = yym1644
					if false {
					} else {
						z.F.EncMapIntFloat64V(x.FMapIntFloat64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntFloat64"))
				r.WriteMapElemValue()
				if x.FMapIntFloat64 == nil {
					r.EncodeNil()
				} else {
					yym1645 := z.EncBinary()
					_ = yym1645
					if false {
					} else {
						z.F.EncMapIntFloat64V(x.FMapIntFloat64, e)
					}
				}
			}
			var yyn1646 bool
			if x.FptrMapIntFloat64 == nil {
				yyn1646 = true
				goto LABEL1646
			}
		LABEL1646:
			if yyr2 || yy2arr2 {
				if yyn1646 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntFloat64 == nil {
						r.EncodeNil()
					} else {
						yy1647 := *x.FptrMapIntFloat64
						yym1648 := z.EncBinary()
						_ = yym1648
						if false {
						} else {
							z.F.EncMapIntFloat64V(yy1647, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntFloat64"))
				r.WriteMapElemValue()
				if yyn1646 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntFloat64 == nil {
						r.EncodeNil()
					} else {
						yy1649 := *x.FptrMapIntFloat64
						yym1650 := z.EncBinary()
						_ = yym1650
						if false {
						} else {
							z.F.EncMapIntFloat64V(yy1649, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapIntBool == nil {
					r.EncodeNil()
				} else {
					yym1652 := z.EncBinary()
					_ = yym1652
					if false {
					} else {
						z.F.EncMapIntBoolV(x.FMapIntBool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapIntBool"))
				r.WriteMapElemValue()
				if x.FMapIntBool == nil {
					r.EncodeNil()
				} else {
					yym1653 := z.EncBinary()
					_ = yym1653
					if false {
					} else {
						z.F.EncMapIntBoolV(x.FMapIntBool, e)
					}
				}
			}
			var yyn1654 bool
			if x.FptrMapIntBool == nil {
				yyn1654 = true
				goto LABEL1654
			}
		LABEL1654:
			if yyr2 || yy2arr2 {
				if yyn1654 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapIntBool == nil {
						r.EncodeNil()
					} else {
						yy1655 := *x.FptrMapIntBool
						yym1656 := z.EncBinary()
						_ = yym1656
						if false {
						} else {
							z.F.EncMapIntBoolV(yy1655, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapIntBool"))
				r.WriteMapElemValue()
				if yyn1654 {
					r.EncodeNil()
				} else {
					if x.FptrMapIntBool == nil {
						r.EncodeNil()
					} else {
						yy1657 := *x.FptrMapIntBool
						yym1658 := z.EncBinary()
						_ = yym1658
						if false {
						} else {
							z.F.EncMapIntBoolV(yy1657, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Intf == nil {
					r.EncodeNil()
				} else {
					yym1660 := z.EncBinary()
					_ = yym1660
					if false {
					} else {
						z.F.EncMapInt8IntfV(x.FMapInt8Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Intf"))
				r.WriteMapElemValue()
				if x.FMapInt8Intf == nil {
					r.EncodeNil()
				} else {
					yym1661 := z.EncBinary()
					_ = yym1661
					if false {
					} else {
						z.F.EncMapInt8IntfV(x.FMapInt8Intf, e)
					}
				}
			}
			var yyn1662 bool
			if x.FptrMapInt8Intf == nil {
				yyn1662 = true
				goto LABEL1662
			}
		LABEL1662:
			if yyr2 || yy2arr2 {
				if yyn1662 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Intf == nil {
						r.EncodeNil()
					} else {
						yy1663 := *x.FptrMapInt8Intf
						yym1664 := z.EncBinary()
						_ = yym1664
						if false {
						} else {
							z.F.EncMapInt8IntfV(yy1663, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Intf"))
				r.WriteMapElemValue()
				if yyn1662 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Intf == nil {
						r.EncodeNil()
					} else {
						yy1665 := *x.FptrMapInt8Intf
						yym1666 := z.EncBinary()
						_ = yym1666
						if false {
						} else {
							z.F.EncMapInt8IntfV(yy1665, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8String == nil {
					r.EncodeNil()
				} else {
					yym1668 := z.EncBinary()
					_ = yym1668
					if false {
					} else {
						z.F.EncMapInt8StringV(x.FMapInt8String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8String"))
				r.WriteMapElemValue()
				if x.FMapInt8String == nil {
					r.EncodeNil()
				} else {
					yym1669 := z.EncBinary()
					_ = yym1669
					if false {
					} else {
						z.F.EncMapInt8StringV(x.FMapInt8String, e)
					}
				}
			}
			var yyn1670 bool
			if x.FptrMapInt8String == nil {
				yyn1670 = true
				goto LABEL1670
			}
		LABEL1670:
			if yyr2 || yy2arr2 {
				if yyn1670 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8String == nil {
						r.EncodeNil()
					} else {
						yy1671 := *x.FptrMapInt8String
						yym1672 := z.EncBinary()
						_ = yym1672
						if false {
						} else {
							z.F.EncMapInt8StringV(yy1671, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8String"))
				r.WriteMapElemValue()
				if yyn1670 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8String == nil {
						r.EncodeNil()
					} else {
						yy1673 := *x.FptrMapInt8String
						yym1674 := z.EncBinary()
						_ = yym1674
						if false {
						} else {
							z.F.EncMapInt8StringV(yy1673, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Uint == nil {
					r.EncodeNil()
				} else {
					yym1676 := z.EncBinary()
					_ = yym1676
					if false {
					} else {
						z.F.EncMapInt8UintV(x.FMapInt8Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Uint"))
				r.WriteMapElemValue()
				if x.FMapInt8Uint == nil {
					r.EncodeNil()
				} else {
					yym1677 := z.EncBinary()
					_ = yym1677
					if false {
					} else {
						z.F.EncMapInt8UintV(x.FMapInt8Uint, e)
					}
				}
			}
			var yyn1678 bool
			if x.FptrMapInt8Uint == nil {
				yyn1678 = true
				goto LABEL1678
			}
		LABEL1678:
			if yyr2 || yy2arr2 {
				if yyn1678 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Uint == nil {
						r.EncodeNil()
					} else {
						yy1679 := *x.FptrMapInt8Uint
						yym1680 := z.EncBinary()
						_ = yym1680
						if false {
						} else {
							z.F.EncMapInt8UintV(yy1679, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Uint"))
				r.WriteMapElemValue()
				if yyn1678 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Uint == nil {
						r.EncodeNil()
					} else {
						yy1681 := *x.FptrMapInt8Uint
						yym1682 := z.EncBinary()
						_ = yym1682
						if false {
						} else {
							z.F.EncMapInt8UintV(yy1681, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1684 := z.EncBinary()
					_ = yym1684
					if false {
					} else {
						z.F.EncMapInt8Uint8V(x.FMapInt8Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Uint8"))
				r.WriteMapElemValue()
				if x.FMapInt8Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1685 := z.EncBinary()
					_ = yym1685
					if false {
					} else {
						z.F.EncMapInt8Uint8V(x.FMapInt8Uint8, e)
					}
				}
			}
			var yyn1686 bool
			if x.FptrMapInt8Uint8 == nil {
				yyn1686 = true
				goto LABEL1686
			}
		LABEL1686:
			if yyr2 || yy2arr2 {
				if yyn1686 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1687 := *x.FptrMapInt8Uint8
						yym1688 := z.EncBinary()
						_ = yym1688
						if false {
						} else {
							z.F.EncMapInt8Uint8V(yy1687, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Uint8"))
				r.WriteMapElemValue()
				if yyn1686 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1689 := *x.FptrMapInt8Uint8
						yym1690 := z.EncBinary()
						_ = yym1690
						if false {
						} else {
							z.F.EncMapInt8Uint8V(yy1689, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1692 := z.EncBinary()
					_ = yym1692
					if false {
					} else {
						z.F.EncMapInt8Uint16V(x.FMapInt8Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Uint16"))
				r.WriteMapElemValue()
				if x.FMapInt8Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1693 := z.EncBinary()
					_ = yym1693
					if false {
					} else {
						z.F.EncMapInt8Uint16V(x.FMapInt8Uint16, e)
					}
				}
			}
			var yyn1694 bool
			if x.FptrMapInt8Uint16 == nil {
				yyn1694 = true
				goto LABEL1694
			}
		LABEL1694:
			if yyr2 || yy2arr2 {
				if yyn1694 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1695 := *x.FptrMapInt8Uint16
						yym1696 := z.EncBinary()
						_ = yym1696
						if false {
						} else {
							z.F.EncMapInt8Uint16V(yy1695, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Uint16"))
				r.WriteMapElemValue()
				if yyn1694 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1697 := *x.FptrMapInt8Uint16
						yym1698 := z.EncBinary()
						_ = yym1698
						if false {
						} else {
							z.F.EncMapInt8Uint16V(yy1697, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1700 := z.EncBinary()
					_ = yym1700
					if false {
					} else {
						z.F.EncMapInt8Uint32V(x.FMapInt8Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Uint32"))
				r.WriteMapElemValue()
				if x.FMapInt8Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1701 := z.EncBinary()
					_ = yym1701
					if false {
					} else {
						z.F.EncMapInt8Uint32V(x.FMapInt8Uint32, e)
					}
				}
			}
			var yyn1702 bool
			if x.FptrMapInt8Uint32 == nil {
				yyn1702 = true
				goto LABEL1702
			}
		LABEL1702:
			if yyr2 || yy2arr2 {
				if yyn1702 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1703 := *x.FptrMapInt8Uint32
						yym1704 := z.EncBinary()
						_ = yym1704
						if false {
						} else {
							z.F.EncMapInt8Uint32V(yy1703, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Uint32"))
				r.WriteMapElemValue()
				if yyn1702 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1705 := *x.FptrMapInt8Uint32
						yym1706 := z.EncBinary()
						_ = yym1706
						if false {
						} else {
							z.F.EncMapInt8Uint32V(yy1705, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1708 := z.EncBinary()
					_ = yym1708
					if false {
					} else {
						z.F.EncMapInt8Uint64V(x.FMapInt8Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Uint64"))
				r.WriteMapElemValue()
				if x.FMapInt8Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1709 := z.EncBinary()
					_ = yym1709
					if false {
					} else {
						z.F.EncMapInt8Uint64V(x.FMapInt8Uint64, e)
					}
				}
			}
			var yyn1710 bool
			if x.FptrMapInt8Uint64 == nil {
				yyn1710 = true
				goto LABEL1710
			}
		LABEL1710:
			if yyr2 || yy2arr2 {
				if yyn1710 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1711 := *x.FptrMapInt8Uint64
						yym1712 := z.EncBinary()
						_ = yym1712
						if false {
						} else {
							z.F.EncMapInt8Uint64V(yy1711, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Uint64"))
				r.WriteMapElemValue()
				if yyn1710 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1713 := *x.FptrMapInt8Uint64
						yym1714 := z.EncBinary()
						_ = yym1714
						if false {
						} else {
							z.F.EncMapInt8Uint64V(yy1713, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1716 := z.EncBinary()
					_ = yym1716
					if false {
					} else {
						z.F.EncMapInt8UintptrV(x.FMapInt8Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Uintptr"))
				r.WriteMapElemValue()
				if x.FMapInt8Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1717 := z.EncBinary()
					_ = yym1717
					if false {
					} else {
						z.F.EncMapInt8UintptrV(x.FMapInt8Uintptr, e)
					}
				}
			}
			var yyn1718 bool
			if x.FptrMapInt8Uintptr == nil {
				yyn1718 = true
				goto LABEL1718
			}
		LABEL1718:
			if yyr2 || yy2arr2 {
				if yyn1718 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1719 := *x.FptrMapInt8Uintptr
						yym1720 := z.EncBinary()
						_ = yym1720
						if false {
						} else {
							z.F.EncMapInt8UintptrV(yy1719, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Uintptr"))
				r.WriteMapElemValue()
				if yyn1718 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1721 := *x.FptrMapInt8Uintptr
						yym1722 := z.EncBinary()
						_ = yym1722
						if false {
						} else {
							z.F.EncMapInt8UintptrV(yy1721, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Int == nil {
					r.EncodeNil()
				} else {
					yym1724 := z.EncBinary()
					_ = yym1724
					if false {
					} else {
						z.F.EncMapInt8IntV(x.FMapInt8Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Int"))
				r.WriteMapElemValue()
				if x.FMapInt8Int == nil {
					r.EncodeNil()
				} else {
					yym1725 := z.EncBinary()
					_ = yym1725
					if false {
					} else {
						z.F.EncMapInt8IntV(x.FMapInt8Int, e)
					}
				}
			}
			var yyn1726 bool
			if x.FptrMapInt8Int == nil {
				yyn1726 = true
				goto LABEL1726
			}
		LABEL1726:
			if yyr2 || yy2arr2 {
				if yyn1726 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Int == nil {
						r.EncodeNil()
					} else {
						yy1727 := *x.FptrMapInt8Int
						yym1728 := z.EncBinary()
						_ = yym1728
						if false {
						} else {
							z.F.EncMapInt8IntV(yy1727, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Int"))
				r.WriteMapElemValue()
				if yyn1726 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Int == nil {
						r.EncodeNil()
					} else {
						yy1729 := *x.FptrMapInt8Int
						yym1730 := z.EncBinary()
						_ = yym1730
						if false {
						} else {
							z.F.EncMapInt8IntV(yy1729, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Int8 == nil {
					r.EncodeNil()
				} else {
					yym1732 := z.EncBinary()
					_ = yym1732
					if false {
					} else {
						z.F.EncMapInt8Int8V(x.FMapInt8Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Int8"))
				r.WriteMapElemValue()
				if x.FMapInt8Int8 == nil {
					r.EncodeNil()
				} else {
					yym1733 := z.EncBinary()
					_ = yym1733
					if false {
					} else {
						z.F.EncMapInt8Int8V(x.FMapInt8Int8, e)
					}
				}
			}
			var yyn1734 bool
			if x.FptrMapInt8Int8 == nil {
				yyn1734 = true
				goto LABEL1734
			}
		LABEL1734:
			if yyr2 || yy2arr2 {
				if yyn1734 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Int8 == nil {
						r.EncodeNil()
					} else {
						yy1735 := *x.FptrMapInt8Int8
						yym1736 := z.EncBinary()
						_ = yym1736
						if false {
						} else {
							z.F.EncMapInt8Int8V(yy1735, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Int8"))
				r.WriteMapElemValue()
				if yyn1734 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Int8 == nil {
						r.EncodeNil()
					} else {
						yy1737 := *x.FptrMapInt8Int8
						yym1738 := z.EncBinary()
						_ = yym1738
						if false {
						} else {
							z.F.EncMapInt8Int8V(yy1737, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Int16 == nil {
					r.EncodeNil()
				} else {
					yym1740 := z.EncBinary()
					_ = yym1740
					if false {
					} else {
						z.F.EncMapInt8Int16V(x.FMapInt8Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Int16"))
				r.WriteMapElemValue()
				if x.FMapInt8Int16 == nil {
					r.EncodeNil()
				} else {
					yym1741 := z.EncBinary()
					_ = yym1741
					if false {
					} else {
						z.F.EncMapInt8Int16V(x.FMapInt8Int16, e)
					}
				}
			}
			var yyn1742 bool
			if x.FptrMapInt8Int16 == nil {
				yyn1742 = true
				goto LABEL1742
			}
		LABEL1742:
			if yyr2 || yy2arr2 {
				if yyn1742 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Int16 == nil {
						r.EncodeNil()
					} else {
						yy1743 := *x.FptrMapInt8Int16
						yym1744 := z.EncBinary()
						_ = yym1744
						if false {
						} else {
							z.F.EncMapInt8Int16V(yy1743, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Int16"))
				r.WriteMapElemValue()
				if yyn1742 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Int16 == nil {
						r.EncodeNil()
					} else {
						yy1745 := *x.FptrMapInt8Int16
						yym1746 := z.EncBinary()
						_ = yym1746
						if false {
						} else {
							z.F.EncMapInt8Int16V(yy1745, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Int32 == nil {
					r.EncodeNil()
				} else {
					yym1748 := z.EncBinary()
					_ = yym1748
					if false {
					} else {
						z.F.EncMapInt8Int32V(x.FMapInt8Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Int32"))
				r.WriteMapElemValue()
				if x.FMapInt8Int32 == nil {
					r.EncodeNil()
				} else {
					yym1749 := z.EncBinary()
					_ = yym1749
					if false {
					} else {
						z.F.EncMapInt8Int32V(x.FMapInt8Int32, e)
					}
				}
			}
			var yyn1750 bool
			if x.FptrMapInt8Int32 == nil {
				yyn1750 = true
				goto LABEL1750
			}
		LABEL1750:
			if yyr2 || yy2arr2 {
				if yyn1750 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Int32 == nil {
						r.EncodeNil()
					} else {
						yy1751 := *x.FptrMapInt8Int32
						yym1752 := z.EncBinary()
						_ = yym1752
						if false {
						} else {
							z.F.EncMapInt8Int32V(yy1751, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Int32"))
				r.WriteMapElemValue()
				if yyn1750 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Int32 == nil {
						r.EncodeNil()
					} else {
						yy1753 := *x.FptrMapInt8Int32
						yym1754 := z.EncBinary()
						_ = yym1754
						if false {
						} else {
							z.F.EncMapInt8Int32V(yy1753, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Int64 == nil {
					r.EncodeNil()
				} else {
					yym1756 := z.EncBinary()
					_ = yym1756
					if false {
					} else {
						z.F.EncMapInt8Int64V(x.FMapInt8Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Int64"))
				r.WriteMapElemValue()
				if x.FMapInt8Int64 == nil {
					r.EncodeNil()
				} else {
					yym1757 := z.EncBinary()
					_ = yym1757
					if false {
					} else {
						z.F.EncMapInt8Int64V(x.FMapInt8Int64, e)
					}
				}
			}
			var yyn1758 bool
			if x.FptrMapInt8Int64 == nil {
				yyn1758 = true
				goto LABEL1758
			}
		LABEL1758:
			if yyr2 || yy2arr2 {
				if yyn1758 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Int64 == nil {
						r.EncodeNil()
					} else {
						yy1759 := *x.FptrMapInt8Int64
						yym1760 := z.EncBinary()
						_ = yym1760
						if false {
						} else {
							z.F.EncMapInt8Int64V(yy1759, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Int64"))
				r.WriteMapElemValue()
				if yyn1758 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Int64 == nil {
						r.EncodeNil()
					} else {
						yy1761 := *x.FptrMapInt8Int64
						yym1762 := z.EncBinary()
						_ = yym1762
						if false {
						} else {
							z.F.EncMapInt8Int64V(yy1761, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Float32 == nil {
					r.EncodeNil()
				} else {
					yym1764 := z.EncBinary()
					_ = yym1764
					if false {
					} else {
						z.F.EncMapInt8Float32V(x.FMapInt8Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Float32"))
				r.WriteMapElemValue()
				if x.FMapInt8Float32 == nil {
					r.EncodeNil()
				} else {
					yym1765 := z.EncBinary()
					_ = yym1765
					if false {
					} else {
						z.F.EncMapInt8Float32V(x.FMapInt8Float32, e)
					}
				}
			}
			var yyn1766 bool
			if x.FptrMapInt8Float32 == nil {
				yyn1766 = true
				goto LABEL1766
			}
		LABEL1766:
			if yyr2 || yy2arr2 {
				if yyn1766 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Float32 == nil {
						r.EncodeNil()
					} else {
						yy1767 := *x.FptrMapInt8Float32
						yym1768 := z.EncBinary()
						_ = yym1768
						if false {
						} else {
							z.F.EncMapInt8Float32V(yy1767, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Float32"))
				r.WriteMapElemValue()
				if yyn1766 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Float32 == nil {
						r.EncodeNil()
					} else {
						yy1769 := *x.FptrMapInt8Float32
						yym1770 := z.EncBinary()
						_ = yym1770
						if false {
						} else {
							z.F.EncMapInt8Float32V(yy1769, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Float64 == nil {
					r.EncodeNil()
				} else {
					yym1772 := z.EncBinary()
					_ = yym1772
					if false {
					} else {
						z.F.EncMapInt8Float64V(x.FMapInt8Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Float64"))
				r.WriteMapElemValue()
				if x.FMapInt8Float64 == nil {
					r.EncodeNil()
				} else {
					yym1773 := z.EncBinary()
					_ = yym1773
					if false {
					} else {
						z.F.EncMapInt8Float64V(x.FMapInt8Float64, e)
					}
				}
			}
			var yyn1774 bool
			if x.FptrMapInt8Float64 == nil {
				yyn1774 = true
				goto LABEL1774
			}
		LABEL1774:
			if yyr2 || yy2arr2 {
				if yyn1774 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Float64 == nil {
						r.EncodeNil()
					} else {
						yy1775 := *x.FptrMapInt8Float64
						yym1776 := z.EncBinary()
						_ = yym1776
						if false {
						} else {
							z.F.EncMapInt8Float64V(yy1775, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Float64"))
				r.WriteMapElemValue()
				if yyn1774 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Float64 == nil {
						r.EncodeNil()
					} else {
						yy1777 := *x.FptrMapInt8Float64
						yym1778 := z.EncBinary()
						_ = yym1778
						if false {
						} else {
							z.F.EncMapInt8Float64V(yy1777, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt8Bool == nil {
					r.EncodeNil()
				} else {
					yym1780 := z.EncBinary()
					_ = yym1780
					if false {
					} else {
						z.F.EncMapInt8BoolV(x.FMapInt8Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt8Bool"))
				r.WriteMapElemValue()
				if x.FMapInt8Bool == nil {
					r.EncodeNil()
				} else {
					yym1781 := z.EncBinary()
					_ = yym1781
					if false {
					} else {
						z.F.EncMapInt8BoolV(x.FMapInt8Bool, e)
					}
				}
			}
			var yyn1782 bool
			if x.FptrMapInt8Bool == nil {
				yyn1782 = true
				goto LABEL1782
			}
		LABEL1782:
			if yyr2 || yy2arr2 {
				if yyn1782 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt8Bool == nil {
						r.EncodeNil()
					} else {
						yy1783 := *x.FptrMapInt8Bool
						yym1784 := z.EncBinary()
						_ = yym1784
						if false {
						} else {
							z.F.EncMapInt8BoolV(yy1783, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt8Bool"))
				r.WriteMapElemValue()
				if yyn1782 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt8Bool == nil {
						r.EncodeNil()
					} else {
						yy1785 := *x.FptrMapInt8Bool
						yym1786 := z.EncBinary()
						_ = yym1786
						if false {
						} else {
							z.F.EncMapInt8BoolV(yy1785, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Intf == nil {
					r.EncodeNil()
				} else {
					yym1788 := z.EncBinary()
					_ = yym1788
					if false {
					} else {
						z.F.EncMapInt16IntfV(x.FMapInt16Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Intf"))
				r.WriteMapElemValue()
				if x.FMapInt16Intf == nil {
					r.EncodeNil()
				} else {
					yym1789 := z.EncBinary()
					_ = yym1789
					if false {
					} else {
						z.F.EncMapInt16IntfV(x.FMapInt16Intf, e)
					}
				}
			}
			var yyn1790 bool
			if x.FptrMapInt16Intf == nil {
				yyn1790 = true
				goto LABEL1790
			}
		LABEL1790:
			if yyr2 || yy2arr2 {
				if yyn1790 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Intf == nil {
						r.EncodeNil()
					} else {
						yy1791 := *x.FptrMapInt16Intf
						yym1792 := z.EncBinary()
						_ = yym1792
						if false {
						} else {
							z.F.EncMapInt16IntfV(yy1791, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Intf"))
				r.WriteMapElemValue()
				if yyn1790 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Intf == nil {
						r.EncodeNil()
					} else {
						yy1793 := *x.FptrMapInt16Intf
						yym1794 := z.EncBinary()
						_ = yym1794
						if false {
						} else {
							z.F.EncMapInt16IntfV(yy1793, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16String == nil {
					r.EncodeNil()
				} else {
					yym1796 := z.EncBinary()
					_ = yym1796
					if false {
					} else {
						z.F.EncMapInt16StringV(x.FMapInt16String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16String"))
				r.WriteMapElemValue()
				if x.FMapInt16String == nil {
					r.EncodeNil()
				} else {
					yym1797 := z.EncBinary()
					_ = yym1797
					if false {
					} else {
						z.F.EncMapInt16StringV(x.FMapInt16String, e)
					}
				}
			}
			var yyn1798 bool
			if x.FptrMapInt16String == nil {
				yyn1798 = true
				goto LABEL1798
			}
		LABEL1798:
			if yyr2 || yy2arr2 {
				if yyn1798 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16String == nil {
						r.EncodeNil()
					} else {
						yy1799 := *x.FptrMapInt16String
						yym1800 := z.EncBinary()
						_ = yym1800
						if false {
						} else {
							z.F.EncMapInt16StringV(yy1799, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16String"))
				r.WriteMapElemValue()
				if yyn1798 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16String == nil {
						r.EncodeNil()
					} else {
						yy1801 := *x.FptrMapInt16String
						yym1802 := z.EncBinary()
						_ = yym1802
						if false {
						} else {
							z.F.EncMapInt16StringV(yy1801, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Uint == nil {
					r.EncodeNil()
				} else {
					yym1804 := z.EncBinary()
					_ = yym1804
					if false {
					} else {
						z.F.EncMapInt16UintV(x.FMapInt16Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Uint"))
				r.WriteMapElemValue()
				if x.FMapInt16Uint == nil {
					r.EncodeNil()
				} else {
					yym1805 := z.EncBinary()
					_ = yym1805
					if false {
					} else {
						z.F.EncMapInt16UintV(x.FMapInt16Uint, e)
					}
				}
			}
			var yyn1806 bool
			if x.FptrMapInt16Uint == nil {
				yyn1806 = true
				goto LABEL1806
			}
		LABEL1806:
			if yyr2 || yy2arr2 {
				if yyn1806 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Uint == nil {
						r.EncodeNil()
					} else {
						yy1807 := *x.FptrMapInt16Uint
						yym1808 := z.EncBinary()
						_ = yym1808
						if false {
						} else {
							z.F.EncMapInt16UintV(yy1807, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Uint"))
				r.WriteMapElemValue()
				if yyn1806 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Uint == nil {
						r.EncodeNil()
					} else {
						yy1809 := *x.FptrMapInt16Uint
						yym1810 := z.EncBinary()
						_ = yym1810
						if false {
						} else {
							z.F.EncMapInt16UintV(yy1809, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1812 := z.EncBinary()
					_ = yym1812
					if false {
					} else {
						z.F.EncMapInt16Uint8V(x.FMapInt16Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Uint8"))
				r.WriteMapElemValue()
				if x.FMapInt16Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1813 := z.EncBinary()
					_ = yym1813
					if false {
					} else {
						z.F.EncMapInt16Uint8V(x.FMapInt16Uint8, e)
					}
				}
			}
			var yyn1814 bool
			if x.FptrMapInt16Uint8 == nil {
				yyn1814 = true
				goto LABEL1814
			}
		LABEL1814:
			if yyr2 || yy2arr2 {
				if yyn1814 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1815 := *x.FptrMapInt16Uint8
						yym1816 := z.EncBinary()
						_ = yym1816
						if false {
						} else {
							z.F.EncMapInt16Uint8V(yy1815, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Uint8"))
				r.WriteMapElemValue()
				if yyn1814 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1817 := *x.FptrMapInt16Uint8
						yym1818 := z.EncBinary()
						_ = yym1818
						if false {
						} else {
							z.F.EncMapInt16Uint8V(yy1817, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1820 := z.EncBinary()
					_ = yym1820
					if false {
					} else {
						z.F.EncMapInt16Uint16V(x.FMapInt16Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Uint16"))
				r.WriteMapElemValue()
				if x.FMapInt16Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1821 := z.EncBinary()
					_ = yym1821
					if false {
					} else {
						z.F.EncMapInt16Uint16V(x.FMapInt16Uint16, e)
					}
				}
			}
			var yyn1822 bool
			if x.FptrMapInt16Uint16 == nil {
				yyn1822 = true
				goto LABEL1822
			}
		LABEL1822:
			if yyr2 || yy2arr2 {
				if yyn1822 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1823 := *x.FptrMapInt16Uint16
						yym1824 := z.EncBinary()
						_ = yym1824
						if false {
						} else {
							z.F.EncMapInt16Uint16V(yy1823, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Uint16"))
				r.WriteMapElemValue()
				if yyn1822 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1825 := *x.FptrMapInt16Uint16
						yym1826 := z.EncBinary()
						_ = yym1826
						if false {
						} else {
							z.F.EncMapInt16Uint16V(yy1825, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1828 := z.EncBinary()
					_ = yym1828
					if false {
					} else {
						z.F.EncMapInt16Uint32V(x.FMapInt16Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Uint32"))
				r.WriteMapElemValue()
				if x.FMapInt16Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1829 := z.EncBinary()
					_ = yym1829
					if false {
					} else {
						z.F.EncMapInt16Uint32V(x.FMapInt16Uint32, e)
					}
				}
			}
			var yyn1830 bool
			if x.FptrMapInt16Uint32 == nil {
				yyn1830 = true
				goto LABEL1830
			}
		LABEL1830:
			if yyr2 || yy2arr2 {
				if yyn1830 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1831 := *x.FptrMapInt16Uint32
						yym1832 := z.EncBinary()
						_ = yym1832
						if false {
						} else {
							z.F.EncMapInt16Uint32V(yy1831, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Uint32"))
				r.WriteMapElemValue()
				if yyn1830 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1833 := *x.FptrMapInt16Uint32
						yym1834 := z.EncBinary()
						_ = yym1834
						if false {
						} else {
							z.F.EncMapInt16Uint32V(yy1833, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1836 := z.EncBinary()
					_ = yym1836
					if false {
					} else {
						z.F.EncMapInt16Uint64V(x.FMapInt16Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Uint64"))
				r.WriteMapElemValue()
				if x.FMapInt16Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1837 := z.EncBinary()
					_ = yym1837
					if false {
					} else {
						z.F.EncMapInt16Uint64V(x.FMapInt16Uint64, e)
					}
				}
			}
			var yyn1838 bool
			if x.FptrMapInt16Uint64 == nil {
				yyn1838 = true
				goto LABEL1838
			}
		LABEL1838:
			if yyr2 || yy2arr2 {
				if yyn1838 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1839 := *x.FptrMapInt16Uint64
						yym1840 := z.EncBinary()
						_ = yym1840
						if false {
						} else {
							z.F.EncMapInt16Uint64V(yy1839, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Uint64"))
				r.WriteMapElemValue()
				if yyn1838 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1841 := *x.FptrMapInt16Uint64
						yym1842 := z.EncBinary()
						_ = yym1842
						if false {
						} else {
							z.F.EncMapInt16Uint64V(yy1841, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1844 := z.EncBinary()
					_ = yym1844
					if false {
					} else {
						z.F.EncMapInt16UintptrV(x.FMapInt16Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Uintptr"))
				r.WriteMapElemValue()
				if x.FMapInt16Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1845 := z.EncBinary()
					_ = yym1845
					if false {
					} else {
						z.F.EncMapInt16UintptrV(x.FMapInt16Uintptr, e)
					}
				}
			}
			var yyn1846 bool
			if x.FptrMapInt16Uintptr == nil {
				yyn1846 = true
				goto LABEL1846
			}
		LABEL1846:
			if yyr2 || yy2arr2 {
				if yyn1846 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1847 := *x.FptrMapInt16Uintptr
						yym1848 := z.EncBinary()
						_ = yym1848
						if false {
						} else {
							z.F.EncMapInt16UintptrV(yy1847, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Uintptr"))
				r.WriteMapElemValue()
				if yyn1846 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1849 := *x.FptrMapInt16Uintptr
						yym1850 := z.EncBinary()
						_ = yym1850
						if false {
						} else {
							z.F.EncMapInt16UintptrV(yy1849, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Int == nil {
					r.EncodeNil()
				} else {
					yym1852 := z.EncBinary()
					_ = yym1852
					if false {
					} else {
						z.F.EncMapInt16IntV(x.FMapInt16Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Int"))
				r.WriteMapElemValue()
				if x.FMapInt16Int == nil {
					r.EncodeNil()
				} else {
					yym1853 := z.EncBinary()
					_ = yym1853
					if false {
					} else {
						z.F.EncMapInt16IntV(x.FMapInt16Int, e)
					}
				}
			}
			var yyn1854 bool
			if x.FptrMapInt16Int == nil {
				yyn1854 = true
				goto LABEL1854
			}
		LABEL1854:
			if yyr2 || yy2arr2 {
				if yyn1854 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Int == nil {
						r.EncodeNil()
					} else {
						yy1855 := *x.FptrMapInt16Int
						yym1856 := z.EncBinary()
						_ = yym1856
						if false {
						} else {
							z.F.EncMapInt16IntV(yy1855, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Int"))
				r.WriteMapElemValue()
				if yyn1854 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Int == nil {
						r.EncodeNil()
					} else {
						yy1857 := *x.FptrMapInt16Int
						yym1858 := z.EncBinary()
						_ = yym1858
						if false {
						} else {
							z.F.EncMapInt16IntV(yy1857, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Int8 == nil {
					r.EncodeNil()
				} else {
					yym1860 := z.EncBinary()
					_ = yym1860
					if false {
					} else {
						z.F.EncMapInt16Int8V(x.FMapInt16Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Int8"))
				r.WriteMapElemValue()
				if x.FMapInt16Int8 == nil {
					r.EncodeNil()
				} else {
					yym1861 := z.EncBinary()
					_ = yym1861
					if false {
					} else {
						z.F.EncMapInt16Int8V(x.FMapInt16Int8, e)
					}
				}
			}
			var yyn1862 bool
			if x.FptrMapInt16Int8 == nil {
				yyn1862 = true
				goto LABEL1862
			}
		LABEL1862:
			if yyr2 || yy2arr2 {
				if yyn1862 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Int8 == nil {
						r.EncodeNil()
					} else {
						yy1863 := *x.FptrMapInt16Int8
						yym1864 := z.EncBinary()
						_ = yym1864
						if false {
						} else {
							z.F.EncMapInt16Int8V(yy1863, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Int8"))
				r.WriteMapElemValue()
				if yyn1862 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Int8 == nil {
						r.EncodeNil()
					} else {
						yy1865 := *x.FptrMapInt16Int8
						yym1866 := z.EncBinary()
						_ = yym1866
						if false {
						} else {
							z.F.EncMapInt16Int8V(yy1865, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Int16 == nil {
					r.EncodeNil()
				} else {
					yym1868 := z.EncBinary()
					_ = yym1868
					if false {
					} else {
						z.F.EncMapInt16Int16V(x.FMapInt16Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Int16"))
				r.WriteMapElemValue()
				if x.FMapInt16Int16 == nil {
					r.EncodeNil()
				} else {
					yym1869 := z.EncBinary()
					_ = yym1869
					if false {
					} else {
						z.F.EncMapInt16Int16V(x.FMapInt16Int16, e)
					}
				}
			}
			var yyn1870 bool
			if x.FptrMapInt16Int16 == nil {
				yyn1870 = true
				goto LABEL1870
			}
		LABEL1870:
			if yyr2 || yy2arr2 {
				if yyn1870 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Int16 == nil {
						r.EncodeNil()
					} else {
						yy1871 := *x.FptrMapInt16Int16
						yym1872 := z.EncBinary()
						_ = yym1872
						if false {
						} else {
							z.F.EncMapInt16Int16V(yy1871, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Int16"))
				r.WriteMapElemValue()
				if yyn1870 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Int16 == nil {
						r.EncodeNil()
					} else {
						yy1873 := *x.FptrMapInt16Int16
						yym1874 := z.EncBinary()
						_ = yym1874
						if false {
						} else {
							z.F.EncMapInt16Int16V(yy1873, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Int32 == nil {
					r.EncodeNil()
				} else {
					yym1876 := z.EncBinary()
					_ = yym1876
					if false {
					} else {
						z.F.EncMapInt16Int32V(x.FMapInt16Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Int32"))
				r.WriteMapElemValue()
				if x.FMapInt16Int32 == nil {
					r.EncodeNil()
				} else {
					yym1877 := z.EncBinary()
					_ = yym1877
					if false {
					} else {
						z.F.EncMapInt16Int32V(x.FMapInt16Int32, e)
					}
				}
			}
			var yyn1878 bool
			if x.FptrMapInt16Int32 == nil {
				yyn1878 = true
				goto LABEL1878
			}
		LABEL1878:
			if yyr2 || yy2arr2 {
				if yyn1878 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Int32 == nil {
						r.EncodeNil()
					} else {
						yy1879 := *x.FptrMapInt16Int32
						yym1880 := z.EncBinary()
						_ = yym1880
						if false {
						} else {
							z.F.EncMapInt16Int32V(yy1879, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Int32"))
				r.WriteMapElemValue()
				if yyn1878 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Int32 == nil {
						r.EncodeNil()
					} else {
						yy1881 := *x.FptrMapInt16Int32
						yym1882 := z.EncBinary()
						_ = yym1882
						if false {
						} else {
							z.F.EncMapInt16Int32V(yy1881, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Int64 == nil {
					r.EncodeNil()
				} else {
					yym1884 := z.EncBinary()
					_ = yym1884
					if false {
					} else {
						z.F.EncMapInt16Int64V(x.FMapInt16Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Int64"))
				r.WriteMapElemValue()
				if x.FMapInt16Int64 == nil {
					r.EncodeNil()
				} else {
					yym1885 := z.EncBinary()
					_ = yym1885
					if false {
					} else {
						z.F.EncMapInt16Int64V(x.FMapInt16Int64, e)
					}
				}
			}
			var yyn1886 bool
			if x.FptrMapInt16Int64 == nil {
				yyn1886 = true
				goto LABEL1886
			}
		LABEL1886:
			if yyr2 || yy2arr2 {
				if yyn1886 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Int64 == nil {
						r.EncodeNil()
					} else {
						yy1887 := *x.FptrMapInt16Int64
						yym1888 := z.EncBinary()
						_ = yym1888
						if false {
						} else {
							z.F.EncMapInt16Int64V(yy1887, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Int64"))
				r.WriteMapElemValue()
				if yyn1886 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Int64 == nil {
						r.EncodeNil()
					} else {
						yy1889 := *x.FptrMapInt16Int64
						yym1890 := z.EncBinary()
						_ = yym1890
						if false {
						} else {
							z.F.EncMapInt16Int64V(yy1889, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Float32 == nil {
					r.EncodeNil()
				} else {
					yym1892 := z.EncBinary()
					_ = yym1892
					if false {
					} else {
						z.F.EncMapInt16Float32V(x.FMapInt16Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Float32"))
				r.WriteMapElemValue()
				if x.FMapInt16Float32 == nil {
					r.EncodeNil()
				} else {
					yym1893 := z.EncBinary()
					_ = yym1893
					if false {
					} else {
						z.F.EncMapInt16Float32V(x.FMapInt16Float32, e)
					}
				}
			}
			var yyn1894 bool
			if x.FptrMapInt16Float32 == nil {
				yyn1894 = true
				goto LABEL1894
			}
		LABEL1894:
			if yyr2 || yy2arr2 {
				if yyn1894 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Float32 == nil {
						r.EncodeNil()
					} else {
						yy1895 := *x.FptrMapInt16Float32
						yym1896 := z.EncBinary()
						_ = yym1896
						if false {
						} else {
							z.F.EncMapInt16Float32V(yy1895, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Float32"))
				r.WriteMapElemValue()
				if yyn1894 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Float32 == nil {
						r.EncodeNil()
					} else {
						yy1897 := *x.FptrMapInt16Float32
						yym1898 := z.EncBinary()
						_ = yym1898
						if false {
						} else {
							z.F.EncMapInt16Float32V(yy1897, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Float64 == nil {
					r.EncodeNil()
				} else {
					yym1900 := z.EncBinary()
					_ = yym1900
					if false {
					} else {
						z.F.EncMapInt16Float64V(x.FMapInt16Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Float64"))
				r.WriteMapElemValue()
				if x.FMapInt16Float64 == nil {
					r.EncodeNil()
				} else {
					yym1901 := z.EncBinary()
					_ = yym1901
					if false {
					} else {
						z.F.EncMapInt16Float64V(x.FMapInt16Float64, e)
					}
				}
			}
			var yyn1902 bool
			if x.FptrMapInt16Float64 == nil {
				yyn1902 = true
				goto LABEL1902
			}
		LABEL1902:
			if yyr2 || yy2arr2 {
				if yyn1902 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Float64 == nil {
						r.EncodeNil()
					} else {
						yy1903 := *x.FptrMapInt16Float64
						yym1904 := z.EncBinary()
						_ = yym1904
						if false {
						} else {
							z.F.EncMapInt16Float64V(yy1903, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Float64"))
				r.WriteMapElemValue()
				if yyn1902 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Float64 == nil {
						r.EncodeNil()
					} else {
						yy1905 := *x.FptrMapInt16Float64
						yym1906 := z.EncBinary()
						_ = yym1906
						if false {
						} else {
							z.F.EncMapInt16Float64V(yy1905, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt16Bool == nil {
					r.EncodeNil()
				} else {
					yym1908 := z.EncBinary()
					_ = yym1908
					if false {
					} else {
						z.F.EncMapInt16BoolV(x.FMapInt16Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt16Bool"))
				r.WriteMapElemValue()
				if x.FMapInt16Bool == nil {
					r.EncodeNil()
				} else {
					yym1909 := z.EncBinary()
					_ = yym1909
					if false {
					} else {
						z.F.EncMapInt16BoolV(x.FMapInt16Bool, e)
					}
				}
			}
			var yyn1910 bool
			if x.FptrMapInt16Bool == nil {
				yyn1910 = true
				goto LABEL1910
			}
		LABEL1910:
			if yyr2 || yy2arr2 {
				if yyn1910 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt16Bool == nil {
						r.EncodeNil()
					} else {
						yy1911 := *x.FptrMapInt16Bool
						yym1912 := z.EncBinary()
						_ = yym1912
						if false {
						} else {
							z.F.EncMapInt16BoolV(yy1911, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt16Bool"))
				r.WriteMapElemValue()
				if yyn1910 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt16Bool == nil {
						r.EncodeNil()
					} else {
						yy1913 := *x.FptrMapInt16Bool
						yym1914 := z.EncBinary()
						_ = yym1914
						if false {
						} else {
							z.F.EncMapInt16BoolV(yy1913, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Intf == nil {
					r.EncodeNil()
				} else {
					yym1916 := z.EncBinary()
					_ = yym1916
					if false {
					} else {
						z.F.EncMapInt32IntfV(x.FMapInt32Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Intf"))
				r.WriteMapElemValue()
				if x.FMapInt32Intf == nil {
					r.EncodeNil()
				} else {
					yym1917 := z.EncBinary()
					_ = yym1917
					if false {
					} else {
						z.F.EncMapInt32IntfV(x.FMapInt32Intf, e)
					}
				}
			}
			var yyn1918 bool
			if x.FptrMapInt32Intf == nil {
				yyn1918 = true
				goto LABEL1918
			}
		LABEL1918:
			if yyr2 || yy2arr2 {
				if yyn1918 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Intf == nil {
						r.EncodeNil()
					} else {
						yy1919 := *x.FptrMapInt32Intf
						yym1920 := z.EncBinary()
						_ = yym1920
						if false {
						} else {
							z.F.EncMapInt32IntfV(yy1919, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Intf"))
				r.WriteMapElemValue()
				if yyn1918 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Intf == nil {
						r.EncodeNil()
					} else {
						yy1921 := *x.FptrMapInt32Intf
						yym1922 := z.EncBinary()
						_ = yym1922
						if false {
						} else {
							z.F.EncMapInt32IntfV(yy1921, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32String == nil {
					r.EncodeNil()
				} else {
					yym1924 := z.EncBinary()
					_ = yym1924
					if false {
					} else {
						z.F.EncMapInt32StringV(x.FMapInt32String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32String"))
				r.WriteMapElemValue()
				if x.FMapInt32String == nil {
					r.EncodeNil()
				} else {
					yym1925 := z.EncBinary()
					_ = yym1925
					if false {
					} else {
						z.F.EncMapInt32StringV(x.FMapInt32String, e)
					}
				}
			}
			var yyn1926 bool
			if x.FptrMapInt32String == nil {
				yyn1926 = true
				goto LABEL1926
			}
		LABEL1926:
			if yyr2 || yy2arr2 {
				if yyn1926 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32String == nil {
						r.EncodeNil()
					} else {
						yy1927 := *x.FptrMapInt32String
						yym1928 := z.EncBinary()
						_ = yym1928
						if false {
						} else {
							z.F.EncMapInt32StringV(yy1927, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32String"))
				r.WriteMapElemValue()
				if yyn1926 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32String == nil {
						r.EncodeNil()
					} else {
						yy1929 := *x.FptrMapInt32String
						yym1930 := z.EncBinary()
						_ = yym1930
						if false {
						} else {
							z.F.EncMapInt32StringV(yy1929, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Uint == nil {
					r.EncodeNil()
				} else {
					yym1932 := z.EncBinary()
					_ = yym1932
					if false {
					} else {
						z.F.EncMapInt32UintV(x.FMapInt32Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Uint"))
				r.WriteMapElemValue()
				if x.FMapInt32Uint == nil {
					r.EncodeNil()
				} else {
					yym1933 := z.EncBinary()
					_ = yym1933
					if false {
					} else {
						z.F.EncMapInt32UintV(x.FMapInt32Uint, e)
					}
				}
			}
			var yyn1934 bool
			if x.FptrMapInt32Uint == nil {
				yyn1934 = true
				goto LABEL1934
			}
		LABEL1934:
			if yyr2 || yy2arr2 {
				if yyn1934 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Uint == nil {
						r.EncodeNil()
					} else {
						yy1935 := *x.FptrMapInt32Uint
						yym1936 := z.EncBinary()
						_ = yym1936
						if false {
						} else {
							z.F.EncMapInt32UintV(yy1935, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Uint"))
				r.WriteMapElemValue()
				if yyn1934 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Uint == nil {
						r.EncodeNil()
					} else {
						yy1937 := *x.FptrMapInt32Uint
						yym1938 := z.EncBinary()
						_ = yym1938
						if false {
						} else {
							z.F.EncMapInt32UintV(yy1937, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1940 := z.EncBinary()
					_ = yym1940
					if false {
					} else {
						z.F.EncMapInt32Uint8V(x.FMapInt32Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Uint8"))
				r.WriteMapElemValue()
				if x.FMapInt32Uint8 == nil {
					r.EncodeNil()
				} else {
					yym1941 := z.EncBinary()
					_ = yym1941
					if false {
					} else {
						z.F.EncMapInt32Uint8V(x.FMapInt32Uint8, e)
					}
				}
			}
			var yyn1942 bool
			if x.FptrMapInt32Uint8 == nil {
				yyn1942 = true
				goto LABEL1942
			}
		LABEL1942:
			if yyr2 || yy2arr2 {
				if yyn1942 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1943 := *x.FptrMapInt32Uint8
						yym1944 := z.EncBinary()
						_ = yym1944
						if false {
						} else {
							z.F.EncMapInt32Uint8V(yy1943, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Uint8"))
				r.WriteMapElemValue()
				if yyn1942 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Uint8 == nil {
						r.EncodeNil()
					} else {
						yy1945 := *x.FptrMapInt32Uint8
						yym1946 := z.EncBinary()
						_ = yym1946
						if false {
						} else {
							z.F.EncMapInt32Uint8V(yy1945, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1948 := z.EncBinary()
					_ = yym1948
					if false {
					} else {
						z.F.EncMapInt32Uint16V(x.FMapInt32Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Uint16"))
				r.WriteMapElemValue()
				if x.FMapInt32Uint16 == nil {
					r.EncodeNil()
				} else {
					yym1949 := z.EncBinary()
					_ = yym1949
					if false {
					} else {
						z.F.EncMapInt32Uint16V(x.FMapInt32Uint16, e)
					}
				}
			}
			var yyn1950 bool
			if x.FptrMapInt32Uint16 == nil {
				yyn1950 = true
				goto LABEL1950
			}
		LABEL1950:
			if yyr2 || yy2arr2 {
				if yyn1950 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1951 := *x.FptrMapInt32Uint16
						yym1952 := z.EncBinary()
						_ = yym1952
						if false {
						} else {
							z.F.EncMapInt32Uint16V(yy1951, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Uint16"))
				r.WriteMapElemValue()
				if yyn1950 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Uint16 == nil {
						r.EncodeNil()
					} else {
						yy1953 := *x.FptrMapInt32Uint16
						yym1954 := z.EncBinary()
						_ = yym1954
						if false {
						} else {
							z.F.EncMapInt32Uint16V(yy1953, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1956 := z.EncBinary()
					_ = yym1956
					if false {
					} else {
						z.F.EncMapInt32Uint32V(x.FMapInt32Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Uint32"))
				r.WriteMapElemValue()
				if x.FMapInt32Uint32 == nil {
					r.EncodeNil()
				} else {
					yym1957 := z.EncBinary()
					_ = yym1957
					if false {
					} else {
						z.F.EncMapInt32Uint32V(x.FMapInt32Uint32, e)
					}
				}
			}
			var yyn1958 bool
			if x.FptrMapInt32Uint32 == nil {
				yyn1958 = true
				goto LABEL1958
			}
		LABEL1958:
			if yyr2 || yy2arr2 {
				if yyn1958 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1959 := *x.FptrMapInt32Uint32
						yym1960 := z.EncBinary()
						_ = yym1960
						if false {
						} else {
							z.F.EncMapInt32Uint32V(yy1959, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Uint32"))
				r.WriteMapElemValue()
				if yyn1958 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Uint32 == nil {
						r.EncodeNil()
					} else {
						yy1961 := *x.FptrMapInt32Uint32
						yym1962 := z.EncBinary()
						_ = yym1962
						if false {
						} else {
							z.F.EncMapInt32Uint32V(yy1961, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1964 := z.EncBinary()
					_ = yym1964
					if false {
					} else {
						z.F.EncMapInt32Uint64V(x.FMapInt32Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Uint64"))
				r.WriteMapElemValue()
				if x.FMapInt32Uint64 == nil {
					r.EncodeNil()
				} else {
					yym1965 := z.EncBinary()
					_ = yym1965
					if false {
					} else {
						z.F.EncMapInt32Uint64V(x.FMapInt32Uint64, e)
					}
				}
			}
			var yyn1966 bool
			if x.FptrMapInt32Uint64 == nil {
				yyn1966 = true
				goto LABEL1966
			}
		LABEL1966:
			if yyr2 || yy2arr2 {
				if yyn1966 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1967 := *x.FptrMapInt32Uint64
						yym1968 := z.EncBinary()
						_ = yym1968
						if false {
						} else {
							z.F.EncMapInt32Uint64V(yy1967, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Uint64"))
				r.WriteMapElemValue()
				if yyn1966 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Uint64 == nil {
						r.EncodeNil()
					} else {
						yy1969 := *x.FptrMapInt32Uint64
						yym1970 := z.EncBinary()
						_ = yym1970
						if false {
						} else {
							z.F.EncMapInt32Uint64V(yy1969, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1972 := z.EncBinary()
					_ = yym1972
					if false {
					} else {
						z.F.EncMapInt32UintptrV(x.FMapInt32Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Uintptr"))
				r.WriteMapElemValue()
				if x.FMapInt32Uintptr == nil {
					r.EncodeNil()
				} else {
					yym1973 := z.EncBinary()
					_ = yym1973
					if false {
					} else {
						z.F.EncMapInt32UintptrV(x.FMapInt32Uintptr, e)
					}
				}
			}
			var yyn1974 bool
			if x.FptrMapInt32Uintptr == nil {
				yyn1974 = true
				goto LABEL1974
			}
		LABEL1974:
			if yyr2 || yy2arr2 {
				if yyn1974 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1975 := *x.FptrMapInt32Uintptr
						yym1976 := z.EncBinary()
						_ = yym1976
						if false {
						} else {
							z.F.EncMapInt32UintptrV(yy1975, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Uintptr"))
				r.WriteMapElemValue()
				if yyn1974 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Uintptr == nil {
						r.EncodeNil()
					} else {
						yy1977 := *x.FptrMapInt32Uintptr
						yym1978 := z.EncBinary()
						_ = yym1978
						if false {
						} else {
							z.F.EncMapInt32UintptrV(yy1977, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Int == nil {
					r.EncodeNil()
				} else {
					yym1980 := z.EncBinary()
					_ = yym1980
					if false {
					} else {
						z.F.EncMapInt32IntV(x.FMapInt32Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Int"))
				r.WriteMapElemValue()
				if x.FMapInt32Int == nil {
					r.EncodeNil()
				} else {
					yym1981 := z.EncBinary()
					_ = yym1981
					if false {
					} else {
						z.F.EncMapInt32IntV(x.FMapInt32Int, e)
					}
				}
			}
			var yyn1982 bool
			if x.FptrMapInt32Int == nil {
				yyn1982 = true
				goto LABEL1982
			}
		LABEL1982:
			if yyr2 || yy2arr2 {
				if yyn1982 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Int == nil {
						r.EncodeNil()
					} else {
						yy1983 := *x.FptrMapInt32Int
						yym1984 := z.EncBinary()
						_ = yym1984
						if false {
						} else {
							z.F.EncMapInt32IntV(yy1983, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Int"))
				r.WriteMapElemValue()
				if yyn1982 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Int == nil {
						r.EncodeNil()
					} else {
						yy1985 := *x.FptrMapInt32Int
						yym1986 := z.EncBinary()
						_ = yym1986
						if false {
						} else {
							z.F.EncMapInt32IntV(yy1985, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Int8 == nil {
					r.EncodeNil()
				} else {
					yym1988 := z.EncBinary()
					_ = yym1988
					if false {
					} else {
						z.F.EncMapInt32Int8V(x.FMapInt32Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Int8"))
				r.WriteMapElemValue()
				if x.FMapInt32Int8 == nil {
					r.EncodeNil()
				} else {
					yym1989 := z.EncBinary()
					_ = yym1989
					if false {
					} else {
						z.F.EncMapInt32Int8V(x.FMapInt32Int8, e)
					}
				}
			}
			var yyn1990 bool
			if x.FptrMapInt32Int8 == nil {
				yyn1990 = true
				goto LABEL1990
			}
		LABEL1990:
			if yyr2 || yy2arr2 {
				if yyn1990 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Int8 == nil {
						r.EncodeNil()
					} else {
						yy1991 := *x.FptrMapInt32Int8
						yym1992 := z.EncBinary()
						_ = yym1992
						if false {
						} else {
							z.F.EncMapInt32Int8V(yy1991, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Int8"))
				r.WriteMapElemValue()
				if yyn1990 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Int8 == nil {
						r.EncodeNil()
					} else {
						yy1993 := *x.FptrMapInt32Int8
						yym1994 := z.EncBinary()
						_ = yym1994
						if false {
						} else {
							z.F.EncMapInt32Int8V(yy1993, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Int16 == nil {
					r.EncodeNil()
				} else {
					yym1996 := z.EncBinary()
					_ = yym1996
					if false {
					} else {
						z.F.EncMapInt32Int16V(x.FMapInt32Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Int16"))
				r.WriteMapElemValue()
				if x.FMapInt32Int16 == nil {
					r.EncodeNil()
				} else {
					yym1997 := z.EncBinary()
					_ = yym1997
					if false {
					} else {
						z.F.EncMapInt32Int16V(x.FMapInt32Int16, e)
					}
				}
			}
			var yyn1998 bool
			if x.FptrMapInt32Int16 == nil {
				yyn1998 = true
				goto LABEL1998
			}
		LABEL1998:
			if yyr2 || yy2arr2 {
				if yyn1998 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Int16 == nil {
						r.EncodeNil()
					} else {
						yy1999 := *x.FptrMapInt32Int16
						yym2000 := z.EncBinary()
						_ = yym2000
						if false {
						} else {
							z.F.EncMapInt32Int16V(yy1999, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Int16"))
				r.WriteMapElemValue()
				if yyn1998 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Int16 == nil {
						r.EncodeNil()
					} else {
						yy2001 := *x.FptrMapInt32Int16
						yym2002 := z.EncBinary()
						_ = yym2002
						if false {
						} else {
							z.F.EncMapInt32Int16V(yy2001, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Int32 == nil {
					r.EncodeNil()
				} else {
					yym2004 := z.EncBinary()
					_ = yym2004
					if false {
					} else {
						z.F.EncMapInt32Int32V(x.FMapInt32Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Int32"))
				r.WriteMapElemValue()
				if x.FMapInt32Int32 == nil {
					r.EncodeNil()
				} else {
					yym2005 := z.EncBinary()
					_ = yym2005
					if false {
					} else {
						z.F.EncMapInt32Int32V(x.FMapInt32Int32, e)
					}
				}
			}
			var yyn2006 bool
			if x.FptrMapInt32Int32 == nil {
				yyn2006 = true
				goto LABEL2006
			}
		LABEL2006:
			if yyr2 || yy2arr2 {
				if yyn2006 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Int32 == nil {
						r.EncodeNil()
					} else {
						yy2007 := *x.FptrMapInt32Int32
						yym2008 := z.EncBinary()
						_ = yym2008
						if false {
						} else {
							z.F.EncMapInt32Int32V(yy2007, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Int32"))
				r.WriteMapElemValue()
				if yyn2006 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Int32 == nil {
						r.EncodeNil()
					} else {
						yy2009 := *x.FptrMapInt32Int32
						yym2010 := z.EncBinary()
						_ = yym2010
						if false {
						} else {
							z.F.EncMapInt32Int32V(yy2009, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Int64 == nil {
					r.EncodeNil()
				} else {
					yym2012 := z.EncBinary()
					_ = yym2012
					if false {
					} else {
						z.F.EncMapInt32Int64V(x.FMapInt32Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Int64"))
				r.WriteMapElemValue()
				if x.FMapInt32Int64 == nil {
					r.EncodeNil()
				} else {
					yym2013 := z.EncBinary()
					_ = yym2013
					if false {
					} else {
						z.F.EncMapInt32Int64V(x.FMapInt32Int64, e)
					}
				}
			}
			var yyn2014 bool
			if x.FptrMapInt32Int64 == nil {
				yyn2014 = true
				goto LABEL2014
			}
		LABEL2014:
			if yyr2 || yy2arr2 {
				if yyn2014 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Int64 == nil {
						r.EncodeNil()
					} else {
						yy2015 := *x.FptrMapInt32Int64
						yym2016 := z.EncBinary()
						_ = yym2016
						if false {
						} else {
							z.F.EncMapInt32Int64V(yy2015, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Int64"))
				r.WriteMapElemValue()
				if yyn2014 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Int64 == nil {
						r.EncodeNil()
					} else {
						yy2017 := *x.FptrMapInt32Int64
						yym2018 := z.EncBinary()
						_ = yym2018
						if false {
						} else {
							z.F.EncMapInt32Int64V(yy2017, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Float32 == nil {
					r.EncodeNil()
				} else {
					yym2020 := z.EncBinary()
					_ = yym2020
					if false {
					} else {
						z.F.EncMapInt32Float32V(x.FMapInt32Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Float32"))
				r.WriteMapElemValue()
				if x.FMapInt32Float32 == nil {
					r.EncodeNil()
				} else {
					yym2021 := z.EncBinary()
					_ = yym2021
					if false {
					} else {
						z.F.EncMapInt32Float32V(x.FMapInt32Float32, e)
					}
				}
			}
			var yyn2022 bool
			if x.FptrMapInt32Float32 == nil {
				yyn2022 = true
				goto LABEL2022
			}
		LABEL2022:
			if yyr2 || yy2arr2 {
				if yyn2022 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Float32 == nil {
						r.EncodeNil()
					} else {
						yy2023 := *x.FptrMapInt32Float32
						yym2024 := z.EncBinary()
						_ = yym2024
						if false {
						} else {
							z.F.EncMapInt32Float32V(yy2023, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Float32"))
				r.WriteMapElemValue()
				if yyn2022 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Float32 == nil {
						r.EncodeNil()
					} else {
						yy2025 := *x.FptrMapInt32Float32
						yym2026 := z.EncBinary()
						_ = yym2026
						if false {
						} else {
							z.F.EncMapInt32Float32V(yy2025, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Float64 == nil {
					r.EncodeNil()
				} else {
					yym2028 := z.EncBinary()
					_ = yym2028
					if false {
					} else {
						z.F.EncMapInt32Float64V(x.FMapInt32Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Float64"))
				r.WriteMapElemValue()
				if x.FMapInt32Float64 == nil {
					r.EncodeNil()
				} else {
					yym2029 := z.EncBinary()
					_ = yym2029
					if false {
					} else {
						z.F.EncMapInt32Float64V(x.FMapInt32Float64, e)
					}
				}
			}
			var yyn2030 bool
			if x.FptrMapInt32Float64 == nil {
				yyn2030 = true
				goto LABEL2030
			}
		LABEL2030:
			if yyr2 || yy2arr2 {
				if yyn2030 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Float64 == nil {
						r.EncodeNil()
					} else {
						yy2031 := *x.FptrMapInt32Float64
						yym2032 := z.EncBinary()
						_ = yym2032
						if false {
						} else {
							z.F.EncMapInt32Float64V(yy2031, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Float64"))
				r.WriteMapElemValue()
				if yyn2030 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Float64 == nil {
						r.EncodeNil()
					} else {
						yy2033 := *x.FptrMapInt32Float64
						yym2034 := z.EncBinary()
						_ = yym2034
						if false {
						} else {
							z.F.EncMapInt32Float64V(yy2033, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt32Bool == nil {
					r.EncodeNil()
				} else {
					yym2036 := z.EncBinary()
					_ = yym2036
					if false {
					} else {
						z.F.EncMapInt32BoolV(x.FMapInt32Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt32Bool"))
				r.WriteMapElemValue()
				if x.FMapInt32Bool == nil {
					r.EncodeNil()
				} else {
					yym2037 := z.EncBinary()
					_ = yym2037
					if false {
					} else {
						z.F.EncMapInt32BoolV(x.FMapInt32Bool, e)
					}
				}
			}
			var yyn2038 bool
			if x.FptrMapInt32Bool == nil {
				yyn2038 = true
				goto LABEL2038
			}
		LABEL2038:
			if yyr2 || yy2arr2 {
				if yyn2038 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt32Bool == nil {
						r.EncodeNil()
					} else {
						yy2039 := *x.FptrMapInt32Bool
						yym2040 := z.EncBinary()
						_ = yym2040
						if false {
						} else {
							z.F.EncMapInt32BoolV(yy2039, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt32Bool"))
				r.WriteMapElemValue()
				if yyn2038 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt32Bool == nil {
						r.EncodeNil()
					} else {
						yy2041 := *x.FptrMapInt32Bool
						yym2042 := z.EncBinary()
						_ = yym2042
						if false {
						} else {
							z.F.EncMapInt32BoolV(yy2041, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Intf == nil {
					r.EncodeNil()
				} else {
					yym2044 := z.EncBinary()
					_ = yym2044
					if false {
					} else {
						z.F.EncMapInt64IntfV(x.FMapInt64Intf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Intf"))
				r.WriteMapElemValue()
				if x.FMapInt64Intf == nil {
					r.EncodeNil()
				} else {
					yym2045 := z.EncBinary()
					_ = yym2045
					if false {
					} else {
						z.F.EncMapInt64IntfV(x.FMapInt64Intf, e)
					}
				}
			}
			var yyn2046 bool
			if x.FptrMapInt64Intf == nil {
				yyn2046 = true
				goto LABEL2046
			}
		LABEL2046:
			if yyr2 || yy2arr2 {
				if yyn2046 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Intf == nil {
						r.EncodeNil()
					} else {
						yy2047 := *x.FptrMapInt64Intf
						yym2048 := z.EncBinary()
						_ = yym2048
						if false {
						} else {
							z.F.EncMapInt64IntfV(yy2047, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Intf"))
				r.WriteMapElemValue()
				if yyn2046 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Intf == nil {
						r.EncodeNil()
					} else {
						yy2049 := *x.FptrMapInt64Intf
						yym2050 := z.EncBinary()
						_ = yym2050
						if false {
						} else {
							z.F.EncMapInt64IntfV(yy2049, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64String == nil {
					r.EncodeNil()
				} else {
					yym2052 := z.EncBinary()
					_ = yym2052
					if false {
					} else {
						z.F.EncMapInt64StringV(x.FMapInt64String, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64String"))
				r.WriteMapElemValue()
				if x.FMapInt64String == nil {
					r.EncodeNil()
				} else {
					yym2053 := z.EncBinary()
					_ = yym2053
					if false {
					} else {
						z.F.EncMapInt64StringV(x.FMapInt64String, e)
					}
				}
			}
			var yyn2054 bool
			if x.FptrMapInt64String == nil {
				yyn2054 = true
				goto LABEL2054
			}
		LABEL2054:
			if yyr2 || yy2arr2 {
				if yyn2054 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64String == nil {
						r.EncodeNil()
					} else {
						yy2055 := *x.FptrMapInt64String
						yym2056 := z.EncBinary()
						_ = yym2056
						if false {
						} else {
							z.F.EncMapInt64StringV(yy2055, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64String"))
				r.WriteMapElemValue()
				if yyn2054 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64String == nil {
						r.EncodeNil()
					} else {
						yy2057 := *x.FptrMapInt64String
						yym2058 := z.EncBinary()
						_ = yym2058
						if false {
						} else {
							z.F.EncMapInt64StringV(yy2057, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Uint == nil {
					r.EncodeNil()
				} else {
					yym2060 := z.EncBinary()
					_ = yym2060
					if false {
					} else {
						z.F.EncMapInt64UintV(x.FMapInt64Uint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Uint"))
				r.WriteMapElemValue()
				if x.FMapInt64Uint == nil {
					r.EncodeNil()
				} else {
					yym2061 := z.EncBinary()
					_ = yym2061
					if false {
					} else {
						z.F.EncMapInt64UintV(x.FMapInt64Uint, e)
					}
				}
			}
			var yyn2062 bool
			if x.FptrMapInt64Uint == nil {
				yyn2062 = true
				goto LABEL2062
			}
		LABEL2062:
			if yyr2 || yy2arr2 {
				if yyn2062 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Uint == nil {
						r.EncodeNil()
					} else {
						yy2063 := *x.FptrMapInt64Uint
						yym2064 := z.EncBinary()
						_ = yym2064
						if false {
						} else {
							z.F.EncMapInt64UintV(yy2063, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Uint"))
				r.WriteMapElemValue()
				if yyn2062 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Uint == nil {
						r.EncodeNil()
					} else {
						yy2065 := *x.FptrMapInt64Uint
						yym2066 := z.EncBinary()
						_ = yym2066
						if false {
						} else {
							z.F.EncMapInt64UintV(yy2065, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Uint8 == nil {
					r.EncodeNil()
				} else {
					yym2068 := z.EncBinary()
					_ = yym2068
					if false {
					} else {
						z.F.EncMapInt64Uint8V(x.FMapInt64Uint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Uint8"))
				r.WriteMapElemValue()
				if x.FMapInt64Uint8 == nil {
					r.EncodeNil()
				} else {
					yym2069 := z.EncBinary()
					_ = yym2069
					if false {
					} else {
						z.F.EncMapInt64Uint8V(x.FMapInt64Uint8, e)
					}
				}
			}
			var yyn2070 bool
			if x.FptrMapInt64Uint8 == nil {
				yyn2070 = true
				goto LABEL2070
			}
		LABEL2070:
			if yyr2 || yy2arr2 {
				if yyn2070 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Uint8 == nil {
						r.EncodeNil()
					} else {
						yy2071 := *x.FptrMapInt64Uint8
						yym2072 := z.EncBinary()
						_ = yym2072
						if false {
						} else {
							z.F.EncMapInt64Uint8V(yy2071, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Uint8"))
				r.WriteMapElemValue()
				if yyn2070 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Uint8 == nil {
						r.EncodeNil()
					} else {
						yy2073 := *x.FptrMapInt64Uint8
						yym2074 := z.EncBinary()
						_ = yym2074
						if false {
						} else {
							z.F.EncMapInt64Uint8V(yy2073, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Uint16 == nil {
					r.EncodeNil()
				} else {
					yym2076 := z.EncBinary()
					_ = yym2076
					if false {
					} else {
						z.F.EncMapInt64Uint16V(x.FMapInt64Uint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Uint16"))
				r.WriteMapElemValue()
				if x.FMapInt64Uint16 == nil {
					r.EncodeNil()
				} else {
					yym2077 := z.EncBinary()
					_ = yym2077
					if false {
					} else {
						z.F.EncMapInt64Uint16V(x.FMapInt64Uint16, e)
					}
				}
			}
			var yyn2078 bool
			if x.FptrMapInt64Uint16 == nil {
				yyn2078 = true
				goto LABEL2078
			}
		LABEL2078:
			if yyr2 || yy2arr2 {
				if yyn2078 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Uint16 == nil {
						r.EncodeNil()
					} else {
						yy2079 := *x.FptrMapInt64Uint16
						yym2080 := z.EncBinary()
						_ = yym2080
						if false {
						} else {
							z.F.EncMapInt64Uint16V(yy2079, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Uint16"))
				r.WriteMapElemValue()
				if yyn2078 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Uint16 == nil {
						r.EncodeNil()
					} else {
						yy2081 := *x.FptrMapInt64Uint16
						yym2082 := z.EncBinary()
						_ = yym2082
						if false {
						} else {
							z.F.EncMapInt64Uint16V(yy2081, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Uint32 == nil {
					r.EncodeNil()
				} else {
					yym2084 := z.EncBinary()
					_ = yym2084
					if false {
					} else {
						z.F.EncMapInt64Uint32V(x.FMapInt64Uint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Uint32"))
				r.WriteMapElemValue()
				if x.FMapInt64Uint32 == nil {
					r.EncodeNil()
				} else {
					yym2085 := z.EncBinary()
					_ = yym2085
					if false {
					} else {
						z.F.EncMapInt64Uint32V(x.FMapInt64Uint32, e)
					}
				}
			}
			var yyn2086 bool
			if x.FptrMapInt64Uint32 == nil {
				yyn2086 = true
				goto LABEL2086
			}
		LABEL2086:
			if yyr2 || yy2arr2 {
				if yyn2086 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Uint32 == nil {
						r.EncodeNil()
					} else {
						yy2087 := *x.FptrMapInt64Uint32
						yym2088 := z.EncBinary()
						_ = yym2088
						if false {
						} else {
							z.F.EncMapInt64Uint32V(yy2087, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Uint32"))
				r.WriteMapElemValue()
				if yyn2086 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Uint32 == nil {
						r.EncodeNil()
					} else {
						yy2089 := *x.FptrMapInt64Uint32
						yym2090 := z.EncBinary()
						_ = yym2090
						if false {
						} else {
							z.F.EncMapInt64Uint32V(yy2089, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Uint64 == nil {
					r.EncodeNil()
				} else {
					yym2092 := z.EncBinary()
					_ = yym2092
					if false {
					} else {
						z.F.EncMapInt64Uint64V(x.FMapInt64Uint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Uint64"))
				r.WriteMapElemValue()
				if x.FMapInt64Uint64 == nil {
					r.EncodeNil()
				} else {
					yym2093 := z.EncBinary()
					_ = yym2093
					if false {
					} else {
						z.F.EncMapInt64Uint64V(x.FMapInt64Uint64, e)
					}
				}
			}
			var yyn2094 bool
			if x.FptrMapInt64Uint64 == nil {
				yyn2094 = true
				goto LABEL2094
			}
		LABEL2094:
			if yyr2 || yy2arr2 {
				if yyn2094 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Uint64 == nil {
						r.EncodeNil()
					} else {
						yy2095 := *x.FptrMapInt64Uint64
						yym2096 := z.EncBinary()
						_ = yym2096
						if false {
						} else {
							z.F.EncMapInt64Uint64V(yy2095, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Uint64"))
				r.WriteMapElemValue()
				if yyn2094 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Uint64 == nil {
						r.EncodeNil()
					} else {
						yy2097 := *x.FptrMapInt64Uint64
						yym2098 := z.EncBinary()
						_ = yym2098
						if false {
						} else {
							z.F.EncMapInt64Uint64V(yy2097, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Uintptr == nil {
					r.EncodeNil()
				} else {
					yym2100 := z.EncBinary()
					_ = yym2100
					if false {
					} else {
						z.F.EncMapInt64UintptrV(x.FMapInt64Uintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Uintptr"))
				r.WriteMapElemValue()
				if x.FMapInt64Uintptr == nil {
					r.EncodeNil()
				} else {
					yym2101 := z.EncBinary()
					_ = yym2101
					if false {
					} else {
						z.F.EncMapInt64UintptrV(x.FMapInt64Uintptr, e)
					}
				}
			}
			var yyn2102 bool
			if x.FptrMapInt64Uintptr == nil {
				yyn2102 = true
				goto LABEL2102
			}
		LABEL2102:
			if yyr2 || yy2arr2 {
				if yyn2102 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Uintptr == nil {
						r.EncodeNil()
					} else {
						yy2103 := *x.FptrMapInt64Uintptr
						yym2104 := z.EncBinary()
						_ = yym2104
						if false {
						} else {
							z.F.EncMapInt64UintptrV(yy2103, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Uintptr"))
				r.WriteMapElemValue()
				if yyn2102 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Uintptr == nil {
						r.EncodeNil()
					} else {
						yy2105 := *x.FptrMapInt64Uintptr
						yym2106 := z.EncBinary()
						_ = yym2106
						if false {
						} else {
							z.F.EncMapInt64UintptrV(yy2105, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Int == nil {
					r.EncodeNil()
				} else {
					yym2108 := z.EncBinary()
					_ = yym2108
					if false {
					} else {
						z.F.EncMapInt64IntV(x.FMapInt64Int, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Int"))
				r.WriteMapElemValue()
				if x.FMapInt64Int == nil {
					r.EncodeNil()
				} else {
					yym2109 := z.EncBinary()
					_ = yym2109
					if false {
					} else {
						z.F.EncMapInt64IntV(x.FMapInt64Int, e)
					}
				}
			}
			var yyn2110 bool
			if x.FptrMapInt64Int == nil {
				yyn2110 = true
				goto LABEL2110
			}
		LABEL2110:
			if yyr2 || yy2arr2 {
				if yyn2110 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Int == nil {
						r.EncodeNil()
					} else {
						yy2111 := *x.FptrMapInt64Int
						yym2112 := z.EncBinary()
						_ = yym2112
						if false {
						} else {
							z.F.EncMapInt64IntV(yy2111, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Int"))
				r.WriteMapElemValue()
				if yyn2110 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Int == nil {
						r.EncodeNil()
					} else {
						yy2113 := *x.FptrMapInt64Int
						yym2114 := z.EncBinary()
						_ = yym2114
						if false {
						} else {
							z.F.EncMapInt64IntV(yy2113, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Int8 == nil {
					r.EncodeNil()
				} else {
					yym2116 := z.EncBinary()
					_ = yym2116
					if false {
					} else {
						z.F.EncMapInt64Int8V(x.FMapInt64Int8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Int8"))
				r.WriteMapElemValue()
				if x.FMapInt64Int8 == nil {
					r.EncodeNil()
				} else {
					yym2117 := z.EncBinary()
					_ = yym2117
					if false {
					} else {
						z.F.EncMapInt64Int8V(x.FMapInt64Int8, e)
					}
				}
			}
			var yyn2118 bool
			if x.FptrMapInt64Int8 == nil {
				yyn2118 = true
				goto LABEL2118
			}
		LABEL2118:
			if yyr2 || yy2arr2 {
				if yyn2118 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Int8 == nil {
						r.EncodeNil()
					} else {
						yy2119 := *x.FptrMapInt64Int8
						yym2120 := z.EncBinary()
						_ = yym2120
						if false {
						} else {
							z.F.EncMapInt64Int8V(yy2119, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Int8"))
				r.WriteMapElemValue()
				if yyn2118 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Int8 == nil {
						r.EncodeNil()
					} else {
						yy2121 := *x.FptrMapInt64Int8
						yym2122 := z.EncBinary()
						_ = yym2122
						if false {
						} else {
							z.F.EncMapInt64Int8V(yy2121, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Int16 == nil {
					r.EncodeNil()
				} else {
					yym2124 := z.EncBinary()
					_ = yym2124
					if false {
					} else {
						z.F.EncMapInt64Int16V(x.FMapInt64Int16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Int16"))
				r.WriteMapElemValue()
				if x.FMapInt64Int16 == nil {
					r.EncodeNil()
				} else {
					yym2125 := z.EncBinary()
					_ = yym2125
					if false {
					} else {
						z.F.EncMapInt64Int16V(x.FMapInt64Int16, e)
					}
				}
			}
			var yyn2126 bool
			if x.FptrMapInt64Int16 == nil {
				yyn2126 = true
				goto LABEL2126
			}
		LABEL2126:
			if yyr2 || yy2arr2 {
				if yyn2126 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Int16 == nil {
						r.EncodeNil()
					} else {
						yy2127 := *x.FptrMapInt64Int16
						yym2128 := z.EncBinary()
						_ = yym2128
						if false {
						} else {
							z.F.EncMapInt64Int16V(yy2127, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Int16"))
				r.WriteMapElemValue()
				if yyn2126 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Int16 == nil {
						r.EncodeNil()
					} else {
						yy2129 := *x.FptrMapInt64Int16
						yym2130 := z.EncBinary()
						_ = yym2130
						if false {
						} else {
							z.F.EncMapInt64Int16V(yy2129, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Int32 == nil {
					r.EncodeNil()
				} else {
					yym2132 := z.EncBinary()
					_ = yym2132
					if false {
					} else {
						z.F.EncMapInt64Int32V(x.FMapInt64Int32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Int32"))
				r.WriteMapElemValue()
				if x.FMapInt64Int32 == nil {
					r.EncodeNil()
				} else {
					yym2133 := z.EncBinary()
					_ = yym2133
					if false {
					} else {
						z.F.EncMapInt64Int32V(x.FMapInt64Int32, e)
					}
				}
			}
			var yyn2134 bool
			if x.FptrMapInt64Int32 == nil {
				yyn2134 = true
				goto LABEL2134
			}
		LABEL2134:
			if yyr2 || yy2arr2 {
				if yyn2134 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Int32 == nil {
						r.EncodeNil()
					} else {
						yy2135 := *x.FptrMapInt64Int32
						yym2136 := z.EncBinary()
						_ = yym2136
						if false {
						} else {
							z.F.EncMapInt64Int32V(yy2135, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Int32"))
				r.WriteMapElemValue()
				if yyn2134 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Int32 == nil {
						r.EncodeNil()
					} else {
						yy2137 := *x.FptrMapInt64Int32
						yym2138 := z.EncBinary()
						_ = yym2138
						if false {
						} else {
							z.F.EncMapInt64Int32V(yy2137, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Int64 == nil {
					r.EncodeNil()
				} else {
					yym2140 := z.EncBinary()
					_ = yym2140
					if false {
					} else {
						z.F.EncMapInt64Int64V(x.FMapInt64Int64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Int64"))
				r.WriteMapElemValue()
				if x.FMapInt64Int64 == nil {
					r.EncodeNil()
				} else {
					yym2141 := z.EncBinary()
					_ = yym2141
					if false {
					} else {
						z.F.EncMapInt64Int64V(x.FMapInt64Int64, e)
					}
				}
			}
			var yyn2142 bool
			if x.FptrMapInt64Int64 == nil {
				yyn2142 = true
				goto LABEL2142
			}
		LABEL2142:
			if yyr2 || yy2arr2 {
				if yyn2142 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Int64 == nil {
						r.EncodeNil()
					} else {
						yy2143 := *x.FptrMapInt64Int64
						yym2144 := z.EncBinary()
						_ = yym2144
						if false {
						} else {
							z.F.EncMapInt64Int64V(yy2143, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Int64"))
				r.WriteMapElemValue()
				if yyn2142 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Int64 == nil {
						r.EncodeNil()
					} else {
						yy2145 := *x.FptrMapInt64Int64
						yym2146 := z.EncBinary()
						_ = yym2146
						if false {
						} else {
							z.F.EncMapInt64Int64V(yy2145, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Float32 == nil {
					r.EncodeNil()
				} else {
					yym2148 := z.EncBinary()
					_ = yym2148
					if false {
					} else {
						z.F.EncMapInt64Float32V(x.FMapInt64Float32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Float32"))
				r.WriteMapElemValue()
				if x.FMapInt64Float32 == nil {
					r.EncodeNil()
				} else {
					yym2149 := z.EncBinary()
					_ = yym2149
					if false {
					} else {
						z.F.EncMapInt64Float32V(x.FMapInt64Float32, e)
					}
				}
			}
			var yyn2150 bool
			if x.FptrMapInt64Float32 == nil {
				yyn2150 = true
				goto LABEL2150
			}
		LABEL2150:
			if yyr2 || yy2arr2 {
				if yyn2150 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Float32 == nil {
						r.EncodeNil()
					} else {
						yy2151 := *x.FptrMapInt64Float32
						yym2152 := z.EncBinary()
						_ = yym2152
						if false {
						} else {
							z.F.EncMapInt64Float32V(yy2151, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Float32"))
				r.WriteMapElemValue()
				if yyn2150 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Float32 == nil {
						r.EncodeNil()
					} else {
						yy2153 := *x.FptrMapInt64Float32
						yym2154 := z.EncBinary()
						_ = yym2154
						if false {
						} else {
							z.F.EncMapInt64Float32V(yy2153, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Float64 == nil {
					r.EncodeNil()
				} else {
					yym2156 := z.EncBinary()
					_ = yym2156
					if false {
					} else {
						z.F.EncMapInt64Float64V(x.FMapInt64Float64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Float64"))
				r.WriteMapElemValue()
				if x.FMapInt64Float64 == nil {
					r.EncodeNil()
				} else {
					yym2157 := z.EncBinary()
					_ = yym2157
					if false {
					} else {
						z.F.EncMapInt64Float64V(x.FMapInt64Float64, e)
					}
				}
			}
			var yyn2158 bool
			if x.FptrMapInt64Float64 == nil {
				yyn2158 = true
				goto LABEL2158
			}
		LABEL2158:
			if yyr2 || yy2arr2 {
				if yyn2158 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Float64 == nil {
						r.EncodeNil()
					} else {
						yy2159 := *x.FptrMapInt64Float64
						yym2160 := z.EncBinary()
						_ = yym2160
						if false {
						} else {
							z.F.EncMapInt64Float64V(yy2159, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Float64"))
				r.WriteMapElemValue()
				if yyn2158 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Float64 == nil {
						r.EncodeNil()
					} else {
						yy2161 := *x.FptrMapInt64Float64
						yym2162 := z.EncBinary()
						_ = yym2162
						if false {
						} else {
							z.F.EncMapInt64Float64V(yy2161, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapInt64Bool == nil {
					r.EncodeNil()
				} else {
					yym2164 := z.EncBinary()
					_ = yym2164
					if false {
					} else {
						z.F.EncMapInt64BoolV(x.FMapInt64Bool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapInt64Bool"))
				r.WriteMapElemValue()
				if x.FMapInt64Bool == nil {
					r.EncodeNil()
				} else {
					yym2165 := z.EncBinary()
					_ = yym2165
					if false {
					} else {
						z.F.EncMapInt64BoolV(x.FMapInt64Bool, e)
					}
				}
			}
			var yyn2166 bool
			if x.FptrMapInt64Bool == nil {
				yyn2166 = true
				goto LABEL2166
			}
		LABEL2166:
			if yyr2 || yy2arr2 {
				if yyn2166 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapInt64Bool == nil {
						r.EncodeNil()
					} else {
						yy2167 := *x.FptrMapInt64Bool
						yym2168 := z.EncBinary()
						_ = yym2168
						if false {
						} else {
							z.F.EncMapInt64BoolV(yy2167, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapInt64Bool"))
				r.WriteMapElemValue()
				if yyn2166 {
					r.EncodeNil()
				} else {
					if x.FptrMapInt64Bool == nil {
						r.EncodeNil()
					} else {
						yy2169 := *x.FptrMapInt64Bool
						yym2170 := z.EncBinary()
						_ = yym2170
						if false {
						} else {
							z.F.EncMapInt64BoolV(yy2169, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolIntf == nil {
					r.EncodeNil()
				} else {
					yym2172 := z.EncBinary()
					_ = yym2172
					if false {
					} else {
						z.F.EncMapBoolIntfV(x.FMapBoolIntf, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolIntf"))
				r.WriteMapElemValue()
				if x.FMapBoolIntf == nil {
					r.EncodeNil()
				} else {
					yym2173 := z.EncBinary()
					_ = yym2173
					if false {
					} else {
						z.F.EncMapBoolIntfV(x.FMapBoolIntf, e)
					}
				}
			}
			var yyn2174 bool
			if x.FptrMapBoolIntf == nil {
				yyn2174 = true
				goto LABEL2174
			}
		LABEL2174:
			if yyr2 || yy2arr2 {
				if yyn2174 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolIntf == nil {
						r.EncodeNil()
					} else {
						yy2175 := *x.FptrMapBoolIntf
						yym2176 := z.EncBinary()
						_ = yym2176
						if false {
						} else {
							z.F.EncMapBoolIntfV(yy2175, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolIntf"))
				r.WriteMapElemValue()
				if yyn2174 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolIntf == nil {
						r.EncodeNil()
					} else {
						yy2177 := *x.FptrMapBoolIntf
						yym2178 := z.EncBinary()
						_ = yym2178
						if false {
						} else {
							z.F.EncMapBoolIntfV(yy2177, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolString == nil {
					r.EncodeNil()
				} else {
					yym2180 := z.EncBinary()
					_ = yym2180
					if false {
					} else {
						z.F.EncMapBoolStringV(x.FMapBoolString, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolString"))
				r.WriteMapElemValue()
				if x.FMapBoolString == nil {
					r.EncodeNil()
				} else {
					yym2181 := z.EncBinary()
					_ = yym2181
					if false {
					} else {
						z.F.EncMapBoolStringV(x.FMapBoolString, e)
					}
				}
			}
			var yyn2182 bool
			if x.FptrMapBoolString == nil {
				yyn2182 = true
				goto LABEL2182
			}
		LABEL2182:
			if yyr2 || yy2arr2 {
				if yyn2182 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolString == nil {
						r.EncodeNil()
					} else {
						yy2183 := *x.FptrMapBoolString
						yym2184 := z.EncBinary()
						_ = yym2184
						if false {
						} else {
							z.F.EncMapBoolStringV(yy2183, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolString"))
				r.WriteMapElemValue()
				if yyn2182 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolString == nil {
						r.EncodeNil()
					} else {
						yy2185 := *x.FptrMapBoolString
						yym2186 := z.EncBinary()
						_ = yym2186
						if false {
						} else {
							z.F.EncMapBoolStringV(yy2185, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolUint == nil {
					r.EncodeNil()
				} else {
					yym2188 := z.EncBinary()
					_ = yym2188
					if false {
					} else {
						z.F.EncMapBoolUintV(x.FMapBoolUint, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolUint"))
				r.WriteMapElemValue()
				if x.FMapBoolUint == nil {
					r.EncodeNil()
				} else {
					yym2189 := z.EncBinary()
					_ = yym2189
					if false {
					} else {
						z.F.EncMapBoolUintV(x.FMapBoolUint, e)
					}
				}
			}
			var yyn2190 bool
			if x.FptrMapBoolUint == nil {
				yyn2190 = true
				goto LABEL2190
			}
		LABEL2190:
			if yyr2 || yy2arr2 {
				if yyn2190 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolUint == nil {
						r.EncodeNil()
					} else {
						yy2191 := *x.FptrMapBoolUint
						yym2192 := z.EncBinary()
						_ = yym2192
						if false {
						} else {
							z.F.EncMapBoolUintV(yy2191, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolUint"))
				r.WriteMapElemValue()
				if yyn2190 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolUint == nil {
						r.EncodeNil()
					} else {
						yy2193 := *x.FptrMapBoolUint
						yym2194 := z.EncBinary()
						_ = yym2194
						if false {
						} else {
							z.F.EncMapBoolUintV(yy2193, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolUint8 == nil {
					r.EncodeNil()
				} else {
					yym2196 := z.EncBinary()
					_ = yym2196
					if false {
					} else {
						z.F.EncMapBoolUint8V(x.FMapBoolUint8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolUint8"))
				r.WriteMapElemValue()
				if x.FMapBoolUint8 == nil {
					r.EncodeNil()
				} else {
					yym2197 := z.EncBinary()
					_ = yym2197
					if false {
					} else {
						z.F.EncMapBoolUint8V(x.FMapBoolUint8, e)
					}
				}
			}
			var yyn2198 bool
			if x.FptrMapBoolUint8 == nil {
				yyn2198 = true
				goto LABEL2198
			}
		LABEL2198:
			if yyr2 || yy2arr2 {
				if yyn2198 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolUint8 == nil {
						r.EncodeNil()
					} else {
						yy2199 := *x.FptrMapBoolUint8
						yym2200 := z.EncBinary()
						_ = yym2200
						if false {
						} else {
							z.F.EncMapBoolUint8V(yy2199, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolUint8"))
				r.WriteMapElemValue()
				if yyn2198 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolUint8 == nil {
						r.EncodeNil()
					} else {
						yy2201 := *x.FptrMapBoolUint8
						yym2202 := z.EncBinary()
						_ = yym2202
						if false {
						} else {
							z.F.EncMapBoolUint8V(yy2201, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolUint16 == nil {
					r.EncodeNil()
				} else {
					yym2204 := z.EncBinary()
					_ = yym2204
					if false {
					} else {
						z.F.EncMapBoolUint16V(x.FMapBoolUint16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolUint16"))
				r.WriteMapElemValue()
				if x.FMapBoolUint16 == nil {
					r.EncodeNil()
				} else {
					yym2205 := z.EncBinary()
					_ = yym2205
					if false {
					} else {
						z.F.EncMapBoolUint16V(x.FMapBoolUint16, e)
					}
				}
			}
			var yyn2206 bool
			if x.FptrMapBoolUint16 == nil {
				yyn2206 = true
				goto LABEL2206
			}
		LABEL2206:
			if yyr2 || yy2arr2 {
				if yyn2206 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolUint16 == nil {
						r.EncodeNil()
					} else {
						yy2207 := *x.FptrMapBoolUint16
						yym2208 := z.EncBinary()
						_ = yym2208
						if false {
						} else {
							z.F.EncMapBoolUint16V(yy2207, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolUint16"))
				r.WriteMapElemValue()
				if yyn2206 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolUint16 == nil {
						r.EncodeNil()
					} else {
						yy2209 := *x.FptrMapBoolUint16
						yym2210 := z.EncBinary()
						_ = yym2210
						if false {
						} else {
							z.F.EncMapBoolUint16V(yy2209, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolUint32 == nil {
					r.EncodeNil()
				} else {
					yym2212 := z.EncBinary()
					_ = yym2212
					if false {
					} else {
						z.F.EncMapBoolUint32V(x.FMapBoolUint32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolUint32"))
				r.WriteMapElemValue()
				if x.FMapBoolUint32 == nil {
					r.EncodeNil()
				} else {
					yym2213 := z.EncBinary()
					_ = yym2213
					if false {
					} else {
						z.F.EncMapBoolUint32V(x.FMapBoolUint32, e)
					}
				}
			}
			var yyn2214 bool
			if x.FptrMapBoolUint32 == nil {
				yyn2214 = true
				goto LABEL2214
			}
		LABEL2214:
			if yyr2 || yy2arr2 {
				if yyn2214 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolUint32 == nil {
						r.EncodeNil()
					} else {
						yy2215 := *x.FptrMapBoolUint32
						yym2216 := z.EncBinary()
						_ = yym2216
						if false {
						} else {
							z.F.EncMapBoolUint32V(yy2215, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolUint32"))
				r.WriteMapElemValue()
				if yyn2214 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolUint32 == nil {
						r.EncodeNil()
					} else {
						yy2217 := *x.FptrMapBoolUint32
						yym2218 := z.EncBinary()
						_ = yym2218
						if false {
						} else {
							z.F.EncMapBoolUint32V(yy2217, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolUint64 == nil {
					r.EncodeNil()
				} else {
					yym2220 := z.EncBinary()
					_ = yym2220
					if false {
					} else {
						z.F.EncMapBoolUint64V(x.FMapBoolUint64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolUint64"))
				r.WriteMapElemValue()
				if x.FMapBoolUint64 == nil {
					r.EncodeNil()
				} else {
					yym2221 := z.EncBinary()
					_ = yym2221
					if false {
					} else {
						z.F.EncMapBoolUint64V(x.FMapBoolUint64, e)
					}
				}
			}
			var yyn2222 bool
			if x.FptrMapBoolUint64 == nil {
				yyn2222 = true
				goto LABEL2222
			}
		LABEL2222:
			if yyr2 || yy2arr2 {
				if yyn2222 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolUint64 == nil {
						r.EncodeNil()
					} else {
						yy2223 := *x.FptrMapBoolUint64
						yym2224 := z.EncBinary()
						_ = yym2224
						if false {
						} else {
							z.F.EncMapBoolUint64V(yy2223, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolUint64"))
				r.WriteMapElemValue()
				if yyn2222 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolUint64 == nil {
						r.EncodeNil()
					} else {
						yy2225 := *x.FptrMapBoolUint64
						yym2226 := z.EncBinary()
						_ = yym2226
						if false {
						} else {
							z.F.EncMapBoolUint64V(yy2225, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolUintptr == nil {
					r.EncodeNil()
				} else {
					yym2228 := z.EncBinary()
					_ = yym2228
					if false {
					} else {
						z.F.EncMapBoolUintptrV(x.FMapBoolUintptr, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolUintptr"))
				r.WriteMapElemValue()
				if x.FMapBoolUintptr == nil {
					r.EncodeNil()
				} else {
					yym2229 := z.EncBinary()
					_ = yym2229
					if false {
					} else {
						z.F.EncMapBoolUintptrV(x.FMapBoolUintptr, e)
					}
				}
			}
			var yyn2230 bool
			if x.FptrMapBoolUintptr == nil {
				yyn2230 = true
				goto LABEL2230
			}
		LABEL2230:
			if yyr2 || yy2arr2 {
				if yyn2230 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolUintptr == nil {
						r.EncodeNil()
					} else {
						yy2231 := *x.FptrMapBoolUintptr
						yym2232 := z.EncBinary()
						_ = yym2232
						if false {
						} else {
							z.F.EncMapBoolUintptrV(yy2231, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolUintptr"))
				r.WriteMapElemValue()
				if yyn2230 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolUintptr == nil {
						r.EncodeNil()
					} else {
						yy2233 := *x.FptrMapBoolUintptr
						yym2234 := z.EncBinary()
						_ = yym2234
						if false {
						} else {
							z.F.EncMapBoolUintptrV(yy2233, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolInt == nil {
					r.EncodeNil()
				} else {
					yym2236 := z.EncBinary()
					_ = yym2236
					if false {
					} else {
						z.F.EncMapBoolIntV(x.FMapBoolInt, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolInt"))
				r.WriteMapElemValue()
				if x.FMapBoolInt == nil {
					r.EncodeNil()
				} else {
					yym2237 := z.EncBinary()
					_ = yym2237
					if false {
					} else {
						z.F.EncMapBoolIntV(x.FMapBoolInt, e)
					}
				}
			}
			var yyn2238 bool
			if x.FptrMapBoolInt == nil {
				yyn2238 = true
				goto LABEL2238
			}
		LABEL2238:
			if yyr2 || yy2arr2 {
				if yyn2238 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolInt == nil {
						r.EncodeNil()
					} else {
						yy2239 := *x.FptrMapBoolInt
						yym2240 := z.EncBinary()
						_ = yym2240
						if false {
						} else {
							z.F.EncMapBoolIntV(yy2239, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolInt"))
				r.WriteMapElemValue()
				if yyn2238 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolInt == nil {
						r.EncodeNil()
					} else {
						yy2241 := *x.FptrMapBoolInt
						yym2242 := z.EncBinary()
						_ = yym2242
						if false {
						} else {
							z.F.EncMapBoolIntV(yy2241, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolInt8 == nil {
					r.EncodeNil()
				} else {
					yym2244 := z.EncBinary()
					_ = yym2244
					if false {
					} else {
						z.F.EncMapBoolInt8V(x.FMapBoolInt8, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolInt8"))
				r.WriteMapElemValue()
				if x.FMapBoolInt8 == nil {
					r.EncodeNil()
				} else {
					yym2245 := z.EncBinary()
					_ = yym2245
					if false {
					} else {
						z.F.EncMapBoolInt8V(x.FMapBoolInt8, e)
					}
				}
			}
			var yyn2246 bool
			if x.FptrMapBoolInt8 == nil {
				yyn2246 = true
				goto LABEL2246
			}
		LABEL2246:
			if yyr2 || yy2arr2 {
				if yyn2246 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolInt8 == nil {
						r.EncodeNil()
					} else {
						yy2247 := *x.FptrMapBoolInt8
						yym2248 := z.EncBinary()
						_ = yym2248
						if false {
						} else {
							z.F.EncMapBoolInt8V(yy2247, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolInt8"))
				r.WriteMapElemValue()
				if yyn2246 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolInt8 == nil {
						r.EncodeNil()
					} else {
						yy2249 := *x.FptrMapBoolInt8
						yym2250 := z.EncBinary()
						_ = yym2250
						if false {
						} else {
							z.F.EncMapBoolInt8V(yy2249, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolInt16 == nil {
					r.EncodeNil()
				} else {
					yym2252 := z.EncBinary()
					_ = yym2252
					if false {
					} else {
						z.F.EncMapBoolInt16V(x.FMapBoolInt16, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolInt16"))
				r.WriteMapElemValue()
				if x.FMapBoolInt16 == nil {
					r.EncodeNil()
				} else {
					yym2253 := z.EncBinary()
					_ = yym2253
					if false {
					} else {
						z.F.EncMapBoolInt16V(x.FMapBoolInt16, e)
					}
				}
			}
			var yyn2254 bool
			if x.FptrMapBoolInt16 == nil {
				yyn2254 = true
				goto LABEL2254
			}
		LABEL2254:
			if yyr2 || yy2arr2 {
				if yyn2254 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolInt16 == nil {
						r.EncodeNil()
					} else {
						yy2255 := *x.FptrMapBoolInt16
						yym2256 := z.EncBinary()
						_ = yym2256
						if false {
						} else {
							z.F.EncMapBoolInt16V(yy2255, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolInt16"))
				r.WriteMapElemValue()
				if yyn2254 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolInt16 == nil {
						r.EncodeNil()
					} else {
						yy2257 := *x.FptrMapBoolInt16
						yym2258 := z.EncBinary()
						_ = yym2258
						if false {
						} else {
							z.F.EncMapBoolInt16V(yy2257, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolInt32 == nil {
					r.EncodeNil()
				} else {
					yym2260 := z.EncBinary()
					_ = yym2260
					if false {
					} else {
						z.F.EncMapBoolInt32V(x.FMapBoolInt32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolInt32"))
				r.WriteMapElemValue()
				if x.FMapBoolInt32 == nil {
					r.EncodeNil()
				} else {
					yym2261 := z.EncBinary()
					_ = yym2261
					if false {
					} else {
						z.F.EncMapBoolInt32V(x.FMapBoolInt32, e)
					}
				}
			}
			var yyn2262 bool
			if x.FptrMapBoolInt32 == nil {
				yyn2262 = true
				goto LABEL2262
			}
		LABEL2262:
			if yyr2 || yy2arr2 {
				if yyn2262 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolInt32 == nil {
						r.EncodeNil()
					} else {
						yy2263 := *x.FptrMapBoolInt32
						yym2264 := z.EncBinary()
						_ = yym2264
						if false {
						} else {
							z.F.EncMapBoolInt32V(yy2263, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolInt32"))
				r.WriteMapElemValue()
				if yyn2262 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolInt32 == nil {
						r.EncodeNil()
					} else {
						yy2265 := *x.FptrMapBoolInt32
						yym2266 := z.EncBinary()
						_ = yym2266
						if false {
						} else {
							z.F.EncMapBoolInt32V(yy2265, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolInt64 == nil {
					r.EncodeNil()
				} else {
					yym2268 := z.EncBinary()
					_ = yym2268
					if false {
					} else {
						z.F.EncMapBoolInt64V(x.FMapBoolInt64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolInt64"))
				r.WriteMapElemValue()
				if x.FMapBoolInt64 == nil {
					r.EncodeNil()
				} else {
					yym2269 := z.EncBinary()
					_ = yym2269
					if false {
					} else {
						z.F.EncMapBoolInt64V(x.FMapBoolInt64, e)
					}
				}
			}
			var yyn2270 bool
			if x.FptrMapBoolInt64 == nil {
				yyn2270 = true
				goto LABEL2270
			}
		LABEL2270:
			if yyr2 || yy2arr2 {
				if yyn2270 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolInt64 == nil {
						r.EncodeNil()
					} else {
						yy2271 := *x.FptrMapBoolInt64
						yym2272 := z.EncBinary()
						_ = yym2272
						if false {
						} else {
							z.F.EncMapBoolInt64V(yy2271, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolInt64"))
				r.WriteMapElemValue()
				if yyn2270 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolInt64 == nil {
						r.EncodeNil()
					} else {
						yy2273 := *x.FptrMapBoolInt64
						yym2274 := z.EncBinary()
						_ = yym2274
						if false {
						} else {
							z.F.EncMapBoolInt64V(yy2273, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolFloat32 == nil {
					r.EncodeNil()
				} else {
					yym2276 := z.EncBinary()
					_ = yym2276
					if false {
					} else {
						z.F.EncMapBoolFloat32V(x.FMapBoolFloat32, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolFloat32"))
				r.WriteMapElemValue()
				if x.FMapBoolFloat32 == nil {
					r.EncodeNil()
				} else {
					yym2277 := z.EncBinary()
					_ = yym2277
					if false {
					} else {
						z.F.EncMapBoolFloat32V(x.FMapBoolFloat32, e)
					}
				}
			}
			var yyn2278 bool
			if x.FptrMapBoolFloat32 == nil {
				yyn2278 = true
				goto LABEL2278
			}
		LABEL2278:
			if yyr2 || yy2arr2 {
				if yyn2278 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolFloat32 == nil {
						r.EncodeNil()
					} else {
						yy2279 := *x.FptrMapBoolFloat32
						yym2280 := z.EncBinary()
						_ = yym2280
						if false {
						} else {
							z.F.EncMapBoolFloat32V(yy2279, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolFloat32"))
				r.WriteMapElemValue()
				if yyn2278 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolFloat32 == nil {
						r.EncodeNil()
					} else {
						yy2281 := *x.FptrMapBoolFloat32
						yym2282 := z.EncBinary()
						_ = yym2282
						if false {
						} else {
							z.F.EncMapBoolFloat32V(yy2281, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolFloat64 == nil {
					r.EncodeNil()
				} else {
					yym2284 := z.EncBinary()
					_ = yym2284
					if false {
					} else {
						z.F.EncMapBoolFloat64V(x.FMapBoolFloat64, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolFloat64"))
				r.WriteMapElemValue()
				if x.FMapBoolFloat64 == nil {
					r.EncodeNil()
				} else {
					yym2285 := z.EncBinary()
					_ = yym2285
					if false {
					} else {
						z.F.EncMapBoolFloat64V(x.FMapBoolFloat64, e)
					}
				}
			}
			var yyn2286 bool
			if x.FptrMapBoolFloat64 == nil {
				yyn2286 = true
				goto LABEL2286
			}
		LABEL2286:
			if yyr2 || yy2arr2 {
				if yyn2286 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolFloat64 == nil {
						r.EncodeNil()
					} else {
						yy2287 := *x.FptrMapBoolFloat64
						yym2288 := z.EncBinary()
						_ = yym2288
						if false {
						} else {
							z.F.EncMapBoolFloat64V(yy2287, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolFloat64"))
				r.WriteMapElemValue()
				if yyn2286 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolFloat64 == nil {
						r.EncodeNil()
					} else {
						yy2289 := *x.FptrMapBoolFloat64
						yym2290 := z.EncBinary()
						_ = yym2290
						if false {
						} else {
							z.F.EncMapBoolFloat64V(yy2289, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.FMapBoolBool == nil {
					r.EncodeNil()
				} else {
					yym2292 := z.EncBinary()
					_ = yym2292
					if false {
					} else {
						z.F.EncMapBoolBoolV(x.FMapBoolBool, e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FMapBoolBool"))
				r.WriteMapElemValue()
				if x.FMapBoolBool == nil {
					r.EncodeNil()
				} else {
					yym2293 := z.EncBinary()
					_ = yym2293
					if false {
					} else {
						z.F.EncMapBoolBoolV(x.FMapBoolBool, e)
					}
				}
			}
			var yyn2294 bool
			if x.FptrMapBoolBool == nil {
				yyn2294 = true
				goto LABEL2294
			}
		LABEL2294:
			if yyr2 || yy2arr2 {
				if yyn2294 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.FptrMapBoolBool == nil {
						r.EncodeNil()
					} else {
						yy2295 := *x.FptrMapBoolBool
						yym2296 := z.EncBinary()
						_ = yym2296
						if false {
						} else {
							z.F.EncMapBoolBoolV(yy2295, e)
						}
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("FptrMapBoolBool"))
				r.WriteMapElemValue()
				if yyn2294 {
					r.EncodeNil()
				} else {
					if x.FptrMapBoolBool == nil {
						r.EncodeNil()
					} else {
						yy2297 := *x.FptrMapBoolBool
						yym2298 := z.EncBinary()
						_ = yym2298
						if false {
						} else {
							z.F.EncMapBoolBoolV(yy2297, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *TestMammoth2) CodecDecodeSelf(d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap19781 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray19781 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr19781)
		}
	}
}

func (x *TestMammoth2) codecDecodeSelfFromMap(l int, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "FIntf":
			if r.TryDecodeAsNil() {
				x.FIntf = nil
			} else {
				yyv4 := &x.FIntf
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					z.DecFallback(yyv4, true)
				}
			}
		case "FptrIntf":
			if x.FptrIntf == nil {
				x.FptrIntf = new(interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrIntf != nil {
					x.FptrIntf = nil
				}
			} else {
				if x.FptrIntf == nil {
					x.FptrIntf = new(interface{})
				}
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					z.DecFallback(x.FptrIntf, true)
				}
			}
		case "FString":
			if r.TryDecodeAsNil() {
				x.FString = ""
			} else {
				yyv8 := &x.FString
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*string)(yyv8)) = r.DecodeString()
				}
			}
		case "FptrString":
			if x.FptrString == nil {
				x.FptrString = new(string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrString != nil {
					x.FptrString = nil
				}
			} else {
				if x.FptrString == nil {
					x.FptrString = new(string)
				}
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*string)(x.FptrString)) = r.DecodeString()
				}
			}
		case "FFloat32":
			if r.TryDecodeAsNil() {
				x.FFloat32 = 0
			} else {
				yyv12 := &x.FFloat32
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					*((*float32)(yyv12)) = float32(r.DecodeFloat(true))
				}
			}
		case "FptrFloat32":
			if x.FptrFloat32 == nil {
				x.FptrFloat32 = new(float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrFloat32 != nil {
					x.FptrFloat32 = nil
				}
			} else {
				if x.FptrFloat32 == nil {
					x.FptrFloat32 = new(float32)
				}
				yym15 := z.DecBinary()
				_ = yym15
				if false {
				} else {
					*((*float32)(x.FptrFloat32)) = float32(r.DecodeFloat(true))
				}
			}
		case "FFloat64":
			if r.TryDecodeAsNil() {
				x.FFloat64 = 0
			} else {
				yyv16 := &x.FFloat64
				yym17 := z.DecBinary()
				_ = yym17
				if false {
				} else {
					*((*float64)(yyv16)) = float64(r.DecodeFloat(false))
				}
			}
		case "FptrFloat64":
			if x.FptrFloat64 == nil {
				x.FptrFloat64 = new(float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrFloat64 != nil {
					x.FptrFloat64 = nil
				}
			} else {
				if x.FptrFloat64 == nil {
					x.FptrFloat64 = new(float64)
				}
				yym19 := z.DecBinary()
				_ = yym19
				if false {
				} else {
					*((*float64)(x.FptrFloat64)) = float64(r.DecodeFloat(false))
				}
			}
		case "FUint":
			if r.TryDecodeAsNil() {
				x.FUint = 0
			} else {
				yyv20 := &x.FUint
				yym21 := z.DecBinary()
				_ = yym21
				if false {
				} else {
					*((*uint)(yyv20)) = uint(r.DecodeUint(codecSelferBitsize19781))
				}
			}
		case "FptrUint":
			if x.FptrUint == nil {
				x.FptrUint = new(uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrUint != nil {
					x.FptrUint = nil
				}
			} else {
				if x.FptrUint == nil {
					x.FptrUint = new(uint)
				}
				yym23 := z.DecBinary()
				_ = yym23
				if false {
				} else {
					*((*uint)(x.FptrUint)) = uint(r.DecodeUint(codecSelferBitsize19781))
				}
			}
		case "FUint8":
			if r.TryDecodeAsNil() {
				x.FUint8 = 0
			} else {
				yyv24 := &x.FUint8
				yym25 := z.DecBinary()
				_ = yym25
				if false {
				} else {
					*((*uint8)(yyv24)) = uint8(r.DecodeUint(8))
				}
			}
		case "FptrUint8":
			if x.FptrUint8 == nil {
				x.FptrUint8 = new(uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrUint8 != nil {
					x.FptrUint8 = nil
				}
			} else {
				if x.FptrUint8 == nil {
					x.FptrUint8 = new(uint8)
				}
				yym27 := z.DecBinary()
				_ = yym27
				if false {
				} else {
					*((*uint8)(x.FptrUint8)) = uint8(r.DecodeUint(8))
				}
			}
		case "FUint16":
			if r.TryDecodeAsNil() {
				x.FUint16 = 0
			} else {
				yyv28 := &x.FUint16
				yym29 := z.DecBinary()
				_ = yym29
				if false {
				} else {
					*((*uint16)(yyv28)) = uint16(r.DecodeUint(16))
				}
			}
		case "FptrUint16":
			if x.FptrUint16 == nil {
				x.FptrUint16 = new(uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrUint16 != nil {
					x.FptrUint16 = nil
				}
			} else {
				if x.FptrUint16 == nil {
					x.FptrUint16 = new(uint16)
				}
				yym31 := z.DecBinary()
				_ = yym31
				if false {
				} else {
					*((*uint16)(x.FptrUint16)) = uint16(r.DecodeUint(16))
				}
			}
		case "FUint32":
			if r.TryDecodeAsNil() {
				x.FUint32 = 0
			} else {
				yyv32 := &x.FUint32
				yym33 := z.DecBinary()
				_ = yym33
				if false {
				} else {
					*((*uint32)(yyv32)) = uint32(r.DecodeUint(32))
				}
			}
		case "FptrUint32":
			if x.FptrUint32 == nil {
				x.FptrUint32 = new(uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrUint32 != nil {
					x.FptrUint32 = nil
				}
			} else {
				if x.FptrUint32 == nil {
					x.FptrUint32 = new(uint32)
				}
				yym35 := z.DecBinary()
				_ = yym35
				if false {
				} else {
					*((*uint32)(x.FptrUint32)) = uint32(r.DecodeUint(32))
				}
			}
		case "FUint64":
			if r.TryDecodeAsNil() {
				x.FUint64 = 0
			} else {
				yyv36 := &x.FUint64
				yym37 := z.DecBinary()
				_ = yym37
				if false {
				} else {
					*((*uint64)(yyv36)) = uint64(r.DecodeUint(64))
				}
			}
		case "FptrUint64":
			if x.FptrUint64 == nil {
				x.FptrUint64 = new(uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrUint64 != nil {
					x.FptrUint64 = nil
				}
			} else {
				if x.FptrUint64 == nil {
					x.FptrUint64 = new(uint64)
				}
				yym39 := z.DecBinary()
				_ = yym39
				if false {
				} else {
					*((*uint64)(x.FptrUint64)) = uint64(r.DecodeUint(64))
				}
			}
		case "FUintptr":
			if r.TryDecodeAsNil() {
				x.FUintptr = 0
			} else {
				yyv40 := &x.FUintptr
				yym41 := z.DecBinary()
				_ = yym41
				if false {
				} else {
					*((*uintptr)(yyv40)) = uintptr(r.DecodeUint(codecSelferBitsize19781))
				}
			}
		case "FptrUintptr":
			if x.FptrUintptr == nil {
				x.FptrUintptr = new(uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrUintptr != nil {
					x.FptrUintptr = nil
				}
			} else {
				if x.FptrUintptr == nil {
					x.FptrUintptr = new(uintptr)
				}
				yym43 := z.DecBinary()
				_ = yym43
				if false {
				} else {
					*((*uintptr)(x.FptrUintptr)) = uintptr(r.DecodeUint(codecSelferBitsize19781))
				}
			}
		case "FInt":
			if r.TryDecodeAsNil() {
				x.FInt = 0
			} else {
				yyv44 := &x.FInt
				yym45 := z.DecBinary()
				_ = yym45
				if false {
				} else {
					*((*int)(yyv44)) = int(r.DecodeInt(codecSelferBitsize19781))
				}
			}
		case "FptrInt":
			if x.FptrInt == nil {
				x.FptrInt = new(int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrInt != nil {
					x.FptrInt = nil
				}
			} else {
				if x.FptrInt == nil {
					x.FptrInt = new(int)
				}
				yym47 := z.DecBinary()
				_ = yym47
				if false {
				} else {
					*((*int)(x.FptrInt)) = int(r.DecodeInt(codecSelferBitsize19781))
				}
			}
		case "FInt8":
			if r.TryDecodeAsNil() {
				x.FInt8 = 0
			} else {
				yyv48 := &x.FInt8
				yym49 := z.DecBinary()
				_ = yym49
				if false {
				} else {
					*((*int8)(yyv48)) = int8(r.DecodeInt(8))
				}
			}
		case "FptrInt8":
			if x.FptrInt8 == nil {
				x.FptrInt8 = new(int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrInt8 != nil {
					x.FptrInt8 = nil
				}
			} else {
				if x.FptrInt8 == nil {
					x.FptrInt8 = new(int8)
				}
				yym51 := z.DecBinary()
				_ = yym51
				if false {
				} else {
					*((*int8)(x.FptrInt8)) = int8(r.DecodeInt(8))
				}
			}
		case "FInt16":
			if r.TryDecodeAsNil() {
				x.FInt16 = 0
			} else {
				yyv52 := &x.FInt16
				yym53 := z.DecBinary()
				_ = yym53
				if false {
				} else {
					*((*int16)(yyv52)) = int16(r.DecodeInt(16))
				}
			}
		case "FptrInt16":
			if x.FptrInt16 == nil {
				x.FptrInt16 = new(int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrInt16 != nil {
					x.FptrInt16 = nil
				}
			} else {
				if x.FptrInt16 == nil {
					x.FptrInt16 = new(int16)
				}
				yym55 := z.DecBinary()
				_ = yym55
				if false {
				} else {
					*((*int16)(x.FptrInt16)) = int16(r.DecodeInt(16))
				}
			}
		case "FInt32":
			if r.TryDecodeAsNil() {
				x.FInt32 = 0
			} else {
				yyv56 := &x.FInt32
				yym57 := z.DecBinary()
				_ = yym57
				if false {
				} else {
					*((*int32)(yyv56)) = int32(r.DecodeInt(32))
				}
			}
		case "FptrInt32":
			if x.FptrInt32 == nil {
				x.FptrInt32 = new(int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrInt32 != nil {
					x.FptrInt32 = nil
				}
			} else {
				if x.FptrInt32 == nil {
					x.FptrInt32 = new(int32)
				}
				yym59 := z.DecBinary()
				_ = yym59
				if false {
				} else {
					*((*int32)(x.FptrInt32)) = int32(r.DecodeInt(32))
				}
			}
		case "FInt64":
			if r.TryDecodeAsNil() {
				x.FInt64 = 0
			} else {
				yyv60 := &x.FInt64
				yym61 := z.DecBinary()
				_ = yym61
				if false {
				} else {
					*((*int64)(yyv60)) = int64(r.DecodeInt(64))
				}
			}
		case "FptrInt64":
			if x.FptrInt64 == nil {
				x.FptrInt64 = new(int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrInt64 != nil {
					x.FptrInt64 = nil
				}
			} else {
				if x.FptrInt64 == nil {
					x.FptrInt64 = new(int64)
				}
				yym63 := z.DecBinary()
				_ = yym63
				if false {
				} else {
					*((*int64)(x.FptrInt64)) = int64(r.DecodeInt(64))
				}
			}
		case "FBool":
			if r.TryDecodeAsNil() {
				x.FBool = false
			} else {
				yyv64 := &x.FBool
				yym65 := z.DecBinary()
				_ = yym65
				if false {
				} else {
					*((*bool)(yyv64)) = r.DecodeBool()
				}
			}
		case "FptrBool":
			if x.FptrBool == nil {
				x.FptrBool = new(bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrBool != nil {
					x.FptrBool = nil
				}
			} else {
				if x.FptrBool == nil {
					x.FptrBool = new(bool)
				}
				yym67 := z.DecBinary()
				_ = yym67
				if false {
				} else {
					*((*bool)(x.FptrBool)) = r.DecodeBool()
				}
			}
		case "FSliceIntf":
			if r.TryDecodeAsNil() {
				x.FSliceIntf = nil
			} else {
				yyv68 := &x.FSliceIntf
				yym69 := z.DecBinary()
				_ = yym69
				if false {
				} else {
					z.F.DecSliceIntfX(yyv68, d)
				}
			}
		case "FptrSliceIntf":
			if x.FptrSliceIntf == nil {
				x.FptrSliceIntf = new([]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceIntf != nil {
					x.FptrSliceIntf = nil
				}
			} else {
				if x.FptrSliceIntf == nil {
					x.FptrSliceIntf = new([]interface{})
				}
				yym71 := z.DecBinary()
				_ = yym71
				if false {
				} else {
					z.F.DecSliceIntfX(x.FptrSliceIntf, d)
				}
			}
		case "FSliceString":
			if r.TryDecodeAsNil() {
				x.FSliceString = nil
			} else {
				yyv72 := &x.FSliceString
				yym73 := z.DecBinary()
				_ = yym73
				if false {
				} else {
					z.F.DecSliceStringX(yyv72, d)
				}
			}
		case "FptrSliceString":
			if x.FptrSliceString == nil {
				x.FptrSliceString = new([]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceString != nil {
					x.FptrSliceString = nil
				}
			} else {
				if x.FptrSliceString == nil {
					x.FptrSliceString = new([]string)
				}
				yym75 := z.DecBinary()
				_ = yym75
				if false {
				} else {
					z.F.DecSliceStringX(x.FptrSliceString, d)
				}
			}
		case "FSliceFloat32":
			if r.TryDecodeAsNil() {
				x.FSliceFloat32 = nil
			} else {
				yyv76 := &x.FSliceFloat32
				yym77 := z.DecBinary()
				_ = yym77
				if false {
				} else {
					z.F.DecSliceFloat32X(yyv76, d)
				}
			}
		case "FptrSliceFloat32":
			if x.FptrSliceFloat32 == nil {
				x.FptrSliceFloat32 = new([]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceFloat32 != nil {
					x.FptrSliceFloat32 = nil
				}
			} else {
				if x.FptrSliceFloat32 == nil {
					x.FptrSliceFloat32 = new([]float32)
				}
				yym79 := z.DecBinary()
				_ = yym79
				if false {
				} else {
					z.F.DecSliceFloat32X(x.FptrSliceFloat32, d)
				}
			}
		case "FSliceFloat64":
			if r.TryDecodeAsNil() {
				x.FSliceFloat64 = nil
			} else {
				yyv80 := &x.FSliceFloat64
				yym81 := z.DecBinary()
				_ = yym81
				if false {
				} else {
					z.F.DecSliceFloat64X(yyv80, d)
				}
			}
		case "FptrSliceFloat64":
			if x.FptrSliceFloat64 == nil {
				x.FptrSliceFloat64 = new([]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceFloat64 != nil {
					x.FptrSliceFloat64 = nil
				}
			} else {
				if x.FptrSliceFloat64 == nil {
					x.FptrSliceFloat64 = new([]float64)
				}
				yym83 := z.DecBinary()
				_ = yym83
				if false {
				} else {
					z.F.DecSliceFloat64X(x.FptrSliceFloat64, d)
				}
			}
		case "FSliceUint":
			if r.TryDecodeAsNil() {
				x.FSliceUint = nil
			} else {
				yyv84 := &x.FSliceUint
				yym85 := z.DecBinary()
				_ = yym85
				if false {
				} else {
					z.F.DecSliceUintX(yyv84, d)
				}
			}
		case "FptrSliceUint":
			if x.FptrSliceUint == nil {
				x.FptrSliceUint = new([]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceUint != nil {
					x.FptrSliceUint = nil
				}
			} else {
				if x.FptrSliceUint == nil {
					x.FptrSliceUint = new([]uint)
				}
				yym87 := z.DecBinary()
				_ = yym87
				if false {
				} else {
					z.F.DecSliceUintX(x.FptrSliceUint, d)
				}
			}
		case "FSliceUint16":
			if r.TryDecodeAsNil() {
				x.FSliceUint16 = nil
			} else {
				yyv88 := &x.FSliceUint16
				yym89 := z.DecBinary()
				_ = yym89
				if false {
				} else {
					z.F.DecSliceUint16X(yyv88, d)
				}
			}
		case "FptrSliceUint16":
			if x.FptrSliceUint16 == nil {
				x.FptrSliceUint16 = new([]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceUint16 != nil {
					x.FptrSliceUint16 = nil
				}
			} else {
				if x.FptrSliceUint16 == nil {
					x.FptrSliceUint16 = new([]uint16)
				}
				yym91 := z.DecBinary()
				_ = yym91
				if false {
				} else {
					z.F.DecSliceUint16X(x.FptrSliceUint16, d)
				}
			}
		case "FSliceUint32":
			if r.TryDecodeAsNil() {
				x.FSliceUint32 = nil
			} else {
				yyv92 := &x.FSliceUint32
				yym93 := z.DecBinary()
				_ = yym93
				if false {
				} else {
					z.F.DecSliceUint32X(yyv92, d)
				}
			}
		case "FptrSliceUint32":
			if x.FptrSliceUint32 == nil {
				x.FptrSliceUint32 = new([]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceUint32 != nil {
					x.FptrSliceUint32 = nil
				}
			} else {
				if x.FptrSliceUint32 == nil {
					x.FptrSliceUint32 = new([]uint32)
				}
				yym95 := z.DecBinary()
				_ = yym95
				if false {
				} else {
					z.F.DecSliceUint32X(x.FptrSliceUint32, d)
				}
			}
		case "FSliceUint64":
			if r.TryDecodeAsNil() {
				x.FSliceUint64 = nil
			} else {
				yyv96 := &x.FSliceUint64
				yym97 := z.DecBinary()
				_ = yym97
				if false {
				} else {
					z.F.DecSliceUint64X(yyv96, d)
				}
			}
		case "FptrSliceUint64":
			if x.FptrSliceUint64 == nil {
				x.FptrSliceUint64 = new([]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceUint64 != nil {
					x.FptrSliceUint64 = nil
				}
			} else {
				if x.FptrSliceUint64 == nil {
					x.FptrSliceUint64 = new([]uint64)
				}
				yym99 := z.DecBinary()
				_ = yym99
				if false {
				} else {
					z.F.DecSliceUint64X(x.FptrSliceUint64, d)
				}
			}
		case "FSliceUintptr":
			if r.TryDecodeAsNil() {
				x.FSliceUintptr = nil
			} else {
				yyv100 := &x.FSliceUintptr
				yym101 := z.DecBinary()
				_ = yym101
				if false {
				} else {
					z.F.DecSliceUintptrX(yyv100, d)
				}
			}
		case "FptrSliceUintptr":
			if x.FptrSliceUintptr == nil {
				x.FptrSliceUintptr = new([]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceUintptr != nil {
					x.FptrSliceUintptr = nil
				}
			} else {
				if x.FptrSliceUintptr == nil {
					x.FptrSliceUintptr = new([]uintptr)
				}
				yym103 := z.DecBinary()
				_ = yym103
				if false {
				} else {
					z.F.DecSliceUintptrX(x.FptrSliceUintptr, d)
				}
			}
		case "FSliceInt":
			if r.TryDecodeAsNil() {
				x.FSliceInt = nil
			} else {
				yyv104 := &x.FSliceInt
				yym105 := z.DecBinary()
				_ = yym105
				if false {
				} else {
					z.F.DecSliceIntX(yyv104, d)
				}
			}
		case "FptrSliceInt":
			if x.FptrSliceInt == nil {
				x.FptrSliceInt = new([]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceInt != nil {
					x.FptrSliceInt = nil
				}
			} else {
				if x.FptrSliceInt == nil {
					x.FptrSliceInt = new([]int)
				}
				yym107 := z.DecBinary()
				_ = yym107
				if false {
				} else {
					z.F.DecSliceIntX(x.FptrSliceInt, d)
				}
			}
		case "FSliceInt8":
			if r.TryDecodeAsNil() {
				x.FSliceInt8 = nil
			} else {
				yyv108 := &x.FSliceInt8
				yym109 := z.DecBinary()
				_ = yym109
				if false {
				} else {
					z.F.DecSliceInt8X(yyv108, d)
				}
			}
		case "FptrSliceInt8":
			if x.FptrSliceInt8 == nil {
				x.FptrSliceInt8 = new([]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceInt8 != nil {
					x.FptrSliceInt8 = nil
				}
			} else {
				if x.FptrSliceInt8 == nil {
					x.FptrSliceInt8 = new([]int8)
				}
				yym111 := z.DecBinary()
				_ = yym111
				if false {
				} else {
					z.F.DecSliceInt8X(x.FptrSliceInt8, d)
				}
			}
		case "FSliceInt16":
			if r.TryDecodeAsNil() {
				x.FSliceInt16 = nil
			} else {
				yyv112 := &x.FSliceInt16
				yym113 := z.DecBinary()
				_ = yym113
				if false {
				} else {
					z.F.DecSliceInt16X(yyv112, d)
				}
			}
		case "FptrSliceInt16":
			if x.FptrSliceInt16 == nil {
				x.FptrSliceInt16 = new([]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceInt16 != nil {
					x.FptrSliceInt16 = nil
				}
			} else {
				if x.FptrSliceInt16 == nil {
					x.FptrSliceInt16 = new([]int16)
				}
				yym115 := z.DecBinary()
				_ = yym115
				if false {
				} else {
					z.F.DecSliceInt16X(x.FptrSliceInt16, d)
				}
			}
		case "FSliceInt32":
			if r.TryDecodeAsNil() {
				x.FSliceInt32 = nil
			} else {
				yyv116 := &x.FSliceInt32
				yym117 := z.DecBinary()
				_ = yym117
				if false {
				} else {
					z.F.DecSliceInt32X(yyv116, d)
				}
			}
		case "FptrSliceInt32":
			if x.FptrSliceInt32 == nil {
				x.FptrSliceInt32 = new([]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceInt32 != nil {
					x.FptrSliceInt32 = nil
				}
			} else {
				if x.FptrSliceInt32 == nil {
					x.FptrSliceInt32 = new([]int32)
				}
				yym119 := z.DecBinary()
				_ = yym119
				if false {
				} else {
					z.F.DecSliceInt32X(x.FptrSliceInt32, d)
				}
			}
		case "FSliceInt64":
			if r.TryDecodeAsNil() {
				x.FSliceInt64 = nil
			} else {
				yyv120 := &x.FSliceInt64
				yym121 := z.DecBinary()
				_ = yym121
				if false {
				} else {
					z.F.DecSliceInt64X(yyv120, d)
				}
			}
		case "FptrSliceInt64":
			if x.FptrSliceInt64 == nil {
				x.FptrSliceInt64 = new([]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceInt64 != nil {
					x.FptrSliceInt64 = nil
				}
			} else {
				if x.FptrSliceInt64 == nil {
					x.FptrSliceInt64 = new([]int64)
				}
				yym123 := z.DecBinary()
				_ = yym123
				if false {
				} else {
					z.F.DecSliceInt64X(x.FptrSliceInt64, d)
				}
			}
		case "FSliceBool":
			if r.TryDecodeAsNil() {
				x.FSliceBool = nil
			} else {
				yyv124 := &x.FSliceBool
				yym125 := z.DecBinary()
				_ = yym125
				if false {
				} else {
					z.F.DecSliceBoolX(yyv124, d)
				}
			}
		case "FptrSliceBool":
			if x.FptrSliceBool == nil {
				x.FptrSliceBool = new([]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrSliceBool != nil {
					x.FptrSliceBool = nil
				}
			} else {
				if x.FptrSliceBool == nil {
					x.FptrSliceBool = new([]bool)
				}
				yym127 := z.DecBinary()
				_ = yym127
				if false {
				} else {
					z.F.DecSliceBoolX(x.FptrSliceBool, d)
				}
			}
		case "FMapIntfIntf":
			if r.TryDecodeAsNil() {
				x.FMapIntfIntf = nil
			} else {
				yyv128 := &x.FMapIntfIntf
				yym129 := z.DecBinary()
				_ = yym129
				if false {
				} else {
					z.F.DecMapIntfIntfX(yyv128, d)
				}
			}
		case "FptrMapIntfIntf":
			if x.FptrMapIntfIntf == nil {
				x.FptrMapIntfIntf = new(map[interface{}]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfIntf != nil {
					x.FptrMapIntfIntf = nil
				}
			} else {
				if x.FptrMapIntfIntf == nil {
					x.FptrMapIntfIntf = new(map[interface{}]interface{})
				}
				yym131 := z.DecBinary()
				_ = yym131
				if false {
				} else {
					z.F.DecMapIntfIntfX(x.FptrMapIntfIntf, d)
				}
			}
		case "FMapIntfString":
			if r.TryDecodeAsNil() {
				x.FMapIntfString = nil
			} else {
				yyv132 := &x.FMapIntfString
				yym133 := z.DecBinary()
				_ = yym133
				if false {
				} else {
					z.F.DecMapIntfStringX(yyv132, d)
				}
			}
		case "FptrMapIntfString":
			if x.FptrMapIntfString == nil {
				x.FptrMapIntfString = new(map[interface{}]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfString != nil {
					x.FptrMapIntfString = nil
				}
			} else {
				if x.FptrMapIntfString == nil {
					x.FptrMapIntfString = new(map[interface{}]string)
				}
				yym135 := z.DecBinary()
				_ = yym135
				if false {
				} else {
					z.F.DecMapIntfStringX(x.FptrMapIntfString, d)
				}
			}
		case "FMapIntfUint":
			if r.TryDecodeAsNil() {
				x.FMapIntfUint = nil
			} else {
				yyv136 := &x.FMapIntfUint
				yym137 := z.DecBinary()
				_ = yym137
				if false {
				} else {
					z.F.DecMapIntfUintX(yyv136, d)
				}
			}
		case "FptrMapIntfUint":
			if x.FptrMapIntfUint == nil {
				x.FptrMapIntfUint = new(map[interface{}]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfUint != nil {
					x.FptrMapIntfUint = nil
				}
			} else {
				if x.FptrMapIntfUint == nil {
					x.FptrMapIntfUint = new(map[interface{}]uint)
				}
				yym139 := z.DecBinary()
				_ = yym139
				if false {
				} else {
					z.F.DecMapIntfUintX(x.FptrMapIntfUint, d)
				}
			}
		case "FMapIntfUint8":
			if r.TryDecodeAsNil() {
				x.FMapIntfUint8 = nil
			} else {
				yyv140 := &x.FMapIntfUint8
				yym141 := z.DecBinary()
				_ = yym141
				if false {
				} else {
					z.F.DecMapIntfUint8X(yyv140, d)
				}
			}
		case "FptrMapIntfUint8":
			if x.FptrMapIntfUint8 == nil {
				x.FptrMapIntfUint8 = new(map[interface{}]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfUint8 != nil {
					x.FptrMapIntfUint8 = nil
				}
			} else {
				if x.FptrMapIntfUint8 == nil {
					x.FptrMapIntfUint8 = new(map[interface{}]uint8)
				}
				yym143 := z.DecBinary()
				_ = yym143
				if false {
				} else {
					z.F.DecMapIntfUint8X(x.FptrMapIntfUint8, d)
				}
			}
		case "FMapIntfUint16":
			if r.TryDecodeAsNil() {
				x.FMapIntfUint16 = nil
			} else {
				yyv144 := &x.FMapIntfUint16
				yym145 := z.DecBinary()
				_ = yym145
				if false {
				} else {
					z.F.DecMapIntfUint16X(yyv144, d)
				}
			}
		case "FptrMapIntfUint16":
			if x.FptrMapIntfUint16 == nil {
				x.FptrMapIntfUint16 = new(map[interface{}]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfUint16 != nil {
					x.FptrMapIntfUint16 = nil
				}
			} else {
				if x.FptrMapIntfUint16 == nil {
					x.FptrMapIntfUint16 = new(map[interface{}]uint16)
				}
				yym147 := z.DecBinary()
				_ = yym147
				if false {
				} else {
					z.F.DecMapIntfUint16X(x.FptrMapIntfUint16, d)
				}
			}
		case "FMapIntfUint32":
			if r.TryDecodeAsNil() {
				x.FMapIntfUint32 = nil
			} else {
				yyv148 := &x.FMapIntfUint32
				yym149 := z.DecBinary()
				_ = yym149
				if false {
				} else {
					z.F.DecMapIntfUint32X(yyv148, d)
				}
			}
		case "FptrMapIntfUint32":
			if x.FptrMapIntfUint32 == nil {
				x.FptrMapIntfUint32 = new(map[interface{}]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfUint32 != nil {
					x.FptrMapIntfUint32 = nil
				}
			} else {
				if x.FptrMapIntfUint32 == nil {
					x.FptrMapIntfUint32 = new(map[interface{}]uint32)
				}
				yym151 := z.DecBinary()
				_ = yym151
				if false {
				} else {
					z.F.DecMapIntfUint32X(x.FptrMapIntfUint32, d)
				}
			}
		case "FMapIntfUint64":
			if r.TryDecodeAsNil() {
				x.FMapIntfUint64 = nil
			} else {
				yyv152 := &x.FMapIntfUint64
				yym153 := z.DecBinary()
				_ = yym153
				if false {
				} else {
					z.F.DecMapIntfUint64X(yyv152, d)
				}
			}
		case "FptrMapIntfUint64":
			if x.FptrMapIntfUint64 == nil {
				x.FptrMapIntfUint64 = new(map[interface{}]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfUint64 != nil {
					x.FptrMapIntfUint64 = nil
				}
			} else {
				if x.FptrMapIntfUint64 == nil {
					x.FptrMapIntfUint64 = new(map[interface{}]uint64)
				}
				yym155 := z.DecBinary()
				_ = yym155
				if false {
				} else {
					z.F.DecMapIntfUint64X(x.FptrMapIntfUint64, d)
				}
			}
		case "FMapIntfUintptr":
			if r.TryDecodeAsNil() {
				x.FMapIntfUintptr = nil
			} else {
				yyv156 := &x.FMapIntfUintptr
				yym157 := z.DecBinary()
				_ = yym157
				if false {
				} else {
					z.F.DecMapIntfUintptrX(yyv156, d)
				}
			}
		case "FptrMapIntfUintptr":
			if x.FptrMapIntfUintptr == nil {
				x.FptrMapIntfUintptr = new(map[interface{}]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfUintptr != nil {
					x.FptrMapIntfUintptr = nil
				}
			} else {
				if x.FptrMapIntfUintptr == nil {
					x.FptrMapIntfUintptr = new(map[interface{}]uintptr)
				}
				yym159 := z.DecBinary()
				_ = yym159
				if false {
				} else {
					z.F.DecMapIntfUintptrX(x.FptrMapIntfUintptr, d)
				}
			}
		case "FMapIntfInt":
			if r.TryDecodeAsNil() {
				x.FMapIntfInt = nil
			} else {
				yyv160 := &x.FMapIntfInt
				yym161 := z.DecBinary()
				_ = yym161
				if false {
				} else {
					z.F.DecMapIntfIntX(yyv160, d)
				}
			}
		case "FptrMapIntfInt":
			if x.FptrMapIntfInt == nil {
				x.FptrMapIntfInt = new(map[interface{}]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfInt != nil {
					x.FptrMapIntfInt = nil
				}
			} else {
				if x.FptrMapIntfInt == nil {
					x.FptrMapIntfInt = new(map[interface{}]int)
				}
				yym163 := z.DecBinary()
				_ = yym163
				if false {
				} else {
					z.F.DecMapIntfIntX(x.FptrMapIntfInt, d)
				}
			}
		case "FMapIntfInt8":
			if r.TryDecodeAsNil() {
				x.FMapIntfInt8 = nil
			} else {
				yyv164 := &x.FMapIntfInt8
				yym165 := z.DecBinary()
				_ = yym165
				if false {
				} else {
					z.F.DecMapIntfInt8X(yyv164, d)
				}
			}
		case "FptrMapIntfInt8":
			if x.FptrMapIntfInt8 == nil {
				x.FptrMapIntfInt8 = new(map[interface{}]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfInt8 != nil {
					x.FptrMapIntfInt8 = nil
				}
			} else {
				if x.FptrMapIntfInt8 == nil {
					x.FptrMapIntfInt8 = new(map[interface{}]int8)
				}
				yym167 := z.DecBinary()
				_ = yym167
				if false {
				} else {
					z.F.DecMapIntfInt8X(x.FptrMapIntfInt8, d)
				}
			}
		case "FMapIntfInt16":
			if r.TryDecodeAsNil() {
				x.FMapIntfInt16 = nil
			} else {
				yyv168 := &x.FMapIntfInt16
				yym169 := z.DecBinary()
				_ = yym169
				if false {
				} else {
					z.F.DecMapIntfInt16X(yyv168, d)
				}
			}
		case "FptrMapIntfInt16":
			if x.FptrMapIntfInt16 == nil {
				x.FptrMapIntfInt16 = new(map[interface{}]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfInt16 != nil {
					x.FptrMapIntfInt16 = nil
				}
			} else {
				if x.FptrMapIntfInt16 == nil {
					x.FptrMapIntfInt16 = new(map[interface{}]int16)
				}
				yym171 := z.DecBinary()
				_ = yym171
				if false {
				} else {
					z.F.DecMapIntfInt16X(x.FptrMapIntfInt16, d)
				}
			}
		case "FMapIntfInt32":
			if r.TryDecodeAsNil() {
				x.FMapIntfInt32 = nil
			} else {
				yyv172 := &x.FMapIntfInt32
				yym173 := z.DecBinary()
				_ = yym173
				if false {
				} else {
					z.F.DecMapIntfInt32X(yyv172, d)
				}
			}
		case "FptrMapIntfInt32":
			if x.FptrMapIntfInt32 == nil {
				x.FptrMapIntfInt32 = new(map[interface{}]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfInt32 != nil {
					x.FptrMapIntfInt32 = nil
				}
			} else {
				if x.FptrMapIntfInt32 == nil {
					x.FptrMapIntfInt32 = new(map[interface{}]int32)
				}
				yym175 := z.DecBinary()
				_ = yym175
				if false {
				} else {
					z.F.DecMapIntfInt32X(x.FptrMapIntfInt32, d)
				}
			}
		case "FMapIntfInt64":
			if r.TryDecodeAsNil() {
				x.FMapIntfInt64 = nil
			} else {
				yyv176 := &x.FMapIntfInt64
				yym177 := z.DecBinary()
				_ = yym177
				if false {
				} else {
					z.F.DecMapIntfInt64X(yyv176, d)
				}
			}
		case "FptrMapIntfInt64":
			if x.FptrMapIntfInt64 == nil {
				x.FptrMapIntfInt64 = new(map[interface{}]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfInt64 != nil {
					x.FptrMapIntfInt64 = nil
				}
			} else {
				if x.FptrMapIntfInt64 == nil {
					x.FptrMapIntfInt64 = new(map[interface{}]int64)
				}
				yym179 := z.DecBinary()
				_ = yym179
				if false {
				} else {
					z.F.DecMapIntfInt64X(x.FptrMapIntfInt64, d)
				}
			}
		case "FMapIntfFloat32":
			if r.TryDecodeAsNil() {
				x.FMapIntfFloat32 = nil
			} else {
				yyv180 := &x.FMapIntfFloat32
				yym181 := z.DecBinary()
				_ = yym181
				if false {
				} else {
					z.F.DecMapIntfFloat32X(yyv180, d)
				}
			}
		case "FptrMapIntfFloat32":
			if x.FptrMapIntfFloat32 == nil {
				x.FptrMapIntfFloat32 = new(map[interface{}]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfFloat32 != nil {
					x.FptrMapIntfFloat32 = nil
				}
			} else {
				if x.FptrMapIntfFloat32 == nil {
					x.FptrMapIntfFloat32 = new(map[interface{}]float32)
				}
				yym183 := z.DecBinary()
				_ = yym183
				if false {
				} else {
					z.F.DecMapIntfFloat32X(x.FptrMapIntfFloat32, d)
				}
			}
		case "FMapIntfFloat64":
			if r.TryDecodeAsNil() {
				x.FMapIntfFloat64 = nil
			} else {
				yyv184 := &x.FMapIntfFloat64
				yym185 := z.DecBinary()
				_ = yym185
				if false {
				} else {
					z.F.DecMapIntfFloat64X(yyv184, d)
				}
			}
		case "FptrMapIntfFloat64":
			if x.FptrMapIntfFloat64 == nil {
				x.FptrMapIntfFloat64 = new(map[interface{}]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfFloat64 != nil {
					x.FptrMapIntfFloat64 = nil
				}
			} else {
				if x.FptrMapIntfFloat64 == nil {
					x.FptrMapIntfFloat64 = new(map[interface{}]float64)
				}
				yym187 := z.DecBinary()
				_ = yym187
				if false {
				} else {
					z.F.DecMapIntfFloat64X(x.FptrMapIntfFloat64, d)
				}
			}
		case "FMapIntfBool":
			if r.TryDecodeAsNil() {
				x.FMapIntfBool = nil
			} else {
				yyv188 := &x.FMapIntfBool
				yym189 := z.DecBinary()
				_ = yym189
				if false {
				} else {
					z.F.DecMapIntfBoolX(yyv188, d)
				}
			}
		case "FptrMapIntfBool":
			if x.FptrMapIntfBool == nil {
				x.FptrMapIntfBool = new(map[interface{}]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntfBool != nil {
					x.FptrMapIntfBool = nil
				}
			} else {
				if x.FptrMapIntfBool == nil {
					x.FptrMapIntfBool = new(map[interface{}]bool)
				}
				yym191 := z.DecBinary()
				_ = yym191
				if false {
				} else {
					z.F.DecMapIntfBoolX(x.FptrMapIntfBool, d)
				}
			}
		case "FMapStringIntf":
			if r.TryDecodeAsNil() {
				x.FMapStringIntf = nil
			} else {
				yyv192 := &x.FMapStringIntf
				yym193 := z.DecBinary()
				_ = yym193
				if false {
				} else {
					z.F.DecMapStringIntfX(yyv192, d)
				}
			}
		case "FptrMapStringIntf":
			if x.FptrMapStringIntf == nil {
				x.FptrMapStringIntf = new(map[string]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringIntf != nil {
					x.FptrMapStringIntf = nil
				}
			} else {
				if x.FptrMapStringIntf == nil {
					x.FptrMapStringIntf = new(map[string]interface{})
				}
				yym195 := z.DecBinary()
				_ = yym195
				if false {
				} else {
					z.F.DecMapStringIntfX(x.FptrMapStringIntf, d)
				}
			}
		case "FMapStringString":
			if r.TryDecodeAsNil() {
				x.FMapStringString = nil
			} else {
				yyv196 := &x.FMapStringString
				yym197 := z.DecBinary()
				_ = yym197
				if false {
				} else {
					z.F.DecMapStringStringX(yyv196, d)
				}
			}
		case "FptrMapStringString":
			if x.FptrMapStringString == nil {
				x.FptrMapStringString = new(map[string]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringString != nil {
					x.FptrMapStringString = nil
				}
			} else {
				if x.FptrMapStringString == nil {
					x.FptrMapStringString = new(map[string]string)
				}
				yym199 := z.DecBinary()
				_ = yym199
				if false {
				} else {
					z.F.DecMapStringStringX(x.FptrMapStringString, d)
				}
			}
		case "FMapStringUint":
			if r.TryDecodeAsNil() {
				x.FMapStringUint = nil
			} else {
				yyv200 := &x.FMapStringUint
				yym201 := z.DecBinary()
				_ = yym201
				if false {
				} else {
					z.F.DecMapStringUintX(yyv200, d)
				}
			}
		case "FptrMapStringUint":
			if x.FptrMapStringUint == nil {
				x.FptrMapStringUint = new(map[string]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringUint != nil {
					x.FptrMapStringUint = nil
				}
			} else {
				if x.FptrMapStringUint == nil {
					x.FptrMapStringUint = new(map[string]uint)
				}
				yym203 := z.DecBinary()
				_ = yym203
				if false {
				} else {
					z.F.DecMapStringUintX(x.FptrMapStringUint, d)
				}
			}
		case "FMapStringUint8":
			if r.TryDecodeAsNil() {
				x.FMapStringUint8 = nil
			} else {
				yyv204 := &x.FMapStringUint8
				yym205 := z.DecBinary()
				_ = yym205
				if false {
				} else {
					z.F.DecMapStringUint8X(yyv204, d)
				}
			}
		case "FptrMapStringUint8":
			if x.FptrMapStringUint8 == nil {
				x.FptrMapStringUint8 = new(map[string]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringUint8 != nil {
					x.FptrMapStringUint8 = nil
				}
			} else {
				if x.FptrMapStringUint8 == nil {
					x.FptrMapStringUint8 = new(map[string]uint8)
				}
				yym207 := z.DecBinary()
				_ = yym207
				if false {
				} else {
					z.F.DecMapStringUint8X(x.FptrMapStringUint8, d)
				}
			}
		case "FMapStringUint16":
			if r.TryDecodeAsNil() {
				x.FMapStringUint16 = nil
			} else {
				yyv208 := &x.FMapStringUint16
				yym209 := z.DecBinary()
				_ = yym209
				if false {
				} else {
					z.F.DecMapStringUint16X(yyv208, d)
				}
			}
		case "FptrMapStringUint16":
			if x.FptrMapStringUint16 == nil {
				x.FptrMapStringUint16 = new(map[string]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringUint16 != nil {
					x.FptrMapStringUint16 = nil
				}
			} else {
				if x.FptrMapStringUint16 == nil {
					x.FptrMapStringUint16 = new(map[string]uint16)
				}
				yym211 := z.DecBinary()
				_ = yym211
				if false {
				} else {
					z.F.DecMapStringUint16X(x.FptrMapStringUint16, d)
				}
			}
		case "FMapStringUint32":
			if r.TryDecodeAsNil() {
				x.FMapStringUint32 = nil
			} else {
				yyv212 := &x.FMapStringUint32
				yym213 := z.DecBinary()
				_ = yym213
				if false {
				} else {
					z.F.DecMapStringUint32X(yyv212, d)
				}
			}
		case "FptrMapStringUint32":
			if x.FptrMapStringUint32 == nil {
				x.FptrMapStringUint32 = new(map[string]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringUint32 != nil {
					x.FptrMapStringUint32 = nil
				}
			} else {
				if x.FptrMapStringUint32 == nil {
					x.FptrMapStringUint32 = new(map[string]uint32)
				}
				yym215 := z.DecBinary()
				_ = yym215
				if false {
				} else {
					z.F.DecMapStringUint32X(x.FptrMapStringUint32, d)
				}
			}
		case "FMapStringUint64":
			if r.TryDecodeAsNil() {
				x.FMapStringUint64 = nil
			} else {
				yyv216 := &x.FMapStringUint64
				yym217 := z.DecBinary()
				_ = yym217
				if false {
				} else {
					z.F.DecMapStringUint64X(yyv216, d)
				}
			}
		case "FptrMapStringUint64":
			if x.FptrMapStringUint64 == nil {
				x.FptrMapStringUint64 = new(map[string]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringUint64 != nil {
					x.FptrMapStringUint64 = nil
				}
			} else {
				if x.FptrMapStringUint64 == nil {
					x.FptrMapStringUint64 = new(map[string]uint64)
				}
				yym219 := z.DecBinary()
				_ = yym219
				if false {
				} else {
					z.F.DecMapStringUint64X(x.FptrMapStringUint64, d)
				}
			}
		case "FMapStringUintptr":
			if r.TryDecodeAsNil() {
				x.FMapStringUintptr = nil
			} else {
				yyv220 := &x.FMapStringUintptr
				yym221 := z.DecBinary()
				_ = yym221
				if false {
				} else {
					z.F.DecMapStringUintptrX(yyv220, d)
				}
			}
		case "FptrMapStringUintptr":
			if x.FptrMapStringUintptr == nil {
				x.FptrMapStringUintptr = new(map[string]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringUintptr != nil {
					x.FptrMapStringUintptr = nil
				}
			} else {
				if x.FptrMapStringUintptr == nil {
					x.FptrMapStringUintptr = new(map[string]uintptr)
				}
				yym223 := z.DecBinary()
				_ = yym223
				if false {
				} else {
					z.F.DecMapStringUintptrX(x.FptrMapStringUintptr, d)
				}
			}
		case "FMapStringInt":
			if r.TryDecodeAsNil() {
				x.FMapStringInt = nil
			} else {
				yyv224 := &x.FMapStringInt
				yym225 := z.DecBinary()
				_ = yym225
				if false {
				} else {
					z.F.DecMapStringIntX(yyv224, d)
				}
			}
		case "FptrMapStringInt":
			if x.FptrMapStringInt == nil {
				x.FptrMapStringInt = new(map[string]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringInt != nil {
					x.FptrMapStringInt = nil
				}
			} else {
				if x.FptrMapStringInt == nil {
					x.FptrMapStringInt = new(map[string]int)
				}
				yym227 := z.DecBinary()
				_ = yym227
				if false {
				} else {
					z.F.DecMapStringIntX(x.FptrMapStringInt, d)
				}
			}
		case "FMapStringInt8":
			if r.TryDecodeAsNil() {
				x.FMapStringInt8 = nil
			} else {
				yyv228 := &x.FMapStringInt8
				yym229 := z.DecBinary()
				_ = yym229
				if false {
				} else {
					z.F.DecMapStringInt8X(yyv228, d)
				}
			}
		case "FptrMapStringInt8":
			if x.FptrMapStringInt8 == nil {
				x.FptrMapStringInt8 = new(map[string]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringInt8 != nil {
					x.FptrMapStringInt8 = nil
				}
			} else {
				if x.FptrMapStringInt8 == nil {
					x.FptrMapStringInt8 = new(map[string]int8)
				}
				yym231 := z.DecBinary()
				_ = yym231
				if false {
				} else {
					z.F.DecMapStringInt8X(x.FptrMapStringInt8, d)
				}
			}
		case "FMapStringInt16":
			if r.TryDecodeAsNil() {
				x.FMapStringInt16 = nil
			} else {
				yyv232 := &x.FMapStringInt16
				yym233 := z.DecBinary()
				_ = yym233
				if false {
				} else {
					z.F.DecMapStringInt16X(yyv232, d)
				}
			}
		case "FptrMapStringInt16":
			if x.FptrMapStringInt16 == nil {
				x.FptrMapStringInt16 = new(map[string]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringInt16 != nil {
					x.FptrMapStringInt16 = nil
				}
			} else {
				if x.FptrMapStringInt16 == nil {
					x.FptrMapStringInt16 = new(map[string]int16)
				}
				yym235 := z.DecBinary()
				_ = yym235
				if false {
				} else {
					z.F.DecMapStringInt16X(x.FptrMapStringInt16, d)
				}
			}
		case "FMapStringInt32":
			if r.TryDecodeAsNil() {
				x.FMapStringInt32 = nil
			} else {
				yyv236 := &x.FMapStringInt32
				yym237 := z.DecBinary()
				_ = yym237
				if false {
				} else {
					z.F.DecMapStringInt32X(yyv236, d)
				}
			}
		case "FptrMapStringInt32":
			if x.FptrMapStringInt32 == nil {
				x.FptrMapStringInt32 = new(map[string]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringInt32 != nil {
					x.FptrMapStringInt32 = nil
				}
			} else {
				if x.FptrMapStringInt32 == nil {
					x.FptrMapStringInt32 = new(map[string]int32)
				}
				yym239 := z.DecBinary()
				_ = yym239
				if false {
				} else {
					z.F.DecMapStringInt32X(x.FptrMapStringInt32, d)
				}
			}
		case "FMapStringInt64":
			if r.TryDecodeAsNil() {
				x.FMapStringInt64 = nil
			} else {
				yyv240 := &x.FMapStringInt64
				yym241 := z.DecBinary()
				_ = yym241
				if false {
				} else {
					z.F.DecMapStringInt64X(yyv240, d)
				}
			}
		case "FptrMapStringInt64":
			if x.FptrMapStringInt64 == nil {
				x.FptrMapStringInt64 = new(map[string]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringInt64 != nil {
					x.FptrMapStringInt64 = nil
				}
			} else {
				if x.FptrMapStringInt64 == nil {
					x.FptrMapStringInt64 = new(map[string]int64)
				}
				yym243 := z.DecBinary()
				_ = yym243
				if false {
				} else {
					z.F.DecMapStringInt64X(x.FptrMapStringInt64, d)
				}
			}
		case "FMapStringFloat32":
			if r.TryDecodeAsNil() {
				x.FMapStringFloat32 = nil
			} else {
				yyv244 := &x.FMapStringFloat32
				yym245 := z.DecBinary()
				_ = yym245
				if false {
				} else {
					z.F.DecMapStringFloat32X(yyv244, d)
				}
			}
		case "FptrMapStringFloat32":
			if x.FptrMapStringFloat32 == nil {
				x.FptrMapStringFloat32 = new(map[string]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringFloat32 != nil {
					x.FptrMapStringFloat32 = nil
				}
			} else {
				if x.FptrMapStringFloat32 == nil {
					x.FptrMapStringFloat32 = new(map[string]float32)
				}
				yym247 := z.DecBinary()
				_ = yym247
				if false {
				} else {
					z.F.DecMapStringFloat32X(x.FptrMapStringFloat32, d)
				}
			}
		case "FMapStringFloat64":
			if r.TryDecodeAsNil() {
				x.FMapStringFloat64 = nil
			} else {
				yyv248 := &x.FMapStringFloat64
				yym249 := z.DecBinary()
				_ = yym249
				if false {
				} else {
					z.F.DecMapStringFloat64X(yyv248, d)
				}
			}
		case "FptrMapStringFloat64":
			if x.FptrMapStringFloat64 == nil {
				x.FptrMapStringFloat64 = new(map[string]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringFloat64 != nil {
					x.FptrMapStringFloat64 = nil
				}
			} else {
				if x.FptrMapStringFloat64 == nil {
					x.FptrMapStringFloat64 = new(map[string]float64)
				}
				yym251 := z.DecBinary()
				_ = yym251
				if false {
				} else {
					z.F.DecMapStringFloat64X(x.FptrMapStringFloat64, d)
				}
			}
		case "FMapStringBool":
			if r.TryDecodeAsNil() {
				x.FMapStringBool = nil
			} else {
				yyv252 := &x.FMapStringBool
				yym253 := z.DecBinary()
				_ = yym253
				if false {
				} else {
					z.F.DecMapStringBoolX(yyv252, d)
				}
			}
		case "FptrMapStringBool":
			if x.FptrMapStringBool == nil {
				x.FptrMapStringBool = new(map[string]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapStringBool != nil {
					x.FptrMapStringBool = nil
				}
			} else {
				if x.FptrMapStringBool == nil {
					x.FptrMapStringBool = new(map[string]bool)
				}
				yym255 := z.DecBinary()
				_ = yym255
				if false {
				} else {
					z.F.DecMapStringBoolX(x.FptrMapStringBool, d)
				}
			}
		case "FMapFloat32Intf":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Intf = nil
			} else {
				yyv256 := &x.FMapFloat32Intf
				yym257 := z.DecBinary()
				_ = yym257
				if false {
				} else {
					z.F.DecMapFloat32IntfX(yyv256, d)
				}
			}
		case "FptrMapFloat32Intf":
			if x.FptrMapFloat32Intf == nil {
				x.FptrMapFloat32Intf = new(map[float32]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Intf != nil {
					x.FptrMapFloat32Intf = nil
				}
			} else {
				if x.FptrMapFloat32Intf == nil {
					x.FptrMapFloat32Intf = new(map[float32]interface{})
				}
				yym259 := z.DecBinary()
				_ = yym259
				if false {
				} else {
					z.F.DecMapFloat32IntfX(x.FptrMapFloat32Intf, d)
				}
			}
		case "FMapFloat32String":
			if r.TryDecodeAsNil() {
				x.FMapFloat32String = nil
			} else {
				yyv260 := &x.FMapFloat32String
				yym261 := z.DecBinary()
				_ = yym261
				if false {
				} else {
					z.F.DecMapFloat32StringX(yyv260, d)
				}
			}
		case "FptrMapFloat32String":
			if x.FptrMapFloat32String == nil {
				x.FptrMapFloat32String = new(map[float32]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32String != nil {
					x.FptrMapFloat32String = nil
				}
			} else {
				if x.FptrMapFloat32String == nil {
					x.FptrMapFloat32String = new(map[float32]string)
				}
				yym263 := z.DecBinary()
				_ = yym263
				if false {
				} else {
					z.F.DecMapFloat32StringX(x.FptrMapFloat32String, d)
				}
			}
		case "FMapFloat32Uint":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Uint = nil
			} else {
				yyv264 := &x.FMapFloat32Uint
				yym265 := z.DecBinary()
				_ = yym265
				if false {
				} else {
					z.F.DecMapFloat32UintX(yyv264, d)
				}
			}
		case "FptrMapFloat32Uint":
			if x.FptrMapFloat32Uint == nil {
				x.FptrMapFloat32Uint = new(map[float32]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Uint != nil {
					x.FptrMapFloat32Uint = nil
				}
			} else {
				if x.FptrMapFloat32Uint == nil {
					x.FptrMapFloat32Uint = new(map[float32]uint)
				}
				yym267 := z.DecBinary()
				_ = yym267
				if false {
				} else {
					z.F.DecMapFloat32UintX(x.FptrMapFloat32Uint, d)
				}
			}
		case "FMapFloat32Uint8":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Uint8 = nil
			} else {
				yyv268 := &x.FMapFloat32Uint8
				yym269 := z.DecBinary()
				_ = yym269
				if false {
				} else {
					z.F.DecMapFloat32Uint8X(yyv268, d)
				}
			}
		case "FptrMapFloat32Uint8":
			if x.FptrMapFloat32Uint8 == nil {
				x.FptrMapFloat32Uint8 = new(map[float32]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Uint8 != nil {
					x.FptrMapFloat32Uint8 = nil
				}
			} else {
				if x.FptrMapFloat32Uint8 == nil {
					x.FptrMapFloat32Uint8 = new(map[float32]uint8)
				}
				yym271 := z.DecBinary()
				_ = yym271
				if false {
				} else {
					z.F.DecMapFloat32Uint8X(x.FptrMapFloat32Uint8, d)
				}
			}
		case "FMapFloat32Uint16":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Uint16 = nil
			} else {
				yyv272 := &x.FMapFloat32Uint16
				yym273 := z.DecBinary()
				_ = yym273
				if false {
				} else {
					z.F.DecMapFloat32Uint16X(yyv272, d)
				}
			}
		case "FptrMapFloat32Uint16":
			if x.FptrMapFloat32Uint16 == nil {
				x.FptrMapFloat32Uint16 = new(map[float32]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Uint16 != nil {
					x.FptrMapFloat32Uint16 = nil
				}
			} else {
				if x.FptrMapFloat32Uint16 == nil {
					x.FptrMapFloat32Uint16 = new(map[float32]uint16)
				}
				yym275 := z.DecBinary()
				_ = yym275
				if false {
				} else {
					z.F.DecMapFloat32Uint16X(x.FptrMapFloat32Uint16, d)
				}
			}
		case "FMapFloat32Uint32":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Uint32 = nil
			} else {
				yyv276 := &x.FMapFloat32Uint32
				yym277 := z.DecBinary()
				_ = yym277
				if false {
				} else {
					z.F.DecMapFloat32Uint32X(yyv276, d)
				}
			}
		case "FptrMapFloat32Uint32":
			if x.FptrMapFloat32Uint32 == nil {
				x.FptrMapFloat32Uint32 = new(map[float32]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Uint32 != nil {
					x.FptrMapFloat32Uint32 = nil
				}
			} else {
				if x.FptrMapFloat32Uint32 == nil {
					x.FptrMapFloat32Uint32 = new(map[float32]uint32)
				}
				yym279 := z.DecBinary()
				_ = yym279
				if false {
				} else {
					z.F.DecMapFloat32Uint32X(x.FptrMapFloat32Uint32, d)
				}
			}
		case "FMapFloat32Uint64":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Uint64 = nil
			} else {
				yyv280 := &x.FMapFloat32Uint64
				yym281 := z.DecBinary()
				_ = yym281
				if false {
				} else {
					z.F.DecMapFloat32Uint64X(yyv280, d)
				}
			}
		case "FptrMapFloat32Uint64":
			if x.FptrMapFloat32Uint64 == nil {
				x.FptrMapFloat32Uint64 = new(map[float32]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Uint64 != nil {
					x.FptrMapFloat32Uint64 = nil
				}
			} else {
				if x.FptrMapFloat32Uint64 == nil {
					x.FptrMapFloat32Uint64 = new(map[float32]uint64)
				}
				yym283 := z.DecBinary()
				_ = yym283
				if false {
				} else {
					z.F.DecMapFloat32Uint64X(x.FptrMapFloat32Uint64, d)
				}
			}
		case "FMapFloat32Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Uintptr = nil
			} else {
				yyv284 := &x.FMapFloat32Uintptr
				yym285 := z.DecBinary()
				_ = yym285
				if false {
				} else {
					z.F.DecMapFloat32UintptrX(yyv284, d)
				}
			}
		case "FptrMapFloat32Uintptr":
			if x.FptrMapFloat32Uintptr == nil {
				x.FptrMapFloat32Uintptr = new(map[float32]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Uintptr != nil {
					x.FptrMapFloat32Uintptr = nil
				}
			} else {
				if x.FptrMapFloat32Uintptr == nil {
					x.FptrMapFloat32Uintptr = new(map[float32]uintptr)
				}
				yym287 := z.DecBinary()
				_ = yym287
				if false {
				} else {
					z.F.DecMapFloat32UintptrX(x.FptrMapFloat32Uintptr, d)
				}
			}
		case "FMapFloat32Int":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Int = nil
			} else {
				yyv288 := &x.FMapFloat32Int
				yym289 := z.DecBinary()
				_ = yym289
				if false {
				} else {
					z.F.DecMapFloat32IntX(yyv288, d)
				}
			}
		case "FptrMapFloat32Int":
			if x.FptrMapFloat32Int == nil {
				x.FptrMapFloat32Int = new(map[float32]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Int != nil {
					x.FptrMapFloat32Int = nil
				}
			} else {
				if x.FptrMapFloat32Int == nil {
					x.FptrMapFloat32Int = new(map[float32]int)
				}
				yym291 := z.DecBinary()
				_ = yym291
				if false {
				} else {
					z.F.DecMapFloat32IntX(x.FptrMapFloat32Int, d)
				}
			}
		case "FMapFloat32Int8":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Int8 = nil
			} else {
				yyv292 := &x.FMapFloat32Int8
				yym293 := z.DecBinary()
				_ = yym293
				if false {
				} else {
					z.F.DecMapFloat32Int8X(yyv292, d)
				}
			}
		case "FptrMapFloat32Int8":
			if x.FptrMapFloat32Int8 == nil {
				x.FptrMapFloat32Int8 = new(map[float32]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Int8 != nil {
					x.FptrMapFloat32Int8 = nil
				}
			} else {
				if x.FptrMapFloat32Int8 == nil {
					x.FptrMapFloat32Int8 = new(map[float32]int8)
				}
				yym295 := z.DecBinary()
				_ = yym295
				if false {
				} else {
					z.F.DecMapFloat32Int8X(x.FptrMapFloat32Int8, d)
				}
			}
		case "FMapFloat32Int16":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Int16 = nil
			} else {
				yyv296 := &x.FMapFloat32Int16
				yym297 := z.DecBinary()
				_ = yym297
				if false {
				} else {
					z.F.DecMapFloat32Int16X(yyv296, d)
				}
			}
		case "FptrMapFloat32Int16":
			if x.FptrMapFloat32Int16 == nil {
				x.FptrMapFloat32Int16 = new(map[float32]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Int16 != nil {
					x.FptrMapFloat32Int16 = nil
				}
			} else {
				if x.FptrMapFloat32Int16 == nil {
					x.FptrMapFloat32Int16 = new(map[float32]int16)
				}
				yym299 := z.DecBinary()
				_ = yym299
				if false {
				} else {
					z.F.DecMapFloat32Int16X(x.FptrMapFloat32Int16, d)
				}
			}
		case "FMapFloat32Int32":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Int32 = nil
			} else {
				yyv300 := &x.FMapFloat32Int32
				yym301 := z.DecBinary()
				_ = yym301
				if false {
				} else {
					z.F.DecMapFloat32Int32X(yyv300, d)
				}
			}
		case "FptrMapFloat32Int32":
			if x.FptrMapFloat32Int32 == nil {
				x.FptrMapFloat32Int32 = new(map[float32]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Int32 != nil {
					x.FptrMapFloat32Int32 = nil
				}
			} else {
				if x.FptrMapFloat32Int32 == nil {
					x.FptrMapFloat32Int32 = new(map[float32]int32)
				}
				yym303 := z.DecBinary()
				_ = yym303
				if false {
				} else {
					z.F.DecMapFloat32Int32X(x.FptrMapFloat32Int32, d)
				}
			}
		case "FMapFloat32Int64":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Int64 = nil
			} else {
				yyv304 := &x.FMapFloat32Int64
				yym305 := z.DecBinary()
				_ = yym305
				if false {
				} else {
					z.F.DecMapFloat32Int64X(yyv304, d)
				}
			}
		case "FptrMapFloat32Int64":
			if x.FptrMapFloat32Int64 == nil {
				x.FptrMapFloat32Int64 = new(map[float32]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Int64 != nil {
					x.FptrMapFloat32Int64 = nil
				}
			} else {
				if x.FptrMapFloat32Int64 == nil {
					x.FptrMapFloat32Int64 = new(map[float32]int64)
				}
				yym307 := z.DecBinary()
				_ = yym307
				if false {
				} else {
					z.F.DecMapFloat32Int64X(x.FptrMapFloat32Int64, d)
				}
			}
		case "FMapFloat32Float32":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Float32 = nil
			} else {
				yyv308 := &x.FMapFloat32Float32
				yym309 := z.DecBinary()
				_ = yym309
				if false {
				} else {
					z.F.DecMapFloat32Float32X(yyv308, d)
				}
			}
		case "FptrMapFloat32Float32":
			if x.FptrMapFloat32Float32 == nil {
				x.FptrMapFloat32Float32 = new(map[float32]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Float32 != nil {
					x.FptrMapFloat32Float32 = nil
				}
			} else {
				if x.FptrMapFloat32Float32 == nil {
					x.FptrMapFloat32Float32 = new(map[float32]float32)
				}
				yym311 := z.DecBinary()
				_ = yym311
				if false {
				} else {
					z.F.DecMapFloat32Float32X(x.FptrMapFloat32Float32, d)
				}
			}
		case "FMapFloat32Float64":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Float64 = nil
			} else {
				yyv312 := &x.FMapFloat32Float64
				yym313 := z.DecBinary()
				_ = yym313
				if false {
				} else {
					z.F.DecMapFloat32Float64X(yyv312, d)
				}
			}
		case "FptrMapFloat32Float64":
			if x.FptrMapFloat32Float64 == nil {
				x.FptrMapFloat32Float64 = new(map[float32]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Float64 != nil {
					x.FptrMapFloat32Float64 = nil
				}
			} else {
				if x.FptrMapFloat32Float64 == nil {
					x.FptrMapFloat32Float64 = new(map[float32]float64)
				}
				yym315 := z.DecBinary()
				_ = yym315
				if false {
				} else {
					z.F.DecMapFloat32Float64X(x.FptrMapFloat32Float64, d)
				}
			}
		case "FMapFloat32Bool":
			if r.TryDecodeAsNil() {
				x.FMapFloat32Bool = nil
			} else {
				yyv316 := &x.FMapFloat32Bool
				yym317 := z.DecBinary()
				_ = yym317
				if false {
				} else {
					z.F.DecMapFloat32BoolX(yyv316, d)
				}
			}
		case "FptrMapFloat32Bool":
			if x.FptrMapFloat32Bool == nil {
				x.FptrMapFloat32Bool = new(map[float32]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat32Bool != nil {
					x.FptrMapFloat32Bool = nil
				}
			} else {
				if x.FptrMapFloat32Bool == nil {
					x.FptrMapFloat32Bool = new(map[float32]bool)
				}
				yym319 := z.DecBinary()
				_ = yym319
				if false {
				} else {
					z.F.DecMapFloat32BoolX(x.FptrMapFloat32Bool, d)
				}
			}
		case "FMapFloat64Intf":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Intf = nil
			} else {
				yyv320 := &x.FMapFloat64Intf
				yym321 := z.DecBinary()
				_ = yym321
				if false {
				} else {
					z.F.DecMapFloat64IntfX(yyv320, d)
				}
			}
		case "FptrMapFloat64Intf":
			if x.FptrMapFloat64Intf == nil {
				x.FptrMapFloat64Intf = new(map[float64]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Intf != nil {
					x.FptrMapFloat64Intf = nil
				}
			} else {
				if x.FptrMapFloat64Intf == nil {
					x.FptrMapFloat64Intf = new(map[float64]interface{})
				}
				yym323 := z.DecBinary()
				_ = yym323
				if false {
				} else {
					z.F.DecMapFloat64IntfX(x.FptrMapFloat64Intf, d)
				}
			}
		case "FMapFloat64String":
			if r.TryDecodeAsNil() {
				x.FMapFloat64String = nil
			} else {
				yyv324 := &x.FMapFloat64String
				yym325 := z.DecBinary()
				_ = yym325
				if false {
				} else {
					z.F.DecMapFloat64StringX(yyv324, d)
				}
			}
		case "FptrMapFloat64String":
			if x.FptrMapFloat64String == nil {
				x.FptrMapFloat64String = new(map[float64]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64String != nil {
					x.FptrMapFloat64String = nil
				}
			} else {
				if x.FptrMapFloat64String == nil {
					x.FptrMapFloat64String = new(map[float64]string)
				}
				yym327 := z.DecBinary()
				_ = yym327
				if false {
				} else {
					z.F.DecMapFloat64StringX(x.FptrMapFloat64String, d)
				}
			}
		case "FMapFloat64Uint":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Uint = nil
			} else {
				yyv328 := &x.FMapFloat64Uint
				yym329 := z.DecBinary()
				_ = yym329
				if false {
				} else {
					z.F.DecMapFloat64UintX(yyv328, d)
				}
			}
		case "FptrMapFloat64Uint":
			if x.FptrMapFloat64Uint == nil {
				x.FptrMapFloat64Uint = new(map[float64]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Uint != nil {
					x.FptrMapFloat64Uint = nil
				}
			} else {
				if x.FptrMapFloat64Uint == nil {
					x.FptrMapFloat64Uint = new(map[float64]uint)
				}
				yym331 := z.DecBinary()
				_ = yym331
				if false {
				} else {
					z.F.DecMapFloat64UintX(x.FptrMapFloat64Uint, d)
				}
			}
		case "FMapFloat64Uint8":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Uint8 = nil
			} else {
				yyv332 := &x.FMapFloat64Uint8
				yym333 := z.DecBinary()
				_ = yym333
				if false {
				} else {
					z.F.DecMapFloat64Uint8X(yyv332, d)
				}
			}
		case "FptrMapFloat64Uint8":
			if x.FptrMapFloat64Uint8 == nil {
				x.FptrMapFloat64Uint8 = new(map[float64]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Uint8 != nil {
					x.FptrMapFloat64Uint8 = nil
				}
			} else {
				if x.FptrMapFloat64Uint8 == nil {
					x.FptrMapFloat64Uint8 = new(map[float64]uint8)
				}
				yym335 := z.DecBinary()
				_ = yym335
				if false {
				} else {
					z.F.DecMapFloat64Uint8X(x.FptrMapFloat64Uint8, d)
				}
			}
		case "FMapFloat64Uint16":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Uint16 = nil
			} else {
				yyv336 := &x.FMapFloat64Uint16
				yym337 := z.DecBinary()
				_ = yym337
				if false {
				} else {
					z.F.DecMapFloat64Uint16X(yyv336, d)
				}
			}
		case "FptrMapFloat64Uint16":
			if x.FptrMapFloat64Uint16 == nil {
				x.FptrMapFloat64Uint16 = new(map[float64]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Uint16 != nil {
					x.FptrMapFloat64Uint16 = nil
				}
			} else {
				if x.FptrMapFloat64Uint16 == nil {
					x.FptrMapFloat64Uint16 = new(map[float64]uint16)
				}
				yym339 := z.DecBinary()
				_ = yym339
				if false {
				} else {
					z.F.DecMapFloat64Uint16X(x.FptrMapFloat64Uint16, d)
				}
			}
		case "FMapFloat64Uint32":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Uint32 = nil
			} else {
				yyv340 := &x.FMapFloat64Uint32
				yym341 := z.DecBinary()
				_ = yym341
				if false {
				} else {
					z.F.DecMapFloat64Uint32X(yyv340, d)
				}
			}
		case "FptrMapFloat64Uint32":
			if x.FptrMapFloat64Uint32 == nil {
				x.FptrMapFloat64Uint32 = new(map[float64]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Uint32 != nil {
					x.FptrMapFloat64Uint32 = nil
				}
			} else {
				if x.FptrMapFloat64Uint32 == nil {
					x.FptrMapFloat64Uint32 = new(map[float64]uint32)
				}
				yym343 := z.DecBinary()
				_ = yym343
				if false {
				} else {
					z.F.DecMapFloat64Uint32X(x.FptrMapFloat64Uint32, d)
				}
			}
		case "FMapFloat64Uint64":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Uint64 = nil
			} else {
				yyv344 := &x.FMapFloat64Uint64
				yym345 := z.DecBinary()
				_ = yym345
				if false {
				} else {
					z.F.DecMapFloat64Uint64X(yyv344, d)
				}
			}
		case "FptrMapFloat64Uint64":
			if x.FptrMapFloat64Uint64 == nil {
				x.FptrMapFloat64Uint64 = new(map[float64]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Uint64 != nil {
					x.FptrMapFloat64Uint64 = nil
				}
			} else {
				if x.FptrMapFloat64Uint64 == nil {
					x.FptrMapFloat64Uint64 = new(map[float64]uint64)
				}
				yym347 := z.DecBinary()
				_ = yym347
				if false {
				} else {
					z.F.DecMapFloat64Uint64X(x.FptrMapFloat64Uint64, d)
				}
			}
		case "FMapFloat64Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Uintptr = nil
			} else {
				yyv348 := &x.FMapFloat64Uintptr
				yym349 := z.DecBinary()
				_ = yym349
				if false {
				} else {
					z.F.DecMapFloat64UintptrX(yyv348, d)
				}
			}
		case "FptrMapFloat64Uintptr":
			if x.FptrMapFloat64Uintptr == nil {
				x.FptrMapFloat64Uintptr = new(map[float64]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Uintptr != nil {
					x.FptrMapFloat64Uintptr = nil
				}
			} else {
				if x.FptrMapFloat64Uintptr == nil {
					x.FptrMapFloat64Uintptr = new(map[float64]uintptr)
				}
				yym351 := z.DecBinary()
				_ = yym351
				if false {
				} else {
					z.F.DecMapFloat64UintptrX(x.FptrMapFloat64Uintptr, d)
				}
			}
		case "FMapFloat64Int":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Int = nil
			} else {
				yyv352 := &x.FMapFloat64Int
				yym353 := z.DecBinary()
				_ = yym353
				if false {
				} else {
					z.F.DecMapFloat64IntX(yyv352, d)
				}
			}
		case "FptrMapFloat64Int":
			if x.FptrMapFloat64Int == nil {
				x.FptrMapFloat64Int = new(map[float64]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Int != nil {
					x.FptrMapFloat64Int = nil
				}
			} else {
				if x.FptrMapFloat64Int == nil {
					x.FptrMapFloat64Int = new(map[float64]int)
				}
				yym355 := z.DecBinary()
				_ = yym355
				if false {
				} else {
					z.F.DecMapFloat64IntX(x.FptrMapFloat64Int, d)
				}
			}
		case "FMapFloat64Int8":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Int8 = nil
			} else {
				yyv356 := &x.FMapFloat64Int8
				yym357 := z.DecBinary()
				_ = yym357
				if false {
				} else {
					z.F.DecMapFloat64Int8X(yyv356, d)
				}
			}
		case "FptrMapFloat64Int8":
			if x.FptrMapFloat64Int8 == nil {
				x.FptrMapFloat64Int8 = new(map[float64]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Int8 != nil {
					x.FptrMapFloat64Int8 = nil
				}
			} else {
				if x.FptrMapFloat64Int8 == nil {
					x.FptrMapFloat64Int8 = new(map[float64]int8)
				}
				yym359 := z.DecBinary()
				_ = yym359
				if false {
				} else {
					z.F.DecMapFloat64Int8X(x.FptrMapFloat64Int8, d)
				}
			}
		case "FMapFloat64Int16":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Int16 = nil
			} else {
				yyv360 := &x.FMapFloat64Int16
				yym361 := z.DecBinary()
				_ = yym361
				if false {
				} else {
					z.F.DecMapFloat64Int16X(yyv360, d)
				}
			}
		case "FptrMapFloat64Int16":
			if x.FptrMapFloat64Int16 == nil {
				x.FptrMapFloat64Int16 = new(map[float64]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Int16 != nil {
					x.FptrMapFloat64Int16 = nil
				}
			} else {
				if x.FptrMapFloat64Int16 == nil {
					x.FptrMapFloat64Int16 = new(map[float64]int16)
				}
				yym363 := z.DecBinary()
				_ = yym363
				if false {
				} else {
					z.F.DecMapFloat64Int16X(x.FptrMapFloat64Int16, d)
				}
			}
		case "FMapFloat64Int32":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Int32 = nil
			} else {
				yyv364 := &x.FMapFloat64Int32
				yym365 := z.DecBinary()
				_ = yym365
				if false {
				} else {
					z.F.DecMapFloat64Int32X(yyv364, d)
				}
			}
		case "FptrMapFloat64Int32":
			if x.FptrMapFloat64Int32 == nil {
				x.FptrMapFloat64Int32 = new(map[float64]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Int32 != nil {
					x.FptrMapFloat64Int32 = nil
				}
			} else {
				if x.FptrMapFloat64Int32 == nil {
					x.FptrMapFloat64Int32 = new(map[float64]int32)
				}
				yym367 := z.DecBinary()
				_ = yym367
				if false {
				} else {
					z.F.DecMapFloat64Int32X(x.FptrMapFloat64Int32, d)
				}
			}
		case "FMapFloat64Int64":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Int64 = nil
			} else {
				yyv368 := &x.FMapFloat64Int64
				yym369 := z.DecBinary()
				_ = yym369
				if false {
				} else {
					z.F.DecMapFloat64Int64X(yyv368, d)
				}
			}
		case "FptrMapFloat64Int64":
			if x.FptrMapFloat64Int64 == nil {
				x.FptrMapFloat64Int64 = new(map[float64]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Int64 != nil {
					x.FptrMapFloat64Int64 = nil
				}
			} else {
				if x.FptrMapFloat64Int64 == nil {
					x.FptrMapFloat64Int64 = new(map[float64]int64)
				}
				yym371 := z.DecBinary()
				_ = yym371
				if false {
				} else {
					z.F.DecMapFloat64Int64X(x.FptrMapFloat64Int64, d)
				}
			}
		case "FMapFloat64Float32":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Float32 = nil
			} else {
				yyv372 := &x.FMapFloat64Float32
				yym373 := z.DecBinary()
				_ = yym373
				if false {
				} else {
					z.F.DecMapFloat64Float32X(yyv372, d)
				}
			}
		case "FptrMapFloat64Float32":
			if x.FptrMapFloat64Float32 == nil {
				x.FptrMapFloat64Float32 = new(map[float64]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Float32 != nil {
					x.FptrMapFloat64Float32 = nil
				}
			} else {
				if x.FptrMapFloat64Float32 == nil {
					x.FptrMapFloat64Float32 = new(map[float64]float32)
				}
				yym375 := z.DecBinary()
				_ = yym375
				if false {
				} else {
					z.F.DecMapFloat64Float32X(x.FptrMapFloat64Float32, d)
				}
			}
		case "FMapFloat64Float64":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Float64 = nil
			} else {
				yyv376 := &x.FMapFloat64Float64
				yym377 := z.DecBinary()
				_ = yym377
				if false {
				} else {
					z.F.DecMapFloat64Float64X(yyv376, d)
				}
			}
		case "FptrMapFloat64Float64":
			if x.FptrMapFloat64Float64 == nil {
				x.FptrMapFloat64Float64 = new(map[float64]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Float64 != nil {
					x.FptrMapFloat64Float64 = nil
				}
			} else {
				if x.FptrMapFloat64Float64 == nil {
					x.FptrMapFloat64Float64 = new(map[float64]float64)
				}
				yym379 := z.DecBinary()
				_ = yym379
				if false {
				} else {
					z.F.DecMapFloat64Float64X(x.FptrMapFloat64Float64, d)
				}
			}
		case "FMapFloat64Bool":
			if r.TryDecodeAsNil() {
				x.FMapFloat64Bool = nil
			} else {
				yyv380 := &x.FMapFloat64Bool
				yym381 := z.DecBinary()
				_ = yym381
				if false {
				} else {
					z.F.DecMapFloat64BoolX(yyv380, d)
				}
			}
		case "FptrMapFloat64Bool":
			if x.FptrMapFloat64Bool == nil {
				x.FptrMapFloat64Bool = new(map[float64]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapFloat64Bool != nil {
					x.FptrMapFloat64Bool = nil
				}
			} else {
				if x.FptrMapFloat64Bool == nil {
					x.FptrMapFloat64Bool = new(map[float64]bool)
				}
				yym383 := z.DecBinary()
				_ = yym383
				if false {
				} else {
					z.F.DecMapFloat64BoolX(x.FptrMapFloat64Bool, d)
				}
			}
		case "FMapUintIntf":
			if r.TryDecodeAsNil() {
				x.FMapUintIntf = nil
			} else {
				yyv384 := &x.FMapUintIntf
				yym385 := z.DecBinary()
				_ = yym385
				if false {
				} else {
					z.F.DecMapUintIntfX(yyv384, d)
				}
			}
		case "FptrMapUintIntf":
			if x.FptrMapUintIntf == nil {
				x.FptrMapUintIntf = new(map[uint]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintIntf != nil {
					x.FptrMapUintIntf = nil
				}
			} else {
				if x.FptrMapUintIntf == nil {
					x.FptrMapUintIntf = new(map[uint]interface{})
				}
				yym387 := z.DecBinary()
				_ = yym387
				if false {
				} else {
					z.F.DecMapUintIntfX(x.FptrMapUintIntf, d)
				}
			}
		case "FMapUintString":
			if r.TryDecodeAsNil() {
				x.FMapUintString = nil
			} else {
				yyv388 := &x.FMapUintString
				yym389 := z.DecBinary()
				_ = yym389
				if false {
				} else {
					z.F.DecMapUintStringX(yyv388, d)
				}
			}
		case "FptrMapUintString":
			if x.FptrMapUintString == nil {
				x.FptrMapUintString = new(map[uint]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintString != nil {
					x.FptrMapUintString = nil
				}
			} else {
				if x.FptrMapUintString == nil {
					x.FptrMapUintString = new(map[uint]string)
				}
				yym391 := z.DecBinary()
				_ = yym391
				if false {
				} else {
					z.F.DecMapUintStringX(x.FptrMapUintString, d)
				}
			}
		case "FMapUintUint":
			if r.TryDecodeAsNil() {
				x.FMapUintUint = nil
			} else {
				yyv392 := &x.FMapUintUint
				yym393 := z.DecBinary()
				_ = yym393
				if false {
				} else {
					z.F.DecMapUintUintX(yyv392, d)
				}
			}
		case "FptrMapUintUint":
			if x.FptrMapUintUint == nil {
				x.FptrMapUintUint = new(map[uint]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintUint != nil {
					x.FptrMapUintUint = nil
				}
			} else {
				if x.FptrMapUintUint == nil {
					x.FptrMapUintUint = new(map[uint]uint)
				}
				yym395 := z.DecBinary()
				_ = yym395
				if false {
				} else {
					z.F.DecMapUintUintX(x.FptrMapUintUint, d)
				}
			}
		case "FMapUintUint8":
			if r.TryDecodeAsNil() {
				x.FMapUintUint8 = nil
			} else {
				yyv396 := &x.FMapUintUint8
				yym397 := z.DecBinary()
				_ = yym397
				if false {
				} else {
					z.F.DecMapUintUint8X(yyv396, d)
				}
			}
		case "FptrMapUintUint8":
			if x.FptrMapUintUint8 == nil {
				x.FptrMapUintUint8 = new(map[uint]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintUint8 != nil {
					x.FptrMapUintUint8 = nil
				}
			} else {
				if x.FptrMapUintUint8 == nil {
					x.FptrMapUintUint8 = new(map[uint]uint8)
				}
				yym399 := z.DecBinary()
				_ = yym399
				if false {
				} else {
					z.F.DecMapUintUint8X(x.FptrMapUintUint8, d)
				}
			}
		case "FMapUintUint16":
			if r.TryDecodeAsNil() {
				x.FMapUintUint16 = nil
			} else {
				yyv400 := &x.FMapUintUint16
				yym401 := z.DecBinary()
				_ = yym401
				if false {
				} else {
					z.F.DecMapUintUint16X(yyv400, d)
				}
			}
		case "FptrMapUintUint16":
			if x.FptrMapUintUint16 == nil {
				x.FptrMapUintUint16 = new(map[uint]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintUint16 != nil {
					x.FptrMapUintUint16 = nil
				}
			} else {
				if x.FptrMapUintUint16 == nil {
					x.FptrMapUintUint16 = new(map[uint]uint16)
				}
				yym403 := z.DecBinary()
				_ = yym403
				if false {
				} else {
					z.F.DecMapUintUint16X(x.FptrMapUintUint16, d)
				}
			}
		case "FMapUintUint32":
			if r.TryDecodeAsNil() {
				x.FMapUintUint32 = nil
			} else {
				yyv404 := &x.FMapUintUint32
				yym405 := z.DecBinary()
				_ = yym405
				if false {
				} else {
					z.F.DecMapUintUint32X(yyv404, d)
				}
			}
		case "FptrMapUintUint32":
			if x.FptrMapUintUint32 == nil {
				x.FptrMapUintUint32 = new(map[uint]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintUint32 != nil {
					x.FptrMapUintUint32 = nil
				}
			} else {
				if x.FptrMapUintUint32 == nil {
					x.FptrMapUintUint32 = new(map[uint]uint32)
				}
				yym407 := z.DecBinary()
				_ = yym407
				if false {
				} else {
					z.F.DecMapUintUint32X(x.FptrMapUintUint32, d)
				}
			}
		case "FMapUintUint64":
			if r.TryDecodeAsNil() {
				x.FMapUintUint64 = nil
			} else {
				yyv408 := &x.FMapUintUint64
				yym409 := z.DecBinary()
				_ = yym409
				if false {
				} else {
					z.F.DecMapUintUint64X(yyv408, d)
				}
			}
		case "FptrMapUintUint64":
			if x.FptrMapUintUint64 == nil {
				x.FptrMapUintUint64 = new(map[uint]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintUint64 != nil {
					x.FptrMapUintUint64 = nil
				}
			} else {
				if x.FptrMapUintUint64 == nil {
					x.FptrMapUintUint64 = new(map[uint]uint64)
				}
				yym411 := z.DecBinary()
				_ = yym411
				if false {
				} else {
					z.F.DecMapUintUint64X(x.FptrMapUintUint64, d)
				}
			}
		case "FMapUintUintptr":
			if r.TryDecodeAsNil() {
				x.FMapUintUintptr = nil
			} else {
				yyv412 := &x.FMapUintUintptr
				yym413 := z.DecBinary()
				_ = yym413
				if false {
				} else {
					z.F.DecMapUintUintptrX(yyv412, d)
				}
			}
		case "FptrMapUintUintptr":
			if x.FptrMapUintUintptr == nil {
				x.FptrMapUintUintptr = new(map[uint]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintUintptr != nil {
					x.FptrMapUintUintptr = nil
				}
			} else {
				if x.FptrMapUintUintptr == nil {
					x.FptrMapUintUintptr = new(map[uint]uintptr)
				}
				yym415 := z.DecBinary()
				_ = yym415
				if false {
				} else {
					z.F.DecMapUintUintptrX(x.FptrMapUintUintptr, d)
				}
			}
		case "FMapUintInt":
			if r.TryDecodeAsNil() {
				x.FMapUintInt = nil
			} else {
				yyv416 := &x.FMapUintInt
				yym417 := z.DecBinary()
				_ = yym417
				if false {
				} else {
					z.F.DecMapUintIntX(yyv416, d)
				}
			}
		case "FptrMapUintInt":
			if x.FptrMapUintInt == nil {
				x.FptrMapUintInt = new(map[uint]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintInt != nil {
					x.FptrMapUintInt = nil
				}
			} else {
				if x.FptrMapUintInt == nil {
					x.FptrMapUintInt = new(map[uint]int)
				}
				yym419 := z.DecBinary()
				_ = yym419
				if false {
				} else {
					z.F.DecMapUintIntX(x.FptrMapUintInt, d)
				}
			}
		case "FMapUintInt8":
			if r.TryDecodeAsNil() {
				x.FMapUintInt8 = nil
			} else {
				yyv420 := &x.FMapUintInt8
				yym421 := z.DecBinary()
				_ = yym421
				if false {
				} else {
					z.F.DecMapUintInt8X(yyv420, d)
				}
			}
		case "FptrMapUintInt8":
			if x.FptrMapUintInt8 == nil {
				x.FptrMapUintInt8 = new(map[uint]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintInt8 != nil {
					x.FptrMapUintInt8 = nil
				}
			} else {
				if x.FptrMapUintInt8 == nil {
					x.FptrMapUintInt8 = new(map[uint]int8)
				}
				yym423 := z.DecBinary()
				_ = yym423
				if false {
				} else {
					z.F.DecMapUintInt8X(x.FptrMapUintInt8, d)
				}
			}
		case "FMapUintInt16":
			if r.TryDecodeAsNil() {
				x.FMapUintInt16 = nil
			} else {
				yyv424 := &x.FMapUintInt16
				yym425 := z.DecBinary()
				_ = yym425
				if false {
				} else {
					z.F.DecMapUintInt16X(yyv424, d)
				}
			}
		case "FptrMapUintInt16":
			if x.FptrMapUintInt16 == nil {
				x.FptrMapUintInt16 = new(map[uint]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintInt16 != nil {
					x.FptrMapUintInt16 = nil
				}
			} else {
				if x.FptrMapUintInt16 == nil {
					x.FptrMapUintInt16 = new(map[uint]int16)
				}
				yym427 := z.DecBinary()
				_ = yym427
				if false {
				} else {
					z.F.DecMapUintInt16X(x.FptrMapUintInt16, d)
				}
			}
		case "FMapUintInt32":
			if r.TryDecodeAsNil() {
				x.FMapUintInt32 = nil
			} else {
				yyv428 := &x.FMapUintInt32
				yym429 := z.DecBinary()
				_ = yym429
				if false {
				} else {
					z.F.DecMapUintInt32X(yyv428, d)
				}
			}
		case "FptrMapUintInt32":
			if x.FptrMapUintInt32 == nil {
				x.FptrMapUintInt32 = new(map[uint]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintInt32 != nil {
					x.FptrMapUintInt32 = nil
				}
			} else {
				if x.FptrMapUintInt32 == nil {
					x.FptrMapUintInt32 = new(map[uint]int32)
				}
				yym431 := z.DecBinary()
				_ = yym431
				if false {
				} else {
					z.F.DecMapUintInt32X(x.FptrMapUintInt32, d)
				}
			}
		case "FMapUintInt64":
			if r.TryDecodeAsNil() {
				x.FMapUintInt64 = nil
			} else {
				yyv432 := &x.FMapUintInt64
				yym433 := z.DecBinary()
				_ = yym433
				if false {
				} else {
					z.F.DecMapUintInt64X(yyv432, d)
				}
			}
		case "FptrMapUintInt64":
			if x.FptrMapUintInt64 == nil {
				x.FptrMapUintInt64 = new(map[uint]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintInt64 != nil {
					x.FptrMapUintInt64 = nil
				}
			} else {
				if x.FptrMapUintInt64 == nil {
					x.FptrMapUintInt64 = new(map[uint]int64)
				}
				yym435 := z.DecBinary()
				_ = yym435
				if false {
				} else {
					z.F.DecMapUintInt64X(x.FptrMapUintInt64, d)
				}
			}
		case "FMapUintFloat32":
			if r.TryDecodeAsNil() {
				x.FMapUintFloat32 = nil
			} else {
				yyv436 := &x.FMapUintFloat32
				yym437 := z.DecBinary()
				_ = yym437
				if false {
				} else {
					z.F.DecMapUintFloat32X(yyv436, d)
				}
			}
		case "FptrMapUintFloat32":
			if x.FptrMapUintFloat32 == nil {
				x.FptrMapUintFloat32 = new(map[uint]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintFloat32 != nil {
					x.FptrMapUintFloat32 = nil
				}
			} else {
				if x.FptrMapUintFloat32 == nil {
					x.FptrMapUintFloat32 = new(map[uint]float32)
				}
				yym439 := z.DecBinary()
				_ = yym439
				if false {
				} else {
					z.F.DecMapUintFloat32X(x.FptrMapUintFloat32, d)
				}
			}
		case "FMapUintFloat64":
			if r.TryDecodeAsNil() {
				x.FMapUintFloat64 = nil
			} else {
				yyv440 := &x.FMapUintFloat64
				yym441 := z.DecBinary()
				_ = yym441
				if false {
				} else {
					z.F.DecMapUintFloat64X(yyv440, d)
				}
			}
		case "FptrMapUintFloat64":
			if x.FptrMapUintFloat64 == nil {
				x.FptrMapUintFloat64 = new(map[uint]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintFloat64 != nil {
					x.FptrMapUintFloat64 = nil
				}
			} else {
				if x.FptrMapUintFloat64 == nil {
					x.FptrMapUintFloat64 = new(map[uint]float64)
				}
				yym443 := z.DecBinary()
				_ = yym443
				if false {
				} else {
					z.F.DecMapUintFloat64X(x.FptrMapUintFloat64, d)
				}
			}
		case "FMapUintBool":
			if r.TryDecodeAsNil() {
				x.FMapUintBool = nil
			} else {
				yyv444 := &x.FMapUintBool
				yym445 := z.DecBinary()
				_ = yym445
				if false {
				} else {
					z.F.DecMapUintBoolX(yyv444, d)
				}
			}
		case "FptrMapUintBool":
			if x.FptrMapUintBool == nil {
				x.FptrMapUintBool = new(map[uint]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintBool != nil {
					x.FptrMapUintBool = nil
				}
			} else {
				if x.FptrMapUintBool == nil {
					x.FptrMapUintBool = new(map[uint]bool)
				}
				yym447 := z.DecBinary()
				_ = yym447
				if false {
				} else {
					z.F.DecMapUintBoolX(x.FptrMapUintBool, d)
				}
			}
		case "FMapUint8Intf":
			if r.TryDecodeAsNil() {
				x.FMapUint8Intf = nil
			} else {
				yyv448 := &x.FMapUint8Intf
				yym449 := z.DecBinary()
				_ = yym449
				if false {
				} else {
					z.F.DecMapUint8IntfX(yyv448, d)
				}
			}
		case "FptrMapUint8Intf":
			if x.FptrMapUint8Intf == nil {
				x.FptrMapUint8Intf = new(map[uint8]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Intf != nil {
					x.FptrMapUint8Intf = nil
				}
			} else {
				if x.FptrMapUint8Intf == nil {
					x.FptrMapUint8Intf = new(map[uint8]interface{})
				}
				yym451 := z.DecBinary()
				_ = yym451
				if false {
				} else {
					z.F.DecMapUint8IntfX(x.FptrMapUint8Intf, d)
				}
			}
		case "FMapUint8String":
			if r.TryDecodeAsNil() {
				x.FMapUint8String = nil
			} else {
				yyv452 := &x.FMapUint8String
				yym453 := z.DecBinary()
				_ = yym453
				if false {
				} else {
					z.F.DecMapUint8StringX(yyv452, d)
				}
			}
		case "FptrMapUint8String":
			if x.FptrMapUint8String == nil {
				x.FptrMapUint8String = new(map[uint8]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8String != nil {
					x.FptrMapUint8String = nil
				}
			} else {
				if x.FptrMapUint8String == nil {
					x.FptrMapUint8String = new(map[uint8]string)
				}
				yym455 := z.DecBinary()
				_ = yym455
				if false {
				} else {
					z.F.DecMapUint8StringX(x.FptrMapUint8String, d)
				}
			}
		case "FMapUint8Uint":
			if r.TryDecodeAsNil() {
				x.FMapUint8Uint = nil
			} else {
				yyv456 := &x.FMapUint8Uint
				yym457 := z.DecBinary()
				_ = yym457
				if false {
				} else {
					z.F.DecMapUint8UintX(yyv456, d)
				}
			}
		case "FptrMapUint8Uint":
			if x.FptrMapUint8Uint == nil {
				x.FptrMapUint8Uint = new(map[uint8]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Uint != nil {
					x.FptrMapUint8Uint = nil
				}
			} else {
				if x.FptrMapUint8Uint == nil {
					x.FptrMapUint8Uint = new(map[uint8]uint)
				}
				yym459 := z.DecBinary()
				_ = yym459
				if false {
				} else {
					z.F.DecMapUint8UintX(x.FptrMapUint8Uint, d)
				}
			}
		case "FMapUint8Uint8":
			if r.TryDecodeAsNil() {
				x.FMapUint8Uint8 = nil
			} else {
				yyv460 := &x.FMapUint8Uint8
				yym461 := z.DecBinary()
				_ = yym461
				if false {
				} else {
					z.F.DecMapUint8Uint8X(yyv460, d)
				}
			}
		case "FptrMapUint8Uint8":
			if x.FptrMapUint8Uint8 == nil {
				x.FptrMapUint8Uint8 = new(map[uint8]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Uint8 != nil {
					x.FptrMapUint8Uint8 = nil
				}
			} else {
				if x.FptrMapUint8Uint8 == nil {
					x.FptrMapUint8Uint8 = new(map[uint8]uint8)
				}
				yym463 := z.DecBinary()
				_ = yym463
				if false {
				} else {
					z.F.DecMapUint8Uint8X(x.FptrMapUint8Uint8, d)
				}
			}
		case "FMapUint8Uint16":
			if r.TryDecodeAsNil() {
				x.FMapUint8Uint16 = nil
			} else {
				yyv464 := &x.FMapUint8Uint16
				yym465 := z.DecBinary()
				_ = yym465
				if false {
				} else {
					z.F.DecMapUint8Uint16X(yyv464, d)
				}
			}
		case "FptrMapUint8Uint16":
			if x.FptrMapUint8Uint16 == nil {
				x.FptrMapUint8Uint16 = new(map[uint8]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Uint16 != nil {
					x.FptrMapUint8Uint16 = nil
				}
			} else {
				if x.FptrMapUint8Uint16 == nil {
					x.FptrMapUint8Uint16 = new(map[uint8]uint16)
				}
				yym467 := z.DecBinary()
				_ = yym467
				if false {
				} else {
					z.F.DecMapUint8Uint16X(x.FptrMapUint8Uint16, d)
				}
			}
		case "FMapUint8Uint32":
			if r.TryDecodeAsNil() {
				x.FMapUint8Uint32 = nil
			} else {
				yyv468 := &x.FMapUint8Uint32
				yym469 := z.DecBinary()
				_ = yym469
				if false {
				} else {
					z.F.DecMapUint8Uint32X(yyv468, d)
				}
			}
		case "FptrMapUint8Uint32":
			if x.FptrMapUint8Uint32 == nil {
				x.FptrMapUint8Uint32 = new(map[uint8]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Uint32 != nil {
					x.FptrMapUint8Uint32 = nil
				}
			} else {
				if x.FptrMapUint8Uint32 == nil {
					x.FptrMapUint8Uint32 = new(map[uint8]uint32)
				}
				yym471 := z.DecBinary()
				_ = yym471
				if false {
				} else {
					z.F.DecMapUint8Uint32X(x.FptrMapUint8Uint32, d)
				}
			}
		case "FMapUint8Uint64":
			if r.TryDecodeAsNil() {
				x.FMapUint8Uint64 = nil
			} else {
				yyv472 := &x.FMapUint8Uint64
				yym473 := z.DecBinary()
				_ = yym473
				if false {
				} else {
					z.F.DecMapUint8Uint64X(yyv472, d)
				}
			}
		case "FptrMapUint8Uint64":
			if x.FptrMapUint8Uint64 == nil {
				x.FptrMapUint8Uint64 = new(map[uint8]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Uint64 != nil {
					x.FptrMapUint8Uint64 = nil
				}
			} else {
				if x.FptrMapUint8Uint64 == nil {
					x.FptrMapUint8Uint64 = new(map[uint8]uint64)
				}
				yym475 := z.DecBinary()
				_ = yym475
				if false {
				} else {
					z.F.DecMapUint8Uint64X(x.FptrMapUint8Uint64, d)
				}
			}
		case "FMapUint8Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapUint8Uintptr = nil
			} else {
				yyv476 := &x.FMapUint8Uintptr
				yym477 := z.DecBinary()
				_ = yym477
				if false {
				} else {
					z.F.DecMapUint8UintptrX(yyv476, d)
				}
			}
		case "FptrMapUint8Uintptr":
			if x.FptrMapUint8Uintptr == nil {
				x.FptrMapUint8Uintptr = new(map[uint8]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Uintptr != nil {
					x.FptrMapUint8Uintptr = nil
				}
			} else {
				if x.FptrMapUint8Uintptr == nil {
					x.FptrMapUint8Uintptr = new(map[uint8]uintptr)
				}
				yym479 := z.DecBinary()
				_ = yym479
				if false {
				} else {
					z.F.DecMapUint8UintptrX(x.FptrMapUint8Uintptr, d)
				}
			}
		case "FMapUint8Int":
			if r.TryDecodeAsNil() {
				x.FMapUint8Int = nil
			} else {
				yyv480 := &x.FMapUint8Int
				yym481 := z.DecBinary()
				_ = yym481
				if false {
				} else {
					z.F.DecMapUint8IntX(yyv480, d)
				}
			}
		case "FptrMapUint8Int":
			if x.FptrMapUint8Int == nil {
				x.FptrMapUint8Int = new(map[uint8]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Int != nil {
					x.FptrMapUint8Int = nil
				}
			} else {
				if x.FptrMapUint8Int == nil {
					x.FptrMapUint8Int = new(map[uint8]int)
				}
				yym483 := z.DecBinary()
				_ = yym483
				if false {
				} else {
					z.F.DecMapUint8IntX(x.FptrMapUint8Int, d)
				}
			}
		case "FMapUint8Int8":
			if r.TryDecodeAsNil() {
				x.FMapUint8Int8 = nil
			} else {
				yyv484 := &x.FMapUint8Int8
				yym485 := z.DecBinary()
				_ = yym485
				if false {
				} else {
					z.F.DecMapUint8Int8X(yyv484, d)
				}
			}
		case "FptrMapUint8Int8":
			if x.FptrMapUint8Int8 == nil {
				x.FptrMapUint8Int8 = new(map[uint8]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Int8 != nil {
					x.FptrMapUint8Int8 = nil
				}
			} else {
				if x.FptrMapUint8Int8 == nil {
					x.FptrMapUint8Int8 = new(map[uint8]int8)
				}
				yym487 := z.DecBinary()
				_ = yym487
				if false {
				} else {
					z.F.DecMapUint8Int8X(x.FptrMapUint8Int8, d)
				}
			}
		case "FMapUint8Int16":
			if r.TryDecodeAsNil() {
				x.FMapUint8Int16 = nil
			} else {
				yyv488 := &x.FMapUint8Int16
				yym489 := z.DecBinary()
				_ = yym489
				if false {
				} else {
					z.F.DecMapUint8Int16X(yyv488, d)
				}
			}
		case "FptrMapUint8Int16":
			if x.FptrMapUint8Int16 == nil {
				x.FptrMapUint8Int16 = new(map[uint8]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Int16 != nil {
					x.FptrMapUint8Int16 = nil
				}
			} else {
				if x.FptrMapUint8Int16 == nil {
					x.FptrMapUint8Int16 = new(map[uint8]int16)
				}
				yym491 := z.DecBinary()
				_ = yym491
				if false {
				} else {
					z.F.DecMapUint8Int16X(x.FptrMapUint8Int16, d)
				}
			}
		case "FMapUint8Int32":
			if r.TryDecodeAsNil() {
				x.FMapUint8Int32 = nil
			} else {
				yyv492 := &x.FMapUint8Int32
				yym493 := z.DecBinary()
				_ = yym493
				if false {
				} else {
					z.F.DecMapUint8Int32X(yyv492, d)
				}
			}
		case "FptrMapUint8Int32":
			if x.FptrMapUint8Int32 == nil {
				x.FptrMapUint8Int32 = new(map[uint8]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Int32 != nil {
					x.FptrMapUint8Int32 = nil
				}
			} else {
				if x.FptrMapUint8Int32 == nil {
					x.FptrMapUint8Int32 = new(map[uint8]int32)
				}
				yym495 := z.DecBinary()
				_ = yym495
				if false {
				} else {
					z.F.DecMapUint8Int32X(x.FptrMapUint8Int32, d)
				}
			}
		case "FMapUint8Int64":
			if r.TryDecodeAsNil() {
				x.FMapUint8Int64 = nil
			} else {
				yyv496 := &x.FMapUint8Int64
				yym497 := z.DecBinary()
				_ = yym497
				if false {
				} else {
					z.F.DecMapUint8Int64X(yyv496, d)
				}
			}
		case "FptrMapUint8Int64":
			if x.FptrMapUint8Int64 == nil {
				x.FptrMapUint8Int64 = new(map[uint8]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Int64 != nil {
					x.FptrMapUint8Int64 = nil
				}
			} else {
				if x.FptrMapUint8Int64 == nil {
					x.FptrMapUint8Int64 = new(map[uint8]int64)
				}
				yym499 := z.DecBinary()
				_ = yym499
				if false {
				} else {
					z.F.DecMapUint8Int64X(x.FptrMapUint8Int64, d)
				}
			}
		case "FMapUint8Float32":
			if r.TryDecodeAsNil() {
				x.FMapUint8Float32 = nil
			} else {
				yyv500 := &x.FMapUint8Float32
				yym501 := z.DecBinary()
				_ = yym501
				if false {
				} else {
					z.F.DecMapUint8Float32X(yyv500, d)
				}
			}
		case "FptrMapUint8Float32":
			if x.FptrMapUint8Float32 == nil {
				x.FptrMapUint8Float32 = new(map[uint8]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Float32 != nil {
					x.FptrMapUint8Float32 = nil
				}
			} else {
				if x.FptrMapUint8Float32 == nil {
					x.FptrMapUint8Float32 = new(map[uint8]float32)
				}
				yym503 := z.DecBinary()
				_ = yym503
				if false {
				} else {
					z.F.DecMapUint8Float32X(x.FptrMapUint8Float32, d)
				}
			}
		case "FMapUint8Float64":
			if r.TryDecodeAsNil() {
				x.FMapUint8Float64 = nil
			} else {
				yyv504 := &x.FMapUint8Float64
				yym505 := z.DecBinary()
				_ = yym505
				if false {
				} else {
					z.F.DecMapUint8Float64X(yyv504, d)
				}
			}
		case "FptrMapUint8Float64":
			if x.FptrMapUint8Float64 == nil {
				x.FptrMapUint8Float64 = new(map[uint8]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Float64 != nil {
					x.FptrMapUint8Float64 = nil
				}
			} else {
				if x.FptrMapUint8Float64 == nil {
					x.FptrMapUint8Float64 = new(map[uint8]float64)
				}
				yym507 := z.DecBinary()
				_ = yym507
				if false {
				} else {
					z.F.DecMapUint8Float64X(x.FptrMapUint8Float64, d)
				}
			}
		case "FMapUint8Bool":
			if r.TryDecodeAsNil() {
				x.FMapUint8Bool = nil
			} else {
				yyv508 := &x.FMapUint8Bool
				yym509 := z.DecBinary()
				_ = yym509
				if false {
				} else {
					z.F.DecMapUint8BoolX(yyv508, d)
				}
			}
		case "FptrMapUint8Bool":
			if x.FptrMapUint8Bool == nil {
				x.FptrMapUint8Bool = new(map[uint8]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint8Bool != nil {
					x.FptrMapUint8Bool = nil
				}
			} else {
				if x.FptrMapUint8Bool == nil {
					x.FptrMapUint8Bool = new(map[uint8]bool)
				}
				yym511 := z.DecBinary()
				_ = yym511
				if false {
				} else {
					z.F.DecMapUint8BoolX(x.FptrMapUint8Bool, d)
				}
			}
		case "FMapUint16Intf":
			if r.TryDecodeAsNil() {
				x.FMapUint16Intf = nil
			} else {
				yyv512 := &x.FMapUint16Intf
				yym513 := z.DecBinary()
				_ = yym513
				if false {
				} else {
					z.F.DecMapUint16IntfX(yyv512, d)
				}
			}
		case "FptrMapUint16Intf":
			if x.FptrMapUint16Intf == nil {
				x.FptrMapUint16Intf = new(map[uint16]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Intf != nil {
					x.FptrMapUint16Intf = nil
				}
			} else {
				if x.FptrMapUint16Intf == nil {
					x.FptrMapUint16Intf = new(map[uint16]interface{})
				}
				yym515 := z.DecBinary()
				_ = yym515
				if false {
				} else {
					z.F.DecMapUint16IntfX(x.FptrMapUint16Intf, d)
				}
			}
		case "FMapUint16String":
			if r.TryDecodeAsNil() {
				x.FMapUint16String = nil
			} else {
				yyv516 := &x.FMapUint16String
				yym517 := z.DecBinary()
				_ = yym517
				if false {
				} else {
					z.F.DecMapUint16StringX(yyv516, d)
				}
			}
		case "FptrMapUint16String":
			if x.FptrMapUint16String == nil {
				x.FptrMapUint16String = new(map[uint16]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16String != nil {
					x.FptrMapUint16String = nil
				}
			} else {
				if x.FptrMapUint16String == nil {
					x.FptrMapUint16String = new(map[uint16]string)
				}
				yym519 := z.DecBinary()
				_ = yym519
				if false {
				} else {
					z.F.DecMapUint16StringX(x.FptrMapUint16String, d)
				}
			}
		case "FMapUint16Uint":
			if r.TryDecodeAsNil() {
				x.FMapUint16Uint = nil
			} else {
				yyv520 := &x.FMapUint16Uint
				yym521 := z.DecBinary()
				_ = yym521
				if false {
				} else {
					z.F.DecMapUint16UintX(yyv520, d)
				}
			}
		case "FptrMapUint16Uint":
			if x.FptrMapUint16Uint == nil {
				x.FptrMapUint16Uint = new(map[uint16]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Uint != nil {
					x.FptrMapUint16Uint = nil
				}
			} else {
				if x.FptrMapUint16Uint == nil {
					x.FptrMapUint16Uint = new(map[uint16]uint)
				}
				yym523 := z.DecBinary()
				_ = yym523
				if false {
				} else {
					z.F.DecMapUint16UintX(x.FptrMapUint16Uint, d)
				}
			}
		case "FMapUint16Uint8":
			if r.TryDecodeAsNil() {
				x.FMapUint16Uint8 = nil
			} else {
				yyv524 := &x.FMapUint16Uint8
				yym525 := z.DecBinary()
				_ = yym525
				if false {
				} else {
					z.F.DecMapUint16Uint8X(yyv524, d)
				}
			}
		case "FptrMapUint16Uint8":
			if x.FptrMapUint16Uint8 == nil {
				x.FptrMapUint16Uint8 = new(map[uint16]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Uint8 != nil {
					x.FptrMapUint16Uint8 = nil
				}
			} else {
				if x.FptrMapUint16Uint8 == nil {
					x.FptrMapUint16Uint8 = new(map[uint16]uint8)
				}
				yym527 := z.DecBinary()
				_ = yym527
				if false {
				} else {
					z.F.DecMapUint16Uint8X(x.FptrMapUint16Uint8, d)
				}
			}
		case "FMapUint16Uint16":
			if r.TryDecodeAsNil() {
				x.FMapUint16Uint16 = nil
			} else {
				yyv528 := &x.FMapUint16Uint16
				yym529 := z.DecBinary()
				_ = yym529
				if false {
				} else {
					z.F.DecMapUint16Uint16X(yyv528, d)
				}
			}
		case "FptrMapUint16Uint16":
			if x.FptrMapUint16Uint16 == nil {
				x.FptrMapUint16Uint16 = new(map[uint16]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Uint16 != nil {
					x.FptrMapUint16Uint16 = nil
				}
			} else {
				if x.FptrMapUint16Uint16 == nil {
					x.FptrMapUint16Uint16 = new(map[uint16]uint16)
				}
				yym531 := z.DecBinary()
				_ = yym531
				if false {
				} else {
					z.F.DecMapUint16Uint16X(x.FptrMapUint16Uint16, d)
				}
			}
		case "FMapUint16Uint32":
			if r.TryDecodeAsNil() {
				x.FMapUint16Uint32 = nil
			} else {
				yyv532 := &x.FMapUint16Uint32
				yym533 := z.DecBinary()
				_ = yym533
				if false {
				} else {
					z.F.DecMapUint16Uint32X(yyv532, d)
				}
			}
		case "FptrMapUint16Uint32":
			if x.FptrMapUint16Uint32 == nil {
				x.FptrMapUint16Uint32 = new(map[uint16]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Uint32 != nil {
					x.FptrMapUint16Uint32 = nil
				}
			} else {
				if x.FptrMapUint16Uint32 == nil {
					x.FptrMapUint16Uint32 = new(map[uint16]uint32)
				}
				yym535 := z.DecBinary()
				_ = yym535
				if false {
				} else {
					z.F.DecMapUint16Uint32X(x.FptrMapUint16Uint32, d)
				}
			}
		case "FMapUint16Uint64":
			if r.TryDecodeAsNil() {
				x.FMapUint16Uint64 = nil
			} else {
				yyv536 := &x.FMapUint16Uint64
				yym537 := z.DecBinary()
				_ = yym537
				if false {
				} else {
					z.F.DecMapUint16Uint64X(yyv536, d)
				}
			}
		case "FptrMapUint16Uint64":
			if x.FptrMapUint16Uint64 == nil {
				x.FptrMapUint16Uint64 = new(map[uint16]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Uint64 != nil {
					x.FptrMapUint16Uint64 = nil
				}
			} else {
				if x.FptrMapUint16Uint64 == nil {
					x.FptrMapUint16Uint64 = new(map[uint16]uint64)
				}
				yym539 := z.DecBinary()
				_ = yym539
				if false {
				} else {
					z.F.DecMapUint16Uint64X(x.FptrMapUint16Uint64, d)
				}
			}
		case "FMapUint16Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapUint16Uintptr = nil
			} else {
				yyv540 := &x.FMapUint16Uintptr
				yym541 := z.DecBinary()
				_ = yym541
				if false {
				} else {
					z.F.DecMapUint16UintptrX(yyv540, d)
				}
			}
		case "FptrMapUint16Uintptr":
			if x.FptrMapUint16Uintptr == nil {
				x.FptrMapUint16Uintptr = new(map[uint16]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Uintptr != nil {
					x.FptrMapUint16Uintptr = nil
				}
			} else {
				if x.FptrMapUint16Uintptr == nil {
					x.FptrMapUint16Uintptr = new(map[uint16]uintptr)
				}
				yym543 := z.DecBinary()
				_ = yym543
				if false {
				} else {
					z.F.DecMapUint16UintptrX(x.FptrMapUint16Uintptr, d)
				}
			}
		case "FMapUint16Int":
			if r.TryDecodeAsNil() {
				x.FMapUint16Int = nil
			} else {
				yyv544 := &x.FMapUint16Int
				yym545 := z.DecBinary()
				_ = yym545
				if false {
				} else {
					z.F.DecMapUint16IntX(yyv544, d)
				}
			}
		case "FptrMapUint16Int":
			if x.FptrMapUint16Int == nil {
				x.FptrMapUint16Int = new(map[uint16]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Int != nil {
					x.FptrMapUint16Int = nil
				}
			} else {
				if x.FptrMapUint16Int == nil {
					x.FptrMapUint16Int = new(map[uint16]int)
				}
				yym547 := z.DecBinary()
				_ = yym547
				if false {
				} else {
					z.F.DecMapUint16IntX(x.FptrMapUint16Int, d)
				}
			}
		case "FMapUint16Int8":
			if r.TryDecodeAsNil() {
				x.FMapUint16Int8 = nil
			} else {
				yyv548 := &x.FMapUint16Int8
				yym549 := z.DecBinary()
				_ = yym549
				if false {
				} else {
					z.F.DecMapUint16Int8X(yyv548, d)
				}
			}
		case "FptrMapUint16Int8":
			if x.FptrMapUint16Int8 == nil {
				x.FptrMapUint16Int8 = new(map[uint16]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Int8 != nil {
					x.FptrMapUint16Int8 = nil
				}
			} else {
				if x.FptrMapUint16Int8 == nil {
					x.FptrMapUint16Int8 = new(map[uint16]int8)
				}
				yym551 := z.DecBinary()
				_ = yym551
				if false {
				} else {
					z.F.DecMapUint16Int8X(x.FptrMapUint16Int8, d)
				}
			}
		case "FMapUint16Int16":
			if r.TryDecodeAsNil() {
				x.FMapUint16Int16 = nil
			} else {
				yyv552 := &x.FMapUint16Int16
				yym553 := z.DecBinary()
				_ = yym553
				if false {
				} else {
					z.F.DecMapUint16Int16X(yyv552, d)
				}
			}
		case "FptrMapUint16Int16":
			if x.FptrMapUint16Int16 == nil {
				x.FptrMapUint16Int16 = new(map[uint16]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Int16 != nil {
					x.FptrMapUint16Int16 = nil
				}
			} else {
				if x.FptrMapUint16Int16 == nil {
					x.FptrMapUint16Int16 = new(map[uint16]int16)
				}
				yym555 := z.DecBinary()
				_ = yym555
				if false {
				} else {
					z.F.DecMapUint16Int16X(x.FptrMapUint16Int16, d)
				}
			}
		case "FMapUint16Int32":
			if r.TryDecodeAsNil() {
				x.FMapUint16Int32 = nil
			} else {
				yyv556 := &x.FMapUint16Int32
				yym557 := z.DecBinary()
				_ = yym557
				if false {
				} else {
					z.F.DecMapUint16Int32X(yyv556, d)
				}
			}
		case "FptrMapUint16Int32":
			if x.FptrMapUint16Int32 == nil {
				x.FptrMapUint16Int32 = new(map[uint16]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Int32 != nil {
					x.FptrMapUint16Int32 = nil
				}
			} else {
				if x.FptrMapUint16Int32 == nil {
					x.FptrMapUint16Int32 = new(map[uint16]int32)
				}
				yym559 := z.DecBinary()
				_ = yym559
				if false {
				} else {
					z.F.DecMapUint16Int32X(x.FptrMapUint16Int32, d)
				}
			}
		case "FMapUint16Int64":
			if r.TryDecodeAsNil() {
				x.FMapUint16Int64 = nil
			} else {
				yyv560 := &x.FMapUint16Int64
				yym561 := z.DecBinary()
				_ = yym561
				if false {
				} else {
					z.F.DecMapUint16Int64X(yyv560, d)
				}
			}
		case "FptrMapUint16Int64":
			if x.FptrMapUint16Int64 == nil {
				x.FptrMapUint16Int64 = new(map[uint16]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Int64 != nil {
					x.FptrMapUint16Int64 = nil
				}
			} else {
				if x.FptrMapUint16Int64 == nil {
					x.FptrMapUint16Int64 = new(map[uint16]int64)
				}
				yym563 := z.DecBinary()
				_ = yym563
				if false {
				} else {
					z.F.DecMapUint16Int64X(x.FptrMapUint16Int64, d)
				}
			}
		case "FMapUint16Float32":
			if r.TryDecodeAsNil() {
				x.FMapUint16Float32 = nil
			} else {
				yyv564 := &x.FMapUint16Float32
				yym565 := z.DecBinary()
				_ = yym565
				if false {
				} else {
					z.F.DecMapUint16Float32X(yyv564, d)
				}
			}
		case "FptrMapUint16Float32":
			if x.FptrMapUint16Float32 == nil {
				x.FptrMapUint16Float32 = new(map[uint16]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Float32 != nil {
					x.FptrMapUint16Float32 = nil
				}
			} else {
				if x.FptrMapUint16Float32 == nil {
					x.FptrMapUint16Float32 = new(map[uint16]float32)
				}
				yym567 := z.DecBinary()
				_ = yym567
				if false {
				} else {
					z.F.DecMapUint16Float32X(x.FptrMapUint16Float32, d)
				}
			}
		case "FMapUint16Float64":
			if r.TryDecodeAsNil() {
				x.FMapUint16Float64 = nil
			} else {
				yyv568 := &x.FMapUint16Float64
				yym569 := z.DecBinary()
				_ = yym569
				if false {
				} else {
					z.F.DecMapUint16Float64X(yyv568, d)
				}
			}
		case "FptrMapUint16Float64":
			if x.FptrMapUint16Float64 == nil {
				x.FptrMapUint16Float64 = new(map[uint16]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Float64 != nil {
					x.FptrMapUint16Float64 = nil
				}
			} else {
				if x.FptrMapUint16Float64 == nil {
					x.FptrMapUint16Float64 = new(map[uint16]float64)
				}
				yym571 := z.DecBinary()
				_ = yym571
				if false {
				} else {
					z.F.DecMapUint16Float64X(x.FptrMapUint16Float64, d)
				}
			}
		case "FMapUint16Bool":
			if r.TryDecodeAsNil() {
				x.FMapUint16Bool = nil
			} else {
				yyv572 := &x.FMapUint16Bool
				yym573 := z.DecBinary()
				_ = yym573
				if false {
				} else {
					z.F.DecMapUint16BoolX(yyv572, d)
				}
			}
		case "FptrMapUint16Bool":
			if x.FptrMapUint16Bool == nil {
				x.FptrMapUint16Bool = new(map[uint16]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint16Bool != nil {
					x.FptrMapUint16Bool = nil
				}
			} else {
				if x.FptrMapUint16Bool == nil {
					x.FptrMapUint16Bool = new(map[uint16]bool)
				}
				yym575 := z.DecBinary()
				_ = yym575
				if false {
				} else {
					z.F.DecMapUint16BoolX(x.FptrMapUint16Bool, d)
				}
			}
		case "FMapUint32Intf":
			if r.TryDecodeAsNil() {
				x.FMapUint32Intf = nil
			} else {
				yyv576 := &x.FMapUint32Intf
				yym577 := z.DecBinary()
				_ = yym577
				if false {
				} else {
					z.F.DecMapUint32IntfX(yyv576, d)
				}
			}
		case "FptrMapUint32Intf":
			if x.FptrMapUint32Intf == nil {
				x.FptrMapUint32Intf = new(map[uint32]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Intf != nil {
					x.FptrMapUint32Intf = nil
				}
			} else {
				if x.FptrMapUint32Intf == nil {
					x.FptrMapUint32Intf = new(map[uint32]interface{})
				}
				yym579 := z.DecBinary()
				_ = yym579
				if false {
				} else {
					z.F.DecMapUint32IntfX(x.FptrMapUint32Intf, d)
				}
			}
		case "FMapUint32String":
			if r.TryDecodeAsNil() {
				x.FMapUint32String = nil
			} else {
				yyv580 := &x.FMapUint32String
				yym581 := z.DecBinary()
				_ = yym581
				if false {
				} else {
					z.F.DecMapUint32StringX(yyv580, d)
				}
			}
		case "FptrMapUint32String":
			if x.FptrMapUint32String == nil {
				x.FptrMapUint32String = new(map[uint32]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32String != nil {
					x.FptrMapUint32String = nil
				}
			} else {
				if x.FptrMapUint32String == nil {
					x.FptrMapUint32String = new(map[uint32]string)
				}
				yym583 := z.DecBinary()
				_ = yym583
				if false {
				} else {
					z.F.DecMapUint32StringX(x.FptrMapUint32String, d)
				}
			}
		case "FMapUint32Uint":
			if r.TryDecodeAsNil() {
				x.FMapUint32Uint = nil
			} else {
				yyv584 := &x.FMapUint32Uint
				yym585 := z.DecBinary()
				_ = yym585
				if false {
				} else {
					z.F.DecMapUint32UintX(yyv584, d)
				}
			}
		case "FptrMapUint32Uint":
			if x.FptrMapUint32Uint == nil {
				x.FptrMapUint32Uint = new(map[uint32]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Uint != nil {
					x.FptrMapUint32Uint = nil
				}
			} else {
				if x.FptrMapUint32Uint == nil {
					x.FptrMapUint32Uint = new(map[uint32]uint)
				}
				yym587 := z.DecBinary()
				_ = yym587
				if false {
				} else {
					z.F.DecMapUint32UintX(x.FptrMapUint32Uint, d)
				}
			}
		case "FMapUint32Uint8":
			if r.TryDecodeAsNil() {
				x.FMapUint32Uint8 = nil
			} else {
				yyv588 := &x.FMapUint32Uint8
				yym589 := z.DecBinary()
				_ = yym589
				if false {
				} else {
					z.F.DecMapUint32Uint8X(yyv588, d)
				}
			}
		case "FptrMapUint32Uint8":
			if x.FptrMapUint32Uint8 == nil {
				x.FptrMapUint32Uint8 = new(map[uint32]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Uint8 != nil {
					x.FptrMapUint32Uint8 = nil
				}
			} else {
				if x.FptrMapUint32Uint8 == nil {
					x.FptrMapUint32Uint8 = new(map[uint32]uint8)
				}
				yym591 := z.DecBinary()
				_ = yym591
				if false {
				} else {
					z.F.DecMapUint32Uint8X(x.FptrMapUint32Uint8, d)
				}
			}
		case "FMapUint32Uint16":
			if r.TryDecodeAsNil() {
				x.FMapUint32Uint16 = nil
			} else {
				yyv592 := &x.FMapUint32Uint16
				yym593 := z.DecBinary()
				_ = yym593
				if false {
				} else {
					z.F.DecMapUint32Uint16X(yyv592, d)
				}
			}
		case "FptrMapUint32Uint16":
			if x.FptrMapUint32Uint16 == nil {
				x.FptrMapUint32Uint16 = new(map[uint32]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Uint16 != nil {
					x.FptrMapUint32Uint16 = nil
				}
			} else {
				if x.FptrMapUint32Uint16 == nil {
					x.FptrMapUint32Uint16 = new(map[uint32]uint16)
				}
				yym595 := z.DecBinary()
				_ = yym595
				if false {
				} else {
					z.F.DecMapUint32Uint16X(x.FptrMapUint32Uint16, d)
				}
			}
		case "FMapUint32Uint32":
			if r.TryDecodeAsNil() {
				x.FMapUint32Uint32 = nil
			} else {
				yyv596 := &x.FMapUint32Uint32
				yym597 := z.DecBinary()
				_ = yym597
				if false {
				} else {
					z.F.DecMapUint32Uint32X(yyv596, d)
				}
			}
		case "FptrMapUint32Uint32":
			if x.FptrMapUint32Uint32 == nil {
				x.FptrMapUint32Uint32 = new(map[uint32]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Uint32 != nil {
					x.FptrMapUint32Uint32 = nil
				}
			} else {
				if x.FptrMapUint32Uint32 == nil {
					x.FptrMapUint32Uint32 = new(map[uint32]uint32)
				}
				yym599 := z.DecBinary()
				_ = yym599
				if false {
				} else {
					z.F.DecMapUint32Uint32X(x.FptrMapUint32Uint32, d)
				}
			}
		case "FMapUint32Uint64":
			if r.TryDecodeAsNil() {
				x.FMapUint32Uint64 = nil
			} else {
				yyv600 := &x.FMapUint32Uint64
				yym601 := z.DecBinary()
				_ = yym601
				if false {
				} else {
					z.F.DecMapUint32Uint64X(yyv600, d)
				}
			}
		case "FptrMapUint32Uint64":
			if x.FptrMapUint32Uint64 == nil {
				x.FptrMapUint32Uint64 = new(map[uint32]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Uint64 != nil {
					x.FptrMapUint32Uint64 = nil
				}
			} else {
				if x.FptrMapUint32Uint64 == nil {
					x.FptrMapUint32Uint64 = new(map[uint32]uint64)
				}
				yym603 := z.DecBinary()
				_ = yym603
				if false {
				} else {
					z.F.DecMapUint32Uint64X(x.FptrMapUint32Uint64, d)
				}
			}
		case "FMapUint32Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapUint32Uintptr = nil
			} else {
				yyv604 := &x.FMapUint32Uintptr
				yym605 := z.DecBinary()
				_ = yym605
				if false {
				} else {
					z.F.DecMapUint32UintptrX(yyv604, d)
				}
			}
		case "FptrMapUint32Uintptr":
			if x.FptrMapUint32Uintptr == nil {
				x.FptrMapUint32Uintptr = new(map[uint32]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Uintptr != nil {
					x.FptrMapUint32Uintptr = nil
				}
			} else {
				if x.FptrMapUint32Uintptr == nil {
					x.FptrMapUint32Uintptr = new(map[uint32]uintptr)
				}
				yym607 := z.DecBinary()
				_ = yym607
				if false {
				} else {
					z.F.DecMapUint32UintptrX(x.FptrMapUint32Uintptr, d)
				}
			}
		case "FMapUint32Int":
			if r.TryDecodeAsNil() {
				x.FMapUint32Int = nil
			} else {
				yyv608 := &x.FMapUint32Int
				yym609 := z.DecBinary()
				_ = yym609
				if false {
				} else {
					z.F.DecMapUint32IntX(yyv608, d)
				}
			}
		case "FptrMapUint32Int":
			if x.FptrMapUint32Int == nil {
				x.FptrMapUint32Int = new(map[uint32]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Int != nil {
					x.FptrMapUint32Int = nil
				}
			} else {
				if x.FptrMapUint32Int == nil {
					x.FptrMapUint32Int = new(map[uint32]int)
				}
				yym611 := z.DecBinary()
				_ = yym611
				if false {
				} else {
					z.F.DecMapUint32IntX(x.FptrMapUint32Int, d)
				}
			}
		case "FMapUint32Int8":
			if r.TryDecodeAsNil() {
				x.FMapUint32Int8 = nil
			} else {
				yyv612 := &x.FMapUint32Int8
				yym613 := z.DecBinary()
				_ = yym613
				if false {
				} else {
					z.F.DecMapUint32Int8X(yyv612, d)
				}
			}
		case "FptrMapUint32Int8":
			if x.FptrMapUint32Int8 == nil {
				x.FptrMapUint32Int8 = new(map[uint32]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Int8 != nil {
					x.FptrMapUint32Int8 = nil
				}
			} else {
				if x.FptrMapUint32Int8 == nil {
					x.FptrMapUint32Int8 = new(map[uint32]int8)
				}
				yym615 := z.DecBinary()
				_ = yym615
				if false {
				} else {
					z.F.DecMapUint32Int8X(x.FptrMapUint32Int8, d)
				}
			}
		case "FMapUint32Int16":
			if r.TryDecodeAsNil() {
				x.FMapUint32Int16 = nil
			} else {
				yyv616 := &x.FMapUint32Int16
				yym617 := z.DecBinary()
				_ = yym617
				if false {
				} else {
					z.F.DecMapUint32Int16X(yyv616, d)
				}
			}
		case "FptrMapUint32Int16":
			if x.FptrMapUint32Int16 == nil {
				x.FptrMapUint32Int16 = new(map[uint32]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Int16 != nil {
					x.FptrMapUint32Int16 = nil
				}
			} else {
				if x.FptrMapUint32Int16 == nil {
					x.FptrMapUint32Int16 = new(map[uint32]int16)
				}
				yym619 := z.DecBinary()
				_ = yym619
				if false {
				} else {
					z.F.DecMapUint32Int16X(x.FptrMapUint32Int16, d)
				}
			}
		case "FMapUint32Int32":
			if r.TryDecodeAsNil() {
				x.FMapUint32Int32 = nil
			} else {
				yyv620 := &x.FMapUint32Int32
				yym621 := z.DecBinary()
				_ = yym621
				if false {
				} else {
					z.F.DecMapUint32Int32X(yyv620, d)
				}
			}
		case "FptrMapUint32Int32":
			if x.FptrMapUint32Int32 == nil {
				x.FptrMapUint32Int32 = new(map[uint32]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Int32 != nil {
					x.FptrMapUint32Int32 = nil
				}
			} else {
				if x.FptrMapUint32Int32 == nil {
					x.FptrMapUint32Int32 = new(map[uint32]int32)
				}
				yym623 := z.DecBinary()
				_ = yym623
				if false {
				} else {
					z.F.DecMapUint32Int32X(x.FptrMapUint32Int32, d)
				}
			}
		case "FMapUint32Int64":
			if r.TryDecodeAsNil() {
				x.FMapUint32Int64 = nil
			} else {
				yyv624 := &x.FMapUint32Int64
				yym625 := z.DecBinary()
				_ = yym625
				if false {
				} else {
					z.F.DecMapUint32Int64X(yyv624, d)
				}
			}
		case "FptrMapUint32Int64":
			if x.FptrMapUint32Int64 == nil {
				x.FptrMapUint32Int64 = new(map[uint32]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Int64 != nil {
					x.FptrMapUint32Int64 = nil
				}
			} else {
				if x.FptrMapUint32Int64 == nil {
					x.FptrMapUint32Int64 = new(map[uint32]int64)
				}
				yym627 := z.DecBinary()
				_ = yym627
				if false {
				} else {
					z.F.DecMapUint32Int64X(x.FptrMapUint32Int64, d)
				}
			}
		case "FMapUint32Float32":
			if r.TryDecodeAsNil() {
				x.FMapUint32Float32 = nil
			} else {
				yyv628 := &x.FMapUint32Float32
				yym629 := z.DecBinary()
				_ = yym629
				if false {
				} else {
					z.F.DecMapUint32Float32X(yyv628, d)
				}
			}
		case "FptrMapUint32Float32":
			if x.FptrMapUint32Float32 == nil {
				x.FptrMapUint32Float32 = new(map[uint32]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Float32 != nil {
					x.FptrMapUint32Float32 = nil
				}
			} else {
				if x.FptrMapUint32Float32 == nil {
					x.FptrMapUint32Float32 = new(map[uint32]float32)
				}
				yym631 := z.DecBinary()
				_ = yym631
				if false {
				} else {
					z.F.DecMapUint32Float32X(x.FptrMapUint32Float32, d)
				}
			}
		case "FMapUint32Float64":
			if r.TryDecodeAsNil() {
				x.FMapUint32Float64 = nil
			} else {
				yyv632 := &x.FMapUint32Float64
				yym633 := z.DecBinary()
				_ = yym633
				if false {
				} else {
					z.F.DecMapUint32Float64X(yyv632, d)
				}
			}
		case "FptrMapUint32Float64":
			if x.FptrMapUint32Float64 == nil {
				x.FptrMapUint32Float64 = new(map[uint32]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Float64 != nil {
					x.FptrMapUint32Float64 = nil
				}
			} else {
				if x.FptrMapUint32Float64 == nil {
					x.FptrMapUint32Float64 = new(map[uint32]float64)
				}
				yym635 := z.DecBinary()
				_ = yym635
				if false {
				} else {
					z.F.DecMapUint32Float64X(x.FptrMapUint32Float64, d)
				}
			}
		case "FMapUint32Bool":
			if r.TryDecodeAsNil() {
				x.FMapUint32Bool = nil
			} else {
				yyv636 := &x.FMapUint32Bool
				yym637 := z.DecBinary()
				_ = yym637
				if false {
				} else {
					z.F.DecMapUint32BoolX(yyv636, d)
				}
			}
		case "FptrMapUint32Bool":
			if x.FptrMapUint32Bool == nil {
				x.FptrMapUint32Bool = new(map[uint32]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint32Bool != nil {
					x.FptrMapUint32Bool = nil
				}
			} else {
				if x.FptrMapUint32Bool == nil {
					x.FptrMapUint32Bool = new(map[uint32]bool)
				}
				yym639 := z.DecBinary()
				_ = yym639
				if false {
				} else {
					z.F.DecMapUint32BoolX(x.FptrMapUint32Bool, d)
				}
			}
		case "FMapUint64Intf":
			if r.TryDecodeAsNil() {
				x.FMapUint64Intf = nil
			} else {
				yyv640 := &x.FMapUint64Intf
				yym641 := z.DecBinary()
				_ = yym641
				if false {
				} else {
					z.F.DecMapUint64IntfX(yyv640, d)
				}
			}
		case "FptrMapUint64Intf":
			if x.FptrMapUint64Intf == nil {
				x.FptrMapUint64Intf = new(map[uint64]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Intf != nil {
					x.FptrMapUint64Intf = nil
				}
			} else {
				if x.FptrMapUint64Intf == nil {
					x.FptrMapUint64Intf = new(map[uint64]interface{})
				}
				yym643 := z.DecBinary()
				_ = yym643
				if false {
				} else {
					z.F.DecMapUint64IntfX(x.FptrMapUint64Intf, d)
				}
			}
		case "FMapUint64String":
			if r.TryDecodeAsNil() {
				x.FMapUint64String = nil
			} else {
				yyv644 := &x.FMapUint64String
				yym645 := z.DecBinary()
				_ = yym645
				if false {
				} else {
					z.F.DecMapUint64StringX(yyv644, d)
				}
			}
		case "FptrMapUint64String":
			if x.FptrMapUint64String == nil {
				x.FptrMapUint64String = new(map[uint64]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64String != nil {
					x.FptrMapUint64String = nil
				}
			} else {
				if x.FptrMapUint64String == nil {
					x.FptrMapUint64String = new(map[uint64]string)
				}
				yym647 := z.DecBinary()
				_ = yym647
				if false {
				} else {
					z.F.DecMapUint64StringX(x.FptrMapUint64String, d)
				}
			}
		case "FMapUint64Uint":
			if r.TryDecodeAsNil() {
				x.FMapUint64Uint = nil
			} else {
				yyv648 := &x.FMapUint64Uint
				yym649 := z.DecBinary()
				_ = yym649
				if false {
				} else {
					z.F.DecMapUint64UintX(yyv648, d)
				}
			}
		case "FptrMapUint64Uint":
			if x.FptrMapUint64Uint == nil {
				x.FptrMapUint64Uint = new(map[uint64]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Uint != nil {
					x.FptrMapUint64Uint = nil
				}
			} else {
				if x.FptrMapUint64Uint == nil {
					x.FptrMapUint64Uint = new(map[uint64]uint)
				}
				yym651 := z.DecBinary()
				_ = yym651
				if false {
				} else {
					z.F.DecMapUint64UintX(x.FptrMapUint64Uint, d)
				}
			}
		case "FMapUint64Uint8":
			if r.TryDecodeAsNil() {
				x.FMapUint64Uint8 = nil
			} else {
				yyv652 := &x.FMapUint64Uint8
				yym653 := z.DecBinary()
				_ = yym653
				if false {
				} else {
					z.F.DecMapUint64Uint8X(yyv652, d)
				}
			}
		case "FptrMapUint64Uint8":
			if x.FptrMapUint64Uint8 == nil {
				x.FptrMapUint64Uint8 = new(map[uint64]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Uint8 != nil {
					x.FptrMapUint64Uint8 = nil
				}
			} else {
				if x.FptrMapUint64Uint8 == nil {
					x.FptrMapUint64Uint8 = new(map[uint64]uint8)
				}
				yym655 := z.DecBinary()
				_ = yym655
				if false {
				} else {
					z.F.DecMapUint64Uint8X(x.FptrMapUint64Uint8, d)
				}
			}
		case "FMapUint64Uint16":
			if r.TryDecodeAsNil() {
				x.FMapUint64Uint16 = nil
			} else {
				yyv656 := &x.FMapUint64Uint16
				yym657 := z.DecBinary()
				_ = yym657
				if false {
				} else {
					z.F.DecMapUint64Uint16X(yyv656, d)
				}
			}
		case "FptrMapUint64Uint16":
			if x.FptrMapUint64Uint16 == nil {
				x.FptrMapUint64Uint16 = new(map[uint64]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Uint16 != nil {
					x.FptrMapUint64Uint16 = nil
				}
			} else {
				if x.FptrMapUint64Uint16 == nil {
					x.FptrMapUint64Uint16 = new(map[uint64]uint16)
				}
				yym659 := z.DecBinary()
				_ = yym659
				if false {
				} else {
					z.F.DecMapUint64Uint16X(x.FptrMapUint64Uint16, d)
				}
			}
		case "FMapUint64Uint32":
			if r.TryDecodeAsNil() {
				x.FMapUint64Uint32 = nil
			} else {
				yyv660 := &x.FMapUint64Uint32
				yym661 := z.DecBinary()
				_ = yym661
				if false {
				} else {
					z.F.DecMapUint64Uint32X(yyv660, d)
				}
			}
		case "FptrMapUint64Uint32":
			if x.FptrMapUint64Uint32 == nil {
				x.FptrMapUint64Uint32 = new(map[uint64]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Uint32 != nil {
					x.FptrMapUint64Uint32 = nil
				}
			} else {
				if x.FptrMapUint64Uint32 == nil {
					x.FptrMapUint64Uint32 = new(map[uint64]uint32)
				}
				yym663 := z.DecBinary()
				_ = yym663
				if false {
				} else {
					z.F.DecMapUint64Uint32X(x.FptrMapUint64Uint32, d)
				}
			}
		case "FMapUint64Uint64":
			if r.TryDecodeAsNil() {
				x.FMapUint64Uint64 = nil
			} else {
				yyv664 := &x.FMapUint64Uint64
				yym665 := z.DecBinary()
				_ = yym665
				if false {
				} else {
					z.F.DecMapUint64Uint64X(yyv664, d)
				}
			}
		case "FptrMapUint64Uint64":
			if x.FptrMapUint64Uint64 == nil {
				x.FptrMapUint64Uint64 = new(map[uint64]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Uint64 != nil {
					x.FptrMapUint64Uint64 = nil
				}
			} else {
				if x.FptrMapUint64Uint64 == nil {
					x.FptrMapUint64Uint64 = new(map[uint64]uint64)
				}
				yym667 := z.DecBinary()
				_ = yym667
				if false {
				} else {
					z.F.DecMapUint64Uint64X(x.FptrMapUint64Uint64, d)
				}
			}
		case "FMapUint64Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapUint64Uintptr = nil
			} else {
				yyv668 := &x.FMapUint64Uintptr
				yym669 := z.DecBinary()
				_ = yym669
				if false {
				} else {
					z.F.DecMapUint64UintptrX(yyv668, d)
				}
			}
		case "FptrMapUint64Uintptr":
			if x.FptrMapUint64Uintptr == nil {
				x.FptrMapUint64Uintptr = new(map[uint64]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Uintptr != nil {
					x.FptrMapUint64Uintptr = nil
				}
			} else {
				if x.FptrMapUint64Uintptr == nil {
					x.FptrMapUint64Uintptr = new(map[uint64]uintptr)
				}
				yym671 := z.DecBinary()
				_ = yym671
				if false {
				} else {
					z.F.DecMapUint64UintptrX(x.FptrMapUint64Uintptr, d)
				}
			}
		case "FMapUint64Int":
			if r.TryDecodeAsNil() {
				x.FMapUint64Int = nil
			} else {
				yyv672 := &x.FMapUint64Int
				yym673 := z.DecBinary()
				_ = yym673
				if false {
				} else {
					z.F.DecMapUint64IntX(yyv672, d)
				}
			}
		case "FptrMapUint64Int":
			if x.FptrMapUint64Int == nil {
				x.FptrMapUint64Int = new(map[uint64]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Int != nil {
					x.FptrMapUint64Int = nil
				}
			} else {
				if x.FptrMapUint64Int == nil {
					x.FptrMapUint64Int = new(map[uint64]int)
				}
				yym675 := z.DecBinary()
				_ = yym675
				if false {
				} else {
					z.F.DecMapUint64IntX(x.FptrMapUint64Int, d)
				}
			}
		case "FMapUint64Int8":
			if r.TryDecodeAsNil() {
				x.FMapUint64Int8 = nil
			} else {
				yyv676 := &x.FMapUint64Int8
				yym677 := z.DecBinary()
				_ = yym677
				if false {
				} else {
					z.F.DecMapUint64Int8X(yyv676, d)
				}
			}
		case "FptrMapUint64Int8":
			if x.FptrMapUint64Int8 == nil {
				x.FptrMapUint64Int8 = new(map[uint64]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Int8 != nil {
					x.FptrMapUint64Int8 = nil
				}
			} else {
				if x.FptrMapUint64Int8 == nil {
					x.FptrMapUint64Int8 = new(map[uint64]int8)
				}
				yym679 := z.DecBinary()
				_ = yym679
				if false {
				} else {
					z.F.DecMapUint64Int8X(x.FptrMapUint64Int8, d)
				}
			}
		case "FMapUint64Int16":
			if r.TryDecodeAsNil() {
				x.FMapUint64Int16 = nil
			} else {
				yyv680 := &x.FMapUint64Int16
				yym681 := z.DecBinary()
				_ = yym681
				if false {
				} else {
					z.F.DecMapUint64Int16X(yyv680, d)
				}
			}
		case "FptrMapUint64Int16":
			if x.FptrMapUint64Int16 == nil {
				x.FptrMapUint64Int16 = new(map[uint64]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Int16 != nil {
					x.FptrMapUint64Int16 = nil
				}
			} else {
				if x.FptrMapUint64Int16 == nil {
					x.FptrMapUint64Int16 = new(map[uint64]int16)
				}
				yym683 := z.DecBinary()
				_ = yym683
				if false {
				} else {
					z.F.DecMapUint64Int16X(x.FptrMapUint64Int16, d)
				}
			}
		case "FMapUint64Int32":
			if r.TryDecodeAsNil() {
				x.FMapUint64Int32 = nil
			} else {
				yyv684 := &x.FMapUint64Int32
				yym685 := z.DecBinary()
				_ = yym685
				if false {
				} else {
					z.F.DecMapUint64Int32X(yyv684, d)
				}
			}
		case "FptrMapUint64Int32":
			if x.FptrMapUint64Int32 == nil {
				x.FptrMapUint64Int32 = new(map[uint64]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Int32 != nil {
					x.FptrMapUint64Int32 = nil
				}
			} else {
				if x.FptrMapUint64Int32 == nil {
					x.FptrMapUint64Int32 = new(map[uint64]int32)
				}
				yym687 := z.DecBinary()
				_ = yym687
				if false {
				} else {
					z.F.DecMapUint64Int32X(x.FptrMapUint64Int32, d)
				}
			}
		case "FMapUint64Int64":
			if r.TryDecodeAsNil() {
				x.FMapUint64Int64 = nil
			} else {
				yyv688 := &x.FMapUint64Int64
				yym689 := z.DecBinary()
				_ = yym689
				if false {
				} else {
					z.F.DecMapUint64Int64X(yyv688, d)
				}
			}
		case "FptrMapUint64Int64":
			if x.FptrMapUint64Int64 == nil {
				x.FptrMapUint64Int64 = new(map[uint64]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Int64 != nil {
					x.FptrMapUint64Int64 = nil
				}
			} else {
				if x.FptrMapUint64Int64 == nil {
					x.FptrMapUint64Int64 = new(map[uint64]int64)
				}
				yym691 := z.DecBinary()
				_ = yym691
				if false {
				} else {
					z.F.DecMapUint64Int64X(x.FptrMapUint64Int64, d)
				}
			}
		case "FMapUint64Float32":
			if r.TryDecodeAsNil() {
				x.FMapUint64Float32 = nil
			} else {
				yyv692 := &x.FMapUint64Float32
				yym693 := z.DecBinary()
				_ = yym693
				if false {
				} else {
					z.F.DecMapUint64Float32X(yyv692, d)
				}
			}
		case "FptrMapUint64Float32":
			if x.FptrMapUint64Float32 == nil {
				x.FptrMapUint64Float32 = new(map[uint64]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Float32 != nil {
					x.FptrMapUint64Float32 = nil
				}
			} else {
				if x.FptrMapUint64Float32 == nil {
					x.FptrMapUint64Float32 = new(map[uint64]float32)
				}
				yym695 := z.DecBinary()
				_ = yym695
				if false {
				} else {
					z.F.DecMapUint64Float32X(x.FptrMapUint64Float32, d)
				}
			}
		case "FMapUint64Float64":
			if r.TryDecodeAsNil() {
				x.FMapUint64Float64 = nil
			} else {
				yyv696 := &x.FMapUint64Float64
				yym697 := z.DecBinary()
				_ = yym697
				if false {
				} else {
					z.F.DecMapUint64Float64X(yyv696, d)
				}
			}
		case "FptrMapUint64Float64":
			if x.FptrMapUint64Float64 == nil {
				x.FptrMapUint64Float64 = new(map[uint64]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Float64 != nil {
					x.FptrMapUint64Float64 = nil
				}
			} else {
				if x.FptrMapUint64Float64 == nil {
					x.FptrMapUint64Float64 = new(map[uint64]float64)
				}
				yym699 := z.DecBinary()
				_ = yym699
				if false {
				} else {
					z.F.DecMapUint64Float64X(x.FptrMapUint64Float64, d)
				}
			}
		case "FMapUint64Bool":
			if r.TryDecodeAsNil() {
				x.FMapUint64Bool = nil
			} else {
				yyv700 := &x.FMapUint64Bool
				yym701 := z.DecBinary()
				_ = yym701
				if false {
				} else {
					z.F.DecMapUint64BoolX(yyv700, d)
				}
			}
		case "FptrMapUint64Bool":
			if x.FptrMapUint64Bool == nil {
				x.FptrMapUint64Bool = new(map[uint64]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUint64Bool != nil {
					x.FptrMapUint64Bool = nil
				}
			} else {
				if x.FptrMapUint64Bool == nil {
					x.FptrMapUint64Bool = new(map[uint64]bool)
				}
				yym703 := z.DecBinary()
				_ = yym703
				if false {
				} else {
					z.F.DecMapUint64BoolX(x.FptrMapUint64Bool, d)
				}
			}
		case "FMapUintptrIntf":
			if r.TryDecodeAsNil() {
				x.FMapUintptrIntf = nil
			} else {
				yyv704 := &x.FMapUintptrIntf
				yym705 := z.DecBinary()
				_ = yym705
				if false {
				} else {
					z.F.DecMapUintptrIntfX(yyv704, d)
				}
			}
		case "FptrMapUintptrIntf":
			if x.FptrMapUintptrIntf == nil {
				x.FptrMapUintptrIntf = new(map[uintptr]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrIntf != nil {
					x.FptrMapUintptrIntf = nil
				}
			} else {
				if x.FptrMapUintptrIntf == nil {
					x.FptrMapUintptrIntf = new(map[uintptr]interface{})
				}
				yym707 := z.DecBinary()
				_ = yym707
				if false {
				} else {
					z.F.DecMapUintptrIntfX(x.FptrMapUintptrIntf, d)
				}
			}
		case "FMapUintptrString":
			if r.TryDecodeAsNil() {
				x.FMapUintptrString = nil
			} else {
				yyv708 := &x.FMapUintptrString
				yym709 := z.DecBinary()
				_ = yym709
				if false {
				} else {
					z.F.DecMapUintptrStringX(yyv708, d)
				}
			}
		case "FptrMapUintptrString":
			if x.FptrMapUintptrString == nil {
				x.FptrMapUintptrString = new(map[uintptr]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrString != nil {
					x.FptrMapUintptrString = nil
				}
			} else {
				if x.FptrMapUintptrString == nil {
					x.FptrMapUintptrString = new(map[uintptr]string)
				}
				yym711 := z.DecBinary()
				_ = yym711
				if false {
				} else {
					z.F.DecMapUintptrStringX(x.FptrMapUintptrString, d)
				}
			}
		case "FMapUintptrUint":
			if r.TryDecodeAsNil() {
				x.FMapUintptrUint = nil
			} else {
				yyv712 := &x.FMapUintptrUint
				yym713 := z.DecBinary()
				_ = yym713
				if false {
				} else {
					z.F.DecMapUintptrUintX(yyv712, d)
				}
			}
		case "FptrMapUintptrUint":
			if x.FptrMapUintptrUint == nil {
				x.FptrMapUintptrUint = new(map[uintptr]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrUint != nil {
					x.FptrMapUintptrUint = nil
				}
			} else {
				if x.FptrMapUintptrUint == nil {
					x.FptrMapUintptrUint = new(map[uintptr]uint)
				}
				yym715 := z.DecBinary()
				_ = yym715
				if false {
				} else {
					z.F.DecMapUintptrUintX(x.FptrMapUintptrUint, d)
				}
			}
		case "FMapUintptrUint8":
			if r.TryDecodeAsNil() {
				x.FMapUintptrUint8 = nil
			} else {
				yyv716 := &x.FMapUintptrUint8
				yym717 := z.DecBinary()
				_ = yym717
				if false {
				} else {
					z.F.DecMapUintptrUint8X(yyv716, d)
				}
			}
		case "FptrMapUintptrUint8":
			if x.FptrMapUintptrUint8 == nil {
				x.FptrMapUintptrUint8 = new(map[uintptr]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrUint8 != nil {
					x.FptrMapUintptrUint8 = nil
				}
			} else {
				if x.FptrMapUintptrUint8 == nil {
					x.FptrMapUintptrUint8 = new(map[uintptr]uint8)
				}
				yym719 := z.DecBinary()
				_ = yym719
				if false {
				} else {
					z.F.DecMapUintptrUint8X(x.FptrMapUintptrUint8, d)
				}
			}
		case "FMapUintptrUint16":
			if r.TryDecodeAsNil() {
				x.FMapUintptrUint16 = nil
			} else {
				yyv720 := &x.FMapUintptrUint16
				yym721 := z.DecBinary()
				_ = yym721
				if false {
				} else {
					z.F.DecMapUintptrUint16X(yyv720, d)
				}
			}
		case "FptrMapUintptrUint16":
			if x.FptrMapUintptrUint16 == nil {
				x.FptrMapUintptrUint16 = new(map[uintptr]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrUint16 != nil {
					x.FptrMapUintptrUint16 = nil
				}
			} else {
				if x.FptrMapUintptrUint16 == nil {
					x.FptrMapUintptrUint16 = new(map[uintptr]uint16)
				}
				yym723 := z.DecBinary()
				_ = yym723
				if false {
				} else {
					z.F.DecMapUintptrUint16X(x.FptrMapUintptrUint16, d)
				}
			}
		case "FMapUintptrUint32":
			if r.TryDecodeAsNil() {
				x.FMapUintptrUint32 = nil
			} else {
				yyv724 := &x.FMapUintptrUint32
				yym725 := z.DecBinary()
				_ = yym725
				if false {
				} else {
					z.F.DecMapUintptrUint32X(yyv724, d)
				}
			}
		case "FptrMapUintptrUint32":
			if x.FptrMapUintptrUint32 == nil {
				x.FptrMapUintptrUint32 = new(map[uintptr]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrUint32 != nil {
					x.FptrMapUintptrUint32 = nil
				}
			} else {
				if x.FptrMapUintptrUint32 == nil {
					x.FptrMapUintptrUint32 = new(map[uintptr]uint32)
				}
				yym727 := z.DecBinary()
				_ = yym727
				if false {
				} else {
					z.F.DecMapUintptrUint32X(x.FptrMapUintptrUint32, d)
				}
			}
		case "FMapUintptrUint64":
			if r.TryDecodeAsNil() {
				x.FMapUintptrUint64 = nil
			} else {
				yyv728 := &x.FMapUintptrUint64
				yym729 := z.DecBinary()
				_ = yym729
				if false {
				} else {
					z.F.DecMapUintptrUint64X(yyv728, d)
				}
			}
		case "FptrMapUintptrUint64":
			if x.FptrMapUintptrUint64 == nil {
				x.FptrMapUintptrUint64 = new(map[uintptr]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrUint64 != nil {
					x.FptrMapUintptrUint64 = nil
				}
			} else {
				if x.FptrMapUintptrUint64 == nil {
					x.FptrMapUintptrUint64 = new(map[uintptr]uint64)
				}
				yym731 := z.DecBinary()
				_ = yym731
				if false {
				} else {
					z.F.DecMapUintptrUint64X(x.FptrMapUintptrUint64, d)
				}
			}
		case "FMapUintptrUintptr":
			if r.TryDecodeAsNil() {
				x.FMapUintptrUintptr = nil
			} else {
				yyv732 := &x.FMapUintptrUintptr
				yym733 := z.DecBinary()
				_ = yym733
				if false {
				} else {
					z.F.DecMapUintptrUintptrX(yyv732, d)
				}
			}
		case "FptrMapUintptrUintptr":
			if x.FptrMapUintptrUintptr == nil {
				x.FptrMapUintptrUintptr = new(map[uintptr]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrUintptr != nil {
					x.FptrMapUintptrUintptr = nil
				}
			} else {
				if x.FptrMapUintptrUintptr == nil {
					x.FptrMapUintptrUintptr = new(map[uintptr]uintptr)
				}
				yym735 := z.DecBinary()
				_ = yym735
				if false {
				} else {
					z.F.DecMapUintptrUintptrX(x.FptrMapUintptrUintptr, d)
				}
			}
		case "FMapUintptrInt":
			if r.TryDecodeAsNil() {
				x.FMapUintptrInt = nil
			} else {
				yyv736 := &x.FMapUintptrInt
				yym737 := z.DecBinary()
				_ = yym737
				if false {
				} else {
					z.F.DecMapUintptrIntX(yyv736, d)
				}
			}
		case "FptrMapUintptrInt":
			if x.FptrMapUintptrInt == nil {
				x.FptrMapUintptrInt = new(map[uintptr]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrInt != nil {
					x.FptrMapUintptrInt = nil
				}
			} else {
				if x.FptrMapUintptrInt == nil {
					x.FptrMapUintptrInt = new(map[uintptr]int)
				}
				yym739 := z.DecBinary()
				_ = yym739
				if false {
				} else {
					z.F.DecMapUintptrIntX(x.FptrMapUintptrInt, d)
				}
			}
		case "FMapUintptrInt8":
			if r.TryDecodeAsNil() {
				x.FMapUintptrInt8 = nil
			} else {
				yyv740 := &x.FMapUintptrInt8
				yym741 := z.DecBinary()
				_ = yym741
				if false {
				} else {
					z.F.DecMapUintptrInt8X(yyv740, d)
				}
			}
		case "FptrMapUintptrInt8":
			if x.FptrMapUintptrInt8 == nil {
				x.FptrMapUintptrInt8 = new(map[uintptr]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrInt8 != nil {
					x.FptrMapUintptrInt8 = nil
				}
			} else {
				if x.FptrMapUintptrInt8 == nil {
					x.FptrMapUintptrInt8 = new(map[uintptr]int8)
				}
				yym743 := z.DecBinary()
				_ = yym743
				if false {
				} else {
					z.F.DecMapUintptrInt8X(x.FptrMapUintptrInt8, d)
				}
			}
		case "FMapUintptrInt16":
			if r.TryDecodeAsNil() {
				x.FMapUintptrInt16 = nil
			} else {
				yyv744 := &x.FMapUintptrInt16
				yym745 := z.DecBinary()
				_ = yym745
				if false {
				} else {
					z.F.DecMapUintptrInt16X(yyv744, d)
				}
			}
		case "FptrMapUintptrInt16":
			if x.FptrMapUintptrInt16 == nil {
				x.FptrMapUintptrInt16 = new(map[uintptr]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrInt16 != nil {
					x.FptrMapUintptrInt16 = nil
				}
			} else {
				if x.FptrMapUintptrInt16 == nil {
					x.FptrMapUintptrInt16 = new(map[uintptr]int16)
				}
				yym747 := z.DecBinary()
				_ = yym747
				if false {
				} else {
					z.F.DecMapUintptrInt16X(x.FptrMapUintptrInt16, d)
				}
			}
		case "FMapUintptrInt32":
			if r.TryDecodeAsNil() {
				x.FMapUintptrInt32 = nil
			} else {
				yyv748 := &x.FMapUintptrInt32
				yym749 := z.DecBinary()
				_ = yym749
				if false {
				} else {
					z.F.DecMapUintptrInt32X(yyv748, d)
				}
			}
		case "FptrMapUintptrInt32":
			if x.FptrMapUintptrInt32 == nil {
				x.FptrMapUintptrInt32 = new(map[uintptr]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrInt32 != nil {
					x.FptrMapUintptrInt32 = nil
				}
			} else {
				if x.FptrMapUintptrInt32 == nil {
					x.FptrMapUintptrInt32 = new(map[uintptr]int32)
				}
				yym751 := z.DecBinary()
				_ = yym751
				if false {
				} else {
					z.F.DecMapUintptrInt32X(x.FptrMapUintptrInt32, d)
				}
			}
		case "FMapUintptrInt64":
			if r.TryDecodeAsNil() {
				x.FMapUintptrInt64 = nil
			} else {
				yyv752 := &x.FMapUintptrInt64
				yym753 := z.DecBinary()
				_ = yym753
				if false {
				} else {
					z.F.DecMapUintptrInt64X(yyv752, d)
				}
			}
		case "FptrMapUintptrInt64":
			if x.FptrMapUintptrInt64 == nil {
				x.FptrMapUintptrInt64 = new(map[uintptr]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrInt64 != nil {
					x.FptrMapUintptrInt64 = nil
				}
			} else {
				if x.FptrMapUintptrInt64 == nil {
					x.FptrMapUintptrInt64 = new(map[uintptr]int64)
				}
				yym755 := z.DecBinary()
				_ = yym755
				if false {
				} else {
					z.F.DecMapUintptrInt64X(x.FptrMapUintptrInt64, d)
				}
			}
		case "FMapUintptrFloat32":
			if r.TryDecodeAsNil() {
				x.FMapUintptrFloat32 = nil
			} else {
				yyv756 := &x.FMapUintptrFloat32
				yym757 := z.DecBinary()
				_ = yym757
				if false {
				} else {
					z.F.DecMapUintptrFloat32X(yyv756, d)
				}
			}
		case "FptrMapUintptrFloat32":
			if x.FptrMapUintptrFloat32 == nil {
				x.FptrMapUintptrFloat32 = new(map[uintptr]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrFloat32 != nil {
					x.FptrMapUintptrFloat32 = nil
				}
			} else {
				if x.FptrMapUintptrFloat32 == nil {
					x.FptrMapUintptrFloat32 = new(map[uintptr]float32)
				}
				yym759 := z.DecBinary()
				_ = yym759
				if false {
				} else {
					z.F.DecMapUintptrFloat32X(x.FptrMapUintptrFloat32, d)
				}
			}
		case "FMapUintptrFloat64":
			if r.TryDecodeAsNil() {
				x.FMapUintptrFloat64 = nil
			} else {
				yyv760 := &x.FMapUintptrFloat64
				yym761 := z.DecBinary()
				_ = yym761
				if false {
				} else {
					z.F.DecMapUintptrFloat64X(yyv760, d)
				}
			}
		case "FptrMapUintptrFloat64":
			if x.FptrMapUintptrFloat64 == nil {
				x.FptrMapUintptrFloat64 = new(map[uintptr]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrFloat64 != nil {
					x.FptrMapUintptrFloat64 = nil
				}
			} else {
				if x.FptrMapUintptrFloat64 == nil {
					x.FptrMapUintptrFloat64 = new(map[uintptr]float64)
				}
				yym763 := z.DecBinary()
				_ = yym763
				if false {
				} else {
					z.F.DecMapUintptrFloat64X(x.FptrMapUintptrFloat64, d)
				}
			}
		case "FMapUintptrBool":
			if r.TryDecodeAsNil() {
				x.FMapUintptrBool = nil
			} else {
				yyv764 := &x.FMapUintptrBool
				yym765 := z.DecBinary()
				_ = yym765
				if false {
				} else {
					z.F.DecMapUintptrBoolX(yyv764, d)
				}
			}
		case "FptrMapUintptrBool":
			if x.FptrMapUintptrBool == nil {
				x.FptrMapUintptrBool = new(map[uintptr]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapUintptrBool != nil {
					x.FptrMapUintptrBool = nil
				}
			} else {
				if x.FptrMapUintptrBool == nil {
					x.FptrMapUintptrBool = new(map[uintptr]bool)
				}
				yym767 := z.DecBinary()
				_ = yym767
				if false {
				} else {
					z.F.DecMapUintptrBoolX(x.FptrMapUintptrBool, d)
				}
			}
		case "FMapIntIntf":
			if r.TryDecodeAsNil() {
				x.FMapIntIntf = nil
			} else {
				yyv768 := &x.FMapIntIntf
				yym769 := z.DecBinary()
				_ = yym769
				if false {
				} else {
					z.F.DecMapIntIntfX(yyv768, d)
				}
			}
		case "FptrMapIntIntf":
			if x.FptrMapIntIntf == nil {
				x.FptrMapIntIntf = new(map[int]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntIntf != nil {
					x.FptrMapIntIntf = nil
				}
			} else {
				if x.FptrMapIntIntf == nil {
					x.FptrMapIntIntf = new(map[int]interface{})
				}
				yym771 := z.DecBinary()
				_ = yym771
				if false {
				} else {
					z.F.DecMapIntIntfX(x.FptrMapIntIntf, d)
				}
			}
		case "FMapIntString":
			if r.TryDecodeAsNil() {
				x.FMapIntString = nil
			} else {
				yyv772 := &x.FMapIntString
				yym773 := z.DecBinary()
				_ = yym773
				if false {
				} else {
					z.F.DecMapIntStringX(yyv772, d)
				}
			}
		case "FptrMapIntString":
			if x.FptrMapIntString == nil {
				x.FptrMapIntString = new(map[int]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntString != nil {
					x.FptrMapIntString = nil
				}
			} else {
				if x.FptrMapIntString == nil {
					x.FptrMapIntString = new(map[int]string)
				}
				yym775 := z.DecBinary()
				_ = yym775
				if false {
				} else {
					z.F.DecMapIntStringX(x.FptrMapIntString, d)
				}
			}
		case "FMapIntUint":
			if r.TryDecodeAsNil() {
				x.FMapIntUint = nil
			} else {
				yyv776 := &x.FMapIntUint
				yym777 := z.DecBinary()
				_ = yym777
				if false {
				} else {
					z.F.DecMapIntUintX(yyv776, d)
				}
			}
		case "FptrMapIntUint":
			if x.FptrMapIntUint == nil {
				x.FptrMapIntUint = new(map[int]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntUint != nil {
					x.FptrMapIntUint = nil
				}
			} else {
				if x.FptrMapIntUint == nil {
					x.FptrMapIntUint = new(map[int]uint)
				}
				yym779 := z.DecBinary()
				_ = yym779
				if false {
				} else {
					z.F.DecMapIntUintX(x.FptrMapIntUint, d)
				}
			}
		case "FMapIntUint8":
			if r.TryDecodeAsNil() {
				x.FMapIntUint8 = nil
			} else {
				yyv780 := &x.FMapIntUint8
				yym781 := z.DecBinary()
				_ = yym781
				if false {
				} else {
					z.F.DecMapIntUint8X(yyv780, d)
				}
			}
		case "FptrMapIntUint8":
			if x.FptrMapIntUint8 == nil {
				x.FptrMapIntUint8 = new(map[int]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntUint8 != nil {
					x.FptrMapIntUint8 = nil
				}
			} else {
				if x.FptrMapIntUint8 == nil {
					x.FptrMapIntUint8 = new(map[int]uint8)
				}
				yym783 := z.DecBinary()
				_ = yym783
				if false {
				} else {
					z.F.DecMapIntUint8X(x.FptrMapIntUint8, d)
				}
			}
		case "FMapIntUint16":
			if r.TryDecodeAsNil() {
				x.FMapIntUint16 = nil
			} else {
				yyv784 := &x.FMapIntUint16
				yym785 := z.DecBinary()
				_ = yym785
				if false {
				} else {
					z.F.DecMapIntUint16X(yyv784, d)
				}
			}
		case "FptrMapIntUint16":
			if x.FptrMapIntUint16 == nil {
				x.FptrMapIntUint16 = new(map[int]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntUint16 != nil {
					x.FptrMapIntUint16 = nil
				}
			} else {
				if x.FptrMapIntUint16 == nil {
					x.FptrMapIntUint16 = new(map[int]uint16)
				}
				yym787 := z.DecBinary()
				_ = yym787
				if false {
				} else {
					z.F.DecMapIntUint16X(x.FptrMapIntUint16, d)
				}
			}
		case "FMapIntUint32":
			if r.TryDecodeAsNil() {
				x.FMapIntUint32 = nil
			} else {
				yyv788 := &x.FMapIntUint32
				yym789 := z.DecBinary()
				_ = yym789
				if false {
				} else {
					z.F.DecMapIntUint32X(yyv788, d)
				}
			}
		case "FptrMapIntUint32":
			if x.FptrMapIntUint32 == nil {
				x.FptrMapIntUint32 = new(map[int]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntUint32 != nil {
					x.FptrMapIntUint32 = nil
				}
			} else {
				if x.FptrMapIntUint32 == nil {
					x.FptrMapIntUint32 = new(map[int]uint32)
				}
				yym791 := z.DecBinary()
				_ = yym791
				if false {
				} else {
					z.F.DecMapIntUint32X(x.FptrMapIntUint32, d)
				}
			}
		case "FMapIntUint64":
			if r.TryDecodeAsNil() {
				x.FMapIntUint64 = nil
			} else {
				yyv792 := &x.FMapIntUint64
				yym793 := z.DecBinary()
				_ = yym793
				if false {
				} else {
					z.F.DecMapIntUint64X(yyv792, d)
				}
			}
		case "FptrMapIntUint64":
			if x.FptrMapIntUint64 == nil {
				x.FptrMapIntUint64 = new(map[int]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntUint64 != nil {
					x.FptrMapIntUint64 = nil
				}
			} else {
				if x.FptrMapIntUint64 == nil {
					x.FptrMapIntUint64 = new(map[int]uint64)
				}
				yym795 := z.DecBinary()
				_ = yym795
				if false {
				} else {
					z.F.DecMapIntUint64X(x.FptrMapIntUint64, d)
				}
			}
		case "FMapIntUintptr":
			if r.TryDecodeAsNil() {
				x.FMapIntUintptr = nil
			} else {
				yyv796 := &x.FMapIntUintptr
				yym797 := z.DecBinary()
				_ = yym797
				if false {
				} else {
					z.F.DecMapIntUintptrX(yyv796, d)
				}
			}
		case "FptrMapIntUintptr":
			if x.FptrMapIntUintptr == nil {
				x.FptrMapIntUintptr = new(map[int]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntUintptr != nil {
					x.FptrMapIntUintptr = nil
				}
			} else {
				if x.FptrMapIntUintptr == nil {
					x.FptrMapIntUintptr = new(map[int]uintptr)
				}
				yym799 := z.DecBinary()
				_ = yym799
				if false {
				} else {
					z.F.DecMapIntUintptrX(x.FptrMapIntUintptr, d)
				}
			}
		case "FMapIntInt":
			if r.TryDecodeAsNil() {
				x.FMapIntInt = nil
			} else {
				yyv800 := &x.FMapIntInt
				yym801 := z.DecBinary()
				_ = yym801
				if false {
				} else {
					z.F.DecMapIntIntX(yyv800, d)
				}
			}
		case "FptrMapIntInt":
			if x.FptrMapIntInt == nil {
				x.FptrMapIntInt = new(map[int]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntInt != nil {
					x.FptrMapIntInt = nil
				}
			} else {
				if x.FptrMapIntInt == nil {
					x.FptrMapIntInt = new(map[int]int)
				}
				yym803 := z.DecBinary()
				_ = yym803
				if false {
				} else {
					z.F.DecMapIntIntX(x.FptrMapIntInt, d)
				}
			}
		case "FMapIntInt8":
			if r.TryDecodeAsNil() {
				x.FMapIntInt8 = nil
			} else {
				yyv804 := &x.FMapIntInt8
				yym805 := z.DecBinary()
				_ = yym805
				if false {
				} else {
					z.F.DecMapIntInt8X(yyv804, d)
				}
			}
		case "FptrMapIntInt8":
			if x.FptrMapIntInt8 == nil {
				x.FptrMapIntInt8 = new(map[int]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntInt8 != nil {
					x.FptrMapIntInt8 = nil
				}
			} else {
				if x.FptrMapIntInt8 == nil {
					x.FptrMapIntInt8 = new(map[int]int8)
				}
				yym807 := z.DecBinary()
				_ = yym807
				if false {
				} else {
					z.F.DecMapIntInt8X(x.FptrMapIntInt8, d)
				}
			}
		case "FMapIntInt16":
			if r.TryDecodeAsNil() {
				x.FMapIntInt16 = nil
			} else {
				yyv808 := &x.FMapIntInt16
				yym809 := z.DecBinary()
				_ = yym809
				if false {
				} else {
					z.F.DecMapIntInt16X(yyv808, d)
				}
			}
		case "FptrMapIntInt16":
			if x.FptrMapIntInt16 == nil {
				x.FptrMapIntInt16 = new(map[int]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntInt16 != nil {
					x.FptrMapIntInt16 = nil
				}
			} else {
				if x.FptrMapIntInt16 == nil {
					x.FptrMapIntInt16 = new(map[int]int16)
				}
				yym811 := z.DecBinary()
				_ = yym811
				if false {
				} else {
					z.F.DecMapIntInt16X(x.FptrMapIntInt16, d)
				}
			}
		case "FMapIntInt32":
			if r.TryDecodeAsNil() {
				x.FMapIntInt32 = nil
			} else {
				yyv812 := &x.FMapIntInt32
				yym813 := z.DecBinary()
				_ = yym813
				if false {
				} else {
					z.F.DecMapIntInt32X(yyv812, d)
				}
			}
		case "FptrMapIntInt32":
			if x.FptrMapIntInt32 == nil {
				x.FptrMapIntInt32 = new(map[int]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntInt32 != nil {
					x.FptrMapIntInt32 = nil
				}
			} else {
				if x.FptrMapIntInt32 == nil {
					x.FptrMapIntInt32 = new(map[int]int32)
				}
				yym815 := z.DecBinary()
				_ = yym815
				if false {
				} else {
					z.F.DecMapIntInt32X(x.FptrMapIntInt32, d)
				}
			}
		case "FMapIntInt64":
			if r.TryDecodeAsNil() {
				x.FMapIntInt64 = nil
			} else {
				yyv816 := &x.FMapIntInt64
				yym817 := z.DecBinary()
				_ = yym817
				if false {
				} else {
					z.F.DecMapIntInt64X(yyv816, d)
				}
			}
		case "FptrMapIntInt64":
			if x.FptrMapIntInt64 == nil {
				x.FptrMapIntInt64 = new(map[int]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntInt64 != nil {
					x.FptrMapIntInt64 = nil
				}
			} else {
				if x.FptrMapIntInt64 == nil {
					x.FptrMapIntInt64 = new(map[int]int64)
				}
				yym819 := z.DecBinary()
				_ = yym819
				if false {
				} else {
					z.F.DecMapIntInt64X(x.FptrMapIntInt64, d)
				}
			}
		case "FMapIntFloat32":
			if r.TryDecodeAsNil() {
				x.FMapIntFloat32 = nil
			} else {
				yyv820 := &x.FMapIntFloat32
				yym821 := z.DecBinary()
				_ = yym821
				if false {
				} else {
					z.F.DecMapIntFloat32X(yyv820, d)
				}
			}
		case "FptrMapIntFloat32":
			if x.FptrMapIntFloat32 == nil {
				x.FptrMapIntFloat32 = new(map[int]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntFloat32 != nil {
					x.FptrMapIntFloat32 = nil
				}
			} else {
				if x.FptrMapIntFloat32 == nil {
					x.FptrMapIntFloat32 = new(map[int]float32)
				}
				yym823 := z.DecBinary()
				_ = yym823
				if false {
				} else {
					z.F.DecMapIntFloat32X(x.FptrMapIntFloat32, d)
				}
			}
		case "FMapIntFloat64":
			if r.TryDecodeAsNil() {
				x.FMapIntFloat64 = nil
			} else {
				yyv824 := &x.FMapIntFloat64
				yym825 := z.DecBinary()
				_ = yym825
				if false {
				} else {
					z.F.DecMapIntFloat64X(yyv824, d)
				}
			}
		case "FptrMapIntFloat64":
			if x.FptrMapIntFloat64 == nil {
				x.FptrMapIntFloat64 = new(map[int]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntFloat64 != nil {
					x.FptrMapIntFloat64 = nil
				}
			} else {
				if x.FptrMapIntFloat64 == nil {
					x.FptrMapIntFloat64 = new(map[int]float64)
				}
				yym827 := z.DecBinary()
				_ = yym827
				if false {
				} else {
					z.F.DecMapIntFloat64X(x.FptrMapIntFloat64, d)
				}
			}
		case "FMapIntBool":
			if r.TryDecodeAsNil() {
				x.FMapIntBool = nil
			} else {
				yyv828 := &x.FMapIntBool
				yym829 := z.DecBinary()
				_ = yym829
				if false {
				} else {
					z.F.DecMapIntBoolX(yyv828, d)
				}
			}
		case "FptrMapIntBool":
			if x.FptrMapIntBool == nil {
				x.FptrMapIntBool = new(map[int]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapIntBool != nil {
					x.FptrMapIntBool = nil
				}
			} else {
				if x.FptrMapIntBool == nil {
					x.FptrMapIntBool = new(map[int]bool)
				}
				yym831 := z.DecBinary()
				_ = yym831
				if false {
				} else {
					z.F.DecMapIntBoolX(x.FptrMapIntBool, d)
				}
			}
		case "FMapInt8Intf":
			if r.TryDecodeAsNil() {
				x.FMapInt8Intf = nil
			} else {
				yyv832 := &x.FMapInt8Intf
				yym833 := z.DecBinary()
				_ = yym833
				if false {
				} else {
					z.F.DecMapInt8IntfX(yyv832, d)
				}
			}
		case "FptrMapInt8Intf":
			if x.FptrMapInt8Intf == nil {
				x.FptrMapInt8Intf = new(map[int8]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Intf != nil {
					x.FptrMapInt8Intf = nil
				}
			} else {
				if x.FptrMapInt8Intf == nil {
					x.FptrMapInt8Intf = new(map[int8]interface{})
				}
				yym835 := z.DecBinary()
				_ = yym835
				if false {
				} else {
					z.F.DecMapInt8IntfX(x.FptrMapInt8Intf, d)
				}
			}
		case "FMapInt8String":
			if r.TryDecodeAsNil() {
				x.FMapInt8String = nil
			} else {
				yyv836 := &x.FMapInt8String
				yym837 := z.DecBinary()
				_ = yym837
				if false {
				} else {
					z.F.DecMapInt8StringX(yyv836, d)
				}
			}
		case "FptrMapInt8String":
			if x.FptrMapInt8String == nil {
				x.FptrMapInt8String = new(map[int8]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8String != nil {
					x.FptrMapInt8String = nil
				}
			} else {
				if x.FptrMapInt8String == nil {
					x.FptrMapInt8String = new(map[int8]string)
				}
				yym839 := z.DecBinary()
				_ = yym839
				if false {
				} else {
					z.F.DecMapInt8StringX(x.FptrMapInt8String, d)
				}
			}
		case "FMapInt8Uint":
			if r.TryDecodeAsNil() {
				x.FMapInt8Uint = nil
			} else {
				yyv840 := &x.FMapInt8Uint
				yym841 := z.DecBinary()
				_ = yym841
				if false {
				} else {
					z.F.DecMapInt8UintX(yyv840, d)
				}
			}
		case "FptrMapInt8Uint":
			if x.FptrMapInt8Uint == nil {
				x.FptrMapInt8Uint = new(map[int8]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Uint != nil {
					x.FptrMapInt8Uint = nil
				}
			} else {
				if x.FptrMapInt8Uint == nil {
					x.FptrMapInt8Uint = new(map[int8]uint)
				}
				yym843 := z.DecBinary()
				_ = yym843
				if false {
				} else {
					z.F.DecMapInt8UintX(x.FptrMapInt8Uint, d)
				}
			}
		case "FMapInt8Uint8":
			if r.TryDecodeAsNil() {
				x.FMapInt8Uint8 = nil
			} else {
				yyv844 := &x.FMapInt8Uint8
				yym845 := z.DecBinary()
				_ = yym845
				if false {
				} else {
					z.F.DecMapInt8Uint8X(yyv844, d)
				}
			}
		case "FptrMapInt8Uint8":
			if x.FptrMapInt8Uint8 == nil {
				x.FptrMapInt8Uint8 = new(map[int8]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Uint8 != nil {
					x.FptrMapInt8Uint8 = nil
				}
			} else {
				if x.FptrMapInt8Uint8 == nil {
					x.FptrMapInt8Uint8 = new(map[int8]uint8)
				}
				yym847 := z.DecBinary()
				_ = yym847
				if false {
				} else {
					z.F.DecMapInt8Uint8X(x.FptrMapInt8Uint8, d)
				}
			}
		case "FMapInt8Uint16":
			if r.TryDecodeAsNil() {
				x.FMapInt8Uint16 = nil
			} else {
				yyv848 := &x.FMapInt8Uint16
				yym849 := z.DecBinary()
				_ = yym849
				if false {
				} else {
					z.F.DecMapInt8Uint16X(yyv848, d)
				}
			}
		case "FptrMapInt8Uint16":
			if x.FptrMapInt8Uint16 == nil {
				x.FptrMapInt8Uint16 = new(map[int8]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Uint16 != nil {
					x.FptrMapInt8Uint16 = nil
				}
			} else {
				if x.FptrMapInt8Uint16 == nil {
					x.FptrMapInt8Uint16 = new(map[int8]uint16)
				}
				yym851 := z.DecBinary()
				_ = yym851
				if false {
				} else {
					z.F.DecMapInt8Uint16X(x.FptrMapInt8Uint16, d)
				}
			}
		case "FMapInt8Uint32":
			if r.TryDecodeAsNil() {
				x.FMapInt8Uint32 = nil
			} else {
				yyv852 := &x.FMapInt8Uint32
				yym853 := z.DecBinary()
				_ = yym853
				if false {
				} else {
					z.F.DecMapInt8Uint32X(yyv852, d)
				}
			}
		case "FptrMapInt8Uint32":
			if x.FptrMapInt8Uint32 == nil {
				x.FptrMapInt8Uint32 = new(map[int8]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Uint32 != nil {
					x.FptrMapInt8Uint32 = nil
				}
			} else {
				if x.FptrMapInt8Uint32 == nil {
					x.FptrMapInt8Uint32 = new(map[int8]uint32)
				}
				yym855 := z.DecBinary()
				_ = yym855
				if false {
				} else {
					z.F.DecMapInt8Uint32X(x.FptrMapInt8Uint32, d)
				}
			}
		case "FMapInt8Uint64":
			if r.TryDecodeAsNil() {
				x.FMapInt8Uint64 = nil
			} else {
				yyv856 := &x.FMapInt8Uint64
				yym857 := z.DecBinary()
				_ = yym857
				if false {
				} else {
					z.F.DecMapInt8Uint64X(yyv856, d)
				}
			}
		case "FptrMapInt8Uint64":
			if x.FptrMapInt8Uint64 == nil {
				x.FptrMapInt8Uint64 = new(map[int8]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Uint64 != nil {
					x.FptrMapInt8Uint64 = nil
				}
			} else {
				if x.FptrMapInt8Uint64 == nil {
					x.FptrMapInt8Uint64 = new(map[int8]uint64)
				}
				yym859 := z.DecBinary()
				_ = yym859
				if false {
				} else {
					z.F.DecMapInt8Uint64X(x.FptrMapInt8Uint64, d)
				}
			}
		case "FMapInt8Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapInt8Uintptr = nil
			} else {
				yyv860 := &x.FMapInt8Uintptr
				yym861 := z.DecBinary()
				_ = yym861
				if false {
				} else {
					z.F.DecMapInt8UintptrX(yyv860, d)
				}
			}
		case "FptrMapInt8Uintptr":
			if x.FptrMapInt8Uintptr == nil {
				x.FptrMapInt8Uintptr = new(map[int8]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Uintptr != nil {
					x.FptrMapInt8Uintptr = nil
				}
			} else {
				if x.FptrMapInt8Uintptr == nil {
					x.FptrMapInt8Uintptr = new(map[int8]uintptr)
				}
				yym863 := z.DecBinary()
				_ = yym863
				if false {
				} else {
					z.F.DecMapInt8UintptrX(x.FptrMapInt8Uintptr, d)
				}
			}
		case "FMapInt8Int":
			if r.TryDecodeAsNil() {
				x.FMapInt8Int = nil
			} else {
				yyv864 := &x.FMapInt8Int
				yym865 := z.DecBinary()
				_ = yym865
				if false {
				} else {
					z.F.DecMapInt8IntX(yyv864, d)
				}
			}
		case "FptrMapInt8Int":
			if x.FptrMapInt8Int == nil {
				x.FptrMapInt8Int = new(map[int8]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Int != nil {
					x.FptrMapInt8Int = nil
				}
			} else {
				if x.FptrMapInt8Int == nil {
					x.FptrMapInt8Int = new(map[int8]int)
				}
				yym867 := z.DecBinary()
				_ = yym867
				if false {
				} else {
					z.F.DecMapInt8IntX(x.FptrMapInt8Int, d)
				}
			}
		case "FMapInt8Int8":
			if r.TryDecodeAsNil() {
				x.FMapInt8Int8 = nil
			} else {
				yyv868 := &x.FMapInt8Int8
				yym869 := z.DecBinary()
				_ = yym869
				if false {
				} else {
					z.F.DecMapInt8Int8X(yyv868, d)
				}
			}
		case "FptrMapInt8Int8":
			if x.FptrMapInt8Int8 == nil {
				x.FptrMapInt8Int8 = new(map[int8]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Int8 != nil {
					x.FptrMapInt8Int8 = nil
				}
			} else {
				if x.FptrMapInt8Int8 == nil {
					x.FptrMapInt8Int8 = new(map[int8]int8)
				}
				yym871 := z.DecBinary()
				_ = yym871
				if false {
				} else {
					z.F.DecMapInt8Int8X(x.FptrMapInt8Int8, d)
				}
			}
		case "FMapInt8Int16":
			if r.TryDecodeAsNil() {
				x.FMapInt8Int16 = nil
			} else {
				yyv872 := &x.FMapInt8Int16
				yym873 := z.DecBinary()
				_ = yym873
				if false {
				} else {
					z.F.DecMapInt8Int16X(yyv872, d)
				}
			}
		case "FptrMapInt8Int16":
			if x.FptrMapInt8Int16 == nil {
				x.FptrMapInt8Int16 = new(map[int8]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Int16 != nil {
					x.FptrMapInt8Int16 = nil
				}
			} else {
				if x.FptrMapInt8Int16 == nil {
					x.FptrMapInt8Int16 = new(map[int8]int16)
				}
				yym875 := z.DecBinary()
				_ = yym875
				if false {
				} else {
					z.F.DecMapInt8Int16X(x.FptrMapInt8Int16, d)
				}
			}
		case "FMapInt8Int32":
			if r.TryDecodeAsNil() {
				x.FMapInt8Int32 = nil
			} else {
				yyv876 := &x.FMapInt8Int32
				yym877 := z.DecBinary()
				_ = yym877
				if false {
				} else {
					z.F.DecMapInt8Int32X(yyv876, d)
				}
			}
		case "FptrMapInt8Int32":
			if x.FptrMapInt8Int32 == nil {
				x.FptrMapInt8Int32 = new(map[int8]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Int32 != nil {
					x.FptrMapInt8Int32 = nil
				}
			} else {
				if x.FptrMapInt8Int32 == nil {
					x.FptrMapInt8Int32 = new(map[int8]int32)
				}
				yym879 := z.DecBinary()
				_ = yym879
				if false {
				} else {
					z.F.DecMapInt8Int32X(x.FptrMapInt8Int32, d)
				}
			}
		case "FMapInt8Int64":
			if r.TryDecodeAsNil() {
				x.FMapInt8Int64 = nil
			} else {
				yyv880 := &x.FMapInt8Int64
				yym881 := z.DecBinary()
				_ = yym881
				if false {
				} else {
					z.F.DecMapInt8Int64X(yyv880, d)
				}
			}
		case "FptrMapInt8Int64":
			if x.FptrMapInt8Int64 == nil {
				x.FptrMapInt8Int64 = new(map[int8]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Int64 != nil {
					x.FptrMapInt8Int64 = nil
				}
			} else {
				if x.FptrMapInt8Int64 == nil {
					x.FptrMapInt8Int64 = new(map[int8]int64)
				}
				yym883 := z.DecBinary()
				_ = yym883
				if false {
				} else {
					z.F.DecMapInt8Int64X(x.FptrMapInt8Int64, d)
				}
			}
		case "FMapInt8Float32":
			if r.TryDecodeAsNil() {
				x.FMapInt8Float32 = nil
			} else {
				yyv884 := &x.FMapInt8Float32
				yym885 := z.DecBinary()
				_ = yym885
				if false {
				} else {
					z.F.DecMapInt8Float32X(yyv884, d)
				}
			}
		case "FptrMapInt8Float32":
			if x.FptrMapInt8Float32 == nil {
				x.FptrMapInt8Float32 = new(map[int8]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Float32 != nil {
					x.FptrMapInt8Float32 = nil
				}
			} else {
				if x.FptrMapInt8Float32 == nil {
					x.FptrMapInt8Float32 = new(map[int8]float32)
				}
				yym887 := z.DecBinary()
				_ = yym887
				if false {
				} else {
					z.F.DecMapInt8Float32X(x.FptrMapInt8Float32, d)
				}
			}
		case "FMapInt8Float64":
			if r.TryDecodeAsNil() {
				x.FMapInt8Float64 = nil
			} else {
				yyv888 := &x.FMapInt8Float64
				yym889 := z.DecBinary()
				_ = yym889
				if false {
				} else {
					z.F.DecMapInt8Float64X(yyv888, d)
				}
			}
		case "FptrMapInt8Float64":
			if x.FptrMapInt8Float64 == nil {
				x.FptrMapInt8Float64 = new(map[int8]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Float64 != nil {
					x.FptrMapInt8Float64 = nil
				}
			} else {
				if x.FptrMapInt8Float64 == nil {
					x.FptrMapInt8Float64 = new(map[int8]float64)
				}
				yym891 := z.DecBinary()
				_ = yym891
				if false {
				} else {
					z.F.DecMapInt8Float64X(x.FptrMapInt8Float64, d)
				}
			}
		case "FMapInt8Bool":
			if r.TryDecodeAsNil() {
				x.FMapInt8Bool = nil
			} else {
				yyv892 := &x.FMapInt8Bool
				yym893 := z.DecBinary()
				_ = yym893
				if false {
				} else {
					z.F.DecMapInt8BoolX(yyv892, d)
				}
			}
		case "FptrMapInt8Bool":
			if x.FptrMapInt8Bool == nil {
				x.FptrMapInt8Bool = new(map[int8]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt8Bool != nil {
					x.FptrMapInt8Bool = nil
				}
			} else {
				if x.FptrMapInt8Bool == nil {
					x.FptrMapInt8Bool = new(map[int8]bool)
				}
				yym895 := z.DecBinary()
				_ = yym895
				if false {
				} else {
					z.F.DecMapInt8BoolX(x.FptrMapInt8Bool, d)
				}
			}
		case "FMapInt16Intf":
			if r.TryDecodeAsNil() {
				x.FMapInt16Intf = nil
			} else {
				yyv896 := &x.FMapInt16Intf
				yym897 := z.DecBinary()
				_ = yym897
				if false {
				} else {
					z.F.DecMapInt16IntfX(yyv896, d)
				}
			}
		case "FptrMapInt16Intf":
			if x.FptrMapInt16Intf == nil {
				x.FptrMapInt16Intf = new(map[int16]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Intf != nil {
					x.FptrMapInt16Intf = nil
				}
			} else {
				if x.FptrMapInt16Intf == nil {
					x.FptrMapInt16Intf = new(map[int16]interface{})
				}
				yym899 := z.DecBinary()
				_ = yym899
				if false {
				} else {
					z.F.DecMapInt16IntfX(x.FptrMapInt16Intf, d)
				}
			}
		case "FMapInt16String":
			if r.TryDecodeAsNil() {
				x.FMapInt16String = nil
			} else {
				yyv900 := &x.FMapInt16String
				yym901 := z.DecBinary()
				_ = yym901
				if false {
				} else {
					z.F.DecMapInt16StringX(yyv900, d)
				}
			}
		case "FptrMapInt16String":
			if x.FptrMapInt16String == nil {
				x.FptrMapInt16String = new(map[int16]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16String != nil {
					x.FptrMapInt16String = nil
				}
			} else {
				if x.FptrMapInt16String == nil {
					x.FptrMapInt16String = new(map[int16]string)
				}
				yym903 := z.DecBinary()
				_ = yym903
				if false {
				} else {
					z.F.DecMapInt16StringX(x.FptrMapInt16String, d)
				}
			}
		case "FMapInt16Uint":
			if r.TryDecodeAsNil() {
				x.FMapInt16Uint = nil
			} else {
				yyv904 := &x.FMapInt16Uint
				yym905 := z.DecBinary()
				_ = yym905
				if false {
				} else {
					z.F.DecMapInt16UintX(yyv904, d)
				}
			}
		case "FptrMapInt16Uint":
			if x.FptrMapInt16Uint == nil {
				x.FptrMapInt16Uint = new(map[int16]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Uint != nil {
					x.FptrMapInt16Uint = nil
				}
			} else {
				if x.FptrMapInt16Uint == nil {
					x.FptrMapInt16Uint = new(map[int16]uint)
				}
				yym907 := z.DecBinary()
				_ = yym907
				if false {
				} else {
					z.F.DecMapInt16UintX(x.FptrMapInt16Uint, d)
				}
			}
		case "FMapInt16Uint8":
			if r.TryDecodeAsNil() {
				x.FMapInt16Uint8 = nil
			} else {
				yyv908 := &x.FMapInt16Uint8
				yym909 := z.DecBinary()
				_ = yym909
				if false {
				} else {
					z.F.DecMapInt16Uint8X(yyv908, d)
				}
			}
		case "FptrMapInt16Uint8":
			if x.FptrMapInt16Uint8 == nil {
				x.FptrMapInt16Uint8 = new(map[int16]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Uint8 != nil {
					x.FptrMapInt16Uint8 = nil
				}
			} else {
				if x.FptrMapInt16Uint8 == nil {
					x.FptrMapInt16Uint8 = new(map[int16]uint8)
				}
				yym911 := z.DecBinary()
				_ = yym911
				if false {
				} else {
					z.F.DecMapInt16Uint8X(x.FptrMapInt16Uint8, d)
				}
			}
		case "FMapInt16Uint16":
			if r.TryDecodeAsNil() {
				x.FMapInt16Uint16 = nil
			} else {
				yyv912 := &x.FMapInt16Uint16
				yym913 := z.DecBinary()
				_ = yym913
				if false {
				} else {
					z.F.DecMapInt16Uint16X(yyv912, d)
				}
			}
		case "FptrMapInt16Uint16":
			if x.FptrMapInt16Uint16 == nil {
				x.FptrMapInt16Uint16 = new(map[int16]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Uint16 != nil {
					x.FptrMapInt16Uint16 = nil
				}
			} else {
				if x.FptrMapInt16Uint16 == nil {
					x.FptrMapInt16Uint16 = new(map[int16]uint16)
				}
				yym915 := z.DecBinary()
				_ = yym915
				if false {
				} else {
					z.F.DecMapInt16Uint16X(x.FptrMapInt16Uint16, d)
				}
			}
		case "FMapInt16Uint32":
			if r.TryDecodeAsNil() {
				x.FMapInt16Uint32 = nil
			} else {
				yyv916 := &x.FMapInt16Uint32
				yym917 := z.DecBinary()
				_ = yym917
				if false {
				} else {
					z.F.DecMapInt16Uint32X(yyv916, d)
				}
			}
		case "FptrMapInt16Uint32":
			if x.FptrMapInt16Uint32 == nil {
				x.FptrMapInt16Uint32 = new(map[int16]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Uint32 != nil {
					x.FptrMapInt16Uint32 = nil
				}
			} else {
				if x.FptrMapInt16Uint32 == nil {
					x.FptrMapInt16Uint32 = new(map[int16]uint32)
				}
				yym919 := z.DecBinary()
				_ = yym919
				if false {
				} else {
					z.F.DecMapInt16Uint32X(x.FptrMapInt16Uint32, d)
				}
			}
		case "FMapInt16Uint64":
			if r.TryDecodeAsNil() {
				x.FMapInt16Uint64 = nil
			} else {
				yyv920 := &x.FMapInt16Uint64
				yym921 := z.DecBinary()
				_ = yym921
				if false {
				} else {
					z.F.DecMapInt16Uint64X(yyv920, d)
				}
			}
		case "FptrMapInt16Uint64":
			if x.FptrMapInt16Uint64 == nil {
				x.FptrMapInt16Uint64 = new(map[int16]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Uint64 != nil {
					x.FptrMapInt16Uint64 = nil
				}
			} else {
				if x.FptrMapInt16Uint64 == nil {
					x.FptrMapInt16Uint64 = new(map[int16]uint64)
				}
				yym923 := z.DecBinary()
				_ = yym923
				if false {
				} else {
					z.F.DecMapInt16Uint64X(x.FptrMapInt16Uint64, d)
				}
			}
		case "FMapInt16Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapInt16Uintptr = nil
			} else {
				yyv924 := &x.FMapInt16Uintptr
				yym925 := z.DecBinary()
				_ = yym925
				if false {
				} else {
					z.F.DecMapInt16UintptrX(yyv924, d)
				}
			}
		case "FptrMapInt16Uintptr":
			if x.FptrMapInt16Uintptr == nil {
				x.FptrMapInt16Uintptr = new(map[int16]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Uintptr != nil {
					x.FptrMapInt16Uintptr = nil
				}
			} else {
				if x.FptrMapInt16Uintptr == nil {
					x.FptrMapInt16Uintptr = new(map[int16]uintptr)
				}
				yym927 := z.DecBinary()
				_ = yym927
				if false {
				} else {
					z.F.DecMapInt16UintptrX(x.FptrMapInt16Uintptr, d)
				}
			}
		case "FMapInt16Int":
			if r.TryDecodeAsNil() {
				x.FMapInt16Int = nil
			} else {
				yyv928 := &x.FMapInt16Int
				yym929 := z.DecBinary()
				_ = yym929
				if false {
				} else {
					z.F.DecMapInt16IntX(yyv928, d)
				}
			}
		case "FptrMapInt16Int":
			if x.FptrMapInt16Int == nil {
				x.FptrMapInt16Int = new(map[int16]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Int != nil {
					x.FptrMapInt16Int = nil
				}
			} else {
				if x.FptrMapInt16Int == nil {
					x.FptrMapInt16Int = new(map[int16]int)
				}
				yym931 := z.DecBinary()
				_ = yym931
				if false {
				} else {
					z.F.DecMapInt16IntX(x.FptrMapInt16Int, d)
				}
			}
		case "FMapInt16Int8":
			if r.TryDecodeAsNil() {
				x.FMapInt16Int8 = nil
			} else {
				yyv932 := &x.FMapInt16Int8
				yym933 := z.DecBinary()
				_ = yym933
				if false {
				} else {
					z.F.DecMapInt16Int8X(yyv932, d)
				}
			}
		case "FptrMapInt16Int8":
			if x.FptrMapInt16Int8 == nil {
				x.FptrMapInt16Int8 = new(map[int16]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Int8 != nil {
					x.FptrMapInt16Int8 = nil
				}
			} else {
				if x.FptrMapInt16Int8 == nil {
					x.FptrMapInt16Int8 = new(map[int16]int8)
				}
				yym935 := z.DecBinary()
				_ = yym935
				if false {
				} else {
					z.F.DecMapInt16Int8X(x.FptrMapInt16Int8, d)
				}
			}
		case "FMapInt16Int16":
			if r.TryDecodeAsNil() {
				x.FMapInt16Int16 = nil
			} else {
				yyv936 := &x.FMapInt16Int16
				yym937 := z.DecBinary()
				_ = yym937
				if false {
				} else {
					z.F.DecMapInt16Int16X(yyv936, d)
				}
			}
		case "FptrMapInt16Int16":
			if x.FptrMapInt16Int16 == nil {
				x.FptrMapInt16Int16 = new(map[int16]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Int16 != nil {
					x.FptrMapInt16Int16 = nil
				}
			} else {
				if x.FptrMapInt16Int16 == nil {
					x.FptrMapInt16Int16 = new(map[int16]int16)
				}
				yym939 := z.DecBinary()
				_ = yym939
				if false {
				} else {
					z.F.DecMapInt16Int16X(x.FptrMapInt16Int16, d)
				}
			}
		case "FMapInt16Int32":
			if r.TryDecodeAsNil() {
				x.FMapInt16Int32 = nil
			} else {
				yyv940 := &x.FMapInt16Int32
				yym941 := z.DecBinary()
				_ = yym941
				if false {
				} else {
					z.F.DecMapInt16Int32X(yyv940, d)
				}
			}
		case "FptrMapInt16Int32":
			if x.FptrMapInt16Int32 == nil {
				x.FptrMapInt16Int32 = new(map[int16]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Int32 != nil {
					x.FptrMapInt16Int32 = nil
				}
			} else {
				if x.FptrMapInt16Int32 == nil {
					x.FptrMapInt16Int32 = new(map[int16]int32)
				}
				yym943 := z.DecBinary()
				_ = yym943
				if false {
				} else {
					z.F.DecMapInt16Int32X(x.FptrMapInt16Int32, d)
				}
			}
		case "FMapInt16Int64":
			if r.TryDecodeAsNil() {
				x.FMapInt16Int64 = nil
			} else {
				yyv944 := &x.FMapInt16Int64
				yym945 := z.DecBinary()
				_ = yym945
				if false {
				} else {
					z.F.DecMapInt16Int64X(yyv944, d)
				}
			}
		case "FptrMapInt16Int64":
			if x.FptrMapInt16Int64 == nil {
				x.FptrMapInt16Int64 = new(map[int16]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Int64 != nil {
					x.FptrMapInt16Int64 = nil
				}
			} else {
				if x.FptrMapInt16Int64 == nil {
					x.FptrMapInt16Int64 = new(map[int16]int64)
				}
				yym947 := z.DecBinary()
				_ = yym947
				if false {
				} else {
					z.F.DecMapInt16Int64X(x.FptrMapInt16Int64, d)
				}
			}
		case "FMapInt16Float32":
			if r.TryDecodeAsNil() {
				x.FMapInt16Float32 = nil
			} else {
				yyv948 := &x.FMapInt16Float32
				yym949 := z.DecBinary()
				_ = yym949
				if false {
				} else {
					z.F.DecMapInt16Float32X(yyv948, d)
				}
			}
		case "FptrMapInt16Float32":
			if x.FptrMapInt16Float32 == nil {
				x.FptrMapInt16Float32 = new(map[int16]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Float32 != nil {
					x.FptrMapInt16Float32 = nil
				}
			} else {
				if x.FptrMapInt16Float32 == nil {
					x.FptrMapInt16Float32 = new(map[int16]float32)
				}
				yym951 := z.DecBinary()
				_ = yym951
				if false {
				} else {
					z.F.DecMapInt16Float32X(x.FptrMapInt16Float32, d)
				}
			}
		case "FMapInt16Float64":
			if r.TryDecodeAsNil() {
				x.FMapInt16Float64 = nil
			} else {
				yyv952 := &x.FMapInt16Float64
				yym953 := z.DecBinary()
				_ = yym953
				if false {
				} else {
					z.F.DecMapInt16Float64X(yyv952, d)
				}
			}
		case "FptrMapInt16Float64":
			if x.FptrMapInt16Float64 == nil {
				x.FptrMapInt16Float64 = new(map[int16]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Float64 != nil {
					x.FptrMapInt16Float64 = nil
				}
			} else {
				if x.FptrMapInt16Float64 == nil {
					x.FptrMapInt16Float64 = new(map[int16]float64)
				}
				yym955 := z.DecBinary()
				_ = yym955
				if false {
				} else {
					z.F.DecMapInt16Float64X(x.FptrMapInt16Float64, d)
				}
			}
		case "FMapInt16Bool":
			if r.TryDecodeAsNil() {
				x.FMapInt16Bool = nil
			} else {
				yyv956 := &x.FMapInt16Bool
				yym957 := z.DecBinary()
				_ = yym957
				if false {
				} else {
					z.F.DecMapInt16BoolX(yyv956, d)
				}
			}
		case "FptrMapInt16Bool":
			if x.FptrMapInt16Bool == nil {
				x.FptrMapInt16Bool = new(map[int16]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt16Bool != nil {
					x.FptrMapInt16Bool = nil
				}
			} else {
				if x.FptrMapInt16Bool == nil {
					x.FptrMapInt16Bool = new(map[int16]bool)
				}
				yym959 := z.DecBinary()
				_ = yym959
				if false {
				} else {
					z.F.DecMapInt16BoolX(x.FptrMapInt16Bool, d)
				}
			}
		case "FMapInt32Intf":
			if r.TryDecodeAsNil() {
				x.FMapInt32Intf = nil
			} else {
				yyv960 := &x.FMapInt32Intf
				yym961 := z.DecBinary()
				_ = yym961
				if false {
				} else {
					z.F.DecMapInt32IntfX(yyv960, d)
				}
			}
		case "FptrMapInt32Intf":
			if x.FptrMapInt32Intf == nil {
				x.FptrMapInt32Intf = new(map[int32]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Intf != nil {
					x.FptrMapInt32Intf = nil
				}
			} else {
				if x.FptrMapInt32Intf == nil {
					x.FptrMapInt32Intf = new(map[int32]interface{})
				}
				yym963 := z.DecBinary()
				_ = yym963
				if false {
				} else {
					z.F.DecMapInt32IntfX(x.FptrMapInt32Intf, d)
				}
			}
		case "FMapInt32String":
			if r.TryDecodeAsNil() {
				x.FMapInt32String = nil
			} else {
				yyv964 := &x.FMapInt32String
				yym965 := z.DecBinary()
				_ = yym965
				if false {
				} else {
					z.F.DecMapInt32StringX(yyv964, d)
				}
			}
		case "FptrMapInt32String":
			if x.FptrMapInt32String == nil {
				x.FptrMapInt32String = new(map[int32]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32String != nil {
					x.FptrMapInt32String = nil
				}
			} else {
				if x.FptrMapInt32String == nil {
					x.FptrMapInt32String = new(map[int32]string)
				}
				yym967 := z.DecBinary()
				_ = yym967
				if false {
				} else {
					z.F.DecMapInt32StringX(x.FptrMapInt32String, d)
				}
			}
		case "FMapInt32Uint":
			if r.TryDecodeAsNil() {
				x.FMapInt32Uint = nil
			} else {
				yyv968 := &x.FMapInt32Uint
				yym969 := z.DecBinary()
				_ = yym969
				if false {
				} else {
					z.F.DecMapInt32UintX(yyv968, d)
				}
			}
		case "FptrMapInt32Uint":
			if x.FptrMapInt32Uint == nil {
				x.FptrMapInt32Uint = new(map[int32]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Uint != nil {
					x.FptrMapInt32Uint = nil
				}
			} else {
				if x.FptrMapInt32Uint == nil {
					x.FptrMapInt32Uint = new(map[int32]uint)
				}
				yym971 := z.DecBinary()
				_ = yym971
				if false {
				} else {
					z.F.DecMapInt32UintX(x.FptrMapInt32Uint, d)
				}
			}
		case "FMapInt32Uint8":
			if r.TryDecodeAsNil() {
				x.FMapInt32Uint8 = nil
			} else {
				yyv972 := &x.FMapInt32Uint8
				yym973 := z.DecBinary()
				_ = yym973
				if false {
				} else {
					z.F.DecMapInt32Uint8X(yyv972, d)
				}
			}
		case "FptrMapInt32Uint8":
			if x.FptrMapInt32Uint8 == nil {
				x.FptrMapInt32Uint8 = new(map[int32]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Uint8 != nil {
					x.FptrMapInt32Uint8 = nil
				}
			} else {
				if x.FptrMapInt32Uint8 == nil {
					x.FptrMapInt32Uint8 = new(map[int32]uint8)
				}
				yym975 := z.DecBinary()
				_ = yym975
				if false {
				} else {
					z.F.DecMapInt32Uint8X(x.FptrMapInt32Uint8, d)
				}
			}
		case "FMapInt32Uint16":
			if r.TryDecodeAsNil() {
				x.FMapInt32Uint16 = nil
			} else {
				yyv976 := &x.FMapInt32Uint16
				yym977 := z.DecBinary()
				_ = yym977
				if false {
				} else {
					z.F.DecMapInt32Uint16X(yyv976, d)
				}
			}
		case "FptrMapInt32Uint16":
			if x.FptrMapInt32Uint16 == nil {
				x.FptrMapInt32Uint16 = new(map[int32]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Uint16 != nil {
					x.FptrMapInt32Uint16 = nil
				}
			} else {
				if x.FptrMapInt32Uint16 == nil {
					x.FptrMapInt32Uint16 = new(map[int32]uint16)
				}
				yym979 := z.DecBinary()
				_ = yym979
				if false {
				} else {
					z.F.DecMapInt32Uint16X(x.FptrMapInt32Uint16, d)
				}
			}
		case "FMapInt32Uint32":
			if r.TryDecodeAsNil() {
				x.FMapInt32Uint32 = nil
			} else {
				yyv980 := &x.FMapInt32Uint32
				yym981 := z.DecBinary()
				_ = yym981
				if false {
				} else {
					z.F.DecMapInt32Uint32X(yyv980, d)
				}
			}
		case "FptrMapInt32Uint32":
			if x.FptrMapInt32Uint32 == nil {
				x.FptrMapInt32Uint32 = new(map[int32]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Uint32 != nil {
					x.FptrMapInt32Uint32 = nil
				}
			} else {
				if x.FptrMapInt32Uint32 == nil {
					x.FptrMapInt32Uint32 = new(map[int32]uint32)
				}
				yym983 := z.DecBinary()
				_ = yym983
				if false {
				} else {
					z.F.DecMapInt32Uint32X(x.FptrMapInt32Uint32, d)
				}
			}
		case "FMapInt32Uint64":
			if r.TryDecodeAsNil() {
				x.FMapInt32Uint64 = nil
			} else {
				yyv984 := &x.FMapInt32Uint64
				yym985 := z.DecBinary()
				_ = yym985
				if false {
				} else {
					z.F.DecMapInt32Uint64X(yyv984, d)
				}
			}
		case "FptrMapInt32Uint64":
			if x.FptrMapInt32Uint64 == nil {
				x.FptrMapInt32Uint64 = new(map[int32]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Uint64 != nil {
					x.FptrMapInt32Uint64 = nil
				}
			} else {
				if x.FptrMapInt32Uint64 == nil {
					x.FptrMapInt32Uint64 = new(map[int32]uint64)
				}
				yym987 := z.DecBinary()
				_ = yym987
				if false {
				} else {
					z.F.DecMapInt32Uint64X(x.FptrMapInt32Uint64, d)
				}
			}
		case "FMapInt32Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapInt32Uintptr = nil
			} else {
				yyv988 := &x.FMapInt32Uintptr
				yym989 := z.DecBinary()
				_ = yym989
				if false {
				} else {
					z.F.DecMapInt32UintptrX(yyv988, d)
				}
			}
		case "FptrMapInt32Uintptr":
			if x.FptrMapInt32Uintptr == nil {
				x.FptrMapInt32Uintptr = new(map[int32]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Uintptr != nil {
					x.FptrMapInt32Uintptr = nil
				}
			} else {
				if x.FptrMapInt32Uintptr == nil {
					x.FptrMapInt32Uintptr = new(map[int32]uintptr)
				}
				yym991 := z.DecBinary()
				_ = yym991
				if false {
				} else {
					z.F.DecMapInt32UintptrX(x.FptrMapInt32Uintptr, d)
				}
			}
		case "FMapInt32Int":
			if r.TryDecodeAsNil() {
				x.FMapInt32Int = nil
			} else {
				yyv992 := &x.FMapInt32Int
				yym993 := z.DecBinary()
				_ = yym993
				if false {
				} else {
					z.F.DecMapInt32IntX(yyv992, d)
				}
			}
		case "FptrMapInt32Int":
			if x.FptrMapInt32Int == nil {
				x.FptrMapInt32Int = new(map[int32]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Int != nil {
					x.FptrMapInt32Int = nil
				}
			} else {
				if x.FptrMapInt32Int == nil {
					x.FptrMapInt32Int = new(map[int32]int)
				}
				yym995 := z.DecBinary()
				_ = yym995
				if false {
				} else {
					z.F.DecMapInt32IntX(x.FptrMapInt32Int, d)
				}
			}
		case "FMapInt32Int8":
			if r.TryDecodeAsNil() {
				x.FMapInt32Int8 = nil
			} else {
				yyv996 := &x.FMapInt32Int8
				yym997 := z.DecBinary()
				_ = yym997
				if false {
				} else {
					z.F.DecMapInt32Int8X(yyv996, d)
				}
			}
		case "FptrMapInt32Int8":
			if x.FptrMapInt32Int8 == nil {
				x.FptrMapInt32Int8 = new(map[int32]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Int8 != nil {
					x.FptrMapInt32Int8 = nil
				}
			} else {
				if x.FptrMapInt32Int8 == nil {
					x.FptrMapInt32Int8 = new(map[int32]int8)
				}
				yym999 := z.DecBinary()
				_ = yym999
				if false {
				} else {
					z.F.DecMapInt32Int8X(x.FptrMapInt32Int8, d)
				}
			}
		case "FMapInt32Int16":
			if r.TryDecodeAsNil() {
				x.FMapInt32Int16 = nil
			} else {
				yyv1000 := &x.FMapInt32Int16
				yym1001 := z.DecBinary()
				_ = yym1001
				if false {
				} else {
					z.F.DecMapInt32Int16X(yyv1000, d)
				}
			}
		case "FptrMapInt32Int16":
			if x.FptrMapInt32Int16 == nil {
				x.FptrMapInt32Int16 = new(map[int32]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Int16 != nil {
					x.FptrMapInt32Int16 = nil
				}
			} else {
				if x.FptrMapInt32Int16 == nil {
					x.FptrMapInt32Int16 = new(map[int32]int16)
				}
				yym1003 := z.DecBinary()
				_ = yym1003
				if false {
				} else {
					z.F.DecMapInt32Int16X(x.FptrMapInt32Int16, d)
				}
			}
		case "FMapInt32Int32":
			if r.TryDecodeAsNil() {
				x.FMapInt32Int32 = nil
			} else {
				yyv1004 := &x.FMapInt32Int32
				yym1005 := z.DecBinary()
				_ = yym1005
				if false {
				} else {
					z.F.DecMapInt32Int32X(yyv1004, d)
				}
			}
		case "FptrMapInt32Int32":
			if x.FptrMapInt32Int32 == nil {
				x.FptrMapInt32Int32 = new(map[int32]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Int32 != nil {
					x.FptrMapInt32Int32 = nil
				}
			} else {
				if x.FptrMapInt32Int32 == nil {
					x.FptrMapInt32Int32 = new(map[int32]int32)
				}
				yym1007 := z.DecBinary()
				_ = yym1007
				if false {
				} else {
					z.F.DecMapInt32Int32X(x.FptrMapInt32Int32, d)
				}
			}
		case "FMapInt32Int64":
			if r.TryDecodeAsNil() {
				x.FMapInt32Int64 = nil
			} else {
				yyv1008 := &x.FMapInt32Int64
				yym1009 := z.DecBinary()
				_ = yym1009
				if false {
				} else {
					z.F.DecMapInt32Int64X(yyv1008, d)
				}
			}
		case "FptrMapInt32Int64":
			if x.FptrMapInt32Int64 == nil {
				x.FptrMapInt32Int64 = new(map[int32]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Int64 != nil {
					x.FptrMapInt32Int64 = nil
				}
			} else {
				if x.FptrMapInt32Int64 == nil {
					x.FptrMapInt32Int64 = new(map[int32]int64)
				}
				yym1011 := z.DecBinary()
				_ = yym1011
				if false {
				} else {
					z.F.DecMapInt32Int64X(x.FptrMapInt32Int64, d)
				}
			}
		case "FMapInt32Float32":
			if r.TryDecodeAsNil() {
				x.FMapInt32Float32 = nil
			} else {
				yyv1012 := &x.FMapInt32Float32
				yym1013 := z.DecBinary()
				_ = yym1013
				if false {
				} else {
					z.F.DecMapInt32Float32X(yyv1012, d)
				}
			}
		case "FptrMapInt32Float32":
			if x.FptrMapInt32Float32 == nil {
				x.FptrMapInt32Float32 = new(map[int32]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Float32 != nil {
					x.FptrMapInt32Float32 = nil
				}
			} else {
				if x.FptrMapInt32Float32 == nil {
					x.FptrMapInt32Float32 = new(map[int32]float32)
				}
				yym1015 := z.DecBinary()
				_ = yym1015
				if false {
				} else {
					z.F.DecMapInt32Float32X(x.FptrMapInt32Float32, d)
				}
			}
		case "FMapInt32Float64":
			if r.TryDecodeAsNil() {
				x.FMapInt32Float64 = nil
			} else {
				yyv1016 := &x.FMapInt32Float64
				yym1017 := z.DecBinary()
				_ = yym1017
				if false {
				} else {
					z.F.DecMapInt32Float64X(yyv1016, d)
				}
			}
		case "FptrMapInt32Float64":
			if x.FptrMapInt32Float64 == nil {
				x.FptrMapInt32Float64 = new(map[int32]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Float64 != nil {
					x.FptrMapInt32Float64 = nil
				}
			} else {
				if x.FptrMapInt32Float64 == nil {
					x.FptrMapInt32Float64 = new(map[int32]float64)
				}
				yym1019 := z.DecBinary()
				_ = yym1019
				if false {
				} else {
					z.F.DecMapInt32Float64X(x.FptrMapInt32Float64, d)
				}
			}
		case "FMapInt32Bool":
			if r.TryDecodeAsNil() {
				x.FMapInt32Bool = nil
			} else {
				yyv1020 := &x.FMapInt32Bool
				yym1021 := z.DecBinary()
				_ = yym1021
				if false {
				} else {
					z.F.DecMapInt32BoolX(yyv1020, d)
				}
			}
		case "FptrMapInt32Bool":
			if x.FptrMapInt32Bool == nil {
				x.FptrMapInt32Bool = new(map[int32]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt32Bool != nil {
					x.FptrMapInt32Bool = nil
				}
			} else {
				if x.FptrMapInt32Bool == nil {
					x.FptrMapInt32Bool = new(map[int32]bool)
				}
				yym1023 := z.DecBinary()
				_ = yym1023
				if false {
				} else {
					z.F.DecMapInt32BoolX(x.FptrMapInt32Bool, d)
				}
			}
		case "FMapInt64Intf":
			if r.TryDecodeAsNil() {
				x.FMapInt64Intf = nil
			} else {
				yyv1024 := &x.FMapInt64Intf
				yym1025 := z.DecBinary()
				_ = yym1025
				if false {
				} else {
					z.F.DecMapInt64IntfX(yyv1024, d)
				}
			}
		case "FptrMapInt64Intf":
			if x.FptrMapInt64Intf == nil {
				x.FptrMapInt64Intf = new(map[int64]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Intf != nil {
					x.FptrMapInt64Intf = nil
				}
			} else {
				if x.FptrMapInt64Intf == nil {
					x.FptrMapInt64Intf = new(map[int64]interface{})
				}
				yym1027 := z.DecBinary()
				_ = yym1027
				if false {
				} else {
					z.F.DecMapInt64IntfX(x.FptrMapInt64Intf, d)
				}
			}
		case "FMapInt64String":
			if r.TryDecodeAsNil() {
				x.FMapInt64String = nil
			} else {
				yyv1028 := &x.FMapInt64String
				yym1029 := z.DecBinary()
				_ = yym1029
				if false {
				} else {
					z.F.DecMapInt64StringX(yyv1028, d)
				}
			}
		case "FptrMapInt64String":
			if x.FptrMapInt64String == nil {
				x.FptrMapInt64String = new(map[int64]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64String != nil {
					x.FptrMapInt64String = nil
				}
			} else {
				if x.FptrMapInt64String == nil {
					x.FptrMapInt64String = new(map[int64]string)
				}
				yym1031 := z.DecBinary()
				_ = yym1031
				if false {
				} else {
					z.F.DecMapInt64StringX(x.FptrMapInt64String, d)
				}
			}
		case "FMapInt64Uint":
			if r.TryDecodeAsNil() {
				x.FMapInt64Uint = nil
			} else {
				yyv1032 := &x.FMapInt64Uint
				yym1033 := z.DecBinary()
				_ = yym1033
				if false {
				} else {
					z.F.DecMapInt64UintX(yyv1032, d)
				}
			}
		case "FptrMapInt64Uint":
			if x.FptrMapInt64Uint == nil {
				x.FptrMapInt64Uint = new(map[int64]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Uint != nil {
					x.FptrMapInt64Uint = nil
				}
			} else {
				if x.FptrMapInt64Uint == nil {
					x.FptrMapInt64Uint = new(map[int64]uint)
				}
				yym1035 := z.DecBinary()
				_ = yym1035
				if false {
				} else {
					z.F.DecMapInt64UintX(x.FptrMapInt64Uint, d)
				}
			}
		case "FMapInt64Uint8":
			if r.TryDecodeAsNil() {
				x.FMapInt64Uint8 = nil
			} else {
				yyv1036 := &x.FMapInt64Uint8
				yym1037 := z.DecBinary()
				_ = yym1037
				if false {
				} else {
					z.F.DecMapInt64Uint8X(yyv1036, d)
				}
			}
		case "FptrMapInt64Uint8":
			if x.FptrMapInt64Uint8 == nil {
				x.FptrMapInt64Uint8 = new(map[int64]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Uint8 != nil {
					x.FptrMapInt64Uint8 = nil
				}
			} else {
				if x.FptrMapInt64Uint8 == nil {
					x.FptrMapInt64Uint8 = new(map[int64]uint8)
				}
				yym1039 := z.DecBinary()
				_ = yym1039
				if false {
				} else {
					z.F.DecMapInt64Uint8X(x.FptrMapInt64Uint8, d)
				}
			}
		case "FMapInt64Uint16":
			if r.TryDecodeAsNil() {
				x.FMapInt64Uint16 = nil
			} else {
				yyv1040 := &x.FMapInt64Uint16
				yym1041 := z.DecBinary()
				_ = yym1041
				if false {
				} else {
					z.F.DecMapInt64Uint16X(yyv1040, d)
				}
			}
		case "FptrMapInt64Uint16":
			if x.FptrMapInt64Uint16 == nil {
				x.FptrMapInt64Uint16 = new(map[int64]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Uint16 != nil {
					x.FptrMapInt64Uint16 = nil
				}
			} else {
				if x.FptrMapInt64Uint16 == nil {
					x.FptrMapInt64Uint16 = new(map[int64]uint16)
				}
				yym1043 := z.DecBinary()
				_ = yym1043
				if false {
				} else {
					z.F.DecMapInt64Uint16X(x.FptrMapInt64Uint16, d)
				}
			}
		case "FMapInt64Uint32":
			if r.TryDecodeAsNil() {
				x.FMapInt64Uint32 = nil
			} else {
				yyv1044 := &x.FMapInt64Uint32
				yym1045 := z.DecBinary()
				_ = yym1045
				if false {
				} else {
					z.F.DecMapInt64Uint32X(yyv1044, d)
				}
			}
		case "FptrMapInt64Uint32":
			if x.FptrMapInt64Uint32 == nil {
				x.FptrMapInt64Uint32 = new(map[int64]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Uint32 != nil {
					x.FptrMapInt64Uint32 = nil
				}
			} else {
				if x.FptrMapInt64Uint32 == nil {
					x.FptrMapInt64Uint32 = new(map[int64]uint32)
				}
				yym1047 := z.DecBinary()
				_ = yym1047
				if false {
				} else {
					z.F.DecMapInt64Uint32X(x.FptrMapInt64Uint32, d)
				}
			}
		case "FMapInt64Uint64":
			if r.TryDecodeAsNil() {
				x.FMapInt64Uint64 = nil
			} else {
				yyv1048 := &x.FMapInt64Uint64
				yym1049 := z.DecBinary()
				_ = yym1049
				if false {
				} else {
					z.F.DecMapInt64Uint64X(yyv1048, d)
				}
			}
		case "FptrMapInt64Uint64":
			if x.FptrMapInt64Uint64 == nil {
				x.FptrMapInt64Uint64 = new(map[int64]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Uint64 != nil {
					x.FptrMapInt64Uint64 = nil
				}
			} else {
				if x.FptrMapInt64Uint64 == nil {
					x.FptrMapInt64Uint64 = new(map[int64]uint64)
				}
				yym1051 := z.DecBinary()
				_ = yym1051
				if false {
				} else {
					z.F.DecMapInt64Uint64X(x.FptrMapInt64Uint64, d)
				}
			}
		case "FMapInt64Uintptr":
			if r.TryDecodeAsNil() {
				x.FMapInt64Uintptr = nil
			} else {
				yyv1052 := &x.FMapInt64Uintptr
				yym1053 := z.DecBinary()
				_ = yym1053
				if false {
				} else {
					z.F.DecMapInt64UintptrX(yyv1052, d)
				}
			}
		case "FptrMapInt64Uintptr":
			if x.FptrMapInt64Uintptr == nil {
				x.FptrMapInt64Uintptr = new(map[int64]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Uintptr != nil {
					x.FptrMapInt64Uintptr = nil
				}
			} else {
				if x.FptrMapInt64Uintptr == nil {
					x.FptrMapInt64Uintptr = new(map[int64]uintptr)
				}
				yym1055 := z.DecBinary()
				_ = yym1055
				if false {
				} else {
					z.F.DecMapInt64UintptrX(x.FptrMapInt64Uintptr, d)
				}
			}
		case "FMapInt64Int":
			if r.TryDecodeAsNil() {
				x.FMapInt64Int = nil
			} else {
				yyv1056 := &x.FMapInt64Int
				yym1057 := z.DecBinary()
				_ = yym1057
				if false {
				} else {
					z.F.DecMapInt64IntX(yyv1056, d)
				}
			}
		case "FptrMapInt64Int":
			if x.FptrMapInt64Int == nil {
				x.FptrMapInt64Int = new(map[int64]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Int != nil {
					x.FptrMapInt64Int = nil
				}
			} else {
				if x.FptrMapInt64Int == nil {
					x.FptrMapInt64Int = new(map[int64]int)
				}
				yym1059 := z.DecBinary()
				_ = yym1059
				if false {
				} else {
					z.F.DecMapInt64IntX(x.FptrMapInt64Int, d)
				}
			}
		case "FMapInt64Int8":
			if r.TryDecodeAsNil() {
				x.FMapInt64Int8 = nil
			} else {
				yyv1060 := &x.FMapInt64Int8
				yym1061 := z.DecBinary()
				_ = yym1061
				if false {
				} else {
					z.F.DecMapInt64Int8X(yyv1060, d)
				}
			}
		case "FptrMapInt64Int8":
			if x.FptrMapInt64Int8 == nil {
				x.FptrMapInt64Int8 = new(map[int64]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Int8 != nil {
					x.FptrMapInt64Int8 = nil
				}
			} else {
				if x.FptrMapInt64Int8 == nil {
					x.FptrMapInt64Int8 = new(map[int64]int8)
				}
				yym1063 := z.DecBinary()
				_ = yym1063
				if false {
				} else {
					z.F.DecMapInt64Int8X(x.FptrMapInt64Int8, d)
				}
			}
		case "FMapInt64Int16":
			if r.TryDecodeAsNil() {
				x.FMapInt64Int16 = nil
			} else {
				yyv1064 := &x.FMapInt64Int16
				yym1065 := z.DecBinary()
				_ = yym1065
				if false {
				} else {
					z.F.DecMapInt64Int16X(yyv1064, d)
				}
			}
		case "FptrMapInt64Int16":
			if x.FptrMapInt64Int16 == nil {
				x.FptrMapInt64Int16 = new(map[int64]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Int16 != nil {
					x.FptrMapInt64Int16 = nil
				}
			} else {
				if x.FptrMapInt64Int16 == nil {
					x.FptrMapInt64Int16 = new(map[int64]int16)
				}
				yym1067 := z.DecBinary()
				_ = yym1067
				if false {
				} else {
					z.F.DecMapInt64Int16X(x.FptrMapInt64Int16, d)
				}
			}
		case "FMapInt64Int32":
			if r.TryDecodeAsNil() {
				x.FMapInt64Int32 = nil
			} else {
				yyv1068 := &x.FMapInt64Int32
				yym1069 := z.DecBinary()
				_ = yym1069
				if false {
				} else {
					z.F.DecMapInt64Int32X(yyv1068, d)
				}
			}
		case "FptrMapInt64Int32":
			if x.FptrMapInt64Int32 == nil {
				x.FptrMapInt64Int32 = new(map[int64]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Int32 != nil {
					x.FptrMapInt64Int32 = nil
				}
			} else {
				if x.FptrMapInt64Int32 == nil {
					x.FptrMapInt64Int32 = new(map[int64]int32)
				}
				yym1071 := z.DecBinary()
				_ = yym1071
				if false {
				} else {
					z.F.DecMapInt64Int32X(x.FptrMapInt64Int32, d)
				}
			}
		case "FMapInt64Int64":
			if r.TryDecodeAsNil() {
				x.FMapInt64Int64 = nil
			} else {
				yyv1072 := &x.FMapInt64Int64
				yym1073 := z.DecBinary()
				_ = yym1073
				if false {
				} else {
					z.F.DecMapInt64Int64X(yyv1072, d)
				}
			}
		case "FptrMapInt64Int64":
			if x.FptrMapInt64Int64 == nil {
				x.FptrMapInt64Int64 = new(map[int64]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Int64 != nil {
					x.FptrMapInt64Int64 = nil
				}
			} else {
				if x.FptrMapInt64Int64 == nil {
					x.FptrMapInt64Int64 = new(map[int64]int64)
				}
				yym1075 := z.DecBinary()
				_ = yym1075
				if false {
				} else {
					z.F.DecMapInt64Int64X(x.FptrMapInt64Int64, d)
				}
			}
		case "FMapInt64Float32":
			if r.TryDecodeAsNil() {
				x.FMapInt64Float32 = nil
			} else {
				yyv1076 := &x.FMapInt64Float32
				yym1077 := z.DecBinary()
				_ = yym1077
				if false {
				} else {
					z.F.DecMapInt64Float32X(yyv1076, d)
				}
			}
		case "FptrMapInt64Float32":
			if x.FptrMapInt64Float32 == nil {
				x.FptrMapInt64Float32 = new(map[int64]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Float32 != nil {
					x.FptrMapInt64Float32 = nil
				}
			} else {
				if x.FptrMapInt64Float32 == nil {
					x.FptrMapInt64Float32 = new(map[int64]float32)
				}
				yym1079 := z.DecBinary()
				_ = yym1079
				if false {
				} else {
					z.F.DecMapInt64Float32X(x.FptrMapInt64Float32, d)
				}
			}
		case "FMapInt64Float64":
			if r.TryDecodeAsNil() {
				x.FMapInt64Float64 = nil
			} else {
				yyv1080 := &x.FMapInt64Float64
				yym1081 := z.DecBinary()
				_ = yym1081
				if false {
				} else {
					z.F.DecMapInt64Float64X(yyv1080, d)
				}
			}
		case "FptrMapInt64Float64":
			if x.FptrMapInt64Float64 == nil {
				x.FptrMapInt64Float64 = new(map[int64]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Float64 != nil {
					x.FptrMapInt64Float64 = nil
				}
			} else {
				if x.FptrMapInt64Float64 == nil {
					x.FptrMapInt64Float64 = new(map[int64]float64)
				}
				yym1083 := z.DecBinary()
				_ = yym1083
				if false {
				} else {
					z.F.DecMapInt64Float64X(x.FptrMapInt64Float64, d)
				}
			}
		case "FMapInt64Bool":
			if r.TryDecodeAsNil() {
				x.FMapInt64Bool = nil
			} else {
				yyv1084 := &x.FMapInt64Bool
				yym1085 := z.DecBinary()
				_ = yym1085
				if false {
				} else {
					z.F.DecMapInt64BoolX(yyv1084, d)
				}
			}
		case "FptrMapInt64Bool":
			if x.FptrMapInt64Bool == nil {
				x.FptrMapInt64Bool = new(map[int64]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapInt64Bool != nil {
					x.FptrMapInt64Bool = nil
				}
			} else {
				if x.FptrMapInt64Bool == nil {
					x.FptrMapInt64Bool = new(map[int64]bool)
				}
				yym1087 := z.DecBinary()
				_ = yym1087
				if false {
				} else {
					z.F.DecMapInt64BoolX(x.FptrMapInt64Bool, d)
				}
			}
		case "FMapBoolIntf":
			if r.TryDecodeAsNil() {
				x.FMapBoolIntf = nil
			} else {
				yyv1088 := &x.FMapBoolIntf
				yym1089 := z.DecBinary()
				_ = yym1089
				if false {
				} else {
					z.F.DecMapBoolIntfX(yyv1088, d)
				}
			}
		case "FptrMapBoolIntf":
			if x.FptrMapBoolIntf == nil {
				x.FptrMapBoolIntf = new(map[bool]interface{})
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolIntf != nil {
					x.FptrMapBoolIntf = nil
				}
			} else {
				if x.FptrMapBoolIntf == nil {
					x.FptrMapBoolIntf = new(map[bool]interface{})
				}
				yym1091 := z.DecBinary()
				_ = yym1091
				if false {
				} else {
					z.F.DecMapBoolIntfX(x.FptrMapBoolIntf, d)
				}
			}
		case "FMapBoolString":
			if r.TryDecodeAsNil() {
				x.FMapBoolString = nil
			} else {
				yyv1092 := &x.FMapBoolString
				yym1093 := z.DecBinary()
				_ = yym1093
				if false {
				} else {
					z.F.DecMapBoolStringX(yyv1092, d)
				}
			}
		case "FptrMapBoolString":
			if x.FptrMapBoolString == nil {
				x.FptrMapBoolString = new(map[bool]string)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolString != nil {
					x.FptrMapBoolString = nil
				}
			} else {
				if x.FptrMapBoolString == nil {
					x.FptrMapBoolString = new(map[bool]string)
				}
				yym1095 := z.DecBinary()
				_ = yym1095
				if false {
				} else {
					z.F.DecMapBoolStringX(x.FptrMapBoolString, d)
				}
			}
		case "FMapBoolUint":
			if r.TryDecodeAsNil() {
				x.FMapBoolUint = nil
			} else {
				yyv1096 := &x.FMapBoolUint
				yym1097 := z.DecBinary()
				_ = yym1097
				if false {
				} else {
					z.F.DecMapBoolUintX(yyv1096, d)
				}
			}
		case "FptrMapBoolUint":
			if x.FptrMapBoolUint == nil {
				x.FptrMapBoolUint = new(map[bool]uint)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolUint != nil {
					x.FptrMapBoolUint = nil
				}
			} else {
				if x.FptrMapBoolUint == nil {
					x.FptrMapBoolUint = new(map[bool]uint)
				}
				yym1099 := z.DecBinary()
				_ = yym1099
				if false {
				} else {
					z.F.DecMapBoolUintX(x.FptrMapBoolUint, d)
				}
			}
		case "FMapBoolUint8":
			if r.TryDecodeAsNil() {
				x.FMapBoolUint8 = nil
			} else {
				yyv1100 := &x.FMapBoolUint8
				yym1101 := z.DecBinary()
				_ = yym1101
				if false {
				} else {
					z.F.DecMapBoolUint8X(yyv1100, d)
				}
			}
		case "FptrMapBoolUint8":
			if x.FptrMapBoolUint8 == nil {
				x.FptrMapBoolUint8 = new(map[bool]uint8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolUint8 != nil {
					x.FptrMapBoolUint8 = nil
				}
			} else {
				if x.FptrMapBoolUint8 == nil {
					x.FptrMapBoolUint8 = new(map[bool]uint8)
				}
				yym1103 := z.DecBinary()
				_ = yym1103
				if false {
				} else {
					z.F.DecMapBoolUint8X(x.FptrMapBoolUint8, d)
				}
			}
		case "FMapBoolUint16":
			if r.TryDecodeAsNil() {
				x.FMapBoolUint16 = nil
			} else {
				yyv1104 := &x.FMapBoolUint16
				yym1105 := z.DecBinary()
				_ = yym1105
				if false {
				} else {
					z.F.DecMapBoolUint16X(yyv1104, d)
				}
			}
		case "FptrMapBoolUint16":
			if x.FptrMapBoolUint16 == nil {
				x.FptrMapBoolUint16 = new(map[bool]uint16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolUint16 != nil {
					x.FptrMapBoolUint16 = nil
				}
			} else {
				if x.FptrMapBoolUint16 == nil {
					x.FptrMapBoolUint16 = new(map[bool]uint16)
				}
				yym1107 := z.DecBinary()
				_ = yym1107
				if false {
				} else {
					z.F.DecMapBoolUint16X(x.FptrMapBoolUint16, d)
				}
			}
		case "FMapBoolUint32":
			if r.TryDecodeAsNil() {
				x.FMapBoolUint32 = nil
			} else {
				yyv1108 := &x.FMapBoolUint32
				yym1109 := z.DecBinary()
				_ = yym1109
				if false {
				} else {
					z.F.DecMapBoolUint32X(yyv1108, d)
				}
			}
		case "FptrMapBoolUint32":
			if x.FptrMapBoolUint32 == nil {
				x.FptrMapBoolUint32 = new(map[bool]uint32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolUint32 != nil {
					x.FptrMapBoolUint32 = nil
				}
			} else {
				if x.FptrMapBoolUint32 == nil {
					x.FptrMapBoolUint32 = new(map[bool]uint32)
				}
				yym1111 := z.DecBinary()
				_ = yym1111
				if false {
				} else {
					z.F.DecMapBoolUint32X(x.FptrMapBoolUint32, d)
				}
			}
		case "FMapBoolUint64":
			if r.TryDecodeAsNil() {
				x.FMapBoolUint64 = nil
			} else {
				yyv1112 := &x.FMapBoolUint64
				yym1113 := z.DecBinary()
				_ = yym1113
				if false {
				} else {
					z.F.DecMapBoolUint64X(yyv1112, d)
				}
			}
		case "FptrMapBoolUint64":
			if x.FptrMapBoolUint64 == nil {
				x.FptrMapBoolUint64 = new(map[bool]uint64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolUint64 != nil {
					x.FptrMapBoolUint64 = nil
				}
			} else {
				if x.FptrMapBoolUint64 == nil {
					x.FptrMapBoolUint64 = new(map[bool]uint64)
				}
				yym1115 := z.DecBinary()
				_ = yym1115
				if false {
				} else {
					z.F.DecMapBoolUint64X(x.FptrMapBoolUint64, d)
				}
			}
		case "FMapBoolUintptr":
			if r.TryDecodeAsNil() {
				x.FMapBoolUintptr = nil
			} else {
				yyv1116 := &x.FMapBoolUintptr
				yym1117 := z.DecBinary()
				_ = yym1117
				if false {
				} else {
					z.F.DecMapBoolUintptrX(yyv1116, d)
				}
			}
		case "FptrMapBoolUintptr":
			if x.FptrMapBoolUintptr == nil {
				x.FptrMapBoolUintptr = new(map[bool]uintptr)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolUintptr != nil {
					x.FptrMapBoolUintptr = nil
				}
			} else {
				if x.FptrMapBoolUintptr == nil {
					x.FptrMapBoolUintptr = new(map[bool]uintptr)
				}
				yym1119 := z.DecBinary()
				_ = yym1119
				if false {
				} else {
					z.F.DecMapBoolUintptrX(x.FptrMapBoolUintptr, d)
				}
			}
		case "FMapBoolInt":
			if r.TryDecodeAsNil() {
				x.FMapBoolInt = nil
			} else {
				yyv1120 := &x.FMapBoolInt
				yym1121 := z.DecBinary()
				_ = yym1121
				if false {
				} else {
					z.F.DecMapBoolIntX(yyv1120, d)
				}
			}
		case "FptrMapBoolInt":
			if x.FptrMapBoolInt == nil {
				x.FptrMapBoolInt = new(map[bool]int)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolInt != nil {
					x.FptrMapBoolInt = nil
				}
			} else {
				if x.FptrMapBoolInt == nil {
					x.FptrMapBoolInt = new(map[bool]int)
				}
				yym1123 := z.DecBinary()
				_ = yym1123
				if false {
				} else {
					z.F.DecMapBoolIntX(x.FptrMapBoolInt, d)
				}
			}
		case "FMapBoolInt8":
			if r.TryDecodeAsNil() {
				x.FMapBoolInt8 = nil
			} else {
				yyv1124 := &x.FMapBoolInt8
				yym1125 := z.DecBinary()
				_ = yym1125
				if false {
				} else {
					z.F.DecMapBoolInt8X(yyv1124, d)
				}
			}
		case "FptrMapBoolInt8":
			if x.FptrMapBoolInt8 == nil {
				x.FptrMapBoolInt8 = new(map[bool]int8)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolInt8 != nil {
					x.FptrMapBoolInt8 = nil
				}
			} else {
				if x.FptrMapBoolInt8 == nil {
					x.FptrMapBoolInt8 = new(map[bool]int8)
				}
				yym1127 := z.DecBinary()
				_ = yym1127
				if false {
				} else {
					z.F.DecMapBoolInt8X(x.FptrMapBoolInt8, d)
				}
			}
		case "FMapBoolInt16":
			if r.TryDecodeAsNil() {
				x.FMapBoolInt16 = nil
			} else {
				yyv1128 := &x.FMapBoolInt16
				yym1129 := z.DecBinary()
				_ = yym1129
				if false {
				} else {
					z.F.DecMapBoolInt16X(yyv1128, d)
				}
			}
		case "FptrMapBoolInt16":
			if x.FptrMapBoolInt16 == nil {
				x.FptrMapBoolInt16 = new(map[bool]int16)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolInt16 != nil {
					x.FptrMapBoolInt16 = nil
				}
			} else {
				if x.FptrMapBoolInt16 == nil {
					x.FptrMapBoolInt16 = new(map[bool]int16)
				}
				yym1131 := z.DecBinary()
				_ = yym1131
				if false {
				} else {
					z.F.DecMapBoolInt16X(x.FptrMapBoolInt16, d)
				}
			}
		case "FMapBoolInt32":
			if r.TryDecodeAsNil() {
				x.FMapBoolInt32 = nil
			} else {
				yyv1132 := &x.FMapBoolInt32
				yym1133 := z.DecBinary()
				_ = yym1133
				if false {
				} else {
					z.F.DecMapBoolInt32X(yyv1132, d)
				}
			}
		case "FptrMapBoolInt32":
			if x.FptrMapBoolInt32 == nil {
				x.FptrMapBoolInt32 = new(map[bool]int32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolInt32 != nil {
					x.FptrMapBoolInt32 = nil
				}
			} else {
				if x.FptrMapBoolInt32 == nil {
					x.FptrMapBoolInt32 = new(map[bool]int32)
				}
				yym1135 := z.DecBinary()
				_ = yym1135
				if false {
				} else {
					z.F.DecMapBoolInt32X(x.FptrMapBoolInt32, d)
				}
			}
		case "FMapBoolInt64":
			if r.TryDecodeAsNil() {
				x.FMapBoolInt64 = nil
			} else {
				yyv1136 := &x.FMapBoolInt64
				yym1137 := z.DecBinary()
				_ = yym1137
				if false {
				} else {
					z.F.DecMapBoolInt64X(yyv1136, d)
				}
			}
		case "FptrMapBoolInt64":
			if x.FptrMapBoolInt64 == nil {
				x.FptrMapBoolInt64 = new(map[bool]int64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolInt64 != nil {
					x.FptrMapBoolInt64 = nil
				}
			} else {
				if x.FptrMapBoolInt64 == nil {
					x.FptrMapBoolInt64 = new(map[bool]int64)
				}
				yym1139 := z.DecBinary()
				_ = yym1139
				if false {
				} else {
					z.F.DecMapBoolInt64X(x.FptrMapBoolInt64, d)
				}
			}
		case "FMapBoolFloat32":
			if r.TryDecodeAsNil() {
				x.FMapBoolFloat32 = nil
			} else {
				yyv1140 := &x.FMapBoolFloat32
				yym1141 := z.DecBinary()
				_ = yym1141
				if false {
				} else {
					z.F.DecMapBoolFloat32X(yyv1140, d)
				}
			}
		case "FptrMapBoolFloat32":
			if x.FptrMapBoolFloat32 == nil {
				x.FptrMapBoolFloat32 = new(map[bool]float32)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolFloat32 != nil {
					x.FptrMapBoolFloat32 = nil
				}
			} else {
				if x.FptrMapBoolFloat32 == nil {
					x.FptrMapBoolFloat32 = new(map[bool]float32)
				}
				yym1143 := z.DecBinary()
				_ = yym1143
				if false {
				} else {
					z.F.DecMapBoolFloat32X(x.FptrMapBoolFloat32, d)
				}
			}
		case "FMapBoolFloat64":
			if r.TryDecodeAsNil() {
				x.FMapBoolFloat64 = nil
			} else {
				yyv1144 := &x.FMapBoolFloat64
				yym1145 := z.DecBinary()
				_ = yym1145
				if false {
				} else {
					z.F.DecMapBoolFloat64X(yyv1144, d)
				}
			}
		case "FptrMapBoolFloat64":
			if x.FptrMapBoolFloat64 == nil {
				x.FptrMapBoolFloat64 = new(map[bool]float64)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolFloat64 != nil {
					x.FptrMapBoolFloat64 = nil
				}
			} else {
				if x.FptrMapBoolFloat64 == nil {
					x.FptrMapBoolFloat64 = new(map[bool]float64)
				}
				yym1147 := z.DecBinary()
				_ = yym1147
				if false {
				} else {
					z.F.DecMapBoolFloat64X(x.FptrMapBoolFloat64, d)
				}
			}
		case "FMapBoolBool":
			if r.TryDecodeAsNil() {
				x.FMapBoolBool = nil
			} else {
				yyv1148 := &x.FMapBoolBool
				yym1149 := z.DecBinary()
				_ = yym1149
				if false {
				} else {
					z.F.DecMapBoolBoolX(yyv1148, d)
				}
			}
		case "FptrMapBoolBool":
			if x.FptrMapBoolBool == nil {
				x.FptrMapBoolBool = new(map[bool]bool)
			}
			if r.TryDecodeAsNil() {
				if x.FptrMapBoolBool != nil {
					x.FptrMapBoolBool = nil
				}
			} else {
				if x.FptrMapBoolBool == nil {
					x.FptrMapBoolBool = new(map[bool]bool)
				}
				yym1151 := z.DecBinary()
				_ = yym1151
				if false {
				} else {
					z.F.DecMapBoolBoolX(x.FptrMapBoolBool, d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *TestMammoth2) codecDecodeSelfFromArray(l int, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj1152 int
	var yyb1152 bool
	var yyhl1152 bool = l >= 0
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FIntf = nil
	} else {
		yyv1153 := &x.FIntf
		yym1154 := z.DecBinary()
		_ = yym1154
		if false {
		} else {
			z.DecFallback(yyv1153, true)
		}
	}
	if x.FptrIntf == nil {
		x.FptrIntf = new(interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrIntf != nil {
			x.FptrIntf = nil
		}
	} else {
		if x.FptrIntf == nil {
			x.FptrIntf = new(interface{})
		}
		yym1156 := z.DecBinary()
		_ = yym1156
		if false {
		} else {
			z.DecFallback(x.FptrIntf, true)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FString = ""
	} else {
		yyv1157 := &x.FString
		yym1158 := z.DecBinary()
		_ = yym1158
		if false {
		} else {
			*((*string)(yyv1157)) = r.DecodeString()
		}
	}
	if x.FptrString == nil {
		x.FptrString = new(string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrString != nil {
			x.FptrString = nil
		}
	} else {
		if x.FptrString == nil {
			x.FptrString = new(string)
		}
		yym1160 := z.DecBinary()
		_ = yym1160
		if false {
		} else {
			*((*string)(x.FptrString)) = r.DecodeString()
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FFloat32 = 0
	} else {
		yyv1161 := &x.FFloat32
		yym1162 := z.DecBinary()
		_ = yym1162
		if false {
		} else {
			*((*float32)(yyv1161)) = float32(r.DecodeFloat(true))
		}
	}
	if x.FptrFloat32 == nil {
		x.FptrFloat32 = new(float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrFloat32 != nil {
			x.FptrFloat32 = nil
		}
	} else {
		if x.FptrFloat32 == nil {
			x.FptrFloat32 = new(float32)
		}
		yym1164 := z.DecBinary()
		_ = yym1164
		if false {
		} else {
			*((*float32)(x.FptrFloat32)) = float32(r.DecodeFloat(true))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FFloat64 = 0
	} else {
		yyv1165 := &x.FFloat64
		yym1166 := z.DecBinary()
		_ = yym1166
		if false {
		} else {
			*((*float64)(yyv1165)) = float64(r.DecodeFloat(false))
		}
	}
	if x.FptrFloat64 == nil {
		x.FptrFloat64 = new(float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrFloat64 != nil {
			x.FptrFloat64 = nil
		}
	} else {
		if x.FptrFloat64 == nil {
			x.FptrFloat64 = new(float64)
		}
		yym1168 := z.DecBinary()
		_ = yym1168
		if false {
		} else {
			*((*float64)(x.FptrFloat64)) = float64(r.DecodeFloat(false))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FUint = 0
	} else {
		yyv1169 := &x.FUint
		yym1170 := z.DecBinary()
		_ = yym1170
		if false {
		} else {
			*((*uint)(yyv1169)) = uint(r.DecodeUint(codecSelferBitsize19781))
		}
	}
	if x.FptrUint == nil {
		x.FptrUint = new(uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrUint != nil {
			x.FptrUint = nil
		}
	} else {
		if x.FptrUint == nil {
			x.FptrUint = new(uint)
		}
		yym1172 := z.DecBinary()
		_ = yym1172
		if false {
		} else {
			*((*uint)(x.FptrUint)) = uint(r.DecodeUint(codecSelferBitsize19781))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FUint8 = 0
	} else {
		yyv1173 := &x.FUint8
		yym1174 := z.DecBinary()
		_ = yym1174
		if false {
		} else {
			*((*uint8)(yyv1173)) = uint8(r.DecodeUint(8))
		}
	}
	if x.FptrUint8 == nil {
		x.FptrUint8 = new(uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrUint8 != nil {
			x.FptrUint8 = nil
		}
	} else {
		if x.FptrUint8 == nil {
			x.FptrUint8 = new(uint8)
		}
		yym1176 := z.DecBinary()
		_ = yym1176
		if false {
		} else {
			*((*uint8)(x.FptrUint8)) = uint8(r.DecodeUint(8))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FUint16 = 0
	} else {
		yyv1177 := &x.FUint16
		yym1178 := z.DecBinary()
		_ = yym1178
		if false {
		} else {
			*((*uint16)(yyv1177)) = uint16(r.DecodeUint(16))
		}
	}
	if x.FptrUint16 == nil {
		x.FptrUint16 = new(uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrUint16 != nil {
			x.FptrUint16 = nil
		}
	} else {
		if x.FptrUint16 == nil {
			x.FptrUint16 = new(uint16)
		}
		yym1180 := z.DecBinary()
		_ = yym1180
		if false {
		} else {
			*((*uint16)(x.FptrUint16)) = uint16(r.DecodeUint(16))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FUint32 = 0
	} else {
		yyv1181 := &x.FUint32
		yym1182 := z.DecBinary()
		_ = yym1182
		if false {
		} else {
			*((*uint32)(yyv1181)) = uint32(r.DecodeUint(32))
		}
	}
	if x.FptrUint32 == nil {
		x.FptrUint32 = new(uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrUint32 != nil {
			x.FptrUint32 = nil
		}
	} else {
		if x.FptrUint32 == nil {
			x.FptrUint32 = new(uint32)
		}
		yym1184 := z.DecBinary()
		_ = yym1184
		if false {
		} else {
			*((*uint32)(x.FptrUint32)) = uint32(r.DecodeUint(32))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FUint64 = 0
	} else {
		yyv1185 := &x.FUint64
		yym1186 := z.DecBinary()
		_ = yym1186
		if false {
		} else {
			*((*uint64)(yyv1185)) = uint64(r.DecodeUint(64))
		}
	}
	if x.FptrUint64 == nil {
		x.FptrUint64 = new(uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrUint64 != nil {
			x.FptrUint64 = nil
		}
	} else {
		if x.FptrUint64 == nil {
			x.FptrUint64 = new(uint64)
		}
		yym1188 := z.DecBinary()
		_ = yym1188
		if false {
		} else {
			*((*uint64)(x.FptrUint64)) = uint64(r.DecodeUint(64))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FUintptr = 0
	} else {
		yyv1189 := &x.FUintptr
		yym1190 := z.DecBinary()
		_ = yym1190
		if false {
		} else {
			*((*uintptr)(yyv1189)) = uintptr(r.DecodeUint(codecSelferBitsize19781))
		}
	}
	if x.FptrUintptr == nil {
		x.FptrUintptr = new(uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrUintptr != nil {
			x.FptrUintptr = nil
		}
	} else {
		if x.FptrUintptr == nil {
			x.FptrUintptr = new(uintptr)
		}
		yym1192 := z.DecBinary()
		_ = yym1192
		if false {
		} else {
			*((*uintptr)(x.FptrUintptr)) = uintptr(r.DecodeUint(codecSelferBitsize19781))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FInt = 0
	} else {
		yyv1193 := &x.FInt
		yym1194 := z.DecBinary()
		_ = yym1194
		if false {
		} else {
			*((*int)(yyv1193)) = int(r.DecodeInt(codecSelferBitsize19781))
		}
	}
	if x.FptrInt == nil {
		x.FptrInt = new(int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrInt != nil {
			x.FptrInt = nil
		}
	} else {
		if x.FptrInt == nil {
			x.FptrInt = new(int)
		}
		yym1196 := z.DecBinary()
		_ = yym1196
		if false {
		} else {
			*((*int)(x.FptrInt)) = int(r.DecodeInt(codecSelferBitsize19781))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FInt8 = 0
	} else {
		yyv1197 := &x.FInt8
		yym1198 := z.DecBinary()
		_ = yym1198
		if false {
		} else {
			*((*int8)(yyv1197)) = int8(r.DecodeInt(8))
		}
	}
	if x.FptrInt8 == nil {
		x.FptrInt8 = new(int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrInt8 != nil {
			x.FptrInt8 = nil
		}
	} else {
		if x.FptrInt8 == nil {
			x.FptrInt8 = new(int8)
		}
		yym1200 := z.DecBinary()
		_ = yym1200
		if false {
		} else {
			*((*int8)(x.FptrInt8)) = int8(r.DecodeInt(8))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FInt16 = 0
	} else {
		yyv1201 := &x.FInt16
		yym1202 := z.DecBinary()
		_ = yym1202
		if false {
		} else {
			*((*int16)(yyv1201)) = int16(r.DecodeInt(16))
		}
	}
	if x.FptrInt16 == nil {
		x.FptrInt16 = new(int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrInt16 != nil {
			x.FptrInt16 = nil
		}
	} else {
		if x.FptrInt16 == nil {
			x.FptrInt16 = new(int16)
		}
		yym1204 := z.DecBinary()
		_ = yym1204
		if false {
		} else {
			*((*int16)(x.FptrInt16)) = int16(r.DecodeInt(16))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FInt32 = 0
	} else {
		yyv1205 := &x.FInt32
		yym1206 := z.DecBinary()
		_ = yym1206
		if false {
		} else {
			*((*int32)(yyv1205)) = int32(r.DecodeInt(32))
		}
	}
	if x.FptrInt32 == nil {
		x.FptrInt32 = new(int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrInt32 != nil {
			x.FptrInt32 = nil
		}
	} else {
		if x.FptrInt32 == nil {
			x.FptrInt32 = new(int32)
		}
		yym1208 := z.DecBinary()
		_ = yym1208
		if false {
		} else {
			*((*int32)(x.FptrInt32)) = int32(r.DecodeInt(32))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FInt64 = 0
	} else {
		yyv1209 := &x.FInt64
		yym1210 := z.DecBinary()
		_ = yym1210
		if false {
		} else {
			*((*int64)(yyv1209)) = int64(r.DecodeInt(64))
		}
	}
	if x.FptrInt64 == nil {
		x.FptrInt64 = new(int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrInt64 != nil {
			x.FptrInt64 = nil
		}
	} else {
		if x.FptrInt64 == nil {
			x.FptrInt64 = new(int64)
		}
		yym1212 := z.DecBinary()
		_ = yym1212
		if false {
		} else {
			*((*int64)(x.FptrInt64)) = int64(r.DecodeInt(64))
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FBool = false
	} else {
		yyv1213 := &x.FBool
		yym1214 := z.DecBinary()
		_ = yym1214
		if false {
		} else {
			*((*bool)(yyv1213)) = r.DecodeBool()
		}
	}
	if x.FptrBool == nil {
		x.FptrBool = new(bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrBool != nil {
			x.FptrBool = nil
		}
	} else {
		if x.FptrBool == nil {
			x.FptrBool = new(bool)
		}
		yym1216 := z.DecBinary()
		_ = yym1216
		if false {
		} else {
			*((*bool)(x.FptrBool)) = r.DecodeBool()
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceIntf = nil
	} else {
		yyv1217 := &x.FSliceIntf
		yym1218 := z.DecBinary()
		_ = yym1218
		if false {
		} else {
			z.F.DecSliceIntfX(yyv1217, d)
		}
	}
	if x.FptrSliceIntf == nil {
		x.FptrSliceIntf = new([]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceIntf != nil {
			x.FptrSliceIntf = nil
		}
	} else {
		if x.FptrSliceIntf == nil {
			x.FptrSliceIntf = new([]interface{})
		}
		yym1220 := z.DecBinary()
		_ = yym1220
		if false {
		} else {
			z.F.DecSliceIntfX(x.FptrSliceIntf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceString = nil
	} else {
		yyv1221 := &x.FSliceString
		yym1222 := z.DecBinary()
		_ = yym1222
		if false {
		} else {
			z.F.DecSliceStringX(yyv1221, d)
		}
	}
	if x.FptrSliceString == nil {
		x.FptrSliceString = new([]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceString != nil {
			x.FptrSliceString = nil
		}
	} else {
		if x.FptrSliceString == nil {
			x.FptrSliceString = new([]string)
		}
		yym1224 := z.DecBinary()
		_ = yym1224
		if false {
		} else {
			z.F.DecSliceStringX(x.FptrSliceString, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceFloat32 = nil
	} else {
		yyv1225 := &x.FSliceFloat32
		yym1226 := z.DecBinary()
		_ = yym1226
		if false {
		} else {
			z.F.DecSliceFloat32X(yyv1225, d)
		}
	}
	if x.FptrSliceFloat32 == nil {
		x.FptrSliceFloat32 = new([]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceFloat32 != nil {
			x.FptrSliceFloat32 = nil
		}
	} else {
		if x.FptrSliceFloat32 == nil {
			x.FptrSliceFloat32 = new([]float32)
		}
		yym1228 := z.DecBinary()
		_ = yym1228
		if false {
		} else {
			z.F.DecSliceFloat32X(x.FptrSliceFloat32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceFloat64 = nil
	} else {
		yyv1229 := &x.FSliceFloat64
		yym1230 := z.DecBinary()
		_ = yym1230
		if false {
		} else {
			z.F.DecSliceFloat64X(yyv1229, d)
		}
	}
	if x.FptrSliceFloat64 == nil {
		x.FptrSliceFloat64 = new([]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceFloat64 != nil {
			x.FptrSliceFloat64 = nil
		}
	} else {
		if x.FptrSliceFloat64 == nil {
			x.FptrSliceFloat64 = new([]float64)
		}
		yym1232 := z.DecBinary()
		_ = yym1232
		if false {
		} else {
			z.F.DecSliceFloat64X(x.FptrSliceFloat64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceUint = nil
	} else {
		yyv1233 := &x.FSliceUint
		yym1234 := z.DecBinary()
		_ = yym1234
		if false {
		} else {
			z.F.DecSliceUintX(yyv1233, d)
		}
	}
	if x.FptrSliceUint == nil {
		x.FptrSliceUint = new([]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceUint != nil {
			x.FptrSliceUint = nil
		}
	} else {
		if x.FptrSliceUint == nil {
			x.FptrSliceUint = new([]uint)
		}
		yym1236 := z.DecBinary()
		_ = yym1236
		if false {
		} else {
			z.F.DecSliceUintX(x.FptrSliceUint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceUint16 = nil
	} else {
		yyv1237 := &x.FSliceUint16
		yym1238 := z.DecBinary()
		_ = yym1238
		if false {
		} else {
			z.F.DecSliceUint16X(yyv1237, d)
		}
	}
	if x.FptrSliceUint16 == nil {
		x.FptrSliceUint16 = new([]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceUint16 != nil {
			x.FptrSliceUint16 = nil
		}
	} else {
		if x.FptrSliceUint16 == nil {
			x.FptrSliceUint16 = new([]uint16)
		}
		yym1240 := z.DecBinary()
		_ = yym1240
		if false {
		} else {
			z.F.DecSliceUint16X(x.FptrSliceUint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceUint32 = nil
	} else {
		yyv1241 := &x.FSliceUint32
		yym1242 := z.DecBinary()
		_ = yym1242
		if false {
		} else {
			z.F.DecSliceUint32X(yyv1241, d)
		}
	}
	if x.FptrSliceUint32 == nil {
		x.FptrSliceUint32 = new([]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceUint32 != nil {
			x.FptrSliceUint32 = nil
		}
	} else {
		if x.FptrSliceUint32 == nil {
			x.FptrSliceUint32 = new([]uint32)
		}
		yym1244 := z.DecBinary()
		_ = yym1244
		if false {
		} else {
			z.F.DecSliceUint32X(x.FptrSliceUint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceUint64 = nil
	} else {
		yyv1245 := &x.FSliceUint64
		yym1246 := z.DecBinary()
		_ = yym1246
		if false {
		} else {
			z.F.DecSliceUint64X(yyv1245, d)
		}
	}
	if x.FptrSliceUint64 == nil {
		x.FptrSliceUint64 = new([]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceUint64 != nil {
			x.FptrSliceUint64 = nil
		}
	} else {
		if x.FptrSliceUint64 == nil {
			x.FptrSliceUint64 = new([]uint64)
		}
		yym1248 := z.DecBinary()
		_ = yym1248
		if false {
		} else {
			z.F.DecSliceUint64X(x.FptrSliceUint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceUintptr = nil
	} else {
		yyv1249 := &x.FSliceUintptr
		yym1250 := z.DecBinary()
		_ = yym1250
		if false {
		} else {
			z.F.DecSliceUintptrX(yyv1249, d)
		}
	}
	if x.FptrSliceUintptr == nil {
		x.FptrSliceUintptr = new([]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceUintptr != nil {
			x.FptrSliceUintptr = nil
		}
	} else {
		if x.FptrSliceUintptr == nil {
			x.FptrSliceUintptr = new([]uintptr)
		}
		yym1252 := z.DecBinary()
		_ = yym1252
		if false {
		} else {
			z.F.DecSliceUintptrX(x.FptrSliceUintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceInt = nil
	} else {
		yyv1253 := &x.FSliceInt
		yym1254 := z.DecBinary()
		_ = yym1254
		if false {
		} else {
			z.F.DecSliceIntX(yyv1253, d)
		}
	}
	if x.FptrSliceInt == nil {
		x.FptrSliceInt = new([]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceInt != nil {
			x.FptrSliceInt = nil
		}
	} else {
		if x.FptrSliceInt == nil {
			x.FptrSliceInt = new([]int)
		}
		yym1256 := z.DecBinary()
		_ = yym1256
		if false {
		} else {
			z.F.DecSliceIntX(x.FptrSliceInt, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceInt8 = nil
	} else {
		yyv1257 := &x.FSliceInt8
		yym1258 := z.DecBinary()
		_ = yym1258
		if false {
		} else {
			z.F.DecSliceInt8X(yyv1257, d)
		}
	}
	if x.FptrSliceInt8 == nil {
		x.FptrSliceInt8 = new([]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceInt8 != nil {
			x.FptrSliceInt8 = nil
		}
	} else {
		if x.FptrSliceInt8 == nil {
			x.FptrSliceInt8 = new([]int8)
		}
		yym1260 := z.DecBinary()
		_ = yym1260
		if false {
		} else {
			z.F.DecSliceInt8X(x.FptrSliceInt8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceInt16 = nil
	} else {
		yyv1261 := &x.FSliceInt16
		yym1262 := z.DecBinary()
		_ = yym1262
		if false {
		} else {
			z.F.DecSliceInt16X(yyv1261, d)
		}
	}
	if x.FptrSliceInt16 == nil {
		x.FptrSliceInt16 = new([]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceInt16 != nil {
			x.FptrSliceInt16 = nil
		}
	} else {
		if x.FptrSliceInt16 == nil {
			x.FptrSliceInt16 = new([]int16)
		}
		yym1264 := z.DecBinary()
		_ = yym1264
		if false {
		} else {
			z.F.DecSliceInt16X(x.FptrSliceInt16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceInt32 = nil
	} else {
		yyv1265 := &x.FSliceInt32
		yym1266 := z.DecBinary()
		_ = yym1266
		if false {
		} else {
			z.F.DecSliceInt32X(yyv1265, d)
		}
	}
	if x.FptrSliceInt32 == nil {
		x.FptrSliceInt32 = new([]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceInt32 != nil {
			x.FptrSliceInt32 = nil
		}
	} else {
		if x.FptrSliceInt32 == nil {
			x.FptrSliceInt32 = new([]int32)
		}
		yym1268 := z.DecBinary()
		_ = yym1268
		if false {
		} else {
			z.F.DecSliceInt32X(x.FptrSliceInt32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceInt64 = nil
	} else {
		yyv1269 := &x.FSliceInt64
		yym1270 := z.DecBinary()
		_ = yym1270
		if false {
		} else {
			z.F.DecSliceInt64X(yyv1269, d)
		}
	}
	if x.FptrSliceInt64 == nil {
		x.FptrSliceInt64 = new([]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceInt64 != nil {
			x.FptrSliceInt64 = nil
		}
	} else {
		if x.FptrSliceInt64 == nil {
			x.FptrSliceInt64 = new([]int64)
		}
		yym1272 := z.DecBinary()
		_ = yym1272
		if false {
		} else {
			z.F.DecSliceInt64X(x.FptrSliceInt64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FSliceBool = nil
	} else {
		yyv1273 := &x.FSliceBool
		yym1274 := z.DecBinary()
		_ = yym1274
		if false {
		} else {
			z.F.DecSliceBoolX(yyv1273, d)
		}
	}
	if x.FptrSliceBool == nil {
		x.FptrSliceBool = new([]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrSliceBool != nil {
			x.FptrSliceBool = nil
		}
	} else {
		if x.FptrSliceBool == nil {
			x.FptrSliceBool = new([]bool)
		}
		yym1276 := z.DecBinary()
		_ = yym1276
		if false {
		} else {
			z.F.DecSliceBoolX(x.FptrSliceBool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfIntf = nil
	} else {
		yyv1277 := &x.FMapIntfIntf
		yym1278 := z.DecBinary()
		_ = yym1278
		if false {
		} else {
			z.F.DecMapIntfIntfX(yyv1277, d)
		}
	}
	if x.FptrMapIntfIntf == nil {
		x.FptrMapIntfIntf = new(map[interface{}]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfIntf != nil {
			x.FptrMapIntfIntf = nil
		}
	} else {
		if x.FptrMapIntfIntf == nil {
			x.FptrMapIntfIntf = new(map[interface{}]interface{})
		}
		yym1280 := z.DecBinary()
		_ = yym1280
		if false {
		} else {
			z.F.DecMapIntfIntfX(x.FptrMapIntfIntf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfString = nil
	} else {
		yyv1281 := &x.FMapIntfString
		yym1282 := z.DecBinary()
		_ = yym1282
		if false {
		} else {
			z.F.DecMapIntfStringX(yyv1281, d)
		}
	}
	if x.FptrMapIntfString == nil {
		x.FptrMapIntfString = new(map[interface{}]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfString != nil {
			x.FptrMapIntfString = nil
		}
	} else {
		if x.FptrMapIntfString == nil {
			x.FptrMapIntfString = new(map[interface{}]string)
		}
		yym1284 := z.DecBinary()
		_ = yym1284
		if false {
		} else {
			z.F.DecMapIntfStringX(x.FptrMapIntfString, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfUint = nil
	} else {
		yyv1285 := &x.FMapIntfUint
		yym1286 := z.DecBinary()
		_ = yym1286
		if false {
		} else {
			z.F.DecMapIntfUintX(yyv1285, d)
		}
	}
	if x.FptrMapIntfUint == nil {
		x.FptrMapIntfUint = new(map[interface{}]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfUint != nil {
			x.FptrMapIntfUint = nil
		}
	} else {
		if x.FptrMapIntfUint == nil {
			x.FptrMapIntfUint = new(map[interface{}]uint)
		}
		yym1288 := z.DecBinary()
		_ = yym1288
		if false {
		} else {
			z.F.DecMapIntfUintX(x.FptrMapIntfUint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfUint8 = nil
	} else {
		yyv1289 := &x.FMapIntfUint8
		yym1290 := z.DecBinary()
		_ = yym1290
		if false {
		} else {
			z.F.DecMapIntfUint8X(yyv1289, d)
		}
	}
	if x.FptrMapIntfUint8 == nil {
		x.FptrMapIntfUint8 = new(map[interface{}]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfUint8 != nil {
			x.FptrMapIntfUint8 = nil
		}
	} else {
		if x.FptrMapIntfUint8 == nil {
			x.FptrMapIntfUint8 = new(map[interface{}]uint8)
		}
		yym1292 := z.DecBinary()
		_ = yym1292
		if false {
		} else {
			z.F.DecMapIntfUint8X(x.FptrMapIntfUint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfUint16 = nil
	} else {
		yyv1293 := &x.FMapIntfUint16
		yym1294 := z.DecBinary()
		_ = yym1294
		if false {
		} else {
			z.F.DecMapIntfUint16X(yyv1293, d)
		}
	}
	if x.FptrMapIntfUint16 == nil {
		x.FptrMapIntfUint16 = new(map[interface{}]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfUint16 != nil {
			x.FptrMapIntfUint16 = nil
		}
	} else {
		if x.FptrMapIntfUint16 == nil {
			x.FptrMapIntfUint16 = new(map[interface{}]uint16)
		}
		yym1296 := z.DecBinary()
		_ = yym1296
		if false {
		} else {
			z.F.DecMapIntfUint16X(x.FptrMapIntfUint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfUint32 = nil
	} else {
		yyv1297 := &x.FMapIntfUint32
		yym1298 := z.DecBinary()
		_ = yym1298
		if false {
		} else {
			z.F.DecMapIntfUint32X(yyv1297, d)
		}
	}
	if x.FptrMapIntfUint32 == nil {
		x.FptrMapIntfUint32 = new(map[interface{}]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfUint32 != nil {
			x.FptrMapIntfUint32 = nil
		}
	} else {
		if x.FptrMapIntfUint32 == nil {
			x.FptrMapIntfUint32 = new(map[interface{}]uint32)
		}
		yym1300 := z.DecBinary()
		_ = yym1300
		if false {
		} else {
			z.F.DecMapIntfUint32X(x.FptrMapIntfUint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfUint64 = nil
	} else {
		yyv1301 := &x.FMapIntfUint64
		yym1302 := z.DecBinary()
		_ = yym1302
		if false {
		} else {
			z.F.DecMapIntfUint64X(yyv1301, d)
		}
	}
	if x.FptrMapIntfUint64 == nil {
		x.FptrMapIntfUint64 = new(map[interface{}]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfUint64 != nil {
			x.FptrMapIntfUint64 = nil
		}
	} else {
		if x.FptrMapIntfUint64 == nil {
			x.FptrMapIntfUint64 = new(map[interface{}]uint64)
		}
		yym1304 := z.DecBinary()
		_ = yym1304
		if false {
		} else {
			z.F.DecMapIntfUint64X(x.FptrMapIntfUint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfUintptr = nil
	} else {
		yyv1305 := &x.FMapIntfUintptr
		yym1306 := z.DecBinary()
		_ = yym1306
		if false {
		} else {
			z.F.DecMapIntfUintptrX(yyv1305, d)
		}
	}
	if x.FptrMapIntfUintptr == nil {
		x.FptrMapIntfUintptr = new(map[interface{}]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfUintptr != nil {
			x.FptrMapIntfUintptr = nil
		}
	} else {
		if x.FptrMapIntfUintptr == nil {
			x.FptrMapIntfUintptr = new(map[interface{}]uintptr)
		}
		yym1308 := z.DecBinary()
		_ = yym1308
		if false {
		} else {
			z.F.DecMapIntfUintptrX(x.FptrMapIntfUintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfInt = nil
	} else {
		yyv1309 := &x.FMapIntfInt
		yym1310 := z.DecBinary()
		_ = yym1310
		if false {
		} else {
			z.F.DecMapIntfIntX(yyv1309, d)
		}
	}
	if x.FptrMapIntfInt == nil {
		x.FptrMapIntfInt = new(map[interface{}]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfInt != nil {
			x.FptrMapIntfInt = nil
		}
	} else {
		if x.FptrMapIntfInt == nil {
			x.FptrMapIntfInt = new(map[interface{}]int)
		}
		yym1312 := z.DecBinary()
		_ = yym1312
		if false {
		} else {
			z.F.DecMapIntfIntX(x.FptrMapIntfInt, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfInt8 = nil
	} else {
		yyv1313 := &x.FMapIntfInt8
		yym1314 := z.DecBinary()
		_ = yym1314
		if false {
		} else {
			z.F.DecMapIntfInt8X(yyv1313, d)
		}
	}
	if x.FptrMapIntfInt8 == nil {
		x.FptrMapIntfInt8 = new(map[interface{}]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfInt8 != nil {
			x.FptrMapIntfInt8 = nil
		}
	} else {
		if x.FptrMapIntfInt8 == nil {
			x.FptrMapIntfInt8 = new(map[interface{}]int8)
		}
		yym1316 := z.DecBinary()
		_ = yym1316
		if false {
		} else {
			z.F.DecMapIntfInt8X(x.FptrMapIntfInt8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfInt16 = nil
	} else {
		yyv1317 := &x.FMapIntfInt16
		yym1318 := z.DecBinary()
		_ = yym1318
		if false {
		} else {
			z.F.DecMapIntfInt16X(yyv1317, d)
		}
	}
	if x.FptrMapIntfInt16 == nil {
		x.FptrMapIntfInt16 = new(map[interface{}]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfInt16 != nil {
			x.FptrMapIntfInt16 = nil
		}
	} else {
		if x.FptrMapIntfInt16 == nil {
			x.FptrMapIntfInt16 = new(map[interface{}]int16)
		}
		yym1320 := z.DecBinary()
		_ = yym1320
		if false {
		} else {
			z.F.DecMapIntfInt16X(x.FptrMapIntfInt16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfInt32 = nil
	} else {
		yyv1321 := &x.FMapIntfInt32
		yym1322 := z.DecBinary()
		_ = yym1322
		if false {
		} else {
			z.F.DecMapIntfInt32X(yyv1321, d)
		}
	}
	if x.FptrMapIntfInt32 == nil {
		x.FptrMapIntfInt32 = new(map[interface{}]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfInt32 != nil {
			x.FptrMapIntfInt32 = nil
		}
	} else {
		if x.FptrMapIntfInt32 == nil {
			x.FptrMapIntfInt32 = new(map[interface{}]int32)
		}
		yym1324 := z.DecBinary()
		_ = yym1324
		if false {
		} else {
			z.F.DecMapIntfInt32X(x.FptrMapIntfInt32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfInt64 = nil
	} else {
		yyv1325 := &x.FMapIntfInt64
		yym1326 := z.DecBinary()
		_ = yym1326
		if false {
		} else {
			z.F.DecMapIntfInt64X(yyv1325, d)
		}
	}
	if x.FptrMapIntfInt64 == nil {
		x.FptrMapIntfInt64 = new(map[interface{}]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfInt64 != nil {
			x.FptrMapIntfInt64 = nil
		}
	} else {
		if x.FptrMapIntfInt64 == nil {
			x.FptrMapIntfInt64 = new(map[interface{}]int64)
		}
		yym1328 := z.DecBinary()
		_ = yym1328
		if false {
		} else {
			z.F.DecMapIntfInt64X(x.FptrMapIntfInt64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfFloat32 = nil
	} else {
		yyv1329 := &x.FMapIntfFloat32
		yym1330 := z.DecBinary()
		_ = yym1330
		if false {
		} else {
			z.F.DecMapIntfFloat32X(yyv1329, d)
		}
	}
	if x.FptrMapIntfFloat32 == nil {
		x.FptrMapIntfFloat32 = new(map[interface{}]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfFloat32 != nil {
			x.FptrMapIntfFloat32 = nil
		}
	} else {
		if x.FptrMapIntfFloat32 == nil {
			x.FptrMapIntfFloat32 = new(map[interface{}]float32)
		}
		yym1332 := z.DecBinary()
		_ = yym1332
		if false {
		} else {
			z.F.DecMapIntfFloat32X(x.FptrMapIntfFloat32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfFloat64 = nil
	} else {
		yyv1333 := &x.FMapIntfFloat64
		yym1334 := z.DecBinary()
		_ = yym1334
		if false {
		} else {
			z.F.DecMapIntfFloat64X(yyv1333, d)
		}
	}
	if x.FptrMapIntfFloat64 == nil {
		x.FptrMapIntfFloat64 = new(map[interface{}]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfFloat64 != nil {
			x.FptrMapIntfFloat64 = nil
		}
	} else {
		if x.FptrMapIntfFloat64 == nil {
			x.FptrMapIntfFloat64 = new(map[interface{}]float64)
		}
		yym1336 := z.DecBinary()
		_ = yym1336
		if false {
		} else {
			z.F.DecMapIntfFloat64X(x.FptrMapIntfFloat64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntfBool = nil
	} else {
		yyv1337 := &x.FMapIntfBool
		yym1338 := z.DecBinary()
		_ = yym1338
		if false {
		} else {
			z.F.DecMapIntfBoolX(yyv1337, d)
		}
	}
	if x.FptrMapIntfBool == nil {
		x.FptrMapIntfBool = new(map[interface{}]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntfBool != nil {
			x.FptrMapIntfBool = nil
		}
	} else {
		if x.FptrMapIntfBool == nil {
			x.FptrMapIntfBool = new(map[interface{}]bool)
		}
		yym1340 := z.DecBinary()
		_ = yym1340
		if false {
		} else {
			z.F.DecMapIntfBoolX(x.FptrMapIntfBool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringIntf = nil
	} else {
		yyv1341 := &x.FMapStringIntf
		yym1342 := z.DecBinary()
		_ = yym1342
		if false {
		} else {
			z.F.DecMapStringIntfX(yyv1341, d)
		}
	}
	if x.FptrMapStringIntf == nil {
		x.FptrMapStringIntf = new(map[string]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringIntf != nil {
			x.FptrMapStringIntf = nil
		}
	} else {
		if x.FptrMapStringIntf == nil {
			x.FptrMapStringIntf = new(map[string]interface{})
		}
		yym1344 := z.DecBinary()
		_ = yym1344
		if false {
		} else {
			z.F.DecMapStringIntfX(x.FptrMapStringIntf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringString = nil
	} else {
		yyv1345 := &x.FMapStringString
		yym1346 := z.DecBinary()
		_ = yym1346
		if false {
		} else {
			z.F.DecMapStringStringX(yyv1345, d)
		}
	}
	if x.FptrMapStringString == nil {
		x.FptrMapStringString = new(map[string]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringString != nil {
			x.FptrMapStringString = nil
		}
	} else {
		if x.FptrMapStringString == nil {
			x.FptrMapStringString = new(map[string]string)
		}
		yym1348 := z.DecBinary()
		_ = yym1348
		if false {
		} else {
			z.F.DecMapStringStringX(x.FptrMapStringString, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringUint = nil
	} else {
		yyv1349 := &x.FMapStringUint
		yym1350 := z.DecBinary()
		_ = yym1350
		if false {
		} else {
			z.F.DecMapStringUintX(yyv1349, d)
		}
	}
	if x.FptrMapStringUint == nil {
		x.FptrMapStringUint = new(map[string]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringUint != nil {
			x.FptrMapStringUint = nil
		}
	} else {
		if x.FptrMapStringUint == nil {
			x.FptrMapStringUint = new(map[string]uint)
		}
		yym1352 := z.DecBinary()
		_ = yym1352
		if false {
		} else {
			z.F.DecMapStringUintX(x.FptrMapStringUint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringUint8 = nil
	} else {
		yyv1353 := &x.FMapStringUint8
		yym1354 := z.DecBinary()
		_ = yym1354
		if false {
		} else {
			z.F.DecMapStringUint8X(yyv1353, d)
		}
	}
	if x.FptrMapStringUint8 == nil {
		x.FptrMapStringUint8 = new(map[string]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringUint8 != nil {
			x.FptrMapStringUint8 = nil
		}
	} else {
		if x.FptrMapStringUint8 == nil {
			x.FptrMapStringUint8 = new(map[string]uint8)
		}
		yym1356 := z.DecBinary()
		_ = yym1356
		if false {
		} else {
			z.F.DecMapStringUint8X(x.FptrMapStringUint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringUint16 = nil
	} else {
		yyv1357 := &x.FMapStringUint16
		yym1358 := z.DecBinary()
		_ = yym1358
		if false {
		} else {
			z.F.DecMapStringUint16X(yyv1357, d)
		}
	}
	if x.FptrMapStringUint16 == nil {
		x.FptrMapStringUint16 = new(map[string]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringUint16 != nil {
			x.FptrMapStringUint16 = nil
		}
	} else {
		if x.FptrMapStringUint16 == nil {
			x.FptrMapStringUint16 = new(map[string]uint16)
		}
		yym1360 := z.DecBinary()
		_ = yym1360
		if false {
		} else {
			z.F.DecMapStringUint16X(x.FptrMapStringUint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringUint32 = nil
	} else {
		yyv1361 := &x.FMapStringUint32
		yym1362 := z.DecBinary()
		_ = yym1362
		if false {
		} else {
			z.F.DecMapStringUint32X(yyv1361, d)
		}
	}
	if x.FptrMapStringUint32 == nil {
		x.FptrMapStringUint32 = new(map[string]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringUint32 != nil {
			x.FptrMapStringUint32 = nil
		}
	} else {
		if x.FptrMapStringUint32 == nil {
			x.FptrMapStringUint32 = new(map[string]uint32)
		}
		yym1364 := z.DecBinary()
		_ = yym1364
		if false {
		} else {
			z.F.DecMapStringUint32X(x.FptrMapStringUint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringUint64 = nil
	} else {
		yyv1365 := &x.FMapStringUint64
		yym1366 := z.DecBinary()
		_ = yym1366
		if false {
		} else {
			z.F.DecMapStringUint64X(yyv1365, d)
		}
	}
	if x.FptrMapStringUint64 == nil {
		x.FptrMapStringUint64 = new(map[string]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringUint64 != nil {
			x.FptrMapStringUint64 = nil
		}
	} else {
		if x.FptrMapStringUint64 == nil {
			x.FptrMapStringUint64 = new(map[string]uint64)
		}
		yym1368 := z.DecBinary()
		_ = yym1368
		if false {
		} else {
			z.F.DecMapStringUint64X(x.FptrMapStringUint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringUintptr = nil
	} else {
		yyv1369 := &x.FMapStringUintptr
		yym1370 := z.DecBinary()
		_ = yym1370
		if false {
		} else {
			z.F.DecMapStringUintptrX(yyv1369, d)
		}
	}
	if x.FptrMapStringUintptr == nil {
		x.FptrMapStringUintptr = new(map[string]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringUintptr != nil {
			x.FptrMapStringUintptr = nil
		}
	} else {
		if x.FptrMapStringUintptr == nil {
			x.FptrMapStringUintptr = new(map[string]uintptr)
		}
		yym1372 := z.DecBinary()
		_ = yym1372
		if false {
		} else {
			z.F.DecMapStringUintptrX(x.FptrMapStringUintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringInt = nil
	} else {
		yyv1373 := &x.FMapStringInt
		yym1374 := z.DecBinary()
		_ = yym1374
		if false {
		} else {
			z.F.DecMapStringIntX(yyv1373, d)
		}
	}
	if x.FptrMapStringInt == nil {
		x.FptrMapStringInt = new(map[string]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringInt != nil {
			x.FptrMapStringInt = nil
		}
	} else {
		if x.FptrMapStringInt == nil {
			x.FptrMapStringInt = new(map[string]int)
		}
		yym1376 := z.DecBinary()
		_ = yym1376
		if false {
		} else {
			z.F.DecMapStringIntX(x.FptrMapStringInt, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringInt8 = nil
	} else {
		yyv1377 := &x.FMapStringInt8
		yym1378 := z.DecBinary()
		_ = yym1378
		if false {
		} else {
			z.F.DecMapStringInt8X(yyv1377, d)
		}
	}
	if x.FptrMapStringInt8 == nil {
		x.FptrMapStringInt8 = new(map[string]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringInt8 != nil {
			x.FptrMapStringInt8 = nil
		}
	} else {
		if x.FptrMapStringInt8 == nil {
			x.FptrMapStringInt8 = new(map[string]int8)
		}
		yym1380 := z.DecBinary()
		_ = yym1380
		if false {
		} else {
			z.F.DecMapStringInt8X(x.FptrMapStringInt8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringInt16 = nil
	} else {
		yyv1381 := &x.FMapStringInt16
		yym1382 := z.DecBinary()
		_ = yym1382
		if false {
		} else {
			z.F.DecMapStringInt16X(yyv1381, d)
		}
	}
	if x.FptrMapStringInt16 == nil {
		x.FptrMapStringInt16 = new(map[string]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringInt16 != nil {
			x.FptrMapStringInt16 = nil
		}
	} else {
		if x.FptrMapStringInt16 == nil {
			x.FptrMapStringInt16 = new(map[string]int16)
		}
		yym1384 := z.DecBinary()
		_ = yym1384
		if false {
		} else {
			z.F.DecMapStringInt16X(x.FptrMapStringInt16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringInt32 = nil
	} else {
		yyv1385 := &x.FMapStringInt32
		yym1386 := z.DecBinary()
		_ = yym1386
		if false {
		} else {
			z.F.DecMapStringInt32X(yyv1385, d)
		}
	}
	if x.FptrMapStringInt32 == nil {
		x.FptrMapStringInt32 = new(map[string]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringInt32 != nil {
			x.FptrMapStringInt32 = nil
		}
	} else {
		if x.FptrMapStringInt32 == nil {
			x.FptrMapStringInt32 = new(map[string]int32)
		}
		yym1388 := z.DecBinary()
		_ = yym1388
		if false {
		} else {
			z.F.DecMapStringInt32X(x.FptrMapStringInt32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringInt64 = nil
	} else {
		yyv1389 := &x.FMapStringInt64
		yym1390 := z.DecBinary()
		_ = yym1390
		if false {
		} else {
			z.F.DecMapStringInt64X(yyv1389, d)
		}
	}
	if x.FptrMapStringInt64 == nil {
		x.FptrMapStringInt64 = new(map[string]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringInt64 != nil {
			x.FptrMapStringInt64 = nil
		}
	} else {
		if x.FptrMapStringInt64 == nil {
			x.FptrMapStringInt64 = new(map[string]int64)
		}
		yym1392 := z.DecBinary()
		_ = yym1392
		if false {
		} else {
			z.F.DecMapStringInt64X(x.FptrMapStringInt64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringFloat32 = nil
	} else {
		yyv1393 := &x.FMapStringFloat32
		yym1394 := z.DecBinary()
		_ = yym1394
		if false {
		} else {
			z.F.DecMapStringFloat32X(yyv1393, d)
		}
	}
	if x.FptrMapStringFloat32 == nil {
		x.FptrMapStringFloat32 = new(map[string]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringFloat32 != nil {
			x.FptrMapStringFloat32 = nil
		}
	} else {
		if x.FptrMapStringFloat32 == nil {
			x.FptrMapStringFloat32 = new(map[string]float32)
		}
		yym1396 := z.DecBinary()
		_ = yym1396
		if false {
		} else {
			z.F.DecMapStringFloat32X(x.FptrMapStringFloat32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringFloat64 = nil
	} else {
		yyv1397 := &x.FMapStringFloat64
		yym1398 := z.DecBinary()
		_ = yym1398
		if false {
		} else {
			z.F.DecMapStringFloat64X(yyv1397, d)
		}
	}
	if x.FptrMapStringFloat64 == nil {
		x.FptrMapStringFloat64 = new(map[string]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringFloat64 != nil {
			x.FptrMapStringFloat64 = nil
		}
	} else {
		if x.FptrMapStringFloat64 == nil {
			x.FptrMapStringFloat64 = new(map[string]float64)
		}
		yym1400 := z.DecBinary()
		_ = yym1400
		if false {
		} else {
			z.F.DecMapStringFloat64X(x.FptrMapStringFloat64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapStringBool = nil
	} else {
		yyv1401 := &x.FMapStringBool
		yym1402 := z.DecBinary()
		_ = yym1402
		if false {
		} else {
			z.F.DecMapStringBoolX(yyv1401, d)
		}
	}
	if x.FptrMapStringBool == nil {
		x.FptrMapStringBool = new(map[string]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapStringBool != nil {
			x.FptrMapStringBool = nil
		}
	} else {
		if x.FptrMapStringBool == nil {
			x.FptrMapStringBool = new(map[string]bool)
		}
		yym1404 := z.DecBinary()
		_ = yym1404
		if false {
		} else {
			z.F.DecMapStringBoolX(x.FptrMapStringBool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Intf = nil
	} else {
		yyv1405 := &x.FMapFloat32Intf
		yym1406 := z.DecBinary()
		_ = yym1406
		if false {
		} else {
			z.F.DecMapFloat32IntfX(yyv1405, d)
		}
	}
	if x.FptrMapFloat32Intf == nil {
		x.FptrMapFloat32Intf = new(map[float32]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Intf != nil {
			x.FptrMapFloat32Intf = nil
		}
	} else {
		if x.FptrMapFloat32Intf == nil {
			x.FptrMapFloat32Intf = new(map[float32]interface{})
		}
		yym1408 := z.DecBinary()
		_ = yym1408
		if false {
		} else {
			z.F.DecMapFloat32IntfX(x.FptrMapFloat32Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32String = nil
	} else {
		yyv1409 := &x.FMapFloat32String
		yym1410 := z.DecBinary()
		_ = yym1410
		if false {
		} else {
			z.F.DecMapFloat32StringX(yyv1409, d)
		}
	}
	if x.FptrMapFloat32String == nil {
		x.FptrMapFloat32String = new(map[float32]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32String != nil {
			x.FptrMapFloat32String = nil
		}
	} else {
		if x.FptrMapFloat32String == nil {
			x.FptrMapFloat32String = new(map[float32]string)
		}
		yym1412 := z.DecBinary()
		_ = yym1412
		if false {
		} else {
			z.F.DecMapFloat32StringX(x.FptrMapFloat32String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Uint = nil
	} else {
		yyv1413 := &x.FMapFloat32Uint
		yym1414 := z.DecBinary()
		_ = yym1414
		if false {
		} else {
			z.F.DecMapFloat32UintX(yyv1413, d)
		}
	}
	if x.FptrMapFloat32Uint == nil {
		x.FptrMapFloat32Uint = new(map[float32]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Uint != nil {
			x.FptrMapFloat32Uint = nil
		}
	} else {
		if x.FptrMapFloat32Uint == nil {
			x.FptrMapFloat32Uint = new(map[float32]uint)
		}
		yym1416 := z.DecBinary()
		_ = yym1416
		if false {
		} else {
			z.F.DecMapFloat32UintX(x.FptrMapFloat32Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Uint8 = nil
	} else {
		yyv1417 := &x.FMapFloat32Uint8
		yym1418 := z.DecBinary()
		_ = yym1418
		if false {
		} else {
			z.F.DecMapFloat32Uint8X(yyv1417, d)
		}
	}
	if x.FptrMapFloat32Uint8 == nil {
		x.FptrMapFloat32Uint8 = new(map[float32]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Uint8 != nil {
			x.FptrMapFloat32Uint8 = nil
		}
	} else {
		if x.FptrMapFloat32Uint8 == nil {
			x.FptrMapFloat32Uint8 = new(map[float32]uint8)
		}
		yym1420 := z.DecBinary()
		_ = yym1420
		if false {
		} else {
			z.F.DecMapFloat32Uint8X(x.FptrMapFloat32Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Uint16 = nil
	} else {
		yyv1421 := &x.FMapFloat32Uint16
		yym1422 := z.DecBinary()
		_ = yym1422
		if false {
		} else {
			z.F.DecMapFloat32Uint16X(yyv1421, d)
		}
	}
	if x.FptrMapFloat32Uint16 == nil {
		x.FptrMapFloat32Uint16 = new(map[float32]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Uint16 != nil {
			x.FptrMapFloat32Uint16 = nil
		}
	} else {
		if x.FptrMapFloat32Uint16 == nil {
			x.FptrMapFloat32Uint16 = new(map[float32]uint16)
		}
		yym1424 := z.DecBinary()
		_ = yym1424
		if false {
		} else {
			z.F.DecMapFloat32Uint16X(x.FptrMapFloat32Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Uint32 = nil
	} else {
		yyv1425 := &x.FMapFloat32Uint32
		yym1426 := z.DecBinary()
		_ = yym1426
		if false {
		} else {
			z.F.DecMapFloat32Uint32X(yyv1425, d)
		}
	}
	if x.FptrMapFloat32Uint32 == nil {
		x.FptrMapFloat32Uint32 = new(map[float32]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Uint32 != nil {
			x.FptrMapFloat32Uint32 = nil
		}
	} else {
		if x.FptrMapFloat32Uint32 == nil {
			x.FptrMapFloat32Uint32 = new(map[float32]uint32)
		}
		yym1428 := z.DecBinary()
		_ = yym1428
		if false {
		} else {
			z.F.DecMapFloat32Uint32X(x.FptrMapFloat32Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Uint64 = nil
	} else {
		yyv1429 := &x.FMapFloat32Uint64
		yym1430 := z.DecBinary()
		_ = yym1430
		if false {
		} else {
			z.F.DecMapFloat32Uint64X(yyv1429, d)
		}
	}
	if x.FptrMapFloat32Uint64 == nil {
		x.FptrMapFloat32Uint64 = new(map[float32]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Uint64 != nil {
			x.FptrMapFloat32Uint64 = nil
		}
	} else {
		if x.FptrMapFloat32Uint64 == nil {
			x.FptrMapFloat32Uint64 = new(map[float32]uint64)
		}
		yym1432 := z.DecBinary()
		_ = yym1432
		if false {
		} else {
			z.F.DecMapFloat32Uint64X(x.FptrMapFloat32Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Uintptr = nil
	} else {
		yyv1433 := &x.FMapFloat32Uintptr
		yym1434 := z.DecBinary()
		_ = yym1434
		if false {
		} else {
			z.F.DecMapFloat32UintptrX(yyv1433, d)
		}
	}
	if x.FptrMapFloat32Uintptr == nil {
		x.FptrMapFloat32Uintptr = new(map[float32]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Uintptr != nil {
			x.FptrMapFloat32Uintptr = nil
		}
	} else {
		if x.FptrMapFloat32Uintptr == nil {
			x.FptrMapFloat32Uintptr = new(map[float32]uintptr)
		}
		yym1436 := z.DecBinary()
		_ = yym1436
		if false {
		} else {
			z.F.DecMapFloat32UintptrX(x.FptrMapFloat32Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Int = nil
	} else {
		yyv1437 := &x.FMapFloat32Int
		yym1438 := z.DecBinary()
		_ = yym1438
		if false {
		} else {
			z.F.DecMapFloat32IntX(yyv1437, d)
		}
	}
	if x.FptrMapFloat32Int == nil {
		x.FptrMapFloat32Int = new(map[float32]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Int != nil {
			x.FptrMapFloat32Int = nil
		}
	} else {
		if x.FptrMapFloat32Int == nil {
			x.FptrMapFloat32Int = new(map[float32]int)
		}
		yym1440 := z.DecBinary()
		_ = yym1440
		if false {
		} else {
			z.F.DecMapFloat32IntX(x.FptrMapFloat32Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Int8 = nil
	} else {
		yyv1441 := &x.FMapFloat32Int8
		yym1442 := z.DecBinary()
		_ = yym1442
		if false {
		} else {
			z.F.DecMapFloat32Int8X(yyv1441, d)
		}
	}
	if x.FptrMapFloat32Int8 == nil {
		x.FptrMapFloat32Int8 = new(map[float32]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Int8 != nil {
			x.FptrMapFloat32Int8 = nil
		}
	} else {
		if x.FptrMapFloat32Int8 == nil {
			x.FptrMapFloat32Int8 = new(map[float32]int8)
		}
		yym1444 := z.DecBinary()
		_ = yym1444
		if false {
		} else {
			z.F.DecMapFloat32Int8X(x.FptrMapFloat32Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Int16 = nil
	} else {
		yyv1445 := &x.FMapFloat32Int16
		yym1446 := z.DecBinary()
		_ = yym1446
		if false {
		} else {
			z.F.DecMapFloat32Int16X(yyv1445, d)
		}
	}
	if x.FptrMapFloat32Int16 == nil {
		x.FptrMapFloat32Int16 = new(map[float32]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Int16 != nil {
			x.FptrMapFloat32Int16 = nil
		}
	} else {
		if x.FptrMapFloat32Int16 == nil {
			x.FptrMapFloat32Int16 = new(map[float32]int16)
		}
		yym1448 := z.DecBinary()
		_ = yym1448
		if false {
		} else {
			z.F.DecMapFloat32Int16X(x.FptrMapFloat32Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Int32 = nil
	} else {
		yyv1449 := &x.FMapFloat32Int32
		yym1450 := z.DecBinary()
		_ = yym1450
		if false {
		} else {
			z.F.DecMapFloat32Int32X(yyv1449, d)
		}
	}
	if x.FptrMapFloat32Int32 == nil {
		x.FptrMapFloat32Int32 = new(map[float32]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Int32 != nil {
			x.FptrMapFloat32Int32 = nil
		}
	} else {
		if x.FptrMapFloat32Int32 == nil {
			x.FptrMapFloat32Int32 = new(map[float32]int32)
		}
		yym1452 := z.DecBinary()
		_ = yym1452
		if false {
		} else {
			z.F.DecMapFloat32Int32X(x.FptrMapFloat32Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Int64 = nil
	} else {
		yyv1453 := &x.FMapFloat32Int64
		yym1454 := z.DecBinary()
		_ = yym1454
		if false {
		} else {
			z.F.DecMapFloat32Int64X(yyv1453, d)
		}
	}
	if x.FptrMapFloat32Int64 == nil {
		x.FptrMapFloat32Int64 = new(map[float32]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Int64 != nil {
			x.FptrMapFloat32Int64 = nil
		}
	} else {
		if x.FptrMapFloat32Int64 == nil {
			x.FptrMapFloat32Int64 = new(map[float32]int64)
		}
		yym1456 := z.DecBinary()
		_ = yym1456
		if false {
		} else {
			z.F.DecMapFloat32Int64X(x.FptrMapFloat32Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Float32 = nil
	} else {
		yyv1457 := &x.FMapFloat32Float32
		yym1458 := z.DecBinary()
		_ = yym1458
		if false {
		} else {
			z.F.DecMapFloat32Float32X(yyv1457, d)
		}
	}
	if x.FptrMapFloat32Float32 == nil {
		x.FptrMapFloat32Float32 = new(map[float32]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Float32 != nil {
			x.FptrMapFloat32Float32 = nil
		}
	} else {
		if x.FptrMapFloat32Float32 == nil {
			x.FptrMapFloat32Float32 = new(map[float32]float32)
		}
		yym1460 := z.DecBinary()
		_ = yym1460
		if false {
		} else {
			z.F.DecMapFloat32Float32X(x.FptrMapFloat32Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Float64 = nil
	} else {
		yyv1461 := &x.FMapFloat32Float64
		yym1462 := z.DecBinary()
		_ = yym1462
		if false {
		} else {
			z.F.DecMapFloat32Float64X(yyv1461, d)
		}
	}
	if x.FptrMapFloat32Float64 == nil {
		x.FptrMapFloat32Float64 = new(map[float32]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Float64 != nil {
			x.FptrMapFloat32Float64 = nil
		}
	} else {
		if x.FptrMapFloat32Float64 == nil {
			x.FptrMapFloat32Float64 = new(map[float32]float64)
		}
		yym1464 := z.DecBinary()
		_ = yym1464
		if false {
		} else {
			z.F.DecMapFloat32Float64X(x.FptrMapFloat32Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat32Bool = nil
	} else {
		yyv1465 := &x.FMapFloat32Bool
		yym1466 := z.DecBinary()
		_ = yym1466
		if false {
		} else {
			z.F.DecMapFloat32BoolX(yyv1465, d)
		}
	}
	if x.FptrMapFloat32Bool == nil {
		x.FptrMapFloat32Bool = new(map[float32]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat32Bool != nil {
			x.FptrMapFloat32Bool = nil
		}
	} else {
		if x.FptrMapFloat32Bool == nil {
			x.FptrMapFloat32Bool = new(map[float32]bool)
		}
		yym1468 := z.DecBinary()
		_ = yym1468
		if false {
		} else {
			z.F.DecMapFloat32BoolX(x.FptrMapFloat32Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Intf = nil
	} else {
		yyv1469 := &x.FMapFloat64Intf
		yym1470 := z.DecBinary()
		_ = yym1470
		if false {
		} else {
			z.F.DecMapFloat64IntfX(yyv1469, d)
		}
	}
	if x.FptrMapFloat64Intf == nil {
		x.FptrMapFloat64Intf = new(map[float64]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Intf != nil {
			x.FptrMapFloat64Intf = nil
		}
	} else {
		if x.FptrMapFloat64Intf == nil {
			x.FptrMapFloat64Intf = new(map[float64]interface{})
		}
		yym1472 := z.DecBinary()
		_ = yym1472
		if false {
		} else {
			z.F.DecMapFloat64IntfX(x.FptrMapFloat64Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64String = nil
	} else {
		yyv1473 := &x.FMapFloat64String
		yym1474 := z.DecBinary()
		_ = yym1474
		if false {
		} else {
			z.F.DecMapFloat64StringX(yyv1473, d)
		}
	}
	if x.FptrMapFloat64String == nil {
		x.FptrMapFloat64String = new(map[float64]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64String != nil {
			x.FptrMapFloat64String = nil
		}
	} else {
		if x.FptrMapFloat64String == nil {
			x.FptrMapFloat64String = new(map[float64]string)
		}
		yym1476 := z.DecBinary()
		_ = yym1476
		if false {
		} else {
			z.F.DecMapFloat64StringX(x.FptrMapFloat64String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Uint = nil
	} else {
		yyv1477 := &x.FMapFloat64Uint
		yym1478 := z.DecBinary()
		_ = yym1478
		if false {
		} else {
			z.F.DecMapFloat64UintX(yyv1477, d)
		}
	}
	if x.FptrMapFloat64Uint == nil {
		x.FptrMapFloat64Uint = new(map[float64]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Uint != nil {
			x.FptrMapFloat64Uint = nil
		}
	} else {
		if x.FptrMapFloat64Uint == nil {
			x.FptrMapFloat64Uint = new(map[float64]uint)
		}
		yym1480 := z.DecBinary()
		_ = yym1480
		if false {
		} else {
			z.F.DecMapFloat64UintX(x.FptrMapFloat64Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Uint8 = nil
	} else {
		yyv1481 := &x.FMapFloat64Uint8
		yym1482 := z.DecBinary()
		_ = yym1482
		if false {
		} else {
			z.F.DecMapFloat64Uint8X(yyv1481, d)
		}
	}
	if x.FptrMapFloat64Uint8 == nil {
		x.FptrMapFloat64Uint8 = new(map[float64]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Uint8 != nil {
			x.FptrMapFloat64Uint8 = nil
		}
	} else {
		if x.FptrMapFloat64Uint8 == nil {
			x.FptrMapFloat64Uint8 = new(map[float64]uint8)
		}
		yym1484 := z.DecBinary()
		_ = yym1484
		if false {
		} else {
			z.F.DecMapFloat64Uint8X(x.FptrMapFloat64Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Uint16 = nil
	} else {
		yyv1485 := &x.FMapFloat64Uint16
		yym1486 := z.DecBinary()
		_ = yym1486
		if false {
		} else {
			z.F.DecMapFloat64Uint16X(yyv1485, d)
		}
	}
	if x.FptrMapFloat64Uint16 == nil {
		x.FptrMapFloat64Uint16 = new(map[float64]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Uint16 != nil {
			x.FptrMapFloat64Uint16 = nil
		}
	} else {
		if x.FptrMapFloat64Uint16 == nil {
			x.FptrMapFloat64Uint16 = new(map[float64]uint16)
		}
		yym1488 := z.DecBinary()
		_ = yym1488
		if false {
		} else {
			z.F.DecMapFloat64Uint16X(x.FptrMapFloat64Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Uint32 = nil
	} else {
		yyv1489 := &x.FMapFloat64Uint32
		yym1490 := z.DecBinary()
		_ = yym1490
		if false {
		} else {
			z.F.DecMapFloat64Uint32X(yyv1489, d)
		}
	}
	if x.FptrMapFloat64Uint32 == nil {
		x.FptrMapFloat64Uint32 = new(map[float64]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Uint32 != nil {
			x.FptrMapFloat64Uint32 = nil
		}
	} else {
		if x.FptrMapFloat64Uint32 == nil {
			x.FptrMapFloat64Uint32 = new(map[float64]uint32)
		}
		yym1492 := z.DecBinary()
		_ = yym1492
		if false {
		} else {
			z.F.DecMapFloat64Uint32X(x.FptrMapFloat64Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Uint64 = nil
	} else {
		yyv1493 := &x.FMapFloat64Uint64
		yym1494 := z.DecBinary()
		_ = yym1494
		if false {
		} else {
			z.F.DecMapFloat64Uint64X(yyv1493, d)
		}
	}
	if x.FptrMapFloat64Uint64 == nil {
		x.FptrMapFloat64Uint64 = new(map[float64]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Uint64 != nil {
			x.FptrMapFloat64Uint64 = nil
		}
	} else {
		if x.FptrMapFloat64Uint64 == nil {
			x.FptrMapFloat64Uint64 = new(map[float64]uint64)
		}
		yym1496 := z.DecBinary()
		_ = yym1496
		if false {
		} else {
			z.F.DecMapFloat64Uint64X(x.FptrMapFloat64Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Uintptr = nil
	} else {
		yyv1497 := &x.FMapFloat64Uintptr
		yym1498 := z.DecBinary()
		_ = yym1498
		if false {
		} else {
			z.F.DecMapFloat64UintptrX(yyv1497, d)
		}
	}
	if x.FptrMapFloat64Uintptr == nil {
		x.FptrMapFloat64Uintptr = new(map[float64]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Uintptr != nil {
			x.FptrMapFloat64Uintptr = nil
		}
	} else {
		if x.FptrMapFloat64Uintptr == nil {
			x.FptrMapFloat64Uintptr = new(map[float64]uintptr)
		}
		yym1500 := z.DecBinary()
		_ = yym1500
		if false {
		} else {
			z.F.DecMapFloat64UintptrX(x.FptrMapFloat64Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Int = nil
	} else {
		yyv1501 := &x.FMapFloat64Int
		yym1502 := z.DecBinary()
		_ = yym1502
		if false {
		} else {
			z.F.DecMapFloat64IntX(yyv1501, d)
		}
	}
	if x.FptrMapFloat64Int == nil {
		x.FptrMapFloat64Int = new(map[float64]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Int != nil {
			x.FptrMapFloat64Int = nil
		}
	} else {
		if x.FptrMapFloat64Int == nil {
			x.FptrMapFloat64Int = new(map[float64]int)
		}
		yym1504 := z.DecBinary()
		_ = yym1504
		if false {
		} else {
			z.F.DecMapFloat64IntX(x.FptrMapFloat64Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Int8 = nil
	} else {
		yyv1505 := &x.FMapFloat64Int8
		yym1506 := z.DecBinary()
		_ = yym1506
		if false {
		} else {
			z.F.DecMapFloat64Int8X(yyv1505, d)
		}
	}
	if x.FptrMapFloat64Int8 == nil {
		x.FptrMapFloat64Int8 = new(map[float64]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Int8 != nil {
			x.FptrMapFloat64Int8 = nil
		}
	} else {
		if x.FptrMapFloat64Int8 == nil {
			x.FptrMapFloat64Int8 = new(map[float64]int8)
		}
		yym1508 := z.DecBinary()
		_ = yym1508
		if false {
		} else {
			z.F.DecMapFloat64Int8X(x.FptrMapFloat64Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Int16 = nil
	} else {
		yyv1509 := &x.FMapFloat64Int16
		yym1510 := z.DecBinary()
		_ = yym1510
		if false {
		} else {
			z.F.DecMapFloat64Int16X(yyv1509, d)
		}
	}
	if x.FptrMapFloat64Int16 == nil {
		x.FptrMapFloat64Int16 = new(map[float64]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Int16 != nil {
			x.FptrMapFloat64Int16 = nil
		}
	} else {
		if x.FptrMapFloat64Int16 == nil {
			x.FptrMapFloat64Int16 = new(map[float64]int16)
		}
		yym1512 := z.DecBinary()
		_ = yym1512
		if false {
		} else {
			z.F.DecMapFloat64Int16X(x.FptrMapFloat64Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Int32 = nil
	} else {
		yyv1513 := &x.FMapFloat64Int32
		yym1514 := z.DecBinary()
		_ = yym1514
		if false {
		} else {
			z.F.DecMapFloat64Int32X(yyv1513, d)
		}
	}
	if x.FptrMapFloat64Int32 == nil {
		x.FptrMapFloat64Int32 = new(map[float64]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Int32 != nil {
			x.FptrMapFloat64Int32 = nil
		}
	} else {
		if x.FptrMapFloat64Int32 == nil {
			x.FptrMapFloat64Int32 = new(map[float64]int32)
		}
		yym1516 := z.DecBinary()
		_ = yym1516
		if false {
		} else {
			z.F.DecMapFloat64Int32X(x.FptrMapFloat64Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Int64 = nil
	} else {
		yyv1517 := &x.FMapFloat64Int64
		yym1518 := z.DecBinary()
		_ = yym1518
		if false {
		} else {
			z.F.DecMapFloat64Int64X(yyv1517, d)
		}
	}
	if x.FptrMapFloat64Int64 == nil {
		x.FptrMapFloat64Int64 = new(map[float64]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Int64 != nil {
			x.FptrMapFloat64Int64 = nil
		}
	} else {
		if x.FptrMapFloat64Int64 == nil {
			x.FptrMapFloat64Int64 = new(map[float64]int64)
		}
		yym1520 := z.DecBinary()
		_ = yym1520
		if false {
		} else {
			z.F.DecMapFloat64Int64X(x.FptrMapFloat64Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Float32 = nil
	} else {
		yyv1521 := &x.FMapFloat64Float32
		yym1522 := z.DecBinary()
		_ = yym1522
		if false {
		} else {
			z.F.DecMapFloat64Float32X(yyv1521, d)
		}
	}
	if x.FptrMapFloat64Float32 == nil {
		x.FptrMapFloat64Float32 = new(map[float64]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Float32 != nil {
			x.FptrMapFloat64Float32 = nil
		}
	} else {
		if x.FptrMapFloat64Float32 == nil {
			x.FptrMapFloat64Float32 = new(map[float64]float32)
		}
		yym1524 := z.DecBinary()
		_ = yym1524
		if false {
		} else {
			z.F.DecMapFloat64Float32X(x.FptrMapFloat64Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Float64 = nil
	} else {
		yyv1525 := &x.FMapFloat64Float64
		yym1526 := z.DecBinary()
		_ = yym1526
		if false {
		} else {
			z.F.DecMapFloat64Float64X(yyv1525, d)
		}
	}
	if x.FptrMapFloat64Float64 == nil {
		x.FptrMapFloat64Float64 = new(map[float64]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Float64 != nil {
			x.FptrMapFloat64Float64 = nil
		}
	} else {
		if x.FptrMapFloat64Float64 == nil {
			x.FptrMapFloat64Float64 = new(map[float64]float64)
		}
		yym1528 := z.DecBinary()
		_ = yym1528
		if false {
		} else {
			z.F.DecMapFloat64Float64X(x.FptrMapFloat64Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapFloat64Bool = nil
	} else {
		yyv1529 := &x.FMapFloat64Bool
		yym1530 := z.DecBinary()
		_ = yym1530
		if false {
		} else {
			z.F.DecMapFloat64BoolX(yyv1529, d)
		}
	}
	if x.FptrMapFloat64Bool == nil {
		x.FptrMapFloat64Bool = new(map[float64]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapFloat64Bool != nil {
			x.FptrMapFloat64Bool = nil
		}
	} else {
		if x.FptrMapFloat64Bool == nil {
			x.FptrMapFloat64Bool = new(map[float64]bool)
		}
		yym1532 := z.DecBinary()
		_ = yym1532
		if false {
		} else {
			z.F.DecMapFloat64BoolX(x.FptrMapFloat64Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintIntf = nil
	} else {
		yyv1533 := &x.FMapUintIntf
		yym1534 := z.DecBinary()
		_ = yym1534
		if false {
		} else {
			z.F.DecMapUintIntfX(yyv1533, d)
		}
	}
	if x.FptrMapUintIntf == nil {
		x.FptrMapUintIntf = new(map[uint]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintIntf != nil {
			x.FptrMapUintIntf = nil
		}
	} else {
		if x.FptrMapUintIntf == nil {
			x.FptrMapUintIntf = new(map[uint]interface{})
		}
		yym1536 := z.DecBinary()
		_ = yym1536
		if false {
		} else {
			z.F.DecMapUintIntfX(x.FptrMapUintIntf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintString = nil
	} else {
		yyv1537 := &x.FMapUintString
		yym1538 := z.DecBinary()
		_ = yym1538
		if false {
		} else {
			z.F.DecMapUintStringX(yyv1537, d)
		}
	}
	if x.FptrMapUintString == nil {
		x.FptrMapUintString = new(map[uint]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintString != nil {
			x.FptrMapUintString = nil
		}
	} else {
		if x.FptrMapUintString == nil {
			x.FptrMapUintString = new(map[uint]string)
		}
		yym1540 := z.DecBinary()
		_ = yym1540
		if false {
		} else {
			z.F.DecMapUintStringX(x.FptrMapUintString, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintUint = nil
	} else {
		yyv1541 := &x.FMapUintUint
		yym1542 := z.DecBinary()
		_ = yym1542
		if false {
		} else {
			z.F.DecMapUintUintX(yyv1541, d)
		}
	}
	if x.FptrMapUintUint == nil {
		x.FptrMapUintUint = new(map[uint]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintUint != nil {
			x.FptrMapUintUint = nil
		}
	} else {
		if x.FptrMapUintUint == nil {
			x.FptrMapUintUint = new(map[uint]uint)
		}
		yym1544 := z.DecBinary()
		_ = yym1544
		if false {
		} else {
			z.F.DecMapUintUintX(x.FptrMapUintUint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintUint8 = nil
	} else {
		yyv1545 := &x.FMapUintUint8
		yym1546 := z.DecBinary()
		_ = yym1546
		if false {
		} else {
			z.F.DecMapUintUint8X(yyv1545, d)
		}
	}
	if x.FptrMapUintUint8 == nil {
		x.FptrMapUintUint8 = new(map[uint]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintUint8 != nil {
			x.FptrMapUintUint8 = nil
		}
	} else {
		if x.FptrMapUintUint8 == nil {
			x.FptrMapUintUint8 = new(map[uint]uint8)
		}
		yym1548 := z.DecBinary()
		_ = yym1548
		if false {
		} else {
			z.F.DecMapUintUint8X(x.FptrMapUintUint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintUint16 = nil
	} else {
		yyv1549 := &x.FMapUintUint16
		yym1550 := z.DecBinary()
		_ = yym1550
		if false {
		} else {
			z.F.DecMapUintUint16X(yyv1549, d)
		}
	}
	if x.FptrMapUintUint16 == nil {
		x.FptrMapUintUint16 = new(map[uint]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintUint16 != nil {
			x.FptrMapUintUint16 = nil
		}
	} else {
		if x.FptrMapUintUint16 == nil {
			x.FptrMapUintUint16 = new(map[uint]uint16)
		}
		yym1552 := z.DecBinary()
		_ = yym1552
		if false {
		} else {
			z.F.DecMapUintUint16X(x.FptrMapUintUint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintUint32 = nil
	} else {
		yyv1553 := &x.FMapUintUint32
		yym1554 := z.DecBinary()
		_ = yym1554
		if false {
		} else {
			z.F.DecMapUintUint32X(yyv1553, d)
		}
	}
	if x.FptrMapUintUint32 == nil {
		x.FptrMapUintUint32 = new(map[uint]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintUint32 != nil {
			x.FptrMapUintUint32 = nil
		}
	} else {
		if x.FptrMapUintUint32 == nil {
			x.FptrMapUintUint32 = new(map[uint]uint32)
		}
		yym1556 := z.DecBinary()
		_ = yym1556
		if false {
		} else {
			z.F.DecMapUintUint32X(x.FptrMapUintUint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintUint64 = nil
	} else {
		yyv1557 := &x.FMapUintUint64
		yym1558 := z.DecBinary()
		_ = yym1558
		if false {
		} else {
			z.F.DecMapUintUint64X(yyv1557, d)
		}
	}
	if x.FptrMapUintUint64 == nil {
		x.FptrMapUintUint64 = new(map[uint]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintUint64 != nil {
			x.FptrMapUintUint64 = nil
		}
	} else {
		if x.FptrMapUintUint64 == nil {
			x.FptrMapUintUint64 = new(map[uint]uint64)
		}
		yym1560 := z.DecBinary()
		_ = yym1560
		if false {
		} else {
			z.F.DecMapUintUint64X(x.FptrMapUintUint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintUintptr = nil
	} else {
		yyv1561 := &x.FMapUintUintptr
		yym1562 := z.DecBinary()
		_ = yym1562
		if false {
		} else {
			z.F.DecMapUintUintptrX(yyv1561, d)
		}
	}
	if x.FptrMapUintUintptr == nil {
		x.FptrMapUintUintptr = new(map[uint]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintUintptr != nil {
			x.FptrMapUintUintptr = nil
		}
	} else {
		if x.FptrMapUintUintptr == nil {
			x.FptrMapUintUintptr = new(map[uint]uintptr)
		}
		yym1564 := z.DecBinary()
		_ = yym1564
		if false {
		} else {
			z.F.DecMapUintUintptrX(x.FptrMapUintUintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintInt = nil
	} else {
		yyv1565 := &x.FMapUintInt
		yym1566 := z.DecBinary()
		_ = yym1566
		if false {
		} else {
			z.F.DecMapUintIntX(yyv1565, d)
		}
	}
	if x.FptrMapUintInt == nil {
		x.FptrMapUintInt = new(map[uint]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintInt != nil {
			x.FptrMapUintInt = nil
		}
	} else {
		if x.FptrMapUintInt == nil {
			x.FptrMapUintInt = new(map[uint]int)
		}
		yym1568 := z.DecBinary()
		_ = yym1568
		if false {
		} else {
			z.F.DecMapUintIntX(x.FptrMapUintInt, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintInt8 = nil
	} else {
		yyv1569 := &x.FMapUintInt8
		yym1570 := z.DecBinary()
		_ = yym1570
		if false {
		} else {
			z.F.DecMapUintInt8X(yyv1569, d)
		}
	}
	if x.FptrMapUintInt8 == nil {
		x.FptrMapUintInt8 = new(map[uint]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintInt8 != nil {
			x.FptrMapUintInt8 = nil
		}
	} else {
		if x.FptrMapUintInt8 == nil {
			x.FptrMapUintInt8 = new(map[uint]int8)
		}
		yym1572 := z.DecBinary()
		_ = yym1572
		if false {
		} else {
			z.F.DecMapUintInt8X(x.FptrMapUintInt8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintInt16 = nil
	} else {
		yyv1573 := &x.FMapUintInt16
		yym1574 := z.DecBinary()
		_ = yym1574
		if false {
		} else {
			z.F.DecMapUintInt16X(yyv1573, d)
		}
	}
	if x.FptrMapUintInt16 == nil {
		x.FptrMapUintInt16 = new(map[uint]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintInt16 != nil {
			x.FptrMapUintInt16 = nil
		}
	} else {
		if x.FptrMapUintInt16 == nil {
			x.FptrMapUintInt16 = new(map[uint]int16)
		}
		yym1576 := z.DecBinary()
		_ = yym1576
		if false {
		} else {
			z.F.DecMapUintInt16X(x.FptrMapUintInt16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintInt32 = nil
	} else {
		yyv1577 := &x.FMapUintInt32
		yym1578 := z.DecBinary()
		_ = yym1578
		if false {
		} else {
			z.F.DecMapUintInt32X(yyv1577, d)
		}
	}
	if x.FptrMapUintInt32 == nil {
		x.FptrMapUintInt32 = new(map[uint]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintInt32 != nil {
			x.FptrMapUintInt32 = nil
		}
	} else {
		if x.FptrMapUintInt32 == nil {
			x.FptrMapUintInt32 = new(map[uint]int32)
		}
		yym1580 := z.DecBinary()
		_ = yym1580
		if false {
		} else {
			z.F.DecMapUintInt32X(x.FptrMapUintInt32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintInt64 = nil
	} else {
		yyv1581 := &x.FMapUintInt64
		yym1582 := z.DecBinary()
		_ = yym1582
		if false {
		} else {
			z.F.DecMapUintInt64X(yyv1581, d)
		}
	}
	if x.FptrMapUintInt64 == nil {
		x.FptrMapUintInt64 = new(map[uint]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintInt64 != nil {
			x.FptrMapUintInt64 = nil
		}
	} else {
		if x.FptrMapUintInt64 == nil {
			x.FptrMapUintInt64 = new(map[uint]int64)
		}
		yym1584 := z.DecBinary()
		_ = yym1584
		if false {
		} else {
			z.F.DecMapUintInt64X(x.FptrMapUintInt64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintFloat32 = nil
	} else {
		yyv1585 := &x.FMapUintFloat32
		yym1586 := z.DecBinary()
		_ = yym1586
		if false {
		} else {
			z.F.DecMapUintFloat32X(yyv1585, d)
		}
	}
	if x.FptrMapUintFloat32 == nil {
		x.FptrMapUintFloat32 = new(map[uint]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintFloat32 != nil {
			x.FptrMapUintFloat32 = nil
		}
	} else {
		if x.FptrMapUintFloat32 == nil {
			x.FptrMapUintFloat32 = new(map[uint]float32)
		}
		yym1588 := z.DecBinary()
		_ = yym1588
		if false {
		} else {
			z.F.DecMapUintFloat32X(x.FptrMapUintFloat32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintFloat64 = nil
	} else {
		yyv1589 := &x.FMapUintFloat64
		yym1590 := z.DecBinary()
		_ = yym1590
		if false {
		} else {
			z.F.DecMapUintFloat64X(yyv1589, d)
		}
	}
	if x.FptrMapUintFloat64 == nil {
		x.FptrMapUintFloat64 = new(map[uint]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintFloat64 != nil {
			x.FptrMapUintFloat64 = nil
		}
	} else {
		if x.FptrMapUintFloat64 == nil {
			x.FptrMapUintFloat64 = new(map[uint]float64)
		}
		yym1592 := z.DecBinary()
		_ = yym1592
		if false {
		} else {
			z.F.DecMapUintFloat64X(x.FptrMapUintFloat64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintBool = nil
	} else {
		yyv1593 := &x.FMapUintBool
		yym1594 := z.DecBinary()
		_ = yym1594
		if false {
		} else {
			z.F.DecMapUintBoolX(yyv1593, d)
		}
	}
	if x.FptrMapUintBool == nil {
		x.FptrMapUintBool = new(map[uint]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintBool != nil {
			x.FptrMapUintBool = nil
		}
	} else {
		if x.FptrMapUintBool == nil {
			x.FptrMapUintBool = new(map[uint]bool)
		}
		yym1596 := z.DecBinary()
		_ = yym1596
		if false {
		} else {
			z.F.DecMapUintBoolX(x.FptrMapUintBool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Intf = nil
	} else {
		yyv1597 := &x.FMapUint8Intf
		yym1598 := z.DecBinary()
		_ = yym1598
		if false {
		} else {
			z.F.DecMapUint8IntfX(yyv1597, d)
		}
	}
	if x.FptrMapUint8Intf == nil {
		x.FptrMapUint8Intf = new(map[uint8]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Intf != nil {
			x.FptrMapUint8Intf = nil
		}
	} else {
		if x.FptrMapUint8Intf == nil {
			x.FptrMapUint8Intf = new(map[uint8]interface{})
		}
		yym1600 := z.DecBinary()
		_ = yym1600
		if false {
		} else {
			z.F.DecMapUint8IntfX(x.FptrMapUint8Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8String = nil
	} else {
		yyv1601 := &x.FMapUint8String
		yym1602 := z.DecBinary()
		_ = yym1602
		if false {
		} else {
			z.F.DecMapUint8StringX(yyv1601, d)
		}
	}
	if x.FptrMapUint8String == nil {
		x.FptrMapUint8String = new(map[uint8]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8String != nil {
			x.FptrMapUint8String = nil
		}
	} else {
		if x.FptrMapUint8String == nil {
			x.FptrMapUint8String = new(map[uint8]string)
		}
		yym1604 := z.DecBinary()
		_ = yym1604
		if false {
		} else {
			z.F.DecMapUint8StringX(x.FptrMapUint8String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Uint = nil
	} else {
		yyv1605 := &x.FMapUint8Uint
		yym1606 := z.DecBinary()
		_ = yym1606
		if false {
		} else {
			z.F.DecMapUint8UintX(yyv1605, d)
		}
	}
	if x.FptrMapUint8Uint == nil {
		x.FptrMapUint8Uint = new(map[uint8]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Uint != nil {
			x.FptrMapUint8Uint = nil
		}
	} else {
		if x.FptrMapUint8Uint == nil {
			x.FptrMapUint8Uint = new(map[uint8]uint)
		}
		yym1608 := z.DecBinary()
		_ = yym1608
		if false {
		} else {
			z.F.DecMapUint8UintX(x.FptrMapUint8Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Uint8 = nil
	} else {
		yyv1609 := &x.FMapUint8Uint8
		yym1610 := z.DecBinary()
		_ = yym1610
		if false {
		} else {
			z.F.DecMapUint8Uint8X(yyv1609, d)
		}
	}
	if x.FptrMapUint8Uint8 == nil {
		x.FptrMapUint8Uint8 = new(map[uint8]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Uint8 != nil {
			x.FptrMapUint8Uint8 = nil
		}
	} else {
		if x.FptrMapUint8Uint8 == nil {
			x.FptrMapUint8Uint8 = new(map[uint8]uint8)
		}
		yym1612 := z.DecBinary()
		_ = yym1612
		if false {
		} else {
			z.F.DecMapUint8Uint8X(x.FptrMapUint8Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Uint16 = nil
	} else {
		yyv1613 := &x.FMapUint8Uint16
		yym1614 := z.DecBinary()
		_ = yym1614
		if false {
		} else {
			z.F.DecMapUint8Uint16X(yyv1613, d)
		}
	}
	if x.FptrMapUint8Uint16 == nil {
		x.FptrMapUint8Uint16 = new(map[uint8]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Uint16 != nil {
			x.FptrMapUint8Uint16 = nil
		}
	} else {
		if x.FptrMapUint8Uint16 == nil {
			x.FptrMapUint8Uint16 = new(map[uint8]uint16)
		}
		yym1616 := z.DecBinary()
		_ = yym1616
		if false {
		} else {
			z.F.DecMapUint8Uint16X(x.FptrMapUint8Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Uint32 = nil
	} else {
		yyv1617 := &x.FMapUint8Uint32
		yym1618 := z.DecBinary()
		_ = yym1618
		if false {
		} else {
			z.F.DecMapUint8Uint32X(yyv1617, d)
		}
	}
	if x.FptrMapUint8Uint32 == nil {
		x.FptrMapUint8Uint32 = new(map[uint8]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Uint32 != nil {
			x.FptrMapUint8Uint32 = nil
		}
	} else {
		if x.FptrMapUint8Uint32 == nil {
			x.FptrMapUint8Uint32 = new(map[uint8]uint32)
		}
		yym1620 := z.DecBinary()
		_ = yym1620
		if false {
		} else {
			z.F.DecMapUint8Uint32X(x.FptrMapUint8Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Uint64 = nil
	} else {
		yyv1621 := &x.FMapUint8Uint64
		yym1622 := z.DecBinary()
		_ = yym1622
		if false {
		} else {
			z.F.DecMapUint8Uint64X(yyv1621, d)
		}
	}
	if x.FptrMapUint8Uint64 == nil {
		x.FptrMapUint8Uint64 = new(map[uint8]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Uint64 != nil {
			x.FptrMapUint8Uint64 = nil
		}
	} else {
		if x.FptrMapUint8Uint64 == nil {
			x.FptrMapUint8Uint64 = new(map[uint8]uint64)
		}
		yym1624 := z.DecBinary()
		_ = yym1624
		if false {
		} else {
			z.F.DecMapUint8Uint64X(x.FptrMapUint8Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Uintptr = nil
	} else {
		yyv1625 := &x.FMapUint8Uintptr
		yym1626 := z.DecBinary()
		_ = yym1626
		if false {
		} else {
			z.F.DecMapUint8UintptrX(yyv1625, d)
		}
	}
	if x.FptrMapUint8Uintptr == nil {
		x.FptrMapUint8Uintptr = new(map[uint8]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Uintptr != nil {
			x.FptrMapUint8Uintptr = nil
		}
	} else {
		if x.FptrMapUint8Uintptr == nil {
			x.FptrMapUint8Uintptr = new(map[uint8]uintptr)
		}
		yym1628 := z.DecBinary()
		_ = yym1628
		if false {
		} else {
			z.F.DecMapUint8UintptrX(x.FptrMapUint8Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Int = nil
	} else {
		yyv1629 := &x.FMapUint8Int
		yym1630 := z.DecBinary()
		_ = yym1630
		if false {
		} else {
			z.F.DecMapUint8IntX(yyv1629, d)
		}
	}
	if x.FptrMapUint8Int == nil {
		x.FptrMapUint8Int = new(map[uint8]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Int != nil {
			x.FptrMapUint8Int = nil
		}
	} else {
		if x.FptrMapUint8Int == nil {
			x.FptrMapUint8Int = new(map[uint8]int)
		}
		yym1632 := z.DecBinary()
		_ = yym1632
		if false {
		} else {
			z.F.DecMapUint8IntX(x.FptrMapUint8Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Int8 = nil
	} else {
		yyv1633 := &x.FMapUint8Int8
		yym1634 := z.DecBinary()
		_ = yym1634
		if false {
		} else {
			z.F.DecMapUint8Int8X(yyv1633, d)
		}
	}
	if x.FptrMapUint8Int8 == nil {
		x.FptrMapUint8Int8 = new(map[uint8]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Int8 != nil {
			x.FptrMapUint8Int8 = nil
		}
	} else {
		if x.FptrMapUint8Int8 == nil {
			x.FptrMapUint8Int8 = new(map[uint8]int8)
		}
		yym1636 := z.DecBinary()
		_ = yym1636
		if false {
		} else {
			z.F.DecMapUint8Int8X(x.FptrMapUint8Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Int16 = nil
	} else {
		yyv1637 := &x.FMapUint8Int16
		yym1638 := z.DecBinary()
		_ = yym1638
		if false {
		} else {
			z.F.DecMapUint8Int16X(yyv1637, d)
		}
	}
	if x.FptrMapUint8Int16 == nil {
		x.FptrMapUint8Int16 = new(map[uint8]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Int16 != nil {
			x.FptrMapUint8Int16 = nil
		}
	} else {
		if x.FptrMapUint8Int16 == nil {
			x.FptrMapUint8Int16 = new(map[uint8]int16)
		}
		yym1640 := z.DecBinary()
		_ = yym1640
		if false {
		} else {
			z.F.DecMapUint8Int16X(x.FptrMapUint8Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Int32 = nil
	} else {
		yyv1641 := &x.FMapUint8Int32
		yym1642 := z.DecBinary()
		_ = yym1642
		if false {
		} else {
			z.F.DecMapUint8Int32X(yyv1641, d)
		}
	}
	if x.FptrMapUint8Int32 == nil {
		x.FptrMapUint8Int32 = new(map[uint8]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Int32 != nil {
			x.FptrMapUint8Int32 = nil
		}
	} else {
		if x.FptrMapUint8Int32 == nil {
			x.FptrMapUint8Int32 = new(map[uint8]int32)
		}
		yym1644 := z.DecBinary()
		_ = yym1644
		if false {
		} else {
			z.F.DecMapUint8Int32X(x.FptrMapUint8Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Int64 = nil
	} else {
		yyv1645 := &x.FMapUint8Int64
		yym1646 := z.DecBinary()
		_ = yym1646
		if false {
		} else {
			z.F.DecMapUint8Int64X(yyv1645, d)
		}
	}
	if x.FptrMapUint8Int64 == nil {
		x.FptrMapUint8Int64 = new(map[uint8]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Int64 != nil {
			x.FptrMapUint8Int64 = nil
		}
	} else {
		if x.FptrMapUint8Int64 == nil {
			x.FptrMapUint8Int64 = new(map[uint8]int64)
		}
		yym1648 := z.DecBinary()
		_ = yym1648
		if false {
		} else {
			z.F.DecMapUint8Int64X(x.FptrMapUint8Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Float32 = nil
	} else {
		yyv1649 := &x.FMapUint8Float32
		yym1650 := z.DecBinary()
		_ = yym1650
		if false {
		} else {
			z.F.DecMapUint8Float32X(yyv1649, d)
		}
	}
	if x.FptrMapUint8Float32 == nil {
		x.FptrMapUint8Float32 = new(map[uint8]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Float32 != nil {
			x.FptrMapUint8Float32 = nil
		}
	} else {
		if x.FptrMapUint8Float32 == nil {
			x.FptrMapUint8Float32 = new(map[uint8]float32)
		}
		yym1652 := z.DecBinary()
		_ = yym1652
		if false {
		} else {
			z.F.DecMapUint8Float32X(x.FptrMapUint8Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Float64 = nil
	} else {
		yyv1653 := &x.FMapUint8Float64
		yym1654 := z.DecBinary()
		_ = yym1654
		if false {
		} else {
			z.F.DecMapUint8Float64X(yyv1653, d)
		}
	}
	if x.FptrMapUint8Float64 == nil {
		x.FptrMapUint8Float64 = new(map[uint8]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Float64 != nil {
			x.FptrMapUint8Float64 = nil
		}
	} else {
		if x.FptrMapUint8Float64 == nil {
			x.FptrMapUint8Float64 = new(map[uint8]float64)
		}
		yym1656 := z.DecBinary()
		_ = yym1656
		if false {
		} else {
			z.F.DecMapUint8Float64X(x.FptrMapUint8Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint8Bool = nil
	} else {
		yyv1657 := &x.FMapUint8Bool
		yym1658 := z.DecBinary()
		_ = yym1658
		if false {
		} else {
			z.F.DecMapUint8BoolX(yyv1657, d)
		}
	}
	if x.FptrMapUint8Bool == nil {
		x.FptrMapUint8Bool = new(map[uint8]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint8Bool != nil {
			x.FptrMapUint8Bool = nil
		}
	} else {
		if x.FptrMapUint8Bool == nil {
			x.FptrMapUint8Bool = new(map[uint8]bool)
		}
		yym1660 := z.DecBinary()
		_ = yym1660
		if false {
		} else {
			z.F.DecMapUint8BoolX(x.FptrMapUint8Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Intf = nil
	} else {
		yyv1661 := &x.FMapUint16Intf
		yym1662 := z.DecBinary()
		_ = yym1662
		if false {
		} else {
			z.F.DecMapUint16IntfX(yyv1661, d)
		}
	}
	if x.FptrMapUint16Intf == nil {
		x.FptrMapUint16Intf = new(map[uint16]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Intf != nil {
			x.FptrMapUint16Intf = nil
		}
	} else {
		if x.FptrMapUint16Intf == nil {
			x.FptrMapUint16Intf = new(map[uint16]interface{})
		}
		yym1664 := z.DecBinary()
		_ = yym1664
		if false {
		} else {
			z.F.DecMapUint16IntfX(x.FptrMapUint16Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16String = nil
	} else {
		yyv1665 := &x.FMapUint16String
		yym1666 := z.DecBinary()
		_ = yym1666
		if false {
		} else {
			z.F.DecMapUint16StringX(yyv1665, d)
		}
	}
	if x.FptrMapUint16String == nil {
		x.FptrMapUint16String = new(map[uint16]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16String != nil {
			x.FptrMapUint16String = nil
		}
	} else {
		if x.FptrMapUint16String == nil {
			x.FptrMapUint16String = new(map[uint16]string)
		}
		yym1668 := z.DecBinary()
		_ = yym1668
		if false {
		} else {
			z.F.DecMapUint16StringX(x.FptrMapUint16String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Uint = nil
	} else {
		yyv1669 := &x.FMapUint16Uint
		yym1670 := z.DecBinary()
		_ = yym1670
		if false {
		} else {
			z.F.DecMapUint16UintX(yyv1669, d)
		}
	}
	if x.FptrMapUint16Uint == nil {
		x.FptrMapUint16Uint = new(map[uint16]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Uint != nil {
			x.FptrMapUint16Uint = nil
		}
	} else {
		if x.FptrMapUint16Uint == nil {
			x.FptrMapUint16Uint = new(map[uint16]uint)
		}
		yym1672 := z.DecBinary()
		_ = yym1672
		if false {
		} else {
			z.F.DecMapUint16UintX(x.FptrMapUint16Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Uint8 = nil
	} else {
		yyv1673 := &x.FMapUint16Uint8
		yym1674 := z.DecBinary()
		_ = yym1674
		if false {
		} else {
			z.F.DecMapUint16Uint8X(yyv1673, d)
		}
	}
	if x.FptrMapUint16Uint8 == nil {
		x.FptrMapUint16Uint8 = new(map[uint16]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Uint8 != nil {
			x.FptrMapUint16Uint8 = nil
		}
	} else {
		if x.FptrMapUint16Uint8 == nil {
			x.FptrMapUint16Uint8 = new(map[uint16]uint8)
		}
		yym1676 := z.DecBinary()
		_ = yym1676
		if false {
		} else {
			z.F.DecMapUint16Uint8X(x.FptrMapUint16Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Uint16 = nil
	} else {
		yyv1677 := &x.FMapUint16Uint16
		yym1678 := z.DecBinary()
		_ = yym1678
		if false {
		} else {
			z.F.DecMapUint16Uint16X(yyv1677, d)
		}
	}
	if x.FptrMapUint16Uint16 == nil {
		x.FptrMapUint16Uint16 = new(map[uint16]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Uint16 != nil {
			x.FptrMapUint16Uint16 = nil
		}
	} else {
		if x.FptrMapUint16Uint16 == nil {
			x.FptrMapUint16Uint16 = new(map[uint16]uint16)
		}
		yym1680 := z.DecBinary()
		_ = yym1680
		if false {
		} else {
			z.F.DecMapUint16Uint16X(x.FptrMapUint16Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Uint32 = nil
	} else {
		yyv1681 := &x.FMapUint16Uint32
		yym1682 := z.DecBinary()
		_ = yym1682
		if false {
		} else {
			z.F.DecMapUint16Uint32X(yyv1681, d)
		}
	}
	if x.FptrMapUint16Uint32 == nil {
		x.FptrMapUint16Uint32 = new(map[uint16]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Uint32 != nil {
			x.FptrMapUint16Uint32 = nil
		}
	} else {
		if x.FptrMapUint16Uint32 == nil {
			x.FptrMapUint16Uint32 = new(map[uint16]uint32)
		}
		yym1684 := z.DecBinary()
		_ = yym1684
		if false {
		} else {
			z.F.DecMapUint16Uint32X(x.FptrMapUint16Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Uint64 = nil
	} else {
		yyv1685 := &x.FMapUint16Uint64
		yym1686 := z.DecBinary()
		_ = yym1686
		if false {
		} else {
			z.F.DecMapUint16Uint64X(yyv1685, d)
		}
	}
	if x.FptrMapUint16Uint64 == nil {
		x.FptrMapUint16Uint64 = new(map[uint16]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Uint64 != nil {
			x.FptrMapUint16Uint64 = nil
		}
	} else {
		if x.FptrMapUint16Uint64 == nil {
			x.FptrMapUint16Uint64 = new(map[uint16]uint64)
		}
		yym1688 := z.DecBinary()
		_ = yym1688
		if false {
		} else {
			z.F.DecMapUint16Uint64X(x.FptrMapUint16Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Uintptr = nil
	} else {
		yyv1689 := &x.FMapUint16Uintptr
		yym1690 := z.DecBinary()
		_ = yym1690
		if false {
		} else {
			z.F.DecMapUint16UintptrX(yyv1689, d)
		}
	}
	if x.FptrMapUint16Uintptr == nil {
		x.FptrMapUint16Uintptr = new(map[uint16]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Uintptr != nil {
			x.FptrMapUint16Uintptr = nil
		}
	} else {
		if x.FptrMapUint16Uintptr == nil {
			x.FptrMapUint16Uintptr = new(map[uint16]uintptr)
		}
		yym1692 := z.DecBinary()
		_ = yym1692
		if false {
		} else {
			z.F.DecMapUint16UintptrX(x.FptrMapUint16Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Int = nil
	} else {
		yyv1693 := &x.FMapUint16Int
		yym1694 := z.DecBinary()
		_ = yym1694
		if false {
		} else {
			z.F.DecMapUint16IntX(yyv1693, d)
		}
	}
	if x.FptrMapUint16Int == nil {
		x.FptrMapUint16Int = new(map[uint16]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Int != nil {
			x.FptrMapUint16Int = nil
		}
	} else {
		if x.FptrMapUint16Int == nil {
			x.FptrMapUint16Int = new(map[uint16]int)
		}
		yym1696 := z.DecBinary()
		_ = yym1696
		if false {
		} else {
			z.F.DecMapUint16IntX(x.FptrMapUint16Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Int8 = nil
	} else {
		yyv1697 := &x.FMapUint16Int8
		yym1698 := z.DecBinary()
		_ = yym1698
		if false {
		} else {
			z.F.DecMapUint16Int8X(yyv1697, d)
		}
	}
	if x.FptrMapUint16Int8 == nil {
		x.FptrMapUint16Int8 = new(map[uint16]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Int8 != nil {
			x.FptrMapUint16Int8 = nil
		}
	} else {
		if x.FptrMapUint16Int8 == nil {
			x.FptrMapUint16Int8 = new(map[uint16]int8)
		}
		yym1700 := z.DecBinary()
		_ = yym1700
		if false {
		} else {
			z.F.DecMapUint16Int8X(x.FptrMapUint16Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Int16 = nil
	} else {
		yyv1701 := &x.FMapUint16Int16
		yym1702 := z.DecBinary()
		_ = yym1702
		if false {
		} else {
			z.F.DecMapUint16Int16X(yyv1701, d)
		}
	}
	if x.FptrMapUint16Int16 == nil {
		x.FptrMapUint16Int16 = new(map[uint16]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Int16 != nil {
			x.FptrMapUint16Int16 = nil
		}
	} else {
		if x.FptrMapUint16Int16 == nil {
			x.FptrMapUint16Int16 = new(map[uint16]int16)
		}
		yym1704 := z.DecBinary()
		_ = yym1704
		if false {
		} else {
			z.F.DecMapUint16Int16X(x.FptrMapUint16Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Int32 = nil
	} else {
		yyv1705 := &x.FMapUint16Int32
		yym1706 := z.DecBinary()
		_ = yym1706
		if false {
		} else {
			z.F.DecMapUint16Int32X(yyv1705, d)
		}
	}
	if x.FptrMapUint16Int32 == nil {
		x.FptrMapUint16Int32 = new(map[uint16]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Int32 != nil {
			x.FptrMapUint16Int32 = nil
		}
	} else {
		if x.FptrMapUint16Int32 == nil {
			x.FptrMapUint16Int32 = new(map[uint16]int32)
		}
		yym1708 := z.DecBinary()
		_ = yym1708
		if false {
		} else {
			z.F.DecMapUint16Int32X(x.FptrMapUint16Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Int64 = nil
	} else {
		yyv1709 := &x.FMapUint16Int64
		yym1710 := z.DecBinary()
		_ = yym1710
		if false {
		} else {
			z.F.DecMapUint16Int64X(yyv1709, d)
		}
	}
	if x.FptrMapUint16Int64 == nil {
		x.FptrMapUint16Int64 = new(map[uint16]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Int64 != nil {
			x.FptrMapUint16Int64 = nil
		}
	} else {
		if x.FptrMapUint16Int64 == nil {
			x.FptrMapUint16Int64 = new(map[uint16]int64)
		}
		yym1712 := z.DecBinary()
		_ = yym1712
		if false {
		} else {
			z.F.DecMapUint16Int64X(x.FptrMapUint16Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Float32 = nil
	} else {
		yyv1713 := &x.FMapUint16Float32
		yym1714 := z.DecBinary()
		_ = yym1714
		if false {
		} else {
			z.F.DecMapUint16Float32X(yyv1713, d)
		}
	}
	if x.FptrMapUint16Float32 == nil {
		x.FptrMapUint16Float32 = new(map[uint16]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Float32 != nil {
			x.FptrMapUint16Float32 = nil
		}
	} else {
		if x.FptrMapUint16Float32 == nil {
			x.FptrMapUint16Float32 = new(map[uint16]float32)
		}
		yym1716 := z.DecBinary()
		_ = yym1716
		if false {
		} else {
			z.F.DecMapUint16Float32X(x.FptrMapUint16Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Float64 = nil
	} else {
		yyv1717 := &x.FMapUint16Float64
		yym1718 := z.DecBinary()
		_ = yym1718
		if false {
		} else {
			z.F.DecMapUint16Float64X(yyv1717, d)
		}
	}
	if x.FptrMapUint16Float64 == nil {
		x.FptrMapUint16Float64 = new(map[uint16]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Float64 != nil {
			x.FptrMapUint16Float64 = nil
		}
	} else {
		if x.FptrMapUint16Float64 == nil {
			x.FptrMapUint16Float64 = new(map[uint16]float64)
		}
		yym1720 := z.DecBinary()
		_ = yym1720
		if false {
		} else {
			z.F.DecMapUint16Float64X(x.FptrMapUint16Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint16Bool = nil
	} else {
		yyv1721 := &x.FMapUint16Bool
		yym1722 := z.DecBinary()
		_ = yym1722
		if false {
		} else {
			z.F.DecMapUint16BoolX(yyv1721, d)
		}
	}
	if x.FptrMapUint16Bool == nil {
		x.FptrMapUint16Bool = new(map[uint16]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint16Bool != nil {
			x.FptrMapUint16Bool = nil
		}
	} else {
		if x.FptrMapUint16Bool == nil {
			x.FptrMapUint16Bool = new(map[uint16]bool)
		}
		yym1724 := z.DecBinary()
		_ = yym1724
		if false {
		} else {
			z.F.DecMapUint16BoolX(x.FptrMapUint16Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Intf = nil
	} else {
		yyv1725 := &x.FMapUint32Intf
		yym1726 := z.DecBinary()
		_ = yym1726
		if false {
		} else {
			z.F.DecMapUint32IntfX(yyv1725, d)
		}
	}
	if x.FptrMapUint32Intf == nil {
		x.FptrMapUint32Intf = new(map[uint32]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Intf != nil {
			x.FptrMapUint32Intf = nil
		}
	} else {
		if x.FptrMapUint32Intf == nil {
			x.FptrMapUint32Intf = new(map[uint32]interface{})
		}
		yym1728 := z.DecBinary()
		_ = yym1728
		if false {
		} else {
			z.F.DecMapUint32IntfX(x.FptrMapUint32Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32String = nil
	} else {
		yyv1729 := &x.FMapUint32String
		yym1730 := z.DecBinary()
		_ = yym1730
		if false {
		} else {
			z.F.DecMapUint32StringX(yyv1729, d)
		}
	}
	if x.FptrMapUint32String == nil {
		x.FptrMapUint32String = new(map[uint32]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32String != nil {
			x.FptrMapUint32String = nil
		}
	} else {
		if x.FptrMapUint32String == nil {
			x.FptrMapUint32String = new(map[uint32]string)
		}
		yym1732 := z.DecBinary()
		_ = yym1732
		if false {
		} else {
			z.F.DecMapUint32StringX(x.FptrMapUint32String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Uint = nil
	} else {
		yyv1733 := &x.FMapUint32Uint
		yym1734 := z.DecBinary()
		_ = yym1734
		if false {
		} else {
			z.F.DecMapUint32UintX(yyv1733, d)
		}
	}
	if x.FptrMapUint32Uint == nil {
		x.FptrMapUint32Uint = new(map[uint32]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Uint != nil {
			x.FptrMapUint32Uint = nil
		}
	} else {
		if x.FptrMapUint32Uint == nil {
			x.FptrMapUint32Uint = new(map[uint32]uint)
		}
		yym1736 := z.DecBinary()
		_ = yym1736
		if false {
		} else {
			z.F.DecMapUint32UintX(x.FptrMapUint32Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Uint8 = nil
	} else {
		yyv1737 := &x.FMapUint32Uint8
		yym1738 := z.DecBinary()
		_ = yym1738
		if false {
		} else {
			z.F.DecMapUint32Uint8X(yyv1737, d)
		}
	}
	if x.FptrMapUint32Uint8 == nil {
		x.FptrMapUint32Uint8 = new(map[uint32]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Uint8 != nil {
			x.FptrMapUint32Uint8 = nil
		}
	} else {
		if x.FptrMapUint32Uint8 == nil {
			x.FptrMapUint32Uint8 = new(map[uint32]uint8)
		}
		yym1740 := z.DecBinary()
		_ = yym1740
		if false {
		} else {
			z.F.DecMapUint32Uint8X(x.FptrMapUint32Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Uint16 = nil
	} else {
		yyv1741 := &x.FMapUint32Uint16
		yym1742 := z.DecBinary()
		_ = yym1742
		if false {
		} else {
			z.F.DecMapUint32Uint16X(yyv1741, d)
		}
	}
	if x.FptrMapUint32Uint16 == nil {
		x.FptrMapUint32Uint16 = new(map[uint32]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Uint16 != nil {
			x.FptrMapUint32Uint16 = nil
		}
	} else {
		if x.FptrMapUint32Uint16 == nil {
			x.FptrMapUint32Uint16 = new(map[uint32]uint16)
		}
		yym1744 := z.DecBinary()
		_ = yym1744
		if false {
		} else {
			z.F.DecMapUint32Uint16X(x.FptrMapUint32Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Uint32 = nil
	} else {
		yyv1745 := &x.FMapUint32Uint32
		yym1746 := z.DecBinary()
		_ = yym1746
		if false {
		} else {
			z.F.DecMapUint32Uint32X(yyv1745, d)
		}
	}
	if x.FptrMapUint32Uint32 == nil {
		x.FptrMapUint32Uint32 = new(map[uint32]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Uint32 != nil {
			x.FptrMapUint32Uint32 = nil
		}
	} else {
		if x.FptrMapUint32Uint32 == nil {
			x.FptrMapUint32Uint32 = new(map[uint32]uint32)
		}
		yym1748 := z.DecBinary()
		_ = yym1748
		if false {
		} else {
			z.F.DecMapUint32Uint32X(x.FptrMapUint32Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Uint64 = nil
	} else {
		yyv1749 := &x.FMapUint32Uint64
		yym1750 := z.DecBinary()
		_ = yym1750
		if false {
		} else {
			z.F.DecMapUint32Uint64X(yyv1749, d)
		}
	}
	if x.FptrMapUint32Uint64 == nil {
		x.FptrMapUint32Uint64 = new(map[uint32]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Uint64 != nil {
			x.FptrMapUint32Uint64 = nil
		}
	} else {
		if x.FptrMapUint32Uint64 == nil {
			x.FptrMapUint32Uint64 = new(map[uint32]uint64)
		}
		yym1752 := z.DecBinary()
		_ = yym1752
		if false {
		} else {
			z.F.DecMapUint32Uint64X(x.FptrMapUint32Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Uintptr = nil
	} else {
		yyv1753 := &x.FMapUint32Uintptr
		yym1754 := z.DecBinary()
		_ = yym1754
		if false {
		} else {
			z.F.DecMapUint32UintptrX(yyv1753, d)
		}
	}
	if x.FptrMapUint32Uintptr == nil {
		x.FptrMapUint32Uintptr = new(map[uint32]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Uintptr != nil {
			x.FptrMapUint32Uintptr = nil
		}
	} else {
		if x.FptrMapUint32Uintptr == nil {
			x.FptrMapUint32Uintptr = new(map[uint32]uintptr)
		}
		yym1756 := z.DecBinary()
		_ = yym1756
		if false {
		} else {
			z.F.DecMapUint32UintptrX(x.FptrMapUint32Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Int = nil
	} else {
		yyv1757 := &x.FMapUint32Int
		yym1758 := z.DecBinary()
		_ = yym1758
		if false {
		} else {
			z.F.DecMapUint32IntX(yyv1757, d)
		}
	}
	if x.FptrMapUint32Int == nil {
		x.FptrMapUint32Int = new(map[uint32]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Int != nil {
			x.FptrMapUint32Int = nil
		}
	} else {
		if x.FptrMapUint32Int == nil {
			x.FptrMapUint32Int = new(map[uint32]int)
		}
		yym1760 := z.DecBinary()
		_ = yym1760
		if false {
		} else {
			z.F.DecMapUint32IntX(x.FptrMapUint32Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Int8 = nil
	} else {
		yyv1761 := &x.FMapUint32Int8
		yym1762 := z.DecBinary()
		_ = yym1762
		if false {
		} else {
			z.F.DecMapUint32Int8X(yyv1761, d)
		}
	}
	if x.FptrMapUint32Int8 == nil {
		x.FptrMapUint32Int8 = new(map[uint32]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Int8 != nil {
			x.FptrMapUint32Int8 = nil
		}
	} else {
		if x.FptrMapUint32Int8 == nil {
			x.FptrMapUint32Int8 = new(map[uint32]int8)
		}
		yym1764 := z.DecBinary()
		_ = yym1764
		if false {
		} else {
			z.F.DecMapUint32Int8X(x.FptrMapUint32Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Int16 = nil
	} else {
		yyv1765 := &x.FMapUint32Int16
		yym1766 := z.DecBinary()
		_ = yym1766
		if false {
		} else {
			z.F.DecMapUint32Int16X(yyv1765, d)
		}
	}
	if x.FptrMapUint32Int16 == nil {
		x.FptrMapUint32Int16 = new(map[uint32]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Int16 != nil {
			x.FptrMapUint32Int16 = nil
		}
	} else {
		if x.FptrMapUint32Int16 == nil {
			x.FptrMapUint32Int16 = new(map[uint32]int16)
		}
		yym1768 := z.DecBinary()
		_ = yym1768
		if false {
		} else {
			z.F.DecMapUint32Int16X(x.FptrMapUint32Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Int32 = nil
	} else {
		yyv1769 := &x.FMapUint32Int32
		yym1770 := z.DecBinary()
		_ = yym1770
		if false {
		} else {
			z.F.DecMapUint32Int32X(yyv1769, d)
		}
	}
	if x.FptrMapUint32Int32 == nil {
		x.FptrMapUint32Int32 = new(map[uint32]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Int32 != nil {
			x.FptrMapUint32Int32 = nil
		}
	} else {
		if x.FptrMapUint32Int32 == nil {
			x.FptrMapUint32Int32 = new(map[uint32]int32)
		}
		yym1772 := z.DecBinary()
		_ = yym1772
		if false {
		} else {
			z.F.DecMapUint32Int32X(x.FptrMapUint32Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Int64 = nil
	} else {
		yyv1773 := &x.FMapUint32Int64
		yym1774 := z.DecBinary()
		_ = yym1774
		if false {
		} else {
			z.F.DecMapUint32Int64X(yyv1773, d)
		}
	}
	if x.FptrMapUint32Int64 == nil {
		x.FptrMapUint32Int64 = new(map[uint32]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Int64 != nil {
			x.FptrMapUint32Int64 = nil
		}
	} else {
		if x.FptrMapUint32Int64 == nil {
			x.FptrMapUint32Int64 = new(map[uint32]int64)
		}
		yym1776 := z.DecBinary()
		_ = yym1776
		if false {
		} else {
			z.F.DecMapUint32Int64X(x.FptrMapUint32Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Float32 = nil
	} else {
		yyv1777 := &x.FMapUint32Float32
		yym1778 := z.DecBinary()
		_ = yym1778
		if false {
		} else {
			z.F.DecMapUint32Float32X(yyv1777, d)
		}
	}
	if x.FptrMapUint32Float32 == nil {
		x.FptrMapUint32Float32 = new(map[uint32]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Float32 != nil {
			x.FptrMapUint32Float32 = nil
		}
	} else {
		if x.FptrMapUint32Float32 == nil {
			x.FptrMapUint32Float32 = new(map[uint32]float32)
		}
		yym1780 := z.DecBinary()
		_ = yym1780
		if false {
		} else {
			z.F.DecMapUint32Float32X(x.FptrMapUint32Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Float64 = nil
	} else {
		yyv1781 := &x.FMapUint32Float64
		yym1782 := z.DecBinary()
		_ = yym1782
		if false {
		} else {
			z.F.DecMapUint32Float64X(yyv1781, d)
		}
	}
	if x.FptrMapUint32Float64 == nil {
		x.FptrMapUint32Float64 = new(map[uint32]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Float64 != nil {
			x.FptrMapUint32Float64 = nil
		}
	} else {
		if x.FptrMapUint32Float64 == nil {
			x.FptrMapUint32Float64 = new(map[uint32]float64)
		}
		yym1784 := z.DecBinary()
		_ = yym1784
		if false {
		} else {
			z.F.DecMapUint32Float64X(x.FptrMapUint32Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint32Bool = nil
	} else {
		yyv1785 := &x.FMapUint32Bool
		yym1786 := z.DecBinary()
		_ = yym1786
		if false {
		} else {
			z.F.DecMapUint32BoolX(yyv1785, d)
		}
	}
	if x.FptrMapUint32Bool == nil {
		x.FptrMapUint32Bool = new(map[uint32]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint32Bool != nil {
			x.FptrMapUint32Bool = nil
		}
	} else {
		if x.FptrMapUint32Bool == nil {
			x.FptrMapUint32Bool = new(map[uint32]bool)
		}
		yym1788 := z.DecBinary()
		_ = yym1788
		if false {
		} else {
			z.F.DecMapUint32BoolX(x.FptrMapUint32Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Intf = nil
	} else {
		yyv1789 := &x.FMapUint64Intf
		yym1790 := z.DecBinary()
		_ = yym1790
		if false {
		} else {
			z.F.DecMapUint64IntfX(yyv1789, d)
		}
	}
	if x.FptrMapUint64Intf == nil {
		x.FptrMapUint64Intf = new(map[uint64]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Intf != nil {
			x.FptrMapUint64Intf = nil
		}
	} else {
		if x.FptrMapUint64Intf == nil {
			x.FptrMapUint64Intf = new(map[uint64]interface{})
		}
		yym1792 := z.DecBinary()
		_ = yym1792
		if false {
		} else {
			z.F.DecMapUint64IntfX(x.FptrMapUint64Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64String = nil
	} else {
		yyv1793 := &x.FMapUint64String
		yym1794 := z.DecBinary()
		_ = yym1794
		if false {
		} else {
			z.F.DecMapUint64StringX(yyv1793, d)
		}
	}
	if x.FptrMapUint64String == nil {
		x.FptrMapUint64String = new(map[uint64]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64String != nil {
			x.FptrMapUint64String = nil
		}
	} else {
		if x.FptrMapUint64String == nil {
			x.FptrMapUint64String = new(map[uint64]string)
		}
		yym1796 := z.DecBinary()
		_ = yym1796
		if false {
		} else {
			z.F.DecMapUint64StringX(x.FptrMapUint64String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Uint = nil
	} else {
		yyv1797 := &x.FMapUint64Uint
		yym1798 := z.DecBinary()
		_ = yym1798
		if false {
		} else {
			z.F.DecMapUint64UintX(yyv1797, d)
		}
	}
	if x.FptrMapUint64Uint == nil {
		x.FptrMapUint64Uint = new(map[uint64]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Uint != nil {
			x.FptrMapUint64Uint = nil
		}
	} else {
		if x.FptrMapUint64Uint == nil {
			x.FptrMapUint64Uint = new(map[uint64]uint)
		}
		yym1800 := z.DecBinary()
		_ = yym1800
		if false {
		} else {
			z.F.DecMapUint64UintX(x.FptrMapUint64Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Uint8 = nil
	} else {
		yyv1801 := &x.FMapUint64Uint8
		yym1802 := z.DecBinary()
		_ = yym1802
		if false {
		} else {
			z.F.DecMapUint64Uint8X(yyv1801, d)
		}
	}
	if x.FptrMapUint64Uint8 == nil {
		x.FptrMapUint64Uint8 = new(map[uint64]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Uint8 != nil {
			x.FptrMapUint64Uint8 = nil
		}
	} else {
		if x.FptrMapUint64Uint8 == nil {
			x.FptrMapUint64Uint8 = new(map[uint64]uint8)
		}
		yym1804 := z.DecBinary()
		_ = yym1804
		if false {
		} else {
			z.F.DecMapUint64Uint8X(x.FptrMapUint64Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Uint16 = nil
	} else {
		yyv1805 := &x.FMapUint64Uint16
		yym1806 := z.DecBinary()
		_ = yym1806
		if false {
		} else {
			z.F.DecMapUint64Uint16X(yyv1805, d)
		}
	}
	if x.FptrMapUint64Uint16 == nil {
		x.FptrMapUint64Uint16 = new(map[uint64]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Uint16 != nil {
			x.FptrMapUint64Uint16 = nil
		}
	} else {
		if x.FptrMapUint64Uint16 == nil {
			x.FptrMapUint64Uint16 = new(map[uint64]uint16)
		}
		yym1808 := z.DecBinary()
		_ = yym1808
		if false {
		} else {
			z.F.DecMapUint64Uint16X(x.FptrMapUint64Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Uint32 = nil
	} else {
		yyv1809 := &x.FMapUint64Uint32
		yym1810 := z.DecBinary()
		_ = yym1810
		if false {
		} else {
			z.F.DecMapUint64Uint32X(yyv1809, d)
		}
	}
	if x.FptrMapUint64Uint32 == nil {
		x.FptrMapUint64Uint32 = new(map[uint64]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Uint32 != nil {
			x.FptrMapUint64Uint32 = nil
		}
	} else {
		if x.FptrMapUint64Uint32 == nil {
			x.FptrMapUint64Uint32 = new(map[uint64]uint32)
		}
		yym1812 := z.DecBinary()
		_ = yym1812
		if false {
		} else {
			z.F.DecMapUint64Uint32X(x.FptrMapUint64Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Uint64 = nil
	} else {
		yyv1813 := &x.FMapUint64Uint64
		yym1814 := z.DecBinary()
		_ = yym1814
		if false {
		} else {
			z.F.DecMapUint64Uint64X(yyv1813, d)
		}
	}
	if x.FptrMapUint64Uint64 == nil {
		x.FptrMapUint64Uint64 = new(map[uint64]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Uint64 != nil {
			x.FptrMapUint64Uint64 = nil
		}
	} else {
		if x.FptrMapUint64Uint64 == nil {
			x.FptrMapUint64Uint64 = new(map[uint64]uint64)
		}
		yym1816 := z.DecBinary()
		_ = yym1816
		if false {
		} else {
			z.F.DecMapUint64Uint64X(x.FptrMapUint64Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Uintptr = nil
	} else {
		yyv1817 := &x.FMapUint64Uintptr
		yym1818 := z.DecBinary()
		_ = yym1818
		if false {
		} else {
			z.F.DecMapUint64UintptrX(yyv1817, d)
		}
	}
	if x.FptrMapUint64Uintptr == nil {
		x.FptrMapUint64Uintptr = new(map[uint64]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Uintptr != nil {
			x.FptrMapUint64Uintptr = nil
		}
	} else {
		if x.FptrMapUint64Uintptr == nil {
			x.FptrMapUint64Uintptr = new(map[uint64]uintptr)
		}
		yym1820 := z.DecBinary()
		_ = yym1820
		if false {
		} else {
			z.F.DecMapUint64UintptrX(x.FptrMapUint64Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Int = nil
	} else {
		yyv1821 := &x.FMapUint64Int
		yym1822 := z.DecBinary()
		_ = yym1822
		if false {
		} else {
			z.F.DecMapUint64IntX(yyv1821, d)
		}
	}
	if x.FptrMapUint64Int == nil {
		x.FptrMapUint64Int = new(map[uint64]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Int != nil {
			x.FptrMapUint64Int = nil
		}
	} else {
		if x.FptrMapUint64Int == nil {
			x.FptrMapUint64Int = new(map[uint64]int)
		}
		yym1824 := z.DecBinary()
		_ = yym1824
		if false {
		} else {
			z.F.DecMapUint64IntX(x.FptrMapUint64Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Int8 = nil
	} else {
		yyv1825 := &x.FMapUint64Int8
		yym1826 := z.DecBinary()
		_ = yym1826
		if false {
		} else {
			z.F.DecMapUint64Int8X(yyv1825, d)
		}
	}
	if x.FptrMapUint64Int8 == nil {
		x.FptrMapUint64Int8 = new(map[uint64]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Int8 != nil {
			x.FptrMapUint64Int8 = nil
		}
	} else {
		if x.FptrMapUint64Int8 == nil {
			x.FptrMapUint64Int8 = new(map[uint64]int8)
		}
		yym1828 := z.DecBinary()
		_ = yym1828
		if false {
		} else {
			z.F.DecMapUint64Int8X(x.FptrMapUint64Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Int16 = nil
	} else {
		yyv1829 := &x.FMapUint64Int16
		yym1830 := z.DecBinary()
		_ = yym1830
		if false {
		} else {
			z.F.DecMapUint64Int16X(yyv1829, d)
		}
	}
	if x.FptrMapUint64Int16 == nil {
		x.FptrMapUint64Int16 = new(map[uint64]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Int16 != nil {
			x.FptrMapUint64Int16 = nil
		}
	} else {
		if x.FptrMapUint64Int16 == nil {
			x.FptrMapUint64Int16 = new(map[uint64]int16)
		}
		yym1832 := z.DecBinary()
		_ = yym1832
		if false {
		} else {
			z.F.DecMapUint64Int16X(x.FptrMapUint64Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Int32 = nil
	} else {
		yyv1833 := &x.FMapUint64Int32
		yym1834 := z.DecBinary()
		_ = yym1834
		if false {
		} else {
			z.F.DecMapUint64Int32X(yyv1833, d)
		}
	}
	if x.FptrMapUint64Int32 == nil {
		x.FptrMapUint64Int32 = new(map[uint64]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Int32 != nil {
			x.FptrMapUint64Int32 = nil
		}
	} else {
		if x.FptrMapUint64Int32 == nil {
			x.FptrMapUint64Int32 = new(map[uint64]int32)
		}
		yym1836 := z.DecBinary()
		_ = yym1836
		if false {
		} else {
			z.F.DecMapUint64Int32X(x.FptrMapUint64Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Int64 = nil
	} else {
		yyv1837 := &x.FMapUint64Int64
		yym1838 := z.DecBinary()
		_ = yym1838
		if false {
		} else {
			z.F.DecMapUint64Int64X(yyv1837, d)
		}
	}
	if x.FptrMapUint64Int64 == nil {
		x.FptrMapUint64Int64 = new(map[uint64]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Int64 != nil {
			x.FptrMapUint64Int64 = nil
		}
	} else {
		if x.FptrMapUint64Int64 == nil {
			x.FptrMapUint64Int64 = new(map[uint64]int64)
		}
		yym1840 := z.DecBinary()
		_ = yym1840
		if false {
		} else {
			z.F.DecMapUint64Int64X(x.FptrMapUint64Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Float32 = nil
	} else {
		yyv1841 := &x.FMapUint64Float32
		yym1842 := z.DecBinary()
		_ = yym1842
		if false {
		} else {
			z.F.DecMapUint64Float32X(yyv1841, d)
		}
	}
	if x.FptrMapUint64Float32 == nil {
		x.FptrMapUint64Float32 = new(map[uint64]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Float32 != nil {
			x.FptrMapUint64Float32 = nil
		}
	} else {
		if x.FptrMapUint64Float32 == nil {
			x.FptrMapUint64Float32 = new(map[uint64]float32)
		}
		yym1844 := z.DecBinary()
		_ = yym1844
		if false {
		} else {
			z.F.DecMapUint64Float32X(x.FptrMapUint64Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Float64 = nil
	} else {
		yyv1845 := &x.FMapUint64Float64
		yym1846 := z.DecBinary()
		_ = yym1846
		if false {
		} else {
			z.F.DecMapUint64Float64X(yyv1845, d)
		}
	}
	if x.FptrMapUint64Float64 == nil {
		x.FptrMapUint64Float64 = new(map[uint64]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Float64 != nil {
			x.FptrMapUint64Float64 = nil
		}
	} else {
		if x.FptrMapUint64Float64 == nil {
			x.FptrMapUint64Float64 = new(map[uint64]float64)
		}
		yym1848 := z.DecBinary()
		_ = yym1848
		if false {
		} else {
			z.F.DecMapUint64Float64X(x.FptrMapUint64Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUint64Bool = nil
	} else {
		yyv1849 := &x.FMapUint64Bool
		yym1850 := z.DecBinary()
		_ = yym1850
		if false {
		} else {
			z.F.DecMapUint64BoolX(yyv1849, d)
		}
	}
	if x.FptrMapUint64Bool == nil {
		x.FptrMapUint64Bool = new(map[uint64]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUint64Bool != nil {
			x.FptrMapUint64Bool = nil
		}
	} else {
		if x.FptrMapUint64Bool == nil {
			x.FptrMapUint64Bool = new(map[uint64]bool)
		}
		yym1852 := z.DecBinary()
		_ = yym1852
		if false {
		} else {
			z.F.DecMapUint64BoolX(x.FptrMapUint64Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrIntf = nil
	} else {
		yyv1853 := &x.FMapUintptrIntf
		yym1854 := z.DecBinary()
		_ = yym1854
		if false {
		} else {
			z.F.DecMapUintptrIntfX(yyv1853, d)
		}
	}
	if x.FptrMapUintptrIntf == nil {
		x.FptrMapUintptrIntf = new(map[uintptr]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrIntf != nil {
			x.FptrMapUintptrIntf = nil
		}
	} else {
		if x.FptrMapUintptrIntf == nil {
			x.FptrMapUintptrIntf = new(map[uintptr]interface{})
		}
		yym1856 := z.DecBinary()
		_ = yym1856
		if false {
		} else {
			z.F.DecMapUintptrIntfX(x.FptrMapUintptrIntf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrString = nil
	} else {
		yyv1857 := &x.FMapUintptrString
		yym1858 := z.DecBinary()
		_ = yym1858
		if false {
		} else {
			z.F.DecMapUintptrStringX(yyv1857, d)
		}
	}
	if x.FptrMapUintptrString == nil {
		x.FptrMapUintptrString = new(map[uintptr]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrString != nil {
			x.FptrMapUintptrString = nil
		}
	} else {
		if x.FptrMapUintptrString == nil {
			x.FptrMapUintptrString = new(map[uintptr]string)
		}
		yym1860 := z.DecBinary()
		_ = yym1860
		if false {
		} else {
			z.F.DecMapUintptrStringX(x.FptrMapUintptrString, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrUint = nil
	} else {
		yyv1861 := &x.FMapUintptrUint
		yym1862 := z.DecBinary()
		_ = yym1862
		if false {
		} else {
			z.F.DecMapUintptrUintX(yyv1861, d)
		}
	}
	if x.FptrMapUintptrUint == nil {
		x.FptrMapUintptrUint = new(map[uintptr]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrUint != nil {
			x.FptrMapUintptrUint = nil
		}
	} else {
		if x.FptrMapUintptrUint == nil {
			x.FptrMapUintptrUint = new(map[uintptr]uint)
		}
		yym1864 := z.DecBinary()
		_ = yym1864
		if false {
		} else {
			z.F.DecMapUintptrUintX(x.FptrMapUintptrUint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrUint8 = nil
	} else {
		yyv1865 := &x.FMapUintptrUint8
		yym1866 := z.DecBinary()
		_ = yym1866
		if false {
		} else {
			z.F.DecMapUintptrUint8X(yyv1865, d)
		}
	}
	if x.FptrMapUintptrUint8 == nil {
		x.FptrMapUintptrUint8 = new(map[uintptr]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrUint8 != nil {
			x.FptrMapUintptrUint8 = nil
		}
	} else {
		if x.FptrMapUintptrUint8 == nil {
			x.FptrMapUintptrUint8 = new(map[uintptr]uint8)
		}
		yym1868 := z.DecBinary()
		_ = yym1868
		if false {
		} else {
			z.F.DecMapUintptrUint8X(x.FptrMapUintptrUint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrUint16 = nil
	} else {
		yyv1869 := &x.FMapUintptrUint16
		yym1870 := z.DecBinary()
		_ = yym1870
		if false {
		} else {
			z.F.DecMapUintptrUint16X(yyv1869, d)
		}
	}
	if x.FptrMapUintptrUint16 == nil {
		x.FptrMapUintptrUint16 = new(map[uintptr]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrUint16 != nil {
			x.FptrMapUintptrUint16 = nil
		}
	} else {
		if x.FptrMapUintptrUint16 == nil {
			x.FptrMapUintptrUint16 = new(map[uintptr]uint16)
		}
		yym1872 := z.DecBinary()
		_ = yym1872
		if false {
		} else {
			z.F.DecMapUintptrUint16X(x.FptrMapUintptrUint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrUint32 = nil
	} else {
		yyv1873 := &x.FMapUintptrUint32
		yym1874 := z.DecBinary()
		_ = yym1874
		if false {
		} else {
			z.F.DecMapUintptrUint32X(yyv1873, d)
		}
	}
	if x.FptrMapUintptrUint32 == nil {
		x.FptrMapUintptrUint32 = new(map[uintptr]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrUint32 != nil {
			x.FptrMapUintptrUint32 = nil
		}
	} else {
		if x.FptrMapUintptrUint32 == nil {
			x.FptrMapUintptrUint32 = new(map[uintptr]uint32)
		}
		yym1876 := z.DecBinary()
		_ = yym1876
		if false {
		} else {
			z.F.DecMapUintptrUint32X(x.FptrMapUintptrUint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrUint64 = nil
	} else {
		yyv1877 := &x.FMapUintptrUint64
		yym1878 := z.DecBinary()
		_ = yym1878
		if false {
		} else {
			z.F.DecMapUintptrUint64X(yyv1877, d)
		}
	}
	if x.FptrMapUintptrUint64 == nil {
		x.FptrMapUintptrUint64 = new(map[uintptr]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrUint64 != nil {
			x.FptrMapUintptrUint64 = nil
		}
	} else {
		if x.FptrMapUintptrUint64 == nil {
			x.FptrMapUintptrUint64 = new(map[uintptr]uint64)
		}
		yym1880 := z.DecBinary()
		_ = yym1880
		if false {
		} else {
			z.F.DecMapUintptrUint64X(x.FptrMapUintptrUint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrUintptr = nil
	} else {
		yyv1881 := &x.FMapUintptrUintptr
		yym1882 := z.DecBinary()
		_ = yym1882
		if false {
		} else {
			z.F.DecMapUintptrUintptrX(yyv1881, d)
		}
	}
	if x.FptrMapUintptrUintptr == nil {
		x.FptrMapUintptrUintptr = new(map[uintptr]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrUintptr != nil {
			x.FptrMapUintptrUintptr = nil
		}
	} else {
		if x.FptrMapUintptrUintptr == nil {
			x.FptrMapUintptrUintptr = new(map[uintptr]uintptr)
		}
		yym1884 := z.DecBinary()
		_ = yym1884
		if false {
		} else {
			z.F.DecMapUintptrUintptrX(x.FptrMapUintptrUintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrInt = nil
	} else {
		yyv1885 := &x.FMapUintptrInt
		yym1886 := z.DecBinary()
		_ = yym1886
		if false {
		} else {
			z.F.DecMapUintptrIntX(yyv1885, d)
		}
	}
	if x.FptrMapUintptrInt == nil {
		x.FptrMapUintptrInt = new(map[uintptr]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrInt != nil {
			x.FptrMapUintptrInt = nil
		}
	} else {
		if x.FptrMapUintptrInt == nil {
			x.FptrMapUintptrInt = new(map[uintptr]int)
		}
		yym1888 := z.DecBinary()
		_ = yym1888
		if false {
		} else {
			z.F.DecMapUintptrIntX(x.FptrMapUintptrInt, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrInt8 = nil
	} else {
		yyv1889 := &x.FMapUintptrInt8
		yym1890 := z.DecBinary()
		_ = yym1890
		if false {
		} else {
			z.F.DecMapUintptrInt8X(yyv1889, d)
		}
	}
	if x.FptrMapUintptrInt8 == nil {
		x.FptrMapUintptrInt8 = new(map[uintptr]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrInt8 != nil {
			x.FptrMapUintptrInt8 = nil
		}
	} else {
		if x.FptrMapUintptrInt8 == nil {
			x.FptrMapUintptrInt8 = new(map[uintptr]int8)
		}
		yym1892 := z.DecBinary()
		_ = yym1892
		if false {
		} else {
			z.F.DecMapUintptrInt8X(x.FptrMapUintptrInt8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrInt16 = nil
	} else {
		yyv1893 := &x.FMapUintptrInt16
		yym1894 := z.DecBinary()
		_ = yym1894
		if false {
		} else {
			z.F.DecMapUintptrInt16X(yyv1893, d)
		}
	}
	if x.FptrMapUintptrInt16 == nil {
		x.FptrMapUintptrInt16 = new(map[uintptr]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrInt16 != nil {
			x.FptrMapUintptrInt16 = nil
		}
	} else {
		if x.FptrMapUintptrInt16 == nil {
			x.FptrMapUintptrInt16 = new(map[uintptr]int16)
		}
		yym1896 := z.DecBinary()
		_ = yym1896
		if false {
		} else {
			z.F.DecMapUintptrInt16X(x.FptrMapUintptrInt16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrInt32 = nil
	} else {
		yyv1897 := &x.FMapUintptrInt32
		yym1898 := z.DecBinary()
		_ = yym1898
		if false {
		} else {
			z.F.DecMapUintptrInt32X(yyv1897, d)
		}
	}
	if x.FptrMapUintptrInt32 == nil {
		x.FptrMapUintptrInt32 = new(map[uintptr]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrInt32 != nil {
			x.FptrMapUintptrInt32 = nil
		}
	} else {
		if x.FptrMapUintptrInt32 == nil {
			x.FptrMapUintptrInt32 = new(map[uintptr]int32)
		}
		yym1900 := z.DecBinary()
		_ = yym1900
		if false {
		} else {
			z.F.DecMapUintptrInt32X(x.FptrMapUintptrInt32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrInt64 = nil
	} else {
		yyv1901 := &x.FMapUintptrInt64
		yym1902 := z.DecBinary()
		_ = yym1902
		if false {
		} else {
			z.F.DecMapUintptrInt64X(yyv1901, d)
		}
	}
	if x.FptrMapUintptrInt64 == nil {
		x.FptrMapUintptrInt64 = new(map[uintptr]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrInt64 != nil {
			x.FptrMapUintptrInt64 = nil
		}
	} else {
		if x.FptrMapUintptrInt64 == nil {
			x.FptrMapUintptrInt64 = new(map[uintptr]int64)
		}
		yym1904 := z.DecBinary()
		_ = yym1904
		if false {
		} else {
			z.F.DecMapUintptrInt64X(x.FptrMapUintptrInt64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrFloat32 = nil
	} else {
		yyv1905 := &x.FMapUintptrFloat32
		yym1906 := z.DecBinary()
		_ = yym1906
		if false {
		} else {
			z.F.DecMapUintptrFloat32X(yyv1905, d)
		}
	}
	if x.FptrMapUintptrFloat32 == nil {
		x.FptrMapUintptrFloat32 = new(map[uintptr]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrFloat32 != nil {
			x.FptrMapUintptrFloat32 = nil
		}
	} else {
		if x.FptrMapUintptrFloat32 == nil {
			x.FptrMapUintptrFloat32 = new(map[uintptr]float32)
		}
		yym1908 := z.DecBinary()
		_ = yym1908
		if false {
		} else {
			z.F.DecMapUintptrFloat32X(x.FptrMapUintptrFloat32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrFloat64 = nil
	} else {
		yyv1909 := &x.FMapUintptrFloat64
		yym1910 := z.DecBinary()
		_ = yym1910
		if false {
		} else {
			z.F.DecMapUintptrFloat64X(yyv1909, d)
		}
	}
	if x.FptrMapUintptrFloat64 == nil {
		x.FptrMapUintptrFloat64 = new(map[uintptr]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrFloat64 != nil {
			x.FptrMapUintptrFloat64 = nil
		}
	} else {
		if x.FptrMapUintptrFloat64 == nil {
			x.FptrMapUintptrFloat64 = new(map[uintptr]float64)
		}
		yym1912 := z.DecBinary()
		_ = yym1912
		if false {
		} else {
			z.F.DecMapUintptrFloat64X(x.FptrMapUintptrFloat64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapUintptrBool = nil
	} else {
		yyv1913 := &x.FMapUintptrBool
		yym1914 := z.DecBinary()
		_ = yym1914
		if false {
		} else {
			z.F.DecMapUintptrBoolX(yyv1913, d)
		}
	}
	if x.FptrMapUintptrBool == nil {
		x.FptrMapUintptrBool = new(map[uintptr]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapUintptrBool != nil {
			x.FptrMapUintptrBool = nil
		}
	} else {
		if x.FptrMapUintptrBool == nil {
			x.FptrMapUintptrBool = new(map[uintptr]bool)
		}
		yym1916 := z.DecBinary()
		_ = yym1916
		if false {
		} else {
			z.F.DecMapUintptrBoolX(x.FptrMapUintptrBool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntIntf = nil
	} else {
		yyv1917 := &x.FMapIntIntf
		yym1918 := z.DecBinary()
		_ = yym1918
		if false {
		} else {
			z.F.DecMapIntIntfX(yyv1917, d)
		}
	}
	if x.FptrMapIntIntf == nil {
		x.FptrMapIntIntf = new(map[int]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntIntf != nil {
			x.FptrMapIntIntf = nil
		}
	} else {
		if x.FptrMapIntIntf == nil {
			x.FptrMapIntIntf = new(map[int]interface{})
		}
		yym1920 := z.DecBinary()
		_ = yym1920
		if false {
		} else {
			z.F.DecMapIntIntfX(x.FptrMapIntIntf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntString = nil
	} else {
		yyv1921 := &x.FMapIntString
		yym1922 := z.DecBinary()
		_ = yym1922
		if false {
		} else {
			z.F.DecMapIntStringX(yyv1921, d)
		}
	}
	if x.FptrMapIntString == nil {
		x.FptrMapIntString = new(map[int]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntString != nil {
			x.FptrMapIntString = nil
		}
	} else {
		if x.FptrMapIntString == nil {
			x.FptrMapIntString = new(map[int]string)
		}
		yym1924 := z.DecBinary()
		_ = yym1924
		if false {
		} else {
			z.F.DecMapIntStringX(x.FptrMapIntString, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntUint = nil
	} else {
		yyv1925 := &x.FMapIntUint
		yym1926 := z.DecBinary()
		_ = yym1926
		if false {
		} else {
			z.F.DecMapIntUintX(yyv1925, d)
		}
	}
	if x.FptrMapIntUint == nil {
		x.FptrMapIntUint = new(map[int]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntUint != nil {
			x.FptrMapIntUint = nil
		}
	} else {
		if x.FptrMapIntUint == nil {
			x.FptrMapIntUint = new(map[int]uint)
		}
		yym1928 := z.DecBinary()
		_ = yym1928
		if false {
		} else {
			z.F.DecMapIntUintX(x.FptrMapIntUint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntUint8 = nil
	} else {
		yyv1929 := &x.FMapIntUint8
		yym1930 := z.DecBinary()
		_ = yym1930
		if false {
		} else {
			z.F.DecMapIntUint8X(yyv1929, d)
		}
	}
	if x.FptrMapIntUint8 == nil {
		x.FptrMapIntUint8 = new(map[int]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntUint8 != nil {
			x.FptrMapIntUint8 = nil
		}
	} else {
		if x.FptrMapIntUint8 == nil {
			x.FptrMapIntUint8 = new(map[int]uint8)
		}
		yym1932 := z.DecBinary()
		_ = yym1932
		if false {
		} else {
			z.F.DecMapIntUint8X(x.FptrMapIntUint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntUint16 = nil
	} else {
		yyv1933 := &x.FMapIntUint16
		yym1934 := z.DecBinary()
		_ = yym1934
		if false {
		} else {
			z.F.DecMapIntUint16X(yyv1933, d)
		}
	}
	if x.FptrMapIntUint16 == nil {
		x.FptrMapIntUint16 = new(map[int]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntUint16 != nil {
			x.FptrMapIntUint16 = nil
		}
	} else {
		if x.FptrMapIntUint16 == nil {
			x.FptrMapIntUint16 = new(map[int]uint16)
		}
		yym1936 := z.DecBinary()
		_ = yym1936
		if false {
		} else {
			z.F.DecMapIntUint16X(x.FptrMapIntUint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntUint32 = nil
	} else {
		yyv1937 := &x.FMapIntUint32
		yym1938 := z.DecBinary()
		_ = yym1938
		if false {
		} else {
			z.F.DecMapIntUint32X(yyv1937, d)
		}
	}
	if x.FptrMapIntUint32 == nil {
		x.FptrMapIntUint32 = new(map[int]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntUint32 != nil {
			x.FptrMapIntUint32 = nil
		}
	} else {
		if x.FptrMapIntUint32 == nil {
			x.FptrMapIntUint32 = new(map[int]uint32)
		}
		yym1940 := z.DecBinary()
		_ = yym1940
		if false {
		} else {
			z.F.DecMapIntUint32X(x.FptrMapIntUint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntUint64 = nil
	} else {
		yyv1941 := &x.FMapIntUint64
		yym1942 := z.DecBinary()
		_ = yym1942
		if false {
		} else {
			z.F.DecMapIntUint64X(yyv1941, d)
		}
	}
	if x.FptrMapIntUint64 == nil {
		x.FptrMapIntUint64 = new(map[int]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntUint64 != nil {
			x.FptrMapIntUint64 = nil
		}
	} else {
		if x.FptrMapIntUint64 == nil {
			x.FptrMapIntUint64 = new(map[int]uint64)
		}
		yym1944 := z.DecBinary()
		_ = yym1944
		if false {
		} else {
			z.F.DecMapIntUint64X(x.FptrMapIntUint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntUintptr = nil
	} else {
		yyv1945 := &x.FMapIntUintptr
		yym1946 := z.DecBinary()
		_ = yym1946
		if false {
		} else {
			z.F.DecMapIntUintptrX(yyv1945, d)
		}
	}
	if x.FptrMapIntUintptr == nil {
		x.FptrMapIntUintptr = new(map[int]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntUintptr != nil {
			x.FptrMapIntUintptr = nil
		}
	} else {
		if x.FptrMapIntUintptr == nil {
			x.FptrMapIntUintptr = new(map[int]uintptr)
		}
		yym1948 := z.DecBinary()
		_ = yym1948
		if false {
		} else {
			z.F.DecMapIntUintptrX(x.FptrMapIntUintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntInt = nil
	} else {
		yyv1949 := &x.FMapIntInt
		yym1950 := z.DecBinary()
		_ = yym1950
		if false {
		} else {
			z.F.DecMapIntIntX(yyv1949, d)
		}
	}
	if x.FptrMapIntInt == nil {
		x.FptrMapIntInt = new(map[int]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntInt != nil {
			x.FptrMapIntInt = nil
		}
	} else {
		if x.FptrMapIntInt == nil {
			x.FptrMapIntInt = new(map[int]int)
		}
		yym1952 := z.DecBinary()
		_ = yym1952
		if false {
		} else {
			z.F.DecMapIntIntX(x.FptrMapIntInt, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntInt8 = nil
	} else {
		yyv1953 := &x.FMapIntInt8
		yym1954 := z.DecBinary()
		_ = yym1954
		if false {
		} else {
			z.F.DecMapIntInt8X(yyv1953, d)
		}
	}
	if x.FptrMapIntInt8 == nil {
		x.FptrMapIntInt8 = new(map[int]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntInt8 != nil {
			x.FptrMapIntInt8 = nil
		}
	} else {
		if x.FptrMapIntInt8 == nil {
			x.FptrMapIntInt8 = new(map[int]int8)
		}
		yym1956 := z.DecBinary()
		_ = yym1956
		if false {
		} else {
			z.F.DecMapIntInt8X(x.FptrMapIntInt8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntInt16 = nil
	} else {
		yyv1957 := &x.FMapIntInt16
		yym1958 := z.DecBinary()
		_ = yym1958
		if false {
		} else {
			z.F.DecMapIntInt16X(yyv1957, d)
		}
	}
	if x.FptrMapIntInt16 == nil {
		x.FptrMapIntInt16 = new(map[int]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntInt16 != nil {
			x.FptrMapIntInt16 = nil
		}
	} else {
		if x.FptrMapIntInt16 == nil {
			x.FptrMapIntInt16 = new(map[int]int16)
		}
		yym1960 := z.DecBinary()
		_ = yym1960
		if false {
		} else {
			z.F.DecMapIntInt16X(x.FptrMapIntInt16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntInt32 = nil
	} else {
		yyv1961 := &x.FMapIntInt32
		yym1962 := z.DecBinary()
		_ = yym1962
		if false {
		} else {
			z.F.DecMapIntInt32X(yyv1961, d)
		}
	}
	if x.FptrMapIntInt32 == nil {
		x.FptrMapIntInt32 = new(map[int]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntInt32 != nil {
			x.FptrMapIntInt32 = nil
		}
	} else {
		if x.FptrMapIntInt32 == nil {
			x.FptrMapIntInt32 = new(map[int]int32)
		}
		yym1964 := z.DecBinary()
		_ = yym1964
		if false {
		} else {
			z.F.DecMapIntInt32X(x.FptrMapIntInt32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntInt64 = nil
	} else {
		yyv1965 := &x.FMapIntInt64
		yym1966 := z.DecBinary()
		_ = yym1966
		if false {
		} else {
			z.F.DecMapIntInt64X(yyv1965, d)
		}
	}
	if x.FptrMapIntInt64 == nil {
		x.FptrMapIntInt64 = new(map[int]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntInt64 != nil {
			x.FptrMapIntInt64 = nil
		}
	} else {
		if x.FptrMapIntInt64 == nil {
			x.FptrMapIntInt64 = new(map[int]int64)
		}
		yym1968 := z.DecBinary()
		_ = yym1968
		if false {
		} else {
			z.F.DecMapIntInt64X(x.FptrMapIntInt64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntFloat32 = nil
	} else {
		yyv1969 := &x.FMapIntFloat32
		yym1970 := z.DecBinary()
		_ = yym1970
		if false {
		} else {
			z.F.DecMapIntFloat32X(yyv1969, d)
		}
	}
	if x.FptrMapIntFloat32 == nil {
		x.FptrMapIntFloat32 = new(map[int]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntFloat32 != nil {
			x.FptrMapIntFloat32 = nil
		}
	} else {
		if x.FptrMapIntFloat32 == nil {
			x.FptrMapIntFloat32 = new(map[int]float32)
		}
		yym1972 := z.DecBinary()
		_ = yym1972
		if false {
		} else {
			z.F.DecMapIntFloat32X(x.FptrMapIntFloat32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntFloat64 = nil
	} else {
		yyv1973 := &x.FMapIntFloat64
		yym1974 := z.DecBinary()
		_ = yym1974
		if false {
		} else {
			z.F.DecMapIntFloat64X(yyv1973, d)
		}
	}
	if x.FptrMapIntFloat64 == nil {
		x.FptrMapIntFloat64 = new(map[int]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntFloat64 != nil {
			x.FptrMapIntFloat64 = nil
		}
	} else {
		if x.FptrMapIntFloat64 == nil {
			x.FptrMapIntFloat64 = new(map[int]float64)
		}
		yym1976 := z.DecBinary()
		_ = yym1976
		if false {
		} else {
			z.F.DecMapIntFloat64X(x.FptrMapIntFloat64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapIntBool = nil
	} else {
		yyv1977 := &x.FMapIntBool
		yym1978 := z.DecBinary()
		_ = yym1978
		if false {
		} else {
			z.F.DecMapIntBoolX(yyv1977, d)
		}
	}
	if x.FptrMapIntBool == nil {
		x.FptrMapIntBool = new(map[int]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapIntBool != nil {
			x.FptrMapIntBool = nil
		}
	} else {
		if x.FptrMapIntBool == nil {
			x.FptrMapIntBool = new(map[int]bool)
		}
		yym1980 := z.DecBinary()
		_ = yym1980
		if false {
		} else {
			z.F.DecMapIntBoolX(x.FptrMapIntBool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Intf = nil
	} else {
		yyv1981 := &x.FMapInt8Intf
		yym1982 := z.DecBinary()
		_ = yym1982
		if false {
		} else {
			z.F.DecMapInt8IntfX(yyv1981, d)
		}
	}
	if x.FptrMapInt8Intf == nil {
		x.FptrMapInt8Intf = new(map[int8]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Intf != nil {
			x.FptrMapInt8Intf = nil
		}
	} else {
		if x.FptrMapInt8Intf == nil {
			x.FptrMapInt8Intf = new(map[int8]interface{})
		}
		yym1984 := z.DecBinary()
		_ = yym1984
		if false {
		} else {
			z.F.DecMapInt8IntfX(x.FptrMapInt8Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8String = nil
	} else {
		yyv1985 := &x.FMapInt8String
		yym1986 := z.DecBinary()
		_ = yym1986
		if false {
		} else {
			z.F.DecMapInt8StringX(yyv1985, d)
		}
	}
	if x.FptrMapInt8String == nil {
		x.FptrMapInt8String = new(map[int8]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8String != nil {
			x.FptrMapInt8String = nil
		}
	} else {
		if x.FptrMapInt8String == nil {
			x.FptrMapInt8String = new(map[int8]string)
		}
		yym1988 := z.DecBinary()
		_ = yym1988
		if false {
		} else {
			z.F.DecMapInt8StringX(x.FptrMapInt8String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Uint = nil
	} else {
		yyv1989 := &x.FMapInt8Uint
		yym1990 := z.DecBinary()
		_ = yym1990
		if false {
		} else {
			z.F.DecMapInt8UintX(yyv1989, d)
		}
	}
	if x.FptrMapInt8Uint == nil {
		x.FptrMapInt8Uint = new(map[int8]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Uint != nil {
			x.FptrMapInt8Uint = nil
		}
	} else {
		if x.FptrMapInt8Uint == nil {
			x.FptrMapInt8Uint = new(map[int8]uint)
		}
		yym1992 := z.DecBinary()
		_ = yym1992
		if false {
		} else {
			z.F.DecMapInt8UintX(x.FptrMapInt8Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Uint8 = nil
	} else {
		yyv1993 := &x.FMapInt8Uint8
		yym1994 := z.DecBinary()
		_ = yym1994
		if false {
		} else {
			z.F.DecMapInt8Uint8X(yyv1993, d)
		}
	}
	if x.FptrMapInt8Uint8 == nil {
		x.FptrMapInt8Uint8 = new(map[int8]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Uint8 != nil {
			x.FptrMapInt8Uint8 = nil
		}
	} else {
		if x.FptrMapInt8Uint8 == nil {
			x.FptrMapInt8Uint8 = new(map[int8]uint8)
		}
		yym1996 := z.DecBinary()
		_ = yym1996
		if false {
		} else {
			z.F.DecMapInt8Uint8X(x.FptrMapInt8Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Uint16 = nil
	} else {
		yyv1997 := &x.FMapInt8Uint16
		yym1998 := z.DecBinary()
		_ = yym1998
		if false {
		} else {
			z.F.DecMapInt8Uint16X(yyv1997, d)
		}
	}
	if x.FptrMapInt8Uint16 == nil {
		x.FptrMapInt8Uint16 = new(map[int8]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Uint16 != nil {
			x.FptrMapInt8Uint16 = nil
		}
	} else {
		if x.FptrMapInt8Uint16 == nil {
			x.FptrMapInt8Uint16 = new(map[int8]uint16)
		}
		yym2000 := z.DecBinary()
		_ = yym2000
		if false {
		} else {
			z.F.DecMapInt8Uint16X(x.FptrMapInt8Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Uint32 = nil
	} else {
		yyv2001 := &x.FMapInt8Uint32
		yym2002 := z.DecBinary()
		_ = yym2002
		if false {
		} else {
			z.F.DecMapInt8Uint32X(yyv2001, d)
		}
	}
	if x.FptrMapInt8Uint32 == nil {
		x.FptrMapInt8Uint32 = new(map[int8]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Uint32 != nil {
			x.FptrMapInt8Uint32 = nil
		}
	} else {
		if x.FptrMapInt8Uint32 == nil {
			x.FptrMapInt8Uint32 = new(map[int8]uint32)
		}
		yym2004 := z.DecBinary()
		_ = yym2004
		if false {
		} else {
			z.F.DecMapInt8Uint32X(x.FptrMapInt8Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Uint64 = nil
	} else {
		yyv2005 := &x.FMapInt8Uint64
		yym2006 := z.DecBinary()
		_ = yym2006
		if false {
		} else {
			z.F.DecMapInt8Uint64X(yyv2005, d)
		}
	}
	if x.FptrMapInt8Uint64 == nil {
		x.FptrMapInt8Uint64 = new(map[int8]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Uint64 != nil {
			x.FptrMapInt8Uint64 = nil
		}
	} else {
		if x.FptrMapInt8Uint64 == nil {
			x.FptrMapInt8Uint64 = new(map[int8]uint64)
		}
		yym2008 := z.DecBinary()
		_ = yym2008
		if false {
		} else {
			z.F.DecMapInt8Uint64X(x.FptrMapInt8Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Uintptr = nil
	} else {
		yyv2009 := &x.FMapInt8Uintptr
		yym2010 := z.DecBinary()
		_ = yym2010
		if false {
		} else {
			z.F.DecMapInt8UintptrX(yyv2009, d)
		}
	}
	if x.FptrMapInt8Uintptr == nil {
		x.FptrMapInt8Uintptr = new(map[int8]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Uintptr != nil {
			x.FptrMapInt8Uintptr = nil
		}
	} else {
		if x.FptrMapInt8Uintptr == nil {
			x.FptrMapInt8Uintptr = new(map[int8]uintptr)
		}
		yym2012 := z.DecBinary()
		_ = yym2012
		if false {
		} else {
			z.F.DecMapInt8UintptrX(x.FptrMapInt8Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Int = nil
	} else {
		yyv2013 := &x.FMapInt8Int
		yym2014 := z.DecBinary()
		_ = yym2014
		if false {
		} else {
			z.F.DecMapInt8IntX(yyv2013, d)
		}
	}
	if x.FptrMapInt8Int == nil {
		x.FptrMapInt8Int = new(map[int8]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Int != nil {
			x.FptrMapInt8Int = nil
		}
	} else {
		if x.FptrMapInt8Int == nil {
			x.FptrMapInt8Int = new(map[int8]int)
		}
		yym2016 := z.DecBinary()
		_ = yym2016
		if false {
		} else {
			z.F.DecMapInt8IntX(x.FptrMapInt8Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Int8 = nil
	} else {
		yyv2017 := &x.FMapInt8Int8
		yym2018 := z.DecBinary()
		_ = yym2018
		if false {
		} else {
			z.F.DecMapInt8Int8X(yyv2017, d)
		}
	}
	if x.FptrMapInt8Int8 == nil {
		x.FptrMapInt8Int8 = new(map[int8]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Int8 != nil {
			x.FptrMapInt8Int8 = nil
		}
	} else {
		if x.FptrMapInt8Int8 == nil {
			x.FptrMapInt8Int8 = new(map[int8]int8)
		}
		yym2020 := z.DecBinary()
		_ = yym2020
		if false {
		} else {
			z.F.DecMapInt8Int8X(x.FptrMapInt8Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Int16 = nil
	} else {
		yyv2021 := &x.FMapInt8Int16
		yym2022 := z.DecBinary()
		_ = yym2022
		if false {
		} else {
			z.F.DecMapInt8Int16X(yyv2021, d)
		}
	}
	if x.FptrMapInt8Int16 == nil {
		x.FptrMapInt8Int16 = new(map[int8]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Int16 != nil {
			x.FptrMapInt8Int16 = nil
		}
	} else {
		if x.FptrMapInt8Int16 == nil {
			x.FptrMapInt8Int16 = new(map[int8]int16)
		}
		yym2024 := z.DecBinary()
		_ = yym2024
		if false {
		} else {
			z.F.DecMapInt8Int16X(x.FptrMapInt8Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Int32 = nil
	} else {
		yyv2025 := &x.FMapInt8Int32
		yym2026 := z.DecBinary()
		_ = yym2026
		if false {
		} else {
			z.F.DecMapInt8Int32X(yyv2025, d)
		}
	}
	if x.FptrMapInt8Int32 == nil {
		x.FptrMapInt8Int32 = new(map[int8]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Int32 != nil {
			x.FptrMapInt8Int32 = nil
		}
	} else {
		if x.FptrMapInt8Int32 == nil {
			x.FptrMapInt8Int32 = new(map[int8]int32)
		}
		yym2028 := z.DecBinary()
		_ = yym2028
		if false {
		} else {
			z.F.DecMapInt8Int32X(x.FptrMapInt8Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Int64 = nil
	} else {
		yyv2029 := &x.FMapInt8Int64
		yym2030 := z.DecBinary()
		_ = yym2030
		if false {
		} else {
			z.F.DecMapInt8Int64X(yyv2029, d)
		}
	}
	if x.FptrMapInt8Int64 == nil {
		x.FptrMapInt8Int64 = new(map[int8]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Int64 != nil {
			x.FptrMapInt8Int64 = nil
		}
	} else {
		if x.FptrMapInt8Int64 == nil {
			x.FptrMapInt8Int64 = new(map[int8]int64)
		}
		yym2032 := z.DecBinary()
		_ = yym2032
		if false {
		} else {
			z.F.DecMapInt8Int64X(x.FptrMapInt8Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Float32 = nil
	} else {
		yyv2033 := &x.FMapInt8Float32
		yym2034 := z.DecBinary()
		_ = yym2034
		if false {
		} else {
			z.F.DecMapInt8Float32X(yyv2033, d)
		}
	}
	if x.FptrMapInt8Float32 == nil {
		x.FptrMapInt8Float32 = new(map[int8]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Float32 != nil {
			x.FptrMapInt8Float32 = nil
		}
	} else {
		if x.FptrMapInt8Float32 == nil {
			x.FptrMapInt8Float32 = new(map[int8]float32)
		}
		yym2036 := z.DecBinary()
		_ = yym2036
		if false {
		} else {
			z.F.DecMapInt8Float32X(x.FptrMapInt8Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Float64 = nil
	} else {
		yyv2037 := &x.FMapInt8Float64
		yym2038 := z.DecBinary()
		_ = yym2038
		if false {
		} else {
			z.F.DecMapInt8Float64X(yyv2037, d)
		}
	}
	if x.FptrMapInt8Float64 == nil {
		x.FptrMapInt8Float64 = new(map[int8]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Float64 != nil {
			x.FptrMapInt8Float64 = nil
		}
	} else {
		if x.FptrMapInt8Float64 == nil {
			x.FptrMapInt8Float64 = new(map[int8]float64)
		}
		yym2040 := z.DecBinary()
		_ = yym2040
		if false {
		} else {
			z.F.DecMapInt8Float64X(x.FptrMapInt8Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt8Bool = nil
	} else {
		yyv2041 := &x.FMapInt8Bool
		yym2042 := z.DecBinary()
		_ = yym2042
		if false {
		} else {
			z.F.DecMapInt8BoolX(yyv2041, d)
		}
	}
	if x.FptrMapInt8Bool == nil {
		x.FptrMapInt8Bool = new(map[int8]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt8Bool != nil {
			x.FptrMapInt8Bool = nil
		}
	} else {
		if x.FptrMapInt8Bool == nil {
			x.FptrMapInt8Bool = new(map[int8]bool)
		}
		yym2044 := z.DecBinary()
		_ = yym2044
		if false {
		} else {
			z.F.DecMapInt8BoolX(x.FptrMapInt8Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Intf = nil
	} else {
		yyv2045 := &x.FMapInt16Intf
		yym2046 := z.DecBinary()
		_ = yym2046
		if false {
		} else {
			z.F.DecMapInt16IntfX(yyv2045, d)
		}
	}
	if x.FptrMapInt16Intf == nil {
		x.FptrMapInt16Intf = new(map[int16]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Intf != nil {
			x.FptrMapInt16Intf = nil
		}
	} else {
		if x.FptrMapInt16Intf == nil {
			x.FptrMapInt16Intf = new(map[int16]interface{})
		}
		yym2048 := z.DecBinary()
		_ = yym2048
		if false {
		} else {
			z.F.DecMapInt16IntfX(x.FptrMapInt16Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16String = nil
	} else {
		yyv2049 := &x.FMapInt16String
		yym2050 := z.DecBinary()
		_ = yym2050
		if false {
		} else {
			z.F.DecMapInt16StringX(yyv2049, d)
		}
	}
	if x.FptrMapInt16String == nil {
		x.FptrMapInt16String = new(map[int16]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16String != nil {
			x.FptrMapInt16String = nil
		}
	} else {
		if x.FptrMapInt16String == nil {
			x.FptrMapInt16String = new(map[int16]string)
		}
		yym2052 := z.DecBinary()
		_ = yym2052
		if false {
		} else {
			z.F.DecMapInt16StringX(x.FptrMapInt16String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Uint = nil
	} else {
		yyv2053 := &x.FMapInt16Uint
		yym2054 := z.DecBinary()
		_ = yym2054
		if false {
		} else {
			z.F.DecMapInt16UintX(yyv2053, d)
		}
	}
	if x.FptrMapInt16Uint == nil {
		x.FptrMapInt16Uint = new(map[int16]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Uint != nil {
			x.FptrMapInt16Uint = nil
		}
	} else {
		if x.FptrMapInt16Uint == nil {
			x.FptrMapInt16Uint = new(map[int16]uint)
		}
		yym2056 := z.DecBinary()
		_ = yym2056
		if false {
		} else {
			z.F.DecMapInt16UintX(x.FptrMapInt16Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Uint8 = nil
	} else {
		yyv2057 := &x.FMapInt16Uint8
		yym2058 := z.DecBinary()
		_ = yym2058
		if false {
		} else {
			z.F.DecMapInt16Uint8X(yyv2057, d)
		}
	}
	if x.FptrMapInt16Uint8 == nil {
		x.FptrMapInt16Uint8 = new(map[int16]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Uint8 != nil {
			x.FptrMapInt16Uint8 = nil
		}
	} else {
		if x.FptrMapInt16Uint8 == nil {
			x.FptrMapInt16Uint8 = new(map[int16]uint8)
		}
		yym2060 := z.DecBinary()
		_ = yym2060
		if false {
		} else {
			z.F.DecMapInt16Uint8X(x.FptrMapInt16Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Uint16 = nil
	} else {
		yyv2061 := &x.FMapInt16Uint16
		yym2062 := z.DecBinary()
		_ = yym2062
		if false {
		} else {
			z.F.DecMapInt16Uint16X(yyv2061, d)
		}
	}
	if x.FptrMapInt16Uint16 == nil {
		x.FptrMapInt16Uint16 = new(map[int16]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Uint16 != nil {
			x.FptrMapInt16Uint16 = nil
		}
	} else {
		if x.FptrMapInt16Uint16 == nil {
			x.FptrMapInt16Uint16 = new(map[int16]uint16)
		}
		yym2064 := z.DecBinary()
		_ = yym2064
		if false {
		} else {
			z.F.DecMapInt16Uint16X(x.FptrMapInt16Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Uint32 = nil
	} else {
		yyv2065 := &x.FMapInt16Uint32
		yym2066 := z.DecBinary()
		_ = yym2066
		if false {
		} else {
			z.F.DecMapInt16Uint32X(yyv2065, d)
		}
	}
	if x.FptrMapInt16Uint32 == nil {
		x.FptrMapInt16Uint32 = new(map[int16]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Uint32 != nil {
			x.FptrMapInt16Uint32 = nil
		}
	} else {
		if x.FptrMapInt16Uint32 == nil {
			x.FptrMapInt16Uint32 = new(map[int16]uint32)
		}
		yym2068 := z.DecBinary()
		_ = yym2068
		if false {
		} else {
			z.F.DecMapInt16Uint32X(x.FptrMapInt16Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Uint64 = nil
	} else {
		yyv2069 := &x.FMapInt16Uint64
		yym2070 := z.DecBinary()
		_ = yym2070
		if false {
		} else {
			z.F.DecMapInt16Uint64X(yyv2069, d)
		}
	}
	if x.FptrMapInt16Uint64 == nil {
		x.FptrMapInt16Uint64 = new(map[int16]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Uint64 != nil {
			x.FptrMapInt16Uint64 = nil
		}
	} else {
		if x.FptrMapInt16Uint64 == nil {
			x.FptrMapInt16Uint64 = new(map[int16]uint64)
		}
		yym2072 := z.DecBinary()
		_ = yym2072
		if false {
		} else {
			z.F.DecMapInt16Uint64X(x.FptrMapInt16Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Uintptr = nil
	} else {
		yyv2073 := &x.FMapInt16Uintptr
		yym2074 := z.DecBinary()
		_ = yym2074
		if false {
		} else {
			z.F.DecMapInt16UintptrX(yyv2073, d)
		}
	}
	if x.FptrMapInt16Uintptr == nil {
		x.FptrMapInt16Uintptr = new(map[int16]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Uintptr != nil {
			x.FptrMapInt16Uintptr = nil
		}
	} else {
		if x.FptrMapInt16Uintptr == nil {
			x.FptrMapInt16Uintptr = new(map[int16]uintptr)
		}
		yym2076 := z.DecBinary()
		_ = yym2076
		if false {
		} else {
			z.F.DecMapInt16UintptrX(x.FptrMapInt16Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Int = nil
	} else {
		yyv2077 := &x.FMapInt16Int
		yym2078 := z.DecBinary()
		_ = yym2078
		if false {
		} else {
			z.F.DecMapInt16IntX(yyv2077, d)
		}
	}
	if x.FptrMapInt16Int == nil {
		x.FptrMapInt16Int = new(map[int16]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Int != nil {
			x.FptrMapInt16Int = nil
		}
	} else {
		if x.FptrMapInt16Int == nil {
			x.FptrMapInt16Int = new(map[int16]int)
		}
		yym2080 := z.DecBinary()
		_ = yym2080
		if false {
		} else {
			z.F.DecMapInt16IntX(x.FptrMapInt16Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Int8 = nil
	} else {
		yyv2081 := &x.FMapInt16Int8
		yym2082 := z.DecBinary()
		_ = yym2082
		if false {
		} else {
			z.F.DecMapInt16Int8X(yyv2081, d)
		}
	}
	if x.FptrMapInt16Int8 == nil {
		x.FptrMapInt16Int8 = new(map[int16]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Int8 != nil {
			x.FptrMapInt16Int8 = nil
		}
	} else {
		if x.FptrMapInt16Int8 == nil {
			x.FptrMapInt16Int8 = new(map[int16]int8)
		}
		yym2084 := z.DecBinary()
		_ = yym2084
		if false {
		} else {
			z.F.DecMapInt16Int8X(x.FptrMapInt16Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Int16 = nil
	} else {
		yyv2085 := &x.FMapInt16Int16
		yym2086 := z.DecBinary()
		_ = yym2086
		if false {
		} else {
			z.F.DecMapInt16Int16X(yyv2085, d)
		}
	}
	if x.FptrMapInt16Int16 == nil {
		x.FptrMapInt16Int16 = new(map[int16]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Int16 != nil {
			x.FptrMapInt16Int16 = nil
		}
	} else {
		if x.FptrMapInt16Int16 == nil {
			x.FptrMapInt16Int16 = new(map[int16]int16)
		}
		yym2088 := z.DecBinary()
		_ = yym2088
		if false {
		} else {
			z.F.DecMapInt16Int16X(x.FptrMapInt16Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Int32 = nil
	} else {
		yyv2089 := &x.FMapInt16Int32
		yym2090 := z.DecBinary()
		_ = yym2090
		if false {
		} else {
			z.F.DecMapInt16Int32X(yyv2089, d)
		}
	}
	if x.FptrMapInt16Int32 == nil {
		x.FptrMapInt16Int32 = new(map[int16]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Int32 != nil {
			x.FptrMapInt16Int32 = nil
		}
	} else {
		if x.FptrMapInt16Int32 == nil {
			x.FptrMapInt16Int32 = new(map[int16]int32)
		}
		yym2092 := z.DecBinary()
		_ = yym2092
		if false {
		} else {
			z.F.DecMapInt16Int32X(x.FptrMapInt16Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Int64 = nil
	} else {
		yyv2093 := &x.FMapInt16Int64
		yym2094 := z.DecBinary()
		_ = yym2094
		if false {
		} else {
			z.F.DecMapInt16Int64X(yyv2093, d)
		}
	}
	if x.FptrMapInt16Int64 == nil {
		x.FptrMapInt16Int64 = new(map[int16]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Int64 != nil {
			x.FptrMapInt16Int64 = nil
		}
	} else {
		if x.FptrMapInt16Int64 == nil {
			x.FptrMapInt16Int64 = new(map[int16]int64)
		}
		yym2096 := z.DecBinary()
		_ = yym2096
		if false {
		} else {
			z.F.DecMapInt16Int64X(x.FptrMapInt16Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Float32 = nil
	} else {
		yyv2097 := &x.FMapInt16Float32
		yym2098 := z.DecBinary()
		_ = yym2098
		if false {
		} else {
			z.F.DecMapInt16Float32X(yyv2097, d)
		}
	}
	if x.FptrMapInt16Float32 == nil {
		x.FptrMapInt16Float32 = new(map[int16]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Float32 != nil {
			x.FptrMapInt16Float32 = nil
		}
	} else {
		if x.FptrMapInt16Float32 == nil {
			x.FptrMapInt16Float32 = new(map[int16]float32)
		}
		yym2100 := z.DecBinary()
		_ = yym2100
		if false {
		} else {
			z.F.DecMapInt16Float32X(x.FptrMapInt16Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Float64 = nil
	} else {
		yyv2101 := &x.FMapInt16Float64
		yym2102 := z.DecBinary()
		_ = yym2102
		if false {
		} else {
			z.F.DecMapInt16Float64X(yyv2101, d)
		}
	}
	if x.FptrMapInt16Float64 == nil {
		x.FptrMapInt16Float64 = new(map[int16]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Float64 != nil {
			x.FptrMapInt16Float64 = nil
		}
	} else {
		if x.FptrMapInt16Float64 == nil {
			x.FptrMapInt16Float64 = new(map[int16]float64)
		}
		yym2104 := z.DecBinary()
		_ = yym2104
		if false {
		} else {
			z.F.DecMapInt16Float64X(x.FptrMapInt16Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt16Bool = nil
	} else {
		yyv2105 := &x.FMapInt16Bool
		yym2106 := z.DecBinary()
		_ = yym2106
		if false {
		} else {
			z.F.DecMapInt16BoolX(yyv2105, d)
		}
	}
	if x.FptrMapInt16Bool == nil {
		x.FptrMapInt16Bool = new(map[int16]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt16Bool != nil {
			x.FptrMapInt16Bool = nil
		}
	} else {
		if x.FptrMapInt16Bool == nil {
			x.FptrMapInt16Bool = new(map[int16]bool)
		}
		yym2108 := z.DecBinary()
		_ = yym2108
		if false {
		} else {
			z.F.DecMapInt16BoolX(x.FptrMapInt16Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Intf = nil
	} else {
		yyv2109 := &x.FMapInt32Intf
		yym2110 := z.DecBinary()
		_ = yym2110
		if false {
		} else {
			z.F.DecMapInt32IntfX(yyv2109, d)
		}
	}
	if x.FptrMapInt32Intf == nil {
		x.FptrMapInt32Intf = new(map[int32]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Intf != nil {
			x.FptrMapInt32Intf = nil
		}
	} else {
		if x.FptrMapInt32Intf == nil {
			x.FptrMapInt32Intf = new(map[int32]interface{})
		}
		yym2112 := z.DecBinary()
		_ = yym2112
		if false {
		} else {
			z.F.DecMapInt32IntfX(x.FptrMapInt32Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32String = nil
	} else {
		yyv2113 := &x.FMapInt32String
		yym2114 := z.DecBinary()
		_ = yym2114
		if false {
		} else {
			z.F.DecMapInt32StringX(yyv2113, d)
		}
	}
	if x.FptrMapInt32String == nil {
		x.FptrMapInt32String = new(map[int32]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32String != nil {
			x.FptrMapInt32String = nil
		}
	} else {
		if x.FptrMapInt32String == nil {
			x.FptrMapInt32String = new(map[int32]string)
		}
		yym2116 := z.DecBinary()
		_ = yym2116
		if false {
		} else {
			z.F.DecMapInt32StringX(x.FptrMapInt32String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Uint = nil
	} else {
		yyv2117 := &x.FMapInt32Uint
		yym2118 := z.DecBinary()
		_ = yym2118
		if false {
		} else {
			z.F.DecMapInt32UintX(yyv2117, d)
		}
	}
	if x.FptrMapInt32Uint == nil {
		x.FptrMapInt32Uint = new(map[int32]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Uint != nil {
			x.FptrMapInt32Uint = nil
		}
	} else {
		if x.FptrMapInt32Uint == nil {
			x.FptrMapInt32Uint = new(map[int32]uint)
		}
		yym2120 := z.DecBinary()
		_ = yym2120
		if false {
		} else {
			z.F.DecMapInt32UintX(x.FptrMapInt32Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Uint8 = nil
	} else {
		yyv2121 := &x.FMapInt32Uint8
		yym2122 := z.DecBinary()
		_ = yym2122
		if false {
		} else {
			z.F.DecMapInt32Uint8X(yyv2121, d)
		}
	}
	if x.FptrMapInt32Uint8 == nil {
		x.FptrMapInt32Uint8 = new(map[int32]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Uint8 != nil {
			x.FptrMapInt32Uint8 = nil
		}
	} else {
		if x.FptrMapInt32Uint8 == nil {
			x.FptrMapInt32Uint8 = new(map[int32]uint8)
		}
		yym2124 := z.DecBinary()
		_ = yym2124
		if false {
		} else {
			z.F.DecMapInt32Uint8X(x.FptrMapInt32Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Uint16 = nil
	} else {
		yyv2125 := &x.FMapInt32Uint16
		yym2126 := z.DecBinary()
		_ = yym2126
		if false {
		} else {
			z.F.DecMapInt32Uint16X(yyv2125, d)
		}
	}
	if x.FptrMapInt32Uint16 == nil {
		x.FptrMapInt32Uint16 = new(map[int32]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Uint16 != nil {
			x.FptrMapInt32Uint16 = nil
		}
	} else {
		if x.FptrMapInt32Uint16 == nil {
			x.FptrMapInt32Uint16 = new(map[int32]uint16)
		}
		yym2128 := z.DecBinary()
		_ = yym2128
		if false {
		} else {
			z.F.DecMapInt32Uint16X(x.FptrMapInt32Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Uint32 = nil
	} else {
		yyv2129 := &x.FMapInt32Uint32
		yym2130 := z.DecBinary()
		_ = yym2130
		if false {
		} else {
			z.F.DecMapInt32Uint32X(yyv2129, d)
		}
	}
	if x.FptrMapInt32Uint32 == nil {
		x.FptrMapInt32Uint32 = new(map[int32]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Uint32 != nil {
			x.FptrMapInt32Uint32 = nil
		}
	} else {
		if x.FptrMapInt32Uint32 == nil {
			x.FptrMapInt32Uint32 = new(map[int32]uint32)
		}
		yym2132 := z.DecBinary()
		_ = yym2132
		if false {
		} else {
			z.F.DecMapInt32Uint32X(x.FptrMapInt32Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Uint64 = nil
	} else {
		yyv2133 := &x.FMapInt32Uint64
		yym2134 := z.DecBinary()
		_ = yym2134
		if false {
		} else {
			z.F.DecMapInt32Uint64X(yyv2133, d)
		}
	}
	if x.FptrMapInt32Uint64 == nil {
		x.FptrMapInt32Uint64 = new(map[int32]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Uint64 != nil {
			x.FptrMapInt32Uint64 = nil
		}
	} else {
		if x.FptrMapInt32Uint64 == nil {
			x.FptrMapInt32Uint64 = new(map[int32]uint64)
		}
		yym2136 := z.DecBinary()
		_ = yym2136
		if false {
		} else {
			z.F.DecMapInt32Uint64X(x.FptrMapInt32Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Uintptr = nil
	} else {
		yyv2137 := &x.FMapInt32Uintptr
		yym2138 := z.DecBinary()
		_ = yym2138
		if false {
		} else {
			z.F.DecMapInt32UintptrX(yyv2137, d)
		}
	}
	if x.FptrMapInt32Uintptr == nil {
		x.FptrMapInt32Uintptr = new(map[int32]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Uintptr != nil {
			x.FptrMapInt32Uintptr = nil
		}
	} else {
		if x.FptrMapInt32Uintptr == nil {
			x.FptrMapInt32Uintptr = new(map[int32]uintptr)
		}
		yym2140 := z.DecBinary()
		_ = yym2140
		if false {
		} else {
			z.F.DecMapInt32UintptrX(x.FptrMapInt32Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Int = nil
	} else {
		yyv2141 := &x.FMapInt32Int
		yym2142 := z.DecBinary()
		_ = yym2142
		if false {
		} else {
			z.F.DecMapInt32IntX(yyv2141, d)
		}
	}
	if x.FptrMapInt32Int == nil {
		x.FptrMapInt32Int = new(map[int32]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Int != nil {
			x.FptrMapInt32Int = nil
		}
	} else {
		if x.FptrMapInt32Int == nil {
			x.FptrMapInt32Int = new(map[int32]int)
		}
		yym2144 := z.DecBinary()
		_ = yym2144
		if false {
		} else {
			z.F.DecMapInt32IntX(x.FptrMapInt32Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Int8 = nil
	} else {
		yyv2145 := &x.FMapInt32Int8
		yym2146 := z.DecBinary()
		_ = yym2146
		if false {
		} else {
			z.F.DecMapInt32Int8X(yyv2145, d)
		}
	}
	if x.FptrMapInt32Int8 == nil {
		x.FptrMapInt32Int8 = new(map[int32]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Int8 != nil {
			x.FptrMapInt32Int8 = nil
		}
	} else {
		if x.FptrMapInt32Int8 == nil {
			x.FptrMapInt32Int8 = new(map[int32]int8)
		}
		yym2148 := z.DecBinary()
		_ = yym2148
		if false {
		} else {
			z.F.DecMapInt32Int8X(x.FptrMapInt32Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Int16 = nil
	} else {
		yyv2149 := &x.FMapInt32Int16
		yym2150 := z.DecBinary()
		_ = yym2150
		if false {
		} else {
			z.F.DecMapInt32Int16X(yyv2149, d)
		}
	}
	if x.FptrMapInt32Int16 == nil {
		x.FptrMapInt32Int16 = new(map[int32]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Int16 != nil {
			x.FptrMapInt32Int16 = nil
		}
	} else {
		if x.FptrMapInt32Int16 == nil {
			x.FptrMapInt32Int16 = new(map[int32]int16)
		}
		yym2152 := z.DecBinary()
		_ = yym2152
		if false {
		} else {
			z.F.DecMapInt32Int16X(x.FptrMapInt32Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Int32 = nil
	} else {
		yyv2153 := &x.FMapInt32Int32
		yym2154 := z.DecBinary()
		_ = yym2154
		if false {
		} else {
			z.F.DecMapInt32Int32X(yyv2153, d)
		}
	}
	if x.FptrMapInt32Int32 == nil {
		x.FptrMapInt32Int32 = new(map[int32]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Int32 != nil {
			x.FptrMapInt32Int32 = nil
		}
	} else {
		if x.FptrMapInt32Int32 == nil {
			x.FptrMapInt32Int32 = new(map[int32]int32)
		}
		yym2156 := z.DecBinary()
		_ = yym2156
		if false {
		} else {
			z.F.DecMapInt32Int32X(x.FptrMapInt32Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Int64 = nil
	} else {
		yyv2157 := &x.FMapInt32Int64
		yym2158 := z.DecBinary()
		_ = yym2158
		if false {
		} else {
			z.F.DecMapInt32Int64X(yyv2157, d)
		}
	}
	if x.FptrMapInt32Int64 == nil {
		x.FptrMapInt32Int64 = new(map[int32]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Int64 != nil {
			x.FptrMapInt32Int64 = nil
		}
	} else {
		if x.FptrMapInt32Int64 == nil {
			x.FptrMapInt32Int64 = new(map[int32]int64)
		}
		yym2160 := z.DecBinary()
		_ = yym2160
		if false {
		} else {
			z.F.DecMapInt32Int64X(x.FptrMapInt32Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Float32 = nil
	} else {
		yyv2161 := &x.FMapInt32Float32
		yym2162 := z.DecBinary()
		_ = yym2162
		if false {
		} else {
			z.F.DecMapInt32Float32X(yyv2161, d)
		}
	}
	if x.FptrMapInt32Float32 == nil {
		x.FptrMapInt32Float32 = new(map[int32]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Float32 != nil {
			x.FptrMapInt32Float32 = nil
		}
	} else {
		if x.FptrMapInt32Float32 == nil {
			x.FptrMapInt32Float32 = new(map[int32]float32)
		}
		yym2164 := z.DecBinary()
		_ = yym2164
		if false {
		} else {
			z.F.DecMapInt32Float32X(x.FptrMapInt32Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Float64 = nil
	} else {
		yyv2165 := &x.FMapInt32Float64
		yym2166 := z.DecBinary()
		_ = yym2166
		if false {
		} else {
			z.F.DecMapInt32Float64X(yyv2165, d)
		}
	}
	if x.FptrMapInt32Float64 == nil {
		x.FptrMapInt32Float64 = new(map[int32]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Float64 != nil {
			x.FptrMapInt32Float64 = nil
		}
	} else {
		if x.FptrMapInt32Float64 == nil {
			x.FptrMapInt32Float64 = new(map[int32]float64)
		}
		yym2168 := z.DecBinary()
		_ = yym2168
		if false {
		} else {
			z.F.DecMapInt32Float64X(x.FptrMapInt32Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt32Bool = nil
	} else {
		yyv2169 := &x.FMapInt32Bool
		yym2170 := z.DecBinary()
		_ = yym2170
		if false {
		} else {
			z.F.DecMapInt32BoolX(yyv2169, d)
		}
	}
	if x.FptrMapInt32Bool == nil {
		x.FptrMapInt32Bool = new(map[int32]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt32Bool != nil {
			x.FptrMapInt32Bool = nil
		}
	} else {
		if x.FptrMapInt32Bool == nil {
			x.FptrMapInt32Bool = new(map[int32]bool)
		}
		yym2172 := z.DecBinary()
		_ = yym2172
		if false {
		} else {
			z.F.DecMapInt32BoolX(x.FptrMapInt32Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Intf = nil
	} else {
		yyv2173 := &x.FMapInt64Intf
		yym2174 := z.DecBinary()
		_ = yym2174
		if false {
		} else {
			z.F.DecMapInt64IntfX(yyv2173, d)
		}
	}
	if x.FptrMapInt64Intf == nil {
		x.FptrMapInt64Intf = new(map[int64]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Intf != nil {
			x.FptrMapInt64Intf = nil
		}
	} else {
		if x.FptrMapInt64Intf == nil {
			x.FptrMapInt64Intf = new(map[int64]interface{})
		}
		yym2176 := z.DecBinary()
		_ = yym2176
		if false {
		} else {
			z.F.DecMapInt64IntfX(x.FptrMapInt64Intf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64String = nil
	} else {
		yyv2177 := &x.FMapInt64String
		yym2178 := z.DecBinary()
		_ = yym2178
		if false {
		} else {
			z.F.DecMapInt64StringX(yyv2177, d)
		}
	}
	if x.FptrMapInt64String == nil {
		x.FptrMapInt64String = new(map[int64]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64String != nil {
			x.FptrMapInt64String = nil
		}
	} else {
		if x.FptrMapInt64String == nil {
			x.FptrMapInt64String = new(map[int64]string)
		}
		yym2180 := z.DecBinary()
		_ = yym2180
		if false {
		} else {
			z.F.DecMapInt64StringX(x.FptrMapInt64String, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Uint = nil
	} else {
		yyv2181 := &x.FMapInt64Uint
		yym2182 := z.DecBinary()
		_ = yym2182
		if false {
		} else {
			z.F.DecMapInt64UintX(yyv2181, d)
		}
	}
	if x.FptrMapInt64Uint == nil {
		x.FptrMapInt64Uint = new(map[int64]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Uint != nil {
			x.FptrMapInt64Uint = nil
		}
	} else {
		if x.FptrMapInt64Uint == nil {
			x.FptrMapInt64Uint = new(map[int64]uint)
		}
		yym2184 := z.DecBinary()
		_ = yym2184
		if false {
		} else {
			z.F.DecMapInt64UintX(x.FptrMapInt64Uint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Uint8 = nil
	} else {
		yyv2185 := &x.FMapInt64Uint8
		yym2186 := z.DecBinary()
		_ = yym2186
		if false {
		} else {
			z.F.DecMapInt64Uint8X(yyv2185, d)
		}
	}
	if x.FptrMapInt64Uint8 == nil {
		x.FptrMapInt64Uint8 = new(map[int64]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Uint8 != nil {
			x.FptrMapInt64Uint8 = nil
		}
	} else {
		if x.FptrMapInt64Uint8 == nil {
			x.FptrMapInt64Uint8 = new(map[int64]uint8)
		}
		yym2188 := z.DecBinary()
		_ = yym2188
		if false {
		} else {
			z.F.DecMapInt64Uint8X(x.FptrMapInt64Uint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Uint16 = nil
	} else {
		yyv2189 := &x.FMapInt64Uint16
		yym2190 := z.DecBinary()
		_ = yym2190
		if false {
		} else {
			z.F.DecMapInt64Uint16X(yyv2189, d)
		}
	}
	if x.FptrMapInt64Uint16 == nil {
		x.FptrMapInt64Uint16 = new(map[int64]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Uint16 != nil {
			x.FptrMapInt64Uint16 = nil
		}
	} else {
		if x.FptrMapInt64Uint16 == nil {
			x.FptrMapInt64Uint16 = new(map[int64]uint16)
		}
		yym2192 := z.DecBinary()
		_ = yym2192
		if false {
		} else {
			z.F.DecMapInt64Uint16X(x.FptrMapInt64Uint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Uint32 = nil
	} else {
		yyv2193 := &x.FMapInt64Uint32
		yym2194 := z.DecBinary()
		_ = yym2194
		if false {
		} else {
			z.F.DecMapInt64Uint32X(yyv2193, d)
		}
	}
	if x.FptrMapInt64Uint32 == nil {
		x.FptrMapInt64Uint32 = new(map[int64]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Uint32 != nil {
			x.FptrMapInt64Uint32 = nil
		}
	} else {
		if x.FptrMapInt64Uint32 == nil {
			x.FptrMapInt64Uint32 = new(map[int64]uint32)
		}
		yym2196 := z.DecBinary()
		_ = yym2196
		if false {
		} else {
			z.F.DecMapInt64Uint32X(x.FptrMapInt64Uint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Uint64 = nil
	} else {
		yyv2197 := &x.FMapInt64Uint64
		yym2198 := z.DecBinary()
		_ = yym2198
		if false {
		} else {
			z.F.DecMapInt64Uint64X(yyv2197, d)
		}
	}
	if x.FptrMapInt64Uint64 == nil {
		x.FptrMapInt64Uint64 = new(map[int64]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Uint64 != nil {
			x.FptrMapInt64Uint64 = nil
		}
	} else {
		if x.FptrMapInt64Uint64 == nil {
			x.FptrMapInt64Uint64 = new(map[int64]uint64)
		}
		yym2200 := z.DecBinary()
		_ = yym2200
		if false {
		} else {
			z.F.DecMapInt64Uint64X(x.FptrMapInt64Uint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Uintptr = nil
	} else {
		yyv2201 := &x.FMapInt64Uintptr
		yym2202 := z.DecBinary()
		_ = yym2202
		if false {
		} else {
			z.F.DecMapInt64UintptrX(yyv2201, d)
		}
	}
	if x.FptrMapInt64Uintptr == nil {
		x.FptrMapInt64Uintptr = new(map[int64]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Uintptr != nil {
			x.FptrMapInt64Uintptr = nil
		}
	} else {
		if x.FptrMapInt64Uintptr == nil {
			x.FptrMapInt64Uintptr = new(map[int64]uintptr)
		}
		yym2204 := z.DecBinary()
		_ = yym2204
		if false {
		} else {
			z.F.DecMapInt64UintptrX(x.FptrMapInt64Uintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Int = nil
	} else {
		yyv2205 := &x.FMapInt64Int
		yym2206 := z.DecBinary()
		_ = yym2206
		if false {
		} else {
			z.F.DecMapInt64IntX(yyv2205, d)
		}
	}
	if x.FptrMapInt64Int == nil {
		x.FptrMapInt64Int = new(map[int64]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Int != nil {
			x.FptrMapInt64Int = nil
		}
	} else {
		if x.FptrMapInt64Int == nil {
			x.FptrMapInt64Int = new(map[int64]int)
		}
		yym2208 := z.DecBinary()
		_ = yym2208
		if false {
		} else {
			z.F.DecMapInt64IntX(x.FptrMapInt64Int, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Int8 = nil
	} else {
		yyv2209 := &x.FMapInt64Int8
		yym2210 := z.DecBinary()
		_ = yym2210
		if false {
		} else {
			z.F.DecMapInt64Int8X(yyv2209, d)
		}
	}
	if x.FptrMapInt64Int8 == nil {
		x.FptrMapInt64Int8 = new(map[int64]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Int8 != nil {
			x.FptrMapInt64Int8 = nil
		}
	} else {
		if x.FptrMapInt64Int8 == nil {
			x.FptrMapInt64Int8 = new(map[int64]int8)
		}
		yym2212 := z.DecBinary()
		_ = yym2212
		if false {
		} else {
			z.F.DecMapInt64Int8X(x.FptrMapInt64Int8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Int16 = nil
	} else {
		yyv2213 := &x.FMapInt64Int16
		yym2214 := z.DecBinary()
		_ = yym2214
		if false {
		} else {
			z.F.DecMapInt64Int16X(yyv2213, d)
		}
	}
	if x.FptrMapInt64Int16 == nil {
		x.FptrMapInt64Int16 = new(map[int64]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Int16 != nil {
			x.FptrMapInt64Int16 = nil
		}
	} else {
		if x.FptrMapInt64Int16 == nil {
			x.FptrMapInt64Int16 = new(map[int64]int16)
		}
		yym2216 := z.DecBinary()
		_ = yym2216
		if false {
		} else {
			z.F.DecMapInt64Int16X(x.FptrMapInt64Int16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Int32 = nil
	} else {
		yyv2217 := &x.FMapInt64Int32
		yym2218 := z.DecBinary()
		_ = yym2218
		if false {
		} else {
			z.F.DecMapInt64Int32X(yyv2217, d)
		}
	}
	if x.FptrMapInt64Int32 == nil {
		x.FptrMapInt64Int32 = new(map[int64]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Int32 != nil {
			x.FptrMapInt64Int32 = nil
		}
	} else {
		if x.FptrMapInt64Int32 == nil {
			x.FptrMapInt64Int32 = new(map[int64]int32)
		}
		yym2220 := z.DecBinary()
		_ = yym2220
		if false {
		} else {
			z.F.DecMapInt64Int32X(x.FptrMapInt64Int32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Int64 = nil
	} else {
		yyv2221 := &x.FMapInt64Int64
		yym2222 := z.DecBinary()
		_ = yym2222
		if false {
		} else {
			z.F.DecMapInt64Int64X(yyv2221, d)
		}
	}
	if x.FptrMapInt64Int64 == nil {
		x.FptrMapInt64Int64 = new(map[int64]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Int64 != nil {
			x.FptrMapInt64Int64 = nil
		}
	} else {
		if x.FptrMapInt64Int64 == nil {
			x.FptrMapInt64Int64 = new(map[int64]int64)
		}
		yym2224 := z.DecBinary()
		_ = yym2224
		if false {
		} else {
			z.F.DecMapInt64Int64X(x.FptrMapInt64Int64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Float32 = nil
	} else {
		yyv2225 := &x.FMapInt64Float32
		yym2226 := z.DecBinary()
		_ = yym2226
		if false {
		} else {
			z.F.DecMapInt64Float32X(yyv2225, d)
		}
	}
	if x.FptrMapInt64Float32 == nil {
		x.FptrMapInt64Float32 = new(map[int64]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Float32 != nil {
			x.FptrMapInt64Float32 = nil
		}
	} else {
		if x.FptrMapInt64Float32 == nil {
			x.FptrMapInt64Float32 = new(map[int64]float32)
		}
		yym2228 := z.DecBinary()
		_ = yym2228
		if false {
		} else {
			z.F.DecMapInt64Float32X(x.FptrMapInt64Float32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Float64 = nil
	} else {
		yyv2229 := &x.FMapInt64Float64
		yym2230 := z.DecBinary()
		_ = yym2230
		if false {
		} else {
			z.F.DecMapInt64Float64X(yyv2229, d)
		}
	}
	if x.FptrMapInt64Float64 == nil {
		x.FptrMapInt64Float64 = new(map[int64]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Float64 != nil {
			x.FptrMapInt64Float64 = nil
		}
	} else {
		if x.FptrMapInt64Float64 == nil {
			x.FptrMapInt64Float64 = new(map[int64]float64)
		}
		yym2232 := z.DecBinary()
		_ = yym2232
		if false {
		} else {
			z.F.DecMapInt64Float64X(x.FptrMapInt64Float64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapInt64Bool = nil
	} else {
		yyv2233 := &x.FMapInt64Bool
		yym2234 := z.DecBinary()
		_ = yym2234
		if false {
		} else {
			z.F.DecMapInt64BoolX(yyv2233, d)
		}
	}
	if x.FptrMapInt64Bool == nil {
		x.FptrMapInt64Bool = new(map[int64]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapInt64Bool != nil {
			x.FptrMapInt64Bool = nil
		}
	} else {
		if x.FptrMapInt64Bool == nil {
			x.FptrMapInt64Bool = new(map[int64]bool)
		}
		yym2236 := z.DecBinary()
		_ = yym2236
		if false {
		} else {
			z.F.DecMapInt64BoolX(x.FptrMapInt64Bool, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolIntf = nil
	} else {
		yyv2237 := &x.FMapBoolIntf
		yym2238 := z.DecBinary()
		_ = yym2238
		if false {
		} else {
			z.F.DecMapBoolIntfX(yyv2237, d)
		}
	}
	if x.FptrMapBoolIntf == nil {
		x.FptrMapBoolIntf = new(map[bool]interface{})
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolIntf != nil {
			x.FptrMapBoolIntf = nil
		}
	} else {
		if x.FptrMapBoolIntf == nil {
			x.FptrMapBoolIntf = new(map[bool]interface{})
		}
		yym2240 := z.DecBinary()
		_ = yym2240
		if false {
		} else {
			z.F.DecMapBoolIntfX(x.FptrMapBoolIntf, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolString = nil
	} else {
		yyv2241 := &x.FMapBoolString
		yym2242 := z.DecBinary()
		_ = yym2242
		if false {
		} else {
			z.F.DecMapBoolStringX(yyv2241, d)
		}
	}
	if x.FptrMapBoolString == nil {
		x.FptrMapBoolString = new(map[bool]string)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolString != nil {
			x.FptrMapBoolString = nil
		}
	} else {
		if x.FptrMapBoolString == nil {
			x.FptrMapBoolString = new(map[bool]string)
		}
		yym2244 := z.DecBinary()
		_ = yym2244
		if false {
		} else {
			z.F.DecMapBoolStringX(x.FptrMapBoolString, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolUint = nil
	} else {
		yyv2245 := &x.FMapBoolUint
		yym2246 := z.DecBinary()
		_ = yym2246
		if false {
		} else {
			z.F.DecMapBoolUintX(yyv2245, d)
		}
	}
	if x.FptrMapBoolUint == nil {
		x.FptrMapBoolUint = new(map[bool]uint)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolUint != nil {
			x.FptrMapBoolUint = nil
		}
	} else {
		if x.FptrMapBoolUint == nil {
			x.FptrMapBoolUint = new(map[bool]uint)
		}
		yym2248 := z.DecBinary()
		_ = yym2248
		if false {
		} else {
			z.F.DecMapBoolUintX(x.FptrMapBoolUint, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolUint8 = nil
	} else {
		yyv2249 := &x.FMapBoolUint8
		yym2250 := z.DecBinary()
		_ = yym2250
		if false {
		} else {
			z.F.DecMapBoolUint8X(yyv2249, d)
		}
	}
	if x.FptrMapBoolUint8 == nil {
		x.FptrMapBoolUint8 = new(map[bool]uint8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolUint8 != nil {
			x.FptrMapBoolUint8 = nil
		}
	} else {
		if x.FptrMapBoolUint8 == nil {
			x.FptrMapBoolUint8 = new(map[bool]uint8)
		}
		yym2252 := z.DecBinary()
		_ = yym2252
		if false {
		} else {
			z.F.DecMapBoolUint8X(x.FptrMapBoolUint8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolUint16 = nil
	} else {
		yyv2253 := &x.FMapBoolUint16
		yym2254 := z.DecBinary()
		_ = yym2254
		if false {
		} else {
			z.F.DecMapBoolUint16X(yyv2253, d)
		}
	}
	if x.FptrMapBoolUint16 == nil {
		x.FptrMapBoolUint16 = new(map[bool]uint16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolUint16 != nil {
			x.FptrMapBoolUint16 = nil
		}
	} else {
		if x.FptrMapBoolUint16 == nil {
			x.FptrMapBoolUint16 = new(map[bool]uint16)
		}
		yym2256 := z.DecBinary()
		_ = yym2256
		if false {
		} else {
			z.F.DecMapBoolUint16X(x.FptrMapBoolUint16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolUint32 = nil
	} else {
		yyv2257 := &x.FMapBoolUint32
		yym2258 := z.DecBinary()
		_ = yym2258
		if false {
		} else {
			z.F.DecMapBoolUint32X(yyv2257, d)
		}
	}
	if x.FptrMapBoolUint32 == nil {
		x.FptrMapBoolUint32 = new(map[bool]uint32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolUint32 != nil {
			x.FptrMapBoolUint32 = nil
		}
	} else {
		if x.FptrMapBoolUint32 == nil {
			x.FptrMapBoolUint32 = new(map[bool]uint32)
		}
		yym2260 := z.DecBinary()
		_ = yym2260
		if false {
		} else {
			z.F.DecMapBoolUint32X(x.FptrMapBoolUint32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolUint64 = nil
	} else {
		yyv2261 := &x.FMapBoolUint64
		yym2262 := z.DecBinary()
		_ = yym2262
		if false {
		} else {
			z.F.DecMapBoolUint64X(yyv2261, d)
		}
	}
	if x.FptrMapBoolUint64 == nil {
		x.FptrMapBoolUint64 = new(map[bool]uint64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolUint64 != nil {
			x.FptrMapBoolUint64 = nil
		}
	} else {
		if x.FptrMapBoolUint64 == nil {
			x.FptrMapBoolUint64 = new(map[bool]uint64)
		}
		yym2264 := z.DecBinary()
		_ = yym2264
		if false {
		} else {
			z.F.DecMapBoolUint64X(x.FptrMapBoolUint64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolUintptr = nil
	} else {
		yyv2265 := &x.FMapBoolUintptr
		yym2266 := z.DecBinary()
		_ = yym2266
		if false {
		} else {
			z.F.DecMapBoolUintptrX(yyv2265, d)
		}
	}
	if x.FptrMapBoolUintptr == nil {
		x.FptrMapBoolUintptr = new(map[bool]uintptr)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolUintptr != nil {
			x.FptrMapBoolUintptr = nil
		}
	} else {
		if x.FptrMapBoolUintptr == nil {
			x.FptrMapBoolUintptr = new(map[bool]uintptr)
		}
		yym2268 := z.DecBinary()
		_ = yym2268
		if false {
		} else {
			z.F.DecMapBoolUintptrX(x.FptrMapBoolUintptr, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolInt = nil
	} else {
		yyv2269 := &x.FMapBoolInt
		yym2270 := z.DecBinary()
		_ = yym2270
		if false {
		} else {
			z.F.DecMapBoolIntX(yyv2269, d)
		}
	}
	if x.FptrMapBoolInt == nil {
		x.FptrMapBoolInt = new(map[bool]int)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolInt != nil {
			x.FptrMapBoolInt = nil
		}
	} else {
		if x.FptrMapBoolInt == nil {
			x.FptrMapBoolInt = new(map[bool]int)
		}
		yym2272 := z.DecBinary()
		_ = yym2272
		if false {
		} else {
			z.F.DecMapBoolIntX(x.FptrMapBoolInt, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolInt8 = nil
	} else {
		yyv2273 := &x.FMapBoolInt8
		yym2274 := z.DecBinary()
		_ = yym2274
		if false {
		} else {
			z.F.DecMapBoolInt8X(yyv2273, d)
		}
	}
	if x.FptrMapBoolInt8 == nil {
		x.FptrMapBoolInt8 = new(map[bool]int8)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolInt8 != nil {
			x.FptrMapBoolInt8 = nil
		}
	} else {
		if x.FptrMapBoolInt8 == nil {
			x.FptrMapBoolInt8 = new(map[bool]int8)
		}
		yym2276 := z.DecBinary()
		_ = yym2276
		if false {
		} else {
			z.F.DecMapBoolInt8X(x.FptrMapBoolInt8, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolInt16 = nil
	} else {
		yyv2277 := &x.FMapBoolInt16
		yym2278 := z.DecBinary()
		_ = yym2278
		if false {
		} else {
			z.F.DecMapBoolInt16X(yyv2277, d)
		}
	}
	if x.FptrMapBoolInt16 == nil {
		x.FptrMapBoolInt16 = new(map[bool]int16)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolInt16 != nil {
			x.FptrMapBoolInt16 = nil
		}
	} else {
		if x.FptrMapBoolInt16 == nil {
			x.FptrMapBoolInt16 = new(map[bool]int16)
		}
		yym2280 := z.DecBinary()
		_ = yym2280
		if false {
		} else {
			z.F.DecMapBoolInt16X(x.FptrMapBoolInt16, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolInt32 = nil
	} else {
		yyv2281 := &x.FMapBoolInt32
		yym2282 := z.DecBinary()
		_ = yym2282
		if false {
		} else {
			z.F.DecMapBoolInt32X(yyv2281, d)
		}
	}
	if x.FptrMapBoolInt32 == nil {
		x.FptrMapBoolInt32 = new(map[bool]int32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolInt32 != nil {
			x.FptrMapBoolInt32 = nil
		}
	} else {
		if x.FptrMapBoolInt32 == nil {
			x.FptrMapBoolInt32 = new(map[bool]int32)
		}
		yym2284 := z.DecBinary()
		_ = yym2284
		if false {
		} else {
			z.F.DecMapBoolInt32X(x.FptrMapBoolInt32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolInt64 = nil
	} else {
		yyv2285 := &x.FMapBoolInt64
		yym2286 := z.DecBinary()
		_ = yym2286
		if false {
		} else {
			z.F.DecMapBoolInt64X(yyv2285, d)
		}
	}
	if x.FptrMapBoolInt64 == nil {
		x.FptrMapBoolInt64 = new(map[bool]int64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolInt64 != nil {
			x.FptrMapBoolInt64 = nil
		}
	} else {
		if x.FptrMapBoolInt64 == nil {
			x.FptrMapBoolInt64 = new(map[bool]int64)
		}
		yym2288 := z.DecBinary()
		_ = yym2288
		if false {
		} else {
			z.F.DecMapBoolInt64X(x.FptrMapBoolInt64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolFloat32 = nil
	} else {
		yyv2289 := &x.FMapBoolFloat32
		yym2290 := z.DecBinary()
		_ = yym2290
		if false {
		} else {
			z.F.DecMapBoolFloat32X(yyv2289, d)
		}
	}
	if x.FptrMapBoolFloat32 == nil {
		x.FptrMapBoolFloat32 = new(map[bool]float32)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolFloat32 != nil {
			x.FptrMapBoolFloat32 = nil
		}
	} else {
		if x.FptrMapBoolFloat32 == nil {
			x.FptrMapBoolFloat32 = new(map[bool]float32)
		}
		yym2292 := z.DecBinary()
		_ = yym2292
		if false {
		} else {
			z.F.DecMapBoolFloat32X(x.FptrMapBoolFloat32, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolFloat64 = nil
	} else {
		yyv2293 := &x.FMapBoolFloat64
		yym2294 := z.DecBinary()
		_ = yym2294
		if false {
		} else {
			z.F.DecMapBoolFloat64X(yyv2293, d)
		}
	}
	if x.FptrMapBoolFloat64 == nil {
		x.FptrMapBoolFloat64 = new(map[bool]float64)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolFloat64 != nil {
			x.FptrMapBoolFloat64 = nil
		}
	} else {
		if x.FptrMapBoolFloat64 == nil {
			x.FptrMapBoolFloat64 = new(map[bool]float64)
		}
		yym2296 := z.DecBinary()
		_ = yym2296
		if false {
		} else {
			z.F.DecMapBoolFloat64X(x.FptrMapBoolFloat64, d)
		}
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.FMapBoolBool = nil
	} else {
		yyv2297 := &x.FMapBoolBool
		yym2298 := z.DecBinary()
		_ = yym2298
		if false {
		} else {
			z.F.DecMapBoolBoolX(yyv2297, d)
		}
	}
	if x.FptrMapBoolBool == nil {
		x.FptrMapBoolBool = new(map[bool]bool)
	}
	yyj1152++
	if yyhl1152 {
		yyb1152 = yyj1152 > l
	} else {
		yyb1152 = r.CheckBreak()
	}
	if yyb1152 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.FptrMapBoolBool != nil {
			x.FptrMapBoolBool = nil
		}
	} else {
		if x.FptrMapBoolBool == nil {
			x.FptrMapBoolBool = new(map[bool]bool)
		}
		yym2300 := z.DecBinary()
		_ = yym2300
		if false {
		} else {
			z.F.DecMapBoolBoolX(x.FptrMapBoolBool, d)
		}
	}
	for {
		yyj1152++
		if yyhl1152 {
			yyb1152 = yyj1152 > l
		} else {
			yyb1152 = r.CheckBreak()
		}
		if yyb1152 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj1152-1, "")
	}
	r.ReadArrayEnd()
}

func (x testMammoth2Binary) CodecEncodeSelf(e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	yym1 := z.EncBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.EncExt(x) {
	} else if yym1 {
		z.EncBinaryMarshal(x)
	} else {
		r.EncodeUint(uint64(x))
	}
}

func (x *testMammoth2Binary) CodecDecodeSelf(d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else if yym1 {
		z.DecBinaryUnmarshal(x)
	} else {
		*((*uint64)(x)) = uint64(r.DecodeUint(64))
	}
}

func (x testMammoth2Text) CodecEncodeSelf(e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	yym1 := z.EncBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.EncExt(x) {
	} else if !yym1 {
		z.EncTextMarshal(x)
	} else {
		r.EncodeUint(uint64(x))
	}
}

func (x *testMammoth2Text) CodecDecodeSelf(d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else if !yym1 {
		z.DecTextUnmarshal(x)
	} else {
		*((*uint64)(x)) = uint64(r.DecodeUint(64))
	}
}

func (x testMammoth2Json) CodecEncodeSelf(e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	yym1 := z.EncBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.EncExt(x) {
	} else if !yym1 && z.IsJSONHandle() {
		z.EncJSONMarshal(x)
	} else {
		r.EncodeUint(uint64(x))
	}
}

func (x *testMammoth2Json) CodecDecodeSelf(d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else if !yym1 && z.IsJSONHandle() {
		z.DecJSONUnmarshal(x)
	} else {
		*((*uint64)(x)) = uint64(r.DecodeUint(64))
	}
}

func (x *testMammoth2Basic) CodecEncodeSelf(e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			h.enctestMammoth2Basic((*testMammoth2Basic)(x), e)
		}
	}
}

func (x *testMammoth2Basic) CodecDecodeSelf(d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		h.dectestMammoth2Basic((*testMammoth2Basic)(x), d)
	}
}

func (x *TestMammoth2Wrapper) CodecEncodeSelf(e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(8)
			} else {
				r.WriteMapStart(8)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yy4 := &x.V
				yy4.CodecEncodeSelf(e)
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("V"))
				r.WriteMapElemValue()
				yy6 := &x.V
				yy6.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				x.T.CodecEncodeSelf(e)
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("T"))
				r.WriteMapElemValue()
				x.T.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				x.B.CodecEncodeSelf(e)
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("B"))
				r.WriteMapElemValue()
				x.B.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				x.J.CodecEncodeSelf(e)
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("J"))
				r.WriteMapElemValue()
				x.J.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yy18 := &x.C
				yy18.CodecEncodeSelf(e)
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("C"))
				r.WriteMapElemValue()
				yy20 := &x.C
				yy20.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.M == nil {
					r.EncodeNil()
				} else {
					yym23 := z.EncBinary()
					_ = yym23
					if false {
					} else {
						h.encMaptestMammoth2BasicTestMammoth2((map[testMammoth2Basic]TestMammoth2)(x.M), e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("M"))
				r.WriteMapElemValue()
				if x.M == nil {
					r.EncodeNil()
				} else {
					yym24 := z.EncBinary()
					_ = yym24
					if false {
					} else {
						h.encMaptestMammoth2BasicTestMammoth2((map[testMammoth2Basic]TestMammoth2)(x.M), e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.L == nil {
					r.EncodeNil()
				} else {
					yym26 := z.EncBinary()
					_ = yym26
					if false {
					} else {
						h.encSliceTestMammoth2(([]TestMammoth2)(x.L), e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("L"))
				r.WriteMapElemValue()
				if x.L == nil {
					r.EncodeNil()
				} else {
					yym27 := z.EncBinary()
					_ = yym27
					if false {
					} else {
						h.encSliceTestMammoth2(([]TestMammoth2)(x.L), e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yy29 := &x.A
				yym30 := z.EncBinary()
				_ = yym30
				if false {
				} else {
					h.encArray4int64((*[4]int64)(yy29), e)
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF819781, string("A"))
				r.WriteMapElemValue()
				yy31 := &x.A
				yym32 := z.EncBinary()
				_ = yym32
				if false {
				} else {
					h.encArray4int64((*[4]int64)(yy31), e)
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *TestMammoth2Wrapper) CodecDecodeSelf(d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap19781 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray19781 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr19781)
		}
	}
}

func (x *TestMammoth2Wrapper) codecDecodeSelfFromMap(l int, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "V":
			if r.TryDecodeAsNil() {
				x.V = TestMammoth2{}
			} else {
				yyv4 := &x.V
				yyv4.CodecDecodeSelf(d)
			}
		case "T":
			if r.TryDecodeAsNil() {
				x.T = 0
			} else {
				yyv5 := &x.T
				yyv5.CodecDecodeSelf(d)
			}
		case "B":
			if r.TryDecodeAsNil() {
				x.B = 0
			} else {
				yyv6 := &x.B
				yyv6.CodecDecodeSelf(d)
			}
		case "J":
			if r.TryDecodeAsNil() {
				x.J = 0
			} else {
				yyv7 := &x.J
				yyv7.CodecDecodeSelf(d)
			}
		case "C":
			if r.TryDecodeAsNil() {
				x.C = testMammoth2Basic{}
			} else {
				yyv8 := &x.C
				yyv8.CodecDecodeSelf(d)
			}
		case "M":
			if r.TryDecodeAsNil() {
				x.M = nil
			} else {
				yyv9 := &x.M
				yym10 := z.DecBinary()
				_ = yym10
				if false {
				} else {
					h.decMaptestMammoth2BasicTestMammoth2((*map[testMammoth2Basic]TestMammoth2)(yyv9), d)
				}
			}
		case "L":
			if r.TryDecodeAsNil() {
				x.L = nil
			} else {
				yyv11 := &x.L
				yym12 := z.DecBinary()
				_ = yym12
				if false {
				} else {
					h.decSliceTestMammoth2((*[]TestMammoth2)(yyv11), d)
				}
			}
		case "A":
			if r.TryDecodeAsNil() {
				x.A = [4]int64{}
			} else {
				yyv13 := &x.A
				yym14 := z.DecBinary()
				_ = yym14
				if false {
				} else {
					h.decArray4int64((*[4]int64)(yyv13), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *TestMammoth2Wrapper) codecDecodeSelfFromArray(l int, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj15 int
	var yyb15 bool
	var yyhl15 bool = l >= 0
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.V = TestMammoth2{}
	} else {
		yyv16 := &x.V
		yyv16.CodecDecodeSelf(d)
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.T = 0
	} else {
		yyv17 := &x.T
		yyv17.CodecDecodeSelf(d)
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.B = 0
	} else {
		yyv18 := &x.B
		yyv18.CodecDecodeSelf(d)
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.J = 0
	} else {
		yyv19 := &x.J
		yyv19.CodecDecodeSelf(d)
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.C = testMammoth2Basic{}
	} else {
		yyv20 := &x.C
		yyv20.CodecDecodeSelf(d)
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.M = nil
	} else {
		yyv21 := &x.M
		yym22 := z.DecBinary()
		_ = yym22
		if false {
		} else {
			h.decMaptestMammoth2BasicTestMammoth2((*map[testMammoth2Basic]TestMammoth2)(yyv21), d)
		}
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.L = nil
	} else {
		yyv23 := &x.L
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			h.decSliceTestMammoth2((*[]TestMammoth2)(yyv23), d)
		}
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.A = [4]int64{}
	} else {
		yyv25 := &x.A
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else {
			h.decArray4int64((*[4]int64)(yyv25), d)
		}
	}
	for {
		yyj15++
		if yyhl15 {
			yyb15 = yyj15 > l
		} else {
			yyb15 = r.CheckBreak()
		}
		if yyb15 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj15-1, "")
	}
	r.ReadArrayEnd()
}

func (x codecSelfer19781) enctestMammoth2Basic(v *testMammoth2Basic, e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteArrayStart(len(v))
	for _, yyv1 := range v {
		r.WriteArrayElem()
		yym2 := z.EncBinary()
		_ = yym2
		if false {
		} else {
			r.EncodeUint(uint64(yyv1))
		}
	}
	r.WriteArrayEnd()
}

func (x codecSelfer19781) dectestMammoth2Basic(v *testMammoth2Basic, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := v
	yyh1, yyl1 := z.DecSliceHelperStart()
	if yyl1 == 0 {

	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1

		var yyj1 int
		// var yydn1 bool
		for ; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {

			yyh1.ElemContainerState(yyj1)
			// yydn1 = r.TryDecodeAsNil()

			// if indefinite, etc, then expand the slice if necessary
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				z.DecArrayCannotExpand(len(v), yyj1+1)
				yydb1 = true

			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = 0
				} else {
					yyv2 := &yyv1[yyj1]
					yym3 := z.DecBinary()
					_ = yym3
					if false {
					} else {
						*((*uint64)(yyv2)) = uint64(r.DecodeUint(64))
					}
				}

			}

		}

	}
	yyh1.End()

}

func (x codecSelfer19781) encMaptestMammoth2BasicTestMammoth2(v map[testMammoth2Basic]TestMammoth2, e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteMapStart(len(v))
	for yyk1, yyv1 := range v {
		r.WriteMapElemKey()
		yy2 := &yyk1
		yy2.CodecEncodeSelf(e)
		r.WriteMapElemValue()
		yy4 := &yyv1
		yy4.CodecEncodeSelf(e)
	}
	r.WriteMapEnd()
}

func (x codecSelfer19781) decMaptestMammoth2BasicTestMammoth2(v *map[testMammoth2Basic]TestMammoth2, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyl1 := r.ReadMapStart()
	yybh1 := z.DecBasicHandle()
	if yyv1 == nil {
		yyrl1 := z.DecInferLen(yyl1, yybh1.MaxInitLen, 4880)
		yyv1 = make(map[testMammoth2Basic]TestMammoth2, yyrl1)
		*v = yyv1
	}
	var yymk1 testMammoth2Basic
	var yymv1 TestMammoth2
	var yymg1, yymdn1 bool
	if yybh1.MapValueReset {
		yymg1 = true
	}
	if yyl1 != 0 {
		yyhl1 := yyl1 > 0
		for yyj1 := 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {
			r.ReadMapElemKey()
			if r.TryDecodeAsNil() {
				yymk1 = testMammoth2Basic{}
			} else {
				yyv2 := &yymk1
				yyv2.CodecDecodeSelf(d)
			}

			if yymg1 {
				yymv1 = yyv1[yymk1]
			} else {
				yymv1 = TestMammoth2{}
			}
			r.ReadMapElemValue()
			yymdn1 = false
			if r.TryDecodeAsNil() {
				yymdn1 = true
			} else {
				yyv3 := &yymv1
				yyv3.CodecDecodeSelf(d)
			}

			if yymdn1 {
				if yybh1.DeleteOnNilMapValue {
					delete(yyv1, yymk1)
				} else {
					yyv1[yymk1] = TestMammoth2{}
				}
			} else if yyv1 != nil {
				yyv1[yymk1] = yymv1
			}
		}
	} // else len==0: TODO: Should we clear map entries?
	r.ReadMapEnd()
}

func (x codecSelfer19781) encSliceTestMammoth2(v []TestMammoth2, e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteArrayStart(len(v))
	for _, yyv1 := range v {
		r.WriteArrayElem()
		yy2 := &yyv1
		yy2.CodecEncodeSelf(e)
	}
	r.WriteArrayEnd()
}

func (x codecSelfer19781) decSliceTestMammoth2(v *[]TestMammoth2, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []TestMammoth2{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 4848)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]TestMammoth2, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		// var yydn1 bool
		for ; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {
			if yyj1 == 0 && len(yyv1) == 0 {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 4848)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]TestMammoth2, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			// yydn1 = r.TryDecodeAsNil()

			// if indefinite, etc, then expand the slice if necessary
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, TestMammoth2{})
				yyc1 = true

			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = TestMammoth2{}
				} else {
					yyv2 := &yyv1[yyj1]
					yyv2.CodecDecodeSelf(d)
				}

			}

		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]TestMammoth2, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}

}

func (x codecSelfer19781) encArray4int64(v *[4]int64, e *Encoder) {
	var h codecSelfer19781
	z, r := GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteArrayStart(len(v))
	for _, yyv1 := range v {
		r.WriteArrayElem()
		yym2 := z.EncBinary()
		_ = yym2
		if false {
		} else {
			r.EncodeInt(int64(yyv1))
		}
	}
	r.WriteArrayEnd()
}

func (x codecSelfer19781) decArray4int64(v *[4]int64, d *Decoder) {
	var h codecSelfer19781
	z, r := GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := v
	yyh1, yyl1 := z.DecSliceHelperStart()
	if yyl1 == 0 {

	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1

		var yyj1 int
		// var yydn1 bool
		for ; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {

			yyh1.ElemContainerState(yyj1)
			// yydn1 = r.TryDecodeAsNil()

			// if indefinite, etc, then expand the slice if necessary
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				z.DecArrayCannotExpand(len(v), yyj1+1)
				yydb1 = true

			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = 0
				} else {
					yyv2 := &yyv1[yyj1]
					yym3 := z.DecBinary()
					_ = yym3
					if false {
					} else {
						*((*int64)(yyv2)) = int64(r.DecodeInt(64))
					}
				}

			}

		}

	}
	yyh1.End()

}
