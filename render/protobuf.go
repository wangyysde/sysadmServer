/**
* SYSADM Server
* @Author  Wayne Wang <net_use@bzhy.com>
* @Copyright Bzhy Network
* @HomePage http://www.sysadm.cn
* @Version 0.1.0
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
* @License GNU Lesser General Public License  https://www.sysadm.cn/lgpl.html
* @Modified Apr 20 2021
**/

package render

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// ProtoBuf contains the given interface object.
type ProtoBuf struct {
	Data interface{}
}

var protobufContentType = []string{"application/x-protobuf"}

// Render (ProtoBuf) marshals the given interface object and writes data with custom ContentType.
func (r ProtoBuf) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := proto.Marshal(r.Data.(proto.Message))
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

// WriteContentType (ProtoBuf) writes ProtoBuf ContentType.
func (r ProtoBuf) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, protobufContentType)
}
