/*
Copyright 2015 James Duncan Davidson

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package base64fix

import (
	"encoding/base64"
	"strings"
)

// Encoding wraps the underlying encoding/base64 implementation, allowing base64fix to intercept a decode method, check the payload, and then forward on.
type Encoding struct {
	encoding *base64.Encoding
}

// StdEncoding provides a wrapped base64.StdEncoding struct
var StdEncoding = Encoding{encoding: base64.StdEncoding}

// URLEncoding provides a wrapped base64.URLEncoding
var URLEncoding = Encoding{encoding: base64.URLEncoding}

// Decode decodes the given base64-encoded src using base64.Decode, first padding the input as needed to comply with base64 standards.
func (enc *Encoding) Decode(dst, src []byte) (int, error) {
	if n := len(src) % 4; n != 0 {
		p := make([]byte, 4-n)
		for i := range p {
			p[i] = '='
		}
		src = append(src, p...)
	}
	return enc.encoding.Decode(dst, src)
}

// DecodeString decodes the given base64-encoded string using base64.DecodeString, first padding the input as needed to comply with base64 standards.
func (enc *Encoding) DecodeString(s string) ([]byte, error) {
	// fix up as suggested by David Symonds in https://github.com/golang/go/issues/4237#issuecomment-66071224
	if n := len(s) % 4; n != 0 {
		s += strings.Repeat("=", 4-n)
	}
	return enc.encoding.DecodeString(s)
}
