/*
 * Copyright (c) 2022-2023 Zander Schwid & Co. LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package base62_test

import (
	"bytes"
	"github.com/codeallergy/base62"
	"testing"
)

var (
	raw5k       = bytes.Repeat([]byte{0xff}, 5000)
	raw100k     = bytes.Repeat([]byte{0xff}, 100*1000)
	encoded5k   = base62.StdEncoding.EncodeToString(raw5k)
	encoded100k = base62.StdEncoding.EncodeToString(raw100k)
)

func BenchmarkBase62Encode_5K(b *testing.B) {
	b.SetBytes(int64(len(raw5k)))
	for i := 0; i < b.N; i++ {
		base62.StdEncoding.EncodeToString(raw5k)
	}
}

func BenchmarkBase62Encode_100K(b *testing.B) {
	b.SetBytes(int64(len(raw100k)))
	for i := 0; i < b.N; i++ {
		base62.StdEncoding.EncodeToString(raw100k)
	}
}

func BenchmarkBase62Decode_5K(b *testing.B) {
	b.SetBytes(int64(len(encoded5k)))
	for i := 0; i < b.N; i++ {
		base62.StdEncoding.DecodeString(encoded5k)
	}
}

func BenchmarkBase62Decode_100K(b *testing.B) {
	b.SetBytes(int64(len(encoded100k)))
	for i := 0; i < b.N; i++ {
		base62.StdEncoding.DecodeString(encoded100k)
	}
}
