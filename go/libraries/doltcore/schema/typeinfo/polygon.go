// Copyright 2020 Dolthub, Inc.
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

package typeinfo

import (
	"context"
	"fmt"

	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/dolt/go/store/geometry"
	"github.com/dolthub/dolt/go/store/types"
)

// This is a dolt implementation of the MySQL type Point, thus most of the functionality
// within is directly reliant on the go-mysql-server implementation.
type polygonType struct {
	sqlPolygonType sql.PolygonType
}

var _ TypeInfo = (*polygonType)(nil)

var PolygonType = &polygonType{sql.PolygonType{}}

// ConvertTypesPolygonToSQLPolygon basically makes a deep copy of sql.Linestring
func ConvertTypesPolygonToSQLPolygon(p types.Polygon) sql.Polygon {
	lines := make([]sql.Linestring, len(p.Lines))
	for i, l := range p.Lines {
		lines[i] = ConvertTypesLinestringToSQLLinestring(l)
	}
	return sql.Polygon{SRID: p.SRID, Lines: lines}
}

// ConvertNomsValueToValue implements TypeInfo interface.
func (ti *polygonType) ConvertNomsValueToValue(v types.Value) (interface{}, error) {
	// Expect a types.Polygon, return a sql.Polygon
	if val, ok := v.(types.Polygon); ok {
		return ConvertTypesPolygonToSQLPolygon(val), nil
	}
	// Check for null
	if _, ok := v.(types.Null); ok || v == nil {
		return nil, nil
	}
	return nil, fmt.Errorf(`"%v" cannot convert NomsKind "%v" to a value`, ti.String(), v.Kind())
}

// ReadFrom reads a go value from a noms types.CodecReader directly
func (ti *polygonType) ReadFrom(nbf *types.NomsBinFormat, reader types.CodecReader) (interface{}, error) {
	k := reader.ReadKind()
	switch k {
	case types.PolygonKind:
		p, err := reader.ReadPolygon()
		if err != nil {
			return nil, err
		}
		return ti.ConvertNomsValueToValue(p)
	case types.NullKind:
		return nil, nil
	}

	return nil, fmt.Errorf(`"%v" cannot convert NomsKind "%v" to a value`, ti.String(), k)
}

func ConvertSQLPolygonToTypesPolygon(p sql.Polygon) types.Polygon {
	lines := make([]types.Linestring, len(p.Lines))
	for i, l := range p.Lines {
		lines[i] = ConvertSQLLinestringToTypesLinestring(l)
	}
	return types.Polygon{SRID: p.SRID, Lines: lines}
}

// ConvertValueToNomsValue implements TypeInfo interface.
func (ti *polygonType) ConvertValueToNomsValue(ctx context.Context, vrw types.ValueReadWriter, v interface{}) (types.Value, error) {
	// Check for null
	if v == nil {
		return types.NullValue, nil
	}

	// Convert to sql.PolygonType
	poly, err := ti.sqlPolygonType.Convert(v)
	if err != nil {
		return nil, err
	}

	return ConvertSQLPolygonToTypesPolygon(poly.(sql.Polygon)), nil
}

// Equals implements TypeInfo interface.
func (ti *polygonType) Equals(other TypeInfo) bool {
	if other == nil {
		return false
	}
	_, ok := other.(*polygonType)
	return ok
}

// FormatValue implements TypeInfo interface.
func (ti *polygonType) FormatValue(v types.Value) (*string, error) {
	if val, ok := v.(types.Polygon); ok {
		size := geometry.EWKBHeaderSize + types.LengthSize
		for _, l := range val.Lines {
			size += types.LengthSize + geometry.PointSize*len(l.Points)
		}
		buf := make([]byte, size)
		types.WriteEWKBHeader(val, buf[:geometry.EWKBHeaderSize])
		types.WriteEWKBPolyData(val, buf[geometry.EWKBHeaderSize:])
		resStr := string(buf)
		return &resStr, nil
	}
	if _, ok := v.(types.Null); ok || v == nil {
		return nil, nil
	}

	return nil, fmt.Errorf(`"%v" has unexpectedly encountered a value of type "%T" from embedded type`, ti.String(), v.Kind())
}

// GetTypeIdentifier implements TypeInfo interface.
func (ti *polygonType) GetTypeIdentifier() Identifier {
	return PolygonTypeIdentifier
}

// GetTypeParams implements TypeInfo interface.
func (ti *polygonType) GetTypeParams() map[string]string {
	return map[string]string{}
}

// IsValid implements TypeInfo interface.
func (ti *polygonType) IsValid(v types.Value) bool {
	if _, ok := v.(types.Polygon); ok {
		return true
	}
	if _, ok := v.(types.Null); ok || v == nil {
		return true
	}
	return false
}

// NomsKind implements TypeInfo interface.
func (ti *polygonType) NomsKind() types.NomsKind {
	return types.PolygonKind
}

// Promote implements TypeInfo interface.
func (ti *polygonType) Promote() TypeInfo {
	return &polygonType{ti.sqlPolygonType.Promote().(sql.PolygonType)}
}

// String implements TypeInfo interface.
func (ti *polygonType) String() string {
	return "Polygon"
}

// ToSqlType implements TypeInfo interface.
func (ti *polygonType) ToSqlType() sql.Type {
	return ti.sqlPolygonType
}

// polygonTypeConverter is an internal function for GetTypeConverter that handles the specific type as the source TypeInfo.
func polygonTypeConverter(ctx context.Context, src *polygonType, destTi TypeInfo) (tc TypeConverter, needsConversion bool, err error) {
	switch dest := destTi.(type) {
	case *bitType:
		return func(ctx context.Context, vrw types.ValueReadWriter, v types.Value) (types.Value, error) {
			return types.Uint(0), nil
		}, true, nil
	case *blobStringType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *boolType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *datetimeType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *decimalType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *enumType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *floatType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *inlineBlobType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *intType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *jsonType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *linestringType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *pointType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *polygonType:
		return identityTypeConverter, false, nil
	case *setType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *timeType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *uintType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *uuidType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *varBinaryType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *varStringType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	case *yearType:
		return wrapConvertValueToNomsValue(dest.ConvertValueToNomsValue)
	default:
		return nil, false, UnhandledTypeConversion.New(src.String(), destTi.String())
	}
}