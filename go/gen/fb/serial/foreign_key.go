// Copyright 2022 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serial

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ForeignKeyReferentialAction byte

const (
	ForeignKeyReferentialActionDefaultAction ForeignKeyReferentialAction = 0
	ForeignKeyReferentialActionCascade       ForeignKeyReferentialAction = 1
	ForeignKeyReferentialActionNoAction      ForeignKeyReferentialAction = 2
	ForeignKeyReferentialActionRestrict      ForeignKeyReferentialAction = 3
	ForeignKeyReferentialActionSetNull       ForeignKeyReferentialAction = 4
)

var EnumNamesForeignKeyReferentialAction = map[ForeignKeyReferentialAction]string{
	ForeignKeyReferentialActionDefaultAction: "DefaultAction",
	ForeignKeyReferentialActionCascade:       "Cascade",
	ForeignKeyReferentialActionNoAction:      "NoAction",
	ForeignKeyReferentialActionRestrict:      "Restrict",
	ForeignKeyReferentialActionSetNull:       "SetNull",
}

var EnumValuesForeignKeyReferentialAction = map[string]ForeignKeyReferentialAction{
	"DefaultAction": ForeignKeyReferentialActionDefaultAction,
	"Cascade":       ForeignKeyReferentialActionCascade,
	"NoAction":      ForeignKeyReferentialActionNoAction,
	"Restrict":      ForeignKeyReferentialActionRestrict,
	"SetNull":       ForeignKeyReferentialActionSetNull,
}

func (v ForeignKeyReferentialAction) String() string {
	if s, ok := EnumNamesForeignKeyReferentialAction[v]; ok {
		return s
	}
	return "ForeignKeyReferentialAction(" + strconv.FormatInt(int64(v), 10) + ")"
}

type ForeignKeyCollection struct {
	_tab flatbuffers.Table
}

func GetRootAsForeignKeyCollection(buf []byte, offset flatbuffers.UOffsetT) *ForeignKeyCollection {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ForeignKeyCollection{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsForeignKeyCollection(buf []byte, offset flatbuffers.UOffsetT) *ForeignKeyCollection {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ForeignKeyCollection{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ForeignKeyCollection) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ForeignKeyCollection) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ForeignKeyCollection) ForeignKeys(obj *ForeignKey, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ForeignKeyCollection) ForeignKeysLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ForeignKeyCollectionStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ForeignKeyCollectionAddForeignKeys(builder *flatbuffers.Builder, foreignKeys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(foreignKeys), 0)
}
func ForeignKeyCollectionStartForeignKeysVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ForeignKeyCollectionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type ForeignKey struct {
	_tab flatbuffers.Table
}

func GetRootAsForeignKey(buf []byte, offset flatbuffers.UOffsetT) *ForeignKey {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ForeignKey{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsForeignKey(buf []byte, offset flatbuffers.UOffsetT) *ForeignKey {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ForeignKey{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ForeignKey) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ForeignKey) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ForeignKey) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ForeignKey) ChildTableName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ForeignKey) ChildTableIndex() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ForeignKey) ChildTableColumns(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *ForeignKey) ChildTableColumnsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ForeignKey) MutateChildTableColumns(j int, n uint64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *ForeignKey) ParentTableName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ForeignKey) ParentTableIndex() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ForeignKey) ParentTableColumns(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *ForeignKey) ParentTableColumnsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ForeignKey) MutateParentTableColumns(j int, n uint64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *ForeignKey) OnUpdate() ForeignKeyReferentialAction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return ForeignKeyReferentialAction(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ForeignKey) MutateOnUpdate(n ForeignKeyReferentialAction) bool {
	return rcv._tab.MutateByteSlot(18, byte(n))
}

func (rcv *ForeignKey) OnDelete() ForeignKeyReferentialAction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return ForeignKeyReferentialAction(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ForeignKey) MutateOnDelete(n ForeignKeyReferentialAction) bool {
	return rcv._tab.MutateByteSlot(20, byte(n))
}

func (rcv *ForeignKey) UnresolvedChildColumns(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *ForeignKey) UnresolvedChildColumnsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ForeignKey) UnresolvedParentColumns(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *ForeignKey) UnresolvedParentColumnsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ForeignKeyStart(builder *flatbuffers.Builder) {
	builder.StartObject(11)
}
func ForeignKeyAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ForeignKeyAddChildTableName(builder *flatbuffers.Builder, childTableName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(childTableName), 0)
}
func ForeignKeyAddChildTableIndex(builder *flatbuffers.Builder, childTableIndex flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(childTableIndex), 0)
}
func ForeignKeyAddChildTableColumns(builder *flatbuffers.Builder, childTableColumns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(childTableColumns), 0)
}
func ForeignKeyStartChildTableColumnsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func ForeignKeyAddParentTableName(builder *flatbuffers.Builder, parentTableName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(parentTableName), 0)
}
func ForeignKeyAddParentTableIndex(builder *flatbuffers.Builder, parentTableIndex flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(parentTableIndex), 0)
}
func ForeignKeyAddParentTableColumns(builder *flatbuffers.Builder, parentTableColumns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(parentTableColumns), 0)
}
func ForeignKeyStartParentTableColumnsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func ForeignKeyAddOnUpdate(builder *flatbuffers.Builder, onUpdate ForeignKeyReferentialAction) {
	builder.PrependByteSlot(7, byte(onUpdate), 0)
}
func ForeignKeyAddOnDelete(builder *flatbuffers.Builder, onDelete ForeignKeyReferentialAction) {
	builder.PrependByteSlot(8, byte(onDelete), 0)
}
func ForeignKeyAddUnresolvedChildColumns(builder *flatbuffers.Builder, unresolvedChildColumns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(unresolvedChildColumns), 0)
}
func ForeignKeyStartUnresolvedChildColumnsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ForeignKeyAddUnresolvedParentColumns(builder *flatbuffers.Builder, unresolvedParentColumns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(unresolvedParentColumns), 0)
}
func ForeignKeyStartUnresolvedParentColumnsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ForeignKeyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}