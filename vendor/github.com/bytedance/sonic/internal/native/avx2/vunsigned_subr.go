// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__vunsigned = 0
)

const (
    _stack__vunsigned = 24
)

const (
    _size__vunsigned = 356
)

var (
    _pcsp__vunsigned = [][2]uint32{
        {0x1, 0},
        {0x6, 8},
        {0x7, 16},
        {0x4b, 24},
        {0x4d, 16},
        {0x4e, 8},
        {0x4f, 0},
        {0x5a, 24},
        {0x5c, 16},
        {0x5d, 8},
        {0x5e, 0},
        {0x75, 24},
        {0x77, 16},
        {0x78, 8},
        {0x79, 0},
        {0x11a, 24},
        {0x11c, 16},
        {0x11d, 8},
        {0x11e, 0},
        {0x151, 24},
        {0x153, 16},
        {0x154, 8},
        {0x155, 0},
        {0x15d, 24},
        {0x15f, 16},
        {0x160, 8},
        {0x164, 0},
    }
)

var _cfunc_vunsigned = []loader.CFunc{
    {"_vunsigned_entry", 0,  _entry__vunsigned, 0, nil},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}