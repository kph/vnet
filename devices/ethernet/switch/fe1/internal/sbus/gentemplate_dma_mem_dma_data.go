// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=sbus -id dma_data -d VecType=dma_data_vec -d Type=dma_data github.com/platinasystems/go/elib/hw/dma_mem.tmpl]

package sbus

import (
	"github.com/platinasystems/go/elib"
	"github.com/platinasystems/go/elib/hw"

	"reflect"
	"unsafe"
)

type dma_data_vec []dma_data

func fromByteSlice_dma_data(b []byte, l, c uint) (x dma_data_vec) {
	s := uint(unsafe.Sizeof(x[0]))
	if l == 0 {
		l = uint(len(b)) / s
		c = uint(cap(b))
	}
	return *(*dma_data_vec)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  int(l),
		Cap:  int(c / s),
	}))
}

func (x dma_data_vec) toByteSlice() []byte {
	l := uint(len(x))
	l *= uint(unsafe.Sizeof(x[0]))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&x[0])),
		Len:  int(l),
		Cap:  int(l)}))
}

func dma_dataAllocAligned(n, a uint) (x dma_data_vec, id elib.Index) {
	var b []byte
	var c uint
	b, id, _, c = hw.DmaAllocAligned(n*uint(unsafe.Sizeof(x[0])), a)
	x = fromByteSlice_dma_data(b, n, c)
	return
}

func dma_dataAlloc(n uint) (x dma_data_vec, id elib.Index) { return dma_dataAllocAligned(n, 0) }

func dma_dataNew() (x dma_data_vec, id elib.Index) { return dma_dataAlloc(1) }

func (x *dma_data_vec) Free(id elib.Index) {
	hw.DmaFree(id)
	*x = nil
}

func (x *dma_data_vec) Get(id elib.Index) {
	*x = fromByteSlice_dma_data(hw.DmaGetData(id), 0, 0)
}

func (x *dma_data) PhysAddress() uintptr {
	return hw.DmaPhysAddress(uintptr(unsafe.Pointer(x)))
}