// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package executors

import (
	"processors"
)

type IExecutor interface {
	Execute() (processors.IProcessor, error)
}
