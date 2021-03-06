// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package datatypes

import (
	"fmt"
	"io"
	"reflect"

	"datavalues"

	"base/binary"
	"base/errors"
)

const (
	DataTypeFloat32Name = "Float32"
)

type Float32DataType struct {
	DataTypeBase
}

func NewFloat32DataType() IDataType {
	return &Float32DataType{}
}

func (datatype *Float32DataType) Type() reflect.Type {
	return reflect.ValueOf(int64(0)).Type()
}

func (datatype *Float32DataType) Name() string {
	return DataTypeFloat32Name
}

func (datatype *Float32DataType) Serialize(writer *binary.Writer, v *datavalues.Value) error {
	if err := writer.Float32(float32(v.GetFloat())); err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func (datatype *Float32DataType) SerializeText(writer io.Writer, v *datavalues.Value) error {
	if _, err := writer.Write([]byte(fmt.Sprintf("%v", float32(v.GetFloat())))); err != nil {
		return errors.Wrap(err)
	}
	return nil
}
