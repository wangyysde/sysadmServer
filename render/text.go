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
	"fmt"
	"net/http"

	"github.com/wangyysde/sysadmServer/internal/bytesconv"
)

// String contains the given interface object slice and its format.
type String struct {
	Format string
	Data   []interface{}
}

var plainContentType = []string{"text/plain; charset=utf-8"}

// Render (String) writes data with custom ContentType.
func (r String) Render(w http.ResponseWriter) error {
	return WriteString(w, r.Format, r.Data)
}

// WriteContentType (String) writes Plain ContentType.
func (r String) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, plainContentType)
}

// WriteString writes data according to its format and write custom ContentType.
func WriteString(w http.ResponseWriter, format string, data []interface{}) (err error) {
	writeContentType(w, plainContentType)
	if len(data) > 0 {
		_, err = fmt.Fprintf(w, format, data...)
		return
	}
	_, err = w.Write(bytesconv.StringToBytes(format))
	return
}
