// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !go1.18

package pqarrow

import (
	"errors"

	"github.com/nidhhoggr/go-arrow/arrow"
	"github.com/nidhhoggr/go-arrow/parquet/file"
)

func writeDictionaryArrow(*arrowWriteContext, file.ColumnChunkWriter, arrow.Array, []int16, []int16, bool) (err error) {
	return errors.New("parquet/pqarrow: go1.18+ required to write arrow dictionary arrays in WriteArrowToColumn")
}
