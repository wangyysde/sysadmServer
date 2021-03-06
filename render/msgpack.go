// sysadmServer
// @Author  Wayne Wang <net_use@bzhy.com>
// @Copyright Bzhy Network
// @HomePage http://www.sysadm.cn
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and 
// limitations under the License.
// @License GNU Lesser General Public License  https://www.sysadm.cn/lgpl.html
//  @Modified on Jul 16 2021

//go:build !nomsgpack
// +build !nomsgpack

package render

import (
	"net/http"

	"github.com/ugorji/go/codec"
)

// Check interface implemented here to support go build tag nomsgpack.
var (
	_ Render = MsgPack{}
)

// MsgPack contains the given interface object.
type MsgPack struct {
	Data interface{}
}

var msgpackContentType = []string{"application/msgpack; charset=utf-8"}

// WriteContentType (MsgPack) writes MsgPack ContentType.
func (r MsgPack) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, msgpackContentType)
}

// Render (MsgPack) encodes the given interface object and writes data with custom ContentType.
func (r MsgPack) Render(w http.ResponseWriter) error {
	return WriteMsgPack(w, r.Data)
}

// WriteMsgPack writes MsgPack ContentType and encodes the given interface object.
func WriteMsgPack(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, msgpackContentType)
	var mh codec.MsgpackHandle
	return codec.NewEncoder(w, &mh).Encode(obj)
}
