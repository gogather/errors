// Copyright 2020. All rights reserved.
// This file is part of gogather project
// I am coding in Tencent
// Created by rainesli on 2020/3/30.

package errors

import (
	"fmt"
	"github.com/go-errors/errors"
)

func ToTrace(err error) (trace string) {
	e, ok := err.(*errors.Error)
	if ok {
		trace = trace + e.ErrorStack()
	} else {
		trace = trace + fmt.Sprintf("%v", err)
	}
	return trace
}

func New(e interface{}) (err error) {
	return errors.New(e)
}
